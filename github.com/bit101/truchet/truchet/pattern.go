package truchet

// Row is one row in a pattern
type Row []int

// Pattern holds a 2d truchet pattern definition
type Pattern struct {
	tiles    []Row
	colCount int
	rowCount int
}

// NewPattern creates and inits a new pattern
func NewPattern(tiles ...Row) Pattern {
	var rows []Row
	for _, row := range tiles {
		rows = append(rows, row)
	}
	return Pattern{
		tiles:    rows,
		colCount: len(tiles),
		rowCount: len(tiles[0]),
	}
}

// GetTile returns a tile for a given column/row
func (p Pattern) GetTile(c, r int) int {
	return p.tiles[c%p.colCount][r%p.rowCount]
}
