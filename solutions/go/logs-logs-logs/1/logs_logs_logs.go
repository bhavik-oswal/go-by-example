package logs
func Application(log string) string {
	for _, r := range log { 
		switch r {
		case '❗': 
			return "recommendation"
		case '🔍': 
			return "search"
		case '☀': 
			return "weather"
		}
	}
	return "default"
}
func Replace(log string, oldRune, newRune rune) string {
	result := []rune(log)
	for i, r := range result {
		if r == oldRune {
			result[i] = newRune
		}
	}
	return string(result)
}
func WithinLimit(log string, limit int) bool {
	return len([]rune(log)) <= limit
}

