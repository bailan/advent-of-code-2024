package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int) {
	var secrets []int
	for _, line := range strings.Split(input, "\n") {
		number, _ := strconv.Atoi(line)
		secrets = append(secrets, number)
	}
	return secrets
}

func nextSecret(secret int) int {
	secret = (secret << 6 ^ secret) % 16777216
	secret = (secret >> 5 ^ secret) % 16777216
	secret = (secret << 11 ^ secret) % 16777216
	return secret
}

func part1(input string) int {
	secrets := parseInput(input)
	total := 0
	for _, secret := range secrets {
		for i := 0; i < 2000; i++ {
			secret = nextSecret(secret)
		}
		total += secret
	}
	return total
}

type Sequence struct {
    A int
	B int
	C int
	D int
}

func part2(input string, iteration int) int {
	secrets := parseInput(input)

	sequenceToTotalPrice := make(map[Sequence]int)
	for _, secret := range secrets {
		changes := make([]int, 4)
		last_secret := secret
		for i := 0; i < 3; i++ {
			secret = nextSecret(secret)
			changes[i + 1] = (secret % 10) - (last_secret % 10)
			last_secret = secret
		}
		sequenceToPrice := make(map[Sequence]int)
		for i := 3; i < iteration; i++ {
			secret = nextSecret(secret)
			changes = append(changes[1:], (secret % 10) - (last_secret % 10))
			last_secret = secret
			sequence := Sequence{changes[0], changes[1], changes[2], changes[3]}
			_, ok := sequenceToPrice[sequence]
			if !ok {
				sequenceToPrice[sequence] = secret % 10
			}
		}
		for sequence, price := range sequenceToPrice {
			sequenceToTotalPrice[sequence] += price
		}
	}
	maxPrice := 0
	for _, price := range sequenceToTotalPrice {
		if price > maxPrice {
			maxPrice = price
		}
	}
	return maxPrice
}

func main() {
	b, err := os.ReadFile("day22.input")
	if err != nil {
		fmt.Print(err)
	}
	input := string(b)
	fmt.Println(part1(input))
	fmt.Println(part2(input, 2000))
}