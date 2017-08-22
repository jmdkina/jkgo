
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

function do_drawinit_image(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute vec2 a_TexCoord;\n' +
        'varying vec2 v_TexCoord;\n' +
        'void main() {\n' +
        ' gl_Position = a_Position;\n' +
        ' v_TexCoord = a_TexCoord;\n' +
        '}\n';
    var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'uniform sampler2D u_Sampler; \n' +
        'uniform sampler2D u_Sampler1;\n' +
        'varying vec2 v_TexCoord; \n' +
        'void main() {\n' + 
        ' vec4 color0 = texture2D(u_Sampler, v_TexCoord);\n' +
        ' vec4 color1 = texture2D(u_Sampler1, v_TexCoord);\n' +
        ' gl_FragColor = color0 * color1;\n' +
        '}\n';

    if (!initShaders(gl, VSHADER_SOURCE_1, FSHADER_SOURCE)) {
        log_print("failed to initialize shaders");
        return;
    }
}

function do_drawinit_matrix(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute vec4 a_Color;\n' +
        'uniform mat4 u_ModelMatrix; \n' +
        'varying vec4 v_Color;\n' +
        'void main() {\n' +
        ' gl_Position = u_ModelMatrix * a_Position;\n' +
        ' v_Color = a_Color;\n' +
        '}\n';
    var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'varying vec4 v_Color;\n' +
        'void main() {\n' + 
        ' gl_FragColor = v_Color;\n' +
        '}\n';

    if (!initShaders(gl, VSHADER_SOURCE_1, FSHADER_SOURCE)) {
        log_print("failed to initialize shaders");
        return;
    }
}

function do_drawinit_matrix_otheo(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute vec4 a_Color;\n' +
        'uniform mat4 u_ProjMatrix; \n' +
        'varying vec4 v_Color;\n' +
        'void main() {\n' +
        ' gl_Position = u_ProjMatrix * a_Position;\n' +
        ' v_Color = a_Color;\n' +
        '}\n';
    var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'varying vec4 v_Color;\n' +
        'void main() {\n' + 
        ' gl_FragColor = v_Color;\n' +
        '}\n';

    if (!initShaders(gl, VSHADER_SOURCE_1, FSHADER_SOURCE)) {
        log_print("failed to initialize shaders");
        return;
    }
}

function do_drawinit_matrix_more(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute vec4 a_Color;\n' +
        'uniform mat4 u_ModelMatrix;\n' +
        'uniform mat4 u_ViewMatrix; \n' +
        'uniform mat4 u_ProjMatrix; \n' +
        'varying vec4 v_Color;\n' +
        'void main() {\n' +
        ' gl_Position = u_ProjMatrix * u_ViewMatrix * u_ModelMatrix * a_Position;\n' +
        ' v_Color = a_Color;\n' +
        '}\n';
    var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'varying vec4 v_Color;\n' +
        'void main() {\n' + 
        ' gl_FragColor = v_Color;\n' +
        '}\n';

    if (!initShaders(gl, VSHADER_SOURCE_1, FSHADER_SOURCE)) {
        log_print("failed to initialize shaders");
        return;
    }
}

function do_drawinit_cube(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute vec4 a_Color;\n' +
        'uniform mat4 u_MvpMatrix; \n' +
        'varying vec4 v_Color;\n' +
        'void main() {\n' +
        ' gl_Position = u_MvpMatrix * a_Position;\n' +
        ' v_Color = a_Color;\n' +
        '}\n';
    var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'varying vec4 v_Color;\n' +
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

function initVertexBuffers_image(gl) {
    // Put position and color together
    var vertices = new Float32Array([
       -0.5, 0.5, 0.0, 1.0,
       -0.5, -0.8, 0.0, 0.0,
        0.5, 0.5, 1.0, 1.0,
        0.5, -0.8, 1.0, 0.0,
    ]);

    var n = 4;

    var vertexBuffer = gl.createBuffer();
    if (!vertexBuffer) {
        log_print("Failed to create buffer");
        return -1;
    }
    gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW);

    var FSIZE = vertices.BYTES_PER_ELEMENT;

    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    gl.vertexAttribPointer(a_Position, 2, gl.FLOAT, false, FSIZE * 4, 0);
    gl.enableVertexAttribArray(a_Position);

    var a_TexCoord = gl.getAttribLocation(gl.program, 'a_TexCoord');
    gl.vertexAttribPointer(a_TexCoord, 2, gl.FLOAT, false, FSIZE * 4, FSIZE * 2);
    gl.enableVertexAttribArray(a_TexCoord);

    return n;
}

