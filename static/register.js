const form = document.getElementById("register-form");
const username = document.getElementById("username");
const password = document.getElementById("password");

form.addEventListener('submit', handleRegister);

async function handleRegister(event) {
    event.preventDefault();

    if (form.checkValidity() === false) {
        event.stopPropagation();
    }
    //if inputs are valids, register user
    else {
        await fetch("/api/user", {
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
                    window.location.assign("/")
                    alert("Registered successfully. Login to continue")
                }
                else {
                    alert("Unable to register. Please try again.")
                }
            })
            .catch(err => {
                console.log(err)
                alert("Unable to register. Please try again.")
            })
    }

    form.classList.add('was-validated');
}