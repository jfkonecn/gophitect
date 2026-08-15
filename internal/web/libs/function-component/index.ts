class FunctionComponent extends HTMLElement {
  connectedCallback() {
    this.innerHTML = `
            <div class="handle">Function Component</div>
            <div>Hello, World!</div>
        `;

    this.style.position = "absolute";
    this.style.left = "100px";
    this.style.top = "100px";

    this.style.width = "300px";
    this.style.height = "200px";

    this.style.resize = "both";
    this.style.overflow = "auto";

    this.style.border = "1px solid black";
    this.style.background = "white";

    const handle = this.querySelector(".handle") as HTMLDivElement;

    handle.style.cursor = "move";
    handle.style.userSelect = "none";

    handle.addEventListener("pointerdown", (event) => {
      const rect = this.getBoundingClientRect();

      const offsetX = event.clientX - rect.left;
      const offsetY = event.clientY - rect.top;

      const move = (event: PointerEvent) => {
        this.style.left = `${event.clientX - offsetX}px`;
        this.style.top = `${event.clientY - offsetY}px`;
      };

      const stop = () => {
        window.removeEventListener("pointermove", move);
        window.removeEventListener("pointerup", stop);
      };

      window.addEventListener("pointermove", move);
      window.addEventListener("pointerup", stop);
    });
  }
}

customElements.define("function-component", FunctionComponent);
