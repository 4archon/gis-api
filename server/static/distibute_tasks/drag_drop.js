all_workers = document.getElementById("list_all_employees");
for (child of all_workers.children) {
    child.addEventListener("dragstart", (event) => {
        data = {
            type_drop: "worker",
            id: event.target.getElementsByTagName("span")[0].innerHTML,
            html_in: event.target.outerHTML
        }
        json_str = JSON.stringify(data);
        event.dataTransfer.setData("application/json", json_str);
    })
}


workers_active = []

list_workers = document.getElementById("list-employees");
list_workers.addEventListener("dragover", (e) => e.preventDefault());
list_workers.addEventListener("drop", (e) => {
    data = JSON.parse(e.dataTransfer.getData("application/json"));
    if (data.type_drop == "worker") {
        num = Number(data.id);
        if (!workers_active.includes(num)) {
            workers_active.push(num);
            list_workers.innerHTML += data.html_in;
        }
    }
})

document.getElementById("clear-list-employees").onclick = clear_employees;

function clear_employees() {
    parent = document.getElementById("list-employees");
    parent.innerHTML = '';
    workers_active = [];
}


selected_points = document.getElementById("list-work");
selected_points.addEventListener("dragover", (e) => e.preventDefault());
selected_points.addEventListener("drop", (e) => {
    data = JSON.parse(e.dataTransfer.getData("application/json"));
    if (data.type_drop == "point") {
        num = Number(data.id);
        if (!work_list.includes(num)) {
            work_list.push(num);
            new_li = document.createElement('li');
            new_li.className = "list-group-item";
            new_li.draggable = true;
            tx = document.createTextNode("ID точки: " + num);
            new_li.appendChild(tx);
            selected_points.appendChild(new_li);
        }
    }
})

document.getElementById("assign_work").onclick = sendWork;

async function sendWork() {
    if (work_list.length != 0 && workers_active.length != 0) {
        deadline = document.getElementById("deadline");
        data = {
            tasks: work_list,
            workers: workers_active,
            deadline: deadline.value
        }
        json_str = JSON.stringify(data);
        url = "/assign_tasks"
        response = await fetch(url, {
            method: "POST",
            cache: "no-cache",
            credentials: "same-origin",
            headers: {
                "Content-Type": "application/json"
            },
            body: json_str
        })
        clear_work();
        clear_employees();
        deadline.value = "";
    }
}