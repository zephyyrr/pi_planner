package model

import (
	"fmt"
	"time"
)

type Activity struct {
	name           string
	responsible    fmt.Stringer
	Approximations map[fmt.Stringer]time.Duration
}
