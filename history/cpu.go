package history

import "math"

type cpuUsageAggregator struct {
	AggregatedValue float64
	Count           float64
}

func (agg *cpuUsageAggregator) AddValue(value interface{}) {
	if usage, ok := value.(float64); ok {
		agg.AggregatedValue += usage
		agg.Count++
	}
}

func (agg *cpuUsageAggregator) GetAverageValue() interface{} {
	avg := agg.AggregatedValue / agg.Count

	// reset
	agg.AggregatedValue = 0
	agg.Count = 0

	return math.Round(avg/0.01) * 0.01
}
