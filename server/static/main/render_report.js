let selectedTasks = [];
let reportData = null;
let reportPointID = null;

function reportClick(event) {
    let id = event.target.getAttribute("data-id");
    reportPointID = id;
    selectedTasks = [];
    reportData = null;
    render_report_header(id);
    render_report_tasks();
    render_report_footer(id, false);
    pointProfile.hide();
    pointReport.show();
}

function render_report_header(id) {
    result = 
    `
    <h1 class="modal-title fs-5">Отчет по точке 
    <span class="badge text-bg-primary">${id}</span>
    </h1>
    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
    `
    
    document.getElementById("point-report-header").innerHTML = result;
}

function render_report_footer(id, able) {
    let cont = document.getElementById("point-report-footer");
    if (able) {
        cont.innerHTML = 
        `
        <button data-id="${id}" type="button" class="btn btn-success" onclick="">
            Отправить отчёт
        </button>
        `
    } else {
        cont.innerHTML = 
        `
        <button data-id="${id}" disabled type="button" class="btn btn-success" onclick="">
            Отправить отчёт
        </button>
        `
    }
}

function render_report_tasks() {
    let result =
    `
    <h5 class="">Отметьте выполненные задачи:</h5>
    <div class="accordion accordion-flush">
        ${tasks === null ? 
            `
            <div class="card mt-1" data-id="-1" onclick="taskClick(event)">
                <div class="card-body">
                    <h5>Проинспектировать</h5>
                </div>
            </div>
            <div class="card mt-1" data-id="-2" onclick="taskClick(event)">
                <div class="card-body">
                    <h5>Произвести сервис</h5>
                </div>
            </div>
            <div class="card mt-1" data-id="-3" onclick="taskClick(event)">
                <div class="card-body">
                    <h5>Невозможно произвести работы</h5>
                </div>
            </div>
            `:
            tasks.reduce((acc, el) => {
            return acc +=
            `
            <div class="card mt-1" data-id="${el.id}" onclick="taskClick(event)">
                <div class="card-body">
                    <h5>${el.type}</h5>
                    <p class="card-text">
                    ${el.comment === null ? "Нет комменатария": el.comment}
                    </p>
                </div>
            </div>
            `
        }, "") +
        `
        <div class="card mt-1" data-id="-3" onclick="taskClick(event)">
            <div class="card-body">
                <h5>Невозможно произвести работы</h5>
            </div>
        </div>
        `
        }
    </div>
    <div class="mt-2" id="report-form"/>
    `

    let body = document.getElementById("point-report-body");
    body.innerHTML = result;
}

function findTask(id) {
    let task;
    if (id == -1) {
        task = {
            id: id,
            type: "Проинспектировать",
            comment: null
        }
    } else if (id == -2) {
        task = {
            id: id,
            type: "Произвести сервис",
            comment: null
        }
    } else if (id == -3) {
        task = {
            id: id,
            type: "Невозможно произвести работы"
        }
    } 
    else {
        task = tasks.find((element) => element.id == id);
    }
    return task
}

function taskClick(event) {
    target = event.currentTarget;
    id = target.getAttribute("data-id");
    let task = findTask(id);
    task["target"] = target;
    let found = selectedTasks.find((element) => element.id == task.id);
    if (found === undefined) {
        if (task.id == -3) {
            selectedTasks.forEach((el) => {
                el.target.classList.remove("text-bg-success");
                el.target.classList.remove("bg-gradient");
            })
            selectedTasks = [];
            selectedTasks.push(task)
            target.classList.add("text-bg-danger");
            target.classList.add("bg-gradient");
        } else {
            decline = selectedTasks.find((element) => element.id == -3);
            if (decline !== undefined) {
                decline.target.classList.remove("text-bg-danger");
                decline.target.classList.remove("bg-gradient");
                selectedTasks = selectedTasks.filter((element) => element.id != -3);
            }
            selectedTasks.push(task)
            target.classList.add("text-bg-success");
            target.classList.add("bg-gradient");
        }
    } else {
        selectedTasks = selectedTasks.filter((element) => element.id != found.id);
        if (found.id == -3) {
            target.classList.remove("text-bg-danger");
            target.classList.remove("bg-gradient");
        } else {
            target.classList.remove("text-bg-success");
            target.classList.remove("bg-gradient");
        }
        
    }
    render_report_form();
    // console.log(selectedTasks);
}

function new_report_data(type) {
    if (type == "decline") {
        return {
            type: "decline",
            reason: null,
            yourself: null,
            pointID: null
        }
    } else if (type == "service") {
        return {
            type: "service",
            carry: null,
            newLocation: null,
            data: null
        }
    } else if (type == "inspection") {
        return {
            type: "inspection",
            data: null
        }
    }
}

