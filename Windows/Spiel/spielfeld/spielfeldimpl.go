package spielfeld

import ("../plättchen"
		"fmt"
		"../feld"
		"../reihe"
		"gfx2"
		"strconv"
		"../seite"
		"../button"
		)

type data struct {									
	s1,s2 seite.Seite							//Seiten für das Buttonhandling - s1: Ansicht vor der Feldwahl des Spielers; s2: Ansicht nach der Feldwahl des Spielers
	aktuellesFeld feld.Feld						//Speicher die Feldwahl des Spielers
	aktuellesPlättchen plättchen.Plättchen		//Speichert das aktuell zu spielende Plättchen
	belegteFelder []feld.Feld					//Speichert welche Felder bereits belegt sind
	gesetztePlättchen [19]plättchen.Plättchen	//Speichert die gelegten Plättchen
	plättchenImSpiel [27]plättchen.Plättchen    //Hier sind die Plättchen abgelegt, die sich noch im Spiel befinden (zu Beginn alle)
	zug uint8									//Hier ist die Zugnummer gespeichert (Da die Zugnummer auch als Index für Slices und Felder verwendet wird beginnt sie bei Null; für die Darstellung wird 1 addiert
	werteReihe [15]reihe.Reihe					//Hier werden die Reihen des Spielfeldes verwaltet, was einen live Punkteauswertung ermöglicht
	punkte int									//Speichert den Punktestand
	spielername string							//Speichert den Spielernamen, der it diesem Spielfeld spielt
}


func New (resolution uint16,spielername string) *data {		//Gibt ein Spielfeld mit Nullinitialisierten feldern zurück in der angegeben Auflösung (160*resolutionx100*resolution)
	var sf *data											//Es wird Variable sf (Spielfeld) initialisisert
	sf = new(data)											//Dieser Variablen wird jetzt eine Datenstruktur vom obengezeigten Typ zugewiesen
	(*sf).s1 = seite.New("Spielfeld")						//Legt die erste Seite an; Spieler hat noch kein feld gewählt für das aktuelle Plättchen
	(*sf).s1.AddButton(resolution,"3-1",390,100,"",0,0)		//Legt die Legefläche und die weiteren Buttons auf der Seite s1 an
	(*sf).s1.AddButton(resolution,"2-1",302,150,"",0,0)
	(*sf).s1.AddButton(resolution,"4-1",478,150,"",0,0)
	(*sf).s1.AddButton(resolution,"1-1",214,200,"",0,0)
	(*sf).s1.AddButton(resolution,"3-2",390,200,"",0,0)
	(*sf).s1.AddButton(resolution,"5-1",566,200,"",0,0)
	(*sf).s1.AddButton(resolution,"2-2",302,250,"",0,0)
	(*sf).s1.AddButton(resolution,"4-2",478,250,"",0,0)
	(*sf).s1.AddButton(resolution,"1-2",214,300,"",0,0)
	(*sf).s1.AddButton(resolution,"3-3",390,300,"",0,0)
	(*sf).s1.AddButton(resolution,"5-2",566,300,"",0,0)
	(*sf).s1.AddButton(resolution,"2-3",302,350,"",0,0)
	(*sf).s1.AddButton(resolution,"4-3",478,350,"",0,0)
	(*sf).s1.AddButton(resolution,"1-3",214,400,"",0,0)
	(*sf).s1.AddButton(resolution,"3-4",390,400,"",0,0)
	(*sf).s1.AddButton(resolution,"5-3",566,400,"",0,0)
	(*sf).s1.AddButton(resolution,"2-4",302,450,"",0,0)
	(*sf).s1.AddButton(resolution,"4-4",478,450,"",0,0)
	(*sf).s1.AddButton(resolution,"3-5",390,500,"",0,0)
	(*sf).s1.AddButton(resolution,"Beenden2",1450,700,"Beenden",-40,-10) //"normale Buttons"
	(*sf).s1.AddButton(resolution,"Konsole",1362,650,"Konsole",-40,-10)
	(*sf).s1.AddButton(resolution,"Menü",1274,700,"Menü",-24,-10)

	(*sf).s2 = seite.New("Spielfeld")							//Legt die zweite Seite an; hier hat der Spieler ein Feld für das aktuelle Plättchen gewählt
	(*sf).s2.AddButton(resolution,"3-1",390,100,"",0,0)			//Legt die Legefläche und die weiteren Buttons auf der Seite s2 an
	(*sf).s2.AddButton(resolution,"2-1",302,150,"",0,0)
	(*sf).s2.AddButton(resolution,"4-1",478,150,"",0,0)
	(*sf).s2.AddButton(resolution,"1-1",214,200,"",0,0)
	(*sf).s2.AddButton(resolution,"3-2",390,200,"",0,0)
	(*sf).s2.AddButton(resolution,"5-1",566,200,"",0,0)
	(*sf).s2.AddButton(resolution,"2-2",302,250,"",0,0)
	(*sf).s2.AddButton(resolution,"4-2",478,250,"",0,0)
	(*sf).s2.AddButton(resolution,"1-2",214,300,"",0,0)
	(*sf).s2.AddButton(resolution,"3-3",390,300,"",0,0)
	(*sf).s2.AddButton(resolution,"5-2",566,300,"",0,0)
	(*sf).s2.AddButton(resolution,"2-3",302,350,"",0,0)
	(*sf).s2.AddButton(resolution,"4-3",478,350,"",0,0)
	(*sf).s2.AddButton(resolution,"1-3",214,400,"",0,0)
	(*sf).s2.AddButton(resolution,"3-4",390,400,"",0,0)
	(*sf).s2.AddButton(resolution,"5-3",566,400,"",0,0)
	(*sf).s2.AddButton(resolution,"2-4",302,450,"",0,0)
	(*sf).s2.AddButton(resolution,"4-4",478,450,"",0,0)
	(*sf).s2.AddButton(resolution,"3-5",390,500,"",0,0)
	(*sf).s2.AddButton(resolution,"Weiter",755,500,"Weiter",-34,-10) //"normaler" Button

	(*sf).aktuellesFeld = feld.New()						//Ein neues aktuelle Feld ist erstellt mit den Werten Spalte 0 und Zeile 0
	(*sf).aktuellesPlättchen = plättchen.New(0,0,0)			//Ein neues aktuelles Plättchen mit den Werten (0,0,0) ist gegeben.
	for i:=0;i<len((*sf).gesetztePlättchen);i++ {			//Das Feld gesetzte Plättchen wird mit Nullinitialisierten Plättchen belegt
		(*sf).gesetztePlättchen[i] = plättchen.New(0,0,0)
	} 
	(*sf).plättchenImSpiel = plättchen.PlättchenGenerator ()	//Erstellt eine Liste aller zu Beginn des Spieles vorhandener Plättchen
	(*sf).werteReihe = reihe.ReihenGenerator ()					//Erstellt eine Liste aller für die Punkteberechnung notwendigen Reihen
																//Zug ist null initialisiert
																//Punkte ist null initialisiert
	(*sf).spielername = spielername								//Trägt den übergebenen Spielername in das Spielfeld ein
	return sf													//Rückgabe des Spielfeldes
}

