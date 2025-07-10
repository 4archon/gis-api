let data;
let gisKey;
let map;
let cluster;
let filteredPoints;


async function getPoinst() {
    url = "/distribute_tasks"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin"
    })
    res = await response.json();
    data = res.points;
    gisKey = res.gisKey;
    createMap();
}


function createMap() {
    map = new mapgl.Map("map", {
        center: [37.6156, 55.7522],
        zoom: 10,
        key: gisKey,
        style: "c080bb6a-8134-4993-93a1-5b4d8c36a59b"
    });

    map.on("click", (event) => {
        console.log(event.lngLat);
    });

    fillPoints()
}

function fillPoints() {
    cluster = new mapgl.Clusterer(map, {
        radius: 60,
        clusterStyle: {
        icon: '/static/svg/cluster.svg',
        labelColor: '#ffffff',
        labelFontSize: 16,
        }
    });

    filteredPoints = data;
    cluster.load(filteredPoints);
    cluster.on("click", clusterClick);
}

getPoinst();