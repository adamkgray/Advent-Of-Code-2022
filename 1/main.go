package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	/* open file */
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	var elves []int
	var cumsum int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		val, err := strconv.Atoi(text)
		if err != nil {
			elves = append(elves, cumsum)
			cumsum = 0
		} else {
			cumsum += val
		}
	}

	/* handle error */
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	/* part 1 */
	maxVal := 0
	for _, val := range elves {
		if val > maxVal {
			maxVal = val
		}
	}

	fmt.Println("Part 1:")
	fmt.Println(maxVal)

	/* part 2 */

	/* better solution */
	maxVals := []int{0, 0, 0}
	for _, val := range elves {
		/* compare val to all current maxVals */
		for i := 0; i < 3; i++ {
			/* if its greater than _any_ of them, replace the smallest one */
			if val > maxVals[i] {
				minIndex := 0
				if maxVals[1] < maxVals[minIndex] {
					minIndex = 1
				}
				if maxVals[2] < maxVals[minIndex] {
					minIndex = 2
				}
				maxVals[minIndex] = val

				/* break to next val */
				break
			}

		}
	}
	maxSum := maxVals[0] + maxVals[1] + maxVals[2]
	fmt.Println("Part 2 (Better):")
	fmt.Println(maxSum)

	/* naive solution */
	sort.Ints(elves)
	maxSum = elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3]
	fmt.Println("Part 2 (Naive):")
	fmt.Println(maxSum)

}
