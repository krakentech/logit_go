package main

import (
	"fmt"
	logit "github.com/krakentech/logit_go"
)

func main() {
	logit.SetIsDebug(true)

	t := logit.NewTracker("track logging")

	logit.Debug("This is a debug message: %d", 42)
	logit.Info("This is an info message")
	logit.Warn("This is a warning message")
	logit.Error("This is an error message")
	logit.Err(fmt.Errorf("this is an error message"), "This is the message")

	logit.DebugData([]string{"A", "B", "C"}, false, "Debug Data (formatted: %t)", false)
	logit.DebugData([]string{"A", "B", "C"}, true, "Debug Data (formatted: %t)", true)

	t.Log()
}
