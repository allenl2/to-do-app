console.log("welcome");

const form = document.getElementById("login-form");
const username = document.getElementById("username");
const password = document.getElementById("password");

$(document).ready(function() {
    form.addEventListener('submit', handleLogin);
})

async function handleLogin(event) {
    event.preventDefault();

    // $.ajax({
    //     url: "/api/login",
    //     type: "POST",
    //     data: {
    //         username: username.value,
    //         password: password.value,
    //     },
    //     success: function(res) {
    //         window.location.assign("/home")
    //     },
    //     error: function(err) {
    //         console.log("error: ")
    //         console.log(err)
    //     }
    // })
    
    await fetch("/api/login", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username.value,
            password: password.value,
        })
    })
        .then(res => {
            console.log(res)
            if(res.ok) {
                window.location.assign("/home")
            }
        })
        .catch((err) => {
            console.log(err)
        })
}