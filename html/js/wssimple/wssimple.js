log_enable(true)

var remote_url = "wssimple";

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];
    var ws = new Vue({
        el: '#wssimple',
        delimiters: ['${', '}'],
        data: {
            cmd: "",
            do: "",
            port: ""
        },
        methods: {
            start_server: function() {
                start_server_do(this.port);
            },
            stop_server:function() {
                stop_server();
            },
            do_stopplaying:function() {
                do_command("stopPlaying");
            },
            do_dataupdate: function() {
                do_command("dataUpdate");
            },
            do_resumeplaying: function() {
                do_command("resumePlaying");
            }
        }
    });
    function start_server_do(port) {
        var str = "cmd=start&port=" + port;
        $.ajax({
            url: remote_url,
            method: "POST",
            data: str,
            success: function (response, textStatus) {
                log_print("query new success" + response + ", result:" + textStatus);   
            }
        });
    }
    
    function stop_server() {
        var str = "cmd=stop"
        $.ajax({
            url: remote_url,
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
            
                } catch(e) {
                    layer.msg("Fail parse , err: " + e);
                }
            }
        });
    }
    
    function do_command(cmd) {
        var str = "cmd=exec&do=" + cmd;
        $.ajax({
            url: remote_url,
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
            
                } catch(e) {
                    layer.msg("Fail parse , err: " + e);
                }
            }
        });
    }
    

});

