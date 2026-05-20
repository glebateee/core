package sessions

import (
	"context"
	"time"

	"github.com/glebateee/core/config"
	"github.com/glebateee/core/pipeline"
	gorilla "github.com/gorilla/sessions"
)

type SessionComponent struct {
	store *gorilla.CookieStore
	config.Config
}

func (s *SessionComponent) Init() {
	cookieKey, found := s.Config.GetString("sessions:key")
	if !found {
		panic("session key not set in config")
	}
	if s.GetBoolDefault("sessions:cyclekey", true) {
		cookieKey += time.Now().String()
	}
	s.store = gorilla.NewCookieStore([]byte(cookieKey))
}

func (s *SessionComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	session, _ := s.store.Get(ctx.Request, SESSION__CONTEXT_KEY)
	c := context.WithValue(ctx.Context(), SESSION__CONTEXT_KEY, session)
	ctx.Request = ctx.Request.WithContext(c)
	next(ctx)
	s.store.Save(ctx.Request, ctx.ResponseWriter, session)
}
