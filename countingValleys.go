func countingValleys(steps int32, path string) int32 {
	// Write your code here
	m := make(map[string]int)
	m["U"], m["D"] = 1, -1

	valleys, cLevel, pLevel := 0, 0, 0
	for _, ch := range strings.Split(path, "") {
		pLevel = cLevel
		cLevel += m[ch]
		if cLevel == 0 && pLevel == -1 {
			valleys++
		}
	}

	return int32(valleys)

}
