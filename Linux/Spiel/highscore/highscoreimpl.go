package highscore

import (
		"../dateien"
		"strconv"
		"fmt"
		"../gfx"
		)

type data struct { //Der Highscore besteht aus zwei Slices eins für die Namen und eins für die Punkte. Gleiche Indices weisen Paare aus
	namen []string
	punkte []int
}

func New() *data {  //Generiert einen leeren Highscore mit "-" als Name und 0 als Punkte auf den 6 Einträgen - 6 wegen Hex-tension
	var h *data
	h = new(data)
	(*h).namen = make([]string,6)   //Slice mit 6 Plätzen im namen-Slice
	for i:=0;i<len((*h).namen);i++{	//Eintragung der "-"-Namensplatzhalter
		(*h).namen[i] = "-"
	}
	(*h).punkte = make([]int,6)		//Slice mit 6 Plätzen im punkte-Slice
	for i:=0;i<len((*h).punkte);i++{//Eintragung der 0-Werte
		(*h).punkte[i] = 0
	}
	return h
}

func (h *data) Einfügen (name string, punkte int) (index int) { //Sortiert einen übergebenes Ergebnis bestehend aus einem Spielernamen und seinem Punktestand in ein bestehenden Highscore ein und gibt den Index an dem er einsortiert wurde zurück
	index = 7													//Indexwert wird der Wert 7 gegeben, der nicht mehr im Highscore enthalten ist
	for i:=0;i<len((*h).punkte);i++{  							//Geht den Punkte-Slice durch und vergleicht die übergeben Punkte mit den Punkten im Highscore
		if (*h).punkte[i]>punkte {								//Sind die Punkte kleiner, gehe zum nächsten Eintrag
			continue
		} else {												//Ist der übergebene Punktestand größer oder gleich...(bedeutet, dass gleicher Punktestand wird weiter oben eingefügt - Letzter Spieler bevorzugt)
			var zergInt []int									//Zwischenergebnis-Slices werden erstellt
			var zergString []string
			zergInt = append(zergInt,(*h).punkte[:i]...)		//Der neue übergeben Punktestand wird in den Punkte-Slice des Higscores an der Indexposition i eingespleißt 
			zergInt = append(zergInt,punkte)
			zergInt = append(zergInt,(*h).punkte[i:len((*h).punkte)-1]...) //Der Letze Eintrag aus dem Highscore fliegt raus
			zergString = append(zergString,(*h).namen[:i]...)				//Gleiches mit dem Namens-Slice
			zergString = append(zergString,name)
			zergString = append(zergString,(*h).namen[i:len((*h).namen)-1]...)
			(*h).punkte = zergInt											//Übertragen der Zwischen Ergebnisse in den Highscore
			(*h).namen = zergString
			index = i														//Index wird auf i gesetzt (i<6)
		}
		break
	}
	return																	//Index wird zurückgegeben
}

func (h *data) Speichern () {												//Speichert den Highscore in zwei Dateien, die Namen in highscore-namen.txt und die Punkte in highscore-punkte.txt und trennt die Einträge mit einem Zeilenumbruch 
	var d dateien.Datei
	var b byte
	d = dateien.Oeffnen ("highscore-namen.txt",'s')							//Die Datei wird zum schreiben geöffnet, der vorherige Inhalt wird überschrieben
	for i:=0;i<len((*h).namen);i++{											// for-Loop, der durch die namens Einträge des Highscores durchläuft
		for j:=0;j<len((*h).namen[i]);j++ {									// for-Loop, der durch die Bytes des String durchgeht
			b = (*h).namen[i][j]
			d.Schreiben(b)													//Die Bytes werden nach einander in die Datei geschrieben
		}
		b = '\n'															//Nach jedem Namen wird ein Zeilenumbruch eingefügt
		d.Schreiben(b)
	}
	d.Schliessen()															//Datei wird geschlossen
	d = dateien.Oeffnen ("highscore-punkte.txt",'s')						//Funktionsgleich, für die Punkte, hier werden nur die Punkte (int) vorher mithilfe der Funktion strconv.Itoa in strings umgewandelt 
	for i:=0;i<len((*h).punkte);i++{
		var intString string
		intString = strconv.Itoa((*h).punkte[i])
		for j:=0;j<len(intString);j++ {
			b = intString[j]
			d.Schreiben(b)
		}
		b = '\n'
		d.Schreiben(b)
	}
	d.Schliessen()
}