//Funktionen für das Spiel nicht verwendet -------------------------------------------------------------------------------

func (sf *data) SetzeSeite1 (s seite.Seite) {				//Ermöglicht es einem Spielfeld eine neue Seite 1 zu zuweisen
	(*sf).s1 = s
}

func (sf *data) SetzeSeite2 (s seite.Seite) {				//Ermöglichte es einem Spielfeld eine neue Seite 2 zu zuweisen
	(*sf).s2 = s
}

func (sf *data) GibSeite1 () seite.Seite {					//Gibt die Seite 1 eines Spielfeldes zurück
	return (*sf).s1
}

func (sf *data) GibSeite2 () seite.Seite {					//Gibt die Seite 2 eines Spielfeldes zurück
	return (*sf).s2
}

//------------------------------------------------------------------------------------------------------------------------

func (sf *data) FeldFrei (f feld.Feld) bool {					//Tested, ob ein gegebenes Feld auf dem Spielfeld frei ist
	if !feld.IstKorrekteEingabe(f.GibtFeld()) {					//Abfangen des Startzustandes, wenn der Spieler noch kein Feld gewählt hat
		return false
	}
	for i:=0;i<len((*sf).belegteFelder); i++ {					//Geht die Liste der belegten Felder(Koordinaten) durch und wenn das Feld in der Liste bereits enthalten ist gibt es ein false zurück
		if feld.IstGleich(f,(*sf).belegteFelder[i]) || feld.IstGleich(f,(*sf).aktuellesFeld) { //tested, ob das Feld bereits besetzt ist oder gerade schon ausgewählt wurde.
			return false
		}
	}
	return true													//Wenn das Feld noch nicht im Slice enthalten ist, dann gibt es ein True zurück
}

func (sf *data) SetzePunkte (punkte int) {						//Trägt ein gegeben Punkte Wert auf dem Spielfeld ein -- nicht verwendet
	(*sf).punkte = punkte
}

//Funktionen für die Brut-Force-Berechnung der Punkte; Berechnung war leider zu langsam ---------------------------------------------

func (sf *data) FreiesLegen (f feld.Feld, p plättchen.Plättchen) {				//Erlaubt es beliebige Plättchen aufs Spielfeld zulegen, solange das Feld noch frei ist
	if (*sf).FeldFrei(f) {														//Tested ob das übergeben Feld frei ist
		(*sf).aktuellesFeld = f													//das aktuellen Feld wird auf das übergebene Feld gesetzt
		(*sf).aktuellesPlättchen = p											//das aktuelle Plättchen wird auf das übergeben Plättchen gesetzt
		(*sf).belegteFelder = append((*sf).belegteFelder, f)					//das aktuelle Feld wird an die belegten Felder angehangen
		sf.reiheEintragen()														//Die Werte der Reihen werden nach dem hinzugefügten Plättchen angepasst
		(*sf).punkte = sf.punkteberechnen()										//Neuen Punktestand berechnen und auf dem Spielfeld eintragen
		(*sf).aktuellesFeld = feld.New()										//Das aktuelle Feld leeren (mit einem Nullerfeld besetzen)
		(*sf).gesetztePlättchen[(*sf).zug],(*sf).aktuellesPlättchen = (*sf).aktuellesPlättchen,(*sf).gesetztePlättchen[(*sf).zug] 	//Das aktuelle Plättchen durch ein Nullerplättchen tauschen 
																																	//- macht zwei Dinge, legt das entsprechende Plättchen an der 
																																	//richtigen Stelle in die gesetzten Plättchen ab und leert das
																																	//aktuelle Plättchen.
		(*sf).zug++ 															//Erhöhung der Zugnummer
	} 
}
	
