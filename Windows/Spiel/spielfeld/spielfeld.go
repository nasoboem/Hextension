package spielfeld

import ("../feld"
		"../plättchen"
		"../seite"
		)

type Spielfeld interface {

//Vor.: -
//Erg.: Ein leeres Spielfeld mit Plättchen und 0 Punkten ist geliefert.
//	New (resolution uint16, spielername string)

//Vor.: -
//Erg.: True ist geliefert, genau dann wenn das gegebene Feld noch frei ist
	FeldFrei (f feld.Feld) bool

//Vor.: Das Feld auf das das Plättchen gesetzt werden soll ist frei. 
//Eff.: Das Plättchen wird gesetzt und die Punkte entsprechend angepasst. Danach wird automatisch ein neues Plättchen gewählt.
	PlättchenSetzen (spiel [19]uint8)
	
//Vor.: -
//Eff.: Wählt ein Feld aus, auf dem das aktuelle Plättchen gelegt werden soll.
	FeldSetzen (f feld.Feld)
	
//Vor.: -
//Erg.: Das aktuelle Feld ist geliefert
	GibaktuellesFeld () feld.Feld
	
//Vor.: -
//Erg.: Die Zugnummer ist geliefert (1-19)
	GibZug () uint8
	
//Vor.: -
//Erg.: Der aktuelle Punktestand ist geliefert.
	GibPunkte () int

//Vor.: -
//Erg.: Das aktuelle Plättchen ist geliefert.
	GibaktuellesPlättchen () plättchen.Plättchen

//Vor.: -
//Eff.: Zieht ein Plättchen aus den noch im Spiel befindlichen Plättchen und macht es zum aktuellen Plättchen auf dem Spielfeld.
	PlättchenZiehen (spiel [19]uint8)
	
//Vor.: -
//Erg.: Der aktuelle Stand des Spielfeldes (Zug, Legefläche, aktuell gewähltes Feld, aktuelle Plättchen und noch im Spiel befindliche Plättchen werden als strukturieter String zurück gegeben
	String () string

//Vor.: Es ist ein gfx-Fenster geöffnet mit der größe 160*resolutionx100*resolution
//Eff.: Das Spielfeld ist im gfx-Fenster dargestellt.
	Draw (resolution uint16)
	
//Vor.: -
//Erg.: Der Name des Spielers, der mit diesem Spielfeld spielt ist geliefert.	
	GibSpielername () string
	
//Vor.: Es ist ein gfx-Fenster geöffnet mit der größe 160*resolutionx100*resolution
//Eff.: Das aktuelle Plättchen wird auf dem vom Spieler ausgewählten Feld dargestellt. 
	PlättchenAnzeigen (resolution uint16)
	
//Vor.: -
//Eff.: Die gegeben Punkte sind auf dem Spielfeld eingetragen.
	SetzePunkte (punkte int)

//Vor.: -
//Eff.: Legt ein geegbenes Plättchen auf das angegeben Feld des Spielfeldes  -  Mit dieser Funktion lassen sich belibige Konstelationen auf dem Spielfeld generrieren (auch welche, die im eigentlichen Spiel unmöglich sind).
	FreiesLegen (f feld.Feld, p plättchen.Plättchen)
	
//Vor.: -
//Eff.: Das letzte Plättchen ist vom Spielfeld entfernt und alle veränderungen des Spielfeldes die damit einhergehen verändern sich 
//		(Reduktion des Punktestands, Reduktion der Zugnummmer), sofern noch ein Plättchen auf dem Spielfeld ist.
//Erg.: Des zuletzt gespielte Plättchen ist geliefert. Befindet sich kein Plättchen mehr auf dem Spielfeld wird ein Nullerplättchen zurückgegeben - Wurde verwendet um ein Brut-Force-Allgorithmus zu bauen. 
	LetztesPlättchenEntfernen () (p plättchen.Plättchen)

//Vor.: -
//Eff.: Die im 	Spiel vorkommenden Plättchen sind aus der Anzeige der "Plättchen im Spiel" entfernt
	PlättchenAusdemSpielnehmen (spiel [19]uint8)
	
//Vor.: Es ist ein gfx-Fenster geöffnet mit der größe 160*resolutionx100*resolution
//Eff.: Die bisher auf die Spielfläche gelegten Plättchen sind im gfx-Fenster dargestellt.
	DrawLegefläche (resolution uint16)

//Vor.: -
//Erg.: Die Seite 1 des Spielfeldes ist geliefert. - Vor der Wahl eines Feldes durch den Spieler
	GibSeite1 () seite.Seite
	
//Vor.: -
//Erg.: Die Seite 2 des Spielfeldes ist geliefert. - Nach der Wahl eines Feldes durch den Spieler.
	GibSeite2 () seite.Seite
	
//Vor.: -
//Eff.: Die gegeben Seiet ist nun die 1. Seite des spielfeldes.
	SetzeSeite1 (s seite.Seite)
	
//Vor.: -
//Eff.: Die gegeben Seite ist nun die 2. Seite des Spielfeldes.
	SetzeSeite2 (s seite.Seite)

}

//Vor.: -
//Erg.: Ein Spielfeld ist geliefert mit dem Spielernamen "Max Mustermann", und dem Spielstand, der durch die im Spiel festgelegte Reihenfolge der Plättchen und anzahl und Auswahl der Felder im Slice f.
func SpielfeldGenerator (resolution uint16, f []feld.Feld, spiel [19]uint8) Spielfeld { //Mithilfe des  SpielfeldGenerator lassen sich beliebige Zustände des Spielfeldes generieren, mit dem Spielernamen Max Mustermann
																						//aber nur Zustände, die auch imn Spiel erlaubt und möglich sind 
	var sf Spielfeld
	sf = New(resolution,"Max Mustermann")
	sf.PlättchenZiehen(spiel)															//Ziehen des ersten Plättchens (Belegung des aktuellen Plättchens
	for i:=0;i<len(f);i++{																//Es werden so viele Plättchen (Abfolge der Plättchen wird durch das Spiel definiert) auf das Spielfeld gelegt, wie Felder in f sind
		if i<len(spiel) {
			sf.FeldSetzen(f[i])															//Feld wird festgesetzt (wie beim Zu des Spielers
			sf.PlättchenSetzen(spiel)													//Dann wird das aktuelle Plättchen auf das Feld gelegt und das nächste Plättchen nach dem Spiel zum aktuellen Plättchen gemacht usw.
		}
	}
	return sf																			//Das so bespielte Spielfeld wird zurück gegeben.  - Wurde verwendet für die Erklärung der Regeln.
}
		
