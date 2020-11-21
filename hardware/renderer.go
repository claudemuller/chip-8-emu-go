package hardware

// Renderer is where the graphics get rendered to.
type Renderer struct {
    scale int
}

// NewRenderer is a factory that returns a new Renderer.
func NewRenderer(s int) *Renderer {
    return &Renderer{
        scale: s,
    }
}
