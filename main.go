package main

import (
	"fmt"
)

func main() {

	portua, komandoa := argumentuakParseatu()

	//Exekutagarria abiarazi, sarrerak eta irteerak berbideratuz
	in, out, errr := exekutagarriaAbiaratu(komandoa)
	//zerbitzatzen hasi
	Listener := zerbitzaria(portua)

	for {
		//bezeroa onartu
		fmt.Println("Zerbitzaria entzuten " + portua + " portuan...\n")
		conn, err := Listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Bezeroarekin konexioa ireki da\n")
		//bezero eta exekutagarri harteko komunikazioa kudeatu
		komunikazioaKudeatu(conn, in, out, errr)
		fmt.Println("Bezeroarekin konexioa itxi da\n")
		conn.Close()
	}

}
