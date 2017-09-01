/**
 * Created by v on 9/1/2017.
 */
$(function(){
    Vue.config.delimiters = ['${', '}#'];
    var add_fileserver = new Vue({
        el: '#dboperation',
        data: {
            host:"127.0.0.1",
            dbs:""
        },
        methods: {
            query_dbs: function() {
                var str = "jk=query_dbs&host=" + this.host;
                console.log(str);
                this.result = "";
                $.ajax({
                    url: "db",
                    method:"POST",
                    data: str,
                    success: function(response, textStatus) {
                        console.log("querydb success" + response + textStatus);
                        ret = $.parseJSON(response)
                        dbs = ret.Result;
                        console.log(dbs[1]);
                        add_fileserver.result = "success";
                    }
                });
            },
        }
    });
});