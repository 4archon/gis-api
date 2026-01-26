let pointHistory = new bootstrap.Modal(document.getElementById("point-history"), null);

async function getHistory(id) {
    url = "/history"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "text/plain"
        },
        body: id
    })
    res = await response.json();
    return res;
}

async function historyClick(event) {
    let id = event.target.getAttribute("data-id");
    let data = await getHistory(id);
    console.log(data);
    render_history_header(data);
    render_history_body(data);
    pointHistory.show();
}

function render_history_header(data) {
    result = `
    <h1 class="modal-title fs-5">История точки 
    <span class="badge text-bg-primary">${data.id}</span>
    </h1>
    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
    `
    let header = document.getElementById("point-history-header");
    header.innerHTML = result;
}

function render_history_body(data) {
    let header = document.getElementById("point-history-body");
    header.innerHTML = "";
    data.storyPoints.forEach((element) => {
        header.innerHTML += `
        <div class="card">
            <div class="card-header">
                Сервис номер: ${element.id}
                ${element.sent != false ? "":
                `<span class="badge text-bg-primary" style="font-size: 11pt;">
                В работе</span>`}
                <button type="button" class="btn btn-outline-secondary
                ${element.invisible ? "active" : ""}
                btn-sm ms-1" data-id="${element.id}"
                data-invisible="${element.invisible ? true : false}"
                onclick="history_hide_service(event);">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-eye-slash" viewBox="0 0 16 16">
                        <path d="M13.359 11.238C15.06 9.72 16 8 16 8s-3-5.5-8-5.5a7 7 0 0 0-2.79.588l.77.771A6 6 0 0 1 8 3.5c2.12 0 3.879 1.168 5.168 2.457A13 13 0 0 1 14.828 8q-.086.13-.195.288c-.335.48-.83 1.12-1.465 1.755q-.247.248-.517.486z"/>
                        <path d="M11.297 9.176a3.5 3.5 0 0 0-4.474-4.474l.823.823a2.5 2.5 0 0 1 2.829 2.829zm-2.943 1.299.822.822a3.5 3.5 0 0 1-4.474-4.474l.823.823a2.5 2.5 0 0 0 2.829 2.829"/>
                        <path d="M3.35 5.47q-.27.24-.518.487A13 13 0 0 0 1.172 8l.195.288c.335.48.83 1.12 1.465 1.755C4.121 11.332 5.881 12.5 8 12.5c.716 0 1.39-.133 2.02-.36l.77.772A7 7 0 0 1 8 13.5C3 13.5 0 8 0 8s.939-1.721 2.641-3.238l.708.709zm10.296 8.884-12-12 .708-.708 12 12z"/>
                    </svg>
                </button>
            </div>
            <div class="card-body">
                <h5 class="card-title">Общая информация:</h5>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item">Исполнители: 
                    ${element.userLogins === null ? "Не указано":
                        element.userLogins.reduce((acc, el, index) => {
                        return acc +=
                        `<span class="badge text-bg-primary">` +
                        element.userIDs[index] +
                        `</span>`
                        + " " + el
                        + (index != element.userLogins.length - 1 ? ", ": "");
                    }, "")}</li>
                    <li class="list-group-item">Дата исполнения:
                    ${element.execution === null ? "Не указано": 
                    new Date(element.execution).toLocaleDateString()}
                    </li>
                    <li class="list-group-item">Статус:
                    ${element.status === null ? "Не указан": element.status}</li>
                    <li class="list-group-item">Комментарий: 
                    ${element.comment === null || element.comment == "" ?
                    "Нет комментария": element.comment}</li>
                    <li class="list-group-item"></li>
                </ul>
                ${element.tasks === null ? "":
                `<h5 class="card-title">Выполненные задачи:</h5>`}
                <div class="accordion accordion-flush">
                    ${element.tasks === null ? "":
                        element.tasks.reduce((acc, el) => {
                        return acc +=
                        `
                        <div class="accordion-item">
                            <h2 class="accordion-header">
                                <button class="accordion-button collapsed" type="button"
                                data-bs-toggle="collapse"
                                data-bs-target="#task${el.id}" aria-expanded="false"
                                aria-controls="task${el.id}">
                                    ${el.type}
                                </button>
                            </h2>
                            <div id="task${el.id}" class="accordion-collapse collapse">
                                <div class="accordion-body">
                                    <span class="badge text-bg-danger">
                                        ${el.deadline === null ? "Без дедлайна":
                                        new Date(el.deadline).toLocaleDateString()}
                                    </span>
                                    <span class="badge text-bg-danger">
                                        ${el.customer === null ? "Заказчик не указан":el.customer}
                                    </span>
                                    <br>
                                    ${el.comment === null ? "Нет комменатария": el.comment}
                                </div>
                            </div>
                        </div>
                        `
                    }, "")}
                </div>
                ${element.works === null ? "":
                `<h5 class="card-title">Результат:</h5>`}
                <ul class="list-group list-group-flush">
                    ${element.works === null ? "":
                        element.works.filter((el) => el.type == "required")
                        .every((el) => el.work == "Работа не требуется") ?
                        `<li class="list-group-item">Работа не требуется</li>`:
                        element.works.filter((el) => el.type == "required").reduce((acc, el) => {
                            return acc +=
                            `
                            <li class="list-group-item">
                            ${el.type == "done" ? "Выполнено: ": "Требуется выполнить: "}    
                            ${el.work}, количество дуг: ${el.arc} 
                            </li>
                            `
                        }, "")
                    }
                    ${element.works === null ? "":
                        element.works.filter((el) => el.type == "done").reduce((acc, el) => {
                            return acc +=
                            `
                            <li class="list-group-item">
                            ${el.type == "done" ? "Выполнено: ": "Требуется выполнить: "}    
                            ${el.work}, количество дуг: ${el.arc} 
                            </li>
                            `
                        }, "")
                    }
                </ul>
                <h5 class="card-title">Материалы:</h5>
                <div class="row">
                    ${element.medias === null ? "":
                        element.medias.reduce((acc, el) => {
                            return acc +=
                            `
                            ${el.type == "mov"? `
                                <div class="col-2 d-flex justify-content-center">
                                    <a class="d-flex align-items-center history-media" data-gall="history-${element.id}" data-autoplay="true"
                                    data-vbtype="video"
                                    href="/media/${el.id}.${el.type}">
                                        <svg style="max-height: 200px; max-width: 100%; border-radius: 5px;" 
                                        xmlns="http://www.w3.org/2000/svg" width="50" height="50" fill="currentColor" class="bi bi-play-circle" viewBox="0 0 16 16">
                                            <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14m0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16"/>
                                            <path d="M6.271 5.055a.5.5 0 0 1 .52.038l3.5 2.5a.5.5 0 0 1 0 .814l-3.5 2.5A.5.5 0 0 1 6 10.5v-5a.5.5 0 0 1 .271-.445"/>
                                        </svg>
                                    </a>
                                </div>
                                `:`
                                <div class="col-2 d-flex justify-content-center">
                                    <a class="d-flex align-items-center history-media" data-gall="history-${element.id}" href="/media/${el.id}.${el.type}">
                                        <img src="/media/${el.id}.${el.type}" loading="lazy" 
                                        alt="loading" style="max-height: 200px; max-width: 100%; border-radius: 5px;"/>
                                    </a>
                                </div>
                                `}
                            `
                        }, "")}
                </div>
            </div>
        </div>
        `
    });
    new VenoBox({
            selector: '.history-media',
            numeration: true,
            infinigall: true,
            share: true,
            spinner: 'circle',
            fitView: true
        });
}


function history_hide_service(event) {
    let target = event.currentTarget;
    target.disabled = true;
    
    let serviceID = target.getAttribute("data-id");
    let invisible = target.getAttribute("data-invisible") == "true" ? true : false;
    
    postChangeServiceInvisible(serviceID, !invisible, target).then(() => {
        target.disabled = false;
    });
}

async function postChangeServiceInvisible(serviceID, invisible, target) {
    let formData = new FormData();
    formData.append("id", serviceID);
    formData.append("value", invisible);

    let url = "/change_service_invisible"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        body: formData
    })
    let res = await response;
    if (res.ok) {
        target.setAttribute("data-invisible", invisible ? "true" : "false");
        if (invisible) {
            target.classList.add("active");
        } else {
            target.classList.remove("active");
        }
    }
}