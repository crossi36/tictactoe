package tictactoe

func mapValue(x int) string {
	switch {
	case x == 0:
		return "X"
	case x == 1:
		return "O"
	default:
		return " "
	}
}

func checkFor(value int, items ...int) bool {
	for _, item := range items {
		if value == item {
			return true
		}
	}
	return false
}
