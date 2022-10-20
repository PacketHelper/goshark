package goshark

import (
	"errors"
	"log"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStructPack(t *testing.T) {
	output, err := structPack("II", 1, 2)
	if err != nil {
		log.Fatal()
	}
	assert.Equal(t, output, 0)
}

func TestNegativeStructPackIncorrectFormat(t *testing.T) {
	header := "IIHHIIII"
	_, err := structPack(header, 1, 2, 3, 4, 5)
	assert.Equal(t, err, errors.New("wrong format or missing arguments"))
}
