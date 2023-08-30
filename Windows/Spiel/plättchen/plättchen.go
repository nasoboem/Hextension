package plättchen

type Plättchen interface {
	//Vor.:-
	//Erg.: Ein neues Plättch ist gegeben.
	//New(zahlLink,zahlOben,zahlRechts uint8) Plättchen
	
	
	//Vor.:-
	//Eff.:Das Plättchen ist auf der Konsole dargestellt.
	String () string
	
	//Vor.: - 
	//Erg.: Die Werte des Plättchen sind geliefert in der Reihenfolge Links, Oben, Rechts.
	GetValues () (uint8, uint8, uint8)
	
	//Vor.: Ein offenes gfx-Fenster. Das Plättchen muss vollständig im Fenster enthalten sein. (x-Koordinate muss min Abstand von Größe des Sechseck besitzen 50*resolution/10, y-Koordinate muss min Abstand von 50*(433/500)*resolution/10
	//Eff.: Das Plättchen ist im gfx-Fenster dargestellt 
	Draw(resolution uint16,x,y uint16)
}

	//Vor.: -
	//Erg.: True ist geliefert, genau dann wenn die beiden Plättchen in allen Werten (l,o,r) jeweils übereinstimmen 
func IstGleich(p1,p2 Plättchen) bool {
	var l1,l2,o1,o2,r1,r2 uint8
	l1,o1,r1 = p1.GetValues()
	l2,o2,r2 = p2.GetValues()
	return l1==l2&&o1==o2&&r1==r2
}

	//Vor.: -
	//Erg.: Ein Slice mit alle für das Spiel benötigten Plättchen (27 verschiedene) ist geliefert.

func PlättchenGenerator () [27]Plättchen {  //erzeugt eine Liste von Plättchen, wie sie zu beginn des Spiels aussieht
var f [27]Plättchen
f[0] = New(2,1,3)
f[1] = New(2,1,4)
f[2] = New(2,1,8)
f[3] = New(2,5,3)
f[4] = New(2,5,4)
f[5] = New(2,5,8)
f[6] = New(2,9,3)
f[7] = New(2,9,4)
f[8] = New(2,9,8)
f[9] = New(6,1,3)
f[10] = New(6,1,4)
f[11] = New(6,1,8)
f[12] = New(6,5,3)
f[13] = New(6,5,4)
f[14] = New(6,5,8)
f[15] = New(6,9,3)
f[16] = New(6,9,4)
f[17] = New(6,9,8)
f[18] = New(7,1,3)
f[19] = New(7,1,4)
f[20] = New(7,1,8)
f[21] = New(7,5,3)
f[22] = New(7,5,4)
f[23] = New(7,5,8)
f[24] = New(7,9,3)
f[25] = New(7,9,4)
f[26] = New(7,9,8)
return f
}


