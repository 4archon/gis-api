let pointEdit = new bootstrap.Modal(document.getElementById("point-edit"), null);


function editPointProfile(event) {
    let id = Number(event.currentTarget.getAttribute("data-id"));
    let point = data.find((el) => el.id == id);
    render_point_edit_header(id);
    render_point_edit_body(point);
    render_point_edit_footer(id);
    pointProfile.hide();
    pointEdit.show();
}

function render_point_edit_header(id) {
    let container = document.getElementById("point-edit-header");
    container.innerHTML = 
    `<h1 class="modal-title fs-5">Редактировать данные точки
        <span class="badge text-bg-primary">${id}</span>
    </h1>
    <button type="button" class="btn-close" data-id="${id}"
    data-bs-dismiss="modal" aria-label="Close" onclick="closeModal(event)"></button>`
}

function closeModal(event) {
    let id = Number(event.currentTarget.getAttribute("data-id"));
    showPointProfile(id);
}

function render_point_edit_footer(id) {
    let container = document.getElementById("point-edit-footer");
    container.innerHTML = 
    `
    <button type="button" class="btn btn-secondary"
    data-id="${id}" onclick="cancelPointEdit(event)">
        Отменить
    </button>
    <button data-id="${id}" type="button" class="btn btn-primary"
    onclick="sendPointEdit(event)">
        Сохранить
    </button>
    `
}

function cancelPointEdit(event) {
    let id = Number(event.currentTarget.getAttribute("data-id"));
    pointEdit.hide();
    showPointProfile(id);
}

function sendPointEdit(event) {
    let id = Number(event.currentTarget.getAttribute("data-id"));
    let point = data.find((el) => el.id == id);
    let pointChange = {
        id: id,
        lat: Number(document.getElementById("inputPointEdit-Lat").value),
        long: Number(document.getElementById("inputPointEdit-Long").value),
        address: document.getElementById("inputPointEdit-Address").value,
        district: document.getElementById("inputPointEdit-District").value,
        active: document.getElementById("inputPointEdit-Active").value == "true" ? true:false,
        status: document.getElementById("inputPointEdit-Status").value,
        externalID: document.getElementById("inputPointEdit-ExternalID").value,
        carpet: document.getElementById("inputPointEdit-Carpet").value,
        numberArc: Number(document.getElementById("inputPointEdit-NumberArc").value),
        arcType: document.getElementById("inputPointEdit-ArcType").value,
        owner: document.getElementById("inputPointEdit-Owner").value,
        operator: document.getElementById("inputPointEdit-Operator").value,
        comment: document.getElementById("inputPointEdit-Comment").value,
        marks: null
    };

    if (point.marks !== null) {
        point.marks.forEach((el) => {
            let mark = {
                id: el.id,
                number: document.getElementById(`inputPointEdit-Mark-Number${el.id}`).value,
                type: document.getElementById(`inputPointEdit-Mark-Type${el.id}`).value,
                active:document.getElementById(`inputPointEdit-Mark-Active${el.id}`).value
                == "true" ? true:false,
            }
            if (pointChange.marks === null) {pointChange.marks = []}
            pointChange.marks.push(mark);
        });
    }

    postPointEdit(pointChange);
}

