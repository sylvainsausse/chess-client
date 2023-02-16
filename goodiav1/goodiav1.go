package goodiav1

import (
	chess "github.com/sylvainsausse/chess-engine"
)

func clone(c chess.Chessboard) chess.Chessboard {
	a := chess.NewChessboard()
	for i := 0;i<64;i++ {
		a[i] = c[i]
	}
	return a
}

func PlayGoodIAv1Calc(c chess.Chessboard,t chess.Team, index int, playteam int) (int,byte,byte,byte,byte) {
	if index == 0 {
		return 0,0,0,0,0
	}
	list := c.GetAllPlays(t)

	minoumax := -playteam*100000000
	var move [2][2]int
	translate := [...]int{0,100000,9,5,3,3,1,0,0,-100000,-9,-5,-3,-3,-1}

	for _,item := range list {
		b := clone(c)
		b.Make_move(t,item[0][0],item[0][1],item[1][0],item[1][1])
		valuenext,_,_,_,_ := PlayGoodIAv1Calc(c,t,index-1,-playteam)
		/*if valuenext*(-playteam) > 10 {
			continue
		}*/
		taken := chess.Piece(byte(c.Sum() - b.Sum()))
		var value int
		if taken < 0 {
			value = valuenext + index*translate[int(-taken)]
		}
		if taken > 14 {
			value = valuenext + index*int(taken)
		}
		value = valuenext + index*translate[int(taken)]
		if playteam > 0 && value > minoumax {
			minoumax = value
			for i,j := range item {
				for k,l := range j {
					move[i][k] = l
				}
			}
		}
		if playteam < 0 && value < minoumax {
			minoumax = value
			for i,j := range item {
				for k,l := range j {
					move[i][k] = l
				}
			}
		}
	}
	return minoumax,byte(move[0][0]),byte(move[0][1]),byte(move[1][0]),byte(move[1][1])
}

	

func PlayGoodIAv1(c chess.Chessboard,t chess.Team, index int) (byte,byte,byte,byte) {
	_,l1,c1,l2,c2 := PlayGoodIAv1Calc(c,t,index,1)
	return l1,c1,l2,c2
}