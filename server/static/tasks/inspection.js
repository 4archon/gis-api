document.getElementById("inputCheckup").onchange = (el) => {
    if (el.target.value == "Точка требует ремонта") {
        add_require();
    } else if (el.target.value == "Точка не требует ремонта") {
        add_not_require();
    } else {
        console.log("xyi")
    }
};

function add_require() {
    parent = document.getElementById("inputRepairType");
    target_html = `<option value="Временный демонтаж" selected>Временный демонтаж</option>
                    <option value="Демонтаж-монтаж">Демонтаж-монтаж</option>
                    <option value="Демонтаж-монтаж + Покраска">Демонтаж-монтаж + Покраска</option>
                    <option value="Забрать дуги срочно">Забрать дуги срочно</option>
                    <option value="Замена на алюминиевую дугу">Замена на алюминиевую дугу</option>
                    <option value="Монтаж старой точки">Монтаж старой точки</option>
                    <option value="Покраска">Покраска</option>`;
    parent.innerHTML = target_html;
}

function add_not_require() {
    parent = document.getElementById("inputRepairType");
    target_html = `<option value="Идет благоустройство" selected>Идет благоустройство</option>
                    <option value="Уточнение не требуется">Уточнение не требуется</option>`;
    parent.innerHTML = target_html;
}