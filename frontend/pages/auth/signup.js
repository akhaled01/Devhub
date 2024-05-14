import { HandleSignup } from "../../funcs/auth";

export const Signup = () => {
  document.getElementById("app").innerHTML = /* html */ `
    <div class="logo">
  <img src="../../assets/logo.svg" id="authlogo">
</div>
<div class="form">
  <form id="signupform" autocomplete="off">
    <img src="../../assets/profile.svg" alt="upload avatar" title="upload avatar" id="uploadAvatarBtn">
    <div class="authgroup">
      <input name="avatar" placeholder="Avatar" type="file" id="avatar" required />
      <input name="first_name" placeholder="First Name" type="text" id="fname" required />
      <input name="last_name" placeholder="Last Name" type="text" id="lname" required />
    </div>
    <div class="authgroup">
      <input name="username" placeholder="Username" type="text" id="uname" required />
      <input name="email" placeholder="Email" type="text" id="email" required />
    </div>
    <div class="authgroup" >
      <div>
      <input name="age" placeholder="Age" type="number" id="age" required />
      </div>
      <div style="color:white;">
      <label for="gender" >Gender:</label><br>
      </div>
      <div style="color:white;" required >
      <input id="gender" type="radio" name="gender" value="M"> Male
      <input id="gender" type="radio" name="gender" value="F"> Female<br>
      </div>
    </div>
    <div class="authgroup">
      <input name="password" placeholder="Password" type="password" id="pass" required />
      <input name="signUpPassConfirm" placeholder="Confirm Password" type="password" id="cpass" required />
    </div>
    <button type="button" class="Signupbtn" id="submit">
      <span>Signup</span>
    </button>
    <a class="ohref" href="/login">Or login Instead</a>
  </form>
</div>
  `;

  document.getElementById("uploadAvatarBtn").addEventListener("click", () => {
    document.getElementById("avatar").click();
  });

  document
    .getElementById("submit")
    .addEventListener("click", async () => await HandleSignup());

  document.addEventListener("keydown", (e) => {
    if (e && e.key && e.key.toLowerCase() === "enter") {
      document.getElementById("submit").click();
    }
  });

  const svg = document.getElementById("authlogo");
  let degree = 0;

  function animate() {
    degree++;
    if (degree === 360) {
      degree = 0;
    }
    svg.style.transform = `rotateY(${degree}deg)`;
    requestAnimationFrame(animate);
  }

  animate();
};
