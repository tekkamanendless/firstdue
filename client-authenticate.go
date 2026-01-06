package firstdue

import (
	"context"
	"fmt"
	"net/http"
)

type PostAuthTokenRequest struct {
	GrantType string `json:"grant_type"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type PostAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (c *Client) Authenticate(ctx context.Context, username string, password string) error {
	input := PostAuthTokenRequest{
		GrantType: "client_credentials",
		Email:     username,
		Password:  password,
	}
	var output PostAuthTokenResponse
	err := c.Raw(ctx, http.MethodPost, "/v1/auth/token", input, &output)
	if err != nil {
		return fmt.Errorf("authenticate: %w", err)
	}
	c.config.Token = output.AccessToken
	return nil
}
