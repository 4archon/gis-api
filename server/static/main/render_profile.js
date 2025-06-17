let tasks;


function render_profile_info(profile) {
    result = `
    <h5 class="mt-2">Данные точки:
    <button data-id="${profile.id}" onclick="historyClick(event)"
    type="button" class="btn btn-info btn-sm">История точки</button>
    </h5>
    <div class="card-body">
        <ul class="list-group list-group-flush">
            <li class="list-group-item">Статус:
            ${profile.active ? "Активная точка" : "Деактивирована"}</li>
            <li class="list-group-item">Адрес:
            ${profile.address === null ? "Не указано" : profile.address}</li>
            <li class="list-group-item">Округ:
            ${profile.district === null ? "Не указано" : profile.district}</li>
            <li class="list-group-item">Координаты:
            ${profile.coordinates}</li>
            <li class="list-group-item">Количество дуг:
            ${profile.numberArc === null ? "Не указано" : profile.numberArc}</li>
            <li class="list-group-item">Тип дуги:
            ${profile.arcType === null ? "Не указано" : profile.arcType}</li>
            <li class="list-group-item">Покрытие:
            ${profile.carpet === null ? "Не указано" : profile.carpet}</li>
            <li class="list-group-item">Дата последних изменений данных:
            ${profile.changeDate === null ? "Не указано" : new Date(profile.changeDate).toLocaleDateString()}</li>
        </ul>
    </div>
    <h5 class="mt-2">Недавние материалы:</h5>
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
    conteiner.className = "row mt-3"
    let medias = await getRecentMedia(id);
    let result = ``;
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
    <span class="badge text-bg-danger">
    ${profile.deadline === null ? "Без дедлайна" : 
        new Date(profile.deadline).toLocaleDateString()}</span>
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
    res = await response.json();
    tasks = res.tasks;
    return res;
}

function sort_works(data) {
    targetWorks = [];
    if (data.works !== null) {
        data.works.forEach((element) => {
            let found = targetWorks.find((el) => el.work == element.work);
            if (found === undefined) {
                targetWorks.push(element);
            } else {
                if (element.arc > found.arc) {
                    found.arc = element.arc;
                }
            }
        });
        data.works = targetWorks;
    }

    if (targetWorks.length > 1) {
        targetWorks = targetWorks.filter((el) => el.work != "Работа не требуется");
    }

    return targetWorks;
}

async function render_profile_tasks(id) {
    let data = await getCurrentTasks(id);
    // console.log(data);
    data.works = sort_works(data);
    let result = data === null ? "": 
    `
    <h5 class="">Задачи:</h5>
    <div class="accordion accordion-flush">
        ${data.tasks === null ? 
            `
            <ul class="list-group list-group-flush">
            <li class="list-group-item">Задачи не выставлены</li>
            </ul>
            `:
            data.tasks.reduce((acc, el) => {
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
                    ${el.comment === null ? "Нет комменатария": el.comment}</div>
                </div>
            </div>
            `
        }, "")}
    </div>
    <h5 class="mt-2">Результаты инспекции:</h5>
    <ul class="list-group list-group-flush">
        ${data.works === null ? "Нет релевантных данных по инспекции":
            data.works.reduce((acc, el) => {
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
    if (data.appointed) {
        cont.innerHTML = `
        <button data-id="${data.id}" type="button" class="btn btn-primary" onclick="reportClick(event)">
            Отправить отчёт
        </button>
        `
    } else {
        cont.innerHTML = "";
    }
}

function pointClick(event) {
    render_profile_header(event.targetData.userData);
    render_profile_info(event.targetData.userData);
    render_profile_footer(event.targetData.userData);
    let body = document.getElementById("point-profile-body");
    render_profile_tasks(event.targetData.userData.id);
    render_profile_media(event.targetData.userData.id).then((element) => {
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