func (sf *data) LetztesPlättchenEntfernen () (p plättchen.Plättchen) { 				//Entfernt das letzte Plättchen vom Spielfeld und gibt es zurück
	p = plättchen.New(0,0,0)
	if (*sf).zug>0 {																//Abfangen des Falles, dass das Spielfeld bereits leer ist
		(*sf).zug--																	//Zugnummer wird reduziert - damit Zeigt der Index auf das letzte Feld und letzte Plättchen 
		p,(*sf).gesetztePlättchen[(*sf).zug] = (*sf).gesetztePlättchen[(*sf).zug],p //Das erzeugte Nullerplättchen wird in die gesetztenPlätten hineingetauscht und dem Ausgabe-Plättchen übergeben
		(*sf).belegteFelder = (*sf).belegteFelder[:len((*sf).belegteFelder)-1]		//Der Slice der belegten Felder wird um eins kleiner gemacht
		sf.reihenNeuEintrage()														//Alle Reihen müssen neu eingetragen werden - hatte keine lust eine eigene Reihenentfernungsfunktion zu schreiben
		(*sf).punkte = sf.punkteberechnen()											//Punkte des Spielfeldes werden neu berechnet und eingetragen
	}
	return p																		//Ausgabe-Plättchen wird zurück gegeben
}


func (sf *data) PlättchenAusdemSpielnehmen (game [19]uint8) {		//"Entfernt" - ersetzt sie mit Nullerplättchen - alle im Spiel verwendeten Plättchen aus plättchenImSpiel - ist nur für die Optik! 
	for i:=0;i<len(game);i++ {										//Geht das Spiel durch 
		var np plättchen.Plättchen									//und ersetzt 
		np = plättchen.New(0,0,0)									//bei den im Spielfeld gespeicherten
		(*sf).plättchenImSpiel[game[i]] = np						//Plättchen die durch das Spiel vorgegeben Plättchen mit Nullerplättchen.
	}
}

//-------------------------------------------------------------------------------------------------------------------------------------

func (sf *data) PlättchenSetzen (spiel [19]uint8) { 	                                                                      					//Schreibt die aktuellen Werte für Plättchen und Feld in die entsprechenden Listen und setzt aktuelles Feld auf (0,0) und läd das nächste Plättchen			
	(*sf).belegteFelder = append((*sf).belegteFelder, (*sf).aktuellesFeld)																		//Hängt das aktuelle Feld an das Slice belegteFelder an
	sf.reiheEintragen()																															//Trägt die Werte aus den Plättchen entsprechend der Felder in die Reihen zur Punkte berechnung ein.
	(*sf).punkte = sf.punkteberechnen()																											//Der Punkte Stand wird in das Spielfeld eingetragen
	(*sf).aktuellesFeld = feld.New()																											//Generiert ein neues aktuelles Felt mit den Werten (0,0)						
	(*sf).gesetztePlättchen[(*sf).zug],(*sf).aktuellesPlättchen = (*sf).aktuellesPlättchen,(*sf).gesetztePlättchen[(*sf).zug]					//Das aktuelle Plättchen wird mit dem Nullerplättchen aus gesetztePlättchen getauscht
	(*sf).zug++																																	//Der Zug(zähler) wird eins rauf gesetzt
	if int((*sf).zug)<len(spiel) {																												//Beim letzten Zug muss kein neues Plättchen getauscht werden oder die Zug erhöht werden
		(*sf).aktuellesPlättchen,(*sf).plättchenImSpiel[spiel[(*sf).zug]] = (*sf).plättchenImSpiel[spiel[(*sf).zug]],(*sf).aktuellesPlättchen	//Das für den nächsten Zug notwendige Plättchen wird mit dem aktuellen Plättchen (0,0,0) getauscht, so dass die Liste der plättchenImSpiel langsam mit (0,0,0) Plättchen befüllt wird
	}
}


func (sf *data) PlättchenZiehen (spiel [19]uint8) {																							//Für den 1.Zug wird das aktuelle Plättchen mit dem ersten Plättchen des Spiels besetzt
	(*sf).aktuellesPlättchen,(*sf).plättchenImSpiel[spiel[(*sf).zug]] = (*sf).plättchenImSpiel[spiel[(*sf).zug]],(*sf).aktuellesPlättchen	//Das für den 1. Zug notwendige Plättchen wird mit dem aktuellen Plättchen (0,0,0) getauscht, so dass die Liste der plättchenImSpiel langsam mit (0,0,0) Plättchen befüllt wird
}


