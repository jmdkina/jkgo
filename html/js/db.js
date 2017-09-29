/**
 * Created by v on 9/1/2017.
 */

log_enable(true);
$(function(){
    // Vue.options.delimiters = ['{$', '#}'];
    var dbop = new Vue({
        el: '#dboperation',
        delimiters: ['${', '}'],
        data: {
            host:"127.0.0.1",
            dbs:[],
            dbname:"",
            colls: [],
            collname: "",
            data: []
        },
        methods: {
            query_dbs: function() {
                var str = "jk=query_dbs&host=" + this.host;
                log_print(str);
                $.ajax({
                    url: "db",
                    method:"POST",
                    data: str,
                    success: function(response, textStatus) {
                        log_print("querydb success" + response + ", result:" + textStatus);
                        ret = $.parseJSON(response);
                        dbop.dbs = [];
                        for (var i = 0; i < ret.Result.length; i++) {
                            var t_v = [];
                            t_v["text"] = ret.Result[i];
                            t_v["value"] = ret.Result[i];
                            dbop.dbs.push(t_v);
                        }
                        $("#collop").show();
                    }
                });
            },
            query_colls: function() {
                var str = "jk=query_colls&host="+this.host + "&dbname=" + this.dbname;
                log_print(str);
                $.ajax({
                    url: "db",
                    method:"POST",
                    data: str,
                    success: function(response, textStatus) {
                        log_print("querycolls success" + response + ", result:" + textStatus);
                        ret = $.parseJSON(response);
                        dbop.colls = [];
                        for (var i = 0; i < ret.Result.length; i++) {
                            var t_v = [];
                            t_v["text"] = ret.Result[i];
                            t_v["value"] = ret.Result[i];
                            dbop.colls.push(t_v);
                        }
                        $("#colldata").show();
                    }
                });
            },
            query_data: function() {
                var str = "jk=query_data&host="+this.host + "&dbname=" + this.dbname + "&collname=" + this.collname;
                log_print(str);
                $.ajax({
                    url: "db",
                    method:"POST",
                    data: str,
                    success: function(response, textStatus) {
                        log_print("query data success" + response + ", result:" + textStatus);
                        ret = $.parseJSON(response);
                        log_print("Result data length " + ret.Result.length);
                 
                        objheader = "<tr>";
                        objbody = "";
                        for (var i = 0; i < ret.Result.length; i++) {
                           
                            objbody += "<tr>";
                            for (var key in ret.Result[i]) {
                                if (i == 0) {
                                    objheader += "<td>" + key + "</td>";
                                }
                                objbody += "<td>" + ret.Result[i][key] + "</td>";
                            }
                            objbody += "</tr>";
                        }
                        objheader += "</tr>";
                        obj = "<table class=\"table table-bordered table-striped\"><thead><tr>" + 
                        objheader + "</tr></thead><tbody>" + objbody + "</tbody></table>";
                        $("#outputdata").html(obj);
                    }
                });
            }
        }
    });
});