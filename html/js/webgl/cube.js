
function do_drawinit_cube(gl) {
    var VSHADER_SOURCE_1 = 
        'attribute vec4 a_Position;\n' +
        'attribute vec4 a_Color;\n' +
        'uniform mat4 u_MvpMatrix; \n' +
        'attribute vec2 a_TexCoord;\n' +
        'attribute vec2 a_TexCoord1;\n' +
        'varying vec2 v_TexCoord;\n' +
        'varying vec2 v_TexCoord1;\n' +
        'void main() {\n' +
        ' gl_Position = u_MvpMatrix * a_Position;\n' +
        ' v_TexCoord = a_TexCoord;\n' +
        ' v_TexCoord1 = a_TexCoord1;\n' +
        '}\n';
    var FSHADER_SOURCE =
        'precision mediump float;\n'+
        'uniform sampler2D u_Sampler; \n' +
        'uniform sampler2D u_Sampler1; \n' +
        'varying vec2 v_TexCoord; \n' +
        'varying vec2 v_TexCoord1; \n' +
        'void main() {\n' + 
        ' vec4 color0 = texture2D(u_Sampler, v_TexCoord);\n' +
        ' vec4 color1 = texture2D(u_Sampler1, v_TexCoord1);\n' +
        ' gl_FragColor = color0* color1;\n' +
        '}\n';

    if (!initShaders(gl, VSHADER_SOURCE_1, FSHADER_SOURCE)) {
        log_print("failed to initialize shaders");
        return;
    }
}


function initVertexBuffers_cube(gl) {
    // Put position and color together
    var vertices = new Float32Array([
        1.0, 1.0, 1.0,
        -1.0, 1.0, 1.0,
        -1.0, -1.0, 1.0,

        1.0, -1.0, 1.0,
        1.0, -1.0, -1.0,
        1.0, 1.0, -1.0,

        -1.0, 1.0, -1.0,
        -1.0, -1.0, -1.0
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
    gl.vertexAttribPointer(a_Position, 3, gl.FLOAT, false, FSIZE * 3, 0);
    gl.enableVertexAttribArray(a_Position);

    var a_TexCoord = gl.getAttribLocation(gl.program, 'a_TexCoord');
    gl.vertexAttribPointer(a_TexCoord, 3, gl.FLOAT, false, FSIZE * 3, 0);
    gl.enableVertexAttribArray(a_TexCoord);

    var a_TexCoord1 = gl.getAttribLocation(gl.program, 'a_TexCoord1');
    gl.vertexAttribPointer(a_TexCoord1, 1, gl.FLOAT, false, FSIZE * 3, FSIZE * 1);
    gl.enableVertexAttribArray(a_TexCoord1);

    gl.bindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer);
    gl.bufferData(gl.ELEMENT_ARRAY_BUFFER, indices, gl.STATIC_DRAW);

    return indices.length;
}

var nf;
var g_eyeX = 0.0, g_eyeY = 0.0, g_eyeZ = 0.0
function do_draw_cube_extern(gl, n) {
    gl.clearColor(0.0, 0.0, 0.0, 1.0);
    gl.enable(gl.DEPTH_TEST);

    var mvpMtrix = new Matrix4();
    mvpMtrix.setPerspective(30, 1, 1, 100);
    mvpMtrix.lookAt(3, 4, 7, 0, 0, 0, 0, 1, 0);

    var mMatrix = new Matrix4();
    mMatrix.setRotate(g_eyeX, g_eyeY, g_eyeZ, 1);
    mvpMtrix = mvpMtrix.multiply(mMatrix);

    var u_MvpMatrix = gl.getUniformLocation(gl.program, 'u_MvpMatrix');
    gl.uniformMatrix4fv(u_MvpMatrix, false, mvpMtrix.elements);

    gl.clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

    gl.drawElements(gl.TRIANGLES, n, gl.UNSIGNED_BYTE, 0);
    nf.innerHTML = 'g_eyeX: ' + g_eyeX + ', g_eyeY: ' + g_eyeY;
}

function keydown(ev, gl, n) {
    switch (ev.keyCode) {
        case 39:
        g_eyeX += 1;
        break;
        case 37:
        g_eyeX -= 1;
        break;
        case 38:
        g_eyeY += 1;
        break;
        case 40:
        g_eyeY -= 1;
        break;
        default:break;
    }

    do_draw_cube_extern(gl, n);
}

var g_texUnit0 = false, g_texUnit1 = false;
function loadTexture_cube(gl, n, texture, u_Sampler, image, texUnit) {
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
        nf = document.getElementById('nearFar');
        document.onkeydown = function(ev) {
            keydown(ev, gl, n);
        }
        do_draw_cube_extern(gl, n);
    }
}


function do_draw_cube(gl) {
    var n = initVertexBuffers_cube(gl);

    var texture = gl.createTexture();
    var texture1 = gl.createTexture();

    var u_Sampler = gl.getUniformLocation(gl.program, 'u_Sampler');
    var u_Sampelr1 = gl.getUniformLocation(gl.program, 'u_Sampelr1');
    var image = new Image();
    var image1 = new Image();
    image.onload = function() {
        log_print("Start to load image " + image.src);
        loadTexture_cube(gl, n, texture, u_Sampler, image, 0);
    }
    image.src = 'images/1.jpg';

    image1.onload = function() {
        log_print("Start to load image: " + image1.src);
        loadTexture_cube(gl, n, texture, u_Sampelr1, image1, 1);
    }
    image1.src = 'images/2.jpg';

    log_print("do init_texture");
    return true;
}

function do_cube(gl) {
    do_drawinit_cube(gl);
    do_draw_cube(gl);
}
