package pipeline

import (
	"net/http"
	"strings"
)

type DeferredWriter struct {
	http.ResponseWriter
	strings.Builder
	statusCode int
}

func (w *DeferredWriter) Write(data []byte) (int, error) {
	return w.Builder.Write(data)
}

func (w *DeferredWriter) WriteHeader(code int) {
	w.statusCode = code
}

func (w *DeferredWriter) FlushData() (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	w.ResponseWriter.WriteHeader(w.statusCode)
	return w.ResponseWriter.Write([]byte(w.Builder.String()))
}
