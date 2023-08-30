package button

import ("../sechseck"
		"gfx2"
		)
		
type data struct {
	s sechseck.Sechseck		//Button greift auf Sechseck zurück. Position, Größe und Farbe werden im Sechseck gespeichert
	sname string			//Hier wird der Name gespeicher, für den der Button zuständig sind
	text string				//Speichert den Text, der auf dem Button angezeigt wird
	tr,tg,tb uint8			//Textfarbe
	hr,hg,hb uint8			//Farbe des Hintergrundwechsels, wenn die Maus drüber fährt
	offset [2]int32			//Offset vom Mittelpunkt des Sechsecks. offset[0] entspricht dem Mittelpunkt. negative Zahlen werden abgezogen, positive dazu addiert.
	highlight bool			//Highlightfarbe gesetzt
}

func New (resolution uint16,sname string,x,y uint16,text string, ox,oy int32) *data {		// In der New Funktion wird zwischen Buttons mit Text und ohne Text unterschieden.
	var bt *data
	bt = new(data)
	switch {
		case text =="":
			(*bt).s = sechseck.New(x*resolution/10,y*resolution/10,50*resolution/10,60,60,60)		//Größe - 50, Farbe (R 60, G 60, B 60) Dunkelgrau - Entspricht dem Aussehen eines Nuller-Plättchens (0,0,0) - Buttons werden für die Legefläche verwendet
		case sname =="":
			(*bt).s = sechseck.New(x*resolution/10,y*resolution/10,50*resolution/10,140,140,140)  	//Größe - 50, Farbe (R 140, G 140, B 140) hellgrau - Darstellung funktionsloser Buttons, die keine Seitennamen als Ausgabe haben
		default:
		(*bt).s = sechseck.New(x*resolution/10,y*resolution/10,50*resolution/10,230,80,80)			//Größe = 50, Farbe (R 230, G 80, B 80) Schmutziges Rot werden ursprünglich vorgegeben
	}
	(*bt).sname = sname								//Name der verknüpften Seite wird übergeben
	(*bt).text = text								//Auf dem Button angezeigter Text wird übergeben
	(*bt).offset[0] = ox*int32(resolution)/10		//Hier wird der Offset des Textes an die Auflösung angepasst
	(*bt).offset[1] = oy*int32(resolution)/10	
	(*bt).tr = 255									//Stiftfarbe für den Text wird auf Weiß gesetzt
	(*bt).tg = 255
	(*bt).tb = 255
	switch {
		case text == "":							//- Bei Buttons der Legefläche
			(*bt).hr = 140							//HighlightHintergrundfarbe für die Buttons der Legefläche - helles Grau
			(*bt).hg = 140
			(*bt).hb = 140
		case sname =="":							//- Bei inaktiven Buttons - noch nicht implementiert, müsste von der PressButton-Funktion aussortiert werden 
			(*bt).hr = 140							//HighlightHintergrundfarbe für funktionslose Buttons ist die gleiche, wie die normale Hintergrundfarbe - helles Grau
			(*bt).hg = 140
			(*bt).hb = 140
		default:									//Standard-Button
			(*bt).hr = 140							//HighlightHintergrundfarbe wird auf ein schmutziges Blau gesetzt 
			(*bt).hg = 20
			(*bt).hb = 180
	}
	(*bt).highlight = false							//Ist Highlight false, dann ist die im Sechseck gespeicherte Hintergrundfarbe die reguläre Hintergrundfarbe und die Highlightfarbe ist gespeichert in hr,hg,hb. Ist highlight true, dann wurden die beiden Farben vertauscht
	return bt
}


//Set-Funktionen

func (bt *data) SetzeKoordinaten (x,y uint16) { 	//Setzt die Koordinaten des Buttons
	(*bt).s.SetzeKoordinaten(x,y)
}

func (bt *data) SetzeGroesse (k uint16) {		//Setzt die Größe des Buttons
	(*bt).s.SetzeGroesse(k)
}

func (bt *data) SetzeFarbe(r,g,b uint8) {		//Setzt die Farbe des Buttons
	(*bt).s.SetzeFarbe(r,g,b)
}

func (bt *data) SetzeText (text string) {		//Setzt den Text, der auf dem Button dargestellt wird
	(*bt).text = text
}

func (bt *data) SetzeTextFarbe (r,g,b uint8) {	//Setzt die Text-Farbe
	(*bt).tr = r
	(*bt).tg = g
	(*bt).tb = b
}

func (bt *data) SetzeSeitenNamen (sname string) { //Setzt den Namen der verknüpften Seite
	(*bt).sname = sname
}

func (bt *data) SetzeOffset (x,y int32) {		//Setzt den Offset des Textes zum Mittelpunkt des Buttons
	(*bt).offset[0] = x
	(*bt).offset[1] = y
}

