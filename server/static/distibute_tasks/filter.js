let filterBar = new bootstrap.Offcanvas(document.getElementById("point-filter"))
let filterOptions = null;

document.getElementById("filter-button").onclick = showFilter;

function showFilter() {
    render_users_filter();
    filterBar.toggle();
}

function render_users_filter() {
    let container = document.getElementById("filter-workers");
    container.innerHTML = users.reduce((acc, el) => {
        return acc +
        `
        <li class="list-group-item">
            <input id="filter-worker${el.id}" data-id="${el.id}" type="checkbox"
            ${filterOptions.appointUsersID.includes(el.id) ? "checked" : ""}
            class="form-check-label" onchange="fillPoints()" />
            <label class="mx-1" for="filter-worker${el.id}">
                <span class="badge text-bg-primary">${el.id}</span>
                <span class="badge text-bg-dark">
                    ${el.subgroup === null ? "н" : el.subgroup === "service" ? "с" : "и"}
                </span>
                ${el.login === null ? "Логин не указан" : el.login}
            </label>
        </li>
        `
    }, "");
}

function filterPoints() {
    getFilterOptions();
    console.log(filterOptions);
    let result = structuredClone(data);

    result = filteringRepairsForPeriod(result);
    result = filteringDaysWithout(result);
    result = filteringPointsInfo(result);
    result = filteringAppoint(result);
    result = filteringUserAppoint(result);
    result = filteringMarkings(result);
    result = filteringWorks(result);
    result = filteringStatuses(result);
    result = filteringTaskCustomers(result);
    result = filteringOnlyDeadline(result);
    result = filteringTasks(result);
    
    console.log(result);
    return result;
}

function filteringRepairsForPeriod(previous) {
    if (filterOptions.repairCount === null) {
        return previous;
    }

    if (!isNaN(filterOptions.repairPeriodStart) && !isNaN(filterOptions.repairPeriodEnd)) {
        previous = previous.filter((el) => {
            let counter = 0;
            if (el.repairs !== null) {
                el.repairs.forEach((element) => {
                    let repairDate = new Date(element);
                    if (repairDate >= filterOptions.repairPeriodStart &&
                        repairDate <= filterOptions.repairPeriodEnd) {
                        counter++;
                    }
                })
            }
            if (counter >= filterOptions.repairCount) {
                return true;
            }
        });
    } else if (!isNaN(filterOptions.repairPeriodStart)) {
        previous = previous.filter((el) => {
            let counter = 0;
            if (el.repairs !== null) {
                el.repairs.forEach((element) => {
                    let repairDate = new Date(element);
                    if (repairDate >= filterOptions.repairPeriodStart) {
                        counter++;
                    }
                })
            }
            if (counter >= filterOptions.repairCount) {
                return true;
            }
        });
    } else if (!isNaN(filterOptions.repairPeriodEnd)) {
        previous = previous.filter((el) => {
            let counter = 0;
            if (el.repairs !== null) {
                el.repairs.forEach((element) => {
                    let repairDate = new Date(element);
                    if (repairDate <= filterOptions.repairPeriodEnd) {
                        counter++;
                    }
                })
            }
            if (counter >= filterOptions.repairCount) {
                return true;
            }
        });
    } else {
        previous = previous.filter((el) => {
            let counter = 0;
            if (el.repairs !== null) {
                el.repairs.forEach((element) => {
                    counter++;
                })
            }
            if (counter >= filterOptions.repairCount) {
                return true;
            }
        });
    }

    return previous;
}

function filteringDaysWithout(previous) {
    if (filterOptions.daysWithoutCheck > 0) {
        previous = previous.filter((el) => {
            if (el.maxCheck ===  null) {
                return true;
            }
            let max = new Date(el.maxCheck);
            let now = Date.now();
            let days = Math.floor((now - max.getTime()) / (1000 * 60 * 60 * 24));
            if (days >= filterOptions.daysWithoutCheck) {
                return true;
            }
        });
    }

    if (filterOptions.daysWithoutRepair > 0) {
        previous = previous.filter((el) => {
            if (el.maxRepair ===  null) {
                return true;
            }
            let max = new Date(el.maxRepair);
            let now = Date.now();
            let days = Math.floor((now - max.getTime()) / (1000 * 60 * 60 * 24));
            if (days >= filterOptions.daysWithoutRepair) {
                return true;
            }
        });
    }

    if (filterOptions.daysWithoutExecutionTasks > 0) {
        previous = previous.filter((el) => {
            if (el.tasks ===  null) {
                return false;
            }
            let min = new Date(Math.min.apply(null, el.tasks.map((element) => new Date(element.entryDate))));
            let now = Date.now();
            let days = Math.floor((now - min.getTime()) / (1000 * 60 * 60 * 24));
            if (days >= filterOptions.daysWithoutExecutionTasks) {
                return true;
            }
        });
    }

    return previous;
}

