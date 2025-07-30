async function reloadData() {
    // console.log("reload");

    if (currentDataJson !== null) {
        let reloadResponse = await fetch("/main", {
            method: "POST",
            cache: "no-cache",
            credentials: "same-origin"
        })
        let reloadRes = await reloadResponse.json();
        let reloadDataJson = JSON.stringify(reloadRes);
        if (currentDataJson != reloadDataJson) {
            getPoinst();
            // console.log("da");
        } else {
            // console.log("net");
        }
    }

    setTimeout(reloadData, 5000);
}