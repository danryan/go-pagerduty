// Package pagerduty provides an API library for interacting with PagerDuty.
// Much of this code is inspired by https://github.com/google/go-github <3
package pagerduty

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"github.com/google/go-querystring/query"
)

// Client is an API client
type Client struct {
	client    *http.Client
	Subdomain string
	APIKey    string
	BaseURL   *url.URL

	Incidents *IncidentsService
	Users     *UsersService
}

// New returns a Client with the default http.Client
func New(sub, key string) *Client {
	return NewClient(sub, key, nil)
}

// NewClient returns a default client
func NewClient(sub, key string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	u, _ := url.Parse(fmt.Sprintf("https://%s.pagerduty.com/api/v1/", sub))
	client := &Client{
		client:    httpClient,
		APIKey:    key,
		Subdomain: sub,
		BaseURL:   u,
	}

	client.Incidents = &IncidentsService{client: client}
	client.Users = &UsersService{client: client}

	return client
}

// addOptions adds the parameters in opt as URL query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()

	return u.String(), nil
}

// NewRequest builds an http.Request, resolves relative URLs, and sets HTTP headers
func (c *Client) NewRequest(meth string, path string, input interface{}) (*http.Request, error) {
	ref, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(ref)

	buf := new(bytes.Buffer)
	if input != nil {
		if err := json.NewEncoder(buf).Encode(input); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(meth, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", `Token token=`+c.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

// Get makes a GET request
func (c *Client) Get(path string, output interface{}) (*http.Response, error) {
	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, output)
}

// Post makes a POST request
func (c *Client) Post(path string, input, output interface{}) (*http.Response, error) {
	req, err := c.NewRequest("POST", path, input)
	if err != nil {
		return nil, err
	}

	return c.Do(req, output)
}

// Put makes a PUT request
func (c *Client) Put(path string, input, output interface{}) (*http.Response, error) {
	req, err := c.NewRequest("PUT", path, input)
	if err != nil {
		return nil, err
	}

	return c.Do(req, output)
}

// Delete makes a DELETE request
func (c *Client) Delete(path string, input, output interface{}) (*http.Response, error) {
	req, err := c.NewRequest("DELETE", path, input)
	if err != nil {
		return nil, err
	}

	return c.Do(req, output)
}

// Do performs the request
func (c *Client) Do(req *http.Request, output interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = CheckResponse(res)
	if err != nil {
		return res, err
	}
	if output != nil {
		if w, ok := output.(io.Writer); ok {
			io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(output)
		}
	}

	return res, err
}

// An ErrorResponse represents one or more errors created by an API request.
type ErrorResponse struct {
	Response *http.Response
	Message  string    `json:"message"`
	Code     ErrorCode `json:"code"`
	Errors   []string  `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message, r.Errors)
}

// CheckResponse checks the API response for errors, returning them if
// present. A response is considered to have an error if it has a status code
// outside the 200 range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	v := &ErrorResponse{Response: r}
	d, err := ioutil.ReadAll(r.Body)
	if err == nil && d != nil {
		json.Unmarshal(d, v)
	}

	return v
}
