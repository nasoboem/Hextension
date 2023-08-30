package reihe

import ("../feld"
		"../plättchen")



type Reihe interface {
	
//Vor.: -
//Erg.: Eine neue Reihe mit der gegeben Orientierung (l = 0, o = 1, r = 2) ist geliefert.
// New (felder []Feld, orientierung uint8) *data { //Generiert eine neue Reihe
	
//Vor.: -
//Erg.: Eine Liste der Felder, die zu der Reihe gehören ist geliefert.
	GibFelder () []feld.Feld

//Vor.: -
//Erg.: Der aktuelle Wert der Liste ist geliefert.
	GibWert () uint8

//Vor.: -
//Erg.: True ist geliefert, genau dann wenn das Feld teil der Reihe ist.
	GehörtFeldzurReihe (f feld.Feld) bool

//Vor.: -
//Erg.: True ist geliefert, genau dann wenn die Reihe voll ist.
	IstVoll () bool

//Vor.: -
//Erg.: Die aktuelle Anzahl der in der Reihe leigenden Plättchen ist geliefert.
	AnzahlPlättcheninReihe () int

//Vor.: -
//Eff.: Ein gegebenes Plättchen wird über die Feldangabe einer Reihe zugeordnet und die Werte der Reihe je nach Orientierung entsprechend angepasst.
//Reihe von linksunten nach rechtsoben = l entspricht Orientierung = 0
//			oben nach unten			   = o entspricht Orientierung = 1
//			rechtsunten nach linksoben = r entspricht Orientierung = 2
// Wenn der Wert der Reihe 0 ist wird er auf den Wert gesetzt, der zu dem Wert der Orientierung des Plättchens entspricht.
// Ist der Wert != 0 und der Wert des Plättchens entspricht nicht dem Wert der Reihe, so wird der Wert der Reihe auf 0 gesetzt.
	SetzePlättchen (f feld.Feld, p plättchen.Plättchen)

//Vor.: -
//Erg.: Ein Slice ist geliefert, der die Koordinaten der zur Reihe gehörenden Felder als ein Feld enthält Index = 0 enthält die x-Koordinate und Index = 1 enthält die y-Koordinate
	GibKoordinaten(resolution uint16) [][2]uint16
}

//Vor.: -
//Erg.: Die Anzahl der zur Reihe gehörenden Felder ist geliefert (Wert zwichen 3 und 5)
func Größe (r Reihe) int {
	var f []feld.Feld
	f = r.GibFelder()
	return len(f)
}

//Vor.: -
//Erg.: Liefert ein Feld aus 15 Reihen, die den Reihen der Legefläche des Spielfeldes entsprechen - werden gebraucht, um die Punkte zu berechnen
func ReihenGenerator () [15]Reihe {
	//                                                                       0   1   2   3   4   5   6   7   8   9  10   11  12  13  14  15  16  17  18
	//Der FeldGenerator liefert ein Slice aus Feldern in dieser Reihenfolge f11,f12,f13,f21,f22,f23,f24,f31,f32,f33,f34,f35,f41,f42,f43,f44,f51,f52,f53
	var alleFelder []feld.Feld
	alleFelder = feld.FeldGenerator()
	var orientierung uint8 = 0										//Festlegen der Orientierung auf 0 --> l bei Plättchen
	var r [15]Reihe													//Initialisierung des Rückgabewertes als ein Feld von 15 Reihe
	var r0,r1,r2,r3,r4,r5,r6,r7,r8,r9,r10,r11,r12,r13,r14 []feld.Feld	//Initialisierung eines Slices von Felder zur Übergabe an das Rückgabefeldes, Name entspricht dem Indexwertes des Rückgabefeldes
	r0 = append(r0,alleFelder[0],alleFelder[3],alleFelder[7])										//Definierung der in der Reihe enthaltenen Felder
	r[0] = New(r0,orientierung)										//Initialisierung und übergabe an das Rückgabefeldes
	r1 = append(r1,alleFelder[1],alleFelder[4],alleFelder[8],alleFelder[12])
	r[1] = New(r1, orientierung)
	r2 = append(r2,alleFelder[2],alleFelder[5],alleFelder[9],alleFelder[13],alleFelder[16])
	r[2] = New(r2, orientierung)
	r3 = append(r3,alleFelder[6],alleFelder[10],alleFelder[14],alleFelder[17])
	r[3] = New(r3, orientierung)
	r4 = append(r4,alleFelder[11],alleFelder[15],alleFelder[18])
	r[4] = New(r4, orientierung)
	orientierung = 1												//Festlegen der Orientierung auf 1 --> o bei Plättchen
	r5 = append(r5,alleFelder[0],alleFelder[1],alleFelder[2])
	r[5] = New(r5,orientierung)
	r6 = append(r6,alleFelder[3],alleFelder[4],alleFelder[5],alleFelder[6])
	r[6] = New(r6,orientierung)
	r7 = append(r7,alleFelder[7],alleFelder[8],alleFelder[9],alleFelder[10],alleFelder[11])
	r[7] = New(r7, orientierung)
	r8 = append(r8,alleFelder[12],alleFelder[13],alleFelder[14],alleFelder[15])
	r[8] = New(r8,orientierung)
	r9 = append(r9,alleFelder[16],alleFelder[17],alleFelder[18])
	r[9] = New(r9, orientierung)
	orientierung = 2												//Festlegung der Orientierung auf 2 --> r bei Plättchen
	r10 = append(r10,alleFelder[7],alleFelder[12],alleFelder[16])
	r[10] = New(r10, orientierung)
	r11 = append(r11,alleFelder[3],alleFelder[8],alleFelder[13],alleFelder[17])
	r[11] = New(r11, orientierung)
	r12 = append(r12,alleFelder[0],alleFelder[4],alleFelder[9],alleFelder[14],alleFelder[18])
	r[12] = New(r12,orientierung)
	r13 = append(r13,alleFelder[1],alleFelder[5],alleFelder[10],alleFelder[15])
	r[13] = New(r13,orientierung)
	r14 = append(r14,alleFelder[2],alleFelder[6],alleFelder[11])
	r[14] = New(r14,orientierung)
	return r
}

//Vor.: -
//Erg.: Die Punkte einer Reihe ist geliefert.
func Punkte (r Reihe) int {			//Berechnet die Punkte einer Reihe. Bepunktet werden nur volle Reihen. Der Punktewert ergibt sich aus dem Wert der Reihe * der Größe der Reihe.
	if r.IstVoll() {
		return Größe(r) * int(r.GibWert())
	} else {
		return 0
	}
}




	
	
	

	


	
	
			


