package LeetCode_Go

func lengthOfLongestSubstring(s string) int {
	charIndex := make([byte]int)

	maxLength := 0
	curLen := 0
	startIndex := 0

	for i := 0; i <len(s); i++ {
		char := s[i]
		if charIdx, ok := charIndex[char]; ok && charIdx >= startIndex {
			if maxLength < curLen {
				maxLength = curLen
			}

			charIndex[char] = i
			startIndex = charIdx + 1
			curLen = i - charIdx
		} else {
			charIndex[char] = i
			curLen++
		}
	}

	if maxLength <curLen {
		maxLength = curLen
	}

	return maxLength
}
