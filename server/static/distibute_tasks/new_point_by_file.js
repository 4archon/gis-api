let newPointByFileModal = new bootstrap.Modal(document.getElementById("point-new-by-file"));

let newPointsFile = null;

document.getElementById("new-points-by-file-button").onclick = newPointByFileMenu;

function newPointByFileMenu() {
    newPointsFile = null;
    render_new_point_by_file_body();
    render_new_point_by_file_footer();
    newPointByFileModal.show();
}

function render_new_point_by_file_body() {
    let container = document.getElementById("point-new-by-file-body");
    container.innerHTML = 
    `
    <div class="container">
        <div class="row">
            <div class="col-sm-6 col-12">
                <div class="card text-center">
                    <div class="card-body d-flex justify-content-center">
                        <div onclick="dowloadNewPointsByFileExample()"
                        class="d-flex justify-content-center">
                            <svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" fill="currentColor" class="bi bi-file-earmark" viewBox="0 0 16 16">
                                <path d="M14 4.5V14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h5.5zm-3 0A1.5 1.5 0 0 1 9.5 3V1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4.5z"/>
                            </svg>
                        </div>
                    </div>
                    <div class="card-footer text-muted">Пример файла</div>
                </div>
            </div>
            <div class="col-sm-6 col-12">
                <div class="card text-center">
                    <div class="card-body d-flex justify-content-center">
                        <div id="uploader-new-point-by-file"
                        onclick="selectNewPointsByFile()"
                        class="d-flex justify-content-center">
                            <svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" fill="currentColor" class="bi bi-file-earmark-arrow-down" viewBox="0 0 16 16">
                                <path d="M8.5 6.5a.5.5 0 0 0-1 0v3.793L6.354 9.146a.5.5 0 1 0-.708.708l2 2a.5.5 0 0 0 .708 0l2-2a.5.5 0 0 0-.708-.708L8.5 10.293z"/>
                                <path d="M14 14V4.5L9.5 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2M9.5 3A1.5 1.5 0 0 0 11 4.5h2V14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.5z"/>
                            </svg>
                        </div>
                    </div>
                    <div class="card-footer text-muted" id="file-name-new-point-by-file">
                        Загрузите csv файл
                    </div>
                </div>
            </div>
        </div>
        <div class="row mt-3">
            <div class="col-12">
                <h5>Инструкция:</h5>
                <ol class="m-0">
                    <h6><li>Скачать файл с примером</li></h6>
                    <h6><li>Создать новую таблицу в Google Spreadsheets</li></h6>
                    <h6><li>Открыть файл с примером в Google Spreadsheets(Файл >> Импортировать >> Добавить >> Вставить листы)</li></h6>
                    <h6><li>Выделить первые две строчки на добавленном листе и выбрать Преобразовать в таблицу</li></h6>
                    <h6><li>Заполнить таблица в соответствии с типами данных в колонках</li></h6>
                    <h6><li>Выгрузить csv(Файл >> Скачать >> Формат CSV)</li></h6>
                    <h6><li>Загрузить csv файл в форму выше</li></h6>
                    <h6><li>Подтвердить кнопкой добавить</li></h6>
                </ol>
            </div>
        </div>
    </div>
    `
}

function dowloadNewPointsByFileExample() {
    window.location.href = "/static/other/example.xlsx";
}

function selectNewPointsByFile() {
    newPointsFile = document.createElement("input");
    newPointsFile.type = "file";
    newPointsFile.onchange = loadNewPointsByFile;
    newPointsFile.click();
}

