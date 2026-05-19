package placeholder

import (
	"sync"

	"github.com/glebateee/core/http"

	"github.com/glebateee/core/pipeline"
	"github.com/glebateee/core/pipeline/basic"
	"github.com/glebateee/core/services"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&SimpleMessageComponent{},
	)
}

func Start() {
	res, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		res[0].(*sync.WaitGroup).Wait()
	} else {
		panic(err)
	}
}
