package main

func isOneBitCharacter(bits []int) bool {
	n := len(bits)

	for i := 0; i < n-1; {
		if bits[i] == 1 && i == n-2 {
			return false
		}
		if bits[i] == 1 {
			i += 2
			continue
		}
		if bits[i] == 0 {
			i += 1
			continue
		}
	}

	return true
}
