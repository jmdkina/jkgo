
function do_drawinit(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute float a_PointSize; \n'+
        'void main() {\n' +
        ' gl_Position = a_Position;\n' +
        ' gl_PointSize = a_PointSize;\n' +
        '}\n';
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

    if (!initShaders(gl, VSHADER_SOURCE_1, FSHADER_SOURCE)) {
    	log_print("failed to initialize shaders");
    	return;
    }
}

function do_drawinit_2(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute vec4 a_Color; \n' +
        'attribute float a_PointSize; \n'+
        'varying vec4 v_Color;\n' +
        'void main() {\n' +
        ' gl_Position = a_Position;\n' +
        ' gl_PointSize = a_PointSize;\n' +
        ' v_Color = a_Color;\n' +
        '}\n';
    var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'varying vec4 v_Color; \n'+
        'void main() {\n' + 
        ' gl_FragColor = v_Color;\n' +
        '}\n';

    if (!initShaders(gl, VSHADER_SOURCE_1, FSHADER_SOURCE)) {
        log_print("failed to initialize shaders");
        return;
    }
}

function do_clearcolor(gl) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
}

function do_vertex_init(gl) {
    var n = initVertexBuffers(gl);
    if (n < 0) {
        log_print("Failed to set vertex");
        return 0;
    }
    return n;
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

function initVertexBuffers_2(gl) {
    // Put position and point size together
    var vertices = new Float32Array([
        0.0, 0.5, 10.0,
       -0.5, -0.5, 20.0,
        0.5, -0.5, 30.0
    ]);
    var n = 3;

    var vertexBuffer = gl.createBuffer();
    if (!vertexBuffer) {
        log_print("Failed to create buffer");
        return -1;
    }
    gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW);

    var FSIZE = vertices.BYTES_PER_ELEMENT;

    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    gl.vertexAttribPointer(a_Position, 2, gl.FLOAT, false, FSIZE * 3, 0);
    gl.enableVertexAttribArray(a_Position);

    var a_PointSize = gl.getAttribLocation(gl.program, 'a_PointSize')
    gl.vertexAttribPointer(a_PointSize, 1, gl.FLOAT, false, FSIZE*3, FSIZE * 2);
    gl.enableVertexAttribArray(a_PointSize);

    return n;
}

function initVertexBuffers_3(gl) {
    // Put position and color together
    var vertices = new Float32Array([
        0.0, 0.5, 1.0, 0.0, 0.0,
       -0.5, -0.5, 0.0, 1.0, 0.0,
        0.5, -0.5, 0.0, 0.0, 1.0
    ]);
    var n = 3;

    var vertexBuffer = gl.createBuffer();
    if (!vertexBuffer) {
        log_print("Failed to create buffer");
        return -1;
    }
    gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW);

    var FSIZE = vertices.BYTES_PER_ELEMENT;

    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    gl.vertexAttribPointer(a_Position, 2, gl.FLOAT, false, FSIZE * 5, 0);
    gl.enableVertexAttribArray(a_Position);

    var a_Color = gl.getAttribLocation(gl.program, 'a_Color');
    gl.vertexAttribPointer(a_Color, 3, gl.FLOAT, false, FSIZE * 5, FSIZE * 2);
    gl.enableVertexAttribArray(a_Color);

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

function do_draw_ext(gl, num) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.clear(gl.COLOR_BUFFER_BIT);
    gl.drawArrays(gl.POINTS, 0, num);
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


function do_set_pointsize_2(gl) {
    // Only point size
    var sizes = new Float32Array([
        10.0, 20.0, 30.0
    ]);

    var sizeBuffer = gl.createBuffer();
    gl.bindBuffer(gl.ARRAY_BUFFER, sizeBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, sizes, gl.STATIC_DRAW);
    var a_PointSize = gl.getAttribLocation(gl.program, 'a_PointSize')
    gl.vertexAttribPointer(a_PointSize, 1, gl.FLOAT, false, 0, 0);
    gl.enableVertexAttribArray(a_PointSize);
}


function do_todraw(gl, n, currentAngle, modelMratrix, u_ModelMatrix) {
    modelMatrix.setRotate(currentAngle, 0, 0, 1);

    gl.uniformMatrix4fv(u_ModelMatrix, false, modelMatrix.elements);

    gl.clear(gl.COLOR_BUFFER_BIT);

    gl.drawArrays(gl.TRIANGLES, 0, n);
}

var g_last = Date.now();
var ANGLE_STEP = 45.0;
function animate(angle) {
    var now = Date.now();
    var elapsed = now - g_last;
    g_last = now;
    var newAngle = angle + (ANGLE_STEP * elapsed) / 1000.0;
    return newAngle %= 360;
}

var currentAngle = 0.0;
var modelMatrix = new Matrix4();

var global_gl;
var u_ModelMatrix;

var tick = function() {
    var n = 3;
    currentAngle = animate(currentAngle);
    do_todraw(global_gl, n, currentAngle, modelMatrix, u_ModelMatrix);
    requestAnimationFrame(tick);
}

/*
 * demo
 */

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

function test_animate(gl) {
    do_vertex_init(gl);
    do_clearcolor(gl);

    global_gl = gl;
    u_ModelMatrix = gl.getUniformLocation(gl.program, 'u_xformMatrix');
    tick();
}

function test_diff_size_1(gl) {
    var n = do_vertex_init(gl);
    do_set_pointsize_2(gl);
    log_print("draw with point " + n);
    do_draw_ext(gl, n);
}

function test_diff_size(gl) {
    var n = initVertexBuffers_2(gl);
    log_print("draw with point " + n);
    do_draw_ext(gl, n);
}

function test_color(gl) {
    do_drawinit_2(gl);
    var n = initVertexBuffers_3(gl);
    do_set_pointsize(gl, 50.0);
    log_print("Draw with point " + n);
    // do_draw_ext(gl, n);
    do_draw_num(gl, n);
}

function test_functions(gl) {
    // test_diff_size(gl);
    test_color(gl);
}
