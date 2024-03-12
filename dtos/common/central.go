package common

type Metrics struct {
	MemAlloc       uint64 `json:"memAlloc"`
	MemFrees       uint64 `json:"memFrees"`
	MemLiveObjects uint64 `json:"memLiveObjects"`
	MemMallocs     uint64 `json:"memMallocs"`
	MemSys         uint64 `json:"memSys"`
	MemTotalAlloc  uint64 `json:"memTotalAlloc"`
	CpuBusyAvg     uint8  `json:"cpuBusyAvg"`
}

// MetricsResponse defines the providing memory and cpu utilization stats of the service.
// This object and its properties correspond to the MetricsResponse object in the API specification
type MetricsResponse struct {
	Versionable `json:",inline"`
	Metrics     Metrics `json:"metrics"`
	ServiceName string  `json:"serviceName"`
}

// NewMetricsResponse creates new MetricsResponse with all fields set appropriately
func NewMetricsResponse(metrics Metrics, serviceName string) MetricsResponse {
	return MetricsResponse{
		Versionable: NewVersionable(),
		Metrics:     metrics,
		ServiceName: serviceName,
	}
}

// BaseWithMetricsResponse defines the base content for response DTOs (data transfer objects).
// This object and its properties correspond to the BaseWithMetricsResponse object in the API specification
type BaseWithMetricsResponse struct {
	BaseResponse `json:",inline"`
	ServiceName  string      `json:"serviceName"`
	Metrics      interface{} `json:"metrics"`
}
