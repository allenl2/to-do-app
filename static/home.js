const logoutBtn = document.getElementById("logoutBtn");
const newTaskForm = document.getElementById("newTaskForm");
const newTask = document.getElementById("newTask");
const deleteTaskBtns = document.getElementsByClassName("delete-btn");
const statusBox = document.getElementsByClassName("form-check-input");

logoutBtn.addEventListener("click", handleLogout);
newTaskForm.addEventListener("submit", addNewTask);

for (var i=0; i< deleteTaskBtns.length; i++) {
    deleteTaskBtns[i].addEventListener("click", deleteTask);
    statusBox[i].addEventListener("click", updateStatus)
}

function handleLogout(event) {
    event.preventDefault();

    fetch("/api/logout", {
        method: 'POST',
    })
    .then(res => {
        if(res.ok) {
            window.location.assign("/");
        }
    })
    .catch((err) => {
        console.log(err);
    });
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
            isDone: false,
        })
    })
    .then(res => {
        if(res.ok) {
            console.log("added new task");
        } else {
            console.log("unable to add new task!");
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
        if(res.ok) {
            console.log("task deleted");
        } else {
            console.log("unable to delete task");
        }
    })
    .catch((err) => {
        console.log(err);
    });
}

function updateStatus(event) {
    event.preventDefault;
    var uri = "/api/tasks/" + this.getAttribute("taskID");
    
    fetch(uri, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            isDone: this.checked,
        })
    })
    .then(res => {
        if(res.ok) {
            console.log("status changed");
        } else {
            console.log("unable to change status");
        }
    })
    .catch((err) => {
        console.log(err);
    });
}
