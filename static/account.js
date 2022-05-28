const logoutBtn = document.getElementById("logoutBtn");
const username = document.getElementById("username");
// const oldPass = document.getElementById("oldPassword");
const newPass = document.getElementById("newPassword");

logoutBtn.addEventListener('click', handleLogout);

//logout the current user
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
    });
}

//update account details when submitted
function handleAccountForm(userID) {
    var uri = "/api/user/" + userID

    fetch(uri, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username.value,
        })
    })
    .then(res => {
        console.log(res);
        if(res.ok) {
            alert("successful update!");
        } else {
            alert("unable to update!");
        }
    })
    .catch((err) => {
        console.log(err);
    });
}

//update password when submitted
function handlePasswordForm(userID) {
    var uri = "/api/user/" + userID;

    fetch(uri, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            password: newPass.value,
        })
    })
    .then(res => {
        console.log(res);
        if(res.ok) {
            alert("successful update!");
            newPass.value = "";
        } else {
            alert("unable to update!");
        }
    })
    .catch((err) => {
        console.log(err);
    });
}