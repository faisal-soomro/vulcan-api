// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": findings Resource Client
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
	"strconv"
)

// FindFindingFindingsPath computes a request path to the Find finding action of findings.
func FindFindingFindingsPath(teamID string, findingID string) string {
	param0 := teamID
	param1 := findingID

	return fmt.Sprintf("/api/v1/teams/%s/findings/%s", param0, param1)
}

// Find a finding.
func (c *Client) FindFindingFindings(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFindFindingFindingsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFindFindingFindingsRequest create the request corresponding to the Find finding action endpoint of the findings resource.
func (c *Client) NewFindFindingFindingsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// FindFindingsFromAIssueFindingsPath computes a request path to the Find findings from a Issue action of findings.
func FindFindingsFromAIssueFindingsPath(teamID string, issueID string) string {
	param0 := teamID
	param1 := issueID

	return fmt.Sprintf("/api/v1/teams/%s/findings/issues/%s", param0, param1)
}

// Find all findings from a team and issue.
func (c *Client) FindFindingsFromAIssueFindings(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, maxScore *float64, minDate *string, minScore *float64, page *float64, size *float64, sortBy *string, status *string) (*http.Response, error) {
	req, err := c.NewFindFindingsFromAIssueFindingsRequest(ctx, path, atDate, identifiers, labels, maxDate, maxScore, minDate, minScore, page, size, sortBy, status)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFindFindingsFromAIssueFindingsRequest create the request corresponding to the Find findings from a Issue action endpoint of the findings resource.
func (c *Client) NewFindFindingsFromAIssueFindingsRequest(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, maxScore *float64, minDate *string, minScore *float64, page *float64, size *float64, sortBy *string, status *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if atDate != nil {
		values.Set("atDate", *atDate)
	}
	if identifiers != nil {
		values.Set("identifiers", *identifiers)
	}
	if labels != nil {
		values.Set("labels", *labels)
	}
	if maxDate != nil {
		values.Set("maxDate", *maxDate)
	}
	if maxScore != nil {
		tmp110 := strconv.FormatFloat(*maxScore, 'f', -1, 64)
		values.Set("maxScore", tmp110)
	}
	if minDate != nil {
		values.Set("minDate", *minDate)
	}
	if minScore != nil {
		tmp111 := strconv.FormatFloat(*minScore, 'f', -1, 64)
		values.Set("minScore", tmp111)
	}
	if page != nil {
		tmp112 := strconv.FormatFloat(*page, 'f', -1, 64)
		values.Set("page", tmp112)
	}
	if size != nil {
		tmp113 := strconv.FormatFloat(*size, 'f', -1, 64)
		values.Set("size", tmp113)
	}
	if sortBy != nil {
		values.Set("sortBy", *sortBy)
	}
	if status != nil {
		values.Set("status", *status)
	}
	u.RawQuery = values.Encode()
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

// FindFindingsFromATargetFindingsPath computes a request path to the Find findings from a Target action of findings.
func FindFindingsFromATargetFindingsPath(teamID string, targetID string) string {
	param0 := teamID
	param1 := targetID

	return fmt.Sprintf("/api/v1/teams/%s/findings/targets/%s", param0, param1)
}

// Find all findings from a team and target.
func (c *Client) FindFindingsFromATargetFindings(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, maxScore *float64, minDate *string, minScore *float64, page *float64, size *float64, sortBy *string, status *string) (*http.Response, error) {
	req, err := c.NewFindFindingsFromATargetFindingsRequest(ctx, path, atDate, identifiers, labels, maxDate, maxScore, minDate, minScore, page, size, sortBy, status)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFindFindingsFromATargetFindingsRequest create the request corresponding to the Find findings from a Target action endpoint of the findings resource.
func (c *Client) NewFindFindingsFromATargetFindingsRequest(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, maxScore *float64, minDate *string, minScore *float64, page *float64, size *float64, sortBy *string, status *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if atDate != nil {
		values.Set("atDate", *atDate)
	}
	if identifiers != nil {
		values.Set("identifiers", *identifiers)
	}
	if labels != nil {
		values.Set("labels", *labels)
	}
	if maxDate != nil {
		values.Set("maxDate", *maxDate)
	}
	if maxScore != nil {
		tmp114 := strconv.FormatFloat(*maxScore, 'f', -1, 64)
		values.Set("maxScore", tmp114)
	}
	if minDate != nil {
		values.Set("minDate", *minDate)
	}
	if minScore != nil {
		tmp115 := strconv.FormatFloat(*minScore, 'f', -1, 64)
		values.Set("minScore", tmp115)
	}
	if page != nil {
		tmp116 := strconv.FormatFloat(*page, 'f', -1, 64)
		values.Set("page", tmp116)
	}
	if size != nil {
		tmp117 := strconv.FormatFloat(*size, 'f', -1, 64)
		values.Set("size", tmp117)
	}
	if sortBy != nil {
		values.Set("sortBy", *sortBy)
	}
	if status != nil {
		values.Set("status", *status)
	}
	u.RawQuery = values.Encode()
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

// ListFindingOverwritesFindingsPath computes a request path to the List Finding Overwrites action of findings.
func ListFindingOverwritesFindingsPath(teamID string, findingID string) string {
	param0 := teamID
	param1 := findingID

	return fmt.Sprintf("/api/v1/teams/%s/findings/%s/overwrites", param0, param1)
}

// List Finding Overwrites.
func (c *Client) ListFindingOverwritesFindings(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListFindingOverwritesFindingsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFindingOverwritesFindingsRequest create the request corresponding to the List Finding Overwrites action endpoint of the findings resource.
func (c *Client) NewListFindingOverwritesFindingsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListFindingsLabelsFindingsPath computes a request path to the List findings labels action of findings.
func ListFindingsLabelsFindingsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/findings/labels", param0)
}

// List all findings labels.
func (c *Client) ListFindingsLabelsFindings(ctx context.Context, path string, atDate *string, identifiers *string, maxDate *string, minDate *string, status *string) (*http.Response, error) {
	req, err := c.NewListFindingsLabelsFindingsRequest(ctx, path, atDate, identifiers, maxDate, minDate, status)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFindingsLabelsFindingsRequest create the request corresponding to the List findings labels action endpoint of the findings resource.
func (c *Client) NewListFindingsLabelsFindingsRequest(ctx context.Context, path string, atDate *string, identifiers *string, maxDate *string, minDate *string, status *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if atDate != nil {
		values.Set("atDate", *atDate)
	}
	if identifiers != nil {
		values.Set("identifiers", *identifiers)
	}
	if maxDate != nil {
		values.Set("maxDate", *maxDate)
	}
	if minDate != nil {
		values.Set("minDate", *minDate)
	}
	if status != nil {
		values.Set("status", *status)
	}
	u.RawQuery = values.Encode()
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

// SubmitAFindingOverwriteFindingsPath computes a request path to the Submit a Finding Overwrite action of findings.
func SubmitAFindingOverwriteFindingsPath(teamID string, findingID string) string {
	param0 := teamID
	param1 := findingID

	return fmt.Sprintf("/api/v1/teams/%s/findings/%s/overwrites", param0, param1)
}

// Overwrite data for a specific finding.
func (c *Client) SubmitAFindingOverwriteFindings(ctx context.Context, path string, payload *FindingOverwritePayload) (*http.Response, error) {
	req, err := c.NewSubmitAFindingOverwriteFindingsRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSubmitAFindingOverwriteFindingsRequest create the request corresponding to the Submit a Finding Overwrite action endpoint of the findings resource.
func (c *Client) NewSubmitAFindingOverwriteFindingsRequest(ctx context.Context, path string, payload *FindingOverwritePayload) (*http.Request, error) {
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

// ListFindingsFindingsPath computes a request path to the list findings action of findings.
func ListFindingsFindingsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/findings", param0)
}

// List all findings from a team.
func (c *Client) ListFindingsFindings(ctx context.Context, path string, atDate *string, identifier *string, identifiers *string, issueID *string, labels *string, maxDate *string, maxScore *float64, minDate *string, minScore *float64, page *float64, size *float64, sortBy *string, status *string, targetID *string) (*http.Response, error) {
	req, err := c.NewListFindingsFindingsRequest(ctx, path, atDate, identifier, identifiers, issueID, labels, maxDate, maxScore, minDate, minScore, page, size, sortBy, status, targetID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFindingsFindingsRequest create the request corresponding to the list findings action endpoint of the findings resource.
func (c *Client) NewListFindingsFindingsRequest(ctx context.Context, path string, atDate *string, identifier *string, identifiers *string, issueID *string, labels *string, maxDate *string, maxScore *float64, minDate *string, minScore *float64, page *float64, size *float64, sortBy *string, status *string, targetID *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if atDate != nil {
		values.Set("atDate", *atDate)
	}
	if identifier != nil {
		values.Set("identifier", *identifier)
	}
	if identifiers != nil {
		values.Set("identifiers", *identifiers)
	}
	if issueID != nil {
		values.Set("issueID", *issueID)
	}
	if labels != nil {
		values.Set("labels", *labels)
	}
	if maxDate != nil {
		values.Set("maxDate", *maxDate)
	}
	if maxScore != nil {
		tmp118 := strconv.FormatFloat(*maxScore, 'f', -1, 64)
		values.Set("maxScore", tmp118)
	}
	if minDate != nil {
		values.Set("minDate", *minDate)
	}
	if minScore != nil {
		tmp119 := strconv.FormatFloat(*minScore, 'f', -1, 64)
		values.Set("minScore", tmp119)
	}
	if page != nil {
		tmp120 := strconv.FormatFloat(*page, 'f', -1, 64)
		values.Set("page", tmp120)
	}
	if size != nil {
		tmp121 := strconv.FormatFloat(*size, 'f', -1, 64)
		values.Set("size", tmp121)
	}
	if sortBy != nil {
		values.Set("sortBy", *sortBy)
	}
	if status != nil {
		values.Set("status", *status)
	}
	if targetID != nil {
		values.Set("targetID", *targetID)
	}
	u.RawQuery = values.Encode()
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

// ListFindingsIssuesFindingsPath computes a request path to the list findings issues action of findings.
func ListFindingsIssuesFindingsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/findings/issues", param0)
}

// List number of findings and max score per issue.
func (c *Client) ListFindingsIssuesFindings(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, minDate *string, page *float64, size *float64, sortBy *string, status *string, targetID *string) (*http.Response, error) {
	req, err := c.NewListFindingsIssuesFindingsRequest(ctx, path, atDate, identifiers, labels, maxDate, minDate, page, size, sortBy, status, targetID)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFindingsIssuesFindingsRequest create the request corresponding to the list findings issues action endpoint of the findings resource.
func (c *Client) NewListFindingsIssuesFindingsRequest(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, minDate *string, page *float64, size *float64, sortBy *string, status *string, targetID *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if atDate != nil {
		values.Set("atDate", *atDate)
	}
	if identifiers != nil {
		values.Set("identifiers", *identifiers)
	}
	if labels != nil {
		values.Set("labels", *labels)
	}
	if maxDate != nil {
		values.Set("maxDate", *maxDate)
	}
	if minDate != nil {
		values.Set("minDate", *minDate)
	}
	if page != nil {
		tmp122 := strconv.FormatFloat(*page, 'f', -1, 64)
		values.Set("page", tmp122)
	}
	if size != nil {
		tmp123 := strconv.FormatFloat(*size, 'f', -1, 64)
		values.Set("size", tmp123)
	}
	if sortBy != nil {
		values.Set("sortBy", *sortBy)
	}
	if status != nil {
		values.Set("status", *status)
	}
	if targetID != nil {
		values.Set("targetID", *targetID)
	}
	u.RawQuery = values.Encode()
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

// ListFindingsTargetsFindingsPath computes a request path to the list findings targets action of findings.
func ListFindingsTargetsFindingsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/findings/targets", param0)
}

// List number of findings and max score per target.
func (c *Client) ListFindingsTargetsFindings(ctx context.Context, path string, atDate *string, identifiers *string, issueID *string, labels *string, maxDate *string, minDate *string, page *float64, size *float64, sortBy *string, status *string) (*http.Response, error) {
	req, err := c.NewListFindingsTargetsFindingsRequest(ctx, path, atDate, identifiers, issueID, labels, maxDate, minDate, page, size, sortBy, status)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFindingsTargetsFindingsRequest create the request corresponding to the list findings targets action endpoint of the findings resource.
func (c *Client) NewListFindingsTargetsFindingsRequest(ctx context.Context, path string, atDate *string, identifiers *string, issueID *string, labels *string, maxDate *string, minDate *string, page *float64, size *float64, sortBy *string, status *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if atDate != nil {
		values.Set("atDate", *atDate)
	}
	if identifiers != nil {
		values.Set("identifiers", *identifiers)
	}
	if issueID != nil {
		values.Set("issueID", *issueID)
	}
	if labels != nil {
		values.Set("labels", *labels)
	}
	if maxDate != nil {
		values.Set("maxDate", *maxDate)
	}
	if minDate != nil {
		values.Set("minDate", *minDate)
	}
	if page != nil {
		tmp124 := strconv.FormatFloat(*page, 'f', -1, 64)
		values.Set("page", tmp124)
	}
	if size != nil {
		tmp125 := strconv.FormatFloat(*size, 'f', -1, 64)
		values.Set("size", tmp125)
	}
	if sortBy != nil {
		values.Set("sortBy", *sortBy)
	}
	if status != nil {
		values.Set("status", *status)
	}
	u.RawQuery = values.Encode()
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
