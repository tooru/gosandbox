(function() {
    "use strict";

    var scene = new THREE.Scene();
    var camera = new THREE.PerspectiveCamera(100, window.innerWidth/window.innerHeight, 0.1, 1000);

    var renderer = new THREE.WebGLRenderer();
    renderer.setSize( window.innerWidth, window.innerHeight );
    document.body.appendChild(renderer.domElement);

//    var geometry = new THREE.BoxBufferGeometry(2, 2, 2, 4, 4, 4);
    var geometry = new THREE.SphereGeometry(3, 16, 16);
    var material = new THREE.MeshBasicMaterial({
        color: 0x00ff00,
        wireframe: true,
    });
    var cube = new THREE.Mesh(geometry, material);
    scene.add(cube);

    var gridHelper = new THREE.GridHelper(20, 20);
    scene.add(gridHelper);

    camera.position.x = 0;
    camera.position.y = 1;
    camera.position.z = 7;

    var animate = function () {
        requestAnimationFrame(animate);

        cube.rotation.x += 0.01;
        cube.rotation.y += 0.01;

        renderer.render(scene, camera);
    };

    //animate();
    renderer.render(scene, camera);
})();