function render_report_form() {
    let form = document.getElementById("report-form");
    if (selectedTasks.length == 0) {
        form.innerHTML = ``;
        reportData = null;
    } else {
        if (selectedTasks.some((el) => el.type == "Невозможно произвести работы")) {
            result = 
            `
            <h5>Выберете причины невозможности выполнения работ:</h5>
            <div id="report-data"/>
            `
            reportData = new_report_data("decline");
            
        } else if (selectedTasks.some((el) => el.type !== "Проинспектировать")) {
            result = 
            `
            <h5>Добавьте выполненные работы:</h5>
            <div id="report-data"/>
            `
            reportData = new_report_data("service");
        } else {
            result = 
            `
            <h5>Заполните отчет инспекции:</h5>
            <div id="report-data"/>
            `
            reportData = new_report_data("inspection");
        }
    form.innerHTML = result;
    render_data_to_form();
    }
}

function render_data_to_form() {
    console.log(reportData);
    if (reportData.type == "decline") {
        render_decline_to_form();
    } else if (reportData.type == "service") {
        render_service_to_form();
    } else if (reportData.type == "inspection") {
        render_inspection_to_form();
    }
}

function render_decline_to_form() {
    let container = document.getElementById("report-data");
    container.innerHTML = 
    `
    <div class="form-check">
        <input class="form-check-input" type="radio" name="radioReason" id="radioReason1"
        onchange="decline_changed(event)"
        ${reportData.reason == "Временно невозможно проверить точку" ? "checked": ""}>
        <label class="form-check-label" for="radioReason1">
            <h6>Временно невозможно проверить точку</h6>
        </label>
    </div>
    <div class="form-check">
        <input class="form-check-input" type="radio" name="radioReason" id="radioReason2"
        onchange="decline_changed(event)"
        ${reportData.reason == "Идет благоустройство" ? "checked": ""}>
        <label class="form-check-label" for="radioReason2">
            <h6>Идет благоустройство</h6>
        </label>
    </div>
    <div class="form-check">
        <input class="form-check-input" type="radio" name="radioReason" id="radioReason3"
        onchange="decline_changed(event)"
        ${reportData.reason == "Идет благоустройство - требуется забрать дуги" ? "checked": ""}>
        <label class="form-check-label" for="radioReason3">
            <h6>Идет благоустройство - требуется забрать дуги</h6>
        </label>
    </div>
    <div class="form-check">
        <input class="form-check-input" type="radio" name="radioReason" id="radioReason4"
        onchange="decline_changed(event)"
        ${reportData.reason == "Идет благоустройство - требуется демонтировать и забрать дуги" ? "checked": ""}>
        <label class="form-check-label" for="radioReason4">
            <h6>Идет благоустройство - требуется демонтировать и забрать дуги</h6>
        </label>
    </div>
    <div class="form-check">
        <input class="form-check-input" type="radio" name="radioReason" id="radioReason5"
        onchange="decline_changed(event)"
        ${reportData.reason == "Точка является дублем" ? "checked": ""}>
        <label class="form-check-label" for="radioReason5">
            <h6>Точка является дублем</h6>
        </label>
    </div>
    <div class="form-check">
        <input class="form-check-input" type="radio" name="radioReason" id="radioReason6"
        onchange="decline_changed(event)"
        ${reportData.reason == "Невозможно установить дуги, необходимо деактивировать" ? "checked": ""}>
        <label class="form-check-label" for="radioReason6">
            <h6>Невозможно установить дуги, необходимо деактивировать</h6>
        </label>
    </div>
    `
    container.innerHTML += reportData.reason != "Идет благоустройство - требуется забрать дуги" ? ``: 
    `
    ${reportData.yourself === null ? 
        `
        <h5>Заберете дуги самостоятельно?</h5>
        <div class="container">
            <div class="row">
                <button id="yourself1" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                onclick="yourself_changed(event)">Да</button>
                <button id="yourself2" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                onclick="yourself_changed(event)">Нет</button>
            </div>
        </div>
        `
        :
        `
        ${reportData.yourself ? 
            `
            <h5>Заберете дуги самостоятельно?</h5>
            <div class="container">
                <div class="row">
                    <button id="yourself1" type="button" class="btn btn-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)" disabled>Да</button>
                    <button id="yourself2" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)">Нет</button>
                </div>
            </div>
            `
            :
            `
            <h5>Заберете дуги самостоятельно?</h5>
            <div class="container">
                <div class="row">
                    <button id="yourself1" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)">Да</button>
                    <button id="yourself2" type="button" class="btn btn-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)" disabled>Нет</button>
                </div>
            </div>
            `}
        `
    }
    `
    container.innerHTML += reportData.reason != "Идет благоустройство - требуется демонтировать и забрать дуги" ? ``: 
    `
    ${reportData.yourself === null ? 
        `
        <h5>Демонтируете и заберете дуги самостоятельно?</h5>
        <div class="container">
            <div class="row">
                <button id="yourself1" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                onclick="yourself_changed(event)">Да</button>
                <button id="yourself2" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                onclick="yourself_changed(event)">Нет</button>
            </div>
        </div>
        `
        :
        `
        ${reportData.yourself ? 
            `
            <h5>Демонтируете и заберете дуги самостоятельно?</h5>
            <div class="container">
                <div class="row">
                    <button id="yourself1" type="button" class="btn btn-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)" disabled>Да</button>
                    <button id="yourself2" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)">Нет</button>
                </div>
            </div>
            `
            :
            `
            <h5>Демонтируете и заберете дуги самостоятельно?</h5>
            <div class="container">
                <div class="row">
                    <button id="yourself1" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)">Да</button>
                    <button id="yourself2" type="button" class="btn btn-secondary col-sm-1 col-5 mx-1"
                    onclick="yourself_changed(event)" disabled>Нет</button>
                </div>
            </div>
            `}
        `
    }
    `
    container.innerHTML += reportData.reason != "Точка является дублем" ? "" : 
    `
    ${reportData.pointID === null ? 
        `
        <h5>Укажите id совпадающей точки</h5>
        <div class="container">
            <div class="row">
                <div class="col-12 col-sm-4">
                    <input class="form-control" type="text" placeholder="ID точки"
                    onchange="check_point_id(event)">
                </div>
            </div>
        </div>
        `
        :
        `
        ${reportData.pointID === false ? 
            `
            <h5>Укажите id совпадающей точки</h5>
            <div class="container">
                <div class="row">
                    <div class="col-12 col-sm-4">
                        <input class="form-control is-invalid" type="text" placeholder="Невалидный ID"
                        onchange="check_point_id(event)">
                    </div>
                </div>
            </div>
            `
            :
            `
            <h5>Укажите id совпадающей точки</h5>
            <div class="container">
                <div class="row">
                    <div class="col-12 col-sm-4">
                        <input class="form-control" type="text" placeholder="ID точки"
                        onchange="check_point_id(event)"
                        value="${reportData.pointID}">
                    </div>
                </div>
            </div>
            `
        }
        `
    }
    `
}

