package core

type Color struct {
	R, G, B float32
	Alpha   float32
}

func ColorRGB(r, g, b float32) Color {
	return ColorRGBA(r, g, b, 1.0)
}

func ColorRGBA(r, g, b, a float32) Color {
	return Color{
		R:     r,
		G:     g,
		B:     b,
		Alpha: a,
	}
}