func (bt *data) SetzeHighlightHintergrundFarbe (r,g,b uint8) {	//Setzt die Highlight-Farbe des Buttons
	(*bt).hr = r
	(*bt).hg = g
	(*bt).hb = b
}

//Get-Funktionen

func (bt *data) GetSeitenName () string {		//Gibt den Seitennamen der verknüpften Seite als String aus - String kann auch für etwas anderes verwendet werden
	return (*bt).sname
}

func (bt *data) GetKoordinaten () (x,y uint16) { //Liest die Koordinaten des Sechsecks aus
	return (*bt).s.GibKoordinaten()
}

func (bt *data) GetGroesse () (k uint16) { //Liest die Größe des Sechsecks aus
	return (*bt).s.GibGroesse()
}

func (bt *data) GetFarbe () (r,g,b uint8) { //liest die Farbe des Sechsecks aus
	return (*bt).s.GibFarbe()
}

func (bt *data) GetText () string { //Liest die Beschriftung des Buttons aus
	return (*bt).text
}

func (bt *data) GetOffset () (x,y int32) { //Liest den x- und Y-Koordinaten offset aus
	x = (*bt).offset[0]
	y = (*bt).offset[1]
	return
}

func (bt *data) GetTextFarbe () (r,g,b uint8) { //Gibt die Textfarbe in r,g,b zurück
	r = (*bt).tr
	g = (*bt).tg
	b = (*bt).tb
	return
}

func (bt *data) GetHighlightFarbe () (r,g,b uint8) { //Gibt ide Highlightfarbe in r,g,b zurück
	r = (*bt).hr
	g = (*bt).hg
	b = (*bt).hb
	return
}

func (bt *data) HighlightGesetzt () bool {		//Tested, ob der Button gerade im Highlightmodus ist oder nicht
	return (*bt).highlight
}

//Weitere Funktionen

func (bt *data) SwitchBackGround () { //Tauscht die eigendliche Hintergrundfarbe des Buttons mit seiner Highlightfarbe aus und setzt das Highlightflag
	var hr,hg,hb,r,g,b uint8
	hr,hg,hb = bt.GetHighlightFarbe()
	r,g,b = (*bt).s.GibFarbe()
	bt.SetzeHighlightHintergrundFarbe(r,g,b)
	(*bt).s.SetzeFarbe(hr,hg,hb)
	if (*bt).highlight {
		(*bt).highlight = false
	} else {
		(*bt).highlight = true
	}
}
	

func (bt *data) Draw (resolution uint16) {		//Zeichenfunktion für Buttons					
	(*bt).s.Draw()											//Zeichnet das Sechseck
	gfx2.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(20*resolution/10))	//Setzt den Font für den Text
	var text string
	text = (*bt).text										//Text wird zwischengeseichert in der Variablen "text"
	gfx2.Stiftfarbe((*bt).GetTextFarbe())					//Textfarbe wird gesetzt
	var x,y,a,b uint16 										//x,y übernehmen die Koordinaten des Sechsecks, a,b sind die Kooerdinaten der linken oberen Ecke ander der Text geschrieben wird	
	var ox,oy int32										   //ox,oy übernehmen den Offset des Textes für x- und y-Koordinate
	x,y = (*bt).GetKoordinaten()
	ox,oy = (*bt).GetOffset()
	switch {												//Berechnet die richtigen Koordinaten des Texts aus den x-.y-Koordinaten und ihrem Offset
		case ox < 0 && oy < 0:
			a = x - uint16(ox*-1)
			b = y - uint16(oy*-1)
		case ox >= 0 && oy < 0:
			a = x + uint16(ox)
			b = y - uint16(oy*-1)
		case ox < 0 && oy >= 0:
			a = x - uint16(ox*-1)
			b = y + uint16(oy)
		case ox >= 0 && oy >= 0:
			a = x + uint16(ox)
			b = y + uint16(oy)
		}
	if i,ok:=enthältZeilenumbruch(text);!ok {	
		gfx2.SchreibeFont(a,b,text)								//Schreibt den Font
	} else {
		var text1,text2 string
		text1 = text[:i]
		text2 = text[i+1:]
		gfx2.SchreibeFont(a,b,text1)								//Schreibt den Font
		gfx2.SchreibeFont(a,b+20*resolution/10,text2)
	}
}


func (bt *data) GehörtPunktzuButton (x, y uint16) bool {	//Prüft, ob eine gegebene Kooerdinate auf dem Buttonliegt oder nicht.
	return (*bt).s.GehörtPunktzuSechseck(x,y)
}

func enthältZeilenumbruch (text string) (int, bool) {	//Zählt die Zeihlenümbrüche in einem String und gibt ein True zurück, wenn sich welche im Text befinden
	for i,s:=range (text) {
		if s=='\n' {
			return i,true
		}
	}
	return 0,false
}
