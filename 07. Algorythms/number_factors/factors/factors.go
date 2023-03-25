// Package factors ...
package factors

// NumberFactors ...
func NumberFactors(n int) (res []int) {
	divisor := 2
	if n == 0 {
		return []int{0}
	}
	for flag := false; !flag; {
		if n%divisor == 0 {
			res = append(res, divisor)
			n /= divisor
		} else {
			divisor++
		}

		if n == 1 {
			flag = true
		}
	}

	return res
}
