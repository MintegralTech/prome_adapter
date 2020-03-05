package client

import (
	"time"

    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/log/level"
    "github.com/prometheus/common/model"
    "github.com/prometheus/prometheus/prompb"
)

type Client struct {
    logger log.Logger

    url     string
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
    conn, err := net.DialTimeout(c.transport, c.address, c.timeout)
    if err != nil {
        return err
    }
    defer conn.Close()

    var buf bytes.Buffer
    for _, s := range samples {
        k := pathFromMetric(s.Metric, c.prefix)
        t := float64(s.Timestamp.UnixNano()) / 1e9
        v := float64(s.Value)
        if math.IsNaN(v) || math.IsInf(v, 0) {
            level.Debug(c.logger).Log("msg", "cannot send value to Graphite, skipping sample", "value", v, "sample", s)
            continue
        }
        fmt.Fprintf(&buf, "%s %f %f\n", k, v, t)
    }

    _, err = conn.Write(buf.Bytes())
    return err
}