function initVertexBuffers_matrix(gl) {
    // Put position and color together
    /*
    var vertices = new Float32Array([
        0.0, 0.5, -0.4, 0.4, 1.0, 0.4,
        -0.5, -0.5, -0.4, 0.4, 1.0, 0.4,
        0.5, -0.5, -0.4, 1.0, 0.4, 0.4,

        0.5, 0.4, -0.2, 1.0, 0.4, 0.4,
        -0.5, 0.4, -0.2, 1.0, 1.0, 0.4,
        0.0, -0.6, -0.2, 1.0, 1.0, 0.4,

        0, 0.5, 0.0, 0.4, 0.4, 1.0,
        -0.5, -0.25, 0.0, 0.4, 0.4, 1.0,
        0.5, -0.25, 0.0, 1.0, 0.4, 0.4
    ]);
    */
    var vertices = new Float32Array([
        0.0, 1, -0.4, 0.4, 1.0, 0.4,
        -0.5, -0.5, -0.4, 0.4, 1.0, 0.4,
        0.25, -0.5, -0.4, 1.0, 0.4, 0.4,

        0.15, 0.5, -0.2, 1.0, 0.4, 0.4,
        -0.05, -0.25, -0.2, 1.0, 1.0, 0.4,
        0.25, -0.25, -0.2, 1.0, 1.0, 0.4,

        0.2, 0.25, 0.0, 0.4, 0.4, 1.0,
        0.1, -0.1, 0.0, 0.4, 0.4, 1.0,
        0.3, -0.1, 0.0, 1.0, 0.4, 0.4
    ]);
    var n = 9;

    var vertexBuffer = gl.createBuffer();
    if (!vertexBuffer) {
        log_print("Failed to create buffer");
        return -1;
    }
    gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW);

    var FSIZE = vertices.BYTES_PER_ELEMENT;

    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    gl.vertexAttribPointer(a_Position, 3, gl.FLOAT, false, FSIZE * 6, 0);
    gl.enableVertexAttribArray(a_Position);

    var a_Color = gl.getAttribLocation(gl.program, 'a_Color');
    gl.vertexAttribPointer(a_Color, 3, gl.FLOAT, false, FSIZE * 6, FSIZE * 3);
    gl.enableVertexAttribArray(a_Color);

    return n;
}

function initVertexBuffers_cube(gl) {
    // Put position and color together
    var vertices = new Float32Array([
        1.0, 1.0, 1.0, 0.4, 1.0, 0.4,
        -1.0, 1.0, 1.0, 0.4, 1.0, 0.4,
        -1.0, -1.0, 1.0, 1.0, 0.4, 0.4,

        1.0, -1.0, 1.0, 1.0, 0.4, 0.4,
        1.0, -1.0, -1.0, 1.0, 1.0, 0.4,
        1.0, 1.0, -1.0, 1.0, 1.0, 0.5,

        -1.0, 1.0, -1.0, 1.0, 1.0, 0.4,
        -1.0, -1.0, -1.0, 0.4, 0.4, 1.0,
    ]);

    var indices = new Uint8Array ([
        0, 1, 2, 0, 2, 3,
        0, 3, 4, 0, 4, 5,
        0, 5, 6, 0, 6, 1, 
        1, 6, 7, 1, 7, 2,
        7, 4, 3, 7, 3, 2,
        4, 7, 6, 4, 6, 5
    ]);

    var vertexBuffer = gl.createBuffer();
    var indexBuffer = gl.createBuffer();
    if (!vertexBuffer) {
        log_print("Failed to create buffer");
        return -1;
    }
    gl.bindBuffer(gl.ARRAY_BUFFER, vertexBuffer);
    gl.bufferData(gl.ARRAY_BUFFER, vertices, gl.STATIC_DRAW);

    var FSIZE = vertices.BYTES_PER_ELEMENT;

    var a_Position = gl.getAttribLocation(gl.program, 'a_Position');
    gl.vertexAttribPointer(a_Position, 3, gl.FLOAT, false, FSIZE * 6, 0);
    gl.enableVertexAttribArray(a_Position);

    var a_Color = gl.getAttribLocation(gl.program, 'a_Color');
    gl.vertexAttribPointer(a_Color, 3, gl.FLOAT, false, FSIZE * 6, FSIZE * 3);
    gl.enableVertexAttribArray(a_Color);

    gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer);
    gl.bufferData(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW);

    return indices.length;
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

