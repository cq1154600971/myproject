package leetcode

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if dfs(i, j, 0, board, word) {
				return true
			}
		}
	}
	return false
}
func dfs(i, j, step int,  board [][]byte, word string) bool {
	if i < 0 || i > len(board) || j < 0 || j > len(board[0]) || board[i][j] != word[step] {
		return false
	}
	if step == len(word) {
		return true
	}
	return dfs(i - 1, j, step + 1, board, word) ||
		dfs(i + 1, j, step + 1, board, word) ||
		dfs(i, j - 1, step + 1, board, word) ||
		dfs(i, j + 1, step + 1, board, word)
}
