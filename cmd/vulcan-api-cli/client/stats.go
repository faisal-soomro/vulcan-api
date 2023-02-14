// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "Vulcan-API": stats Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// CoverageStatsPath computes a request path to the coverage action of stats.
func CoverageStatsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/stats/coverage", param0)
}

// Get asset coverage for a team.
func (c *Client) CoverageStats(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewCoverageStatsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCoverageStatsRequest create the request corresponding to the coverage action endpoint of the stats resource.
func (c *Client) NewCoverageStatsRequest(ctx context.Context, path string) (*http.Request, error) {
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

// CurrentExposureStatsPath computes a request path to the current exposure action of stats.
func CurrentExposureStatsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/stats/exposure/current", param0)
}

// Get current exposure statistics for a team. This metric takes into account only the exposure for open vulnerabilities since the last time they were detected.
func (c *Client) CurrentExposureStats(ctx context.Context, path string, maxScore *float64, minScore *float64) (*http.Response, error) {
	req, err := c.NewCurrentExposureStatsRequest(ctx, path, maxScore, minScore)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCurrentExposureStatsRequest create the request corresponding to the current exposure action endpoint of the stats resource.
func (c *Client) NewCurrentExposureStatsRequest(ctx context.Context, path string, maxScore *float64, minScore *float64) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if maxScore != nil {
		tmp135 := strconv.FormatFloat(*maxScore, 'f', -1, 64)
		values.Set("maxScore", tmp135)
	}
	if minScore != nil {
		tmp136 := strconv.FormatFloat(*minScore, 'f', -1, 64)
		values.Set("minScore", tmp136)
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

// ExposureStatsPath computes a request path to the exposure action of stats.
func ExposureStatsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/stats/exposure", param0)
}

// Get exposure statistics for a team. This metric takes into account the exposure across all lifecycle of vulnerabilities.
func (c *Client) ExposureStats(ctx context.Context, path string, atDate *string, maxScore *float64, minScore *float64) (*http.Response, error) {
	req, err := c.NewExposureStatsRequest(ctx, path, atDate, maxScore, minScore)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewExposureStatsRequest create the request corresponding to the exposure action endpoint of the stats resource.
func (c *Client) NewExposureStatsRequest(ctx context.Context, path string, atDate *string, maxScore *float64, minScore *float64) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if atDate != nil {
		values.Set("atDate", *atDate)
	}
	if maxScore != nil {
		tmp137 := strconv.FormatFloat(*maxScore, 'f', -1, 64)
		values.Set("maxScore", tmp137)
	}
	if minScore != nil {
		tmp138 := strconv.FormatFloat(*minScore, 'f', -1, 64)
		values.Set("minScore", tmp138)
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

// FixedStatsPath computes a request path to the fixed action of stats.
func FixedStatsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/stats/fixed", param0)
}

// Get fixed issues statistics for a team.
func (c *Client) FixedStats(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, minDate *string) (*http.Response, error) {
	req, err := c.NewFixedStatsRequest(ctx, path, atDate, identifiers, labels, maxDate, minDate)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFixedStatsRequest create the request corresponding to the fixed action endpoint of the stats resource.
func (c *Client) NewFixedStatsRequest(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, minDate *string) (*http.Request, error) {
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

// MttrStatsPath computes a request path to the mttr action of stats.
func MttrStatsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/stats/mttr", param0)
}

// Get MTR statistics for a team.
func (c *Client) MttrStats(ctx context.Context, path string, maxDate *string, minDate *string) (*http.Response, error) {
	req, err := c.NewMttrStatsRequest(ctx, path, maxDate, minDate)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewMttrStatsRequest create the request corresponding to the mttr action endpoint of the stats resource.
func (c *Client) NewMttrStatsRequest(ctx context.Context, path string, maxDate *string, minDate *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if maxDate != nil {
		values.Set("maxDate", *maxDate)
	}
	if minDate != nil {
		values.Set("minDate", *minDate)
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

// OpenStatsPath computes a request path to the open action of stats.
func OpenStatsPath(teamID string) string {
	param0 := teamID

	return fmt.Sprintf("/api/v1/teams/%s/stats/open", param0)
}

// Get open issues statistics for a team.
func (c *Client) OpenStats(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, minDate *string) (*http.Response, error) {
	req, err := c.NewOpenStatsRequest(ctx, path, atDate, identifiers, labels, maxDate, minDate)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewOpenStatsRequest create the request corresponding to the open action endpoint of the stats resource.
func (c *Client) NewOpenStatsRequest(ctx context.Context, path string, atDate *string, identifiers *string, labels *string, maxDate *string, minDate *string) (*http.Request, error) {
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
