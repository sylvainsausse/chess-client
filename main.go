package main

import (
	"log"
	"math/rand"
	"net"
	"os"

	manplay "github.com/sylvainsausse/chess-client/manualPlay"
	chess "github.com/sylvainsausse/chess-engine"
)

func playRandom(c chess.Chessboard,t chess.Team) (byte,byte,byte,byte){
	list := c.GetAllPlays(t)
	chosen := list[rand.Int()%len(list)]
	return byte(chosen[0][0]),byte(chosen[0][1]),byte(chosen[1][0]),byte(chosen[1][1])
}

func main(){
	conn,err := net.Dial("tcp","localhost:3000")
	if err != nil {
		log.Fatal(err.Error())
	}
	buff := make([]byte,1024)
	n,err := conn.Read(buff)
	if err != nil || n != 1{
		log.Fatalln("Error :",err.Error())
	}
	t := chess.Team(buff[0]==1)
	println(t)
	println("---------")
	for true {
		var c chess.Chessboard
		n,err := conn.Read(buff)
		if err != nil || n != 64{
			if n == 1 && buff[0] == 0x11 {
				os.Exit(0)
			}
			println(n)
			panic("Error : wrong bits")
		}
		c.LoadFromBytes(buff)
		n,err = conn.Read(buff)
		if err != nil || n != 1{
		 	panic("Error :"+err.Error())
		}
		if buff[0] == 0xff {
			buff[0] = 0x00
			for buff[0] != 0xff {
				//c.Disp()
				l1,c1,l2,c2 := manplay.ManualPlay(c)
				ibuf := []byte{l1,c1,l2,c2}
				conn.Write(ibuf)
				conn.Read(buff)
			}
		}else if buff[0] == 0x11 {
			os.Exit(0)
		}
	}
}