func (sf *data) FeldSetzen (f feld.Feld) {			//Setzt ein gegebenes Feld als aktuelles Feld --> durch gereicht von Spielen
	(*sf).aktuellesFeld = f
}

func (sf *data) GibaktuellesFeld () feld.Feld {		//Gibt das aktuelle Feld zurück, das bei dem Zug aktuell ist --> (0,0) wenn Spieler noch keines gewählt hat
	return (*sf).aktuellesFeld
}

func (sf *data) GibZug () uint8 {					//Gibt die aktuelle Zugnummer zurück (1-19) - Endbildschirm zeigt Zug 20
	return (*sf).zug
}

func (sf *data) GibPunkte () int {					//Gibt die im Spielfeld gespeicherten Punkte zurück
	return (*sf).punkte
}


func (sf *data) GibaktuellesPlättchen () plättchen.Plättchen {  //Gibt das aktuelle Plättchen zurück
	return (*sf).aktuellesPlättchen
}

func (sf *data) punkteberechnen () int {							//Berechnet die aktuellen Punkte
	var aktpunkte int												//Initialisierung eines Akkus für den aktuellen Punktestand
	for i:=0;i<len((*sf).werteReihe);i++ {							//Durchläuft alle Reihen und addiert die Punkte der einzelnen Reihen zum aktuellen Punktestand dazu
		aktpunkte = aktpunkte + reihe.Punkte((*sf).werteReihe[i])
	}
	return aktpunkte												//Gibt den Punktestand zurück
}

func (sf *data) GibSpielername () string {							//Gibt den Spielernamen zurück
	return (*sf).spielername
}

func (sf *data) reiheEintragen () {															//Trägt das aktuelle Plättchen entsprechend dem aktuellen Feld in die Reihen ein
	for i:=0;i<len((*sf).werteReihe);i++ {													//Geht alle Reihen durch und ändert nach dem aktuellenFeld und aktuellem Plättchen die Werte der Reihen
		(*sf).werteReihe[i].SetzePlättchen((*sf).aktuellesFeld,(*sf).aktuellesPlättchen)
	}
}
	
func (sf *data) reihenNeuEintrage () {         									//Trägt alle bisher gespieleten Plättchen (gesetzePlättchen) und nach den belegten Feldern ein (Ein neu Eintragung - wurde für den Brut-Force-Allgorythmus verwendet)
	(*sf).werteReihe = reihe.ReihenGenerator ()
	for j:=0;j<len((*sf).belegteFelder);j++{									//Geht alle belegten Felder und die dazu gehörenden gelegten Plättchen durch
			for i:=0;i<len((*sf).werteReihe);i++ {								//Geht alle Reihen durch
				(*sf).werteReihe[i].SetzePlättchen((*sf).belegteFelder[j],(*sf).gesetztePlättchen[j])
		}
	}
}


//Die Stringfunktion für das Spielfeld greift auf die Stringfunktion von Plättchen zurück und erlaubt es 
//so alle Teile des Spielfeldes auf der Konsole auszugeben. Nach jedem Zug und nach der Wahl des Feldes 
//wird das Spielfeld in seiner gänze dargestellt. Dies ist für die Konsolenversion des Spiels gedacht.
//Festlegung des Designs für die Konsolenausgabe:
// Plättchen im Spiel:
// Die Plättchen im Spiel sind gespeicher in der Variablen plättchenImSpiel als eine Feld von 27 Plättchen.
//Die darzustellenden Reihen sind [:9], [9:18] & [18:]
//Zugnummer:
//Die Zugnummer ist ein einfacher uint8-Wert, der durch strconv.Itoa umgewandelt werden kann.
//Aktuelle Feld:
//Das Aktuelle Feld, wird vom Spieler gewählt und wird als Text erneut abgefragt.
//Aktuelles Plättchen:
//Das aktuelle Plättchen ist das was gerade gespielt wird. Es sollte neben der Legefläche angezeigt werden. Die Orientierung ist hier 
//wichtig, damit der Spieler das Spiel auch verstehen kann. (lu-ro,o-u,lo-ru)
//                                                          (  2  , 9 ,  8  )
//Legefeld:
//Das Legefeld ist etwas komplexer. Da der Spieler die Felder in einer zufälligen Reihenfolge wählt und diese dann in eine Druckbare Reihenfolge gebracht werden muss.
//Dafür sollte eine kleine Funktion geschrieben werden, die die Felder in die richtige Reihenfolge bringt.
//Das Legefeld muss ein Sechseck generieren, so dass man die Felder wählen kann. Die Feldnummer sollte darüberstehen.
//Beispiel:
//                                   3-1
//                                 ( , , )
//                           2-1              4-1
//                         ( , , )          ( , , )
//                   1-1             3-2              5-1
//                 (2,9,8)         ( , , )          ( , , )
//                           2-2              4-2
//                         (7,3,8)          ( , , )
//                   1-2             3-3              5-2
//                 ( , , )         ( , , )          ( , , )
//                           2-3              4-3
//                         (7,3,5)          ( , , )   
//                   1-3             3-4              5-3
//                 (7,9,4)         ( , , )          ( , , )
//                           2-4              4-4
//                         (2,3,4)          ( , , )
//                                   3-5
//                                 ( , , )

