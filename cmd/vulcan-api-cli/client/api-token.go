// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": api-token Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateAPITokenPath computes a request path to the create action of api-token.
func CreateAPITokenPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/api/v1/users/%s/token", param0)
}

// Generate an API token for an user.
func (c *Client) CreateAPIToken(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCreateAPITokenRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateAPITokenRequest create the request corresponding to the create action endpoint of the api-token resource.
func (c *Client) NewCreateAPITokenRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
