package utils_test

import (
	"Moonlight/utils"
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	got := utils.Util("a")
	fmt.Println(got)
}
