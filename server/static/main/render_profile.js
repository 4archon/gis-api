function render_profile_info(profile) {
    result = `
    <button data-id="${profile.id}" onclick="historyClick(event)"
    type="button" class="btn btn-info btn-sm mb-2">История точки</button>
    <h5>Данные точки:</h5>
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
    return result;
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
            res = `
            <div class="col-3 d-flex justify-content-center">
                <a class="profileMedia" data-gall="gallery-profile" data-autoplay="true"
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
            res = `
            <div class="col-3 d-flex justify-content-center">
                <a class="profileMedia" data-gall="gallery-profile" href="/media/${element.id}.${element.type}">
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
    return result;
}


function pointClick(event) {
    document.getElementById("point-profile-header").innerHTML = render_profile_header(event.targetData.userData);
    res = render_profile_info(event.targetData.userData);
    let body = document.getElementById("point-profile-body")
    body.innerHTML = res;
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
