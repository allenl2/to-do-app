const logoutBtn = document.getElementById("logoutBtn");
const newTaskForm = document.getElementById("newTaskForm");
const newTask = document.getElementById("newTask");
const statusBox = document.getElementsByClassName("form-check-input");

const editTaskBtns = document.getElementsByClassName("edit-btn");
const taskContent = document.getElementsByClassName("task-content");
const editModalText = document.getElementById("edit-modal-text");
const saveTaskBtn = document.getElementById("save-edit-btn");
const deleteTaskBtns = document.getElementsByClassName("delete-btn");

logoutBtn.addEventListener("click", handleLogout);
newTaskForm.addEventListener("submit", addNewTask);
saveTaskBtn.addEventListener("click", updateTask);

var currentTaskID;
var currenTaskContent;

for (var i=0; i< deleteTaskBtns.length; i++) {
    statusBox[i].addEventListener("click", updateStatus);
    editTaskBtns[i].addEventListener("click", renderEditModal);
    deleteTaskBtns[i].addEventListener("click", deleteTask);
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
            alert("task deleted");
            this.parentElement.parentElement.remove();
        } else {
            alert("unable to delete task");
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


function renderEditModal(event) {
    event.preventDefault;
    currentTaskID = this.getAttribute("taskID");
    currenTaskContent = this.parentElement.parentElement.childNodes[3];
    editModalText.value = currenTaskContent.innerHTML;
}


function updateTask() {
    var uri = "/api/tasks/" + currentTaskID;

    fetch(uri, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            taskname: editModalText.value,
        })
    })
    .then(res => {
        if(res.ok) {
            console.log("updated task");
            currenTaskContent.innerHTML = editModalText.value;

        } else {
            console.log("unable to update task!");
        }
    })
    .catch((err) => {
        console.log(err);
    });
}