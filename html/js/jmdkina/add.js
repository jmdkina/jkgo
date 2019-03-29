log_enable(true);

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];

    // init get info

    var jmdkinaadd = new Vue({
        el: '#jmdkinaadd',
        delimiters: ['${', '}'],
        data: {
            path: "",
            tag: "",
            author: "jmd",
            code:"",
            content:""
        },
        beforeCreate: function() {
            
        },
        methods: {
            add: function() {
                log_print("Start to send upload")
                var form = new FormData(document.forms.namedItem("jmdkinaadd"))
                var xhr = new XMLHttpRequest();
                var url_args = "cmd=add&code=" + this.code +
                    "&path=" + this.path + "&tag=" + this.tag + "&author=" + this.author +
                    "&content=" + this.content;
                xhr.open("post", "/jmdkinaadd?" + url_args, true);
                xhr.upload.onprogress = function(ev) {
                    var percent = 0;
                    if (ev.lengthComputable) {
                        percent = 100 * ev.loaded / ev.total;
                        $("#uploadfile").width(percent + "%");
                    }
                };
                xhr.onload = function(oEvent) {
                    if (xhr.status == 200) {
                        layer.msg("Upload Success");
                    } else {
                        layer.msg("Upload failed " + xhr.status);
                    }
                };
                xhr.send(form);
            }
        }
    });
});