function do_draw_strip(gl, num) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.clear(gl.COLOR_BUFFER_BIT);
    gl.drawArrays(gl.TRIANGLE_STRIP, 0, num);
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


function do_init_textures(gl, n) {
    var texture = gl.createTexture();
    var texture1 = gl.createTexture();

    var u_Sampler = gl.getUniformLocation(gl.program, 'u_Sampler');
    var u_Sampelr1 = gl.getUniformLocation(gl.program, 'u_Sampelr1');
    var image = new Image();
    var image1 = new Image();
    image.onload = function() {
        log_print("Start to load image " + image.src);
        loadTexture(gl, n, texture, u_Sampler, image, 0);
    }
    image.src = 'images/1.jpg';

    image1.onload = function() {
        log_print("Start to load image: " + image1.src);
        loadTexture(gl, n, texture, u_Sampelr1, image1, 1);
    }
    image1.src = 'images/2.jpg';
    log_print("do init_texture");
    return true;
}

var g_texUnit0 = false, g_texUnit1 = false;
function loadTexture(gl, n, texture, u_Sampler, image, texUnit) {
    gl.pixelStorei(gl.UNPACK_FLIP_Y_WEBGL, 1);

    if (texUnit == 0) {
        gl.activeTexture(gl.TEXTURE0);
        g_texUnit0 = true;
    } else {
        gl.activeTexture(gl.TEXTURE1);
        g_texUnit1 = true;
    }

    gl.bindTexture(gl.TEXTURE_2D, texture);

    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR);
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR);
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE);
    gl.texParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE);
    gl.texImage2D(gl.TEXTURE_2D, 0, gl.RGB, gl.RGB, gl.UNSIGNED_BYTE, image);

    gl.uniform1i(u_Sampler, texUnit);

    log_print("draw strip " + n);
    if (g_texUnit0 && g_texUnit1) {
        do_draw_strip(gl, n);
    }
}

var g_eyeX = 0.20, g_eyeY = 0.25, g_eyeZ = 0.25
function keydown(ev, gl, n, u_ModelMatrix, viewMatrix) {
    if (ev.keyCode == 39) {
        g_eyeX += 0.01;
    } else if (ev.keyCode == 37) {
        g_eyeX -= 0.01;
    } else {
        return;
    }
    do_draw_matrix_extern(gl, n, u_ModelMatrix, viewMatrix);
}

var g_near = 0.0, g_far = 0.5;
var nf;
function keydown_ortho(ev, gl, n, u_ProjMatrix, projMatrix, nf) {
    switch(ev.keyCode) {
        case 39: g_near += 0.01; break;
        case 37: g_near -= 0.01; break;
        case 38: g_far += 0.01; break;
        case 40: g_far -= 0.01; break;
        default: return;
    }
    do_draw_matrix_ortho_e(gl, n, u_ProjMatrix, projMatrix, nf);
}

function do_draw_matrix_extern(gl, n, u_ModelMatrix, viewMatrix) {
    viewMatrix.setLookAt(g_eyeX, g_eyeY, g_eyeZ, 0.0, -1.0, -1.0, 0, 1, 0);

    var mMatrix = new Matrix4();
    mMatrix.setRotate(0, 0, 0, 1);

    var modelMatrix = viewMatrix.multiply(mMatrix);

    gl.uniformMatrix4fv(u_ModelMatrix, false, modelMatrix.elements);

    do_draw_num(gl, n);
}

function do_draw_matrix_ortho_e(gl, n, u_ProjMatrix, viewMatrix, nf) {

    viewMatrix.setOrtho(-1, 1, -1, 1, g_near, g_far);
    gl.uniformMatrix4fv(u_ProjMatrix, false, viewMatrix.elements);

    nf.innerHTML = 'near: ' + Math.round(g_near * 100)/100 + ', far: ' + Math.round(g_far*100)/100;
    do_draw_num(gl, n);
}

