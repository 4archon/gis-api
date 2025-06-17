let userMarker = null;

function search(event) {
    let val = event.currentTarget.value;
    const re = /-?\d*\.?\d+/g
    res = val.match(re);
    if (res.length == 2 && res.every((el) => el <= 180 && el >=-180)) {
        res = res.map((el) => Number(el));
        event.currentTarget.classList.remove("is-invalid");
        if (userMarker !== null) {userMarker.destroy()};
        userMarker = new mapgl.Marker(map, {
            coordinates: res,
            icon: `/static/svg/secondary.svg`,
            anchor: [15, 46]
        })
        map.setCenter(res);
        map.setZoom(13);
    } else {
        event.currentTarget.classList.add("is-invalid");
    }
}

document.getElementById("search").onchange = search;
