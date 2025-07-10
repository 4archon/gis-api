function renderCards(data) {
    let root = document.getElementById("root");
    root.innerHTML = "";
    data.forEach(element => {
        let card = document.createElement("div")
        card.className = "card mt-1";
        card.innerHTML = render_card(element)
        root.appendChild(card);
    });
}

let data;

async function getUsers() {
    let url = "/employees"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin"
    })
    res = await response.json();
    data = res.users;
    renderCards(data);
}

getUsers();

const dialog = new bootstrap.Modal(document.getElementById("dialog"), null);

function changeUser(event) {
    let id = event.target.getAttribute("data-id");
    element = data.find(element => element.id == id);
    document.getElementById("dialog-title").innerHTML = "Изменить данные пользователя id: " + id;
    body = document.getElementById("dialog-body");
    body.innerHTML = render_edit(element);
    save = document.getElementById("dialog-save");
    save.setAttribute("data-id", id);
    document.getElementById("dialog-save").onclick = changeUserData;
    document.getElementById("dialog-save").innerHTML = "Сохранить изменения";
    dialog.show();
}

function changeUserData(event) {
    let id = event.target.getAttribute("data-id");
    element = data.find(element => element.id == id);
    element.login = document.getElementById("inputLogin").value
    element.role = document.getElementById("inputRole").value;
    element.active = document.getElementById("inputActive").value == "true" ? true : false;
    element.name = document.getElementById("inputName").value;
    element.surname = document.getElementById("inputSurname").value;
    element.patronymic = document.getElementById("inputPatronymic").value;
    element.tgID = Number(document.getElementById("inputTgID").value);
    element.subgroup = document.getElementById("inputSubgroup").value;
    element.trust = document.getElementById("inputTrust").value == "true" ? true : false;
    console.log(element);

    element["password"] = document.getElementById("inputPassword").value;
    changeUserBackend(element).then(() => {
        renderCards(data);
        dialog.hide();
    })
    // backend change user data
    // backend change password
    
}

async function changeUserBackend(element) {
    let url = "/change_user"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(element)
    })
    res = await response;
}

function newUser() {
    body = document.getElementById("dialog-body");
    body.innerHTML = render_new_edit();
    document.getElementById("dialog-title").innerHTML = "Создать нового пользователя";
    document.getElementById("dialog-save").onclick = createNewUser;
    document.getElementById("dialog-save").innerHTML = "Создать пользователя";
    dialog.show();
}

function createNewUser() {
    let element = {
        login: document.getElementById("inputLogin").value,
        password: document.getElementById("inputPassword").value,
        role: document.getElementById("inputRole").value,
        active: document.getElementById("inputActive").value == "true" ? true : false,
        name: document.getElementById("inputName").value,
        surname: document.getElementById("inputSurname").value,
        patronymic: document.getElementById("inputPatronymic").value,
        tgID: Number(document.getElementById("inputTgID").value),
        subgroup: document.getElementById("inputSubgroup").value,
        trust: document.getElementById("inputTrust").value == "true" ? true : false
    }
    createNewUserBackend(element).then(() => {
        data.unshift(element);
        renderCards(data);
        dialog.hide();
    });
    // backend add new user data
    // backend new user password
}

async function createNewUserBackend(element) {
    let url = "/create_new_user"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(element)
    })
    res = await response.text();
    element["id"] = Number(res);
}

document.getElementById("new-user-button").onclick = newUser;