import { HandleLogin } from "../../funcs/auth"

export const Login = () => {
  document.getElementById("app").innerHTML = /* html */`
    <div class="logo">
      <img src="../../assets/logo.svg" id="authlogo">
    </div>
    <div class="form">
      <form id="signupform" autocomplete="off">
          <input
          name="Name"
          placeholder="Username/email"
          type="text"
          id="cred"
          required
        />
          <input
          name="Password"
          placeholder="Password"
          type="password"
          id="pass"
          required
        />
        <button
          type="button"
          class="SubmissonBtn"
          id="submit"
        > 
          <span>Login</span>
        </button>
        <a class="ohref" href="/signup">Or Signup Instead</a>
      </form>
    </div>
  `

  document.getElementById("submit").addEventListener("click", async () => await HandleLogin())
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