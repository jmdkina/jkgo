log_enable(true);

var current_image = 0;
var query_length = 10;
var image_path_prefix = "/imgs/q10/";

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];

    // init get info

    var jmdkina = new Vue({
        el: '#jmdkina',
        delimiters: ['${', '}'],
        data: {
            images: []
        },
        beforeCreate: function() {
            var str = "jk=jmdkina&cmd=query_images" + "&index=" + current_image +
                "&length=" + query_length;
            log_print(str);
            $.ajax({
                url: "jmdkina",
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
                        // log_print("length: " + ret.Result.length)
                        jmdkina.images = ret.Result;
                        for (var i = 0; i < jmdkina.images.length; i++) {
                            jmdkina.images[i]["imageurl"] = image_path_prefix + jmdkina.images[i]["path"]
                        }
                        current_image += jmdkina.images.length;
                    } catch(e) {
                        layer.msg("Fail parse response " + response);
                    }
                }
             });
        },
        methods: {
            showmore: function() {
                var str = "jk=jmdkina&cmd=query_images_more&" +
                    "index=" + current_image + "&length=" + query_length;
                log_print(str);
                $.ajax({
                    url: "jmdkina",
                    method: "POST",
                    data: str,
                    success: function (response, textStatus) {
                        log_print("addnew success" + response + ", result:" + textStatus);
                        ret = $.parseJSON(response);
                        try {
                            ret = $.parseJSON(response);
                            if (ret.Status != 200) {
                                layer.msg("Fail of " + ret.Status);
                                return;
                            }
                            jmdkina.images = ret.Result;
                            current_image += jmdkina.images.length;
                            layer.msg("Add New Success!");
                        } catch(e) {
                            layer.msg("Fail parse response " + response);
                        }
                    }
                });
            }
        }
    });
});