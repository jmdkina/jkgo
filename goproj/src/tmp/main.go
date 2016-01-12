package main

import (
	"flag"
	"jk/jklog"
)

var (
	style = flag.String("style", "", "what to execute")
)

func main() {
	flag.Parse()
	if *style == "" {
		jklog.L().Warnln("Give me some args to Use: ./main -style xxx")
		return
	}
	jklog.L().Infoln("use_style: ", *style)

	switch *style {
	case "time":
		timeprint()
		break
	case "net":
		interface_use()
		break
	}

}
