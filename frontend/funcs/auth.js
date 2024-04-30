import { BACKENDURL } from "./vars";
import { Flogin } from "./utils"

// TODO - INCLUDE GENDER, AGE AND OTHER DATA IN VALIDATION AND POST

const ValidateData = () => {
  const { value: username, value: email } = document.getElementById("uname")
    .querySelectorAll("input[type='text']"); // Assuming both fields are within a wrapper
  const password = document.getElementById("pass").value;
  const confirmPassword = document.getElementById("cpass").value;

  // Check for empty fields concisely
  if ([username, email].some(value => !value.trim())) {
    return "Username and email are required.";
  }

  // Validate length for both name and email consistently
  if (username.length > 20 || email.length > 30) {
    return "Username and email should each be up to 20 and 30 characters long, respectively.";
  }

  // Validate password equality and length
  if (password !== confirmPassword) {
    return "Passwords don't match.";
  }

  if (password.length < 6 || password.length > 20) {
    return "Password should be between 6 and 20 characters long.";
  }

  return "";
};

export const HandleSignup = async () => {
  const { value: username, value: email } = document.getElementById("uname")
    .querySelectorAll("input[type='text']"); // Assuming both fields are within a wrapper
  const password = document.getElementById("pass").value;

  const validationMessage = ValidateData();
  if (validationMessage) {
    alert(validationMessage);
    return;
  }

  const signUpData = {
    email,
    username,
    password,
  };

  try {
    const response = await fetch(BACKENDURL + "/signup", {
      method: "POST",
      body: JSON.stringify(signUpData),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      const errorText = await response.text();
      alert(errorText);
    } else {
      await Flogin()
    }
  } catch (error) {
    console.error("Signup error:", error);
    alert("An error occurred during signup. Please try again.");
  }
};


export const HandleLogin = async () => {
  const cred = document.getElementById("cred")
  const pass = document.getElementById("pass")

  if (!cred.value.trim() || !pass.value.trim()) {
    alert("please enter correct values")
    return
  }

  const res = await fetch(BACKENDURL + "/login", {
    method: "POST",
    body: JSON.stringify({
      cred,
      pass
    })
  })

  if (res.ok) {
    window.location.assign("/")
  } else {
    alert(res.text)
  }
}