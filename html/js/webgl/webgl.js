
function do_drawinit(gl) {
	var VSHADER_SOURCE = 
	    'attribute vec4 a_Position;\n' +
        'attribute float a_PointSize; \n'+
        'uniform mat4 u_xformMatrix; \n' +
	    'void main() {\n' +
	    ' gl_Position = a_Position * u_xformMatrix;\n' +
	    ' gl_PointSize = a_PointSize;\n' +
	    '}\n';
	var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'uniform vec4 u_FragColor;\n' +
	    'void main() {\n' + 
	    ' gl_FragColor = u_FragColor;\n' +
	    '}\n';

    if (!initShaders(gl, VSHADER_SOURCE, FSHADER_SOURCE)) {
    	log_print("failed to initialize shaders");
    	return;
    }
}

function do_vertex_init(gl) {
    var n = initVertexBuffers(gl);
    if (n < 0) {
        log_print("Failed to set vertex");
        return ;
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

function initVertexBuffers(gl) {
    var vertices = new Float32Array([
        0.0, 0.5, -0.5, -0.5, 0.5, -0.5
    ]);
    var n = 3;

    var vertexBuffer = gl.createBuffer();
    if (!vertexBuffer) {
        log_print("Failed to create buffer");
        return -1;
    }
    gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW);

    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');

    gl.vertexAttribPointer(a_Position, 2, gl.FLOAT, false, 0, 0);

    gl.enableVertexAttribArray(a_Position);

    return n;
}

function do_rotate(gl) {
    var ANGLE = 90.0;
    var radian = Math.PI * ANGLE /180.0;
    var cosB = Math.cos(radian), sinB = Math.sin(radian);

    var xformMatrix = new Float32Array ([
        cosB, sinB, 0.0, 0.0,
       -sinB, cosB, 0.0, 0.0,
        0.0, 0.0, 1.0, 0.0,
        0.0, 0.0, 0.0, 1.0
    ]);

    var u_xformMatrix = gl.getUniformLocation(gl.program, 'u_xformMatrix');
    gl.uniformMatrix4fv(u_xformMatrix, false, xformMatrix);
}

function do_move(gl) {
    var tx = 0.5, ty = 0.5, tz = 0.0;

    var xformMatrix = new Float32Array ([
        1.0, 0.0, 0.0, 0.0,
        0.0, 1.0, 0.0, 0.0,
        0.0, 0.0, 1.0, 0.0,
        tx, ty, tz, 1.0
    ]);

    var u_xformMatrix = gl.getUniformLocation(gl.program, 'u_xformMatrix');
    gl.uniformMatrix4fv(u_xformMatrix, false, xformMatrix);
}

function do_zoom(gl) {
    var tx = 1.5, ty = 1.5, tz = 0.5;

    var xformMatrix = new Float32Array ([
        tx, 0.0, 0.0, 0.0,
        0.0, ty, 0.0, 0.0,
        0.0, 0.0, tz, 0.0,
        0.0, 0.0, 0.0, 1.0
    ]);

    var u_xformMatrix = gl.getUniformLocation(gl.program, 'u_xformMatrix');
    gl.uniformMatrix4fv(u_xformMatrix, false, xformMatrix);
}

function do_color(gl) {
    var u_FragColor = gl.getUniformLocation(gl.program, 'u_FragColor');
    var rgba = [1.0, 0.0, 1.0, 1.0];
    gl.uniform4f(u_FragColor, rgba[0], rgba[1], rgba[2], rgba[3]);
}

function do_draw(gl) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.clear(gl.COLOR_BUFFER_BIT);
    gl.drawArrays(gl.POINTS, 0, 1);
}

function do_draw_num(gl, num) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.clear(gl.COLOR_BUFFER_BIT);
    gl.drawArrays(gl.TRIANGLES, 0, num);
}

function glclick(ev, gl, canvas, g_points, u_rgba) {
    // var g_points = [];

    var x = ev.clientX;
    var y = ev.clientY;
    var rect = ev.target.getBoundingClientRect();
    x = ((x - rect.left) - canvas.width/2)/(canvas.width/2);
    y = (canvas.height/2 - (y - rect.top))/(canvas.height/2);
    // g_points.push(x); g_points.push(y);
    g_points.push([x,y]);

    if (x >= 0.0 && y >= 0.0) {
        u_rgba.push([1.0, 0.0, 0.0, 1.0]);
    } else if (x < 0.0 && y < 0.0) {
        u_rgba.push([0.0, 1.0, 0.0, 1.0]);
    } else {
        u_rgba.push([1.0, 1.0, 1.0, 1.0]);
    }

    gl.clear(gl.COLOR_BUFFER_BIT);

    var len = g_points.length;
    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    var u_FragColor = gl.getUniformLocation(gl.program, 'u_FragColor');
    for (var i = 0; i < len; i++) {
        var points = g_points[i];
        var rgba = u_rgba[i];
        gl.vertexAttrib3f(a_Position, points[0], points[1], 0.0);
        gl.uniform4f(u_FragColor, rgba[0], rgba[1], rgba[2], rgba[3]);
        gl.drawArrays(gl.POINTS, 0, 1);
    }
}

function test_base(gl) {
    do_set_position(gl);
    do_set_pointsize(gl, 50.0);
    do_draw(gl);
}

function test_triangle(gl) {
    do_vertex_init(gl);
    do_set_pointsize(gl, 50.0);
    do_color(gl);
    do_draw_num(gl, 3);
}

function test_rotate(gl) {
    do_vertex_init(gl);
    do_color(gl);
    // do_rotate(gl);
    // do_move(gl);
    do_zoom(gl);
    do_draw_num(gl, 3);
}

var g_points = [];
var u_rgba = [];
function test_click(gl) {
    do_set_pointsize(gl, 50.0);
    canvas.onmousedown = function(ev) { glclick(ev, gl, canvas, g_points, u_rgba); }
}