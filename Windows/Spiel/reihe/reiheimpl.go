package reihe

import ("../feld"
		"../plättchen"
		//"../gfx"
		)

type data struct {
	felder []feld.Feld	//Der Slice enthält die Felder, die zur Reihe gehören
	besetzt bool		//Merkt sich ob bereits Plättchen in der Reihe liegen 
	wert uint8			//Werte der Plättchen - 0-9
	orientierung uint8 	// 0 = l, 1 = o, 2 = r
	zähler int			//Enthält die Anzahl der in der Reihe abgelegten Plättchen
	voll bool			//Wirdt auf true gesetzt, dann, wenn zähler == len(felder) ist
}

func New (felder []feld.Feld, orientierung uint8) *data { //Generiert eine neue Reihe aus einem Slice von Feldern und der Angabe der Orientierung
	var reihe *data
	reihe = new(data)
	(*reihe).felder = felder								//Die Felder werden übergeben
	(*reihe).besetzt = false								//Es liegt noch kein Plättchen in der Reihe
	(*reihe).wert = 0										//Die Reihe besitzt den Wert 0 entweder wenn sie leer ist (hier der Fall) oder wenn verschiedenwertige Plättchen sich in ihr befinden
	(*reihe).orientierung = orientierung					//Setzt die Orientierung
	(*reihe).zähler = 0										//Keine Plättchen enthalten --> Zähler = 0
	(*reihe).voll = false									//Reihe ist nicht voll
	return reihe
}

func (r *data) GibFelder () []feld.Feld {                   //Gibt die zur Reihe gehörenden Felder zurück
	return (*r).felder
}

func (r *data) GibWert () uint8 {							//Gibt den aktuellen Wert der Reihe zurück
	return (*r).wert
}

func (r *data) GehörtFeldzurReihe (f feld.Feld) bool { 		// Tested ob ein Feld zu der Reihe gehört, wenn das Feld in der Reihe enthalten ist, dann wird true geliefert, sonst false (jedes Feld gehört immer zu drei Reihen für die oreintierungen l,o und r) 
	for i:=0;i<len((*r).felder);i++{
		if feld.IstGleich((*r).felder[i],f) {				//Vergleicht das übergebene Feld mit allen Feldern einer Reihe 
			return true
		}
	}
	return false
}

func (r *data) IstVoll () bool {							//Prüft, ob bei der Reihe das Voll-Flag gesetzt ist.
	return (*r).voll
}

func (r *data) AnzahlPlättcheninReihe () int {				//Gibt den Zählerstand zurück und damit die Anzahl der auf dieser Reihe befindlichen Plättchen
	return (*r).zähler
}

func (r *data) SetzePlättchen (f feld.Feld, p plättchen.Plättchen) {  //Setzt und verwaltet die Werte der Reihen, je nach dem welchen Wert die Plättchen haben und auf welches Feld es gelegt wird
	if r.GehörtFeldzurReihe(f) {						//Es wird geprüft, ob das Feld auf dass das Plättchen gelegt wird zur Reihe gehört
		(*r).zähler++									//Wenn ja, dann wird der Zähler der Reihe hoch gestezt
		if (*r).zähler == len((*r).felder) {			//Es wird geschaut, ob durch das erhöhen des Zählers die Reihe voll ist, und wenn ja
			(*r).voll = true							//Wird das voll Flag gesetzt
		}
		var ll, oo, rr uint8 = p.GetValues()			//Werte der Richtungen des Plättchen werden ausgelesen
		var orientierung uint8 = (*r).orientierung		//Orientierung der Reie wird ausgelesen
		switch orientierung {
			case 0:						//Reihe hat die Orientierung l = 0:
				if !(*r).besetzt {		//Wenn die Reihe leer ist (1.Plättchen, das in die Reihe gelegt wird
					(*r).wert = ll		//Setzt den Wert der Reihe auf den Wert des Plättchens, der an l steht
					(*r).besetzt = true	//die Reihe ist jetzt besetzt
				} else {				//Wenn die Reihe bereits besetzt ist ...
					if ll!=(*r).wert && (*r).wert != 0 {		// testet, ob das neue Plättchen den gleichen Wert hat, oder bereits 0 ist
						(*r).wert = 0						//wenn dem nicht so ist, dann wird der Wert der Reihe auf 0 gesetzt.
					}
				}
			case 1:						//Reihe hat die Orientierung o = 1:
				if !(*r).besetzt {
					(*r).wert = oo
					(*r).besetzt = true
				} else {
					if oo!=(*r).wert && (*r).wert != 0 {
						(*r).wert = 0
					}
				}
			case 2:						//Reihe hat die Orientierung r = 2:
				if !(*r).besetzt {
					(*r).wert = rr
					(*r).besetzt = true
				} else {
					if rr!=(*r).wert && (*r).wert != 0 {
						(*r).wert = 0
					}
				}
		}
	}
}


//Die Zeichenfunktion war ursprünglich angedacht, um ein Reihenhighlighting zu ermöglichen, die dem Spieler die Auswahl der besten Feldes für das aktuellzuspielende Plättchen ermöglichen würde - nicht realisiert, aber mögliche Erweiterung
/*
func (r *data) Draw() { 
	r.farbWahl()
	
}

	
func (r *data) farbWahl () {
	var f uint8
	f = (*r).wert
	switch f {
		case 2:
		gfx.Stiftfarbe (255,0,255)
		case 6:
		gfx.Stiftfarbe (255,0,0)
		case 7:
		gfx.Stiftfarbe (0,160,0)
		case 3:
		gfx.Stiftfarbe (0,255,255)
		case 4:
		gfx.Stiftfarbe (0,0,255)
		case 8:
		gfx.Stiftfarbe (255,160,0)
		case 1:
		gfx.Stiftfarbe (100,100,100)
		case 5:
		gfx.Stiftfarbe (160,0,160)
		case 9:
		gfx.Stiftfarbe (255,255,0)
		default:
		gfx.Stiftfarbe (255,255,255)
	}
}
*/


func (r *data) GibKoordinaten (resolution uint16) [][2]uint16 { //Liefert die Kooerdinaten der zur Reihe gehörenden Felder - ist auch eine Funktion, die nur für das Reihenhighlighting benötigt wird.
	var erg [][2]uint16
	for i:=0;i<len((*r).felder);i++ {
		var x,y uint16
		x,y = (*r).felder[i].GibKoordinaten(resolution)
		var k [2]uint16
		k[0] = x
		k[1] = y
		erg = append (erg,k)
	}
	return erg
}


		

