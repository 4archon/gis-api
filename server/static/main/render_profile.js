let pointProfile = new bootstrap.Modal(document.getElementById("point-profile"), null);
let selectPoint = new bootstrap.Modal(document.getElementById("select-point"), null);


function render_profile_info(profile) {
    result = `
    <h5 class="mt-2">Данные точки:</h5>
    <div class="card-body">
        <ul class="list-group list-group-flush">
            <li class="list-group-item">Статус активации:
            ${profile.active ? "Активная точка" : "Деактивирована"}</li>
            <li class="list-group-item">Статус точки:
            ${profile.status === null ? "Не указано" : profile.status}</li>
            <li class="list-group-item">Адрес:
            ${profile.address === null ? "Не указано" : profile.address}</li>
            <li class="list-group-item">Округ:
            ${profile.district === null || profile.district === "" ? 
                "Не указано" : profile.district}</li>
            <li class="list-group-item">Координаты:
            ${profile.coordinates.toReversed()}</li>
            <li class="list-group-item">Количество дуг:
            ${profile.numberArc === null ? "Не указано" : profile.numberArc}</li>
            <li class="list-group-item">Тип дуги:
            ${profile.arcType === null || profile.arcType === "" ?
                "Не указано" : profile.arcType}</li>
            <li class="list-group-item">Покрытие:
            ${profile.carpet === null || profile.carpet === "" ?
                "Не указано" : profile.carpet}</li>
            <li class="list-group-item">Владелец:
            ${profile.owner === null ? "Не указано" : profile.owner}</li>
            <li class="list-group-item">Оператор:
            ${profile.operator === null ? "Не указано" : profile.operator}</li>
            <li class="list-group-item">Дата последних изменений данных:
            ${profile.changeDate === null ? "Не указано" :
                new Date(profile.changeDate).toLocaleDateString()}</li>
            <li class="list-group-item">Комментарий к точке:
            ${profile.comment === null || profile.comment == "" ?
                "Не указано" : profile.comment}</li>
        </ul>
    </div>
    `

    let body = document.getElementById("point-profile-body");
    body.innerHTML = result;
}

async function getRecentMedia(id) {
    url = "/recent_media"
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

async function render_profile_media(id) {
    let conteiner = document.createElement("div");
    conteiner.className = "row"
    let medias = await getRecentMedia(id);
    let result = ``
    if (medias.medias !== null) {
        result = `<h5 class="mt-2">Недавние материалы:</h5>`; 
    }
    medias.medias.forEach((element) => {
        let res;
        if (element.type == "mov") {
            res = 
            `
            <div class="col-3 d-flex justify-content-center">
                <a class="d-flex align-items-center profileMedia" data-gall="gallery-profile" data-autoplay="true"
                data-vbtype="video"
                href="/media/${element.id}.${element.type}">
                <video preload="metadata"
                style="max-height: 200px; max-width: 100%; border-radius: 5px;">
                    <source src="/media/${element.id}.${element.type}#t=0.5" type="video/mp4" />
                </video>
                </a>
            </div>
            `
        } else {
            res = 
            `
            <div class="col-3 d-flex justify-content-center">
                <a class="profileMedia d-flex align-items-center" data-gall="gallery-profile" href="/media/${element.id}.${element.type}">
                    <img src="/media/${element.id}.${element.type}" 
                    alt="loading" style="max-height: 200px; max-width: 100%; border-radius: 5px;"/>
                </a>
            </div>
            `
        }
        result += res
    })
    conteiner.innerHTML = result;
    return conteiner;
}

function render_profile_header(profile) {
    result = `
    <h1 class="modal-title fs-5">Профиль точки 
    <span class="badge text-bg-primary">${profile.id}</span>
    <span class="badge text-bg-primary"
    data-id="${profile.id}" onclick="historyClick(event)">
        История точки
    </span>
    </h1>
    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>`
    
    document.getElementById("point-profile-header").innerHTML = result;
}

async function getCurrentTasks(id) {
    url = "/current_tasks"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "text/plain"
        },
        body: id
    })
    let res = await response.json();
    return res;
}

function filterWorks(works) {
    let targetWorks = [];
    if (works !== null) {
        works.forEach((element) => {
            let found = targetWorks.find((el) => el.work == element.work);
            if (found === undefined) {
                targetWorks.push(element);
            } else {
                if (element.arc > found.arc) {
                    found.arc = element.arc;
                }
            }
        });
    }

    if (targetWorks.length > 1) {
        targetWorks = targetWorks.filter((el) => el.work != "Работа не требуется");
    }

    return targetWorks;
}

