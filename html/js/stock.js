log_enable(true);

ops = [ "todayall", "toptoday" ]

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];
    var dbop = new Vue({
        el: '#stockoperation',
        delimiters: ['${', '}'],
        data: {
            op: "",
            code: "",
            result: "",
            operations: [ {"value":"todayall"}, {"value":"toptoday"}]
        },
        methods: {
            execute: function() {
                var str = "jk=stockoperation&op=" + this.op + "&code=" + this.code;
                log_print(str);
                $.ajax({
                    url: "stock",
                    method:"POST",
                    data: str,
                    success: function(response, textStatus) {
                        log_print("stockoperation success" + response + ", result:" + textStatus);
                        ret = $.parseJSON(response);
                        if (ret.Status != 200) {
                            layer.msg("Fail of " + ret.Status);
                            return;
                        }
                        layer.msg("Success! " + ret.Result);
                    }
                });
            },
        }
    });
});