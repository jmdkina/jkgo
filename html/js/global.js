/**
 * Created by jmdvi on 2017/7/22.
 */

var log_out = false;

function log_enable(e) {
	log_out = e
}

function log_print(str) {
	if (log_out) {
		console.log(str);
	}
}

function time_unix2string(unix) {
    var unixTimestamp = new Date(unix* 1000); 
    commonTime = unixTimestamp.toLocaleString();
    return commonTime
}
