package sechseck


type Sechseck interface {

//Vor.:-
//Erg.: Ein Zeiger auf ein Secheck ist geliefert, mit den Kooerdinaten x,y (Mittelpunkt), der Breite 2k und der Höhe 2k*433/500 und der Farbe (r,g,b) - Sechseck liegt auf der flachen Seite   
//New(x,y,k uint16, r,g,b uint8)

//Vor.: -
//Eff.: Das Sechseck besitz jetzt die übergeben Farbe (r,g,b).
	SetzeFarbe (r,g,b uint8)

//Vor.: -
//Eff.: Das Sechseck besitzt jetzt die neuen übergebenen Koordinaten (x,y).
	SetzeKoordinaten (x,y uint16)

//Vor.: -
//Eff.: Das Sechseck besitzt jetzt die neue Breite 2k und die neue Höhe 2k*433/500
	SetzeGroesse (k uint16)

//Vor.: -
//Erg.:
	GibFarbe () (r,g,b uint8)

//Vor.: -
//Erg.: Die Koordinaten des Mittelpunnkt des Sechsecks sind geliefert.
	GibKoordinaten () (x,y uint16)

//Vor.: -
//Erg.: Die Größe k ist geliefert, die die größe des Sechsecks bestimmt. (Breite 2k, Höhe 2k*433/500; Ausdehnung des Sechsecks vom Mittelpunkt (x,y) je k in x-Dimension und je k*433/500 in y-Dimension)
	GibGroesse () (k uint16)
	
//Vor: -
//Erg.: Die Koordinaten der Eckpunkte der linken unteren Kante und rechten oberen Kante an.	
	GibLEcken () (lu1,lu2,ro1,ro2 [2]uint16)

//Vor: -
//Erg.: Die Koordinaten der Eckpunkte der oberen Kante und unteren Kante an.
	GibOEcken () (o1,o2,u1,u2 [2]uint16)

//Vor: -
//Erg.: Die Koordinaten der Eckpunkte der linken oberen Kante und rechten unteren Kante an.
	GibREcken () (lo1,lo2,ru1,ru2 [2]uint16)

//Vor.: -
//Erg.: True ist geliefert, wenn die übergebenen Kooerdinaten sich innerhalb des Sechsecks befinden. False ist geliefert, wenn die Kooerdinaten ausserhalb liegen.
	GehörtPunktzuSechseck (xp, yp uint16) bool

//Vor.: Ein gfx-Fenster ist geöffnet. Das Sechseck befindet sich vollständig im Fenster: Mittelpunkt linker bzw. rechter Rand: x  +/- (k+/-1); oberer bzw. unterer Rand y  +/- (k*433/500 +/- 1).
//Eff.: Das Sechseck ist als Vollsechseck im gfx-Fenster dargestellt.
	Draw ()
}

