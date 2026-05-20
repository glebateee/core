package placeholder

import (
	"fmt"
	"time"

	"github.com/glebateee/core/http/actionresults"
	"github.com/glebateee/core/logging"
)

type DayHandler struct {
	logging.Logger
}

func (d DayHandler) GetDay() actionresults.ActionResult {
	return actionresults.NewTemplateAction("day.html", fmt.Sprintf("Day: %d", time.Now().Day()))
}
