package feld


type Feld interface {
	
//Vor.: -
//Erg.: Ein neues Feld mit der Spalte 0 - uint8, Zeile 0 - uint8 ist geliefert.
//New()
 
//Vor.: Ein gabe Werte Spalte und Zeile sind korrekt.
//Eff.: Das Feld besitzt die neuen Werte für Spalte und Zeile.
SetzeFeld (spalte, zeile uint8)


//Vor.: -
//Erg.: Die Werte von Spalte und Zeile sind geliefert.
GibtFeld () (spalte, zeile uint8)

//Vor.: -
//Erg.: Die x- und y-Koordinaten des Feldes sind geliefert. (Mittelpunkt des Sechsecks für das das Feld steht.
GibKoordinaten(resolution uint16) (x,y uint16)
}

//Vor.: -
//Erg.: True ist geliefert, wenn das gegebene Feld mit dem untersuchten Feld in Zeile und Spalte übereinstimmt.
func IstGleich (g,f Feld) bool {
	var fspalte,fzeile,gspalte,gzeile uint8
	fspalte,fzeile = f.GibtFeld()
	gspalte,gzeile = g.GibtFeld()
	return fspalte==gspalte && fzeile==gzeile
}

//Vor.: -
//Erg.: True ist geliefert, wenn die Kombination von Spalte und Zeile auf dem Spielfeld liegt.

func IstKorrekteEingabe (spalte, zeile uint8) bool {		//Überprüft, ob die Eingabe zweier uint8-Werten für Spalte und Zeile im Wertebereich des Spielfeldes liegt
	switch spalte {
		case 1: 
			if zeile <= 3 && zeile >=1 {
				return true
			}
		case 2:
			if zeile <= 4 && zeile >=1 {
				return true
			}
		case 3: 
			if zeile <= 5 && zeile >=1 {
				return true
			}
		case 4:
			if zeile <= 4 && zeile >=1 {
				return true
			}
		case 5:
			if zeile <= 3 && zeile >=1 {
				return true
			}
		}
	return false
}

//Vor.: -
//Erg.: Ein Slice aller Felder, die das Spielfeld ausmachen ist geliefert (19 Felder)

func FeldGenerator () []Feld {
	var erg []Feld																			//Initialisierung des Ergebnisvektors
	var f11,f12,f13,f21,f22,f23,f24,f31,f32,f33,f34,f35,f41,f42,f43,f44,f51,f52,f53 Feld  	//Initialisierung der Felder 
	f11 = New()																				//Zuweisung eines neuen Feldes mit den werten 0/0
	f11.SetzeFeld(1,1)																		//Wertzuweisung 1/1
	f12 = New()
	f12.SetzeFeld(1,2)
	f13 = New()
	f13.SetzeFeld(1,3)
	f21 = New()
	f21.SetzeFeld(2,1)
	f22 = New()
	f22.SetzeFeld(2,2)
	f23 = New()
	f23.SetzeFeld(2,3)
	f24 = New()
	f24.SetzeFeld(2,4)
	f31 = New()
	f31.SetzeFeld(3,1)
	f32 = New()
	f32.SetzeFeld(3,2)
	f33 = New()
	f33.SetzeFeld(3,3)
	f34 = New()
	f34.SetzeFeld(3,4)
	f35 = New()
	f35.SetzeFeld(3,5)
	f41 = New()
	f41.SetzeFeld(4,1)
	f42 = New()
	f42.SetzeFeld(4,2)
	f43 = New()
	f43.SetzeFeld(4,3)
	f44 = New()
	f44.SetzeFeld(4,4)
	f51 = New()
	f51.SetzeFeld(5,1)
	f52 = New()
	f52.SetzeFeld(5,2)
	f53 = New()
	f53.SetzeFeld(5,3)
	erg = append(erg,f11,f12,f13,f21,f22,f23,f24,f31,f32,f33,f34,f35,f41,f42,f43,f44,f51,f52,f53)	//Alle generierten Felder werden an den leeren Ergebnisvektor angehangen
	return erg																						//Ergebnisvektor wird zurück gegeben
}


