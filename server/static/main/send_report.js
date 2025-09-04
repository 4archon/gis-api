function extractDataWorks() {
    if (reportData.type == "service") {
        return parseServiceWorks();
    } else if (reportData.type == "inspection") {
        return parseInspectionWorks();
    }
    return null;
}

function extractStatusFromGroup(group) {
    switch (group) {
    case "group1":
        return "Временно сняты дуги";
    case "group2":
        return "Идет благоустройство";
    case "group3":
        return "Точка доступна";
    case "group4":
        return "Точка недоступна";
    case "group5":
        return "Точка доступна";
    case "group6":
        return null;
    case "group10":
        return null;
    }
}

function getDeclineReportJson() {
    let report = {
        appoint: appoint,
        pointID: Number(reportPointID),
        tasks: selectedTasks.reduce((acc, el) => {
            acc.push({
                id: ((i) => Number(i) < 0 ? 0 : Number(i))(el.id),
                type: el.type
            });
            return acc
        }, []),
        reason: reportData.reason,
        yourself: reportData.yourself,
        comment: reportData.comment,
        duplicate: ((el) => el.original === null ? null : el)({
            original: reportData.pointID,
            duplicate: Number(reportPointID)
        })
    }
    console.log(report);
    return JSON.stringify(report);
}

function getServiceReportJson() {
    let report = {
        appoint: appoint,
        pointID: Number(reportPointID),
        tasks: selectedTasks.reduce((acc, el) => {
            acc.push({
                id: ((i) => Number(i) < 0 ? 0 : Number(i))(el.id),
                type: el.type
            });
            return acc
        }, []),
        done: extractDataWorks(),
        required: reportData.left ? reportData.left : null,
        location: reportData.carry ? reportData.newLocation : null,
        carpet: reportData.carry ? reportData.newCarpet : null,
        numberArc: reportData.numberArc,
        status: extractStatusFromGroup(reportGroup),
        comment: reportData.comment
    }
    console.log(report);
    return JSON.stringify(report);
}

function getInspectionReportJson() {
    let report = {
        appoint: appoint,
        pointID: Number(reportPointID),
        tasks: selectedTasks.reduce((acc, el) => {
            acc.push({
                id: ((i) => Number(i) < 0 ? 0 : Number(i))(el.id),
                type: el.type
            });
            return acc
        }, []),
        required: extractDataWorks(),
        paint: reportData.yourself ? reportData.paintCount : null,
        comment: reportData.comment
    }
    console.log(report);
    return JSON.stringify(report);
}

function sendReport(event) {
    event.currentTarget.disabled = true;
    pointReport.hide();
    let report = null;
    if (reportData.type == "decline") {
        report = getDeclineReportJson();
    } else if (reportData.type == "service") {
        report = getServiceReportJson();
    } else if (reportData.type == "inspection") {
        report = getInspectionReportJson();
    }

    if (report !== null && reportData.type) {
        postReportBackend(report, reportData.type);
    }
}

async function postReportBackend(report, reportType) {
    console.log(mediaCounter);
    let formData = new FormData();
    formData.append("report", report);
    formData.append("reportType", reportType);
    formData.append("count", mediaCounter);
    for (let i = 0; i < mediaCounter; i++) {
        let element = document.getElementById(`file${i}`);
        
        let name = element.getAttribute("data-name");
        formData.append(`name${i}`, name);

        let mediaType = element.files[0].type.split("/")[0];
        if (mediaType == "image") mediaType = "jpeg";
        if (mediaType == "video") mediaType = "mov";
        formData.append(`type${i}`, mediaType);

        formData.append(`file${i}`, element.files[0])
    }

    let url = "/report"
    let response = await fetch(url, {
        method: "POST",
        cache: "no-cache",
        credentials: "same-origin",
        body: formData
    })
    let resMedia = await response;
    console.log(resMedia.ok);
    newNotification(resMedia.ok);
}

function newNotification(success) {
    let alertContainer = document.createElement("div");
    if (success) {
        alertContainer.className = "alert alert-success my-1";
        alertContainer.innerHTML = `<h6 class="m-0">Отчет успешно отправлен</h6>`;
    } else {
        alertContainer.className = "alert alert-danger my-1";
        alertContainer.innerHTML = `<h6 class="m-0">Произошла ошибка сервера</h6>`;
    }
    let container = document.getElementById("notification-bar");
    container.appendChild(alertContainer);
    getPoinst();

    setTimeout(() => {
        alertContainer.remove();
    }, 10000);
}