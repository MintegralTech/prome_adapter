package model

import (
    "time"
    "github.com/prometheus/common/model"
    "github.com/prometheus/prometheus/prompb"
    "github.com/hashicorp/terraform/helper/hashcode"
)

type block struct {
    timestamp time.Duration
    data map[string]*prompb.TimeSeries
}

type cache struct {
    data map[int]float64
}

type Aggregator struct {
    whiteList []string
    jobName string
    prevCache *cache
    sumCache *cache
    pack *block
    ready *block
}

type Aggregators struct {
    jobNum int
    aggregators map[string]*Aggregator
}

var Collection Aggregators

func init() {
    Collection = NewAggregators(Conf.jobNames)
}

func NewAggregator(jobName string) *Aggregator {
    return &Aggregator{
        whiteList: Conf.whitelist,
        jobName:   jobName,
        prevCache: &cache{data: make(map[int]*prompb.Sample)},
        sumCache:  &cache{data: make(map[int]*prompb.Sample)},
        pack:   &block{data: make(map[string]*prompb.TimeSeries)},
        ready:   &block{data: make(map[string]*prompb.TimeSeries)},
    }
}

func NewAggregators(jobNames []string) *Aggregators {
    aggs := &Aggregators{
        jobNum: len(jobNames),
        aggregators: make(map[string]*Aggregator, len(jobNames)),
    }
    for _, jobName := range jobNames {
        aggs.aggregators[jobName] = NewAggregator(jobName)
    }
    return aggs
}

func (collection *Aggregators) updatePrevCache(hc int, prompb.Sample) float64 {
    //
}

func (collection *Aggregators) updateSumCache(hc int, prompb.Sample, incVal float64) int {
    //
}

func (collection *Aggregators) pack(sumVal float64) {
    //
}

func (collection *Aggregators) send() {
    //
}

func (collection *Aggregators) MergeMetric(ts *prompb.TimeSeries) {
    m := make(model.Metric, len(ts.Labels))
    for _, l := range *ts.Labels {
        m[model.LabelName(l.Name)] = model.LabelValue(l.Value)
    }
    metrics := fmt.Sprintf(m)
    hc := hashcode(metrics)
    incVal := collection.updatePrevCache(hc, ts.Sample)
    sumVal := collection.updateSumCache(hc, ts.Sample, incVal)
    colleciton.pack(sumVal)
}
