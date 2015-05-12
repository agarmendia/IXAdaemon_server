package main

import (
	"fmt"
)

func main() {
	//Exekutagarria abiarazi, sarrerak eta irteerak berbideratuz
	in, out, errr := exekutagarriaAbiaratu()
	//zerbitzatzen hasi
	Listener := zerbitzaria()

	fmt.Println("Prozesua martxan da \n")

	for {
		//bezeroa onartu
		conn, err := Listener.Accept()
		if err != nil {
			continue
		}
		//bezero eta exekutagarri harteko komunikazioa kudeatu
		komunikazioaKudeatu(conn, in, out, errr)
		conn.Close()
	}

}
