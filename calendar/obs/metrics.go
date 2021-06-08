package obs

type Metrics interface {
	Record(metricsName string, value float64, tags map[string]string)
}
