package main

func LCS(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	if s1[0] == s2[0] {
		return string(s1[0]) + LCS(s1[1:], s2[1:])
	}

	l1 := LCS(s1[1:], s2)
	l2 := LCS(s1, s2[1:])

	if len(l1) > len(l2) {
		return l1
	}

	return l2
}

// Reference: https://www.codewars.com/kata/52756e5ad454534f220001ef
