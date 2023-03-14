package common

import (
	"testing"
)

func FuzzNewBlockPosition(f *testing.F) {
	f.Add(1, 2, 3)
	f.Add(-33_554_432, -2_048, -33_554_432)
	f.Add(33_554_431, 2_047, 33_554_431)
	f.Fuzz(func(t *testing.T, x, y, z int) {

		if x < -33_554_432 || x > 33_554_431 {
			t.Skip()
		}
		if y < -2_048 || y > 2_047 {
			t.Skip()
		}
		if z < -33_554_432 || z > 33_554_431 {
			t.Skip()
		}

		position := NewBlockPosition(int32(x), int32(y), int32(z))
		if position.X() != int32(x) {
			t.Errorf("Expected X to be %d, got %d", x, position.X())
		}
		if position.Y() != int32(y) {
			t.Errorf("Expected Y to be %d, got %d", y, position.Y())
		}
		if position.Z() != int32(z) {
			t.Errorf("Expected Z to be %d, got %d", z, position.Z())
		}
	})
}
