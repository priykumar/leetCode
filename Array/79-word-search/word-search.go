import "fmt"
var found bool 

func backtrack(board [][]byte, word string, wordIdx, i, j int) bool {
    // word found
    if wordIdx == len(word) || found == true {
        found = true
        return true
    }

    // indices check
    if i<0 || i>=len(board) || j<0 || j>=len(board[0]) || board[i][j] != word[wordIdx] {
        return false
    }

    // visited
    if board[i][j] == '#' {
        return false
    }
    
    // mark visited
    val := board[i][j]
    board[i][j] = '#'

    x:=backtrack(board, word, wordIdx+1, i+1, j) || backtrack(board, word, wordIdx+1, i, j+1) || backtrack(board, word, wordIdx+1, i-1, j) || backtrack(board, word, wordIdx+1, i, j-1)

    board[i][j] = val

    return x
}

func exist(board [][]byte, word string) bool {
    found = false
    for i:=0; i<len(board); i++ {
        for j:=0; j<len(board[0]); j++ {
            if backtrack(board, word, 0, i, j) {
                return true
            }
        }
    }
    return false
}