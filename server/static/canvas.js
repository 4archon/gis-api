document.getElementById("button1").onclick = click_canvas;

var canvasOptions = {
    strokeStyle: '#0000ff',
    lineWidth: 4,
    opacity: 0.7
};

function simple(coordinates) {
    var len = coordinates.length;
    var simplecat = Math.ceil(len / 30);
    coordinates = coordinates.filter((value, index) => {
        if (index % simplecat == 0) {
            return value;
        }
    });
    return coordinates;
}

function click_canvas() {
    var canvas = document.querySelector('#draw-canvas');
    var ctx2d = canvas.getContext('2d');
    var drawing = false;
    var coordinates = [];

    var rect = map.getContainer().getBoundingClientRect();
    canvas.style.width = rect.width + 'px';
    canvas.style.height = rect.height + 'px';
    canvas.style.top = rect.top + 2 + 'px';
    canvas.width = rect.width;
    canvas.height = rect.height;

    ctx2d.strokeStyle = canvasOptions.strokeStyle;
    ctx2d.lineWidth = canvasOptions.lineWidth;
    canvas.style.opacity = canvasOptions.opacity;

    ctx2d.clearRect(0, 0, canvas.width, canvas.height);

    canvas.style.display = 'block';

    canvas.onmousedown = function(e) {
        drawing = true;
        coordinates.push([e.offsetX, e.offsetY]);
    };

    canvas.onmousemove = function(e) {
        if (drawing) {
          var last = coordinates[coordinates.length - 1];
          ctx2d.beginPath();
          ctx2d.moveTo(last[0], last[1]);
          ctx2d.lineTo(e.offsetX, e.offsetY);
          ctx2d.stroke();

          coordinates.push([e.offsetX, e.offsetY]);
        }
    };
    
    canvas.onmouseup = function(e) {
        coordinates.push([e.offsetX, e.offsetY]);
        canvas.style.display = 'none';
        drawing = false;
        coordinates = coordinates.map((x) => {
            return [x[0] / canvas.width, x[1] / canvas.height];
        });
        bounds = map.getBounds();
        coordinates = coordinates.map((x) => {
            return [bounds.southWest[0] + x[0] * (bounds.northEast[0] - bounds.southWest[0]),
                bounds.southWest[1] + (1 - x[1]) * (bounds.northEast[1] - bounds.southWest[1]),
            ];
        });

        coordinates = simple(coordinates);

        const polygon = new mapgl.Polygon(map, {
            coordinates: [
                coordinates
            ],
            color: '#00FF0020',
            strokeColor: '#00AB00',
        });

        markers_in = markers.filter((x) => {
            if (inside(x.coordinates, coordinates)) {
                return x;
            }
        });
        id_in = markers_in.map((x) => {
            return x.ID;
        })
        console.log(id_in);

        clear_elements();
        getPoints(id_in);
        document.getElementById("flex1").style.display = "flex";
    }
}

function inside(point, vs) {
    var x = point[0], y = point[1];
    
    var inside = false;
    for (var i = 0, j = vs.length - 1; i < vs.length; j = i++) {
        var xi = vs[i][0], yi = vs[i][1];
        var xj = vs[j][0], yj = vs[j][1];
        
        var intersect = ((yi > y) != (yj > y))
            && (x < (xj - xi) * (y - yi) / (yj - yi) + xi);
        if (intersect) inside = !inside;
    }
    
    return inside;
};

