function activePage() {
    let path = window.location.pathname.substring(1, window.location.pathname.length);
    switch(path) {
        case "employees":
            document.getElementById(path).classList.add("active");
            return
        case "main":
            document.getElementById(path).classList.add("active");
            return
        case "analytics":
            document.getElementById(path).classList.add("active");
            return
        case "distribute_tasks":
            document.getElementById(path).classList.add("active");
            return
        case "tasks":
            document.getElementById(path).classList.add("active");
            return
    }
}

async function getLogin() {
    url = "/account/login";
    response = await fetch(url);
    login = await response.text();
    return login;
}

async function getRole() {
    url = "/account/role";
    response = await fetch(url);
    role = await response.text();
    return role;
}

async function fetchHeader() {
    try {
        role = await getRole()
        if (role == "admin") {
            response = await fetch("/static/header/header_admin.html");
        }
        else {
            response = await fetch("/static/header/header_worker.html");
        }
        let el = document.createElement("div");
        el.innerHTML = await response.text();
        login = await getLogin();
        document.body.prepend(el);
        document.getElementById("profile").innerHTML = login;
        document.getElementById("profile_sidebar").innerHTML = login;
        activePage()
        nav_but = document.getElementById("nav-bar-open");
        nav_but.onclick = show_nav;
    }
    catch(err) {
        console.log(err)
    }
}

function show_nav() {
    navCanvas = document.getElementById("offcanvasNav");
    off_canvas_nav = new bootstrap.Offcanvas(navCanvas);
    off_canvas_nav.show();
}


fetchHeader()

