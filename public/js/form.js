date = new Date("2018-02-13");
year = date.getFullYear();
month = date.getMonth();
day = date.getDay();
function getDate() {
    document.getElementById("date").innerHTML = day + "-" + month+ "-" + year;
}

