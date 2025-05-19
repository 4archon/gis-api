let data;

async function getProfile() {
    url = "/profile"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin"
    })
    res = await response.json();
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
    document.getElementById("inputName").value;
    document.getElementById("inputSurname").value;
    document.getElementById("inputPatronymic").value;
    document.getElementById("inputTgID").value;
    // backend change user data
    document.getElementById("inputPassword").value;
    // backend change user password
    window.history.back();
}

document.getElementById("save").onclick = changeProfile;