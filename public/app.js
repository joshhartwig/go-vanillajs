import { API } from "./services/api.js";
import { HomePage } from "./components/HomePage.js";
import "./components/AnimatedLoading.js";



window.addEventListener('DOMContentLoaded', (event) => {
  document.querySelector('main').appendChild(new HomePage());
});

window.app = {
  search: (event) => {
    event.preventDefault();
    const query = document.querySelector("input[type='search']").value;
    console.log(`Searching for: ${query}`);
  },
  api: API
}