func (sf *data) String () string {					//Erzeugt ein Ergebnis-String, der die gesamte Information des Spielfeldes anzeigt.
	var erg string									//Initialisierung des ErgebnisStrings
	var gelegtePlättchen [19]plättchen.Plättchen	//Initialisierung einer Variablen, die die für die darstellung notwendige Sortierung der gelegten Plättchen entgegen nimmt
	gelegtePlättchen = sf.sortBelegteFelder()		//Sortiert die gelegten Plättchen nach der Reihenfolge der Felder wie sie im String erscheinen müssen
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprint("   Zug: ")								//Darstellung der Zugnummer
	erg = erg + fmt.Sprint(sf.GibZug()+1)
	erg = erg + fmt.Sprint("                          Punkte: ")	//Darstellung des Punktestandes
	erg = erg + fmt.Sprintln(sf.GibPunkte())
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln("                  Spielfeld")			//Überschrift fürs Spielfeld
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln("                     3-1") 					//21 Leerzeichen + Feldbezeichung
	erg = erg + fmt.Sprint("                   ")							//19 Leerzeichen (differenz = 2)
	erg = erg + fmt.Sprintln(gelegtePlättchen[0])							//Ausgabe des an 3-1 gelegten Plättchen
	erg = erg + fmt.Sprintln("             2-1              4-1")        	//13 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung
	erg = erg + fmt.Sprint("           ")									//11 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[1])								//Ausgabe des an 2-1 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 Leerzeichen (14 - 2 - 2)
	erg = erg + fmt.Sprintln(gelegtePlättchen[2])							//Ausgabe des an 4-1 gelegten Plättchen
	erg = erg + fmt.Sprintln("    1-1              3-2              5-1")	//4 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung
	erg = erg + fmt.Sprint("  ")											//2 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[3])								//Ausgabe des an 1-1 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[4])								//Ausgabe des an 3-2 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 leerzeichen
	erg = erg + fmt.Sprintln(gelegtePlättchen[5])							//Ausgabe des an 5-1 gelegten Plättchen
	erg = erg + fmt.Sprintln("             2-2              4-2")        	//13 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung
	erg = erg + fmt.Sprint("           ")									//11 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[6])								//Ausgabe des an 2-2 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 Leerzeichen (14 - 2 - 2)
	erg = erg + fmt.Sprintln(gelegtePlättchen[7])							//Ausgabe des an 4-2 gelegten Plättchen
	erg = erg + fmt.Sprintln("    1-2              3-3              5-2")	//4 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung
	erg = erg + fmt.Sprint("  ")											//2 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[8])								//Ausgabe des an 1-2 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[9])								//Ausgabe des an 3-3 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 leerzeichen
	erg = erg + fmt.Sprintln(gelegtePlättchen[10])							//Ausgabe des an 5-2 gelegten Plättchen
	erg = erg + fmt.Sprintln("             2-3              4-3")        	//13 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung
	erg = erg + fmt.Sprint("           ")									//11 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[11])							//Ausgabe des an 2-3 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 Leerzeichen (14 - 2 - 2)
	erg = erg + fmt.Sprintln(gelegtePlättchen[12])							//Ausgabe des an 4-3 gelegten Plättchen
	erg = erg + fmt.Sprintln("    1-3              3-4              5-3")	//4 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung
	erg = erg + fmt.Sprint("  ")											//2 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[13])							//Ausgabe des an 1-3 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[14])							//Ausgabe des an 3-4 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 leerzeichen
	erg = erg + fmt.Sprintln(gelegtePlättchen[15])							//Ausgabe des an 5-3 gelegten Plättchen
	erg = erg + fmt.Sprintln("             2-4              4-4")        	//13 Leerzeichen + Feldbezeichnung + 14 Leerzeichen + Feldbezeichnung
	erg = erg + fmt.Sprint("           ")									//11 Leerzeichen
	erg = erg + fmt.Sprint(gelegtePlättchen[16])							//Ausgabe des an 2-4 gelegten Plättchen
	erg = erg + fmt.Sprint("          ")									//10 Leerzeichen (14 - 2 - 2)
	erg = erg + fmt.Sprintln(gelegtePlättchen[17])							//Ausgabe des an 4-4 gelegten Plättchen
	erg = erg + fmt.Sprintln("                     3-5") 					//21 Leerzeichen + Feldbezeichung
	erg = erg + fmt.Sprint("                   ")							//19 Leerzeichen (differenz = 2)
	erg = erg + fmt.Sprintln(gelegtePlättchen[18])							//Ausgabe des an 3-5 gelegten Plättchen
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprint("Das aktuelle Plättchen ist: ")				//Darstellung des aktuellen Plättchens 
	erg = erg + fmt.Sprint((*sf).aktuellesPlättchen)
	erg = erg + fmt.Sprint("           Spieler: ")						//Darstellung des Spielernamens
	erg = erg + fmt.Sprintln((*sf).spielername)
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln("Noch im Spiel befindliche Plättchen:") 	//Zeigt die noch im Spiel befindlichen Plättchen an bezogen auf die 2er, 6er und 7er
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln((*sf).plättchenImSpiel[:9])
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln((*sf).plättchenImSpiel[9:18])
	erg = erg + fmt.Sprintln("")
	erg = erg + fmt.Sprintln((*sf).plättchenImSpiel[18:])
	return erg
}

