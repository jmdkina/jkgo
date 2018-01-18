log_enable(false)

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
            load_shicis();
        },
        methods: {
            showmore: function() {
                load_shicis();
            }
        }
    });


function load_shicis() {
    var str = "jk=shici&cmd=query_shicis&" +
        "index=" + current_shici + "&length=" + query_length;
    log_print(str);
    $.ajax({
        url: "shici",
        method: "POST",
        data: str,
        success: function (response, textStatus) {
            log_print("query new success" + response + ", result:" + textStatus);
            try {
                ret = $.parseJSON(response);
                if (ret.Status != 200) {
                    layer.msg("Fail of " + ret.Status);
                    return;
                }
                if (ret.Result.length == 0) {
                    $(".showmore").attr("disabled", true);
                    $(".showmore").html("No more");
                    return;
                }
                for (var i = 0; i < ret.Result.length; i++) {
                    ret.Result[i]["createtimestr"] = time_unix2string(ret.Result[i]["createtime"]);
                    ret.Result[i]["index"] = i;
                    shici.shicis.push(ret.Result[i])
                }
                current_shici = shici.shicis.length;
            } catch(e) {
                layer.msg("Fail parse , err: " + e);
            }
        }
    });
}
});
