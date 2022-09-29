package decode

import (
	"reflect"
	"testing"
)

func TestHex2Array(t *testing.T) {
	want := []string{"00", "11", "22", "aa", "bb", "cc"}
	got, err := Hex2Array("001122aabbcc")

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHex2Array_IncorrectLengthOfHexOnInput(t *testing.T) {
	hexWithWrongSize := "001122a"

	_, err := Hex2Array(hexWithWrongSize)
	if err == nil {
		t.Error("func should return an error")
	}
}
