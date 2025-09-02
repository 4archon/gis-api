function render_card(element) {
    res = ` <div class="card-header">
                <div class="row">
                    <div class="col-8">
                        <h5>
                        Логин: ${element.login === null ? "не указан": element.login}
                        </h5>
                    </div>
                    <div class="col-sm-4 col-12">
                        <button type="button" class="btn btn-primary float-end"
                        data-id="${element.id}" onclick="changeUser(event)">
                        Редактировать
                        </button>
                    </div>
                </div>
            </div>
            <div class="card-body">
                <h5 class="card-title">ФИО:
                ${element.surname === null ? "": element.surname} 
                ${element.name === null ? "": element.name}
                ${element.patronymic === null ? "": element.patronymic}</h5>
                <ul class="list-group list-group-flush">
                    <li class="list-group-item">Роль: 
                    ${element.role === null ? "не указана": element.role == "admin" ? "Администратор": "Работник"}</li>
                    <li class="list-group-item">Статус:
                    ${element.active ? "Активный": "Деактивирован"}</li>
                    <li class="list-group-item">Telegram id:
                    ${element.tgID === null ? "не указан": element.tgID}</li>
                    <li class="list-group-item">Группа:
                    ${element.subgroup === null ? "не указана": 
                        element.subgroup == "service" ? "Сервис":"Инспекция"}</li>
                    <li class="list-group-item">Доверять сотруднику:
                    ${element.trust === null ? "не указано": element.trust ? "Да":"Нет"}</li>
                </ul>
            </div>`
    return res;
}

function render_edit(element) {
    res = ` <div class="card-body mb-4">
                <div class="row g-3">
                    <div class="col-md-6">
                        <label for="inputLogin" class="form-label">Логин</label>
                        <input id="inputLogin"
                        value="${element.login === null ? "": element.login}" 
                        class="form-control"  placeholder="Введите логин">
                    </div>
                    <div class="col-md-6">
                        <label for="inputPassword" class="form-label">Пароль</label>
                        <input id="inputPassword" class="form-control"
                        placeholder="Пароль останется прежним">
                    </div>
                    <div class="col-md-6">
                        <label for="inputRole" class="form-label">Роль</label>
                        <select id="inputRole" class="form-select">
                            <option value="${element.role == "admin" ? "admin": "worker"}" selected>
                            ${element.role == "admin" ? "Администратор": "Работник"}
                            </option>
                            <option value="${element.role == "admin" ? "worker": "admin"}">
                            ${element.role == "admin" ? "Работник": "Администратор"}
                            </option>
                        </select>
                    </div>
                    <div class="col-md-6">
                        <label for="inputActive" class="form-label">Статус</label>
                        <select id="inputActive" class="form-select">
                            <option value="${element.active ? "true": "false"}" selected>
                            ${element.active ? "Активирован": "Деактивирован"}
                            </option>
                            <option value="${element.active ? "false": "true"}">
                            ${element.active ? "Деактивирован": "Активирован"}
                            </option>
                        </select>
                    </div>
                    <div class="col-md-6">
                        <label for="inputSubgroup" class="form-label">Группа</label>
                        <select id="inputSubgroup" class="form-select">
                            <option 
                            value="${element.subgroup == "inspection" ? "inspection": "service"}"
                            selected>
                            ${element.subgroup == "inspection" ? "Инспекция": "Сервис"}
                            </option>
                            <option 
                            value="${element.subgroup == "inspection" ? "service": "inspection"}"
                            >
                            ${element.subgroup == "inspection" ? "Сервис": "Инспекция"}
                            </option>
                        </select>
                    </div>
                    <div class="col-md-6">
                        <label for="inputTrust" class="form-label">Доверять сотруднику?</label>
                        <select id="inputTrust" class="form-select">
                            <option value="${element.trust ? "true": "false"}" selected>
                            ${element.trust ? "Да": "Нет"}
                            </option>
                            <option value="${element.trust ? "false": "true"}">
                            ${element.trust ? "Нет": "Да"}
                            </option>
                        </select>
                    </div>
                    <label class="form-label mt-5"><h4>Личные данные</h4></label>
                    <div class="col-md-6">
                        <label for="inputName" class="form-label">Имя</label>
                        <input id="inputName" class="form-control" placeholder="Введите имя"
                        value="${element.name === null ? "": element.name}">
                    </div>
                    <div class="col-md-6">
                        <label for="inputSurname" class="form-label">Фамилия</label>
                        <input id="inputSurname" class="form-control" placeholder="Введите фамилию"
                        value="${element.surname === null ? "": element.surname}">
                    </div>
                    <div class="col-md-6">
                        <label for="inputPatronymic" class="form-label">Отчество</label>
                        <input id="inputPatronymic" class="form-control" placeholder="Введите отчество"
                        value="${element.patronymic === null ? "": element.patronymic}">
                    </div>
                    <div class="col-md-6">
                        <label for="inputTgID" class="form-label">Telegram ID</label>
                        <input id="inputTgID" class="form-control" placeholder="Введите телеграм ID"
                        value="${element.tgID === null ? "": element.tgID}">
                    </div>
                  </div>
            </div>`
    return res;
}

function render_new_edit() {
    res = ` <div class="card-body mb-4">
                <div class="row g-3">
                    <div class="col-md-6">
                        <label for="inputLogin" class="form-label">Логин</label>
                        <input id="inputLogin"
                        value="" 
                        class="form-control"  placeholder="Введите логин">
                    </div>
                    <div class="col-md-6">
                        <label for="inputPassword" class="form-label">Пароль</label>
                        <input id="inputPassword" class="form-control" 
                        placeholder="Пароль останется прежним">
                    </div>
                    <div class="col-md-6">
                        <label for="inputRole" class="form-label">Роль</label>
                        <select id="inputRole" class="form-select">
                            <option value="worker" selected>Работник</option>
                            <option value="admin">Администратор</option>
                        </select>
                    </div>
                    <div class="col-md-6">
                        <label for="inputActive" class="form-label">Статус</label>
                        <select id="inputActive" class="form-select">
                            <option value="true" selected>Активирован</option>
                            <option value="false">Деактивирован</option>
                        </select>
                    </div>
                    <div class="col-md-6">
                        <label for="inputSubgroup" class="form-label">Группа</label>
                        <select id="inputSubgroup" class="form-select">
                            <option value="service" selected>Сервис</option>
                            <option value="inspection">Инспекция</option>
                        </select>
                    </div>
                    <div class="col-md-6">
                        <label for="inputTrust" class="form-label">Доверять сотруднику?</label>
                        <select id="inputTrust" class="form-select">
                            <option value="true" selected>Да</option>
                            <option value="false">Нет</option>
                        </select>
                    </div>
                    <label class="form-label mt-5"><h4>Личные данные</h4></label>
                    <div class="col-md-6">
                        <label for="inputName" class="form-label">Имя</label>
                        <input id="inputName" class="form-control" placeholder="Введите имя"
                        value="">
                    </div>
                    <div class="col-md-6">
                        <label for="inputSurname" class="form-label">Фамилия</label>
                        <input id="inputSurname" class="form-control" placeholder="Введите фамилию"
                        value="">
                    </div>
                    <div class="col-md-6">
                        <label for="inputPatronymic" class="form-label">Отчество</label>
                        <input id="inputPatronymic" class="form-control" placeholder="Введите отчество"
                        value="">
                    </div>
                    <div class="col-md-6">
                        <label for="inputTgID" class="form-label">Telegram ID</label>
                        <input id="inputTgID" class="form-control" placeholder="Введите телеграм ID"
                        value="">
                    </div>
                  </div>
            </div>`
    return res
}