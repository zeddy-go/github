package sdk

import (
	"errors"
	"github.com/zeddy-go/github/payload"
	"github.com/zeddy-go/zeddy/http/client"
	"github.com/zeddy-go/zeddy/mapx"
	"net/url"
	"strconv"
)

const (
	starredListUri = "user/starred"
)

func WithPage(page uint) func(url.Values) {
	return func(values url.Values) {
		values.Set("page", strconv.FormatUint(uint64(page), 10))
	}
}

func WithPerPage(perPage uint) func(values url.Values) {
	return func(values url.Values) {
		values.Set("per_page", strconv.FormatUint(uint64(perPage), 10))
	}
}

var defaultQuery = url.Values{
	"sort":      []string{"created"},
	"direction": []string{"desc"},
	"per_page":  []string{"30"},
	"page":      []string{"1"},
}

const defaultVersion = "2022-11-28"

func NewClient(token string) *Client {
	return &Client{
		httpClient: client.NewClient(
			client.WithBaseUrl("https://api.github.com"),
		).Debug(),
		version: defaultVersion,
		token:   token,
	}
}

type Client struct {
	httpClient *client.Client
	version    string
	token      string
}

func (c *Client) GetToken() string {
	return c.token
}

func (c *Client) GetClient() *client.Client {
	return c.httpClient.
		AddHeader("X-GitHub-Api-Version", c.version).
		AddHeader("Authorization", "Bearer "+c.GetToken())
}

func (c *Client) StarredRepo(opts ...func(url.Values)) (resp payload.StarredRepoResp, err error) {
	query := mapx.CloneSimpleMapSlice(defaultQuery)
	for _, opt := range opts {
		opt(query)
	}

	res, err := c.GetClient().
		AddHeader("Accept", "application/vnd.github.v3.star+json").
		SetQuery(query).
		Get(starredListUri)
	if err != nil {
		return
	}
	if res.IsError() {
		var e payload.Error
		err = res.ScanJsonBody(&e)
		if err != nil {
			return
		}
		err = errors.New(e.Message)
	} else {
		err = res.ScanJsonBody(&resp)
	}

	return
}
