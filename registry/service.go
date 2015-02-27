package registry

// Service is a Registry service sporting an URL and an access error count
type Service struct {
	Name  string
	URL   string
	fails int
}
