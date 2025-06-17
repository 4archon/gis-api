let filterBar = new bootstrap.Offcanvas(document.getElementById("point-filter"))

document.getElementById("filter-button").onclick = showFilter;

function showFilter() {
    filterBar.show();
}