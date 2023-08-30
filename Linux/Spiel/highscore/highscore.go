package highscore

type Highscore interface {
	
//Vor.: -
//Erg.: Ein neuer leerer Highscore mit 6 Einträgen ist geliefert.
//	Name	Punkte
//	- 		0
//	-		0
//	-		0
//	-		0
//	-		0
//	-		0
// New()


//Vor.: -
//Eff.: Das übergebene Spielergebnis ist in den Highscore ingetragen, wenn es gleichwertig oder besser war als die Einträge im Highscore.
//Erg.: Der Index der Position an der das übergebene Spielergebnis in den Highscore eingetragen wurde ist gegeben.
	Einfügen (name string, punkte int) (index int)

//Vor.: Eine Datei highscore-namen.txt und highscore-punkte.txt existieren im Ordner, der aufrufenden Funktionsdatei.
//Eff.: Der Highscore ist in den jeweiligen Datein abgelegt.
	Speichern ()

//Vor.: Eine Datei highscore-namen.txt und highscore-punkte.txt existieren im ordner, der aufrufenden Funktionsdatei. Die Dateien enthalten 6 Einträge, die mit Zeilenumprüche separiert sind. Die Einträge in der highscore-punkte.txt müssen positive Integer-Werte sein.
//Eff.: Der Highscore ist aus den Datein ausgelesen und in den Highscore eingetragen.
	Lesen ()

//Vor.: -
//Erg.: Der Highscore ist als String gegeben, in der Form
// Platz   Spieler   Punkte
//++++++++++++++++++++++++++
//   1.  | -       |   0
//   2.  | -       |   0
//   3.  | -       |   0
//   4.  | -       |   0
//   5.  | -       |   0
//   6.  | -       |   0
//++++++++++++++++++++++++++
	String () (erg string)

//Vor.: Ein gfx-Fenster ist geöffnet, in dem der Highscor dargestellt werden kann (1600*resolution/10x1000*resolution/10).
//Eff.: Der Highscore ist im gfx-Fenster angezeigt. Der übergebe Slice der Plazierungen wird roter Schrift angezeigt, der rest in schwarzer Schrift. Index bei 0 beginnend.
//dient zum Highlighting der neuen Einträge. Slice aus Plätzen sollte aus den Indices der Einfügen-Funktion gebildet werden. 
	Draw (resolution uint16,plaetze []int) 
}
