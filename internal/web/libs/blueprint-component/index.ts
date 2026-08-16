type Point = { x: number; y: number };
type BlueprintShape = BlueprintRect | BlueprintCircle;

const numberAttribute = (element: HTMLElement, name: string, fallback: number) => {
  const value = Number(element.getAttribute(name));
  return Number.isFinite(value) ? value : fallback;
};

const setNumberAttribute = (element: HTMLElement, name: string, value: number) => {
  element.setAttribute(name, String(Math.round(value * 100) / 100));
};

const cssValue = (styles: CSSStyleDeclaration, name: string, fallback: string) => {
  const value = styles.getPropertyValue(name).trim();
  return value || fallback;
};

class BlueprintComponent extends HTMLElement {
  static get observedAttributes() {
    return ["grid", "grid-size", "pan-x", "pan-y", "zoom"];
  }

  private canvas: HTMLCanvasElement;
  private context: CanvasRenderingContext2D;
  private resizeObserver: ResizeObserver;
  private mutationObserver: MutationObserver;
  private animationFrame = 0;
  private panStart = { clientX: 0, clientY: 0, panX: 0, panY: 0 };
  private dragStart = { offsetX: 0, offsetY: 0 };
  private draggedShape: BlueprintShape | null = null;
  private isPanning = false;

  constructor() {
    super();

    this.canvas = document.createElement("canvas");
    const context = this.canvas.getContext("2d");
    if (!context) {
      throw new Error("Blueprint canvas context is not available.");
    }

    this.context = context;
    this.resizeObserver = new ResizeObserver(() => this.requestDraw());
    this.mutationObserver = new MutationObserver(() => this.requestDraw());
  }

  connectedCallback() {
    this.style.display ||= "block";
    this.style.position ||= "relative";
    this.style.overflow ||= "hidden";
    this.style.minHeight ||= "70vh";
    this.style.background ||= "var(--blueprint-background, var(--base-color, #ffffff))";
    this.style.border ||= "var(--border-width-0, 1px) solid var(--stroke-weak-color, #d9d9d9)";
    this.style.touchAction = "none";

    this.canvas.style.display = "block";
    this.canvas.style.width = "100%";
    this.canvas.style.height = "100%";
    this.canvas.style.cursor = "grab";
    this.canvas.style.touchAction = "none";

    if (!this.canvas.isConnected) {
      this.prepend(this.canvas);
    }

    this.resizeObserver.observe(this);
    this.mutationObserver.observe(this, {
      attributes: true,
      childList: true,
      characterData: true,
      subtree: true,
    });

    this.addEventListener("blueprintchange", this.handleBlueprintChange);
    this.canvas.addEventListener("pointerdown", this.handlePointerDown);
    this.canvas.addEventListener("wheel", this.handleWheel, { passive: false });
    window.addEventListener("resize", this.handleResize);
    this.requestDraw();
  }

  disconnectedCallback() {
    this.resizeObserver.disconnect();
    this.mutationObserver.disconnect();
    this.removeEventListener("blueprintchange", this.handleBlueprintChange);
    this.canvas.removeEventListener("pointerdown", this.handlePointerDown);
    this.canvas.removeEventListener("wheel", this.handleWheel);
    window.removeEventListener("resize", this.handleResize);
    cancelAnimationFrame(this.animationFrame);
  }

  attributeChangedCallback() {
    this.requestDraw();
  }

  get grid() {
    return this.hasAttribute("grid");
  }

  set grid(value: boolean) {
    this.toggleAttribute("grid", value);
  }

  get gridSize() {
    return numberAttribute(this, "grid-size", 32);
  }

  set gridSize(value: number) {
    setNumberAttribute(this, "grid-size", value);
  }

  get panX() {
    return numberAttribute(this, "pan-x", 0);
  }

  set panX(value: number) {
    setNumberAttribute(this, "pan-x", value);
  }

  get panY() {
    return numberAttribute(this, "pan-y", 0);
  }

  set panY(value: number) {
    setNumberAttribute(this, "pan-y", value);
  }

  get zoom() {
    return numberAttribute(this, "zoom", 1);
  }

  set zoom(value: number) {
    setNumberAttribute(this, "zoom", Math.min(Math.max(value, 0.25), 3));
  }

  clientToWorld(clientX: number, clientY: number): Point {
    const rect = this.canvas.getBoundingClientRect();
    return {
      x: (clientX - rect.left - this.panX) / this.zoom,
      y: (clientY - rect.top - this.panY) / this.zoom,
    };
  }

