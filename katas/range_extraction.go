package main

import "fmt"

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

// Reference: https://www.codewars.com/kata/51ba717bb08c1cd60f00002f
