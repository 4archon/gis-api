let geolocationActivate = false;
let geolocationWatcher = null;
let userPosition;
let userPositionMarker;

document.getElementById("navigation-button").onclick = userGeolocation;

function userGeolocation() {
    if (geolocationActivate) {
        if (navigator.geolocation) {
            navigator.geolocation.clearWatch(geolocationWatcher);
        }
        if (userPositionMarker) {
            userPositionMarker.destroy();
        }
        geolocationActivate = false;
    } else {
        if (navigator.geolocation) {
            geolocationWatcher = navigator.geolocation.
            watchPosition(successPosition, errorPosition);
            console.log(geolocationWatcher);
        } else {
            geoNotification("geolocation is unavailable");
        }
    }
}

function successPosition(position) {
    userPosition = {
        lat: position.coords.latitude,
        long: position.coords.longitude,
        rotation: position.coords.heading === null ? 0 : position.coords.heading,
        coordinates: [position.coords.longitude, position.coords.latitude],
        icon: `/static/svg/cursor.svg`,
        anchor: [16, 16]
    }

    createUserPositionMarker(userPosition);
}

function createUserPositionMarker(position) {
    if (userPositionMarker) {
        userPositionMarker.destroy();
    }

    if (!geolocationActivate) {
        geolocationActivate = true;
        map.setCenter(position.coordinates);
        map.setZoom(14);
    }

    userPositionMarker = new mapgl.Marker(map, position);
}

function errorPosition(err) {
    geoNotification(err.message);
}

function geoNotification(error) {
    let alertContainer = document.createElement("div");
    alertContainer.className = "alert alert-danger my-1";
    alertContainer.innerHTML = `<h6 class="m-0">${error}</h6>`;
    let container = document.getElementById("notification-bar");
    container.appendChild(alertContainer);

    setTimeout(() => {
        alertContainer.remove();
    }, 10000);
}