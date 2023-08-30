package feld

import ("fmt")

type data struct {  //Ein Feld ist definiert durch die Angabe der Koordinaten des Feldes bestehend aus Spalte und Zeile
	spalte uint8
	zeile uint8
}

func New() *data {		//Gibt ein neues Feld mit den Werten Spalte 0 und Zeile 0, was nicht einer korrekten Eingabe für ein Feld entspricht!!!
	var f *data
	f = new(data)
	return f
}

func (f *data) SetzeFeld (spalte, zeile uint8) {	//SetzeFeld ändert die Werte eines gegeben Feldes auf die angegeben Feldparameter Spalte und Zeile
		(*f).spalte = spalte
		(*f).zeile = zeile
}

func (f *data) GibtFeld () (spalte, zeile uint8) { //Gibt die Werte eines gegebenen Feldes zurück
	spalte = (*f).spalte
	zeile = (*f).zeile
	return
}






func (f *data) String () string { //Generiert einen String, der ein Feld repräsentiert aus <Spalte-Zeile> Bsp. 3-1
	var erg string
	var zeile,spalte uint8
	zeile, spalte = f.GibtFeld()
	erg = erg + fmt.Sprint(zeile)
	erg = erg + fmt.Sprint("-")
	erg = erg + fmt.Sprint(spalte)
	return erg
}

func (f *data) GibKoordinaten (resolution uint16) (x,y uint16) { //Gibt die Koordinaten aus, an denen die Felder liegen
	var zeile,spalte uint8
	zeile, spalte = f.GibtFeld()
	switch  {												//Switchstatment abhängig von Zeile und Spalte (insgesamt 19)
			case spalte==3 && zeile==1:
				x = 390*resolution/10
				y = 100*resolution/10
				return
			case spalte==2 && zeile==1:
				x = 302*resolution/10
				y = 150*resolution/10
				return
			case spalte==4 && zeile==1:
				x = 478*resolution/10
				y = 150*resolution/10
				return
			case spalte==1 && zeile==1:
				x = 214*resolution/10
				y = 200*resolution/10
				return
			case spalte==3 && zeile==2:
				x = 390*resolution/10
				y = 200*resolution/10
				return
			case spalte==5 && zeile==1:
				x = 566*resolution/10
				y = 200*resolution/10
				return
			case spalte==2 && zeile==2:
				x = 302*resolution/10
				y = 250*resolution/10
				return
			case spalte==4 && zeile==2:
				x = 478*resolution/10
				y = 250*resolution/10
				return
			case spalte==1 && zeile==2:
				x = 214*resolution/10
				y = 300*resolution/10
				return
			case spalte==3 && zeile==3:
				x = 390*resolution/10
				y = 300*resolution/10
				return
			case spalte==5 && zeile==2:
				x = 566*resolution/10
				y = 300*resolution/10
				return
			case spalte==2 && zeile==3:
				x = 302*resolution/10
				y = 350*resolution/10
				return
			case spalte==4 && zeile==3:
				x = 478*resolution/10
				y = 350*resolution/10
				return
			case spalte==1 && zeile==3:
				x = 214*resolution/10
				y = 400*resolution/10
				return
			case spalte==3 && zeile==4:
				x = 390*resolution/10
				y = 400*resolution/10
				return
			case spalte==5 && zeile==3:
				x = 566*resolution/10
				y = 400*resolution/10
				return
			case spalte==2 && zeile==4:
				x = 302*resolution/10
				y = 450*resolution/10
				return
			case spalte==4 && zeile==4:
				x = 478*resolution/10
				y = 450*resolution/10
				return
			case spalte==3 && zeile==5:
				x = 390*resolution/10
				y = 500*resolution/10
				return
		}
		return
}
		

