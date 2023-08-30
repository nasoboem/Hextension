package button



type Button interface {
	
//Vor.:-
//Erg.:Ein neuer sechseckiger Button ist geliefert, mit den Kooerdinaten x,y, mit der Größe = 50, Hintergrundfarbe (R 230, G 80, B 80), dem übergebenen Text in weißer Farbe (R 255, G 255, B 255), dem gegeben Offset des Texttes zum Mittelpunkt und der Highlighthintergrundfarbe (R 140, G 20, B 180).
//	New (resolution uint16, seitenname string,x,y uint16, text string, ox,oy int32)

//Set-Funktionen --------------------------------------------------------------------------------------------------------------------------------------

//Vor.:-
//Eff.:Der Mittelpunkt des Buttons hat jetzt die Koordinaten x,y
	SetzeKoordinaten (x,y uint16)

//Vor.:-
//Eff.:Der sechseckige Button hat jetzt die maximale Ausdehnung k vom Mittelpunkt (x-k, x+k und y-433/500k, y+433/500k)
	SetzeGroesse (k uint16)

//Vor.: HighlightGesetzt ist false, sonst wird die Highlighthintergrundfarbe verändert.
//Eff.: Die Hintergrundfarbe des Button ist jetzt auf r,g,b gesetzt.
	SetzeFarbe(r,g,b uint8)
	
//Vor.: -
//Eff.: Der übergeben Text ist jetzt der Text, der auf dem Button angezeigt wird.
	SetzeText (text string)

//Vor.: -
//Eff.: Die Textfarbe ist jetzt auf r,g,b gesetzt.
	SetzeTextFarbe (r,g,b uint8)
	
//Vor.:-
//Eff.: Der Offset der linken oberen Ecke destextes zum Mittelpunkt des Buttons ist jetzt auf x,y gesetzt. Negative Werte werden vom Mittelpunkt abgezogen, positive Werte zu den Koordinaten des Mittelpunkt dazu addiert.
	SetzeOffset (x,y int32)
	
//Vor.: HighlightGesetzt ist false, sonst wird die reguläre Hintergrundfarbe verändert.
//Eff.: Die Highlight-Farbe des Buttons ist auf r,g,b gesetzt.
	SetzeHighlightHintergrundFarbe (r,g,b uint8)
	
//Vor.:-
//Eff.: Der Seitenname, mit dem der Button verknüpft ist ist auf den gegebenen string veränder.
	SetzeSeitenNamen(sname string)

//Get-Funktionen -----------------------------------------------------------------------------------------------------------------------------------

//Vor.:-
//Erg.: Die x- und y-Koordinaten des Mittelpunkts des Buttons sind geliefert.
	GetKoordinaten () (x,y uint16)
	
//Vor.:-
//Erg.: Die Größe k des Buttons ist geliefert (In x-Dimension 2*k und in y-Dimension 2*433/500k)
	GetGroesse () (k uint16)
	
//Vor.:-
//Erg.:Die Hintergrundfarbe des Buttons ist in r,g,b geliefert.
	GetFarbe () (r,g,b uint8)

//Vor.:-
//Erg.:Der auf dem Button angezeigte Text ist geliefert.
	GetText () string
	
//Vor.:-
//Erg.:Der Offset der linken oberen Ecke des Textes zu den Koordinaten des Mittelpunkts des Buttons ist geliefert.
	GetOffset () (x,y int32)
	
//Vor.:-
//Erg.:Die Farbe des Textes ist in r,g,b geliefert.	
	GetTextFarbe () (r,g,b uint8)

//Vor.:-
//Erg.:Die Highlight-Farne des Buttons ist in r,g,b geliefert.
	GetHighlightFarbe () (r,g,b uint8)
	
//Vor.: -
//Erg.: Der Name der Seite, die mit dem Buttonverknüpft ist ist gegeben.
	GetSeitenName () string
	
// Andere ------------------------------------------------------------------------------------------------------------------------------------------
	
//Vor.:-
//Erg.: Ein True ist geliefert, wenn die Highlightfarbe mit der Regulären Hintergrundfarbe getauscht wurde. False ist geliefert, wenn Highlight und
	HighlightGesetzt () bool
	
//Vor.: Ein gfx-Fenster ist offen. Die Koordinaten des zu Zeichnenden Buttons liegen vollständig innerhalb des Fensters. (x-Koordinate +/- Größe & y-Kooerdinate +/- 433/500 * Größe)
//Eff.: Der Button ist im gfx-Fenster dargestellt.
	Draw (resolution uint16) 
	
//Vor.:-
//Erg.: True ist geliefert, genau dann wenn die gegeben x-/y-Kooerdinaten sich innerhalb des Buttons befinden.
	GehörtPunktzuButton (x,y uint16) bool
	
//Vor.:-	
//Eff.: Tauscht die Highlighthintergrundfarbe mit der regulären Hintergrundfarbe aus und setzt ein Highlight-Flag	
	SwitchBackGround ()
}

//Vor.:-
//Erg.: True ist geliefert, genau dann wenn zwei Buttons in ihrer verknüpften Seite, ihren Kooerdinaten, ihrem angezeigten Text und dem Offset des Textes übereinstimmen.

func IstGleich(b1,b2 Button) bool {
	var b1x,b1y,b2x,b2y uint16
	var b1s,b2s,b1t,b2t string
	var b1ox,b1oy,b2ox,b2oy int32
	b1x,b1y = b1.GetKoordinaten()
	b1s = b1.GetSeitenName()
	b1t = b1.GetText()
	b1ox,b1oy = b1.GetOffset()
	b2x,b2y = b2.GetKoordinaten()
	b2s = b2.GetSeitenName()
	b2t = b2.GetText()
	b2ox,b2oy = b2.GetOffset()
	return b1x==b2x && b1y==b2y && b1s==b2s && b1t==b2t && b1ox==b2ox && b1oy==b2oy
}
