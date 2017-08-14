
function do_drawinit(gl) {
	var VSHADER_SOURCE = 
	    'attribute vec4 a_Position;\n' +
        'attribute float a_PointSize; \n'+
	    'void main() {\n' +
	    ' gl_Position = a_Position;\n' +
	    ' gl_PointSize = a_PointSize;\n' +
	    '}\n';
	var FSHADER_SOURCE =
	    'void main() {\n' + 
	    ' gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);\n' +
	    '}\n';

    if (!initShaders(gl, VSHADER_SOURCE, FSHADER_SOURCE)) {
    	log_print("failed to initialize shaders");
    	return;
    }
}

function do_set_position(gl, p1, p2, p3) {
	var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    if (a_Position < 0) {
    	log_print("failed to get the storage location of a_Position");
    	return;
    }
    gl.vertexAttrib3f(a_Position, p1, p2, p3);
}

function do_set_pointsize(gl, size) {
    var a_PointSize = gl.getAttribLocation(gl.program, 'a_PointSize');
    gl.vertexAttrib1f(a_PointSize, size);
}

function do_draw(gl) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.clear(gl.COLOR_BUFFER_BIT);
    gl.drawArrays(gl.POINTS, 0, 1);
}

function glclick(ev, gl, canvas, g_points) {
    // var g_points = [];

    var x = ev.clientX;
    var y = ev.clientY;
    var rect = ev.target.getBoundingClientRect();
    x = ((x - rect.left) - canvas.width/2)/(canvas.width/2);
    y = (canvas.height/2 - (y - rect.top))/(canvas.height/2);
    g_points.push(x); g_points.push(y);

    gl.clear(gl.COLOR_BUFFER_BIT);

    var len = g_points.length;
    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    for (var i = 0; i < len; i+=2) {
        gl.vertexAttrib3f(a_Position, g_points[i], g_points[i+1], 0.0);
        gl.drawArrays(gl.POINTS, 0, 1);
    }
}