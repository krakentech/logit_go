package main

import logit "github.com/krakentech/logit_go"

func main() {
	logit.SetIsDebug(true)

	t := logit.NewTracker("track logging")

	logit.Debug("This is a debug message: %d", 42)
	logit.Info("This is an info message")
	logit.Warn("This is a warning message")
	logit.Error("This is an error message")

	logit.DebugData([]string{"A", "B", "C"}, false, "Debug Data (formated: %t)", false)
	logit.DebugData([]string{"A", "B", "C"}, true, "Debug Data (formated: %t)", true)

	t.Log()
}
