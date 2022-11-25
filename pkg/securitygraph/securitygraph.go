/*
Copyright 2022 Adevinta
*/

package securitygraph

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ErrAssetDoesNotExists   = errors.New("asset does not exists in the security graph")
	ErrNotEnoughInformation = errors.New("the security graph does not have enough info")
)

// HttpStatusError is returned by the method [IntelAPIClient.BlastRadius] when
// it receives a response with a status code different to  200.
type HttpStatusError struct {
	Status int    `json:"-"`
	Msg    string `json:"msg"`
}

func (h HttpStatusError) Error() string {
	msg := fmt.Sprintf("invalid http status received from the intel API: %d, details: %s", h.Status, h.Msg)
	return msg
}

// Config defines the config parameters needed by a [Client].
type Config struct {
	IntelAPI    string `mapstructure:"intel_api"`
	InsecureTLS string `mapstructure:"insecure_tls"`
}

// BlastRadiusRequest defines the parameters required by the blast radius endpoint.
type BlastRadiusRequest struct {
	AssetIdentifier string `json:"asset_identifier" validate:"required" urlquery:"asset_identifier"`
	AssetType       string `json:"asset_type" validate:"required" urlquery:"asset_type"`
}

// BlastRadiusResponse defines the output of a blast radius calculation.
type BlastRadiusResponse struct {
	// Score contains the blast radius score for a given asset.
	Score float32 `json:"score"`
	// Metadata contains information about how a blast radius was calculated.
	Metadata string `json:"string"`
}

// IntelAPIClient allows to communicates with Intel API exposed by the Security Graph.
type IntelAPIClient struct {
	c        http.Client
	endpoint *url.URL
}

// Returns IntelAPIClient that uses the given config parameters.
func NewIntelAPIClient(cfg Config) (*IntelAPIClient, error) {
	insecure := cfg.InsecureTLS == "1"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	hClient := http.Client{Transport: tr}
	endpoint, err := url.Parse(cfg.IntelAPI)
	if err != nil {
		return nil, fmt.Errorf("invalid intel API endpoint url %s", endpoint)
	}
	client := &IntelAPIClient{
		c:        hClient,
		endpoint: endpoint,
	}
	return client, nil
}

func (i *IntelAPIClient) urlBlastRadius(identifier string, asset_type string) string {
	u := i.endpoint.JoinPath("/v1/blast-radius")
	q := u.Query()
	if identifier != "" {
		q.Set("team_identifier", identifier)
	}
	if asset_type != "" {
		q.Set("asset_type", asset_type)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

// BlastRadius calls the blast radius endpoint of the Intel API using the given
// parameters.
func (i *IntelAPIClient) BlastRadius(req BlastRadiusRequest) (BlastRadiusResponse, error) {
	u := i.urlBlastRadius(req.AssetIdentifier, req.AssetType)
	resp, err := i.c.Get(u)
	if err != nil {
		return BlastRadiusResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var b BlastRadiusResponse
		err := json.NewDecoder(resp.Body).Decode(&b)
		if err != nil {
			return BlastRadiusResponse{}, fmt.Errorf("invalid response: %w", err)
		}
		return b, nil
	}

	intelAPIErr := HttpStatusError{Status: resp.StatusCode}
	// Some errors can also return a json payload with extended info.
	if resp.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(resp.Body).Decode(&intelAPIErr)
		if err != nil {
			intelAPIErr.Msg = fmt.Sprintf("error decoding extended info from the response %v", err)
		}
	}
	return BlastRadiusResponse{}, &intelAPIErr
}
