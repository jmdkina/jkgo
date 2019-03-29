log_enable(true)

var remote_url = "sctek";

$(function(){
    // Vue.options.delimiters = ['{$', '#}'];
    var ws = new Vue({
        el: '#sctek',
        delimiters: ['${', '}'],
        data: {
            cmd: "",
        },
        methods: {
            start_server: function() {
            }
        }
    });
});