  private get shapes() {
    return Array.from(this.children).filter(
      (element): element is BlueprintShape =>
        element instanceof BlueprintRect || element instanceof BlueprintCircle,
    );
  }

  private get connections() {
    return Array.from(this.children).filter(
      (element): element is BlueprintConnection => element instanceof BlueprintConnection,
    );
  }

  private requestDraw() {
    cancelAnimationFrame(this.animationFrame);
    this.animationFrame = requestAnimationFrame(() => this.draw());
  }

  private draw() {
    const rect = this.getBoundingClientRect();
    const pixelRatio = window.devicePixelRatio || 1;

    this.canvas.width = Math.max(1, Math.floor(rect.width * pixelRatio));
    this.canvas.height = Math.max(1, Math.floor(rect.height * pixelRatio));
    this.context.setTransform(pixelRatio, 0, 0, pixelRatio, 0, 0);
    this.context.clearRect(0, 0, rect.width, rect.height);

    const styles = getComputedStyle(this);
    this.drawBackground(rect, styles);

    this.context.save();
    this.context.translate(this.panX, this.panY);
    this.context.scale(this.zoom, this.zoom);
    this.drawConnections(styles);
    this.drawShapes(styles);
    this.context.restore();
  }

  private drawBackground(rect: DOMRect, styles: CSSStyleDeclaration) {
    this.context.fillStyle = cssValue(
      styles,
      "--blueprint-background",
      cssValue(styles, "--base-color", "#ffffff"),
    );
    this.context.fillRect(0, 0, rect.width, rect.height);

    if (!this.grid) {
      return;
    }

    const gridSize = Math.max(4, this.gridSize * this.zoom);
    const offsetX = ((this.panX % gridSize) + gridSize) % gridSize;
    const offsetY = ((this.panY % gridSize) + gridSize) % gridSize;

    this.context.beginPath();
    for (let x = offsetX; x <= rect.width; x += gridSize) {
      this.context.moveTo(x, 0);
      this.context.lineTo(x, rect.height);
    }
    for (let y = offsetY; y <= rect.height; y += gridSize) {
      this.context.moveTo(0, y);
      this.context.lineTo(rect.width, y);
    }
    this.context.strokeStyle = cssValue(
      styles,
      "--blueprint-grid-color",
      cssValue(styles, "--stroke-weak-color", "#e5e5e5"),
    );
    this.context.lineWidth = 1;
    this.context.stroke();
  }

  private drawConnections(styles: CSSStyleDeclaration) {
    const lineColor = cssValue(
      styles,
      "--blueprint-line-color",
      cssValue(styles, "--brand-color", "#265899"),
    );

    for (const connection of this.connections) {
      const from = this.findShape(connection.from);
      const to = this.findShape(connection.to);
      if (!from || !to) {
        continue;
      }

      const start = from.connectionPointTo(to.centerX, to.centerY);
      const end = to.connectionPointTo(from.centerX, from.centerY);

      this.context.beginPath();
      this.context.moveTo(start.x, start.y);
      this.context.lineTo(end.x, end.y);
      this.context.strokeStyle = lineColor;
      this.context.lineWidth = 2 / this.zoom;
      this.context.stroke();

      if (connection.arrowStart) {
        this.drawArrow(end, start, lineColor);
      }
      if (connection.arrowEnd) {
        this.drawArrow(start, end, lineColor);
      }
    }
  }

  private drawShapes(styles: CSSStyleDeclaration) {
    const stroke = cssValue(
      styles,
      "--blueprint-shape-stroke",
      cssValue(styles, "--stroke-strong-color", "#555555"),
    );
    const fill = cssValue(
      styles,
      "--blueprint-shape-fill",
      cssValue(styles, "--raised-color", "#ffffff"),
    );
    const textColor = cssValue(
      styles,
      "--blueprint-shape-color",
      cssValue(styles, "--text-strong-color", "#111111"),
    );
    const fontFamily = cssValue(styles, "--sans-font", "sans-serif");

    for (const shape of this.shapes) {
      this.context.beginPath();
      this.context.fillStyle = fill;
      this.context.strokeStyle = stroke;
      this.context.lineWidth = 2 / this.zoom;

      if (shape instanceof BlueprintCircle) {
        this.context.arc(shape.x, shape.y, shape.radius, 0, Math.PI * 2);
      } else {
        this.context.roundRect(shape.x, shape.y, shape.width, shape.height, 8);
      }

      this.context.fill();
      this.context.stroke();
      this.drawShapeText(shape, textColor, fontFamily);
    }
  }

