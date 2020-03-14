package truchet

// PatternA creates truchet pattern A
var PatternA = NewPattern(
	Row{0},
)

// PatternB creates truchet pattern B
var PatternB = NewPattern(
	Row{0},
	Row{2},
)

// PatternC creates truchet pattern C
var PatternC = NewPattern(
	Row{0, 2},
	Row{2, 0},
)

// PatternD creates truchet pattern D
var PatternD = NewPattern(
	Row{1, 0},
	Row{2, 3},
)

// PatternE creates truchet pattern E
var PatternE = NewPattern(
	Row{2, 1, 0, 3},
	Row{3, 0, 1, 2},
	Row{0, 3, 2, 1},
	Row{1, 2, 3, 0},
)

// PatternF creates truchet pattern F
var PatternF = NewPattern(
	Row{1, 2},
	Row{0, 3},
)

// PatternG creates truchet pattern G
var PatternG = NewPattern(
	Row{0, 2},
	Row{2, 0},
	Row{2, 0},
	Row{0, 2},
)

// PatternH creates truchet pattern H
var PatternH = NewPattern(
	Row{1, 3, 3},
	Row{3, 1, 3},
	Row{3, 3, 1},
	Row{1, 3, 3},
	Row{3, 1, 3},
	Row{3, 3, 1},
)

// PatternI creates truchet pattern I
var PatternI = NewPattern(
	Row{2, 0, 0, 2},
	Row{0, 2, 2, 0},
	Row{0, 2, 2, 0},
	Row{2, 0, 0, 2},
)

// PatternL creates truchet pattern L
var PatternL = NewPattern(
	Row{3, 0, 3, 0},
	Row{2, 3, 0, 1},
	Row{3, 2, 1, 0},
	Row{2, 1, 2, 1},
)

// PatternM creates truchet pattern M
var PatternM = NewPattern(
	Row{1, 2},
	Row{2, 1},
)

// PatternN creates truchet pattern N
var PatternN = NewPattern(
	Row{1, 0, 3, 2},
	Row{2, 3, 0, 1},
	Row{3, 2, 1, 0},
	Row{0, 1, 2, 3},
)

// PatternO creates truchet pattern O
var PatternO = NewPattern(
	Row{3, 0, 1, 2},
	Row{1, 2, 3, 0},
)

// PatternP creates truchet pattern P
var PatternP = NewPattern(
	Row{1, 2},
	Row{3, 0},
)

// PatternQ creates truchet pattern Q
var PatternQ = NewPattern(
	Row{3, 3, 0, 0},
	Row{1, 1, 2, 2},
)

// PatternR creates truchet pattern R
var PatternR = NewPattern(
	Row{3, 0},
)

// PatternS creates truchet pattern S
var PatternS = NewPattern(
	Row{3, 0, 1, 2},
)

// PatternT creates truchet pattern T
var PatternT = NewPattern(
	Row{1, 3, 0, 2},
)

// PatternV creates truchet pattern V
var PatternV = NewPattern(
	Row{2, 0, 3, 0, 2, 0, 1, 0},
	Row{0, 3, 0, 2, 0, 1, 0, 2},
	Row{3, 0, 2, 0, 1, 0, 2, 0},
	Row{0, 2, 0, 1, 0, 2, 0, 3},
	Row{2, 0, 1, 0, 2, 0, 3, 0},
	Row{0, 1, 0, 2, 0, 3, 0, 2},
	Row{1, 0, 2, 0, 3, 0, 2, 0},
	Row{0, 2, 0, 3, 0, 2, 0, 1},
)

// PatternU creates truchet pattern U
var PatternU = NewPattern(
	Row{1, 3, 0, 2},
	Row{3, 0, 2, 1},
	Row{0, 2, 1, 3},
	Row{2, 1, 3, 0},
)

// PatternX creates truchet pattern X
var PatternX = NewPattern(
	Row{2, 0, 3, 0, 0, 2},
	Row{2, 1, 2, 2, 0, 0},
	Row{3, 0, 0, 2, 2, 0},
	Row{2, 2, 0, 0, 2, 1},
	Row{0, 2, 2, 0, 3, 0},
	Row{0, 0, 2, 1, 2, 2},
)

// PatternY creates truchet pattern Y
var PatternY = NewPattern(
	Row{1, 3, 0, 2},
	Row{3, 1, 2, 0},
	Row{2, 0, 3, 1},
	Row{0, 2, 1, 3},
)

// PatternZ creates truchet pattern Z
var PatternZ = NewPattern(
	Row{1, 2, 0, 3},
	Row{3, 1, 2, 0},
)

// Pattern0 creates truchet pattern 0
var Pattern0 = NewPattern(
	Row{1, 3, 2, 0},
	Row{3, 1, 0, 2},
	Row{0, 2, 3, 1},
	Row{2, 0, 1, 3},
)

// PatternDouat72 creates truchet pattern Douat 72
var PatternDouat72 = NewPattern(
	Row{1, 3, 3, 1, 3, 3, 0, 0, 2, 0, 0, 2},
	Row{3, 1, 3, 3, 1, 3, 0, 2, 0, 0, 2, 0},
	Row{3, 3, 1, 3, 3, 1, 2, 0, 0, 2, 0, 0},
	Row{1, 3, 3, 1, 3, 3, 0, 0, 2, 0, 0, 2},
	Row{3, 1, 3, 3, 1, 3, 0, 2, 0, 0, 2, 0},
	Row{3, 3, 1, 3, 3, 1, 2, 0, 0, 2, 0, 0},

	Row{2, 2, 0, 2, 2, 0, 3, 1, 1, 3, 1, 1},
	Row{2, 0, 2, 2, 0, 2, 1, 3, 1, 1, 3, 1},
	Row{0, 2, 2, 0, 2, 2, 1, 1, 3, 1, 1, 3},
	Row{2, 2, 0, 2, 2, 0, 3, 1, 1, 3, 1, 1},
	Row{2, 0, 2, 2, 0, 2, 1, 3, 1, 1, 3, 1},
	Row{0, 2, 2, 0, 2, 2, 1, 1, 3, 1, 1, 3},
)
