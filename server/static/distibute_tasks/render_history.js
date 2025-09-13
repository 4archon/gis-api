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