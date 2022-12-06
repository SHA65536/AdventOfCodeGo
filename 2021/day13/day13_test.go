package day13

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input string

func TestStar1(t *testing.T) {
	res, err := FoldInstructions(input)
	if err != nil || res != "765" {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FoldInstructions(input)
	}
}

var part_two_ans = `###  #### #  # #### #    ###   ##  #  #
#  #    # # #     # #    #  # #  # #  #
#  #   #  ##     #  #    #  # #    ####
###   #   # #   #   #    ###  # ## #  #
# #  #    # #  #    #    #    #  # #  #
#  # #### #  # #### #### #     ### #  #
`

func TestStar2(t *testing.T) {
	res, err := FoldAll(input)
	if err != nil || res != part_two_ans {
		fmt.Println(res)
		t.FailNow()
	}
}

func BenchmarkStar2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FoldAll(input)
	}
}