func (h *data) Lesen () {													//Lesen ließt die beiden Dateien aus und Überträgt die Informationen in einen Highscore
	var d dateien.Datei
	d = dateien.Oeffnen ("highscore-namen.txt",'l')							//Öffnen der Datei zum lesen
	var b byte
	var name []byte
	var namen []string
	for !d.Ende() {															//solange das Dateiende noch nicht erreicht ist
		b = d.Lesen()														//Lese ein Byte ein
		if b!='\n' {														//Wenn das zuletzt gelesene Byte kein Zeilenumbruch ist
			name = append(name,b)											//hänge das byte an den namen an
		} else {															//wurde ein Zeilenumbruch erkannt, dann ist der Name vollständig
			namen = append(namen,string(name))								//Hänge den Namen ([]byte) als String in die Namensliste ein
			name = make([]byte,0)											//Leere den name []byte, so dass wieder neue Bytes entgegen genommen werden können
		}		
	}																		//Wurde die Datei vollständig ausgelesen, dann ...
	(*h).namen = namen														//Namensliste wird in de Highscore eingetragen
	d.Schliessen()															//Datei wird geschlossen
	
	d = dateien.Oeffnen ("highscore-punkte.txt",'l')						//Baugleich für die Punkte
	var punkte int
	var err error
	var c byte
	var spunkte []byte
	var punkteSlice []int
	for !d.Ende() {															//solange das Dateiende noch nicht erreicht ist
		c = d.Lesen()														//lese ein Byte
		if c!='\n'{															//wenn dieses Byte kein Zeilenumbruch ist, dann
				spunkte = append(spunkte,c)									//hänge das Byte an einen Slice von Byte an
		} else {
			punkte,err = strconv.Atoi(string(spunkte))						//Wurde ein Zeilenumbruch erkannt, dann ist die Zahl vollständig und der Zahlen-String kann mithilve von strconv.Atoi umgewandelt werden - Sollte die Umwandlung nicht möglichsein, dann bricht die Funktion in einer panic ab 
			spunkte = make([]byte,0)										//Leeren des Byte-Slices
			if err==nil {
				punkteSlice = append(punkteSlice,punkte)					//Anhängen der Punkte an den Punkte-Slice
			} else {
				panic ("In der highscore-punkte.txt Datei befinden sich keine Zahlenwerte!!")
			}
		}
	}
	(*h).punkte = punkteSlice												//Punkte-Slice wird in den Highscore eingetragen
	d.Schliessen()															//Datei wird geschlossen
}


func (h *data) String () (erg string){				//Wandelt den Highscore in einen druckbaren String um
	var nn []string
	var lname uint
	lname,nn = nameSize((*h).namen)					//Die Namenslliste aus dem Highscore wird an dieFunktion nameSize übergeben, die eine Namensliste zurück gibt, in der alle Strings die gleiche länge haben, es wird mit Leerzeichen aufgefüllt und die Länge wird in lname zurück gegeben
	erg = erg + fmt.Sprint(" Platz   Spieler")		//Kopfzeile wird gebildet
	for i:=0;i<(int(lname)-7);i++ {					//Fügt entsprechend leerzeilen in Abhängigkeit der Namenslänge ein
		erg = erg + fmt.Sprint(" ")
	}
	erg = erg + fmt.Sprintln("   Punkte")			//Fügt den rest der Kopfzeile an
	länge:=len(erg)									//Der aus der Kopfzeile bestehende String wird gemessen und 
	for i:=0;i<länge-1;i++{							//eine zwischen zeile mit der entsprechenden Länge aus ++++++ eingefügt.
		erg = erg + "+"
	}
	erg = erg + fmt.Sprintln("") 					//Inhalt der Highscore Tabelle wird eingetragen
	for i:=0;i<len(nn);i++{
		erg = erg + fmt.Sprint("   ")				//Platzhalter
		erg = erg + fmt.Sprint(i+1)					//Plazierung
		erg = erg + fmt.Sprint(".  |  ")			//Punkt und Trennungsstrich
		erg = erg + fmt.Sprint(nn[i])				//Name
		if lname<7 {								//Sind die Namen kürzer al 7, so wird ein Mindestabstand eingefügt, damit die Tabelle unter ihren überschriftrn steht und nicht alles zu weit nach rechts rutscht.
			for i:=uint(0);i<7-lname;i++{
				erg = erg + fmt.Sprint(" ")
			}
		}
		erg = erg + fmt.Sprint(" |  ")				//Einfügen des zweiten trennungsstrich
		erg = erg + fmt.Sprintln((*h).punkte[i])	//Einfügen der Punkte
	}
	for i:=0;i<länge-2;i++{							//Einfügen einer abschließenden ++++++ Zeile
		erg = erg + "+"
	}
	return
}