function do_draw_matrix(gl) {
    var n = initVertexBuffers_matrix(gl);

    var u_ModelMatrix = gl.getUniformLocation(gl.program, 'u_ModelMatrix');
    var viewMatrix = new Matrix4();

    document.onkeydown = function(ev) {
        keydown(ev, gl, n, u_ModelMatrix, viewMatrix);
    }

    do_draw_matrix_extern(gl, n, u_ModelMatrix, viewMatrix);
}

function do_draw_matrix_ortho(gl) {
    var n = initVertexBuffers_matrix(gl);
    nf = document.getElementById('nearFar');

    var u_ProjMatrix = gl.getUniformLocation(gl.program, 'u_ProjMatrix');
    var viewMatrix = new Matrix4();

    document.onkeydown = function(ev) {
        keydown_ortho(ev, gl, n, u_ProjMatrix, viewMatrix, nf);
    }

    do_draw_matrix_ortho_e(gl, n, u_ProjMatrix, viewMatrix, nf);
}

function do_draw_matrix_more(gl, canvas) {
    var n = initVertexBuffers_matrix(gl);

    var u_ModelMatrix = gl.getUniformLocation(gl.program, 'u_ModelMatrix');
    var u_ViewMatrix = gl.getUniformLocation(gl.program, 'u_ViewMatrix');
    var u_ProjMatrix = gl.getUniformLocation(gl.program, 'u_ProjMatrix');

    var modelMatrix = new Matrix4();
    var viewMatrix = new Matrix4();
    var projMatrix = new Matrix4();

    modelMatrix.setTranslate(0.75, 0, 0);
    viewMatrix.setLookAt(0, 0, 0, 0, -1, -1, 0, 1, 0);
    projMatrix.setPerspective(30, canvas.width/canvas.height, 1, 100);

    gl.uniformMatrix4fv(u_ModelMatrix, false, modelMatrix.elements);
    gl.uniformMatrix4fv(u_ViewMatrix, false, viewMatrix.elements);
    gl.uniformMatrix4fv(u_ProjMatrix, false, projMatrix.elements);

    do_draw_num(gl, n);

    modelMatrix.setTranslate(-0.75, 0, 0);
    gl.uniformMatrix4fv(u_ModelMatrix, false, modelMatrix.elements);

    gl.drawArrays(gl.TRIANGLES, 0, n);
}

function do_draw_cube(gl) {
    var n = initVertexBuffers_cube(gl);

    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.enable(gl.DEPTH_TEST);

    var mvpMtrix = new Matrix4();
    mvpMtrix.setPerspective(30, 1, 1, 100);
    mvpMtrix.lookAt(3, 3, 7, 0, 0, 0, 0, 1, 0);

    var u_MvpMatrix = gl.getUniformLocation(gl.program, 'u_MvpMatrix');
    gl.uniformMatrix4fv(u_MvpMatrix, false, mvpMtrix.elements);

    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

    gl.drawElements(gl.TRIANGLES, n, gl.UNSIGNED_BYTE, 0);
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

function test_image(gl) {
    do_drawinit_image(gl);
    var n = initVertexBuffers_image(gl);
    do_init_textures(gl, n);
    // log_print("Draw with point " + n);
    // do_draw_strip(gl, n);
}

function test_matrix(gl) {
    do_drawinit_matrix(gl);
    do_draw_matrix(gl);
}

function test_matrix_ortho(gl) {
    do_drawinit_matrix_otheo(gl);
    do_draw_matrix_ortho(gl);
}

function test_matrix_more(gl) {
    do_drawinit_matrix(gl);
    do_draw_matrix_more(gl, canvas);
}

function test_cube(gl) {
    do_drawinit_cube(gl);
    do_draw_cube(gl);
}

function test_functions(gl, canvas) {
    // test_diff_size(gl);
    // test_color(gl);
    // test_image(gl);
    // test_matrix(gl);
    // test_matrix_ortho(gl);
    // test_matrix_more(gl, canvas);
    test_cube(gl);
}
