package dice

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
)

// Result is the interface for the result of
// one or more dice rolls
type Result interface {
	Rolls() []int
	Total() int
}

// result is the result of one or more dice
// rolls
type result struct {
	total int
	rolls []int
}

func (r *result) Total() int {
	return r.total
}

// Rolls returns the individual rolls
func (r *result) Rolls() []int {
	return r.rolls
}

// Roll rolls the dice and returns the result if successful
func Roll(expression string) (Result, error) {

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
		return &result{total: total, rolls: rolls}, nil
	}

	if disadvantage {
		total = keepLowest(rolls)
		return &result{total: total, rolls: rolls}, nil
	}

	return &result{rolls: rolls, total: sum(rolls)}, nil
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
