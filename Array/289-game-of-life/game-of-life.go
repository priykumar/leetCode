func gameOfLife(board [][]int)  {
    m, n := len(board), len(board[0])

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            liveCount:=0
            for ii:=-1; ii<=1; ii++ {
                for jj:=-1; jj<=1; jj++ {
                    if ii==0 && jj==0 {
                        continue
                    }

                    if i+ii>=0 && i+ii<m && j+jj>=0 && j+jj<n && board[i+ii][j+jj]>=1{
                        liveCount++
                    }
                }
            }

            // if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1] >= 1 {
            //     liveCount++
            // }
            // if i-1 >= 0 && board[i-1][j] >= 1 {
            //     liveCount++
            // }
            // if i-1 >= 0 && j+1 < n && board[i-1][j+1] >= 1 {
            //     liveCount++
            // }
            // if j+1 < n && board[i][j+1] >= 1 {
            //     liveCount++
            // }
            // if i+1 < m && j+1 < n &&board[i+1][j+1] >= 1 {
            //     liveCount++
            // }
            // if i+1 < m && board[i+1][j] >= 1 {
            //     liveCount++
            // }
            // if i+1 < m && j-1 >=0 && board[i+1][j-1] >= 1 {
            //     liveCount++
            // }
            // if j-1 >=0 && board[i][j-1] >= 1 {
            //     liveCount++
            // }

            if board[i][j] == 1 {
                // live -> dead
                if liveCount < 2 || liveCount > 3 {
                    board[i][j]=2
                }
            } else {
                // dead -> live
                if liveCount == 3 {
                    board[i][j]=-1
                }
            }
        }
    }

    for i:=0; i<m; i++ {
        for j:=0; j<n; j++ {
            if board[i][j] > 1 {
                board[i][j] = 0
            } else if board[i][j] < 0 {
                board[i][j] = 1
            }
        }
    }
}

// >1 : earlier live: live -> live(1), live -> dead(2)
// <1 : earlier dead: dead -> live(-1), dead -> dead(0)