export class MovieItem extends HTMLElement {
  constructor(movie){
    super();
    this.movie = movie;
  }
  connectedCallback(){
    this.innerHTML = `
      <a href="#${this.movie.id}">
        <article>
          <img src="${this.movie.poster_url}" alt="${this.movie.title} poster">
          <p>${this.movie.title}</p>
        </article>
      </a>
    `
  }
}
customElements.define("movie-item", MovieItem);