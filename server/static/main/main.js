let data;
let gisKey;
let userSubgroup;
let userTrust;
let notAppointedPoints = [];
let appointedPoints = [];
let map = null;
let clusterAppointed = null;
let clusterNotAppointed = null;
let shown = false;
let currentDataJson = null;


async function getPoinst() {
    url = "/main"
    let response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin"
    })
    if (response.ok) {
        try {
            let res = await response.json();
            data = res.points;
            gisKey = res.gisKey;
            userSubgroup = res.subgroup;
            userTrust = res.trust;
            currentDataJson = JSON.stringify(res);
        } catch (error) {
            newNotification(false, error);
        }
        fillPoints();
    } else {
        newNotification(false, `${response.status} ${response.statusText}`);
    }
}

function fillPoints() {
    if (map === null) {
        let animation = document.getElementById("data-loading-animation-container");
        if (animation !== null) {
            animation.remove();
        }
        map = new mapgl.Map("map", {
            center: [37.6156, 55.7522],
            zoom: 10,
            key: gisKey,
            style: "c080bb6a-8134-4993-93a1-5b4d8c36a59b"
        });

        let container = document.getElementById("map");
        let canvas = container.getElementsByTagName("canvas")[0];

        canvas.addEventListener("webglcontextlost",
            (event) => {event.preventDefault();}, false);
        
        canvas.addEventListener("webglcontextrestored", () => {
            map.destroy();
            map = null;
            getPoinst();
        }, false);
    }

    fillAppointedPoints();
    fillNotAppointedPoints();
}

function fillNotAppointedPoints() {
    notAppointedPoints = [];
    data.forEach((element) => {
        if (element.appoint === null && element.owner == "yandex") {
            element["icon"] = `/static/svg/marker.svg`;
            element["anchor"] = [8, 8];
            notAppointedPoints.push(element);
        }
    });

    shown = !shown;
    showNotAppointedPoints();
}

function fillAppointedPoints() {
    appointedPoints = [];
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
    if (clusterAppointed !== null) {
        clusterAppointed.destroy();
    }
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
        if (clusterNotAppointed !== null) {
            clusterNotAppointed.destroy();
        }
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
        if (clusterNotAppointed !== null) {
            clusterNotAppointed.destroy();
            clusterNotAppointed = null;
        }
        shown = false;
    }
}

getPoinst();
reloadData();

document.getElementById("show-hidden").onclick = showNotAppointedPoints;