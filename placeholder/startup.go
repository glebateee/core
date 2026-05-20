package placeholder

import (
	"sync"

	"github.com/glebateee/core/http"
	"github.com/glebateee/core/http/handling"
	"github.com/glebateee/core/sessions"

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
		&sessions.SessionComponent{},
		//&SimpleMessageComponent{},
		handling.NewRouter(
			handling.HandlerEntry{Prefix: "", Handler: NameHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: DayHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: MonthHandler{}},
			handling.HandlerEntry{Prefix: "", Handler: CounterHandler{}},
		).AddMethodAlias("/", NameHandler.GetNames),
	)
}

func Start() {
	sessions.RegisterSessionService()
	res, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		res[0].(*sync.WaitGroup).Wait()
	} else {
		panic(err)
	}
}
