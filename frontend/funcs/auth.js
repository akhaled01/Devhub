import { BACKENDURL } from "./vars";
import { EncodeBase64Image, SetSessionStorage } from "./utils";

/**
 * Validates signup form data
 * @returns error string if there is
 */
const ValidateData = () => {
  const fname = document.getElementById("fname").value;
  const lname = document.getElementById("lname").value;
  const age = document.getElementById("age").value;
  const gender = document.getElementById("gender").value;
  const username = document.getElementById("uname").value;
  const email = document.getElementById("email").value;
  const password = document.getElementById("pass").value;
  const confirmPassword = document.getElementById("cpass").value;

  if (
    !fname ||
    !lname ||
    !age ||
    !gender ||
    !username ||
    !password ||
    !confirmPassword ||
    !email
  ) {
    return "There is a missing required field";
  }

  if (age < 0 || age > 100) {
    return "please enter a sensible age";
  }

  if (gender.toUpperCase() !== "M" && gender.toUpperCase() !== "F") {
    return "gender can either be M or F";
  }

  if (username.length > 20 || email.length > 30) {
    return "Username and email should each be up to 20 and 30 characters long, respectively.";
  }

  if (password.length < 6) {
    return "Password should be atleast 6 charachters long";
  }

  if (password !== confirmPassword) {
    return "Passwords don't match.";
  }

  return "";
};

/**
 * Invokes a async fetch request to create a new user
 * and log him in
 */
export const HandleSignup = async () => {
  const validationMessage = ValidateData();

  if (validationMessage) {
    alert(validationMessage);
    return;
  }

  EncodeBase64Image(async (EncodedAvatar) => {
    const formData = {};
    const authGroups = document.querySelectorAll(".authgroup");

    authGroups.forEach((group) => {
      const inputs = group.querySelectorAll("input");
      inputs.forEach((input) => {
        if (input.name !== "avatar" && input.name !== "signUpPassConfirm") {
          formData[input.name] = input.value;
        }
      });
    });

    formData["image"] = EncodedAvatar;
    formData["age"] = parseInt(formData["age"]);

    console.log(formData);

    try {
      const response = await fetch(BACKENDURL + "/auth/signup", {
        method: "POST",
        body: JSON.stringify(formData),
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.status === 201) {
        switch (response.status) {
          case 409:
            alert("some of the credetials have already been used");
            break;
          default:
            alert("error on signup, check logs");
            break;
        }
      } else {
        window.location.assign("/login"); //TODO: If you can, remimplement FLogin
      }
    } catch (error) {
      console.error("Signup error:", error);
      alert("An error occurred during signup. Please try again.");
    }
  });
};

/**
 * Invokes async request that handles login
 *
 * Path of backend is /auth/login
 */
export const HandleLogin = async () => {
  const cred = document.getElementById("cred");
  const pass = document.getElementById("pass");

  if (!cred.value.trim() || !pass.value.trim()) {
    alert("please enter correct values");
    return;
  }

  const res = await fetch(BACKENDURL + "/auth/login", {
    method: "POST",
    body: JSON.stringify({
      credential: cred.value,
      password: pass.value,
      credential: cred.value,
      password: pass.value,
    }),
    credentials: "include",
  });

  if (res.ok) {
    const data = await res.json();
    // Store data in sessionStorage
    SetSessionStorage(data);
    window.location.assign("/");
  } else {
    switch (res.status) {
      case 401:
        alert("Incorrect Password");
        break;
      case 404:
        alert("User not found");
        break;
      default:
        alert("error logging in");
        break;
    }
  }
};