function filteringPointsInfo(previous) {
    if (!filterOptions.deactivatedPoints) {
        previous = previous.filter((el) => el.active);
    }

    if (filterOptions.deactivatedPointsOnly) {
        previous = previous.filter((el) => !el.active);
    }

    if (filterOptions.numArcFrom !== null && filterOptions.numArcFrom > -1) {
        previous = previous.filter((el) => Number(el.numberArc) >= filterOptions.numArcFrom);
    }

    if (filterOptions.numArcTo !== null && filterOptions.numArcTo > -1) {
        previous = previous.filter((el) => Number(el.numberArc) <= filterOptions.numArcTo);
    }

    if (filterOptions.arcTypes.length != 0) {
        let noArcType = filterOptions.arcTypes.includes("Не указано");
        previous = previous.filter((el) => {
            if (noArcType && (el.arcType === null || el.arcType == "")) {
                return true;
            }
            return filterOptions.arcTypes.includes(el.arcType);
        });
    }

    if (filterOptions.carpets.length != 0) {
        let noCarpet = filterOptions.carpets.includes("Не указано");
        previous = previous.filter((el) => {
            if (noCarpet && (el.carpet === null || el.carpet == "")) {
                return true;
            }
            return filterOptions.carpets.includes(el.carpet);
        });
    }

    if (filterOptions.owners.length != 0) {
        previous = previous.filter((el) => filterOptions.owners.includes(el.owner));
    }

    if (filterOptions.operators.length != 0) {
        previous = previous.filter((el) => filterOptions.operators.includes(el.operator));
    }

    return previous;
}

function filteringAppoint(previous) {
    let noAppoint = filterOptions.appointTo.includes("Без назначений");
    let targetUserGroups = filterOptions.appointTo.map((el) => {
        if (el == "Инспекция") {
            return "inspection";
        } else if (el == "Сервис") {
            return "service";
        } else {
            return el;
        }
    });
    if (filterOptions.appointTo.length > 0 && filterOptions.appointOperationAnd) {
        if (noAppoint) {
            if (filterOptions.appointTo.length > 1) {
                return []
            } else {
                previous = previous.filter((el) => {
                    if (el.appoint === null) {
                        return true;
                    }
                })
            }
        } else {
            let userGroupsSet = new Set(targetUserGroups);
            previous = previous.filter((el) => {
                let pointSet = new Set(el.appoint === null ? null : 
                    el.appoint.map((element) => element.subgroup));
                if (userGroupsSet.isSubsetOf(pointSet)) {
                    return true;
                }
            });
        }
    } else if (filterOptions.appointTo.length > 0) {
        previous = previous.filter((el) => {
            if (el.appoint === null) {
                if (noAppoint) {
                    return true;
                } else {
                    return false;
                }
            }
            return el.appoint.some((element) => targetUserGroups.includes(element.subgroup));
        });
    }

    return previous;
}

function filteringUserAppoint(previous) {
    if (filterOptions.appointUsersID.length == 0) {
        return previous;
    }
    if (filterOptions.appointUsersOperationNot) {
        previous = previous.filter((el) => {
            if (el.appoint === null) {
                return true;
            } else {
                if (el.appoint.some((element) =>
                    filterOptions.appointUsersID.includes(element.id))) {
                    return false;
                } else {
                    return true;
                }
            }
        });
    } else {
        previous = previous.filter((el) => {
            if (el.appoint !== null) {
                if (el.appoint.some((element) => 
                    filterOptions.appointUsersID.includes(element.id))) {
                    return true;
                }
            }
        });
    }

    return previous;
}

