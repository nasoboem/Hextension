package plättchen

import ("fmt"
		"../sechseck"
		"strconv"
		"gfx2"
		)

type data struct {
	//ID uint8			//Plättchen ID geht von 1-27 - Kann auch durch die Positions im feld selbst 
	l uint8				//Links dargestellte Zahl kann die Werte 2,6,7 annehmen 
	o uint8				//Oben dargestellte Zahl kann die Werte 1,5,9 annehmen
	r uint8				//Rechts dargestellte Zahl kann die Werte 3,4,8 annehmen
}
/*
func (p *data) String() string {														//Für die Showfunktion sind nur die Zahlen in l,o und r gespeichert wichtig
	var erg string																		//Ergebnisstring Variable wird initialisiert
	erg = "("+fmt.Sprint((*p).l)+","+fmt.Sprint((*p).o)+","+fmt.Sprint((*p).r)+")"		//Die Werte, die im Plättchen gespeichert werden werden in einen String "(l,o,r)" umgewandelt.
	return erg																			//String wird zurückgegeben
}
*/

func New(l,o,r uint8) *data {					//Ein neues Plättch wird gegeben durch seine es definierenden Werte für l, o und r
	var p *data
	p = new(data)
	(*p).l = l
	(*p).o = o
	(*p).r = r
	return p 
}

func (p *data) GetValues () (l,o,r uint8) { //Gibt die Werte für l, o und r zurück, die im Plättchen gespeichert sind
	l = (*p).l
	o = (*p).o
	r = (*p).r
	return
}

func (p *data) String() string { //Ist die gleiche Stringfunktion wie oben, nur dass Plättchen mit dem Wert (0,0,0) nicht dargestellt werden (mit Leerzeichen aufgefüllt, dargestellt werden)
	var erg string
	var l,o,r uint8
	l,o,r = p.GetValues()
	if l==0 && o==0 && r==0 {
		erg = "       "
	} else {
		erg = "("+fmt.Sprint(l)+","+fmt.Sprint(o)+","+fmt.Sprint(r)+")"
	}
	return erg
}

