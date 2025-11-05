let newPointModal = new bootstrap.Modal(document.getElementById("point-new"));

function newUserMarker(event) {
    if (userMarker !== null) {userMarker.destroy();}
    userMarker = new mapgl.Marker(map, {
        coordinates: event.lngLat,
        icon: `/static/svg/secondary.svg`,
        anchor: [15, 46],
        userData: [...event.lngLat]
    });

    userMarker.on("click", newPointMenu);
}


function newPointMenu(event) {
    let long = event.targetData.userData[0];
    let lat = event.targetData.userData[1];
    render_new_point_body(long, lat);
    render_new_point_footer();
    newPointModal.show();
}

function render_new_point_footer() {
    let container = document.getElementById("point-new-footer");
    container.innerHTML = 
    `
    <button type="button" class="btn btn-secondary" onclick="closeNewPointMenu()">
        Отменить
    </button>
    <button type="button" class="btn btn-primary" onclick="sendNewPoint()">
        Добавить
    </button>
    `
}

function closeNewPointMenu() {
    newPointModal.hide();
}

function sendNewPoint() {
    let newPoint = {
        lat: Number(document.getElementById("inputNewPoint-Lat").value),
        long: Number(document.getElementById("inputNewPoint-Long").value),
        address: document.getElementById("inputNewPoint-Address").value,
        district: document.getElementById("inputNewPoint-District").value,
        externalID: document.getElementById("inputNewPoint-ExternalID").value,
        carpet: document.getElementById("inputNewPoint-Carpet").value,
        numberArc: Number(document.getElementById("inputNewPoint-NumberArc").value),
        arcType: document.getElementById("inputNewPoint-ArcType").value,
        owner: document.getElementById("inputNewPoint-Owner").value,
        operator: document.getElementById("inputNewPoint-Operator").value,
        customer: document.getElementById("inputNewPoint-Customer").value,
        comment: document.getElementById("inputNewPoint-Comment").value,
    };

    let newPoints = {
        newPoints: [newPoint]
    }
    postNewPoints(newPoints);
}

async function postNewPoints(data) {
    let url = "/new_points"
    let response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    let res = await response;
    if (res.ok) {
        newPointModal.hide();
        if (userMarker !== null) {userMarker.destroy();}
        await getPoinst();
    }
}

function render_new_point_body(long, lat) {
    let container = document.getElementById("point-new-body");
    container.innerHTML = 
    `
    <div class="card-body mb-4">
        <div class="row g-3">
            <div class="col-md-6">
                <label for="inputNewPoint-Lat" class="form-label">Широта</label>
                <input id="inputNewPoint-Lat"
                value="${lat}" 
                class="form-control"  placeholder="Введите широту">
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-Long" class="form-label">Долгота</label>
                <input id="inputNewPoint-Long"
                value="${long}" 
                class="form-control"  placeholder="Введите долготу">
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-Address" class="form-label">Адрес</label>
                <input id="inputNewPoint-Address"
                class="form-control"  placeholder="Введите адрес">
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-District" class="form-label">Округ</label>
                <input id="inputNewPoint-District"
                class="form-control"  placeholder="Введите округ">
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-ExternalID" class="form-label">Внешний id</label>
                <input id="inputNewPoint-ExternalID"
                class="form-control"  placeholder="Введите внешний id">
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-Carpet" class="form-label">Покрытие</label>
                <select id="inputNewPoint-Carpet" class="form-select">
                    ${["Асфальт", "Плитка"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${"Асфальт" == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-NumberArc" class="form-label">Количество дуг</label>
                <input id="inputNewPoint-NumberArc" type="number" min="0"
                class="form-control"  placeholder="Введите количество дуг">
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-ArcType" class="form-label">Тип дуги</label>
                <select id="inputNewPoint-ArcType" class="form-select">
                    ${["Алюминиевая", "Металлическая", "Разметка", "Другое"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${"Алюминиевая" == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-Owner" class="form-label">Владелец</label>
                <select id="inputNewPoint-Owner" class="form-select">
                    ${[ "yandex", "whoosh", "yabike"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${"yandex" == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-Operator" class="form-label">Оператор</label>
                <select id="inputNewPoint-Operator" class="form-select">
                    ${["ultradop"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${"ultradop" == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputNewPoint-Customer" class="form-label">Заказчик задачи</label>
                <select id="inputNewPoint-Customer" class="form-select">
                    ${["Яндекс", "Whoosh", "Ultradop", "Яндекс Байк", "Другое"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}">${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-12">
                <label for="inputNewPoint-Comment" class="form-label">
                    Комментарий к задаче
                </label>
                <textarea id="inputNewPoint-Comment" class="form-control" rows="4"></textarea>
            </div>
        </div>
    </div>
    `
}