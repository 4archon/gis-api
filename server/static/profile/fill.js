let data;

async function getProfile() {
    let url = "/profile"
    let response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin"
    })
    let res = await response.json();
    data = res;
    renderProfile(data);
}

function renderProfile(data) {
    document.getElementById("inputLogin").value = data.login;
    document.getElementById("inputRole").value = data.role === null ? "": 
    data.role == "admin" ? "Администратор": "Работник";
    document.getElementById("inputActive").value = data.active ? "Активирован": "Деактивирован";
    document.getElementById("inputName").value = data.name;
    document.getElementById("inputSurname").value = data.surname;
    document.getElementById("inputPatronymic").value = data.patronymic;
    document.getElementById("inputTgID").value = data.tgID;
}

getProfile();

document.getElementById("cancel").onclick = () => {window.history.back();}

function changeProfile() {
    let element = {
        name: document.getElementById("inputName").value,
        surname: document.getElementById("inputSurname").value,
        patronymic: document.getElementById("inputPatronymic").value,
        tgID: Number(document.getElementById("inputTgID").value),
        password: document.getElementById("inputPassword").value
    };
    changeProfileBackend(element).then(() => {
        window.history.back();
    })
    // backend change user data
    // backend change user password
}

async function changeProfileBackend(element) {
    let url = "/change_user_profile"
    let response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(element)
    })
    let res = await response;
}

document.getElementById("save").onclick = changeProfile;