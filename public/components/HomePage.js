import { API } from '../services/api.js';
import { MovieItem } from './MovieItem.js';

export class HomePage extends HTMLElement {
  async render() {
    const topMovies = await API.getTopMovies();
    const randomMovies = await API.getRandomMovies();
    renderMoviesInList(topMovies, this.querySelector('#top-movies-list'));
    renderMoviesInList(randomMovies, this.querySelector('#random-movies-list'));


    function renderMoviesInList(movies, listElement) {
      listElement.innerHTML = '';
      movies.forEach(movie => {
        const li = document.createElement('li');
        li.appendChild(new MovieItem(movie));
        listElement.appendChild(li);
      });
  }
  }
  connectedCallback(){
    const template = document.getElementById('template-home');
    const content = template.content.cloneNode(true);
    this.appendChild(content);
    this.render();
  }
}

customElements.define("home-page", HomePage);