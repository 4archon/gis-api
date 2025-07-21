let data;
let gisKey;
let map;
let notAppointedPoints = [];
let appointedPoints = [];
let shown = false;
let clusterAppointed;
let clusterNotAppointed;
let userSubgroup;
let userTrust;

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
    userSubgroup = res.subgroup;
    userTrust = res.trust;
    console.log(data);
    
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
        if (element.appoint === null && element.owner == "yandex") {
            element["icon"] = `/static/svg/marker.svg`;
            element["anchor"] = [8, 8];
            notAppointedPoints.push(element);
        }
    });
}

function fillAppointedPoints() {
    data.forEach((element) => {
        if (element.appoint !== null) {
            if (element.deadline !== null) {
                element["icon"] = `/static/svg/danger_marker.svg`;
                element["anchor"] = [15, 46]
            }
            appointedPoints.push(element);
        }
    });
    showAppointedPoints();
}

function showAppointedPoints() {
    clusterAppointed = new mapgl.Clusterer(map, {
        radius: 60,
        clusterStyle: clusterAppointedStyle
    });
    clusterAppointed.load(appointedPoints);
    clusterAppointed.on("click", clusterClick);
}

function clusterAppointedStyle(pointsCount, target) {
    let points = target.data;
    if (points.some(el => el.deadline !== null)) {
        return {
            icon: '/static/svg/cluster_danger.svg',
            labelColor: '#ffffff',
            labelFontSize: 16,
        }    
    } else {
        return {
            icon: '/static/svg/cluster.svg',
            labelColor: '#ffffff',
            labelFontSize: 16,
        }
    }
}

function showNotAppointedPoints() {
    if (shown == false) {
        clusterNotAppointed = new mapgl.Clusterer(map, {
            radius: 60,
            clusterStyle: {
                icon: '/static/svg/cluster_not_appointed.svg',
                labelColor: '#ffffff',
                labelFontSize: 16,
            }
        });
        clusterNotAppointed.load(notAppointedPoints);
        clusterNotAppointed.on("click", clusterClick);
        shown = true;
    } else {
        clusterNotAppointed.destroy();
        shown = false;
    }
}

getPoinst();

document.getElementById("show-hidden").onclick = showNotAppointedPoints;