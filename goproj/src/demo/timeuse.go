package main

import (
	"flag"
	"jk/jklog"
	"time"
)

var (
	tm = flag.Int64("tm", 0, "long value")
)

func timeprint() {
	flag.Parse()
	jklog.L().Infoln("begin from here")
	t := time.Now()
	jklog.L().Infoln("time.Now()=", t)

	tU := t.Unix()
	tUN := t.UnixNano()
	tUTC := t.UTC()
	jklog.L().Infoln("t.Unix()=", tU, ", t.UnixNano()=", tUN, ", t.UTC()=", tUTC)

	lo := t.Location()
	jklog.L().Infoln("t.Location()=", lo)

	wd := t.Weekday().String()
	jklog.L().Infoln("weekday()=", wd)

	// duration
	tt := time.Unix(*tm, 0)

	jklog.L().Infoln("The time vaule ", tt.UTC())
	d := time.Since(t)
	dh := d.Hours()
	dm := d.Minutes()
	dn := d.Nanoseconds()
	ds := d.Seconds()
	dstr := d.String()
	jklog.L().Infoln("d.String()=", dstr, ", d.Hours()=",
		dh, ", d.Minutes()=", dm, " d.Nanoseconds()=", dn,
		", d.Seconds()=", ds)
}

func main() {
	timeprint()
}
