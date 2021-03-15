// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": policy-settings Resource Client
//
// Command:
// $ main

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreatePolicySettingsPath computes a request path to the create action of policy-settings.
func CreatePolicySettingsPath(teamID string, policyID string) string {
	param0 := teamID
	param1 := policyID

	return fmt.Sprintf("/api/v1/teams/%s/policies/%s/settings", param0, param1)
}

// Create a new policy setting.
func (c *Client) CreatePolicySettings(ctx context.Context, path string, payload *PolicySettingPayload) (*http.Response, error) {
	req, err := c.NewCreatePolicySettingsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreatePolicySettingsRequest create the request corresponding to the create action endpoint of the policy-settings resource.
func (c *Client) NewCreatePolicySettingsRequest(ctx context.Context, path string, payload *PolicySettingPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// DeletePolicySettingsPath computes a request path to the delete action of policy-settings.
func DeletePolicySettingsPath(teamID string, policyID string, settingsID string) string {
	param0 := teamID
	param1 := policyID
	param2 := settingsID

	return fmt.Sprintf("/api/v1/teams/%s/policies/%s/settings/%s", param0, param1, param2)
}

// Delete a policy setting.
func (c *Client) DeletePolicySettings(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeletePolicySettingsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeletePolicySettingsRequest create the request corresponding to the delete action endpoint of the policy-settings resource.
func (c *Client) NewDeletePolicySettingsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "DELETE", u.String(), nil)
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

// ListPolicySettingsPath computes a request path to the list action of policy-settings.
func ListPolicySettingsPath(teamID string, policyID string) string {
	param0 := teamID
	param1 := policyID

	return fmt.Sprintf("/api/v1/teams/%s/policies/%s/settings", param0, param1)
}

// List settings for a policy.
func (c *Client) ListPolicySettings(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListPolicySettingsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListPolicySettingsRequest create the request corresponding to the list action endpoint of the policy-settings resource.
func (c *Client) NewListPolicySettingsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
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

// ShowPolicySettingsPath computes a request path to the show action of policy-settings.
func ShowPolicySettingsPath(teamID string, policyID string, settingsID string) string {
	param0 := teamID
	param1 := policyID
	param2 := settingsID

	return fmt.Sprintf("/api/v1/teams/%s/policies/%s/settings/%s", param0, param1, param2)
}

// Describe a policy setting.
func (c *Client) ShowPolicySettings(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowPolicySettingsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowPolicySettingsRequest create the request corresponding to the show action endpoint of the policy-settings resource.
func (c *Client) NewShowPolicySettingsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
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

// UpdatePolicySettingsPath computes a request path to the update action of policy-settings.
func UpdatePolicySettingsPath(teamID string, policyID string, settingsID string) string {
	param0 := teamID
	param1 := policyID
	param2 := settingsID

	return fmt.Sprintf("/api/v1/teams/%s/policies/%s/settings/%s", param0, param1, param2)
}

// Update a policy setting.
func (c *Client) UpdatePolicySettings(ctx context.Context, path string, payload *PolicySettingUploadPayload) (*http.Response, error) {
	req, err := c.NewUpdatePolicySettingsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdatePolicySettingsRequest create the request corresponding to the update action endpoint of the policy-settings resource.
func (c *Client) NewUpdatePolicySettingsRequest(ctx context.Context, path string, payload *PolicySettingUploadPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequestWithContext(ctx, "PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.BearerSigner != nil {
		if err := c.BearerSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
