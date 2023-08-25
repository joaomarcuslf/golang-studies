package main

func Cakes(recipe, available map[string]int) int {
	count := 0

	for k, v := range recipe {
		if _, ok := available[k]; !ok {
			return 0
		}

		if available[k] >= v {
			available[k] = available[k] / v
		} else {
			return 0
		}

		if count == 0 || available[k] < count {
			count = available[k]
		}
	}

	return count
}

// Reference: https://www.codewars.com/kata/525c65e51bf619685c000059
