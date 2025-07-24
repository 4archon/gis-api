let tasks;
let appoint;
let selectedTasks = [];
let reportData = null;
let reportPointID = null;

function filterTasks() {
    if (tasks === null) {
        return
    }
    if (userSubgroup == "inspection") {
        tasks = tasks.filter((el) => el.type == "Проинспектировать");
    } else if (userSubgroup == "service") {
        tasks = tasks.filter((el) => el.type != "Проинспектировать");
    }
    if (tasks.length == 0) {
        tasks = null;
    }
}

async function reportClick(event) {
    let id = event.target.getAttribute("data-id");
    reportPointID = id;
    selectedTasks = [];
    reportData = null;
    tasks = (await getCurrentTasks(id)).tasks;
    filterTasks();
    appoint = data.find((el) => el.id == id).appoint;
    render_report_header(id);
    render_report_tasks();
    render_report_footer(false);
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

function render_report_footer(able) {
    let container = document.getElementById("point-report-footer");
    if (able) {
        container.innerHTML = 
        `
        <button type="button" class="btn btn-success" onclick="sendReport(event)">
            Отправить отчёт
        </button>
        `
    } else {
        container.innerHTML = 
        `
        <button disabled type="button" class="btn btn-success">
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
            ${userSubgroup == "service" ? 
                `
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
                `
                :
                `
                <div class="card mt-1" data-id="-1" onclick="taskClick(event)">
                    <div class="card-body">
                        <h5>Проинспектировать</h5>
                    </div>
                </div>
                <div class="card mt-1" data-id="-3" onclick="taskClick(event)">
                    <div class="card-body">
                        <h5>Невозможно произвести работы</h5>
                    </div>
                </div>
                `
            }
            `
            :
            tasks.reduce((acc, el) => {
            return acc +=
            `
            <div class="card mt-1" data-id="${el.id}" onclick="taskClick(event)">
                <div class="card-body">
                    <h5>${el.type}</h5>
                    <p class="card-text">
                    ${el.comment === null ? "": el.comment}
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
    switch(Number(id)) {
    case -1:
        task = {
            id: id,
            type: "Проинспектировать",
            comment: null
        };
        break;
    case -2:
        task = {
            id: id,
            type: "Произвести сервис",
            comment: null
        };
        break;
    case -3:
        task = {
            id: id,
            type: "Невозможно произвести работы"
        };
        break;
    default:
        task = tasks.find((element) => element.id == id);
        break;
    }
    return task
}

function taskClick(event) {
    let target = event.currentTarget;
    let id = target.getAttribute("data-id");
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
            pointID: null,
            comment: null
        }
    } else if (type == "service") {
        return {
            type: "service",
            data: null,
            carry: null,
            newLocation: null,
            left: null,
            comment: null,
            notRequire: false
        }
    } else if (type == "inspection") {
        return {
            type: "inspection",
            data: null,
            yourself: null,
            paintCount: null,
            comment: null,
            notRequire: false
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
            <h5>Выберите причины невозможности выполнения работ:</h5>
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
    render_report_footer(false);
    let valid = validateForm();
    if (valid) {
        // console.log(valid);
        render_load_media();
    }
}

function validateForm() {
    if (reportData.type == "service") {
        if (reportData.data !== null && reportData.data.length > 0) {
            if (reportData.left !== null && reportData.carry !== null
                && !reportData.notRequire) {
                if (reportData.carry) {
                    if(reportData.newLocation !== null &&
                        reportData.data.some((el) => el.type == "Демонтаж") &&
                        reportData.data.some((el) => el.type == "Монтаж")) {
                        if (reportData.left === false) {
                            return true
                        } else {
                            if (reportData.left !== null && reportData.left.length > 0) {
                                return true;
                            }
                        }
                    }
                } else {
                    if (reportData.left === false) {
                        return true
                    } else {
                        if (reportData.left !== null && reportData.left.length > 0) {
                            return true;
                        }
                    }
                }
            } else if (reportData.notRequire) {
                return true;
            }
        }
    } else if (reportData.type == "decline" && reportData.reason !== null) {
        if (reportData.reason == "Идет благоустройство - требуется забрать дуги" ||
            reportData.reason == "Идет благоустройство - требуется демонтировать и забрать дуги") {
                if (reportData.yourself !== null) {
                    return true;
                }
        } else if (reportData.reason == "Точка является дублем") {
            if (reportData.pointID) {
                return true;
            }
        } else {
            return true;
        }
    } else if (reportData.type == "inspection" && 
        reportData.data !== null && reportData.data.length > 0) {
        if (reportData.data.some((el) => el.type == "Покраска")) {
            if (reportData.yourself !== null && reportData.count !== null) {
                return true;
            }
        } else {
            return true;
        }

    }
    return false;
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
    container.innerHTML +=
    `
    <h5 class="mt-1">Оставить комментарий</h5>
    <textarea class="form-control" onchange="commentChange(event)"
    rows="3" >${reportData.comment === null ? "": reportData.comment}</textarea>
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

function commentChange(event) {
    reportData.comment = event.currentTarget.value;
}

function render_service_to_form() {
    let container = document.getElementById("report-data");
    container.innerHTML = reportData.data === null ? ``:
    `
    <div class="container">
        <div class="row">
            ${reportData.data.reduce((acc, el, index) => {
                return acc += 
                `
                <div class="card mt-1 col-lg-6 col-12 p-0">
                    <div class="card-header">
                        <div class="row">
                            <div class="col-8">
                                ${el.type}
                            </div>
                            <div class="col-4">
                                <button data-index="${index}" type="button"
                                class="btn btn-outline-danger btn-sm float-end" 
                                onclick="deleteWork(event)">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash" viewBox="0 0 16 16">
                                        <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0z"/>
                                        <path d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4zM2.5 3h11V2h-11z"/>
                                    </svg>
                                </button>
                            </div>
                        </div>
                    </div>
                    <div class="card-body row">
                        ${el.type == "Нанесение разметки" ? 
                            `
                            <label class="col-auto col-form-label">Номер разметки</label>
                            <div class="col-4 p-0">
                                <input type="text" class="form-control" 
                                value="${el.number === undefined ? "": el.number}"
                                data-index="${index}" onchange="markingChangeNumber(event)">
                            </div>
                            `:""
                        }
                        ${el.type == "Демаркировка" ? 
                            `
                            <label class="col-auto col-form-label">Выберите разметку:</label>
                            ${data.find((point) => point.id == reportPointID).marks === null ? "":
                                data.find((point) => point.id == reportPointID).marks.
                                reduce((acc2, mark) => {
                                    return acc2 + 
                                    `
                                    <div class="card mt-1
                                    ${reportData.data[index].selectedMarks !== undefined &&
                                        reportData.data[index].selectedMarks.
                                        some((markEl) => markEl == mark.id) ? 
                                        "text-bg-success bg-gradient":""}" 
                                    data-id="${mark.id}"
                                    onclick="markingClick(event)" data-index="${index}">
                                        <div class="card-body">
                                            <h6>${mark.type}
                                            <span class="badge text-bg-primary">${mark.number}</span>
                                            </h6>
                                        </div>
                                    </div>
                                    `
                            }, "")}
                            `:""
                        }
                        ${el.type == "Работа не требуется" ? 
                            `
                            <h6>Точка не требует выполнения работ</h6>
                            `
                            :
                            ``
                        }
                        ${el.type != "Нанесение разметки" && el.type != "Демаркировка"
                            && el.type != "Работа не требуется"? 
                            `
                            <label class="col-auto col-form-label">Количество:</label>
                            <div class="col-4 p-0">
                                <input type="number" class="form-control" value="${el.count}" min="1"
                                data-index="${index}" onchange="changeWorkCounter(event)">
                            </div>
                            `:``
                        }
                    </div>
                </div>
                `
            }, "")}
        </div>
    </div>
    `
    render_add_button_to_form();
    if (!reportData.notRequire) {
    // плохой перенос строки
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
            <h5>Выбрано новое местоположение</h5>
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
    container.innerHTML += reportData.left === null ? 
    `
    <h5 class="mt-1">Остались невыполненные работы?</h5>
    <div class="container">
        <div class="row">
            <button id="left-button-yes" type="button"
            class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
            onclick="left_work_changed(event)">Да</button>
            <button id="left-button-no" type="button"
            class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
            onclick="left_work_changed(event)">Нет</button>
        </div>
    </div>
    `
    :
    `
    ${reportData.left ? 
        `
        <h5 class="mt-1">Остались невыполненные работы?</h5>
        <div class="container">
            <div class="row">
                <button id="left-button-yes" type="button"
                class="btn btn-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="left_work_changed(event)" disabled>Да</button>
                <button id="left-button-no" type="button"
                class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="left_work_changed(event)">Нет</button>
            </div>
        </div>
        `
        :
        `
        <h5 class="mt-1">Остались невыполненные работы?</h5>
        <div class="container">
            <div class="row">
                <button id="left-button-yes" type="button"
                class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="left_work_changed(event)">Да</button>
                <button id="left-button-no" type="button"
                class="btn btn-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="left_work_changed(event)" disabled>Нет</button>
            </div>
        </div>
        `
    }
    `
    if (reportData.left) {
        let re = reportData.left.reduce((acc, el, index) => {
            return acc + `
            <div class="card mt-1 col-lg-6 col-12 p-0">
                <div class="card-header">
                    <div class="row">
                        <div class="col-8">
                            ${el.type}
                        </div>
                        <div class="col-4">
                            <button data-index="${index}" type="button"
                            class="btn btn-outline-danger btn-sm float-end" 
                            onclick="deleteLeftWork(event)">
                                <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash" viewBox="0 0 16 16">
                                    <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0z"/>
                                    <path d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4zM2.5 3h11V2h-11z"/>
                                </svg>
                            </button>
                        </div>
                    </div>
                </div>
                <div class="card-body row">
                    <label class="col-auto col-form-label">Количество:</label>
                    <div class="col-4 p-0">
                        <input type="number" class="form-control" value="${el.count}" min="1"
                        data-index="${index}" onchange="changeLeftWorkCounter(event)">
                    </div>
                </div>
            </div>
            `
        }, "")
        container.innerHTML += 
        `
        <div class="container">
            <div class="row">
            ${re}
            </div>
        </div>
        `
        render_service_letf_add_button();
    }
    }
    container.innerHTML +=
    `
    <h5 class="mt-1">Оставить комментарий</h5>
    <textarea class="form-control" onchange="commentChange(event)"
    rows="3" >${reportData.comment === null ? "": reportData.comment}</textarea>
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

function left_work_changed(event) {
    let id = event.currentTarget.id;
    switch(id) {
        case "left-button-yes":
            reportData.left = [];
            break;
        case "left-button-no":
            reportData.left = false;
            break;
    }
    render_data_to_form();
}

function render_add_button_to_form() {
    let container = document.getElementById("report-data");
    result = ``
    if (reportData.data !== null && 
        reportData.data.some((el) => el.type == "Работа не требуется")) {return}
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
                            <option value="Демонтаж" selected>Демонтаж</option>
                            <option value="Монтаж">Монтаж</option>
                            <option value="Покраска">Покраска</option>
                            <option value="Нанесение разметки">Нанесение разметки</option>
                            <option value="Частичное нанесение">Частичное нанесение</option>
                            <option value="Демаркировка">Демаркировка</option>
                            <option value="Работа не требуется">Работа не требуется</option>
                        </select>
                    </div>
                </div>
            </div>
        </div>
        `
    } else if (reportData.type == "inspection") {
        result = 
        `
        <div class="container my-2">
            <div class="row">
                <div class="col-12 col-sm-6">
                    <div class="input-group">
                        <button class="btn btn-outline-secondary" type="button"
                        onclick="add_new_unit()">Добавить</button>
                        <select id="unit-selector" class="form-select">
                            <option value="Демонтаж-монтаж" selected>Демонтаж-монтаж</option>
                            <option value="Монтаж">Монтаж</option>
                            <option value="Покраска">Покраска</option>
                            <option value="Частичное нанесение">Частичное нанесение</option>
                            <option value="Работа не требуется">Работа не требуется</option>
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
    if (reportData.data === null) {reportData.data = []}
    if (val == "Работа не требуется") {
        reportData = new_report_data(reportData.type);
        reportData.data = [];
        reportData.notRequire = true;
        // bag with further form
    }
    if (reportData.data.some((el) => el.type == val && val != "Нанесение разметки")) {
         return
    }
    reportData.data.push({
        type: val,
        count: 1
    });
    render_data_to_form();
}

function render_service_letf_add_button() {
    let container = document.getElementById("report-data");
    result = ``
    if (reportData.data !== null && 
        reportData.data.some((el) => el.type == "Работа не требуется")) {return}
    if (reportData.type == "service") {
        result = 
        `
        <div class="container my-2">
            <div class="row">
                <div class="col-12 col-sm-6">
                    <div class="input-group">
                        <button class="btn btn-outline-secondary" type="button"
                        onclick="add_new_left()">Добавить</button>
                        <select id="left-selector" class="form-select">
                            <option value="Демонтаж-монтаж" selected>Демонтаж-монтаж</option>
                            <option value="Монтаж">Монтаж</option>
                            <option value="Покраска">Покраска</option>
                            <option value="Частичное нанесение">Частичное нанесение</option>
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

function add_new_left() {
    let val = document.getElementById("left-selector").value;
    if (reportData.left === null) {reportData.left = []}
    if (val == "Работа не требуется") {
        reportData.left = [];
    }
    if (reportData.left.some((el) => el.type == val)) {
         return
    }
    reportData.left.push({
        type: val,
        count: 1
    });
    render_data_to_form();
}

function changeWorkCounter(event) {
    let index = Number(event.currentTarget.getAttribute("data-index"));
    let val = Number(event.currentTarget.value);
    reportData.data[index].count = val;
    render_data_to_form();
}

function changeLeftWorkCounter(event) {
    let index = Number(event.currentTarget.getAttribute("data-index"));
    let val = Number(event.currentTarget.value);
    reportData.left[index].count = val;
}

function deleteWork(event) {
    let index = Number(event.currentTarget.getAttribute("data-index"));
    if (reportData.data[index].type == "Работа не требуется") {
        reportData.notRequire = false;
    }
    reportData.data.splice(index, 1);
    render_data_to_form();
}

function deleteLeftWork(event) {
    let index = Number(event.currentTarget.getAttribute("data-index"));
    reportData.left.splice(index, 1);
    render_data_to_form();
}

function markingChangeNumber(event) {
    let val = event.currentTarget.value;
    let index = Number(event.currentTarget.getAttribute("data-index"));
    reportData.data[index]["number"] = val;
    render_data_to_form();
}

function markingClick(event) {
    let markID = Number(event.currentTarget.getAttribute("data-id"));
    let index = Number(event.currentTarget.getAttribute("data-index"));
    if (reportData.data[index].selectedMarks === undefined) {
        reportData.data[index].selectedMarks = [];
    }
    if (reportData.data[index].selectedMarks.some((el) => el == markID)) {
        reportData.data[index].selectedMarks = reportData.data[index].selectedMarks.
        filter((id) => id != markID);
    } else {
        reportData.data[index].selectedMarks.push(markID);
    }
    render_data_to_form();
}

function render_inspection_to_form() {
    let container = document.getElementById("report-data");
    container.innerHTML = reportData.data === null ? ``:
    `
    <div class="container">
        <div class="row">
            ${reportData.data.reduce((acc, el, index) => {
                return acc + 
                `
                <div class="card mt-1 col-lg-6 col-12 p-0">
                    <div class="card-header">
                        <div class="row">
                            <div class="col-8">
                                ${el.type}
                            </div>
                            <div class="col-4">
                                <button data-index="${index}" type="button"
                                class="btn btn-outline-danger btn-sm float-end" 
                                onclick="deleteWork(event)">
                                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash" viewBox="0 0 16 16">
                                        <path d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5m3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0z"/>
                                        <path d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4zM2.5 3h11V2h-11z"/>
                                    </svg>
                                </button>
                            </div>
                        </div>
                    </div>
                    <div class="card-body row">
                        ${el.type == "Работа не требуется" ? 
                            `
                            <h6>Точка не требует выполнения работ</h6>
                            `
                            :
                            ``
                        }
                        ${el.type != "Работа не требуется" ? 
                            `
                            <label class="col-auto col-form-label">Количество:</label>
                            <div class="col-4 p-0">
                                <input type="number" class="form-control" value="${el.count}" min="1"
                                data-index="${index}" onchange="changeWorkCounter(event)">
                            </div>
                            `:``
                        }
                    </div>
                </div>
                `
            }, "")}
        </div>
    </div>
    `
    render_add_button_to_form();
    render_inspection_yourself_paint();
    container.innerHTML +=
    `
    <h5 class="mt-1">Оставить комментарий</h5>
    <textarea class="form-control" onchange="commentChange(event)"
    rows="3" >${reportData.comment === null ? "": reportData.comment}</textarea>
    `
}

function render_inspection_yourself_paint() {
    let container = document.getElementById("report-data");
    if (reportData.data !== null && reportData.data.some((el) => el.type == "Покраска")) {
        let found = reportData.data.find((el) => el.type == "Покраска");
        let paintCount = found.count;
        container.innerHTML += reportData.yourself === null ?
        `
        <h5 class="mt-1">Выполните покраску самостоятельно?</h5>
        <div class="container">
            <div class="row">
                <button id="paint-yourself-button-yes" type="button"
                class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="paint_yourself_changed(event)">Да</button>
                <button id="paint-yourself-button-no" type="button"
                class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
                onclick="paint_yourself_changed(event)">Нет</button>
            </div>
        </div>
        `
        :
        `
        ${reportData.yourself ? 
            `
            <h5 class="mt-1">Выполните покраску самостоятельно?</h5>
            <div class="container">
                <div class="row">
                    <button id="paint-yourself-button-yes" type="button"
                    class="btn btn-secondary col-sm-1 col-5 mx-1 mb-2"
                    onclick="paint_yourself_changed(event)" disabled>Да</button>
                    <button id="paint-yourself-button-no" type="button"
                    class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
                    onclick="paint_yourself_changed(event)">Нет</button>
                </div>
            </div>
            <label class="col-auto col-form-label">Укажите количество покрасок:</label>
            <div class="col-6 col-sm-2">
                <input type="number" class="form-control" min="1" max="${paintCount}"
                value="${reportData.paintCount}"
                onchange="changeYourselfCounter(event)">
            </div>
            `
            :
            `
            <h5 class="mt-1">Выполните покраску самостоятельно?</h5>
            <div class="container">
                <div class="row">
                    <button id="paint-yourself-button-yes" type="button"
                    class="btn btn-outline-secondary col-sm-1 col-5 mx-1 mb-2"
                    onclick="paint_yourself_changed(event)">Да</button>
                    <button id="paint-yourself-button-no" type="button"
                    class="btn btn-secondary col-sm-1 col-5 mx-1 mb-2"
                    onclick="paint_yourself_changed(event)" disabled>Нет</button>
                </div>
            </div>
            `
        }
        `
    };
}

function paint_yourself_changed(event) {
    let id = event.currentTarget.id;
    switch(id) {
        case "paint-yourself-button-yes":
            reportData.yourself = true;
            reportData.paintCount = 1;
            break;
        case "paint-yourself-button-no":
            reportData.yourself = false;
            break;
    }
    render_data_to_form();
}

function changeYourselfCounter(event) {
    let paintCount = reportData.data.find((el) => el.type == "Покраска").count;
    let val = Number(event.currentTarget.value);
    if (val <= paintCount && val > 0) {
        reportData.paintCount = val;
    }
    render_data_to_form();
}

let mediaCounter;

function render_load_media() {
    let container = document.getElementById("report-data");
    mediaCounter = 0;
    container.innerHTML += 
    `
    <h5 class="mt-1" id="media-tittle">Загрузите медиафайлы:</h5>
    <div class="container">
        <div class="row" id="report-media"></div>
    </div>
    `
    render_different_media_sets();
}

function render_different_media_sets() {
    if (reportData.type == "decline") {
        switch(reportData.reason) {
        case "Идет благоустройство - требуется забрать дуги":
            render_set_service_media();
            break;
        case "Идет благоустройство - требуется демонтировать и забрать дуги":
            render_set_service_media();
            break;
        case "Точка является дублем":
            document.getElementById("media-tittle").remove();
            break;
        default:
            render_set_inspection_media();
            break;
        }
    } else if (reportData.type == "service") {
        // console.log(selectedTasks);
        let parsedServiceWorks = parseServiceWorks();
        // console.log(parsedServiceWorks);
        if (parsedServiceWorks.some((el) => el.type == "Демонтаж-монтаж") && 
        reportData.carry) {
            render_set_service_extended_media();
        } else if (parsedServiceWorks.some((el) => el.type == "Демонтаж-монтаж")) {
            render_set_service_media();
        } else if (parsedServiceWorks.some((el) => el.type == "Демонтаж")) {
            render_set_service_media();
        } else if (parsedServiceWorks.some((el) => el.type == "Покраска комплекс")) {
            render_set_service_media();
        } else if (parsedServiceWorks.some((el) => el.type == "Покраска")) {
            render_set_service_media();
        } else {
            render_set_inspection_media();
        }
        if (selectedTasks.some((el) => el.type == "Заделать отверстия")) {
            render_set_hole_media();
        }
        if (parsedServiceWorks.some((el) => el.type == "Частичное нанесение")) {
            render_set_demark_media();
        }
        if (parsedServiceWorks.some((el) => el.type == "Демаркировка")) {
            render_set_demark_media();
        }
        parsedServiceWorks.filter((el) => el.type == "Нанесение разметки").forEach((el) => {
            render_set_mark_media();
        });
    } else if (reportData.type == "inspection") {
        // console.log(selectedTasks);
        let parsedInspectionWorks = parseInspectionWorks();
        // console.log(parsedInspectionWorks);
        render_set_inspection_media();
        if (reportData.data.some((el) => el.type == "Покраска") && reportData.yourself) {
            render_set_paint_inspection_media();
        }
    }
    let valid = validateMedia();
    render_report_footer(valid);
}

function parseServiceWorks() {
    let serviceWorks = [];
    reportData.data.forEach((el) => {
        serviceWorks.push(structuredClone(el));
    });
    if (serviceWorks.some((el) => el.type == "Демонтаж") &&
    serviceWorks.some((el) => el.type == "Монтаж")) {
        let dismount = serviceWorks.find((el) => el.type == "Демонтаж");
        let mount = serviceWorks.find((el) => el.type == "Монтаж");
        if (dismount.count > mount.count) {
            dismount.count = dismount.count - mount.count;
            mount.type = "Демонтаж-монтаж";
        } else if (mount.count > dismount.count) {
            mount.count = mount.count - dismount.count;
            dismount.type = "Демонтаж-монтаж";
        } else {
            dismount.type = "Демонтаж-монтаж";
            serviceWorks = serviceWorks.filter((el) => el.type != "Монтаж");
        }
    }
    if (serviceWorks.some((el) => el.type == "Покраска") &&
    serviceWorks.some((el) => el.type != "Покраска")) {
        let paint = serviceWorks.find((el) => el.type == "Покраска");
        paint.type = "Покраска комплекс";
    }
    return serviceWorks;
}

function parseInspectionWorks() {
    let inspectionWorks = [];
    reportData.data.forEach((el) => {
        inspectionWorks.push(structuredClone(el));
    });
    if (inspectionWorks.some((el) => el.type == "Покраска") && reportData.yourself) {
        let paint = inspectionWorks.find((el) => el.type == "Покраска");
        paint.count = paint.count - reportData.paintCount;
        if (paint.count == 0) {
            inspectionWorks = inspectionWorks.filter((el) => el != paint);
        }
    }
    return inspectionWorks;
}

function render_set_service_extended_media() {
    render_loader_by_name("Фото до", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото слева(старое место)", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото спереди(старое место)", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото справа(старое место)", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото слева(новое место)", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото спереди(новое место)", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото справа(новое место)", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Видео", mediaCounter);
    mediaCounter++;
}

function render_set_service_media() {
    render_loader_by_name("Фото до", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото слева", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото спереди", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото справа", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Видео", mediaCounter);
    mediaCounter++;
}

function render_set_inspection_media() {
    render_loader_by_name("Фото слева", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото спереди", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото справа", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Видео", mediaCounter);
    mediaCounter++;
}

function render_set_paint_inspection_media() {
    render_loader_by_name("Фото до покраски", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото слева после покраски", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото спереди после покраски", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото справа после покраски", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Видео после покраски", mediaCounter);
    mediaCounter++;
}

function render_set_hole_media() {
    render_loader_by_name("Фото отверстия", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото отверстия", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото отверстия", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото отверстия", mediaCounter);
    mediaCounter++;
}

function render_set_mark_media() {
    render_loader_by_name("Фото разметки слева", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото разметки спереди", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото разметки справа", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Видео разметки", mediaCounter);
    mediaCounter++;
}

function render_set_demark_media() {
    render_loader_by_name("Фото разметки до", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото разметки слева", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото разметки спереди", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Фото разметки справа", mediaCounter);
    mediaCounter++;
    render_loader_by_name("Видео разметки", mediaCounter);
    mediaCounter++;
}

function render_loader_by_name(name, mediaCounter) {
    let res = 
    `
    <div class="col-sm-3 col-6">
        <div class="card">
            <div class="card-header">${name}</div>
            <div class="card-body d-flex justify-content-center">
                <input type="file" id="file${mediaCounter}"
                data-id="${mediaCounter}" data-name="${name}" onchange="inputMediaChanged(event)"
                style="display: none;" />
                <div onclick="loadMedia(event)" data-id="${mediaCounter}" id="loader${mediaCounter}"
                class="d-flex justify-content-center">
                    <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" fill="currentColor" class="bi bi-upload" viewBox="0 0 16 16">
                      <path d="M.5 9.9a.5.5 0 0 1 .5.5v2.5a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.5a.5.5 0 0 1 1 0v2.5a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2v-2.5a.5.5 0 0 1 .5-.5"/>
                      <path d="M7.646 1.146a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1-.708.708L8.5 2.707V11.5a.5.5 0 0 1-1 0V2.707L5.354 4.854a.5.5 0 1 1-.708-.708z"/>
                    </svg>
                </div>
            </div>
        </div>
    </div>
    `
    let container = document.getElementById("report-media");
    container.innerHTML += res;
}

function loadMedia(event) {
    let id = Number(event.currentTarget.getAttribute("data-id"));
    let input = document.getElementById("file" + id);
    if (input) {
        input.click();
    }
}

function inputMediaChanged(event) {
    let loaderID = Number(event.currentTarget.getAttribute("data-id"));
    const mediaURL = window.URL.createObjectURL(event.currentTarget.files[0]);
    let mediaType = event.currentTarget.files[0].type.split("/")[0];
    uploadMediaToForm(loaderID, mediaURL, mediaType);
}

function uploadMediaToForm(id, mediaURL, mediaType) {
    let container = document.getElementById("loader" + id);
    container.onclick = null;
    let res = "";
    if (mediaType == "image") {
        res = 
        `
        <div class="d-flex justify-content-center">
            <a class="reportMedia d-flex align-items-center" data-gall="gallery-report"
            href="${mediaURL}">
                <img src="${mediaURL}" 
                alt="loading" style="max-height: 200px; max-width: 100%; border-radius: 5px;"/>
            </a>
        </div>
        `
    } else if (mediaType == "video") {
        res = 
        `
        <div class="d-flex justify-content-center">
            <a class="reportMedia d-flex align-items-center"
            data-gall="gallery-report" data-autoplay="true"
            data-vbtype="video"
            href="${mediaURL}">
                <video preload="metadata"
                style="max-height: 200px; max-width: 100%; border-radius: 5px;">
                    <source src="${mediaURL}#t=0.5" type="video/mp4" />
                </video>
            </a>
        </div>
        `
    }
    container.innerHTML = res;
    new VenoBox({
        selector: '.reportMedia',
        numeration: true,
        infinigall: true,
        share: true,
        spinner: 'circle'
    });
    let valid = validateMedia();
    render_report_footer(valid);
}

function validateMedia() {
    for (let j = 0; j < mediaCounter; j++) {
        let input = document.getElementById("file" + j);
        if (input.value === null || input.value == "") {
            return false;
        }
    }
    return true;
}