function filterProfileTasks(tasks) {
    if (tasks === null) {
        return tasks
    }
    if (userSubgroup == "inspection") {
        tasks = tasks.filter((el) => el.type == "Проинспектировать");
    } else if (userSubgroup == "service") {
        tasks = tasks.filter((el) => el.type != "Проинспектировать");
    }
    if (tasks.length == 0) {
        return null
    }
    return tasks;
}

async function render_profile_tasks(id) {
    if (data.find((el) => el.id == id).appoint === null) {
        return
    }
    let taskData = await getCurrentTasks(id);
    taskData.works = filterWorks(taskData.works);
    taskData.tasks = filterProfileTasks(taskData.tasks);
    let result = taskData === null ? "": 
    `
    <h5 class="">Задачи:</h5>
    <div class="accordion accordion-flush">
        ${taskData.tasks === null ? 
            `
            <ul class="list-group list-group-flush">
                <li class="list-group-item">Задачи не выставлены</li>
            </ul>
            `:
            taskData.tasks.reduce((acc, el) => {
            if (userSubgroup == "inspection" && el.type != "Проинспектировать") {
                return acc;
            } else if (userSubgroup == "service" && el.type == "Проинспектировать") {
                return acc;
            }
            return acc +=
            `
            <div class="card mt-1">
                <div class="card-body">
                    <h5>
                        ${el.type}
                        <span class="badge text-bg-danger">
                            ${el.deadline === null ? "Без дедлайна":
                                new Date(el.deadline).toLocaleDateString()}
                        </span>
                    </h5>
                    ${el.comment === null || el.comment == "" ? "":
                        `<p class="card-text">Комментарий: ${el.comment}</p>`}
                </div>
            </div>
            `
        }, "")}
    </div>
    <h5 class="mt-2">Результаты инспекции:</h5>
    <ul class="list-group list-group-flush">
        ${taskData.works === null || taskData.works.length == 0 ?
            `
            <ul class="list-group list-group-flush">
                <li class="list-group-item">Нет релевантных данных по инспекции</li>
            </ul>
            `
            :
            taskData.works.reduce((acc, el) => {
            return acc +=
            `
            <li class="list-group-item">    
            ${el.work}, количество дуг: ${el.arc} 
            </li>
            `
        }, "")}
    </ul>
    `

    let body = document.getElementById("point-profile-body");
    let cont = document.createElement("div");
    cont.innerHTML = result;
    body.prepend(cont);
}

function render_profile_footer(data) {
    let cont = document.getElementById("point-profile-footer")
    if (data.appoint || userTrust) {
        cont.innerHTML = `
        <button data-id="${data.id}" type="button" class="btn btn-primary" onclick="reportClick(event)">
            Отправить отчёт
        </button>
        `
    } else {
        cont.innerHTML = "";
    }
}

function render_profile_marking(data) {
    let result = `<h5 class="mt-2">Разметка:</h5>`;
    if (data.marks !== null) {
        result += data.marks.reduce((acc, el) => {
            return acc + 
            `
            <ul class="list-group list-group-flush">
                <li class="list-group-item">
                ${el.type}
                <span class="badge text-bg-primary">${el.number}</span>
                </li>
            </ul>
            `
        }, "");
        let body = document.getElementById("point-profile-body");
        body.innerHTML += result;
    }
}

function showPointProfile(pointID) {
    let point = data.find((el) => el.id == pointID);
    render_profile_header(point);
    render_profile_info(point);
    render_profile_marking(point);
    render_profile_footer(point);
    let body = document.getElementById("point-profile-body");
    render_profile_tasks(point.id);
    render_profile_media(point.id).then((element) => {
        body.append(element);
    }).then(() => {
        new VenoBox({
            selector: '.profileMedia',
            numeration: true,
            infinigall: true,
            share: true,
            spinner: 'circle'
        });
    });
    pointProfile.show();
}

function clusterClick(event) {
    if (event.target.type == "marker") {
        showPointProfile(event.target.data.id)
    } else {
        fillPointSelection(event.target.data);
        selectPoint.show();
    }
}

function fillPointSelection(points) {
    container = document.getElementById("select-point-body");
    container.innerHTML = 
    `
    <div class="accordion accordion-flush">
        ${points.reduce((acc, el) => {
            return acc +=
            `
            <div class="card mt-1
            ${el.deadline !== null ? "text-bg-danger bg-gradient":""}"
            data-id="${el.id}" onclick="pointSelected(event)">
                <div class="card-body">
                    <h5>Внутренний id: ${el.id}, внешний id: ${el.externalID}</h5>
                    <p class="card-text">
                    Адрес: ${el.address}
                    </p>
                </div>
            </div>
            `
        }, "")}
    </div>
    `
}

function pointSelected(event) {
    selectPoint.hide();
    showPointProfile(Number(event.currentTarget.getAttribute("data-id")));
}