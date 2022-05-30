const logoutBtn = document.getElementById("logoutBtn");
const newTaskForm = document.getElementById("newTaskForm");
const newTask = document.getElementById("newTask");
const deleteTaskBtns = document.getElementsByClassName("delete-btn")

logoutBtn.addEventListener("click", handleLogout);
newTaskForm.addEventListener("submit", addNewTask);

for (var i=0; i< deleteTaskBtns.length; i++) {
    deleteTaskBtns[i].addEventListener("click", deleteTask);
}

function handleLogout(event) {
    event.preventDefault();

    fetch("/api/logout", {
        method: 'POST',
    })
        .then(res => {
            console.log(res);
            if(res.ok) {
                window.location.assign("/");
            }
        })
        .catch((err) => {
            console.log(err);
        })
}

function addNewTask(event) {
    event.preventDefault;

    fetch("/api/tasks", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            taskname: newTask.value,
            status: "inprogress"
        })
    })
    .then(res => {
        console.log(res);
        if(res.ok) {
            alert("added new task");
        } else {
            alert("unable to add new task!");
        }
    })
    .catch((err) => {
        console.log(err);
    });
}

function deleteTask(event) {
    event.preventDefault;
    var uri = "/api/tasks/" + this.getAttribute("taskID");
    
    fetch(uri, {
        method: 'DELETE',
    })
    .then(res => {
        console.log(res);
        if(res.ok) {
            alert("task deleted");
        } else {
            alert("unable to delete task");
        }
    })
    .catch((err) => {
        console.log(err);
    });

}
