package sessions

import (
	"context"
	"fmt"

	"github.com/glebateee/core/services"
	gorilla "github.com/gorilla/sessions"
)

const SESSION__CONTEXT_KEY = "pro_go_session"

type Session interface {
	GetValue(key string) any
	GetValueDefault(key string, defVal any) any
	SetValue(key string, val any)
}

type SessionAdaptor struct {
	gSession *gorilla.Session
}

func (a *SessionAdaptor) GetValue(key string) any {
	return a.gSession.Values[key]
}

func (a *SessionAdaptor) GetValueDefault(key string, defVal any) any {
	if val, ok := a.gSession.Values[key]; ok {
		return val
	}
	return defVal
}

func (a *SessionAdaptor) SetValue(key string, val any) {
	if val == nil {
		a.gSession.Values[key] = nil
	} else {
		switch valType := val.(type) {
		case int, float64, bool, string:
			a.gSession.Values[key] = valType
		default:
			panic(fmt.Sprintf("session store only int, float64, bool, string: %T", valType))
		}
	}
}

func RegisterSessionService() {
	if err := services.AddScoped(func(ctx context.Context) Session {
		session := ctx.Value(SESSION__CONTEXT_KEY)
		if s, ok := session.(*gorilla.Session); ok {
			return &SessionAdaptor{gSession: s}
		}
		panic("session not found in context")
	}); err != nil {
		panic(err)
	}
}
