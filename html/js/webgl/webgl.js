
function do_drawpointer(gl) {
	var VSHADER_SOURCE = 
	    'void main() {\n' +
	    ' gl_Position = vec4(0.0, 0.0, 0.0, 1.0);\n' +
	    ' gl_PointSize = 20.0;\n' +
	    '}\n';
	var FSHADER_SOURCE =
	    'void main() {\n' + 
	    ' gl_FragColor = vec4(1.0, 0.0, 0.0, 1.0);\n' +
	    '}\n';

    if (!initShaders(gl, VSHADER_SOURCE, FSHADER_SOURCE)) {
    	console.log("failed to initialize shaders");
    	return;
    }

    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.clear(gl.COLOR_BUFFER_BIT);
    gl.drawArrays(gl.POINTS, 0, 1);
}