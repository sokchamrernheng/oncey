package capture

import "net/http"

type HttpCaptureWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func NewHttpCaptureWriter(w http.ResponseWriter) *HttpCaptureWriter {
	return &HttpCaptureWriter{
		ResponseWriter: w,
	}
}

func (hcw *HttpCaptureWriter) GetResult() []byte {
	return hcw.body
}

func (hcw *HttpCaptureWriter) WriteHeader(statusCode int) {
	hcw.status = statusCode
	hcw.ResponseWriter.WriteHeader(statusCode)
}

func (hcw *HttpCaptureWriter) Write(b []byte) (int, error) {
	hcw.body = b
	return hcw.ResponseWriter.Write(b)
}
