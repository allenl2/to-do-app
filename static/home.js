const logoutBtn = document.getElementById("logoutBtn");

logoutBtn.addEventListener("click", handleLogout);

function handleLogout(event) {
    event.preventDefault();

    
    fetch("/api/logout", {
        method: 'POST',
    })
        .then(res => {
            console.log(res)
            if(res.ok) {
                window.location.assign("/")
            }
        })
        .catch((err) => {
            console.log(err)
        })
}
