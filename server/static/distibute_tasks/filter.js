document.getElementById("active").onchange = filter;
document.getElementById("not-active").onchange = filter;
document.getElementById("repair").onchange = filter;
document.getElementById("not-repair").onchange = filter;
document.getElementById("assigned").onchange = filter;
document.getElementById("not-assigned").onchange = filter;
document.getElementById("longtime").onchange = filter;
document.getElementById("not-longtime").onchange = filter;

function active_filter(points, active, not_active) {
    let filtered = [];
    if ((active && not_active) || !(active || not_active)) {
        return points;
    } else {
        if (active) {
            points.forEach((el) => {
                if (el.Active) {
                    filtered.push(el);
                }
            })
        } else {
            points.forEach((el) => {
                if (!(el.Active)) {
                    filtered.push(el);
                }
            })
        }
    }
    return filtered;
}

function repair_filter(points, repair, not_repair) {
    let filtered = [];
    if ((repair && not_repair) || !(repair || not_repair)) {
        return points;
    } else {
        if (repair) {
            points.forEach((el) => {
                if (el.Repair) {
                    filtered.push(el);
                }
            })
        } else {
            points.forEach((el) => {
                if (!(el.Repair)) {
                    filtered.push(el);
                }
            })
        }
    }
    return filtered;
}

function assigned_filter(points, assigned, not_assigned) {
    let filtered = [];
    if ((assigned && not_assigned) || !(assigned || not_assigned)) {
        return points;
    } else {
        if (assigned) {
            points.forEach((el) => {
                if (el.Assigned) {
                    filtered.push(el);
                }
            })
        } else {
            points.forEach((el) => {
                if (!(el.Assigned)) {
                    filtered.push(el);
                }
            })
        }
    }
    return filtered;
}

function longtime_filter(points, longtime, not_longtime) {
    let filtered = [];
    if ((longtime && not_longtime) || !(longtime || not_longtime)) {
        return points;
    } else {
        if (longtime) {
            filtered = points.filter((el) => el.LongTime)
        } else {
            filtered = points.filter((el) => !el.LongTime)
        }
    }
    return filtered;
}

function filter() {
    let filtered = all_points.slice();

    let active = document.getElementById("active").checked;
    let not_active = document.getElementById("not-active").checked;
    let repair = document.getElementById("repair").checked;
    let not_repair = document.getElementById("not-repair").checked;
    let assigned = document.getElementById("assigned").checked;
    let not_assigned = document.getElementById("not-assigned").checked;
    let longtime = document.getElementById("longtime").checked;
    let not_longtime = document.getElementById("not-longtime").checked;

    if (active || not_active || repair || not_repair
        || assigned || not_assigned || longtime || not_longtime
    ) {
        filtered = active_filter(filtered, active, not_active);
        filtered = repair_filter(filtered, repair, not_repair);
        filtered = assigned_filter(filtered, assigned, not_assigned);
        filtered = longtime_filter(filtered, longtime, not_longtime);
    }

    clusterer.load(filtered);
    markers = filtered;
}

markers = all_points;