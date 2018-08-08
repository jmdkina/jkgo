log_enable(true)

$(function(){
    var resume = new Vue({
        el: '#resume',
        delimiters: ['${', '}'],
        data: {
			info: "",
			code: "",
			retinfo: []
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
				resume.retinfo = $.parseJSON(resume.info);
            } catch(e) {
                layer.msg("Fail parse , err: " + e);
            }
		}
	});
}

function change_new() {
	var str = "jk=resume&cmd=change_new&code=" + resume.code + "&content=" + resume.info;
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

