export class AnimatedLoading extends HTMLElement {
  constructor(){
    super();
  }
  connectedCallback(){
    let qty = this.dataset.elements ?? 1;
    let width = this.dataset.width ?? 100;
    let height = this.dataset.height ?? 20;
    for (let i = 0; i < qty; i++) {
      const wrapper = document.createElement('div');
      wrapper.setAttribute('class', 'loading-wave');
      wrapper.style.width = width + 'px';
      wrapper.style.height = height + 'px';
      wrapper.style.margin = '10px';
      wrapper.style.display = 'inline-block';
      this.appendChild(wrapper);
    }
  }
}
customElements.define("animated-loading", AnimatedLoading);