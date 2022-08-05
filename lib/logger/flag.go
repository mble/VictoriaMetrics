package logger

import (
	"flag"
	"strings"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/buildinfo"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/flagutil"
)

func logAllFlags() {
	Infof("build version: %s", buildinfo.Version)
	Infof("command line flags")
	flag.Visit(func(f *flag.Flag) {
		lname := strings.ToLower(f.Name)
		value := f.Value.String()
		if flagutil.IsSecretFlag(lname) {
			value = "secret"
		}
		if flagutil.IsURLFlag(lname) {
			value = flagutil.RedactURLFlagPassword(value)
		}
		Infof("flag %q=%q", f.Name, value)
	})
}
