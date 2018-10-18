log_enable(true)

var remote_url = "wsclient";

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];
    var ws = new Vue({
        el: '#wsclient',
        delimiters: ['${', '}'],
        data: {
            addr: "",
            port: "",
            url: "",
            sendmsg: "",
            recvmsg: ""
        },
        methods: {
            start_server: function() {
                start_server_do(this.addr, this.port, this.url);
            },
            stop_server:function() {
                stop_server();
            },
            send_msg:function() {
                send_msg(this.sendmsg);
            },
            recv_msg: function() {
                recv_msg();
            }
        }
    });
    function start_server_do(addr, port, url) {
        var str = "cmd=start&addr=" + addr + "&port=" + port + "&url=" + url;
        $.ajax({
            url: remote_url,
            method: "POST",
            data: str,
            success: function (response, textStatus) {
                log_print("query new success" + response + ", result:" + textStatus);  
                try {
                    ret = $.parseJSON(response);
                    if (ret.Status != 200) {
                        layer.msg("Fail start websocket " + ret.Result);
                    } else {
                        layer.msg("connect websocket " + ret.Result);
                    }
                } catch(e) {
                    layer.msg("Fail parse response message ", response);
                }
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
                        layer.msg("Fail of " + ret.Result);
                        return;
                    } else {
                        layer.msg("stop wsclient success ", ret.Result);
                    }
            
                } catch(e) {
                    layer.msg("Fail parse , err: " + e);
                }
            }
        });
    }
    
    function send_msg(msg) {
        var str = "cmd=send&msg=" + msg;
        $.ajax({
            url: remote_url,
            method: "POST",
            data: str,
            success: function (response, textStatus) {
                log_print("query new success" + response + ", result:" + textStatus);
                try {
                    ret = $.parseJSON(response);
                    if (ret.Status != 200) {
                        layer.msg("Fail of " + ret.Result);
                        return;
                    } else {
                        layer.msg("wsclient send ", ret.Result);
                    }
            
                } catch(e) {
                    layer.msg("Fail parse , err: " + e);
                }
            }
        });
    }
    function recv_msg() {
        var str = "cmd=recv";
        $.ajax({
            url: remote_url,
            method: "POST",
            data: str,
            success: function (response, textStatus) {
                log_print("query new success" + response + ", result:" + textStatus);
                try {
                    ret = $.parseJSON(response);
                    if (ret.Status != 200) {
                        layer.msg("Fail of " + ret.Result);
                        return;
                    }
                    this.recvmsg = ret["Result"];
            
                } catch(e) {
                    layer.msg("Fail parse , err: " + e);
                }
            }
        });
    }
    

});

