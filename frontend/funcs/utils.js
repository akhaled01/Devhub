import { BACKENDURL } from "./vars";

/**
 * 
 * Follow up login after signup
 * 
 * @param {string} email 
 * @param {string} pass 
 */
export const Flogin = async (email, pass) => {
  const res = await fetch(BACKENDURL + "/login", {
    method: "POST",
    body: JSON.stringify({
      Email: email,
      Password: pass,
    }),
  })

  if (res.ok) {
    window.location.assign("/")
  }
};

/**
 * Function to Update CSS en routing
 * @param {Stylesheet} stylesheet - Path to css file
 */
export const UpdateCSS = (stylesheet) => {
  const linkElement = document.getElementById('page-styles');
  if (linkElement) {
    linkElement.href = stylesheet;
    const styleTags = document.querySelectorAll('style');
    styleTags.forEach(tag => tag.remove());
  } else {
    console.error('Page stylesheet link not found');
  }
};
