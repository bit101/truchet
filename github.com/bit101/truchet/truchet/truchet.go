package truchet

import (
	"fmt"
	"math"
	"os"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
)

// Truchet draw a truchet tile
func Truchet(surface *blgo.Surface, p string, x, y, tileSize, brightness float64) {
	pattern, ok := getPattern(p)
	if !ok {
		fmt.Printf("Pattern '%s' not found.\n", p)
		fmt.Println("Valid patterns: a, b, c, d, e, f, g, h, i, l, m, n, o, p, q, r, s, t, u, v, x, y, z, 0")
		os.Exit(1)
	}
	quarterTile := tileSize * 0.25
	halfTile := tileSize * 0.5
	tile := pattern.GetTile(int(y/tileSize), int(x/tileSize))

	t := blmath.Clamp(2*brightness-0.5, 0, 1)
	midx := blmath.Map(t, 0, 1, -quarterTile, quarterTile)
	midy := blmath.Map(t, 0, 1, quarterTile, -quarterTile)

	surface.Save()
	surface.Translate(x+tileSize/2, y+tileSize/2)
	surface.Rotate(float64(tile) * math.Pi / 2)

	surface.MoveTo(-halfTile, -halfTile)
	surface.LineTo(midx, midy)
	surface.LineTo(halfTile, halfTile)
	surface.LineTo(-halfTile, halfTile)
	surface.Fill()

	surface.Restore()
}

func getPattern(p string) (Pattern, bool) {
	m := map[string]Pattern{
		"a": PatternA,
		"b": PatternB,
		"c": PatternC,
		"d": PatternD,
		"e": PatternE,
		"f": PatternF,
		"g": PatternG,
		"h": PatternH,
		"i": PatternI,
		"l": PatternL,
		"m": PatternM,
		"n": PatternN,
		"o": PatternO,
		"p": PatternP,
		"q": PatternQ,
		"r": PatternR,
		"s": PatternS,
		"t": PatternT,
		"u": PatternU,
		"v": PatternV,
		"x": PatternX,
		"y": PatternY,
		"z": PatternZ,
		"0": Pattern0,
	}
	result, ok := m[p]
	return result, ok
}
