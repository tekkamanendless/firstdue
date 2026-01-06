package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/kirsle/configdir"
	"github.com/tekkamanendless/firstdue"
)

type Args struct {
	Debug bool `arg:"--debug,env:DEBUG" help:"Enable debug mode"`

	Config *ConfigCommand `arg:"subcommand" help:"Configuration commands"`
	API    *APICommand    `arg:"subcommand" help:"API commands"`
}

type ConfigCommand struct {
	Activate  *ConfigActivateCommand  `arg:"subcommand" help:"Activate a configuration"`
	Configure *ConfigConfigureCommand `arg:"subcommand" help:"Configure a new configuration"`
	List      *ConfigListCommand      `arg:"subcommand" help:"List all configurations"`
}

type ConfigActivateCommand struct {
	Name string `arg:"positional,required" help:"Name of the configuration"`
}

type ConfigConfigureCommand struct {
	Name     string `arg:"positional,required" help:"Name of the configuration"`
	BaseURL  string `arg:"--base-url" help:"Base URL for the configuration"`
	Username string `arg:"--username" help:"Username for the configuration"`
	Password string `arg:"--password" help:"Password for the configuration"`
}

type ConfigListCommand struct{}

type APICommand struct {
	Endpoint string `arg:"--endpoint,required" help:"API endpoint to call"`
	Method   string `arg:"--method" default:"GET" help:"HTTP method to use"`
	Body     string `arg:"--body" help:"Request body for POST/PUT methods"`
}

type Config struct {
	DefaultAccount string             `json:"defaultAccount"`
	AccountMap     map[string]Account `json:"accountMap"`
}

type Account struct {
	BaseURL  string `json:"baseURL"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	ctx := context.Background()

	var args Args
	argsParser, err := arg.NewParser(arg.Config{}, &args)
	if err != nil {
		panic(err)
	}

	err = argsParser.Parse(os.Args[1:])
	switch {
	case err == arg.ErrHelp: // found "--help" on command line
		argsParser.WriteHelpForSubcommand(os.Stdout, argsParser.SubcommandNames()...)
		os.Exit(0)
	case err != nil:
		fmt.Printf("error: %v\n", err)
		argsParser.WriteUsageForSubcommand(os.Stdout, argsParser.SubcommandNames()...)
		os.Exit(1)
	}

	if args.Debug {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	}

	configDirectory := configdir.LocalConfig("firstdue")
	err = configdir.MakePath(configDirectory)
	if err != nil {
		fmt.Printf("error creating config directory: %v\n", err)
		os.Exit(1)
	}
	var config Config
	{
		configFilePath := configDirectory + string(os.PathSeparator) + "config.json"
		contents, err := os.ReadFile(configFilePath)
		if err != nil {
			if !os.IsNotExist(err) {
				fmt.Printf("error reading config file: %v\n", err)
				os.Exit(1)
			}
		} else {
			err = json.Unmarshal(contents, &config)
			if err != nil {
				fmt.Printf("error parsing config file: %v\n", err)
			}
		}

		if config.AccountMap == nil {
			config.AccountMap = make(map[string]Account)
		}
	}

	switch {
	case args.Config != nil:
		switch {
		case args.Config.Activate != nil:
			_, exists := config.AccountMap[args.Config.Activate.Name]
			if !exists {
				fmt.Printf("configuration '%s' does not exist\n", args.Config.Activate.Name)
				os.Exit(1)
			}
			config.DefaultAccount = args.Config.Activate.Name
			configFilePath := configDirectory + string(os.PathSeparator) + "config.json"
			contents, err := json.MarshalIndent(config, "", "  ")
			if err != nil {
				fmt.Printf("error serializing config file: %v\n", err)
				os.Exit(1)
			}
			err = os.WriteFile(configFilePath, contents, 0600)
			if err != nil {
				fmt.Printf("error writing config file: %v\n", err)
				os.Exit(1)
			}
		case args.Config.Configure != nil:
			config.AccountMap[args.Config.Configure.Name] = Account{
				BaseURL:  args.Config.Configure.BaseURL,
				Username: args.Config.Configure.Username,
				Password: args.Config.Configure.Password,
			}
			configFilePath := configDirectory + string(os.PathSeparator) + "config.json"
			contents, err := json.MarshalIndent(config, "", "  ")
			if err != nil {
				fmt.Printf("error serializing config file: %v\n", err)
				os.Exit(1)
			}
			err = os.WriteFile(configFilePath, contents, 0600)
			if err != nil {
				fmt.Printf("error writing config file: %v\n", err)
				os.Exit(1)
			}
		case args.Config.List != nil:
			for name := range config.AccountMap {
				fmt.Println(name)
			}
		default:
			argsParser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
	case args.API != nil:
		slog.InfoContext(ctx, "For a list of API endpoints, see https://support.firstduesizeup.com/portal/en/kb/articles/first-due-rest-api-documentation")

		var options []firstdue.ClientOption
		if args.Debug {
			options = append(options, firstdue.WithDebug(true))
		}
		if config.AccountMap[config.DefaultAccount].BaseURL != "" {
			options = append(options, firstdue.WithBaseURL(config.AccountMap[config.DefaultAccount].BaseURL))
		}
		client := firstdue.NewClient(options...)
		err = client.Authenticate(ctx, config.AccountMap[config.DefaultAccount].Username, config.AccountMap[config.DefaultAccount].Password)
		if err != nil {
			fmt.Printf("error authenticating: %v\n", err)
			os.Exit(1)
		}

		var jsonRequest json.RawMessage
		if args.API.Body != "" {
			jsonRequest = json.RawMessage(args.API.Body)
		}
		var jsonResponse json.RawMessage
		err = client.Raw(ctx, args.API.Method, args.API.Endpoint, jsonRequest, &jsonResponse)
		if err != nil {
			fmt.Printf("error making API request: %v\n", err)
			os.Exit(1)
		}

		var prettyJSON []byte
		prettyJSON, err = json.MarshalIndent(jsonResponse, "", "  ")
		if err != nil {
			fmt.Printf("error formatting JSON response: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(prettyJSON))
	default:
		argsParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}
