package internal

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/entrustcorporation/dv/dns01"
	"github.com/entrustcorporation/dv/providers/internal/errutils"
	querystring "github.com/google/go-querystring/query"
)

const apiURL = "https://api.nearlyfreespeech.net"

const authenticationHeader = "X-NFSN-Authentication"

const saltBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Client struct {
	login  string
	apiKey string

	baseURL    *url.URL
	HTTPClient *http.Client
}

func NewClient(login string, apiKey string) *Client {
	baseURL, _ := url.Parse(apiURL)

	return &Client{
		login:      login,
		apiKey:     apiKey,
		baseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c Client) AddRecord(ctx context.Context, domain string, record Record) error {
	endpoint := c.baseURL.JoinPath("dns", dns01.UnFqdn(domain), "addRR")

	params, err := querystring.Values(record)
	if err != nil {
		return err
	}

	return c.doRequest(ctx, endpoint, params)
}

func (c Client) RemoveRecord(ctx context.Context, domain string, record Record) error {
	endpoint := c.baseURL.JoinPath("dns", dns01.UnFqdn(domain), "removeRR")

	params, err := querystring.Values(record)
	if err != nil {
		return err
	}

	return c.doRequest(ctx, endpoint, params)
}

func (c Client) doRequest(ctx context.Context, endpoint *url.URL, params url.Values) error {
	payload := params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(payload))
	if err != nil {
		return fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(authenticationHeader, c.createSignature(endpoint.Path, payload))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return errutils.NewHTTPDoError(req, err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return parseError(req, resp)
	}

	return nil
}

func (c Client) createSignature(uri string, body string) string {
	// This is the only part of this that needs to be serialized.
	salt := make([]byte, 16)
	for i := 0; i < 16; i++ {
		salt[i] = saltBytes[rand.Intn(len(saltBytes))]
	}

	// Header is "login;timestamp;salt;hash".
	// hash is SHA1("login;timestamp;salt;api-key;request-uri;body-hash")
	// and body-hash is SHA1(body).

	bodyHash := sha1.Sum([]byte(body))
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	hashInput := fmt.Sprintf("%s;%s;%s;%s;%s;%02x", c.login, timestamp, salt, c.apiKey, uri, bodyHash)

	return fmt.Sprintf("%s;%s;%s;%02x", c.login, timestamp, salt, sha1.Sum([]byte(hashInput)))
}

func parseError(req *http.Request, resp *http.Response) error {
	raw, _ := io.ReadAll(resp.Body)

	errAPI := &APIError{}
	err := json.Unmarshal(raw, errAPI)
	if err != nil {
		return errutils.NewUnexpectedStatusCodeError(req, resp.StatusCode, raw)
	}

	return errAPI
}
