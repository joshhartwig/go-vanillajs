import { api } from "./services/api.js"

window.app = {
  search: (event) => {
    event.preventDefault();
    const query = document.querySelector("input[type='search']").value;
    console.log(`Searching for: ${query}`);
  },
  api: api
}