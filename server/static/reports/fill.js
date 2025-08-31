let data;

async function getReports() {
    let url = window.location.pathname;
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin"
    })
    data = await response.json();
    console.log(data);
    renderCards();
    render_pages();
    new VenoBox({
        selector: '.history-media',
        numeration: true,
        infinigall: true,
        share: true,
        spinner: 'circle'
    });
}

function renderCards() {
    let root = document.getElementById("root");
    root.innerHTML = "";
    data.services.forEach(element => {
        let card = document.createElement("div")
        card.className = "card mt-1";
        card.innerHTML = render_card(element)
        root.appendChild(card);
    });
}

getReports();

function render_pages() {
    let container = document.getElementById("root");
    container.innerHTML +=
    `
    <nav class="mt-4" aria-label="Page navigation example">
        <ul class="pagination justify-content-center">
            <li class="page-item ${data.currentPage == 1 ? "disabled" : ""}">
                <a class="page-link" href="/reports/1">Первая</a>
            </li>
            <li class="page-item ${data.currentPage == 1 ? "disabled" : ""}">
                <a class="page-link" href="/reports/${data.currentPage - 1}">&laquo;</a>
            </li>
            ${data.currentPage == data.lastPage && data.currentPage - 2 >= 1 ?
                `
                <li class="page-item">
                    <a class="page-link"
                    href="/reports/${data.currentPage + 2}">${data.currentPage + 2}</a>
                </li>
                ` : ""
            }
            ${data.currentPage - 1 < 1 ? "" : 
                `
                <li class="page-item">
                    <a class="page-link"
                    href="/reports/${data.currentPage - 1}">${data.currentPage - 1}</a>
                </li>
                `
            }
            <li class="page-item active">
                <a class="page-link"
                href="/reports/${data.currentPage}">${data.currentPage}</a>
            </li>
            ${data.currentPage + 1 > data.lastPage ? "" :
                `
                <li class="page-item">
                    <a class="page-link"
                    href="/reports/${data.currentPage + 1}">${data.currentPage + 1}</a>
                </li>
                `
            }
            ${data.currentPage == 1 && data.currentPage + 2 <= data.lastPage ?
                `
                <li class="page-item">
                    <a class="page-link"
                    href="/reports/${data.currentPage + 2}">${data.currentPage + 2}</a>
                </li>
                ` : ""
            }
            <li class="page-item ${data.currentPage == data.lastPage ? "disabled" : ""}">
                <a class="page-link" href="/reports/${data.currentPage + 1}">&raquo;</a>
            </li>
            <li class="page-item ${data.currentPage == data.lastPage ? "disabled" : ""}">
                <a class="page-link" href="/reports/${data.lastPage}">Последняя</a>
            </li>
        </ul>
    </nav>
    `
}

function render_card(element) {
    res = 
    `
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
                ${element.users.reduce((acc, el, indx) => {
                    return acc +
                    `
                    <span class="badge text-bg-primary">${el.id}</span>
                    <span class="badge text-bg-primary">
                        ${el.subgroup === null ? "не указано" : el.subgroup}
                    </span>
                    ${el.login !== null ? el.login : ""}
                    ${element.users.length - 1 != indx ? ", " : ""}
                    `
                }, "")}
                </li>
                <li class="list-group-item">Отчет отправлен пользователем:
                    <span class="badge text-bg-primary">${element.sentBy.id}</span>
                    <span class="badge text-bg-primary">
                        ${element.sentBy.subgroup === null ? "не указано" : element.sentBy.subgroup}
                    </span>
                    ${element.sentBy.login !== null ? el.sentBy.login : ""}
                </li>
                <li class="list-group-item">Дата исполнения:
                ${element.execution === null ? "Не указано":
                new Date(element.execution).toLocaleDateString()}
                </li>
                <li class="list-group-item">Дата назначения:
                ${element.appoint === null ? "Не указано":
                new Date(element.appoint).toLocaleDateString()}
                </li>
                <li class="list-group-item">Статус:
                ${element.status === null ? "Не указан": element.status}</li>
                <li class="list-group-item">Выполнено с назначением: 
                ${element.withoutTask == false ? "да" : "нет"}
                </li>
                <li class="list-group-item">Комментарий: 
                ${element.comment === null || element.comment == "" ?
                "Нет комментария": element.comment}</li>
                <li class="list-group-item"></li>
            </ul>
            <h5 class="card-title">Данные точки
            <span class="badge text-bg-primary">${element.point.id}</span>
            </h5>
            <ul class="list-group list-group-flush">
                <li class="list-group-item">Координаты точки: ${element.point.lat}, ${element.point.long}</li>
                <li class="list-group-item">Адрес: ${element.point.address}</li>
                <li class="list-group-item">Количество дуг: ${element.point.numberArc}</li>
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
    return res;
}