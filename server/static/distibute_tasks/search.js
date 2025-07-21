let userMarker = null;

function search(event) {
    let val = event.currentTarget.value;
    const reCoordinates = /-?\d+\.?\d+/g;
    res = val.match(reCoordinates);
    if (res.length == 2 && res.every((el) => el <= 180 && el >=-180)) {
        showSearchedCoordinates(res);
        event.currentTarget.classList.remove("is-invalid");
    } else {
        const reID = /[yw]?\d+/;
        res = val.match(reID);
        if (res.length == 1 && res.every((el) => el.length > 1 &&
        (el[0] == "w" || el[0] == "y"))) {
            found = findPointByExternalID(res[0]);
            event.currentTarget.classList.remove("is-invalid");
            if (!found) {
                event.currentTarget.classList.add("is-invalid");
            }
        } else if (res.length == 1) {
            found = findPointByID(res[0]);
            event.currentTarget.classList.remove("is-invalid");
            if (!found) {
                event.currentTarget.classList.add("is-invalid");
            }
        } else {
            event.currentTarget.classList.add("is-invalid");
        }
    }
}

function showSearchedCoordinates(coordinates) {
    coordinates = coordinates.map((el) => Number(el));
    coordinates = coordinates.reverse();
    if (userMarker !== null) {userMarker.destroy()};
    userMarker = new mapgl.Marker(map, {
        coordinates: coordinates,
        icon: `/static/svg/secondary.svg`,
        anchor: [15, 46]
    });
    map.setCenter(coordinates);
    map.setZoom(13);
}

function findPointByExternalID(externalID) {
    found = data.find((el) => el.externalID == externalID);
    console.log(found);
    if (found === undefined) {
        return false;
    } else {
        map.setCenter(found.coordinates.map((el) => Number(el)));
        map.setZoom(18);
        return true;
    }
}

function findPointByID(id) {
    found = data.find((el) => el.id == id);
    console.log(found);
    if (found === undefined) {
        return false;
    } else {
        map.setCenter(found.coordinates.map((el) => Number(el)));
        map.setZoom(18);
        return true;
    }
}

document.getElementById("search").onchange = search;
