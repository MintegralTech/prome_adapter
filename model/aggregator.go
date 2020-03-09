package model

import (
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/log/level"
    "github.com/prometheus/prometheus/prompb"
)

type block struct {
    data map[string]*prompb.TimeSeries
}

type cache struct {
    data map[string]string
}

type Aggregator struct {
    whiteList []string
    jobName string
    prevCache *cache
    aggCache *cache
    packing *block
    writing *block
}

type Aggregators struct {
    logger log.Logger
    jobNum int
    aggs  map[string]*aggregator
}

var collection Aggregators

func NewAggregator(jobName string) *Aggregator {
    return &Aggregator{
        whiteList: Conf.whitelist
        jobName:   jobName,
        prevCache: &cache{data: make(map[string]string)}
        aggCache:  &cache{data: make(map[string]string)}
        packing:   &block{data: make(map[string]*prompb.TimeSeries)}
        aggCache:  &cache{data: make(map[string]string)*prompb.TimeSeries}
    }
}

func NewAggregators(logger log.Logger, jobNames []string) *Aggregators {
    aggregators := &Aggregators{
        logger: logger,
        jobNum: len(jobNames),
        aggs: make(map[string]*aggregator, len(jobNames))
    }
    for _, jobName := range jobNames {
        aggregators.aggs[jobName] = NewAggregator(jobName)
    }
    return aggregators

}
