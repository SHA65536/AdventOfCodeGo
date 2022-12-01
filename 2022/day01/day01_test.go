package day01

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var calories string

func TestStar1(t *testing.T) {
	res, err := MaxCalories(calories)
	if err != nil {
		t.FailNow()
	}
	if res != "71924" {
		t.FailNow()
	}
}

func TestStar2(t *testing.T) {
	res, err := MaxCaloriesThree(calories)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(res)
	if res != "210406" {
		t.FailNow()
	}
}
