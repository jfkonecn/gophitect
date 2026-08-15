class FunctionComponent extends HTMLElement {
    connectedCallback() {
        this.textContent = "Hello, World!";
    }
}

customElements.define("function-component", FunctionComponent);
