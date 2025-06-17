let data;
let gisKey;
let map;
let notAppointedPoints = [];
let shown = false;
let pointProfile = new bootstrap.Modal(document.getElementById("point-profile"), null);
let pointHistory = new bootstrap.Modal(document.getElementById("point-history"), null);
let pointReport = new bootstrap.Modal(document.getElementById("point-report"), null);


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
    fillPoints();
}

function fillPoints() {
    map = new mapgl.Map("map", {
        center: [37.6156, 55.7522],
        zoom: 10,
        key: gisKey,
        style: "c080bb6a-8134-4993-93a1-5b4d8c36a59b"
    });

    fillAppointedPoints();
    fillNotAppointedPoints();
}

function fillNotAppointedPoints() {
    data.forEach((element) => {
        if (!element.appointed) {
            element["icon"] = `/static/svg/marker.svg`;
            // element["anchor"] = [1, 1]; have to
            marker = new mapgl.Marker(map, element);
            marker.userData = element;
            marker.on("click", pointClick);
            marker.hide();
            notAppointedPoints.push(marker);
        }
    });
}

function fillAppointedPoints() {
    data.forEach((element) => {
        if (element.appointed) {
            if (element.deadline !== null) {
                element["icon"] = `/static/svg/danger_marker.svg`;
                element["anchor"] = [15, 46]
            }
            marker = new mapgl.Marker(map, element);
            marker.userData = element;
            marker.on("click", pointClick);
        }
    });
}

function showNotAppointedPoints() {
    if (shown == false) {
        notAppointedPoints.forEach((element) => {
            element.show();
        });
        shown = true;
    } else {
        notAppointedPoints.forEach((element) => {
            element.hide();
        });
        shown = false;
    }
}

getPoinst();

document.getElementById("show-hidden").onclick = showNotAppointedPoints;