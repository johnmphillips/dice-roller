package dice

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
)

type rollresult struct {
	result int
	rolls  []int
}

func Roll(expression string) (*rollresult, error) {

	re := regexp.MustCompile(`^([0-9]*)d([0-9]*)(kh|kl|!|)$`)

	if !re.MatchString(expression) {
		return nil, fmt.Errorf("%s does not appear to be a valid dice roll expression", expression)
	}
	matches := re.FindStringSubmatch(expression)
	rollQuantity, _ := strconv.ParseInt(matches[1], 0, 64)
	diceSize, _ := strconv.ParseInt(matches[2], 0, 64)

	explode := matches[3] == "!"
	advantage := matches[3] == "kh"
	disadvantage := matches[3] == "kl"

	var rolls []int

	for i := 0; i < int(rollQuantity); i++ {
		currentRoll := rand.Intn(int(diceSize)) + 1
		if currentRoll == int(diceSize) && explode {
			// if we roll max value on the dice, the dice "explodes" and we roll another
			i--
		}
		rolls = append(rolls, currentRoll)
	}

	total := 0
	if advantage {
		total = keepHighest(rolls)
	} else if disadvantage {
		total = keepLowest(rolls)
	} else {
		total = sum(rolls)
	}

	result := rollresult{rolls: rolls, result: total}
	return &result, nil
}

func sum(rolls []int) int {
	total := 0
	for _, v := range rolls {
		total += v
	}
	return total
}

func keepHighest(rolls []int) int {
	highest := 0
	for _, v := range rolls {
		if v > highest {
			highest = v
		}
	}
	return highest
}

func keepLowest(rolls []int) int {
	lowest := math.MaxInt
	for _, v := range rolls {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}