  private drawShapeText(shape: BlueprintShape, color: string, fontFamily: string) {
    const lines = (shape.textContent || "")
      .split(/\s+/)
      .filter(Boolean)
      .reduce<string[]>((result, word) => {
        const current = result[result.length - 1] || "";
        const next = current ? `${current} ${word}` : word;

        if (next.length > 22 && current) {
          result.push(word);
        } else if (result.length === 0) {
          result.push(next);
        } else {
          result[result.length - 1] = next;
        }

        return result;
      }, [])
      .slice(0, 4);

    if (lines.length === 0) {
      return;
    }

    this.context.fillStyle = color;
    this.context.font = `14px ${fontFamily}`;
    this.context.textAlign = "center";
    this.context.textBaseline = "middle";

    const lineHeight = 18;
    const startY = shape.centerY - ((lines.length - 1) * lineHeight) / 2;
    lines.forEach((line, index) => {
      this.context.fillText(line, shape.centerX, startY + index * lineHeight, shape.textWidth);
    });
  }

  private drawArrow(from: Point, to: Point, color: string) {
    const angle = Math.atan2(to.y - from.y, to.x - from.x);
    const size = 12 / this.zoom;

    this.context.beginPath();
    this.context.moveTo(to.x, to.y);
    this.context.lineTo(
      to.x - size * Math.cos(angle - Math.PI / 6),
      to.y - size * Math.sin(angle - Math.PI / 6),
    );
    this.context.lineTo(
      to.x - size * Math.cos(angle + Math.PI / 6),
      to.y - size * Math.sin(angle + Math.PI / 6),
    );
    this.context.closePath();
    this.context.fillStyle = color;
    this.context.fill();
  }

  private findShape(id: string) {
    return this.shapes.find((shape) => shape.id === id || shape.getAttribute("name") === id);
  }

  private shapeAt(point: Point) {
    return this.shapes.toReversed().find((shape) => shape.containsPoint(point));
  }

  private handleBlueprintChange = () => {
    this.requestDraw();
  };

  private handleResize = () => {
    this.requestDraw();
  };

  private handlePointerDown = (event: PointerEvent) => {
    if (event.button !== 0) {
      return;
    }

    const point = this.clientToWorld(event.clientX, event.clientY);
    const shape = this.shapeAt(point);

    this.canvas.setPointerCapture(event.pointerId);
    if (shape) {
      this.draggedShape = shape;
      this.dragStart = { offsetX: point.x - shape.x, offsetY: point.y - shape.y };
      this.canvas.style.cursor = "move";
    } else {
      this.isPanning = true;
      this.panStart = {
        clientX: event.clientX,
        clientY: event.clientY,
        panX: this.panX,
        panY: this.panY,
      };
      this.canvas.style.cursor = "grabbing";
    }

    this.canvas.addEventListener("pointermove", this.handlePointerMove);
    this.canvas.addEventListener("pointerup", this.handlePointerEnd);
    this.canvas.addEventListener("pointercancel", this.handlePointerEnd);
  };

  private handlePointerMove = (event: PointerEvent) => {
    if (this.draggedShape) {
      const point = this.clientToWorld(event.clientX, event.clientY);
      this.draggedShape.x = point.x - this.dragStart.offsetX;
      this.draggedShape.y = point.y - this.dragStart.offsetY;
      return;
    }

    if (this.isPanning) {
      this.panX = this.panStart.panX + event.clientX - this.panStart.clientX;
      this.panY = this.panStart.panY + event.clientY - this.panStart.clientY;
    }
  };

  private handlePointerEnd = (event: PointerEvent) => {
    this.canvas.releasePointerCapture(event.pointerId);
    this.draggedShape = null;
    this.isPanning = false;
    this.canvas.style.cursor = "grab";
    this.canvas.removeEventListener("pointermove", this.handlePointerMove);
    this.canvas.removeEventListener("pointerup", this.handlePointerEnd);
    this.canvas.removeEventListener("pointercancel", this.handlePointerEnd);
  };

  private handleWheel = (event: WheelEvent) => {
    event.preventDefault();

    const rect = this.canvas.getBoundingClientRect();
    const worldX = (event.clientX - rect.left - this.panX) / this.zoom;
    const worldY = (event.clientY - rect.top - this.panY) / this.zoom;
    const nextZoom = Math.min(Math.max(this.zoom * Math.exp(-event.deltaY * 0.001), 0.25), 3);

    this.zoom = nextZoom;
    this.panX = event.clientX - rect.left - worldX * nextZoom;
    this.panY = event.clientY - rect.top - worldY * nextZoom;
  };
}