async function postPointEdit(data) {
    let url = "/point_edit"
    response = await fetch(url, {
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
        pointEdit.hide();
        await getPoinst();
        showPointProfile(data.id);
    }
}

function render_point_edit_body(point) {
    let container = document.getElementById("point-edit-body");
    container.innerHTML = 
    `
    <div class="card-body mb-4">
        <div class="row g-3">
            <div class="col-md-6">
                <label for="inputPointEdit-Lat" class="form-label">Широта</label>
                <input id="inputPointEdit-Lat"
                value="${point.lat === null ? "": point.lat}" 
                class="form-control"  placeholder="Введите широту">
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-Long" class="form-label">Долгота</label>
                <input id="inputPointEdit-Long"
                value="${point.long === null ? "": point.long}" 
                class="form-control"  placeholder="Введите долготу">
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-Address" class="form-label">Адрес</label>
                <input id="inputPointEdit-Address"
                value="${point.address === null ? "": point.address}" 
                class="form-control"  placeholder="Введите адрес">
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-District" class="form-label">Округ</label>
                <input id="inputPointEdit-District"
                value="${point.district === null ? "": point.district}" 
                class="form-control"  placeholder="Введите округ">
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-Active" class="form-label">Статус активации</label>
                <select id="inputPointEdit-Active" class="form-select">
                    <option value="true" ${point.active ? "selected":""}>Активна</option>
                    <option value="false" ${point.active ? "":"selected"}>Деактивирована</option>
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-Status" class="form-label">Статус</label>
                <select id="inputPointEdit-Status" class="form-select">
                    ${["Временно сняты дуги", "Точка доступна", "Точка недоступна",
                        "Временно невозможно проверить точку", "Идет благоустройство",
                        "Идет благоустройство - требуется забрать дуги",
                        "Идет благоустройство - требуется демонтировать и забрать дуги"
                    ].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${point.status == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-ExternalID" class="form-label">Внешний id</label>
                <input id="inputPointEdit-ExternalID"
                value="${point.externalID === null ? "": point.externalID}" 
                class="form-control"  placeholder="Введите внешний id">
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-Carpet" class="form-label">Покрытие</label>
                <select id="inputPointEdit-Carpet" class="form-select">
                    ${["Асфальт", "Плитка"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${point.carpet == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-NumberArc" class="form-label">Количество дуг</label>
                <input id="inputPointEdit-NumberArc" type="number" min="0"
                value="${point.numberArc === null ? "0": point.numberArc}" 
                class="form-control"  placeholder="Введите количество дуг">
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-ArcType" class="form-label">Тип дуги</label>
                <select id="inputPointEdit-ArcType" class="form-select">
                    ${["Алюминиевая", "Металлическая", "Другое"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${point.arcType == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-Owner" class="form-label">Владелец</label>
                <select id="inputPointEdit-Owner" class="form-select">
                    ${["whoosh", "yandex"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${point.owner == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputPointEdit-Operator" class="form-label">Оператор</label>
                <select id="inputPointEdit-Operator" class="form-select">
                    ${["ultradop"].reduce((acc, el) => {
                        return acc +
                        `
                        <option value="${el}" ${point.operator == el ? "selected":""}>${el}</option>
                        `
                    }, "")}
                </select>
            </div>
            ${point.marks === null ? ``:
                `<h4 class="my-3">Разметка</h4>` +
                point.marks.reduce((acc, el) => {
                    return acc + 
                    `
                    <h6 class="m-0">
                        Разметка
                        <span class="badge text-bg-primary">${el.id}</span>
                    </h6>
                    <div class="col-md-4">
                        <label for="inputPointEdit-Mark-Number${el.id}" class="form-label">Номер разметки</label>
                        <input id="inputPointEdit-Mark-Number${el.id}"
                        value="${el.number === null ? "": el.number}" 
                        class="form-control"  placeholder="Введите номер разметки">
                    </div>
                    <div class="col-md-5">
                        <label for="inputPointEdit-Mark-Type${el.id}" class="form-label">Тип дуги</label>
                        <select id="inputPointEdit-Mark-Type${el.id}" class="form-select">
                            ${["Дорожная краска", "Термопластик", "Баллончик"].reduce((acc2, el2) => {
                                return acc2 +
                                `
                                <option value="${el2}" ${el.type == el2 ? "selected":""}>${el2}</option>
                                `
                            }, "")}
                        </select>
                    </div>
                    <div class="col-md-3">
                        <label for="inputPointEdit-Mark-Active${el.id}" class="form-label">Активна</label>
                        <select id="inputPointEdit-Mark-Active${el.id}" class="form-select">
                            <option value="true" ${el.active ? "selected":""}>Да</option>
                            <option value="false" ${el.active ? "":"selected"}>Нет</option>
                        </select>
                    </div>
                    `
                }, "")
            }
            <div class="col-12">
                <label for="inputPointEdit-Comment" class="form-label">
                    Комментарий изменения
                </label>
                <textarea id="inputPointEdit-Comment" class="form-control" rows="4">${point.comment === null ? "" : point.comment}</textarea>
            </div>
        </div>
    </div>
    `
}