package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/tekkamanendless/firstdue"
)

func main() {
	ctx := context.Background()

	var username string
	var password string
	var token string
	var debug bool
	flag.StringVar(&username, "username", "", "FirstDue API username (email)")
	flag.StringVar(&password, "password", "", "FirstDue API password")
	flag.StringVar(&token, "token", "", "FirstDue API token")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")

	flag.Parse()

	client := firstdue.NewClient()
	if debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		client.Debug = true
	}
	if token != "" {
		client.Token = token
	} else if username != "" && password != "" {
		if err := client.Authenticate(ctx, username, password); err != nil {
			panic(err)
		}
	} else {
		panic("must provide either token or username and password")
	}

	fmt.Printf("Authenticated.\n")
	fmt.Printf("Token: %s\n", client.Token)

	fmt.Printf("Getting logs settings...\n")
	logSettings, err := client.GetLogsSettings(ctx, firstdue.GetLogsSettingsRequest{})
	if err != nil {
		fmt.Printf("GetLogsSettings error: %v\n", err)
	} else {
		fmt.Printf("Log Settings: %+v\n", logSettings)
	}

	fmt.Printf("Getting dispatches...\n")
	dispatches, err := client.GetDispatches(ctx, firstdue.GetDispatchesRequest{})
	if err != nil {
		fmt.Printf("GetDispatches error: %v\n", err)
	} else {
		fmt.Printf("Dispatches: %+v\n", dispatches)
	}
}