func (sf *data) sortBelegteFelder () [19]plättchen.Plättchen { 	//Die Stringfunktion benötigt die Ergebnisse in der entsprechenden Reihenfolge des erscheinens auf dem Spielfeld, daher muss die vom Spieler belegte
																//Variable belegteFelder auf die Darstellung in der Stringfunktion "gemapt" werden. Da die Felder auf dem Spielfeld fix sind bedarf es nur der Zuordung der 
																//gelegten Plättchen zu den jeweiligen Feldern. Diese ist gegeben durch die Indexzahl zwischen belegteFelder und gesetztePlättchen. Daher muss nur die Reihenfolge
																//der gesetztePlättchen in Abhängigkeit der belegtenFelder zugeordnet werden.
	
	var erg [19]plättchen.Plättchen								//Initialisierung des Ergebnisvektors
	for i:=0;i<len(erg);i++ {									//Der Ergebnisvektor wird mit Nullerplättchen vorbelegt, dies ist wichtig, da nur die bereits gelegten Felder neu belegt werden
		erg[i] = plättchen.New(0,0,0)
	}
	for i:=0;i<len((*sf).belegteFelder);i++ {					//For-Schleife, die durch die belegtenFelder durch geht
		var zeile,spalte uint8									//Initialisierung der Variablen Zeile und Spalte
		spalte,zeile = (*sf).belegteFelder[i].GibtFeld()		//Zeile und Spalte des aktuell Untersuchten Feldes wird an die Variablen Zeile und Spalte ausgegeben
		switch  {												//Switchstatment abhängig von Zeile und Spalte (insgesamt 19 in der Reihenfolge, die die Stringfunktion benötigt
			case spalte==3 && zeile==1:
				erg[0] = (*sf).gesetztePlättchen[i]
				
			case spalte==2 && zeile==1:
				erg[1] = (*sf).gesetztePlättchen[i]
			case spalte==4 && zeile==1:
				erg[2] = (*sf).gesetztePlättchen[i]
				
			case spalte==1 && zeile==1:
				erg[3] = (*sf).gesetztePlättchen[i]
			case spalte==3 && zeile==2:
				erg[4] = (*sf).gesetztePlättchen[i]
			case spalte==5 && zeile==1:
				erg[5] = (*sf).gesetztePlättchen[i]
				
			case spalte==2 && zeile==2:
				erg[6] = (*sf).gesetztePlättchen[i]
			case spalte==4 && zeile==2:
				erg[7] = (*sf).gesetztePlättchen[i]
				
			case spalte==1 && zeile==2:
				erg[8] = (*sf).gesetztePlättchen[i]
			case spalte==3 && zeile==3:
				erg[9] = (*sf).gesetztePlättchen[i]
			case spalte==5 && zeile==2:
				erg[10] = (*sf).gesetztePlättchen[i]
				
			case spalte==2 && zeile==3:
				erg[11] = (*sf).gesetztePlättchen[i]
			case spalte==4 && zeile==3:
				erg[12] = (*sf).gesetztePlättchen[i]
				
			case spalte==1 && zeile==3:
				erg[13] = (*sf).gesetztePlättchen[i]
			case spalte==3 && zeile==4:
				erg[14] = (*sf).gesetztePlättchen[i]
			case spalte==5 && zeile==3:
				erg[15] = (*sf).gesetztePlättchen[i]
				
			case spalte==2 && zeile==4:
				erg[16] = (*sf).gesetztePlättchen[i]
			case spalte==4 && zeile==4:
				erg[17] = (*sf).gesetztePlättchen[i]
				
			case spalte==3 && zeile==5:
				erg[18] = (*sf).gesetztePlättchen[i]
			}
		}
	return erg
}
	

