package model

import (
    "github.com/prometheus/prometheus/prompb"
    "github.com/hashicorp/terraform/helper/hashcode"
)

type block struct {
    data map[string]*prompb.TimeSeries
}

type cache struct {
    data map[int]string
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
    jobNum int
    aggs  map[string]*Aggregator
}

var Collection Aggregators

func init() {
    Collection = NewAggregators(Conf.jobNames)
}

func NewAggregator(jobName string) *Aggregator {
    return &Aggregator{
        whiteList: Conf.whitelist,
        jobName:   jobName,
        prevCache: &cache{data: make(map[int]string)},
        aggCache:  &cache{data: make(map[int]string)},
        packing:   &block{data: make(map[string]*prompb.TimeSeries)},
        writing:   &block{data: make(map[string]*prompb.TimeSeries)},
    }
}

func NewAggregators(jobNames []string) *Aggregators {
    aggregators := &Aggregators{
        jobNum: len(jobNames),
        aggs: make(map[string]*Aggregator, len(jobNames)),
    }
    for _, jobName := range jobNames {
        aggregators.aggs[jobName] = NewAggregator(jobName)
    }
    return aggregators
}

func (collection *Aggregators) UpdateOne(ts *prompb.TimeSeries) {
     


}