function decline_changed(event) {
    let id = event.currentTarget.id;
    reportData = new_report_data("decline");
    switch(id) {
        case "radioReason1":
            reportData.reason = "Временно невозможно проверить точку";
            break;
        case "radioReason2":
            reportData.reason = "Идет благоустройство";
            break;
        case "radioReason3":
            reportData.reason = "Идет благоустройство - требуется забрать дуги";
            break;
        case "radioReason4":
            reportData.reason = "Идет благоустройство - требуется демонтировать и забрать дуги";
            break;
        case "radioReason5":
            reportData.reason = "Точка является дублем";
            break;
        case "radioReason6":
            reportData.reason = "Невозможно установить дуги, необходимо деактивировать";
            break;
    }
    render_data_to_form();
}

function yourself_changed(event) {
    let id = event.currentTarget.id;
    switch(id) {
        case "yourself1":
            reportData.yourself = true;
            break;
        case "yourself2":
            reportData.yourself = false;
            break;
    }
    render_data_to_form();
}

function check_point_id(event) {
    let val = event.currentTarget.value;
    const re = /\d+/;
    let res = val.match(re);
    if (res !== null && res[0].length == val.length) {
        let id = Number(res);
        let found = data.find((el) => el.id == id);
        if (found !== undefined) {
            reportData.pointID = id;
        } else {
            reportData.pointID = false;    
        }
    } else {
        reportData.pointID = false;
    }
    render_data_to_form();
}

