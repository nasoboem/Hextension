package sechseck

import ("../gfx"
		"../core"
		//"fmt"
		)

type data struct {
	r,g,b uint8		//Farbe des Vollsechseck
	x,y uint16		//Kooerdinaten des Mittelpunkts
	k uint16		//Kantenlänge der gleichseitigen Dreicke aus denen das Sechseck besteht. Ausdehnung des Sechsecks in x-Dimension
}

func New(x,y,k uint16, r,g,b uint8) *data { //Generiert ein Sechseck aus den übergebenen Parametern
	var s *data
	s = new(data)
	(*s).x = x
	(*s).y = y
	(*s).k = k
	(*s).r = r
	(*s).g = g
	(*s).b = b
	return s
}

func (s *data) SetzeFarbe (r,g,b uint8) { //Setzt die Farbe des Sechsecks
	(*s).r = r
	(*s).g = g
	(*s).b = b
}

func (s *data) SetzeKoordinaten (x,y uint16) {	//Setzt die Kooerdinaten des Sechsecks
	(*s).x = x
	(*s).y = y
}

func (s *data) SetzeGroesse (k uint16) {	//Setzt die Kantenlänge der gleichseitigen Dreiecke aus denen das Sechseck aufgebaut ist
	(*s).k = k
}

func (s *data) GibFarbe () (r,g,b uint8) { //Gibt die Farbe des Sehcsecks wieder in r,g,b
	r = (*s).r
	g = (*s).g
	b = (*s).b
	return
}

func (s *data) GibKoordinaten () (x,y uint16) { //Gibt die Kooerdinaten des Mittelpunktes des Sechsecks zurück
	x = (*s).x
	y = (*s).y
	return
}

func (s *data) GibGroesse () (k uint16) { //Gibt die Kantenlänge und damit die Ausdehnung des Sechsecks in der x-Dimension wieder
	k = (*s).k
	return
}

//Eckpunkte werden benötigt, um die Streifen auf die Plättchen zu zeichnen

// L-Ecken sind die Ecken für den Streifen, der von links unten nach rechts oben geht 

func (s *data) GibLEcken () (lu1,lu2,ro1,ro2 [2]uint16) { //lu1[0] = x Wert; lu1[1] = y Wert
	var x,y,k,h uint16
	x,y = s.GibKoordinaten()
	k = s.GibGroesse()
	h = k*433/500
	
	lu1[0] = x-k		//Berechnung der Kooerdinaten der entsprechenden Punkte
	lu1[1] = y
	lu2[0] = x-(k/2)
	lu2[1] = y + h
	ro1[0] = x +(k/2)
	ro1[1] = y - h
	ro2[0] = x + k
	ro2[1] = y
	return
}

// O-Ecken sind die Ecken für den senkrechten Streifen 

func (s *data) GibOEcken () (o1,o2,u1,u2 [2]uint16) {
	var x,y,k,h uint16
	x,y = s.GibKoordinaten()
	k = s.GibGroesse()
	h = k*433/500
	
	o1[0] = x -(k/2)
	o1[1] = y - h
	o2[0] = x + (k/2)
	o2[1] = y - h
	u1[0] = x - (k/2)
	u1[1] = y + h
	u2[0] = x + (k/2)
	u2[1] = y + h
	return
}

// R-Ecken sind die Ecken für den Streifen, der von rechts unten nach links oben geht 

func (s *data) GibREcken () (lo1,lo2,ru1,ru2 [2]uint16) {
	var x,y,k,h uint16
	x,y = s.GibKoordinaten()
	k = s.GibGroesse()
	h = k*433/500
	
	lo1[0] = x - k
	lo1[1] = y
	lo2[0] = x - (k/2)
	lo2[1] = y - h
	ru1[0] = x + (k/2)
	ru1[1] = y + h
	ru2[0] = x + k
	ru2[1] = y
	return
}

 

func (s *data) Draw () {					//Zeichnet das Sechseck durch die Angabe der 6 gleichseitigen Dreiecken aus denen es besteht
	var x,y,k,h uint16
	
	gfx.Stiftfarbe (s.GibFarbe())			//Abrufen der Informationen
	x,y = s.GibKoordinaten()
	k = s.GibGroesse()
	h = k*433/500
	
	gfx.Volldreieck(x,y,x-k,y,x-(k/2),y-h) //Das Sechseck besteht aus sechs gleichseitigen Dreiecken, die alle einzeln gezeichnet werden.
	gfx.Volldreieck(x,y,x-k,y,x-(k/2),y+h)
	gfx.Volldreieck(x,y,x+k,y,x+(k/2),y-(h))
	gfx.Volldreieck(x,y,x+k,y,x+(k/2),y+(h))
	gfx.Volldreieck(x,y,x+(k/2),y+(h),x-(k/2),y+(h))
	gfx.Volldreieck(x,y,x+(k/2),y-(h),x-(k/2),y-(h))
}

func (s *data) GehörtPunktzuSechseck (xp, yp uint16) bool {
	var p1,p2,p3,p4,p5,p6 [2]uint16							//Eckpunkte des Sechsecks
	p1,p6,p3,p4 = s.GibLEcken()								//Liefert die Punkte 1,6,3 & 4
	p2,_,_,p5 = s.GibOEcken()								//Liefert die Punkte 2,3,6,5, es werden aber nur 2 & 5 benötigt
	var x [2]uint16											//Umwandlung der x- und Y-Kooerdinaten in einen Punkt ([2]uint16)
	x[0] = xp
	x[1] = yp
	return selbeSeite(p1,p2,p5,x) && selbeSeite(p2,p3,p6,x) && selbeSeite(p3,p4,p1,x) && selbeSeite(p4,p5,p2,x) && selbeSeite(p5,p6,p3,x) && selbeSeite(p6,p1,p4,x) //Test ob der Punkt auf der Innenseite aller Sechseckkanten befindet.
}

func selbeSeite (a,b,c,p [2]uint16) bool { //Mittels Vektorrechnung wird getestet, ob ein gegebener Punkt auf der selben Seite der Geraden Liegt wie ein anderer Punk
	var cp1,cp2 core.Vector
	var v1, v2, v3 core.Vector
	v1.X = float64(int(b[0])-int(a[0])) //Umwandlung der Punkte in die entsprechenden Vektoren 1.) Punkte a & b bilden die Gerade
	v1.Y = float64(int(b[1])-int(a[1]))
	v1.Z = 0
	v2.X = float64(int(c[0])-int(a[0]))	//Richtungs Vektor des zum Dreieck gehören den Punktes (oder in diesem Fall zum Sechseck gehörenden Punktes) 
	v2.Y = float64(int(c[1])-int(a[1]))
	v2.Z = 0
	v3.X = float64(int(p[0])-int(a[0]))	//Richtungs Vektor des zu testenden Punktes
	v3.Y = float64(int(p[1])-int(a[1]))
	v3.Z = 0
    cp1 = v1.Cross(v2)	//Erstellt den Vektor, der die Ebene identifiziert, die sich aus den beiden Vektoren ergibt (Gerade und zugehöriger Punkt)
    cp2 = v1.Cross(v3)  // Gerade und zu testender Punkt
    return cp1.Dot(cp2) >= 0 // Befinden sich die Punkte auf der selben Seite, so ist das Ergebnis größer gleich null, andernfalls kleiner als Null 
}
	
	
	
	
	
