package client

// Request はHTTPリクエストの引数です
type Request struct {
	Timeout float64
	Method  string
	URL     string
}
