import { HandleSignup } from "../../funcs/auth";


//TODO - Handle new signup from backend 

export const Signup = () => {
  document.getElementById("app").innerHTML = /* html */`
    <div class="logo">
      <img src="../../images/logo.svg" id="authlogo">
    </div>
    <div class="form">
      <form id="signupform" autocomplete="off">
        <div class="authgroup">
          <input
          name="fname"
          placeholder="First Name"
          type="text"
          id="fname"
          required
        />
        <input
          name="lname"
          placeholder="Last Name"
          type="text"
          id="lname"
          required
        />
        </div>
        <div class="authgroup">
          <input
          name="Name"
          placeholder="Username"
          type="text"
          id="uname"
          required
        />
        <input
          name="Email"
          placeholder="Email"
          type="text"
          id="email"
          required
        />
        </div>
        <div class="authgroup">
          <input
          name="age"
          placeholder="Age"
          type="number"
          id="pass"
          required
        />
        <input
          name="gender"
          placeholder="Gender"
          type="text"
          id="cpass"
          required
        />
        <input
          name="avatar"
          placeholder="Avatar"
          type="file"
          id="avatar"
          required
        />
        </div>
        <div class="authgroup">
          <input
          name="Password"
          placeholder="Password"
          type="password"
          id="pass"
          required
        />
        <input
          name="signUpPassConfirm"
          placeholder="Confirm Password"
          type="password"
          id="cpass"
          required
        />
        </div>
        <button
          type="button"
          class="Signupbtn"
          id="submit"
        >
          <span>SignUp</span>
        </button>
        <a class="ohref" href="/login">Or login Instead</a>
      </form>
    </div>
  `

  document.getElementById("submit").addEventListener("click", async () => await HandleSignup())
  document.addEventListener("keydown", (event) => {
    if (event.key === "Enter") {
      document.getElementById("submit").click()
    }
  })
  const svg = document.getElementById('authlogo');
  let degree = 0;


  function animate() {
    degree++;
    if (degree === 360) {
      degree = 0
    }
    svg.style.transform = `rotateY(${degree}deg)`;
    requestAnimationFrame(animate);
  }

  animate();
}
