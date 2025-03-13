document.getElementById("inputServiceType1").onchange = (el) => {
    parent = document.getElementById("inputSubtype1");
    if (el.target.value == "Ремонт") {
        add_repaire(parent);
    } else if (el.target.value == "Демонтаж") {
        add_dismantling(parent);
    } else if (el.target.value == "Монтаж новой точки") {
        add_new_point(parent);
    } else {
        console.log("xyi")
    }
};

counter = 1;

function add_repaire(parent) {
    target_html = `<option value="Покраска" selected>Покраска</option>
                    <option value="Перенос">Перенос</option>
                    <option value="Демонтаж-монтаж">Демонтаж-монтаж</option>
                    <option value="Монтаж старой точки">Монтаж старой точки</option>
                    <option value="Покраска комплекс">Покраска комплекс</option>`;
    parent.innerHTML = target_html;
}

function add_dismantling(parent) {
    target_html = `<option value="Демонтаж временный" selected>Демонтаж временный</option>
                    <option value="Демонтаж НЕ временный">Демонтаж НЕ временный</option>`;
    parent.innerHTML = target_html;
}

function add_new_point(parent) {
    target_html = `<option value="Монтаж новой точки" selected>Монтаж новой точки</option>`;
    parent.innerHTML = target_html;
}

function add_minus_button() {
    but = document.createElement("button");
    but.id = "sub-button";
    but.type = "button";
    but.className = "btn";
    but.onclick = (el) => {
        delete_element(counter);
        sub_counter();
        if (counter == 1) {
            el.target.remove();
        }
    };
    but.innerHTML  = `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="bi bi-dash-square" viewBox="0 0 16 16">
                            <path d="M14 1a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1zM2 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2z"/>
                            <path d="M4 8a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7A.5.5 0 0 1 4 8"/>
                        </svg>`
    

    document.getElementById("add-button").after(but)
}

function delete_element(index) {
    document.getElementById(`serviceType${index}`).remove();
    document.getElementById(`subtype${index}`).remove();
    document.getElementById(`before${index}`).remove();
    document.getElementById(`front${index}`).remove();
    document.getElementById(`left${index}`).remove();
    document.getElementById(`right${index}`).remove();
    document.getElementById(`video_div${index}`).remove();
    document.getElementById(`extra${index}`).remove();
    document.getElementById(`comment${index}`).remove();
}

function add_service_type(parent, entry) {
    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "serviceType" + counter; 
    
    label = document.createElement("label");
    label.htmlFor = "inputServiceType" + counter;
    label.className = "form-label";
    label.innerHTML = "Тип услуги";
    div.appendChild(label);

    select = document.createElement("select");
    select.setAttribute("name", "inputServiceType" + counter);
    select.className = "form-select";
    select.setAttribute("required", "");
    select.id = "inputServiceType" + counter;
    select.innerHTML = `<option value="Ремонт" selected>Ремонт</option>
                        <option value="Демонтаж">Демонтаж</option>
                        <option value="Монтаж новой точки">Монтаж новой точки</option>`;
    div.appendChild(select);
    
    parent.insertBefore(div, entry);
}

function add_subtype(parent, entry) {
    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "subtype" + counter; 
    
    label = document.createElement("label");
    label.htmlFor = "inputSubtype" + counter;
    label.className = "form-label";
    label.innerHTML = "Тип ремонта";
    div.appendChild(label);

    select = document.createElement("select");
    select.setAttribute("name", "inputSubtype" + counter);
    select.className = "form-select";
    select.setAttribute("required", "");
    select.id = "inputSubtype" + counter;
    select.innerHTML = `<option value="Покраска" selected>Покраска</option>
                        <option value="Перенос">Перенос</option>
                        <option value="Демонтаж-монтаж">Демонтаж-монтаж</option>
                        <option value="Монтаж старой точки">Монтаж старой точки</option>
                        <option value="Покраска комплекс">Покраска комплекс</option>`;
    div.appendChild(select);
    
    document.getElementById("inputServiceType" + counter).onchange = (el) => {
        l = "inputServiceType"
        num = el.target.id.substring(l.length);
        parent = document.getElementById("inputSubtype" + num);
        if (el.target.value == "Ремонт") {
            add_repaire(parent);
        } else if (el.target.value == "Демонтаж") {
            add_dismantling(parent);
        } else if (el.target.value == "Монтаж новой точки") {
            add_new_point(parent);
        } else {
            console.log("xyi")
        }
    };

    parent.insertBefore(div, entry);
}

function add_file_nodes(parent, entry) {
    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "before" + counter;
    div.innerHTML = `<label for="photo_before${counter}" class="form-label">Фото до</label>
                    <input name="photo_before${counter}" class="form-control" type="file"
                    id="photo_before${counter}" required>`;
    parent.insertBefore(div, entry);

    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "front" + counter;
    div.innerHTML = `<label for="photo_front${counter}" class="form-label">Фото спереди</label>
                    <input name="photo_front${counter}" class="form-control" type="file"
                    id="photo_front${counter}" required>`;
    parent.insertBefore(div, entry);

    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "left" + counter;
    div.innerHTML = `<label for="photo_left${counter}" class="form-label">Фото слева</label>
                    <input name="photo_left${counter}" class="form-control" type="file"
                    id="photo_left${counter}" required>`;
    parent.insertBefore(div, entry);

    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "right" + counter;
    div.innerHTML = `<label for="photo_right${counter}" class="form-label">Фото справа</label>
                    <input name="photo_right${counter}" class="form-control" type="file"
                    id="photo_right${counter}" required>`;
    parent.insertBefore(div, entry);

    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "video_div" + counter;
    div.innerHTML = `<label for="video${counter}" class="form-label">Видео</label>
                    <input name="video${counter}" class="form-control" type="file"
                    id="video${counter}" required>`;
    parent.insertBefore(div, entry);

    div = document.createElement("div");
    div.className = "col-sm-6 col-12";
    div.id = "extra" + counter;
    div.innerHTML = `<label for="photo_extra${counter}" class="form-label">Дополнительное фото</label>
                    <input name="photo_extra${counter}" class="form-control" type="file"
                    id="photo_extra${counter}">`;
    parent.insertBefore(div, entry);

    div = document.createElement("div");
    div.className = "col-12";
    div.id = "comment" + counter;
    div.innerHTML = `<label for="inputComment${counter}" class="form-label">Комментарий</label>
                    <textarea name="inputComment${counter}" class="form-control"
                    id="inputComment${counter}" rows="3" required></textarea>`;
    parent.insertBefore(div, entry);
}

function add_nodes() {
    parent = document.getElementById("form");
    entry = document.getElementById("entry-point");
    add_service_type(parent, entry);
    add_subtype(parent, entry);
    add_file_nodes(parent, entry);
}

function add_counter() {
    counter++;
    document.getElementById("serviceCounter").value = counter;
}

function sub_counter() {
    counter--;
    document.getElementById("serviceCounter").value = counter;
}

document.getElementById("add-button").onclick = () => {
    add_counter();
    if (counter == 2) {
        add_minus_button();
    }
    add_nodes();
};