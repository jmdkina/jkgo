log_enable(true)

var current_shici = 0;
var query_length = 5;
$(function(){
    // Vue.options.delimiters = ['{$', '#}'];

    // init get info

    
    // $("#fulldisplay").height($(window).height());
    // $(".carousel").carousel('pause');
    var shici = new Vue({
        el: '#shici',
        delimiters: ['${', '}'],
        data: {
            shicis: []
        },
        beforeCreate: function() {
            var str = "jk=shici&cmd=query_shicis" + "&index=" + current_shici +
                "&length=" + query_length;
            log_print(str);
            $.ajax({
                url: "shici",
                method: "POST",
                data: str,
                success: function (response, textStatus) {
                    // log_print("stockquery success" + response + ", result:" + textStatus);
                    ret = $.parseJSON(response);
                    try {
                        ret = $.parseJSON(response);
                        if (ret == undefined || ret.Status != 200) {
                            layer.msg("Fail of " + ret.Status);
                            return;
                        }
                        for (var i = 0; i < ret.Result.length; i++) {
                            ret.Result[i]["createtimestr"] = time_unix2string(ret.Result[i]["createtime"]);
                            ret.Result[i]["index"] = i;
                            shici.shicis.push(ret.Result[i])
                        }
                        current_shici = shici.shicis.length;
                    } catch(e) {
                        layer.msg("Fail parse response " + response);
                    }
                }
             });
        },
    });
});
