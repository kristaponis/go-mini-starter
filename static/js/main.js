// toggleAlert() hides error message in the /signup or /login form.
// These errors come from the server side validation.
function toggleAlert() {
    var toggleEl = document.getElementById("alertId")
    toggleEl.style.display = "none"
}

// Client side /signup and /login form validations and error control.
// It validates Name, Email and Passwords inputs 
// and displays error messages if there are any errors.
const signupForm = document.getElementById("signup-form")
const loginForm = document.getElementById("login-form")
const userName = document.getElementById("name")
const userEmail = document.getElementById("email")
const userPassword = document.getElementById("password")

function validateName() {
    const small = userName.parentElement.querySelector("small")
    if (userName.value.trim() === "") {
        small.innerText = "Name cannot be blank"
        return true
    }
    if (userName.value.trim().length < 2) {
        small.innerText = "Name is less than 2 char"
        return true
    }
    if (userName.value.trim().length > 100) {
        small.innerText = "Name is more than 100 char"
        return true
    }
}

function validateEmail() {
    const small = userEmail.parentElement.querySelector("small")
    if (userEmail.value.trim() === "") {
        small.innerText = "Email cannot be blank"
        return true
    }
    if (userEmail.value.trim().length < 3 || userEmail.value.trim().length > 100) {
        small.innerText = "Invalid email"
        return true
    }
}

function validatePassword() {
    const small = userPassword.parentElement.querySelector("small")
    if (userPassword.value.trim() === "") {
        small.innerText = "Password cannot be blank"
        return true
    }
    if (userPassword.value.trim().length < 8) {
        small.innerText = "Password is less than 8 char"
        return true
    }
    if (userPassword.value.trim().length > 100) {
        small.innerText = "Password is more than 100 char"
        return true
    }
}

if (signupForm) {
    signupForm.addEventListener("submit", (e) => {
        document.getElementById("signup-name").innerText = ""
        document.getElementById("signup-email").innerText = ""
        document.getElementById("signup-password").innerText = ""
        if (validateName()) {
            e.preventDefault()
        }
        if (validateEmail()) {
            e.preventDefault()
        }
        if (validatePassword()) {
            e.preventDefault()
        }
    })
}

if (loginForm) {
    loginForm.addEventListener("submit", (e) => {
        document.getElementById("login-email").innerText = ""
        document.getElementById("login-password").innerText = ""
        if (validateEmail()) {
            e.preventDefault()
        }
        if (validatePassword()) {
            e.preventDefault()
        }
    })
}