func streifen (p1, p2, p3, p4 [2]uint16, abstand uint16) (dreieck1, dreieck2 [6]uint16) { 	//Diese Funktion bekommt die Eckpunkte des Sechseck entweder für l, o oder r und den Abstand (Dicke des Streifens und soll die beiden dazu notwendigen Dreiecke ermitteln.
																							//Die Geradengleichungen der gegenüberliegenden Seiten müssen ermittelt werden und dann von der Seitenhalbierenden soll der Abstand in beide Richtungen dazu addiert werden
	var mp1, mp2 [2]uint16																	//Berechnung des Mittelpunkts der beiden Seiten des Sechsecks
	mp1[0] = (p2[0]+p1[0])/2
	mp1[1] = (p2[1]+p1[1])/2
	mp2[0] = (p4[0]+p3[0])/2
	mp2[1] = (p4[1]+p3[1])/2
	if abstand==0 {																			//Bei Abstand == 0 wird die Linie der Seitenhalbierenden ausgegeben
		dreieck1[0] = mp1[0]
		dreieck1[1] = mp1[1]
		dreieck1[2] = mp1[0]
		dreieck1[3] = mp1[1]
		dreieck1[4] = mp2[0]
		dreieck1[5] = mp2[1]
		dreieck2[0] = mp2[0]
		dreieck2[1] = mp2[1]
		dreieck2[2] = mp2[0]
		dreieck2[3] = mp2[1]
		dreieck2[4] = mp1[0]
		dreieck2[5] = mp1[1]
		return
	} else {													//Der X-Koordinate wird der Abstand drauf gerechnet oder abgezogen und anschließend die richtige y-Koordinate berechnet
		var versatz1,versatz2, steigung float64
		steigung = (float64(p2[1])-float64(p1[1]))/(float64(p2[0])-float64(p1[0]))
		versatz1 = ((float64(p2[0])*float64(p1[1]))-(float64(p1[0])*float64(p2[1])))/(float64(p2[0])-float64(p1[0]))
		versatz2 = ((float64(p4[0])*float64(p3[1]))-(float64(p3[0])*float64(p4[1])))/(float64(p4[0])-float64(p3[0]))
		dreieck1[0] = mp1[0]-abstand					//x1									
		dreieck1[1] = uint16(steigung*float64(dreieck1[0])+versatz1)		//y1
		dreieck1[2] = mp1[0]+abstand					//x2
		dreieck1[3] = uint16(steigung*float64(dreieck1[2])+versatz1)		//y2
		dreieck1[4] = mp2[0]-abstand					//x3
		dreieck1[5] = uint16(steigung*float64(dreieck1[4])+versatz2)		//y3
		
		dreieck2[0] = mp2[0]-abstand
		dreieck2[1] = uint16(steigung*float64(dreieck2[0])+versatz2)
		dreieck2[2] = mp2[0]+abstand
		dreieck2[3] = uint16(steigung*float64(dreieck2[2])+versatz2)
		dreieck2[4] = mp1[0]+abstand
		dreieck2[5] = uint16(steigung*float64(dreieck2[4])+versatz1)
		return
	}
}
	
		
func (p *data) Draw (resolution uint16,x,y uint16) { //Zeichnet die Plättchen
	var l,o,r uint8
	l,o,r = p.GetValues()
	var s sechseck.Sechseck
	s = sechseck.New(x,y,5*resolution,160,160,160)
	if l==0 {							//Wenn es sich um ein Nuller-Plättchen handelt, dann wird die Farbe des Plättchen auf dunkelgrau gesetzt
		s.SetzeFarbe(60,60,60)
	}
	s.Draw()							//Zu erst wird ein Sechseck gezeichnet
	var dreieck1, dreieck2 [6]uint16	// dann die Streifen, die jeweils aus zwei Dreiecken bestehen
	var p1,p2,p3,p4 [2]uint16
	p1,p2,p3,p4 = s.GibLEcken()			//Aus den Ecken der Seiten für links unten und rechts oben werden die Werte für die beiden Dreiecke ermittelt und anschließend gezeichnet
	if resolution < 5 {					//Die Breite der Streifen wird nach unten begrenzt, weil sonst nur eine einfacher Strich übrig bleiben würde
		dreieck1, dreieck2 = streifen(p1,p2,p3,p4,1)
	} else {
		dreieck1, dreieck2 = streifen(p1,p2,p3,p4,2*resolution/10)
	}
	switch {							//Wahl der Stiftfarbe in Abhängigkeit der Werte des Plättchens - bei Nullerplättchen werden die Streifen in der Plättchenfarbe gezeichnet, sie bleiben unsichtbar
		case l==2:
		gfx2.Stiftfarbe (255,0,255)
		case l==6:
		gfx2.Stiftfarbe (255,0,0)
		case l==7:
		gfx2.Stiftfarbe (0,160,0)
		default:
		gfx2.Stiftfarbe (60,60,60)
	}
	gfx2.Volldreieck(dreieck1[0],dreieck1[1],dreieck1[2],dreieck1[3],dreieck1[4],dreieck1[5])		//Streifen von links unten nach recht oben wird gezeichnet
	gfx2.Volldreieck(dreieck2[0],dreieck2[1],dreieck2[2],dreieck2[3],dreieck2[4],dreieck2[5])
	
	p1,p2,p3,p4 = s.GibREcken()					//Prozess wiederholt sich für den Streifen von links oben nach rechts unten
	if resolution < 5 {
		dreieck1, dreieck2 = streifen(p1,p2,p3,p4,1)
	} else {
		dreieck1, dreieck2 = streifen(p1,p2,p3,p4,2*resolution/10)
	}
	switch {
		case r==3:
		gfx2.Stiftfarbe (0,255,255)
		case r==4:
		gfx2.Stiftfarbe (0,0,255)
		case r==8:
		gfx2.Stiftfarbe (255,160,0)
		default:
		gfx2.Stiftfarbe (60,60,60)
	}
	gfx2.Volldreieck(dreieck1[0],dreieck1[1],dreieck1[2],dreieck1[3],dreieck1[4],dreieck1[5])
	gfx2.Volldreieck(dreieck2[0],dreieck2[1],dreieck2[2],dreieck2[3],dreieck2[4],dreieck2[5])
	p1,p2,p3,p4 = s.GibOEcken()										//Prozess wiederholt sich für den senkrechten Streifen 
	dreieck1, dreieck2 = streifen(p1,p2,p3,p4,4*resolution/10)
	switch {
		case o==1:
		gfx2.Stiftfarbe (100,100,100)
		case o==5:
		gfx2.Stiftfarbe (160,0,160)
		case o==9:
		gfx2.Stiftfarbe (255,255,0)
		default:
		gfx2.Stiftfarbe (60,60,60)
	}
	gfx2.Volldreieck(dreieck1[0],dreieck1[1],dreieck1[2],dreieck1[3],dreieck1[4],dreieck1[5])
	gfx2.Volldreieck(dreieck2[0],dreieck2[1],dreieck2[2],dreieck2[3],dreieck2[4],dreieck2[5])
	if l!=0 && o!=0 && r!=0 {										//Beschriftung der Plättchen findet nur statt, wenn es kein Nullerplättchen ist
		var k uint16
		k = s.GibGroesse()											//je nach Größe des Plättchen
		gfx2.Stiftfarbe(0,0,0)
		gfx2.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(30*resolution/10))
		var text string
		text = strconv.Itoa(int(l))									
		gfx2.SchreibeFont(x-2*k/3,y+k/10,text)						// Lage der Zahl kommt es zu einem unterschiedlichen Versatz der Schrift zum Mittelpunkt
		text = strconv.Itoa(int(o))
		gfx2.SchreibeFont(x-k/6,y-(k*433/500)+k/25,text)
		text = strconv.Itoa(int(r))									//für l,o und r müssen die Zahlen zuerst in einen string umgewandelt werden
		gfx2.SchreibeFont(x+k/3,y+k/10,text)							//und dann im gfx2-fenster angezeigt werden
	}
}


	
	
	
	 
	  


	

