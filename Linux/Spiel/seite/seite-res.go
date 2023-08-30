package seite

import (
		"../button"
		)

type Seite interface {
//Vor.: -
//Erg.: Ein Zeiger auf eine Seite mit dem gegeben Namen ist geliefert.
//New (name string) *data

//Vor.:-
//Eff.: Ein Button mit den übergeben Parametern ist der Seite hinzugefügt worden. Für die funktionalität der Knöpfe ist es wichtig, dass der Name des Buttons und der Name der Seite gleich sind.
	AddButton (resolution uint16,seitenname string,x,y uint16,name string,ox,oy int32)

//Vor.:
//Eff.: Der Button zu der angegeben Seite ist nicht auf der Seite vorhanden.
//Erg.: Der zu entfernede Button ist zurückgegeben, wenn ein Button gefunden wurde. Wenn kein Button gefunden wurde, dann ist ein Nullerbutton zurück gegeben (Seitenname,x,y,Text,ox,oy) ("",0,0,"",0,0)
	RemoveButton (sname string) button.Button

	
//Vor.:-
//Erg.: Der name der Seite ist geliefert.
	GetName () string

//Vor.:-
//Eff.:Der Name der Seite wird auf den gegeben Namen verändert.
	ChangeName (name string)

//Vor.: -
//Erg.: Ein Slice mit den Beschriftungstexten der zur Seite gehörenden Knöpfe ist geliefert.
	GetButtonsTexts () []string

//Vor.:Ein gfx-Fenster ist geöffnet mit der Größe Breite 1600 und Höhe 1000 
//Eff.: Die Seite ist im gfx-Fenster dargestellt.
	Draw (resolution uint16)

//Vor.: Ein gfx-Fenster ist geöffnet in der größe 1600,1000.
//Erg.: Der Name der nächste darzustellenden Seite ist geliefert.
//Eff.. Alle Buttons einer Seite sind dargestellt und werden bis ein Button gedrückt wurde mit Moushower animiert.
	PressButton (resolution uint16) (erg string)
	
//Vor.: Ein gfx-Fenster ist geöffnet in der größe 1600,1000.
//Eff.: Alle Buttons einer Seite sind im Fenster dargestellt.
	ShowButton(resolution uint16)

//Vor.: -
//Eff.: Ein gegebener Button ist der Seite hinzuzugefügt. Nullerbuttons werden ignoriert.	
	ReturnButton (resolution uint16,b button.Button)
}

//Vor.: -
//Erg.: Die gesuchte Seite ist aus einem Slice von Seiten ausgewählt und zurückgegeben. Ist die Seite in dem Slice nicht enthalten, so wird eine panic ausgelößt.
//Dies geschieht auch, wenn der Seiten-Slice leer ist. 
func SeitenAuswahl (sname string, seiten []Seite) Seite {
	for i:=0;i<len(seiten);i++{
		if seiten[i].GetName()==sname {
			return seiten[i]
		}
	}
	panic ("Gesuchte Seite ist keine bekannte Seite!!!")
} 

