log_enable(true);

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];

    // init get info

    var shiciadd = new Vue({
        el: '#shiciadd',
        delimiters: ['${', '}'],
        data: {
            title: "",
            subtitle: "",
            createtime: "",
            author: "诀梦颠",
            code:"",
            content:""
        },
        beforeCreate: function() {
            
        },
        methods: {
            add: function() {
                log_print("Start to send shici add")
                
                var tmlist = this.createtime.split("-");
                var url_args = "cmd=add&code=" + this.code +
                    "&title=" + this.title + "&subtitle=" + this.subtitle + "&author=" + this.author +
                    "&createtime=" + $.ConvertTime.DateToUnix(tmlist[0], tmlist[1]-1, tmlist[2], tmlist[3]-8, tmlist[4], tmlist[5]) +
                    "&content=" + this.content;
                log_print(url_args);
                $.ajax({
                    url: "shiciadd",
                    method: "POST",
                    data: url_args,
                    success: function (response, textStatus) {
                        try {
                            ret = $.parseJSON(response);
                            layer.msg("add success!");
                        } catch(e) {
                            layer.msg("Fail parse response " + response);
                        }
                    }
                });
            }
        }
    });
});