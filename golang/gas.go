package main

import "fmt"
import "strings"
import "strconv"

func isCircular(stations []int, costs []int) int {
	size := len(stations)
	if size == 0 || size != len(costs) {
		return -1
	}
	count := 0
	index := 0
	balance := 0
	max := size * 2
	for {
		balance = balance + stations[index] - costs[index]
		if balance == 0 {
			return (index + 1) % size
		}
		if balance < 0 {
			balance = 0
		}
		index = (index + 1) % size
		count++
		if count == max {
			break
		}
	}
	return -1
}

func extractStationsAndCosts(input []string) ([]int, []int) {
	stations := make([]int, 0)
	costs := make([]int, 0)
	for i := 1; i < len(input); i++ {
		elem := strings.Split(input[i], ":")
		if len(elem) != 2 {
			continue
		}
		if v, err := strconv.Atoi(elem[0]); err == nil {
			stations = append(stations, v)
		}
		if v, err := strconv.Atoi(elem[1]); err == nil {
			costs = append(costs, v)
		}
	}
	return stations, costs
}

func main() {
	input := []string{"5", "1:3", "2:4", "3:5", "4:1", "5:2"}
	stations, costs := extractStationsAndCosts(input)
	fmt.Println(isCircular(stations[:], costs[:]))
}
