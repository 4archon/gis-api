let taskBar = new bootstrap.Offcanvas(document.getElementById("tasks-bar"));
let users;
let selectedUsers = [];
let selectedPoints = [];

document.getElementById("task-bar-button").onclick = showTaskBar;

function showTaskBar() {
    render_all_users();
    taskBar.toggle();
}

async function getUsers() {
    let url = "/employees"
    response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin"
    })
    users = (await response.json()).users;
}

getUsers();

function render_all_users() {
    let container = document.getElementById("employees-select-body");
    container.innerHTML = users.reduce((acc, el) => {
        return acc +
        `
        <li class="list-group-item" draggable="true" data-id="${el.id}">
            <span class="badge text-bg-primary">${el.id}</span>
            ${el.login === null ? "Логин не указан" : el.login}
        </li>
        `
    }, "")
    for (child of container.children) {
        child.addEventListener("dragstart", userDragStart);
    }
}

function render_group_users(group) {
    let container = document.getElementById("employees-select-body");
    container.innerHTML = users.reduce((acc, el) => {
        if (el.subgroup == group) {
            return acc +
            `
            <li class="list-group-item" draggable="true" data-id="${el.id}">
                <span class="badge text-bg-primary">${el.id}</span>
                ${el.login === null ? "Логин не указан" : el.login}
            </li>
            `
        } else {
            return acc
        }
    }, "")
    for (child of container.children) {
        child.addEventListener("dragstart", userDragStart);
    }
}

function changeUsersMenu(event) {
    let clicked = event.currentTarget;
    document.getElementById("employees-select-all").classList.remove("active");
    document.getElementById("employees-select-service").classList.remove("active");
    document.getElementById("employees-select-inspection").classList.remove("active");
    clicked.classList.add("active");
    if (clicked.id == "employees-select-all") {
        render_all_users();
    } else if (clicked.id == "employees-select-service") {
        render_group_users("service");
    } else if (clicked.id == "employees-select-inspection") {
        render_group_users("inspection");
    }
}

function userDragStart(event) {
    let data = {
        type: "worker",
        id: event.currentTarget.getAttribute("data-id")
    }
    let jsonStr = JSON.stringify(data);
    event.dataTransfer.setData("application/json", jsonStr);
}

let containerSelectedEmployees = document.getElementById("selected-employees");
containerSelectedEmployees.addEventListener("dragover", (e) => e.preventDefault());
containerSelectedEmployees.addEventListener("drop", (event) => {
    let data = JSON.parse(event.dataTransfer.getData("application/json"));
    if (data.type == "worker") {
        let id = Number(data.id);
        if (!selectedUsers.includes(id)) {
            selectedUsers.push(id);
            render_selected_users();
        }
    }
})

function render_selected_users() {
    let container = containerSelectedEmployees;
    container.innerHTML = ``;
    selectedUsers.forEach((el) => {
        let user = users.find((element) => element.id == el);
        container.innerHTML += 
        `
        <li class="list-group-item" draggable="true" data-id="${user.id}">
            <span class="badge text-bg-primary">${user.id}</span>
            ${user.login === null ? "Логин не указан" : user.login}
        </li>
        `
    });
    for (child of container.children) {
        child.addEventListener("dragstart", selectedUserDragStart);
        child.addEventListener("dragend", selectedUserDragEnd);
    }
}

function selectedUserDragStart(event) {
    let data = {
        type: "selectedWorker",
        id: event.currentTarget.getAttribute("data-id")
    }
    let jsonStr = JSON.stringify(data);
    event.dataTransfer.setData("application/json", jsonStr);
}

function selectedUserDragEnd(event) {
    if (event.dataTransfer.dropEffect == "none") {
        let id = Number(event.currentTarget.getAttribute("data-id"));
        selectedUsers = selectedUsers.filter((el) => el != id);
        render_selected_users();
    }
}

document.getElementById("clear-selected-users").onclick = clearSelectedUsers;

function clearSelectedUsers() {
    selectedUsers = [];
    render_selected_users();
}

function render_selected_points() {
    let container = document.getElementById("selected-points");
    container.innerHTML = selectedPoints.reduce((acc, el) => {
        return acc +
        `
        <li class="list-group-item" draggable="true" data-id="${el.id}">
            <span class="badge text-bg-primary">${el.id}</span>
            <span class="badge text-bg-danger">${el.externalID}</span>
            ${el.address}
        </li>
        `
    }, "");
    for (child of container.children) {
        child.addEventListener("dragstart", selectedPointsDragStart);
        child.addEventListener("dragend", selectedPointsDragEnd);
    }
}

function selectedPointsDragStart(event) {
    let data = {
        type: "selectedPoint",
        id: event.currentTarget.getAttribute("data-id")
    }
    let jsonStr = JSON.stringify(data);
    event.dataTransfer.setData("application/json", jsonStr);
}

function selectedPointsDragEnd(event) {
    if (event.dataTransfer.dropEffect == "none") {
        let id = Number(event.currentTarget.getAttribute("data-id"));
        selectedPoints = selectedPoints.filter((el) => el.id != id);
        render_selected_points();
    }
}

let containerSelectedPoints = document.getElementById("selected-points");
containerSelectedPoints.addEventListener("dragover", (e) => e.preventDefault());
containerSelectedPoints.addEventListener("drop", (event) => {});

document.getElementById("clear-selected-points").onclick = clearSelectedPoints;

function clearSelectedPoints() {
    selectedPoints = [];
    polygons.forEach((el) => el.destroy());
    render_selected_points();
}

document.getElementById("appoint").onclick = appoint;

async function appoint() {
    if (selectedUsers.length == 0 || selectedPoints.length == 0) {return}
    console.log("da");
    let data = {
        users: selectedUsers,
        points: selectedPoints.map((el) => el.id)
    };

    console.log(data);
    let url = "/appoint"
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
}