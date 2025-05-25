async function getHistory(id) {
    url = "/history"
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

async function historyClick(event) {
    let id = event.target.getAttribute("data-id");
    data = await getHistory(id);
    
}

function render_history_header(data) {
    result = `
    <h1 class="modal-title fs-5">Профиль точки 
    <span class="badge text-bg-primary">${profile.id}</span>
    <button "
    type="button" class="btn btn-info btn-sm">История точки</button>
    </h1>
    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>`
    return result;
}