const logoutBtn = document.getElementById("logoutBtn");
const username = document.getElementById("username");
// const oldPass = document.getElementById("oldPassword");
const newPass = document.getElementById("newPassword");
const accountForm = document.getElementById("account-form");
const passForm = document.getElementById("password-form");

logoutBtn.addEventListener('click', handleLogout);

//disable enter key
$(document).keypress(
    function(event){
      if (event.which == '13') {
        event.preventDefault();
      }
  });

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
function handleAccountForm(event, userID) {
    event.preventDefault();

    if (accountForm.checkValidity() === false) {
        event.stopPropagation();
    }
    else {
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
   
    accountForm.classList.add('was-validated');
}

//update password when submitted
function handlePasswordForm(event, userID) {
    event.preventDefault();

    if (passForm.checkValidity() === false) {
        event.stopPropagation();
    }
    else {
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
                passForm.classList.remove('was-validated');
            } else {
                alert("unable to update!");
            }
        })
        .catch((err) => {
            console.log(err);
        });
    }
    
    passForm.classList.add('was-validated');
}