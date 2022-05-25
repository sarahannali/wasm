package greetings

import "fmt"

func filter(ss [9]int, test func(int) bool) (ret []int) {
	for _, s := range ss {
			if test(s) {
					ret = append(ret, s)
			}
	}
	return
}

func validate(board [9][9]int) bool {
	fmt.Println("HAVE BOARD: ", board)
	filterEmptyCells := func(cell int) bool { return cell != 0 }
	s2 := filter(board[0], filterEmptyCells)
	fmt.Println("FILTERED: ", s2)
	return true
}

func Solve(board [9][9]int) string {
    // Return a greeting that embeds the name in a message.
		validate(board)
    message := fmt.Sprintf("Hi, %v. Welcome!", board)
    return message
}