package registry

// Service ...
type Service struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	fails int
}
