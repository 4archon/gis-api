async function getPoints(array_points) {
    url = "http://127.0.0.1:56001/points"
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

function create_window(row_json, parent) {
    newEl = document.createElement('div');
    newEl.className = "flex-item2";
    list = document.createElement('ul');
    newEl.appendChild(list)

    list_row = document.createElement('li');
    list.appendChild(list_row)
    tx = document.createTextNode("Идентификатор: " + row_json.id);
    list_row.appendChild(tx);

    list_row = document.createElement('li');
    list.appendChild(list_row)
    tx = document.createTextNode("Адрес: " + row_json.address);
    list_row.appendChild(tx);

    list_row = document.createElement('li');
    list.appendChild(list_row)
    tx = document.createTextNode("Дата: " + row_json.date);
    list_row.appendChild(tx);

    list_row = document.createElement('li');
    list.appendChild(list_row)
    tx = document.createTextNode("Количество дуг: " + row_json.amount);
    list_row.appendChild(tx);
    
    list_row = document.createElement('li');
    list.appendChild(list_row)
    tx = document.createTextNode("Изображение: ");
    list_row.appendChild(tx);
    link = document.createElement('a');
    link.href = row_json.img;
    list_row.appendChild(link);
    tx = document.createTextNode(row_json.img);
    link.appendChild(tx);

    parent.appendChild(newEl);
}

function handler_json(data_json) {
    parent = document.getElementById('flex2')
    data_json.forEach(element => {
        create_window(element, parent);
    });
}

function clear_elements() {
    elements = document.getElementsByClassName("flex-item2");
    while (elements.length > 0) elements[0].remove();
}

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
    document.getElementById("flex1").style.width = "100%";
    document.getElementById("container").style.display = "none";
    document.getElementById("flex1").style.display = "flex";
});

map.on('click', (event) => {
    document.getElementById("flex1").style.display = "none";
});

function click_close() {
    document.getElementById("flex1").style.display = "none";
    document.getElementById("container").style.display = "flex";
}

document.getElementById("close").onclick = click_close;