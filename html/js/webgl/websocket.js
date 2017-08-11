
function init_websocket(str) {
	var vs = new WebSocket(str);
	return vs;
}


$(function(){
    var ws;
    ws = init_websocket("ws://localhost:12306/websocket");

    ws.onopen = function(evt) {
        log_print("websocket open");
    }
    ws.onclose = function(evt) {
        log_print("websocket close");
        ws = null;
    }
    ws.onmessage = function(evt) {
        log_print("websocket response: " + evt.data);
    }
    ws.onerror = function(evt) {
        log_print("ERROR: " + evt.data);
    }
    
    var ws_option = new Vue({
        el: '#ws',
        data: {
            message: ""
        },
        methods: {
            send_msg: function() {
                var content = this.message;
                log_print("send content: " + content);
                ws.send(content);
            }
        }
    });
});
