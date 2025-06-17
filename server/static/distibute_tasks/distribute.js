let data;
let gisKey;
let map;


async function getPoinst() {
    url = "/main"
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

    fillPoints()
}

function fillPoints() {
    data.forEach((element) => {
        marker = new mapgl.Marker(map, element);
        marker.userData = element;
    });
}

getPoinst();