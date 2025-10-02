// GHIRIMOLDI LUCA 31974A

package utils

import "math"

type Coordinata struct {
	x, y int
}

func (p0 Coordinata) Distanza(p1 Coordinata) int {
	return int(math.Abs(float64(p1.x-p0.x)) + math.Abs(float64(p1.y-p0.y)))
}

// ritorna true se la coordinata c è contenuta nel rettangolo definito da k e v
func (c Coordinata) isInCoord(k Coordinata, v Coordinata) bool {
	return c.x >= k.x && c.x <= v.x && c.y >= k.y && c.y <= v.y
}
