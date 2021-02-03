package live

import (
	"testing"

	"github.com/edwingeng/live/internal"
)

func TestData_Persistent(t *testing.T) {
	d1 := Nil
	if _, ok := d1.Persistent(); !ok {
		t.Fatal("Persistent() does not work as expected")
	}

	var ptr *internal.Data
	d2 := Data{v: ptr}
	if _, ok := d2.Persistent(); !ok {
		t.Fatal("Persistent() does not work as expected")
	}
}