class BlueprintShapeElement extends HTMLElement {
  static get observedAttributes() {
    return ["x", "y", "width", "height", "radius"];
  }

  connectedCallback() {
    this.hidden = true;
    this.style.display = "none";
    this.dispatchChange();
  }

  attributeChangedCallback() {
    this.dispatchChange();
  }

  get x() {
    return numberAttribute(this, "x", 0);
  }

  set x(value: number) {
    setNumberAttribute(this, "x", value);
  }

  get y() {
    return numberAttribute(this, "y", 0);
  }

  set y(value: number) {
    setNumberAttribute(this, "y", value);
  }

  get centerX() {
    return this.x;
  }

  get centerY() {
    return this.y;
  }

  get textWidth() {
    return 140;
  }

  containsPoint(point: Point) {
    return point.x === this.x && point.y === this.y;
  }

  connectionPointTo(x: number, y: number): Point {
    return { x: this.centerX, y: this.centerY };
  }

  protected dispatchChange() {
    this.dispatchEvent(
      new CustomEvent("blueprintchange", {
        bubbles: true,
        detail: { element: this, x: this.x, y: this.y },
      }),
    );
  }
}

class BlueprintRect extends BlueprintShapeElement {
  get width() {
    return numberAttribute(this, "width", 180);
  }

  set width(value: number) {
    setNumberAttribute(this, "width", value);
  }

  get height() {
    return numberAttribute(this, "height", 96);
  }

  set height(value: number) {
    setNumberAttribute(this, "height", value);
  }

  get centerX() {
    return this.x + this.width / 2;
  }

  get centerY() {
    return this.y + this.height / 2;
  }

  get textWidth() {
    return Math.max(24, this.width - 24);
  }

  containsPoint(point: Point) {
    return (
      point.x >= this.x &&
      point.x <= this.x + this.width &&
      point.y >= this.y &&
      point.y <= this.y + this.height
    );
  }

  connectionPointTo(x: number, y: number): Point {
    const halfWidth = this.width / 2;
    const halfHeight = this.height / 2;
    const dx = x - this.centerX;
    const dy = y - this.centerY;
    const scale = Math.min(
      dx === 0 ? Infinity : Math.abs(halfWidth / dx),
      dy === 0 ? Infinity : Math.abs(halfHeight / dy),
    );

    return {
      x: this.centerX + dx * scale,
      y: this.centerY + dy * scale,
    };
  }
}

class BlueprintCircle extends BlueprintShapeElement {
  get radius() {
    return numberAttribute(this, "radius", 56);
  }

  set radius(value: number) {
    setNumberAttribute(this, "radius", value);
  }

  get textWidth() {
    return Math.max(24, this.radius * 1.6);
  }

  containsPoint(point: Point) {
    const dx = point.x - this.x;
    const dy = point.y - this.y;
    return Math.hypot(dx, dy) <= this.radius;
  }

  connectionPointTo(x: number, y: number): Point {
    const angle = Math.atan2(y - this.y, x - this.x);
    return {
      x: this.x + Math.cos(angle) * this.radius,
      y: this.y + Math.sin(angle) * this.radius,
    };
  }
}

class BlueprintConnection extends HTMLElement {
  static get observedAttributes() {
    return ["from", "to", "arrow-start", "arrow-end"];
  }

  connectedCallback() {
    this.hidden = true;
    this.style.display = "none";
    this.dispatchChange();
  }

  attributeChangedCallback() {
    this.dispatchChange();
  }

  get from() {
    return this.getAttribute("from") || "";
  }

  set from(value: string) {
    this.setAttribute("from", value);
  }

  get to() {
    return this.getAttribute("to") || "";
  }

  set to(value: string) {
    this.setAttribute("to", value);
  }

  get arrowStart() {
    return this.hasAttribute("arrow-start");
  }

  set arrowStart(value: boolean) {
    this.toggleAttribute("arrow-start", value);
  }

  get arrowEnd() {
    return this.hasAttribute("arrow-end");
  }

  set arrowEnd(value: boolean) {
    this.toggleAttribute("arrow-end", value);
  }

  private dispatchChange() {
    this.dispatchEvent(new CustomEvent("blueprintchange", { bubbles: true, detail: { element: this } }));
  }
}

customElements.define("blueprint-component", BlueprintComponent);
customElements.define("blueprint-rect", BlueprintRect);
customElements.define("blueprint-circle", BlueprintCircle);
customElements.define("blueprint-connection", BlueprintConnection);