function loadNewPointsByFile(event) {
    let container = document.getElementById("uploader-new-point-by-file");
    if (event.target.files[0].type != "text/csv") {
        container.innerHTML = 
        `
        <svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" fill="currentColor" class="bi bi-file-earmark-arrow-down" viewBox="0 0 16 16">
            <path d="M8.5 6.5a.5.5 0 0 0-1 0v3.793L6.354 9.146a.5.5 0 1 0-.708.708l2 2a.5.5 0 0 0 .708 0l2-2a.5.5 0 0 0-.708-.708L8.5 10.293z"/>
            <path d="M14 14V4.5L9.5 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2M9.5 3A1.5 1.5 0 0 0 11 4.5h2V14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.5z"/>
        </svg>
        `
        document.getElementById("file-name-new-point-by-file").innerHTML = 
        `Загрузите csv файл`
    } else {
        container.innerHTML = 
        `
        <svg xmlns="http://www.w3.org/2000/svg" width="100" height="100" fill="green" class="bi bi-filetype-csv" viewBox="0 0 16 16">
            <path fill-rule="evenodd" d="M14 4.5V14a2 2 0 0 1-2 2h-1v-1h1a1 1 0 0 0 1-1V4.5h-2A1.5 1.5 0 0 1 9.5 3V1H4a1 1 0 0 0-1 1v9H2V2a2 2 0 0 1 2-2h5.5zM3.517 14.841a1.13 1.13 0 0 0 .401.823q.195.162.478.252.284.091.665.091.507 0 .859-.158.354-.158.539-.44.187-.284.187-.656 0-.336-.134-.56a1 1 0 0 0-.375-.357 2 2 0 0 0-.566-.21l-.621-.144a1 1 0 0 1-.404-.176.37.37 0 0 1-.144-.299q0-.234.185-.384.188-.152.512-.152.214 0 .37.068a.6.6 0 0 1 .246.181.56.56 0 0 1 .12.258h.75a1.1 1.1 0 0 0-.2-.566 1.2 1.2 0 0 0-.5-.41 1.8 1.8 0 0 0-.78-.152q-.439 0-.776.15-.337.149-.527.421-.19.273-.19.639 0 .302.122.524.124.223.352.367.228.143.539.213l.618.144q.31.073.463.193a.39.39 0 0 1 .152.326.5.5 0 0 1-.085.29.56.56 0 0 1-.255.193q-.167.07-.413.07-.175 0-.32-.04a.8.8 0 0 1-.248-.115.58.58 0 0 1-.255-.384zM.806 13.693q0-.373.102-.633a.87.87 0 0 1 .302-.399.8.8 0 0 1 .475-.137q.225 0 .398.097a.7.7 0 0 1 .272.26.85.85 0 0 1 .12.381h.765v-.072a1.33 1.33 0 0 0-.466-.964 1.4 1.4 0 0 0-.489-.272 1.8 1.8 0 0 0-.606-.097q-.534 0-.911.223-.375.222-.572.632-.195.41-.196.979v.498q0 .568.193.976.197.407.572.626.375.217.914.217.439 0 .785-.164t.55-.454a1.27 1.27 0 0 0 .226-.674v-.076h-.764a.8.8 0 0 1-.118.363.7.7 0 0 1-.272.25.9.9 0 0 1-.401.087.85.85 0 0 1-.478-.132.83.83 0 0 1-.299-.392 1.7 1.7 0 0 1-.102-.627zm8.239 2.238h-.953l-1.338-3.999h.917l.896 3.138h.038l.888-3.138h.879z"/>
        </svg>
        `
        document.getElementById("file-name-new-point-by-file").innerHTML =
        `${event.target.files[0].name}`
    }
}

function render_new_point_by_file_footer() {
    let container = document.getElementById("point-new-by-file-footer");
    container.innerHTML = 
    `
    <button type="button" class="btn btn-secondary" onclick="closeNewPointByFileMenu()">
        Отменить
    </button>
    <button type="button" class="btn btn-primary" onclick="sendNewPointByFile()">
        Добавить
    </button>
    `
}

function closeNewPointByFileMenu() {
    newPointByFileModal.hide();
}

async function sendNewPointByFile() {
    if (newPointsFile) {
        let file = newPointsFile.files[0];
        if (file && file.type == "text/csv") {
            let formData = new FormData();
            formData.append("file", file);
            newPointByFileModal.hide();

            let url = "/new_points_by_file"
            let response = await fetch(url, {
                method: "POST",
                cache: "no-cache",
                credentials: "same-origin",
                body: formData
            })

            if (response.ok) {
                alert("Точки успешно добавлены")
            } else {
                alert("Произошла ошибка добавления точке через файл")
            }
        }
    }
}