//Vor.: -
//Erg.: Ein Slice von Seiten mit samt der dazugehörenden Buttons ist geliefert. (Hier handelt es sich um die Seiten, die für das Startmenüe benötigt werden.
func SeitenGenerator (resolution uint16) []Seite{
	var willkommen,about,regeln1,regeln2,regeln3,regeln4,beenden, beenden2, guiSeite, konsoleSeite, highscoreSeite Seite //Generieren der Seiten
	willkommen = New("Willkommen")
	about = New("About")
	beenden2 = New("Beenden2")
	beenden = New("Beenden")
	regeln1 = New("Regeln1")
	regeln2 = New("Regeln2")
	regeln3 = New("Regeln3")
	regeln4 = New("Regeln4")
	guiSeite = New("GUI")
	konsoleSeite = New("Konsole")
	highscoreSeite = New("Highscore")

	
	willkommen.AddButton(resolution,"GUI",624,650,"GUI",-15,-10)			//Anhängen der Buttons
	willkommen.AddButton(resolution,"Konsole",712,700,"Konsole",-40,-10)
	willkommen.AddButton(resolution,"Regeln1",800,650,"Regeln",-35,-10)
	willkommen.AddButton(resolution,"Highscore",800,750,"High-\nscore",-27,-20)
	willkommen.AddButton(resolution,"About",888,700,"About",-28,-10)
	willkommen.AddButton(resolution,"Beenden",976,750,"Beenden",-40,-10)
	
	highscoreSeite.AddButton(resolution,"Willkommen",1300,790,"Zurück",-35,-10)
	
	beenden2.AddButton(resolution,"Willkommen",712,700,"Beenden",-39,-10)
	beenden2.AddButton(resolution,"zurück",888,700,"Zurück",-35,-10)
	
	about.AddButton(resolution,"Willkommen",1300,790,"Zurück",-35,-10)
	
	regeln1.AddButton(resolution,"Willkommen",1274,920,"Zurück",-35,-10)
	regeln1.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)
	regeln1.AddButton(resolution,"Regeln2",1450,920,"Weiter",-34,-10)
	
	regeln2.AddButton(resolution,"Regeln1",1274,920,"Zurück",-35,-10)
	regeln2.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)
	regeln2.AddButton(resolution,"Regeln3",1450,920,"Weiter",-34,-10)
	
	regeln3.AddButton(resolution,"Regeln2",1274,920,"Zurück",-35,-10)
	regeln3.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)
	regeln3.AddButton(resolution,"Regeln4",1450,920,"Weiter",-34,-10)
	
	regeln4.AddButton(resolution,"Regeln3",1274,920,"Zurück",-35,-10)
	regeln4.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)	

	
	var erg []Seite
	erg = append(erg,willkommen,regeln1,regeln2,regeln3,regeln4,beenden,beenden2,about,guiSeite,konsoleSeite, highscoreSeite) //Anhängen an den erg-Slice
	return erg
}

//Vor.: -
//Erg.: Ein Slice von Seiten mit samt der dazugehörenden Buttons ist geliefert. (Hier handelt es sich um die Seiten, die für das Startmenüe benötigt werden, die während dem SPielaufgerufen werdne kann.
func SeitenGenerator2 (resolution uint16) []Seite{
	var willkommen,about,regeln1,regeln2,regeln3,regeln4,beenden, beenden2,zurück,highscoreSeite Seite //Generieren der Seiten
	willkommen = New("Willkommen")
	zurück = New("Zurück")
	about = New("About")
	beenden2 = New("Beenden2")
	beenden = New("Beenden")
	regeln1 = New("Regeln1")
	regeln2 = New("Regeln2")
	regeln3 = New("Regeln3")
	regeln4 = New("Regeln4")
	highscoreSeite = New("Highscore")
	
	willkommen.AddButton(resolution,"Regeln1",800,650,"Regeln",-35,-10)		//Anhängen der Buttons
	willkommen.AddButton(resolution,"About",888,700,"About",-28,-10)
	willkommen.AddButton(resolution,"Beenden",976,650,"Beenden",-39,-10)
	willkommen.AddButton(resolution,"Zurück",536,700,"Zurück",-35,-10)
	willkommen.AddButton(resolution,"Highscore",800,750,"High-\nscore",-27,-20)
	
	beenden2.AddButton(resolution,"Willkommen",712,700,"Beenden",-39,-10)
	beenden2.AddButton(resolution,"zurück",888,700,"Zurück",-35,-10)
	
	about.AddButton(resolution,"Willkommen",1300,790,"Zurück",-35,-10)
	
	highscoreSeite.AddButton(resolution,"Willkommen",1300,790,"Zurück",-35,-10)
	
	regeln1.AddButton(resolution,"Willkommen",1274,920,"Zurück",-35,-10)
	regeln1.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)
	regeln1.AddButton(resolution,"Regeln2",1450,920,"Weiter",-34,-10)
	
	regeln2.AddButton(resolution,"Regeln1",1274,920,"Zurück",-35,-10)
	regeln2.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)
	regeln2.AddButton(resolution,"Regeln3",1450,920,"Weiter",-34,-10)
	
	regeln3.AddButton(resolution,"Regeln2",1274,920,"Zurück",-35,-10)
	regeln3.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)
	regeln3.AddButton(resolution,"Regeln4",1450,920,"Weiter",-34,-10)
	
	regeln4.AddButton(resolution,"Regeln3",1274,920,"Zurück",-35,-10)
	regeln4.AddButton(resolution,"Willkommen",1362,870,"Start",-28,-10)	

	
	var erg []Seite
	erg = append(erg,willkommen,regeln1,regeln2,regeln3,regeln4,beenden,beenden2,about,zurück,highscoreSeite) //Anhängen an den erg-Slice
	return erg
}



