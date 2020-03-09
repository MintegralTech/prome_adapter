package client

import (
	"time"
    "net/url"
    "net/http"

    "github.com/spf13/viper"
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/log/level"
    "github.com/prometheus/common/model"
    "github.com/prometheus/prometheus/prompb"
)

type Client struct {
    logger log.Logger

    url     url.URL
    client *http.Client
    timeout time.Duration
}

func NewClient(logger log.Logger, url string, timeout time.Duration) *Client {
    return &Client{
        logger:  logger,
        url:     url,
        timeout: timeout,
    }
}

func (c *Client) Write(samples model.Samples) error {
}

