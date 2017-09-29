log_enable(true);
$(function(){
    // Vue.options.delimiters = ['{$', '#}'];
    var osop = new Vue({
        el: '#osoperation',
        delimiters: ['${', '}'],
        data: {
            filename: "",
            content: "",
            code: ""
        },
        methods: {
            open: function() {
                var str = "jk=osop&method=open&filename=" + this.filename;
                log_print(str);
                $.ajax({
                    url: "/index",
                    method:"POST",
                    data: str,
                    success: function(response, textStatus) {
                        log_print("os op success" + response + ", result:" + textStatus);
                        ret = $.parseJSON(response);
                        osop.content = ret.Result;
                    }
                });
            },
            modify: function() {
                var str = "jk=osop&method=modify&filename=" + this.filename + "&content=" + this.content + "&code=" + this.code;
                log_print(str);
                $.ajax({
                    url: "/index",
                    method:"POST",
                    data: str,
                    success: function(response, textStatus) {
                        log_print("querydb success" + response + ", result:" + textStatus);
                        ret = $.parseJSON(response);
                        if (ret.Status != 200) {
                            alert("Fail of " + ret.Status);
                            return;
                        }
                        log_print(ret.Result);
                    }
                });
            },
        }
    });
});