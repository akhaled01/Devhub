/**
 * Function that expires all document cookies, 
 * clears sessionStorage, and redirects
 * to `/login`
 */
export const Logout = () => {
  sessionStorage.clear();
  document.cookie.split(";").forEach(function (c) {
    document.cookie = c
      .replace(/^ +/, "")
      .replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
  });
  window.location.assign("/login");
};
