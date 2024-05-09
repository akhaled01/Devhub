import { BACKENDURL } from "./vars";

/**
 * Function that expires all document cookies,
 * logs out from the backned,
 * clears sessionStorage, and redirects
 * to `/login`
 */
export const Logout = async () => {
  sessionStorage.clear();
  const res = await fetch(BACKENDURL + "/auth/logout", {
    method: "POST",
    credentials: "include",
  });
  document.cookie.split(";").forEach(function (c) {
    document.cookie = c
      .replace(/^ +/, "")
      .replace(/=.*/, "=;expires=" + new Date().toUTCString() + ";path=/");
  });
  window.location.assign("/login");
};
