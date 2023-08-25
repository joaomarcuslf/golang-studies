package main

import "fmt"

// A format for expressing an ordered list of integers is to use a comma separated list of either

//   individual integers
//   or a range of integers denoted by the starting integer separated from the end integer in the range by a dash, '-'. The range includes all integers in the interval including both endpoints. It is not considered a range unless it spans at least 3 numbers. For example "12,13,15-17"

// Complete the solution so that it takes a list of integers in increasing order and returns a correctly formatted string in the range format.

// Example:
// solution([-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-10--8,-6,-3-1,3-5,7-11,14,15,17-20"

func Support(list, aux []int, response string, item int) ([]int, string) {
	if len(aux) > 0 {
		if item != aux[len(aux)-1]+1 {
			if len(aux) >= 3 {
				if len(response) > 0 {
					response += ","
				}

				response += fmt.Sprintf("%d-%d", aux[0], aux[len(aux)-1])
			} else {
				for j := 0; j < len(aux); j++ {
					if len(response) > 0 {
						response += ","
					}

					response += fmt.Sprintf("%d", aux[j])
				}
			}

			aux = []int{}
		}
	}

	aux = append(aux, item)

	return aux, response
}

func RangeExtraction(list []int) string {
	response := ""
	aux := []int{}

	for i := 0; i < len(list); i++ {
		item := list[i]

		aux, response = Support(list, aux, response, item)
	}

	if len(aux) > 0 {
		aux, response = Support(list, aux, response, aux[0])
	}

	return response
}
