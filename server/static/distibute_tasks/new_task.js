let newTask = new bootstrap.Modal(document.getElementById("new-task"));

document.getElementById("selected-points-new-task").onclick = newTaskSelectedPoints;

function newTaskSelectedPoints() {
    if (selectedPoints.length == 0) {return};
    render_new_task_header_selected_points();
    render_new_task_body_selected_points();
    render_new_task_footer_selected_points();
    newTask.show();
}

function render_new_task_header_selected_points() {
    let container = document.getElementById("new-task-header");
    container.innerHTML = 
    `
    <h1 class="modal-title fs-5">Добавить новую задачу выбранным точкам</h1>
    <button type="button" class="btn-close"
    data-bs-dismiss="modal" aria-label="Close"></button>
    `
}

function render_new_task_body_selected_points() {
    let container = document.getElementById("new-task-body");
    container.innerHTML =
    `
    <div class="card-body mb-4">
        <div class="row g-3">
            <div class="col-md-6">
                <label for="inputTaskSubtype" class="form-label">Подтип</label>
                <select id="inputTaskSubtype" class="form-select"
                onchange="changeInputTaskType(event)">
                    <option value="Демонтаж">Демонтаж</option>
                    <option value="Сезонные">Сезонные</option>
                    <option value="Манипуляции над точкой">Манипуляции над точкой</option>
                    <option value="Разметка">Разметка</option>
                    <option value="Деактивация">Деактивация</option>
                    <option value="Стандартные" selected>Стандартные</option>
                    <option value="Другое">Другое</option>
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputTaskType" class="form-label">Задача</label>
                <select id="inputTaskType" class="form-select">
                    <option value="Проинспектировать" selected>Проинспектировать</option>
                    <option value="Произвести сервис">Произвести сервис</option>
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputTaskCustomer" class="form-label">Заказчик</label>
                <select id="inputTaskCustomer" class="form-select">
                    <option value="Яндекс" selected>Яндекс</option>
                    <option value="Whoosh">Whoosh</option>
                    <option value="Ultradop">Ultradop</option>
                    <option value="Другое">Другое</option>
                </select>
            </div>
            <div class="col-md-6">
                <label for="inputTaskDeadline" class="form-label">Срок исполнения до</label>
                <input id="inputTaskDeadline" class="form-control" type="date">
            </div>
            <div class="col-12">
                <label for="inputTaskComment" class="form-label">Комментарий для исполнителя</label>
                <textarea id="inputTaskComment" class="form-control" rows="4"></textarea>
            </div>
        </div>
    </div>
    `
}

function render_new_task_footer_selected_points() {
    let container = document.getElementById("new-task-footer");
    container.innerHTML = 
    `
    <button type="button" class="btn btn-primary" onclick="applyTaskToSelectedPoints()">
        Добавить
    </button>
    `
}

function applyTaskToSelectedPoints() {
    let data = {
        task: document.getElementById("inputTaskType").value,
        customer: document.getElementById("inputTaskCustomer").value,
        deadline: new Date(document.getElementById("inputTaskDeadline").value),
        comment: document.getElementById("inputTaskComment").value,
        points: selectedPoints.map((el) => el.id)
    }

    applyTaskBackend(data);
}

async function applyTaskBackend(data) {
    let url = "/new_task"
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
    newTask.hide();
}


function changeInputTaskType(event) {
    let container = document.getElementById("inputTaskType");
    let val = event.currentTarget.value;
    switch(val) {
    case "Демонтаж":
        container.innerHTML = 
        `
        <option value="Временный демонтаж по разным причинам" selected>
            Временный демонтаж по разным причинам
        </option>
        <option value="Благоустройство - временный демонтаж">
            Благоустройство - временный демонтаж
        </option>
        <option value="Заделать отверстия">
            Заделать отверстия
        </option>
        `
        break;
    case "Сезонные":
        container.innerHTML = 
        `
        <option value="Снятие дуг в конце сезона" selected>
            Снятие дуг в конце сезона
        </option>
        <option value="Поставить дуги в начале сезона">
            Поставить дуги в начале сезона
        </option>
        <option value="Частичный демонтаж">
            Частичный демонтаж
        </option>
        `
        break;
    case "Манипуляции над точкой":
        container.innerHTML = 
        `
        <option value="Добавить дугу" selected>
            Добавить дугу
        </option>
        <option value="Убрать дугу">
            Убрать дугу
        </option>
        <option value="Монтаж новой точки">
            Монтаж новой точки
        </option>
        <option value="Монтаж старой точки">
            Монтаж старой точки
        </option>
        <option value="Перенос точки">
            Перенос точки
        </option>
        `
        break;
    case "Разметка":
        container.innerHTML = 
        `
        <option value="Нанести разметку - Дорожная краска" selected>
            Нанести разметку - Дорожная краска
        </option>
        <option value="Нанести разметку - Термопластик">
            Нанести разметку - Термопластик
        </option>
        <option value="Нанести разметку - Баллончик">
            Нанести разметку - Баллончик
        </option>
        <option value="Демаркировка">
            Демаркировка
        </option>
        <option value="Частичное нанесение">
            Частичное нанесение
        </option>
        `
        break;
    case "Деактивация":
        container.innerHTML = 
        `
        <option value="Полная деактивация точки" selected>
            Полная деактивация точки
        </option>
        <option value="Снятие всех дуг">
            Снятие всех дуг
        </option>
        <option value="Демаркировка всех разметок">
            Демаркировка всех разметок
        </option>
        `
        break;
    case "Стандартные":
        container.innerHTML = 
        `
        <option value="Проинспектировать" selected>
            Проинспектировать
        </option>
        <option value="Произвести сервис">
            Произвести сервис
        </option>
        `
        break;
    case "Другое":
        container.innerHTML = 
        `
        <option value="Замена дуги на алюминиевую" selected>
            Замена дуги на алюминиевую
        </option>
        <option value="Сделать не свою дугу">
            Сделать не свою дугу
        </option>
        `
        break;
    }

}