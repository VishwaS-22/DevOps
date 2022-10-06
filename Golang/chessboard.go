/*
Instructions
As a chess enthusiast, you would like to write your own version of the game. 
Yes, there may be plenty of implementations of chess available online already, but yours will be unique!
Each square of the chessboard is identified by a letter-number pair.

1. Given a Chessboard and a File, count how many squares are occupied
2. Given a Chessboard and a Rank, count how many squares are occupied
3. Count how many squares are present in the given chessboard
4. Count how many squares are occupied in the given chessboard
*/

package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File[]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,within the given file.
func CountInFile(cb Chessboard, file string) int {
    var ctc int
    for _,occupied:=range cb[file]{
        if occupied{
            ctc++;
        }
    }
	return ctc;

}

// CountInRank returns how many squares are occupied in the chessboard,within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    var ctr int
    if rank>=1 && rank<=8{
        
        for _,check:=range cb{
            if check[rank-1]{
            	ctr++
            }
        }
    }
	return ctr
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var c int
    for _,p:=range cb{
        for range p{
            c++
        }
    }
	return c
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	res:=0
    for _,file:=range cb{
        for _,squares:=range file{
            if squares{
                res+=1
            }
        }
    }
	return res
}