function render_service_to_form() {
    let container = document.getElementById("report-data");
    container.innerHTML = reportData.data === null ? ``:
    `

    `
    render_add_button_to_form();
    container.innerHTML += reportData.carry === null ? 
    `
    <h5>Был ли перенос точки?</h5>
    <div class="container">
        <div class="row">
            <button id="carry1" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
            onclick="carry_changed(event)">Да</button>
            <button id="carry2" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
            onclick="carry_changed(event)">Нет</button>
        </div>
    </div>
    `
    :
    `
    ${reportData.carry ? 
        `
        <h5>Был ли перенос точки?</h5>
        <div class="container">
            <div class="row">
                <button id="carry1" type="button" class="btn btn-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="carry_changed(event)" disabled>Да</button>
                <button id="carry2" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="carry_changed(event)">Нет</button>
            </div>
        </div>
        ${reportData.newLocation === null ? 
            `
            <h5>Укажите новое местоположение</h5>
            <button type="button" class="btn btn-secondary" data-id="${reportPointID}"
            onclick="new_location(event)">Указать</button>
            `
            :
            `
            <h5>Новое местоположение (${reportData.newLocation})</h5>
            `
        }
        `
        :
        `
        <h5>Был ли перенос точки?</h5>
        <div class="container">
            <div class="row">
                <button id="carry1" type="button" class="btn btn-outline-secondary col-sm-1 col-5 mx-1"
                onclick="carry_changed(event)">Да</button>
                <button id="carry2" type="button" class="btn btn-secondary col-sm-1 col-5 mx-1"
                onclick="carry_changed(event)" disabled>Нет</button>
            </div>
        </div>
        `
    }
    `
}

function carry_changed(event) {
    let id = event.currentTarget.id;
    reportData.newLocation = null;
    switch(id) {
        case "carry1":
            reportData.carry = true;
            break;
        case "carry2":
            reportData.carry = false;
            break;
    }
    render_data_to_form();
}

function render_add_button_to_form() {
    let container = document.getElementById("report-data");
    result = ``
    if (reportData.type == "service") {
        result = 
        `
        <div class="container my-2">
            <div class="row">
                <div class="col-12 col-sm-6">
                    <div class="input-group">
                        <button class="btn btn-outline-secondary" type="button"
                        onclick="add_new_unit()">Добавить</button>
                        <select id="unit-selector" class="form-select">
                            <option value="dismantling" selected>Демотнаж</option>
                            <option value="mounting">Монтаж</option>
                            <option value="dyeing">Покраска</option>
                            <option value="marking">Нанесение разметки</option>
                        </select>
                    </div>
                </div>
            </div>
        </div>
        `
    } else if (reportData.type == "inspection") {
        result = 
        `
        <div class="container mt-4">
            <div class="row">
                <div class="col-12 col-sm-6">
                    <div class="input-group">
                        <button class="btn btn-outline-secondary" type="button"
                        onclick="add_new_unit()">Добавить</button>
                        <select id="unit-selector" class="form-select">
                            <option value="dismantling" selected>Демотнаж-монтаж</option>
                            <option value="mounting">Монтаж</option>
                            <option value="dyeing">Покраска</option>
                            <option value="marking">Исправить разметку</option>
                        </select>
                    </div>
                </div>
            </div>
        </div>
        `
    } else {
        result = ``
    }
    container.innerHTML += result;
}

function add_new_unit() {
    let val = document.getElementById("unit-selector").value;
    if (reportData.type == "service") {
        if (reportData.data === null) {reportData.data = []}
        if (val == "dismantling") {
            reportData.data.push({
                type: "dismantling"

            });
        } else if (val == "mounting") {
            reportData.data.push({
                type: "mounting"
            });
        } else if (val == "dyeing") {
            reportData.data.push({
                type: "dyeing"
            });
        } else {
            reportData.data.push({
                type: "marking"
            });
        }
    } else if (reportData.type == "inspection") {
        if (reportData.data === null) {reportData.data = []}
        if (val == "dismantling") {
            reportData.data.push({
                type: "dismantling"
            });
        } else if (val == "mounting") {
            reportData.data.push({
                type: "mounting"
            });
        } else if (val == "dyeing") {
            reportData.data.push({
                type: "dyeing",
                yourself: null
            });
        } else {
            reportData.data.push({
                type: "marking"
            });
        }
    }
    console.log(reportData);
    render_data_to_form();
}


function loadMedia(event) {
    console.log(event.currentTarget);
    console.log(event.currentTarget.files);
    const ur = window.URL.createObjectURL(event.currentTarget.files[0]);
    appendReportBody(ur);
}

function appendReportBody(ur) {
    let form = document.getElementById("report-form");
    res = 
    `
    <div class="col-3 d-flex justify-content-center">
        <a class="profileMedia d-flex align-items-center" data-gall="gallery-profile" href="${ur}">
            <img src="${ur}" 
            alt="loading" style="max-height: 200px; max-width: 100%; border-radius: 5px;"/>
        </a>
    </div>
    `
    form.innerHTML += res;
}