function filteringMarkings(previous) {
    if (filterOptions.markingOnly) {
        previous = previous.filter((el) => el.marks !== null);
    }

    if (filterOptions.activeMarkings) {
        previous = previous.filter((el) => el.marks !== null).filter((el) => {
            return el.marks.some((element) => element.active);
        });
    }

    if (filterOptions.markingTypes.length != 0) {
        previous = previous.filter((el) => el.marks !== null).filter((el) => {
            return el.marks.some((element) => filterOptions.markingTypes.includes(element.type));
        });
    }

    return previous;
}

function filteringWorks(previous) {
    if (filterOptions.works.length == 0) {
        return previous;
    }
    let targetWorksSet = new Set(filterOptions.works);
    if (filterOptions.worksOperationAnd) {
        return previous.filter((el) => {
            let pointWorksSet = new Set(filteringWorksHelper(el.works));
            if (targetWorksSet.isSubsetOf(pointWorksSet)) {
                return true;
            }
        });
    } else {
        return previous.filter((el) => {
            let pointWorksSet = new Set(filteringWorksHelper(el.works));
            if (targetWorksSet.intersection(pointWorksSet).size != 0) {
                return true;
            }
        });
    }
}

function filteringWorksHelper(works) {
    let targetWorks = [];
    if (works !== null) {
        works.forEach((element) => {
            if (!targetWorks.includes(element.work)) {
                if (element.work != "Демонтаж") {
                    targetWorks.push(element.work);
                }
            }
        });
    }

    if (targetWorks.length > 1) {
        targetWorks = targetWorks.filter((el) => el != "Работа не требуется");
    }

    if (targetWorks.length == 0) {
        targetWorks.push("Неизвестно")
    }

    return targetWorks;
}

function filteringStatuses(previous) {
    if (filterOptions.statuses.length == 0) {
        return previous;
    }
    const noStatus = filterOptions.statuses.includes("Статус неизвестен");

    let res = previous.filter((el) => {
        if (filterOptions.statuses.includes(el.status)) {
            return true;
        }
    });

    if (noStatus) {
        res = res.concat(previous.filter((el) => el.status === null));
    }

    return res;
}

function filteringTaskCustomers(previous) {
    if (filterOptions.customers.length == 0) {
        return previous;
    }
    return previous.filter((el) => {
        if (el.tasks === null) {
            return false
        }
        return el.tasks.some((element) => {
            return filterOptions.customers.includes(element.customer);
        });
    });
}

function filteringOnlyDeadline(previous) {
    if (filterOptions.deadlineOnly) {
        return previous.filter((el) => {
            if (el.tasks === null) {
                return false;
            }
            return el.tasks.some((element) => element.deadline !== null);
        })
    } else {
        return previous;
    }
}

function filteringTasks(previous) {
    if (filterOptions.tasks.length == 0) {
        return previous;
    }
    
    let res = [];
    const noTasks = filterOptions.tasks.includes("Нет задач");
    if (filterOptions.tasksOperationAnd && noTasks) {
        if (filterOptions.tasks.length == 1) {
            res = previous.filter((el) => el.tasks === null);
        }
    } else if (filterOptions.tasksOperationAnd) {
        let filterTasksSet = new Set(filterOptions.tasks);
        res = previous.filter((el) => {
            if (el.tasks === null) {
                return false;
            }
            let pointSet = new Set(el.tasks.map((element) => element.type));
            if (filterTasksSet.isSubsetOf(pointSet)) {
                return true;
            } else {
                false;
            }
        });
    } else {
        let filterTasksSet = new Set(filterOptions.tasks);
        res = previous.filter((el) => {
            if (el.tasks === null) {
                return false;
            }
            let pointSet = new Set(el.tasks.map((element) => element.type));
            if (filterTasksSet.intersection(pointSet).size != 0) {
                return true;
            } else {
                false;
            }
        });
        if (noTasks) {
            res = res.concat(previous.filter((el) => el.tasks === null));
        }
    }

    return res;
}

