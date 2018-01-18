log_enable(false);

var current_image = 0;
var query_length = 10;
var image_path_prefix = "/images/imgs/q10/";

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];

    // init get info

    
    // $("#fulldisplay").height($(window).height());
    // $(".carousel").carousel('pause');
    var jmdkina = new Vue({
        el: '#jmdkina',
        delimiters: ['${', '}'],
        data: {
            images: []
        },
        beforeCreate: function() {
            load_images();
        },
        methods: {
            showmore: function() {
                load_images();
            },
            fulldisplay: function(img) {
                var name = img.name;
                var content = img.content;
                $(".carousel-inner").html(""); // clear
                for (var i = 0; i < jmdkina.images.length; i++) {
                    var thisimg = jmdkina.images[i];
                    var imageurl = image_path_prefix + thisimg.path + "/" + thisimg.name;
                    var obj = "<div class='carousel-item'>" + 
                    "<img class='d-block w-100' alt='" + thisimg.name +
                    "' src='"+ imageurl +"'/>" + 
                    "</div></div>";
                    $(".carousel-inner").append(obj);
                }
                $(".modal-title").text(name);
                $(".carousel-inner").find(".active").removeClass("active");
                log_print("now index  " + img.index);
                $(".carousel-inner").find(".carousel-item").eq(img.index).addClass("active");
                $(".modal").modal('show');
            }
        }
    });

function load_images() {
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
                if (ret.Result.length == 0) {
                    $(".showmore").attr("disabled", true);
                    $(".showmore").html("No more");
                    return;
                }
                for (var i = 0; i < ret.Result.length; i++) {
                    ret.Result[i]["imageurl"] = image_path_prefix + 
                        ret.Result[i]["path"] + "/" + ret.Result[i]["name"];
                    ret.Result[i]["timestr"] = time_unix2string(ret.Result[i]["createtime"]);
                    ret.Result[i]["index"] = current_image + i;
                    jmdkina.images.push(ret.Result[i])
                }
                current_image = jmdkina.images.length;
            } catch(e) {
                layer.msg("Fail parse response " + response);
            }
        }
    });
}
});