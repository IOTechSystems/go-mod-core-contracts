package xrtmodels

type Schedule struct {
	Name     string   `json:"name"`
	Device   string   `json:"device"`
	Resource []string `json:"resource"`
	Interval uint64   `json:"interval"`
}
