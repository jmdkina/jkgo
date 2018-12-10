package jkmisc

type JKWetherCity struct {
}

type JKWetherInfoToday struct {
	Wind        string
	Week        string
	City        string
	Temperature string
}

type JKWetherInfoFuture struct {
}

type JKWetherInfoResultSk struct {
	Temp     string
	Humidity string
}

type JKWetherInfoResult struct {
	Sk     JKWetherInfoResultSk
	Today  JKWetherInfoToday
	Future JKWetherInfoFuture
}

//[sk:map[temp:13 wind_direction:东北风 wind_strength:3级 humidity:75% time:10:52]
// today:map[wind:持续无风向微风 week: 星期一 city:深圳 comfort_index: temperature:14℃~17℃ dressing_index:较冷
//   exercise_index:较适宜 weather:阴转小雨 weather_id:map[fa:02 fb:07] date_y:2018年12月10日
//   dressing_advice:建议着厚外套加毛衣等服装。年老体弱者宜着大衣、呢外套加羊毛衫 。
//   uv_index:最弱 wash_index:较适宜 travel_index:较适宜 drying_index:]
// future:map[day_20181210:map[temperature:14℃~17℃ weather:阴转小雨 weather_id:map[fa:02 fb:07] wind:持续无风向微风 week:星期一 date:20181210] day_20181211:map[temperature:11℃~19℃ weather:阴 weather_id:map[fa:02 fb:02] wind:北风3-5级 week:星期二 date:20181211] day_20181212:map[temperature:11℃~16℃ weather:多云 weather_id:map[fa:01 fb:01] wind:持续无风向微风 week:星期三 date:20181212] day_20181213:map[temperature:13℃~17℃ weather:晴转多云 weather_id:map[fa:00 fb:01] wind:东北风3-5级 week:星期四 date:20181213] day_20181214:map[date:20181214 temperature:15℃~20℃ weather:多云 weather_id:map[fa:01 fb:01] wind:东北风3-5级 week:星期五] day_20181215:map[week:星期六 date:20181215 temperature:11℃~19℃ weather:阴 weather_id:map[fa:02 fb:02] wind:北风3-5级] day_20181216:map[date:20181216 temperature:11℃~16℃ weather:多云 weather_id:map[fa:01 fb:01] wind:持续无风向微风 week:星期日]]]

type JKWetherInfo struct {
	Error_code int
	Result     JKWetherInfoResult
}

type JKWetherBase interface {
	Query(location string) (*JKWetherInfo, error)
	QueryCity() (*[]JKWetherCity, error)
}
