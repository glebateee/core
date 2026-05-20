package placeholder

import (
	"fmt"

	"github.com/glebateee/core/sessions"
)

type CounterHandler struct {
	sessions.Session
}

func (c CounterHandler) GetCounter() string {
	counter := c.Session.GetValueDefault("counter", 0).(int)
	c.Session.SetValue("counter", counter+1)
	return fmt.Sprintf("--> counter: %d", counter)
}
