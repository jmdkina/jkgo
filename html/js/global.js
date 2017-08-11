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