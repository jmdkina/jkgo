
function do_drawinit(gl) {
	var VSHADER_SOURCE = 
	    'attribute vec4 a_Position;\n' +
	    'void main() {\n' +
	    ' gl_Position = a_Position;\n' +
	    ' gl_PointSize = 20.0;\n' +
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

function do_set_position(gl) {
	var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    if (a_Position < 0) {
    	log_print("failed to get the storage location of a_Position");
    	return;
    }
    gl.vertexAttrib3f(a_Position, 0.0, 0.0, 0.0);
}

function do_draw(gl) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.clear(gl.COLOR_BUFFER_BIT);
    gl.drawArrays(gl.POINTS, 0, 1);
}