func (sf *data) Draw (resolution uint16) {								//Stellt ein Spielfeld in einem gfx2-Fenster (160*resolutionx100*resolution) dar 
	var leeresFeld feld.Feld											//Es wird ein neues Nullerfeld generiert um zwischen den zwei zuständen des Spielfeldes zu unterscheiden
	leeresFeld = feld.New()												//1. Spieler hat noch kein Feld gewählt & 2. Spieler hat ein Feld gewählt
	
	if feld.IstGleich(leeresFeld,(*sf).aktuellesFeld) {					//Fall 1
		(*sf).s1.Draw(resolution)										//s1 des Spielfeldes wird angezeigt
		sf.drawZugnr(resolution)										//Zugnummer
		sf.drawPunkte(resolution)										//Punktestand
		sf.drawAktPlättchen(resolution)									//aktuelles Plättchen 
		sf.drawPlättchenimSpiel(resolution)								//die noch im Spielbefindlichen Plättchen 
		sf.drawSpielername(resolution)									//und der Spielernamen werden angezeigt
		
	} else {															//Fall 2
		var inaktB1,inaktB2,inaktB3 button.Button						//Die Menü-Button sollen inaktive sein, daher werden sie unabhängig von der Seite erzeugt und werden daher von PressButton nicht erfasst
		inaktB1 = button.New(resolution,"",1450,700,"Beenden",-40,-10)	//Erzeugung der inaktiven Buttons
		inaktB2 = button.New(resolution,"",1362,650,"Konsole",-40,-10)
		inaktB3 = button.New(resolution,"",1274,700,"Menü",-24,-10)
		(*sf).s2.Draw(resolution)										//s2 des Spielfeldes Wird dargestellt
		inaktB1.Draw(resolution)										//Die drei Buttons werden gezeichnet
		inaktB2.Draw(resolution)
		inaktB3.Draw(resolution)
		sf.drawZugnr(resolution)										//Zugnummer
		sf.drawPunkte(resolution)										//Punkte					
		sf.drawPlättchenimSpiel(resolution)								//Plättchen im Spiel
		sf.drawSpielername(resolution)									//und der Spielername werden dargestellt
		
	}
}

func (sf *data) PlättchenAnzeigen(resolution uint16) {					//Zeigt das Plättchen auf dem Feld der Legefläche an, der Spieler gewählt hat
	var f feld.Feld														//Generiert ein Feld, das das vom Spieler gewählte Feld entgegen nimmt <-- sf.GibaktuellesFeld()
	f = sf.GibaktuellesFeld()
	var ap plättchen.Plättchen											//Generiet ein Plättchen, das das aktuelle Plättchen entgegen nimmt <-- sf.GibaktuellesPlättchen() 
	ap = sf.GibaktuellesPlättchen ()
	var spalte, zeile uint8												//Um die Position zu bestimmen muss Spalte und Zeile des Feldes Ausgelesen werden
	spalte, zeile = f.GibtFeld()										
	switch  {															//Switchstatment abhängig von Zeile und Spalte (insgesamt 19) realisieren die Zuurdnung zwischen Feld
																		//und gfx2-Fenster-Koordinaten wo das Pättchen hingezeichnet werden soll
			case spalte==3 && zeile==1:
				ap.Draw(resolution,390*resolution/10,100*resolution/10)
			case spalte==2 && zeile==1:
				ap.Draw(resolution,302*resolution/10,150*resolution/10)
			case spalte==4 && zeile==1:
				ap.Draw(resolution,478*resolution/10,150*resolution/10)
			case spalte==1 && zeile==1:
				ap.Draw(resolution,214*resolution/10,200*resolution/10)
			case spalte==3 && zeile==2:
				ap.Draw(resolution,390*resolution/10,200*resolution/10)
			case spalte==5 && zeile==1:
				ap.Draw(resolution,566*resolution/10,200*resolution/10)
			case spalte==2 && zeile==2:
				ap.Draw(resolution,302*resolution/10,250*resolution/10)
			case spalte==4 && zeile==2:
				ap.Draw(resolution,478*resolution/10,250*resolution/10)
			case spalte==1 && zeile==2:
				ap.Draw(resolution,214*resolution/10,300*resolution/10)
			case spalte==3 && zeile==3:
				ap.Draw(resolution,390*resolution/10,300*resolution/10)	
			case spalte==5 && zeile==2:
				ap.Draw(resolution,566*resolution/10,300*resolution/10)
			case spalte==2 && zeile==3:
				ap.Draw(resolution,302*resolution/10,350*resolution/10)
			case spalte==4 && zeile==3:
				ap.Draw(resolution,478*resolution/10,350*resolution/10)
			case spalte==1 && zeile==3:
				ap.Draw(resolution,214*resolution/10,400*resolution/10)
			case spalte==3 && zeile==4:
				ap.Draw(resolution,390*resolution/10,400*resolution/10)
			case spalte==5 && zeile==3:
				ap.Draw(resolution,566*resolution/10,400*resolution/10)
			case spalte==2 && zeile==4:
				ap.Draw(resolution,302*resolution/10,450*resolution/10)
			case spalte==4 && zeile==4:
				ap.Draw(resolution,478*resolution/10,450*resolution/10)
			case spalte==3 && zeile==5:
				ap.Draw(resolution,390*resolution/10,500*resolution/10)
	}
}
			
	

func (sf *data) drawPlättchenimSpiel (resolution uint16) {                      //Zeichnet die noch im Spiel befindlichen Plättchen
	for i:=uint16(0);i<uint16(len((*sf).plättchenImSpiel));i++ {   				//Geht durch die Liste der Plättchen durch und zerlegt sie in drei Reihen a 9 Plättchen und ordnet ihnen regelmäßig die Koordinaten zu
		(*sf).plättchenImSpiel[i].Draw(resolution,(i%9)*120*resolution/10+75*resolution/10,(i/9)*110*resolution/10+700*resolution/10)
	}
}

