package sdk

import "github.com/zeddy-go/zeddy/http/client"

const (
	starredListUri = "user/starred"
)

type Client struct {
	httpClient *client.Client
}
