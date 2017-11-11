log_enable(true);

ops = [ "todayall", "toptoday" ];

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];

    // init get info

    var stockstory = new Vue({
        el: '#stockstory',
        delimiters: ['${', '}'],
        data: {
            ideas: [],
            code: "",
            analyse:"",
            remember:"",
            date:"",
            key:"",
            showadd_diag: false
        },
        beforeCreate: function() {
            var str = "jk=stockoperation&op=query_ideas" + "&code=jmdstock";
            log_print(str);
            $.ajax({
                url: "stock",
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
                        stockstory.ideas = ret.Result;
                    } catch(e) {
                        layer.msg("Fail parse response " + response);
                    }
                }
             });
        },
        methods: {
            showdiag: function() {
                this.showadd_diag = !this.showadd_diag;
            },
            addnew: function() {
                var str = "jk=stockoperation&op=addnew&" +
                    "date=" + this.date + "&analyse=" + encodeURIComponent(Base64.encode(this.analyse)) + "&remember=" + this.remember +
                    "&key=" + this.key + "&code=" + this.code;
                log_print(str);
                $.ajax({
                    url: "stock",
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
                            layer.msg("Add New Success!");
                        } catch(e) {
                            layer.msg("Fail parse response " + response);
                        }
                    }
                });
            }
        }
    });

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
                var str = "jk=stockoperation&op=" + this.op + "&code=jmdstock" + this.code;
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
            }
        }
    });
});