func (sf *data) DrawLegefläche (resolution uint16) {				//Zeichnet die bereits gelegten Plättchen an ihre richtige Position auf dem Spielfeld
	var legefläche [19]plättchen.Plättchen							//Generiert eine Variable, die die gelegten Plättchen aufnimmt 
	legefläche = sf.sortBelegteFelder()								//dazu wird die Funktion sortBelegteFelder verwendet, die die gelegten Plättchen sortiert
	var leeresPlättchen plättchen.Plättchen							//Da jetzt noch Nullerplättchen enthalten sind und diese nicht dargestellt werden sollen
	leeresPlättchen = plättchen.New(0,0,0)							//wird ein Nullerplättchen zum vergleich generiert
	for i:=0;i<len(legefläche);i++ {								//Die sortierten Plättchen werden mit der for-Schleife durchgegangen
		if 	plättchen.IstGleich(legefläche[i],leeresPlättchen) {	//Wenn es sich nicht um ein Nullerplättchen handelt
			continue
		} else {
		switch i {													//dann zeichne das Plättchen an der entsprechenden stelle auf dem Spielbrett an
			case 0:
				legefläche[i].Draw(resolution,390*resolution/10,100*resolution/10)		//3-1
			case 1:
				legefläche[i].Draw(resolution,302*resolution/10,150*resolution/10)		//2-1
			case 2:
				legefläche[i].Draw(resolution,478*resolution/10,150*resolution/10)		//4-1
			case 3:
				legefläche[i].Draw(resolution,214*resolution/10,200*resolution/10)		//1-1
			case 4:
				legefläche[i].Draw(resolution,390*resolution/10,200*resolution/10)		//3-2
			case 5:
				legefläche[i].Draw(resolution,566*resolution/10,200*resolution/10)		//5-1
			case 6:
				legefläche[i].Draw(resolution,302*resolution/10,250*resolution/10)		//2-2
			case 7:
				legefläche[i].Draw(resolution,478*resolution/10,250*resolution/10)		//4-2
			case 8:
				legefläche[i].Draw(resolution,214*resolution/10,300*resolution/10)		//1-2
			case 9:
				legefläche[i].Draw(resolution,390*resolution/10,300*resolution/10)		//3-3
			case 10:
				legefläche[i].Draw(resolution,566*resolution/10,300*resolution/10)		//5-2
			case 11:
				legefläche[i].Draw(resolution,302*resolution/10,350*resolution/10)		//2-3
			case 12:
				legefläche[i].Draw(resolution,478*resolution/10,350*resolution/10)		//4-3
			case 13:
				legefläche[i].Draw(resolution,214*resolution/10,400*resolution/10)		//1-3
			case 14:
				legefläche[i].Draw(resolution,390*resolution/10,400*resolution/10)		//3-4
			case 15:
				legefläche[i].Draw(resolution,566*resolution/10,400*resolution/10)		//5-3
			case 16:
				legefläche[i].Draw(resolution,302*resolution/10,450*resolution/10)		//2-4
			case 17:
				legefläche[i].Draw(resolution,478*resolution/10,450*resolution/10)		//4-4
			case 18:
				legefläche[i].Draw(resolution,390*resolution/10,500*resolution/10)		//3-5
			}
		}
	}
}


	

func (sf *data) drawZugnr (resolution uint16) {								//Zeichnet die Zugnummer mit dem zugehörigen Text
		gfx2.Stiftfarbe(0,0,0)								
		gfx2.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(5*resolution))
		var text string
		text = strconv.Itoa(int((*sf).zug)+1)								//Umwandlung der Zugnummer in einen String (+1), da der Zug auch als Index verwendet wird und daher bei 0 beginnt 
		gfx2.SchreibeFont(2*resolution,2*resolution,"Zug:"+text)
}	

func (sf *data) drawPunkte (resolution uint16) {							//Zeichnet die Punkte mit dem zugehörigen Text
		gfx2.Stiftfarbe(0,0,0)
		gfx2.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(5*resolution))
		var text string
		text = strconv.Itoa(int((*sf).punkte))								//Umwandlung der Punkte in einen String
		gfx2.SchreibeFont(60*resolution,2*resolution,"Punkte:"+text)
}

func (sf *data) drawAktPlättchen (resolution uint16) {						//Darstellung des aktuellen Plättchens zusammen mit dem zugehörigen Text
		gfx2.Stiftfarbe(0,0,0)
		gfx2.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
		var text string
		text = "Aktuelles"
		gfx2.SchreibeFont(82*resolution,475*resolution/10,text)
		text = "Plättchen"
		gfx2.SchreibeFont(82*resolution,50*resolution,text)
		(*sf).aktuellesPlättchen.Draw(resolution,755*resolution/10,500*resolution/10)
}

func (sf *data) drawSpielername (resolution uint16) {						//Darstellung des Spielernamens
	gfx2.Stiftfarbe(0,0,0)
	gfx2.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(5*resolution))
	var text string
	text = sf.GibSpielername()
	gfx2.SchreibeFont(100*resolution,2*resolution,"Spieler:")
	gfx2.SchreibeFont(100*resolution,7*resolution,text)
}	

