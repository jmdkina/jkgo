package helper

import (
	//"fmt"
	"strings"
	"time"
)

const (
	JKFormatString         = "2006-01-02-15:04:05"
	jkformatStringFile     = "2006-01-02-15-04-05"
	jkformatStringFileDate = "2006-01-02"
)

var Month = map[string]string{
	"Jan": "1",
	"Feb": "2",
	"Mar": "3",
	"Apr": "4",
	"May": "5",
	"Jun": "6",
	"Jul": "7",
	"Aug": "8",
	"Sep": "9",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

//convert from "Fri Aug 10 09:40:40 +0800 2012" to "2012.8.10 09:40"
func TimeConv(origin string) (after string) {
	arr1 := strings.Split(origin, " ")
	//Fri Aug 10 09:40:40 +0800 2012
	// 0   1   2     3      4     5
	arr2 := strings.Split(arr1[3], ":")

	after = arr1[5] + "." +
		Month[arr1[1]] + "." +
		arr1[2] + " " +
		arr2[0] + ":" +
		arr2[1]

	return
}

func ConvertStringToUnixTime(tm string) (int64, error) {
	t, err := time.Parse(JKFormatString, tm)
	if err != nil {
		return -1, err
	} else {
		return t.Unix(), nil
	}
}

func FormUnixTimeToString(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format(jkformatStringFile)
}

func FormUnixTimeToStringDate(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format(jkformatStringFileDate)
}
