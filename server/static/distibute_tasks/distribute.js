const map = new mapgl.Map('map', {
    center: [37.6156, 55.7522],
    zoom: 10,
    key: gisKey,
    style: 'c080bb6a-8134-4993-93a1-5b4d8c36a59b'
});

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
    clear_elements();
    getPoints(array_points);
});

function clear_elements() {
    elements = document.getElementsByClassName("card mb-3");
    while (elements.length > 0) elements[0].remove();
}

async function getPoints(array_points) {
    url = "/points"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "text/plain"
        },
        body: array_points
    })
    res = await response.json();
    handler_json(res);
}

function handler_json(data_json) {
    parent = document.getElementsByClassName('offcanvas-body')[0]
    data_json.forEach(element => {
        create_card(element, parent);
    });
}

function create_card(row_json, parent) {
    new_card = document.createElement("div");
    new_card.className = "card mb-3";
    new_card.style.maxWidth = "540px";
    parent.appendChild(new_card);

    new_row = document.createElement("div");
    new_row.className = "row g-0";
    new_card.appendChild(new_row);

    new_col = document.createElement("div");
    new_col.className = "col-md-4";
    new_row.appendChild(new_col);

    new_img = document.createElement("img");
    new_img.className = "img-fluid rounded-start";
    new_img.src = row_json.img;
    new_col.appendChild(new_img);

    new_col_body = document.createElement("div");
    new_col_body.className = "col-md-8";
    new_row.appendChild(new_col_body);

    new_body = document.createElement("div");
    new_body.className = "card-body";
    new_col_body.appendChild(new_body);

    new_title = document.createElement("h5");
    new_title.className = "card-title";
    tx = document.createTextNode("ID точки: " + row_json.id);
    new_title.appendChild(tx);
    new_body.appendChild(new_title);

    new_id = document.createElement("div")
    new_id.className = "card-text";
    tx = document.createTextNode("Адрес: " + row_json.address);
    new_id.appendChild(tx);
    new_body.appendChild(new_id);

    new_date = document.createElement("div")
    new_date.className = "card-text";
    tx = document.createTextNode("Дата: " + row_json.date);
    new_date.appendChild(tx);
    new_body.appendChild(new_date);

    new_arc_num = document.createElement("div")
    new_arc_num.className = "card-text";
    tx = document.createTextNode("Количество дуг: " + row_json.amount);
    new_arc_num.appendChild(tx);
    new_body.appendChild(new_arc_num);

    el = document.getElementById("offcanvasPoints");
    off_canvas = new bootstrap.Offcanvas(el);
    off_canvas.show();

    el2 = document.getElementById("offcanvasTasks");
    off_canvas_tasks = new bootstrap.Offcanvas(el2);
    off_canvas_tasks.show();
}

map.on('click', (event) => {
    off_canvas.hide();
});
