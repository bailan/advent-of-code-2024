package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
)

func solve1(stones []int) {
    for times := 0; times < 25; times++ {
        var next_stones []int
        for i := range(stones) {
            if stones[i] == 0 {
                next_stones = append(next_stones, 1)
            } else if len(strconv.Itoa(stones[i])) % 2 == 0 {
                str := strconv.Itoa(stones[i])
                left, _ := strconv.Atoi(str[:len(str) / 2])
                next_stones = append(next_stones, left)
                right, _ := strconv.Atoi(str[len(str) / 2:])
                next_stones = append(next_stones, right)
            } else {
                next_stones = append(next_stones, stones[i] * 2024)
            }
        }
        stones = next_stones
    }
    fmt.Println(len(stones))
}

type Key struct {
    N int
    TIMES int
}
var cache = make(map[Key]int)

func stone_count(n int, times int) int {
    if times == 0 {
        return 1
    }
    v, ok := cache[Key{n, times}]
    if ok {
        return v
    }
    var result int
    if n == 0 {
        result = stone_count(1, times - 1)
    } else if len(strconv.Itoa(n)) % 2 == 0 {
        str := strconv.Itoa(n)
        left, _ := strconv.Atoi(str[:len(str) / 2])
        right, _ := strconv.Atoi(str[len(str) / 2:])
        result = stone_count(left, times - 1) + stone_count(right, times - 1)
    } else {
        result = stone_count(n * 2024, times - 1)
    }
    cache[Key{n, times}] = result
    return result
}

func solve2(stones []int) {
    count := 0
    for i := range(stones) {
        count += stone_count(stones[i], 75)
    }
    fmt.Println(count)
}

func main() {
    file, _ := os.Open("day11.input")
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var stones []int
    for scanner.Scan() {
        split := strings.Split(scanner.Text(), " ")
        for i := range(split) {
            value, _ := strconv.Atoi(split[i])
            stones = append(stones, value)
        }
    }
    solve1(stones)
    solve2(stones)
}