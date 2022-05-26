table = document.getElementById("task-table");

window.onload = function() {
    console.log("HELLO")

    $.ajax({
    url: "/api/tasks",
    type: "GET",
    success: function(res) {
         for (i=0; i< res.data.length; i++) {
            var row = table.insertRow();
            var cell1 = row.insertCell();
            var cell2 = row.insertCell();

            cell1.innerHTML = res.data[i].TaskName;
            cell2.innerHTML = res.data[i].Status;
        }
    },
    error: function(err) {
        console.log("error: ")
        console.log(err)
    }
})
}