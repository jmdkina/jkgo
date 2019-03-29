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


(function($) {  
    $.extend({  
        ConvertTime: {  
            /**  
             * 日期 转换为 Unix时间戳  
             * @param <int> year    年  
             * @param <int> month   月  
             * @param <int> day     日  
             * @param <int> hour    时  
             * @param <int> minute  分  
             * @param <int> second  秒  
             * @return <int>        unix时间戳(秒)  
             */ 
            DateToUnix: function(year, month, day, hour, minute, second){  
                var oDate = new Date(Date.UTC(parseInt(year),  
                        parseInt(month),  
                        parseInt(day),  
                        parseInt(hour),  
                        parseInt(minute),  
                        parseInt(second)  
                    )  
                );  
                return (oDate.getTime()/1000);  
            },  
            /**  
             * 时间戳转换日期  
             * @param <int>     unixTime    待时间戳(秒)  
             * @param <bool>    isFull      返回完整时间(Y-m-d 或者 Y-m-d H:i:s)  
             * @param <int>     timeZone    时区  
             */ 
            UnixToDate: function(unixTime, isFull, timeZone){  
                if (typeof(timeZone) == 'number')  
                {  
                    unixTime = parseInt(unixTime) + parseInt(timeZone) * 60 * 60;  
                }  
                var time = new Date(unixTime*1000);  
                var ymdhis = "";  
                ymdhis += time.getUTCFullYear() + "-";  
                ymdhis += time.getUTCMonth() + 1 + "-";  
                ymdhis += time.getUTCDate();  
                if ( isFull === true )  
                {  
                    ymdhis += " " + time.getUTCHours() + ":";  
                    ymdhis += time.getUTCMinutes() + ":";  
                    ymdhis += time.getUTCSeconds();  
                }  
                return ymdhis;  
            }  
        }  
    });  
})(jQuery); 
