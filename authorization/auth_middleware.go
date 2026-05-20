package authorization

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/glebateee/core/authorization/identity"
	"github.com/glebateee/core/config"
	"github.com/glebateee/core/http/handling"
	"github.com/glebateee/core/pipeline"
)

type AuthMiddlewareComponent struct {
	prefix    string
	condition identity.AuthorizationCondition
	pipeline.RequestPipeline
	config.Config
	authFailURL string
	fallbacks   map[*regexp.Regexp]string
}

func NewAuthComponent(
	prefix string,
	condition identity.AuthorizationCondition,
	requestHandlers ...any,
) *AuthMiddlewareComponent {
	entries := []handling.HandlerEntry{}
	for _, handler := range requestHandlers {
		entries = append(entries, handling.HandlerEntry{Prefix: prefix, Handler: handler})
	}
	router := handling.NewRouter(entries...)
	return &AuthMiddlewareComponent{
		prefix:          "/" + prefix,
		condition:       condition,
		RequestPipeline: pipeline.CreatePipeline(router),
		fallbacks:       map[*regexp.Regexp]string{},
	}
}

func (*AuthMiddlewareComponent) ImplementsProcessRequestWithServices() {}

func (c *AuthMiddlewareComponent) Init() {
	c.authFailURL, _ = c.Config.GetString("authorization:failUrl")
}

func (c *AuthMiddlewareComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	user identity.User,
) {
	if strings.HasPrefix(ctx.Request.URL.Path, c.prefix) {
		for expr, target := range c.fallbacks {
			if expr.MatchString(ctx.Request.URL.Path) {
				http.Redirect(ctx.ResponseWriter, ctx.Request, target, http.StatusSeeOther)
				return
			}
		}
		if c.condition.Validate(user) {
			c.RequestPipeline.ProcessRequest(ctx.Request, ctx.ResponseWriter)
		} else {
			if c.authFailURL != "" {
				http.Redirect(ctx.ResponseWriter, ctx.Request, c.authFailURL, http.StatusSeeOther)
			} else if user.IsAuthenticated() {
				ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
			} else {
				ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			}
		}
	} else {
		next(ctx)
	}
}

func (c *AuthMiddlewareComponent) AddFallback(target string, patterns ...string) *AuthMiddlewareComponent {
	for _, p := range patterns {
		c.fallbacks[regexp.MustCompile(p)] = target
	}
	return c
}