func nameSize (namen []string) (uint, []string) {	//Bekommt ein Slice von string in dem die Strings unterschiedliche länge haben und gibt die länge des längsten Strings zurück, sowie ein Slice von string, bei dem die kürzeren Strings mit Leerzeichen aufgefüllt sind, so dass alle gleich lang sind
	var max uint
	var count []uint
	var erg []string
	if namen==nil {									//Bei der Übergabe eines leeren Slices wird 0 und ein leerer Slice zurück gegeben
		return max, erg
	}
	for i:=0;i<len(namen);i++{						//Geht den Slice Namen für namen durch 
		var zähler uint
		for range namen[i] {						// und Zählt die Buchstaben im den Strings
			zähler++
		}
		count = append(count,zähler)				//Die Zählerstände für jeden String werden an den Count-Slice angehangen
	}
	max = count[0]									//Suche nach dem längsten String
	for i:=1;i<len(count);i++ {
		if max>=count[i] {							//Wenn max >= dem Eintrag in count ist nimm den nächsten Eintrag
			continue
		} else {
			max = count[i]							//wenn max kleiner ist setze max gleich den aktuellen Eintrag
		}
	}
	for i:=0;i<len(namen);i++{						//Geht die Namensliste durch um die Leerzeichen einzufügen
		var lname string
		lname = namen[i]							//Der aktuelle Name wird der Variablen lname übergeben
		for j:=uint(0);j<(max-count[i]);j++ {		//Es werden soviele Leerzeichen eingefügt, wie die Differen zwischen dem längsten Namen (max) und der eigenen Länge (count[i])
			lname = lname + " "
		}
		erg = append(erg,lname)						//Der durch Leerzeichen verlängerte Name wird an den Ergebnis-Slice angehangen
	}
	return max, erg									//Rückgabe der maximalen Länge und der auf die gleichen länge gebrachten Namensliste
}

func (h *data) Draw (resolution uint16,plaetze []int) {						//Zeichenfunktion für den Highscore; die plaetze eingetragenen Plazierungen werden rot hervorgehoben (neueintragungen in den Highscore), resolution passt die Ausgabe im gfx.Fenster an
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(4*resolution))		//Festlegung von Font und Stiftfarbe
	gfx.SchreibeFont(40*resolution,30*resolution,"Platz")					//Kopfzeile wird geschrieben
	gfx.SchreibeFont(67*resolution,30*resolution,"Spieler")
	gfx.SchreibeFont(100*resolution,30*resolution,"Punkte")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(3*resolution))		//Reduktion der Schriftgröße für die Einträge im Highscore
	for i:=uint16(0);i<uint16(len((*h).namen));i++ {						//Für die Länge der Namensliste
		if enthalten(i,plaetze) {											//Wenn der Index in der plaetze-Liste enthalten ist, dann
			gfx.Stiftfarbe(255,0,0)											//dann nehme die Stiftfarbe rot
		} else {
			gfx.Stiftfarbe(0,0,0)											//wenn nicht dann die Stiftfarbe schwarz
		}	
		gfx.SchreibeFont(43*resolution,39*resolution+(7*resolution*i),fmt.Sprint(i+1)+".")	//Zeichnen der eigentlichen Einträge
		gfx.SchreibeFont(70*resolution,39*resolution+(7*resolution*i),(*h).namen[i])
		gfx.SchreibeFont(103*resolution,39*resolution+(7*resolution*i),fmt.Sprint((*h).punkte[i]))
	}
}	

func enthalten (wert uint16,werteliste []int) bool {						//Überprüft, ob ein Wert in einer gegeben Liste enthalten ist und liefert ein true, wenn der Wert enthalten ist und ein false wenn dem nicht so ist
	for i:=0;i<len(werteliste);i++{											//Geht die Werte-Liste durch und vergleicht den übergeben Wert mit einträgn in der Liste
		if int(wert)==werteliste[i] {
			return true														//Bei einer Übereinstimmung gebe true zurück 
		} else {
			continue
		}
	}
	return false															//Ist die Liste Abgearbeitet ohne Treffer gebe false zurück
}

	
