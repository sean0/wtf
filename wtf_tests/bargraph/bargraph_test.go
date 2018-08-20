package bargraph_test

import (
	"testing"

	. "github.com/senorprogrammer/wtf/wtf"
	. "github.com/stretchr/testify/assert"
)

// MakeData - Create sample data
func makeData() [][2]int64 {

	//this could come from config
	const lineCount = 2
	var stats [lineCount][2]int64

	stats[0][1] = 1530122942000
	stats[0][0] = 100

	stats[1][1] = 1531142942000
	stats[1][0] = 210

	return stats[:]

}

//TestOutput of the bargraph make string (BuildStars) function
func TestOutput(t *testing.T) {

	result := BuildStars(makeData(), 20, "*")

	Equal(t, "Jun 27, 2018 -\t [red]*[white] - (100)\nJul 09, 2018 -\t [red]********************[white] - (210)\n", result)
}
