package main

import (
	"time"
	"flag"
	"strconv"
	"jk/jklog"
)

func ts_to_string(ts int64) string {
	tt:= time.Unix(ts, 0)
	return tt.Format("2006-01-02-15-04-05 Z0700")
}

var (
	t = flag.String("type", "ts", "timestamp to string")
	v = flag.String("value", "", "timestamp or string")
)

func main() {
	flag.Parse()

	switch *t {
	case "ts":
		i, err := strconv.ParseInt(*v, 10, 32)
		if err != nil {
			jklog.L().Errorln("invalid value, ", err)
			return
		}
		result := ts_to_string(i)
		jklog.L().Infof("%v = %s\n", *v, result)
		return
	case "str":
		t, err := time.Parse("2006-01-02-15-04-05 Z0700", *v)
		if err != nil {
			jklog.L().Errorln("parse failed ", err)
			return
		}
		result := t.Unix()
		jklog.L().Infof("%s = %d\n", *v, result)
	}
}
