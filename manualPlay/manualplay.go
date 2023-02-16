package manualplay

import (
	"fmt"

	chess "github.com/sylvainsausse/chess-engine"
)

func ManualPlay(c chess.Chessboard) (byte,byte,byte,byte){
	c.Disp()
	var a,b string
	_,err := fmt.Scanf("%s %s\n",&a,&b)

	if err != nil {
		panic(err.Error())
	}

	l1,c1 := chess.PosToCord(a)
	l2,c2 := chess.PosToCord(b)
	fmt.Println(a,b,l1,c1,l2,c2)
	return byte(l1),byte(c1),byte(l2),byte(c2)
}