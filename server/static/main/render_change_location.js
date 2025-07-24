let change_location = new bootstrap.Modal(document.getElementById("change-location"), null);
let change_location_map;
let originalPoitMarker;
let userMarker = null;


document.getElementById("change-location").addEventListener("shown.bs.modal", ()=> {
    change_location_map.invalidateSize();
})

document.addEventListener("DOMContentLoaded", function () {
    document.addEventListener('hide.bs.modal', function (event) {
        if (document.activeElement) {
            document.activeElement.blur();
        }
    });
});

function new_location(event){
    change_location.show();
    let id = event.currentTarget.getAttribute("data-id");
    let found = data.find((el) => el.id == id);
    let center = found.coordinates.map((el) => Number(el));
    change_location_map = new mapgl.Map("map-change-location", {
        center: center,
        zoom: 13,
        key: gisKey,
        style: "c080bb6a-8134-4993-93a1-5b4d8c36a59b"
    });
    originalPoitMarker = new mapgl.Marker(change_location_map, found);
    change_location_map.on("click", click_on_map);
}

function click_on_map(event) {
    let coordinates = event.lngLat;
    console.log(coordinates);
    if (userMarker !== null) {userMarker.destroy()};
    userMarker = new mapgl.Marker(change_location_map, {
        coordinates: coordinates,
        icon: `/static/svg/secondary.svg`,
        anchor: [15, 46]
    })
}

let originalMarkerVisible = true;
document.getElementById("change-location-hide").onclick = hide_original_point;

function hide_original_point() {
    if (originalMarkerVisible) {
        originalPoitMarker.hide();
        originalMarkerVisible = false;
    } else {
        originalPoitMarker.show();
        originalMarkerVisible = true;
    }
}

document.getElementById("change-location-save").onclick = save_new_location;

function save_new_location() {
    if (userMarker !== null) {
        reportData.newLocation = userMarker.getCoordinates();
        change_location.hide();
        change_location_map.destroy();
        render_data_to_form();
    } else {
        alert("Новая точка не указана на карте");
    }
}
