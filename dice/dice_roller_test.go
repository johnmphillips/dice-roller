package dice

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func Test_RollsCorrectNumberOfDice(t *testing.T) {

	for i := 0; i < 1000; i++ {
		numberOfDice := rand.Intn(20) + 1
		diceSize := rand.Intn(20) + 1

		expression := fmt.Sprintf("%dd%d", numberOfDice, diceSize)
		result, _ := Roll(expression)

		if len(result.rolls) != numberOfDice {
			t.Errorf("%s: Expected %d dice to be rolled but got %d", expression, numberOfDice, len(result.rolls))
		}
	}

}

func Test_RollsCorrectSizeOfDice(t *testing.T) {

	for i := 0; i < 1000; i++ {
		numberOfDice := rand.Intn(20) + 1
		diceSize := rand.Intn(20) + 1

		expression := fmt.Sprintf("%dd%d", numberOfDice, diceSize)
		result, _ := Roll(expression)

		for _, v := range result.rolls {
			if v > diceSize || v < 1 {
				t.Errorf("%s: Dice values should be between 1 and %d. Got %v", expression, diceSize, result.rolls)
			}
		}
	}
}

func Test_RollSum(t *testing.T) {

	for i := 0; i < 1000; i++ {
		numberOfDice := rand.Intn(20) + 1
		diceSize := rand.Intn(20) + 1

		expression := fmt.Sprintf("%dd%d", numberOfDice, diceSize)
		result, _ := Roll(expression)

		sumOfRolls := 0

		for _, v := range result.rolls {
			sumOfRolls += v
		}

		if result.result != sumOfRolls {
			t.Errorf("%s: Result should equal the sum of the rolls. Got %v", expression, result.rolls)
		}

	}

}

func Test_RollKeepLowest(t *testing.T) {
	for i := 0; i < 1000; i++ {
		numberOfDice := rand.Intn(20) + 1
		diceSize := rand.Intn(20) + 1

		expression := fmt.Sprintf("%dd%dkl", numberOfDice, diceSize)
		result, _ := Roll(expression)

		lowestRoll := math.MaxInt

		for _, v := range result.rolls {
			if v < lowestRoll {
				lowestRoll = v
			}
		}

		if result.result != lowestRoll {
			t.Errorf("%s: Result should equal the lowest roll. Got %v", expression, result.rolls)
		}
	}
}

func Test_RollKeepHighest(t *testing.T) {
	for i := 0; i < 1000; i++ {
		numberOfDice := rand.Intn(20) + 1
		diceSize := rand.Intn(20) + 1

		expression := fmt.Sprintf("%dd%dkh", numberOfDice, diceSize)
		result, _ := Roll(expression)

		highestRoll := 0

		for _, v := range result.rolls {
			if v > highestRoll {
				highestRoll = v
			}
		}

		if result.result != highestRoll {
			t.Errorf("%s: Result should equal the highestRoll roll. Got %v", expression, result.rolls)
		}
	}
}

func Test_ErrorsWhenGivenInvalidExpression(t *testing.T) {
	_, err := Roll("john")
	if err == nil {
		t.Errorf("Expected an error to be thrown but it wasn't")
	}
	want := "john does not appear to be a valid dice roll expression"
	if err.Error() != want {
		t.Errorf("Wanted error message: [%s] Got: [%s]", want, err.Error())
	}
}

func Test_RollExplodingDice(t *testing.T) {
	for i := 0; i < 1000; i++ {
		result, _ := Roll("2d4!")
		numberOfFours := 0

		fmt.Println(result)
		for _, v := range result.rolls {
			if v == 4 {
				numberOfFours++
			}
		}

		want := numberOfFours + 2
		if len(result.rolls) != want {
			t.Errorf("Dice did not explode correctly: Wanted [%d total rolls] Got: [%d total rolls]", want, len(result.rolls))
		}

	}
}
