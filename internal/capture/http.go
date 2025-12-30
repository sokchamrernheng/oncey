package capture

import "net/http"

type HttpCaptureWriter struct {
	http.ResponseWriter
	Status int
	Body []byte
}

func NewHttpCaptureWriter(w http.ResponseWriter) *HttpCaptureWriter {
	return &HttpCaptureWriter{
		ResponseWriter: w,
	}
}

func (hcw *HttpCaptureWriter) WriteHeader(statusCode int) {
	hcw.Status = statusCode
	hcw.ResponseWriter.WriteHeader(statusCode)
}

func (hcw *HttpCaptureWriter) Write(b []byte) (int, error) {
	hcw.Body = append(hcw.Body, b...)
	return hcw.ResponseWriter.Write(b)
}
