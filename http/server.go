package http

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/glebateee/core/config"
	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/pipeline"
)

type pipelineAdaptor struct {
	pipeline.RequestPipeline
}

func (p pipelineAdaptor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.ProcessRequest(r, w)
}

func Serve(
	pl pipeline.RequestPipeline,
	cfg config.Config,
	logger logging.Logger,
) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	adaptor := pipelineAdaptor{RequestPipeline: pl}
	enableHTTP := cfg.GetBoolDefault("http:enableHTTP", true)
	if enableHTTP {
		httpPort := cfg.GetIntDefault("http:port", 5000)
		logger.Debugf("starting HTTP server on port: %d", httpPort)
		wg.Add(1)
		go func() {
			if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), adaptor); err != nil {
				panic(err)
			}
		}()
	}
	enableHTTPS := cfg.GetBoolDefault("http:enableHTTPS", false)
	if enableHTTPS {
		httpsPort := cfg.GetIntDefault("http:httpsPort", 5500)
		certFile, cok := cfg.GetString("http:httpsCert")
		keyFile, kok := cfg.GetString("http:httpsKey")
		if cok && kok {
			logger.Debugf("starting HTTPS server on port: %d", httpsPort)
			wg.Add(1)
			go func() {
				if err := http.ListenAndServeTLS(fmt.Sprintf(":%d", httpsPort), certFile, keyFile, adaptor); err != nil {
					panic(err)
				}
			}()
		} else {
			panic("HTTPS settings invalid")
		}
	}
	return &wg
}
