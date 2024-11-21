function getPoints(array_points) {

}

const clusterer = new mapgl.Clusterer(map, {
    redius: 120,
});
clusterer.load(markers);

clusterer.on('click', (event) => {
    array_points = [];
    if (event.target.type == "cluster") {
        event.target.data.forEach(element => {
            array_points.push(element.ID);
        });
    }
    if (event.target.type == "marker") {
        array_points.push(event.target.data.ID);
    }
    alert(array_points);
    document.getElementById("flex1").style.display = "flex";
});

map.on('click', (event) => {
    document.getElementById("flex1").style.display = "none";
});

