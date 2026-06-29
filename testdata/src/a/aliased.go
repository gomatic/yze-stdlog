package a

import (
	applog "log" // want `standard log package`
	_ "log"      // want `standard log package`
)

var _ = applog.Println
