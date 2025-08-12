async function deleteTaskFromPoint(event) {
    let data = {
        id: Number(event.currentTarget.getAttribute("data-id"))
    }
    let parent = event.currentTarget.parentElement.parentElement.
    parentElement.parentElement.parentElement;
    let tasksContainer = parent.parentElement;

    let url = "/delete_task"
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
    if (res.ok) {
        parent.remove();
        getPoinst();
        if (tasksContainer.childElementCount == 0) {
            tasksContainer.innerHTML +=
            `
            <ul class="list-group list-group-flush">
                <li class="list-group-item">Задачи не выставлены</li>
            </ul>
            `
        }
    }
}