function getFilterOptions() {
    filterOptions = {
        repairPeriodStart: null,
        repairPeriodEnd: null,
        repairCount: null,
        daysWithoutCheck: null,
        daysWithoutRepair: null,
        daysWithoutExecutionTasks: null,
        deactivatedPoints: null,
        deactivatedPointsOnly: null,
        numArcFrom: null,
        numArcTo: null,
        arcTypes: [],
        carpets: [],
        owners: [],
        operators: [],
        appointOperationAnd: null,
        appointTo: [],
        appointUsersOperationNot: null,
        appointUsersID: [],
        markingOnly: null,
        activeMarkings: null,
        markingTypes: [],
        worksOperationAnd: null,
        works: [],
        statuses: [],
        deadlineOnly: null,
        customers: [],
        tasksOperationAnd: null,
        tasks: []
    }

    filterOptions.repairPeriodStart = new Date(document.getElementById("repair-period-start").value);

    filterOptions.repairPeriodEnd = new Date(document.getElementById("repair-period-end").value);

    filterOptions.repairCount = document.getElementById("repair-count").value == "" ? null :
    Number(document.getElementById("repair-count").value);

    filterOptions.daysWithoutCheck = Number(document.getElementById("days-without-checks").value);

    filterOptions.daysWithoutRepair = Number(document.getElementById("days-without-repair").value);

    filterOptions.daysWithoutExecutionTasks = Number(document.getElementById("days-without-execution-tasks").value);

    filterOptions.deactivatedPoints = document.getElementById("deactivated-points").checked;
    filterOptions.deactivatedPointsOnly = document.getElementById("deactivated-points-only").checked;

    filterOptions.numArcFrom = document.getElementById("num-arc-from").value == "" ? null :
    Number(document.getElementById("num-arc-from").value);
    filterOptions.numArcTo = document.getElementById("num-arc-to").value == "" ? null :
    Number(document.getElementById("num-arc-to").value);

    for (let i = 1; i <= 5; i++) {
        if (document.getElementById(`at${i}`).checked) {
            filterOptions.arcTypes.push(document.getElementById(`l-at${i}`).innerHTML)
        }
    }

    for (let i = 1; i <= 3; i++) {
        if (document.getElementById(`cr${i}`).checked) {
            filterOptions.carpets.push(document.getElementById(`l-cr${i}`).innerHTML)
        }
    }

    for (let i = 1; i <= 3; i++) {
        if (document.getElementById(`ow${i}`).checked) {
            filterOptions.owners.push(document.getElementById(`l-ow${i}`).innerHTML)
        }
    }

    for (let i = 1; i <= 1; i++) {
        if (document.getElementById(`op${i}`).checked) {
            filterOptions.operators.push(document.getElementById(`l-op${i}`).innerHTML)
        }
    }

    filterOptions.appointOperationAnd = document.getElementById("ap-boolean").checked;

    for (let i = 1; i <= 3; i++) {
        if (document.getElementById(`ap${i}`).checked) {
            filterOptions.appointTo.push(document.getElementById(`l-ap${i}`).innerHTML)
        }
    }

    filterOptions.appointUsersOperationNot = document.getElementById("filter-workers-boolean").checked;

    for (const el of document.getElementById("filter-workers").getElementsByTagName("input")) {
        if (el.checked) {
            filterOptions.appointUsersID.push(Number(el.getAttribute("data-id")));
        }
    }

    filterOptions.markingOnly = document.getElementById("marking-only").checked;
    filterOptions.activeMarkings = document.getElementById("active-markings").checked;

    for (let i = 1; i <= 3; i++) {
        if (document.getElementById(`ma${i}`).checked) {
            filterOptions.markingTypes.push(document.getElementById(`l-ma${i}`).innerHTML)
        }
    }

    filterOptions.worksOperationAnd = document.getElementById("wr-boolean").checked;

    for (let i = 1; i <= 6; i++) {
        if (document.getElementById(`wr${i}`).checked) {
            filterOptions.works.push(document.getElementById(`l-wr${i}`).innerHTML)
        }
    }

    for (let i = 1; i <= 9; i++) {
        if (document.getElementById(`st${i}`).checked) {
            filterOptions.statuses.push(document.getElementById(`l-st${i}`).innerHTML)
        }
    }

    filterOptions.deadlineOnly = document.getElementById("deadline-only").checked;

    for (let i = 1; i <= 5; i++) {
        if (document.getElementById(`cu${i}`).checked) {
            filterOptions.customers.push(document.getElementById(`l-cu${i}`).innerHTML)
        }
    }

    filterOptions.tasksOperationAnd = document.getElementById("ex-boolean").checked;

    for (let i = 1; i <= 24; i++) {
        if (document.getElementById(`ex${i}`).checked) {
            filterOptions.tasks.push(document.getElementById(`l-ex${i}`).innerHTML)
        }
    }
}