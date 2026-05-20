package placeholder

import (
	"fmt"
	"time"

	"github.com/glebateee/core/logging"
)

type MonthHandler struct {
	logging.Logger
}

func (d MonthHandler) GetMonth() string {
	return fmt.Sprintf("Month: %s", time.Now().Month().String())
}
