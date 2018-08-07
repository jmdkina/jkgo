log_enable(false)

$(function(){
    var resume = new Vue({
        el: '#resume',
        delimiters: ['${', '}'],
        data: {
			info: ""
        },
        beforeCreate: function() {
            load_startup();
        },
		methods: {
			change: function() {
			    change_new();
			}
		}
    });

function load_startup() {
	var str = "jk=resume&cmd=query_info";
	log_print(str);
	$.ajax({
		url: "resume_set",
		method: "POST",
		data: str,
		success: function(response, textStatus) {
			log_print("query result " + response + ", status: " + textStatus);
			try {
				ret = $.parseJSON(response);
				if (ret.Status != 200) {
					layer.msg("Fail of " + ret.Status);
					return;
				}
				resume.info = ret.Result;
            } catch(e) {
                layer.msg("Fail parse , err: " + e);
            }
		}
	});
}

function change_new() {
	var str = "jk=resume&cmd=change_new&content=" + resume.info;
	log_print(str);
	$.ajax({
		url: "resume_set",
		method: "POST",
		data: str,
		success: function(response, textStatus) {
			log_print("change new " + response + ", status: " + textStatus);
			try {
				ret = $.parseJSON(response);
				if (ret.Status != 200) {
					layer.msg("Fail of " + ret.Status + ", " + ret.Result);
					return;
				} else {
					layer.msg("Write success");
				}
			} catch(e) {
				layer.msg("Fail parse, err : " + e);
			}
		}
	});
}

});

