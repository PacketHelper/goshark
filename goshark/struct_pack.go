package goshark

import (
	"errors"
	"fmt"
	"strings"
)

func structPack(format string, a ...int) (int, error) {
	if len(format) != len(a) {
		return 0, errors.New("wrong format or missing arguments")
	}

	var formatList = strings.Split(format, "")

	for i := range a {
		fmt.Println(formatList[i], a[i])
	}
	return 0, nil
}
