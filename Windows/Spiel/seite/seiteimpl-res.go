package seite

import ("gfx2"
		"../button"
		"../plättchen"
		. "../zufallszahlen"
		)

//Eine Seite besteht aus dem Seitennamen und einem Slice von Buttons, die wiederum in der Regel Namen von Seiten zurückgeben.
//Die Verknüpfung zwischen Seite und Button kkommt dadurch zustande, dass der Seitennamen der im Button gespeicher ist und 
//der Seitenname der seite, zu der Button führen soll übereinstimmen.
//Da die Seitenauswahlfunktion bestimmt nach welchen kriterien auf die nächste Seite führen, so könnte dies theoretisch angepasst werden.
// Daher ist es wichtig, dass die Namen der Seiten und die in den Buttons gespeicherten Seitennamen nicht verändert werden.

type data struct {
	buttons []button.Button //Slice von Buttons mit denen man auf andere Seiten gelangen kann. (Text des Buttons muss dem Namen der Seite entsprechen, zu dem er führen soll)
	name string //Seitenname zur Identifizierung
}

func New (name string) *data {	//Trägt nur den Seitennamen ein
	var s *data
	s = new(data)
	(*s).name = name
	return s
}

func (s *data) AddButton (resolution uint16,seitenname string,x,y uint16,name string,ox,oy int32) { //Funktion, mit der sich neue Buttons auf der Seite angelegt werden können
		var bt button.Button
		bt = button.New(resolution,seitenname,x,y,name,ox,oy)										//Dem Button muss die resolution mitgegeben werden, da hier Koordinaten an das Sechseckweiter
																									//gegeben wird und auch die Schriftgröße und der Versatz der Schriftz auf dem Button angepasst werden 
		(*s).buttons = append((*s).buttons,bt)														//Der neue Button wird in den Slice buttons eingehangen (Reihenfolge der Buttons spielt keine Rolle
}

func (s *data) RemoveButton (sname string) button.Button{											//Entfernt einen Button - welcher Button entfernt wird bestimmt der gespeicherte Seitenname
																									//Inaktive Buttons (keinen Seitennamen) können so nur gemeinsam entfert werden
																									//-- irrelevant, da die "inaktiven Buttons" noch nicht ordentlich implementiert sind
	var nbuttons []button.Button																	//Es wird ein neuer Slice von Buttons generiert, damit der vorherige gespleist werden kann
	var erg button.Button																			//Der Entferte Button wird zurück gegeben, falls ergebraucht wird
	for i:=0;i<len((*s).buttons);i++ {																//Geht die Buttons-Liste durch und überprüft, ob einer mit dem entsprechenden Seitennamen enthalten ist
		if (*s).buttons[i].GetSeitenName() == sname {
			erg = (*s).buttons[i]																	//Ist das der Fall, dann wird der Button im erg-Button gespeichert  
			nbuttons = append(nbuttons,(*s).buttons[:i]...)											//und durch Spleisen aus dem ursprünglichen Slice entfernt
			nbuttons = append(nbuttons,(*s).buttons[i+1:]...)
			(*s).buttons = nbuttons																	//Anschleißend wird der so generierte Slice an die Seite zurück gegeben
		}
	}
	return erg																						//Der erg-Button wird zurückgegeben. Falls kein Button mit der Seite gefunden wurde, wird der Buttonslice der Seite nicht verändert und ein
																									//Nuller-Button zurückgegeben.
}

func (s *data) ReturnButton (resolution uint16, b button.Button) {									//Hängt ein gegebenen Butten an die Seite an, aber nur dann, wenn es sich nicht um ein Nullerbutton handelt.
	var nb button.Button
	nb = button.New(resolution,"",0,0,"",0,0)
	if !button.IstGleich(nb,b) {
		(*s).buttons = append((*s).buttons,b)
	}
}

func (s *data) ChangeName (name string) {															//Verändert den Namen einer Seite
	(*s).name = name
}

func (s *data) GetButtonsTexts () []string {														//Liest die Beschriftung der Buttons aus und gibt sie in einem Slice zurück
	var erg []string
	for i:=0;i<len((*s).buttons);i++{
		erg = append(erg,(*s).buttons[i].GetText())
	}
	return erg
}

func (s *data) Draw (resolution uint16) {  															//Zeichen Funktion für Seiten
	if (*s).name!="Beenden" && (*s).name!="Konsole"{												//Beenden und Konsole führt zum schließen des gfx2-Fensters, daher wird hier keine Seite gezeichnet
		gfx2.Stiftfarbe(255,255,255)																	//Ein weißer Hintergrund wird generiert 
		gfx2.Cls()
		if (*s).name!="Regeln1" && (*s).name!="Regeln2" && (*s).name!="Regeln3" && (*s).name!="Regeln4" && (*s).name!="Spielfeld" { //Diese seiten brauchen keinen aus Plättchen generierten Rand
			drawHintergrund (resolution)															//Zeichen des Plättchen randes
		}
	}
}


func (s *data) GetName () string {																	//Gibt den Namen der Seite zurück
	return (*s).name
}



func (s *data) PressButton (resolution uint16) (erg string) {										//PressButton organisiert das gesamte Buttonhandling vom Zeichnen, über Highlighting und dem Auslesen
																									//des Seitennamens, wenn ein Button gedrückt wird 
	var x,y uint16
	var taste uint8
	var status int8
	for i:=0;i<len((*s).buttons);i++{																//Wird die Funktion aufgerufen, so werden zunächst alle Buttons in ihre nicht gehighlightete Version überführt
		if (*s).buttons[i].HighlightGesetzt()==true{
			(*s).buttons[i].SwitchBackGround()
		}
	}
	gfx2.UpdateAus()																					//Alle Buttons der Seite werden gezeichnet
	for i:=0;i<len((*s).buttons);i++{
		(*s).buttons[i].Draw(resolution)
	}
	gfx2.UpdateAn()
	for {																							//In einer Endlosschleife wird das Highlighting an und ausgeschaltet für jeden Button der Seite
		taste,status,x,y = gfx2.MausLesen1()
		for i:=0;i<len((*s).buttons);i++{
			switch {
				case (*s).buttons[i].HighlightGesetzt() == false && (*s).buttons[i].GehörtPunktzuButton (x,y) == false:
					continue
				case (*s).buttons[i].HighlightGesetzt() == true && (*s).buttons[i].GehörtPunktzuButton (x,y) == false:
					(*s).buttons[i].SwitchBackGround ()
					gfx2.UpdateAus()
					(*s).buttons[i].Draw(resolution)
					gfx2.UpdateAn()
					continue
				case (*s).buttons[i].HighlightGesetzt() == false && (*s).buttons[i].GehörtPunktzuButton (x,y) == true:
					(*s).buttons[i].SwitchBackGround ()
					gfx2.UpdateAus()
					(*s).buttons[i].Draw(resolution)
					gfx2.UpdateAn()
					continue
				case (*s).buttons[i].HighlightGesetzt() == true && (*s).buttons[i].GehörtPunktzuButton (x,y) == true:
					continue
			}
		}
		if taste==1 && status==1 {																	//bis ein Button gedrückt wird.
			for i:=0;i<len((*s).buttons);i++{
				if (*s).buttons[i].GehörtPunktzuButton (x,y) {
						erg = (*s).buttons[i].GetSeitenName()										//Dann wird der Seitenname, der im Buttongespeichert ist zurückgegeben und die Funktion beendet. 
							return
				}
			}
		}
	}
}


func (s *data) ShowButton (resolution uint16) {														//Funktion zum Zeichnen von Buttons - wird genutzt, wenn die Buttons der Seite zusehen sein sollen, diese aber keine Funktion haben
	for i:=0;i<len((*s).buttons);i++{
		if (*s).buttons[i].HighlightGesetzt()==true{												//Auch heir wird zunächst das Highlighting entfernt
			(*s).buttons[i].SwitchBackGround()
		}
	}
	for i:=0;i<len((*s).buttons);i++{																//Anschließend alle Buttons gezeichent
		(*s).buttons[i].Draw(resolution)
	}
}


func drawHintergrund (resolution uint16) {															//Dies Funktion generiert einen leicht zufällige Verteilung von Spielplättchen um den Rand einer Seite.
	var p [27]plättchen.Plättchen
	p = plättchen.PlättchenGenerator()
	gfx2.Stiftfarbe(255,255,255)
	gfx2.Cls()
	var i int
	for j:=0;j<36;j++{  //Es sind 18 Kästchen des Spielrandes definiert, in die jeweils zufällig ausgewählt von den 27 Plättchen zwei Plättchen an zufällige Koordinate gelegt wird.
		Randomisieren()
		var x,y uint16
		switch i {																			//Definiert die Kooerdinaten der gedachten Kästchen begrenzt durch den Range der gewählten Zufallszahlenbereiche
			case 0:
				x = uint16(Zufallszahl(int64(51*resolution/10),int64(240*resolution/10)))
				y = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
			case 1:
				x = uint16(Zufallszahl(int64(260*resolution/10),int64(500*resolution/10)))
				y = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
			case 2:
				x = uint16(Zufallszahl(int64(520*resolution/10),int64(760*resolution/10)))
				y = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
			case 3:
				x = uint16(Zufallszahl(int64(780*resolution/10),int64(1020*resolution/10)))
				y = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
			case 4:
				x = uint16(Zufallszahl(int64(1040*resolution/10),int64(1280*resolution/10)))
				y = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
			case 5:
				x = uint16(Zufallszahl(int64(1300*resolution/10),int64(1549*resolution/10)))
				y = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
			case 6:
				x = uint16(Zufallszahl(int64(51*resolution/10),int64(240*resolution/10)))
				y = uint16(Zufallszahl(int64(900*resolution/10),int64(949*resolution/10)))
			case 7:
				x = uint16(Zufallszahl(int64(260*resolution/10),int64(500*resolution/10)))
				y = uint16(Zufallszahl(int64(900*resolution/10),int64(949*resolution/10)))
			case 8:
				x = uint16(Zufallszahl(int64(520*resolution/10),int64(760*resolution/10)))
				y = uint16(Zufallszahl(int64(900*resolution/10),int64(949*resolution/10)))
			case 9:
				x = uint16(Zufallszahl(int64(780*resolution/10),int64(1020*resolution/10)))
				y = uint16(Zufallszahl(int64(900*resolution/10),int64(949*resolution/10)))
			case 10:
				x = uint16(Zufallszahl(int64(1040*resolution/10),int64(1280*resolution/10)))
				y = uint16(Zufallszahl(int64(900*resolution/10),int64(949*resolution/10)))
			case 11:
				x = uint16(Zufallszahl(int64(1300*resolution/10),int64(1549*resolution/10)))
				y = uint16(Zufallszahl(int64(900*resolution/10),int64(949*resolution/10)))
			case 12:
				x = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
				y = uint16(Zufallszahl(int64(120*resolution/10),int64(370*resolution/10)))
			case 13:
				x = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
				y = uint16(Zufallszahl(int64(390*resolution/10),int64(630*resolution/10)))
			case 14:
				x = uint16(Zufallszahl(int64(51*resolution/10),int64(100*resolution/10)))
				y = uint16(Zufallszahl(int64(650*resolution/10),int64(880*resolution/10)))
			case 15:
				x = uint16(Zufallszahl(int64(1500*resolution/10),int64(1549*resolution/10)))
				y = uint16(Zufallszahl(int64(120*resolution/10),int64(370*resolution/10)))
			case 16:
				x = uint16(Zufallszahl(int64(1500*resolution/10),int64(1549*resolution/10)))
				y = uint16(Zufallszahl(int64(390*resolution/10),int64(630*resolution/10)))
			case 17:
				x = uint16(Zufallszahl(int64(1500*resolution/10),int64(1549*resolution/10)))
				y = uint16(Zufallszahl(int64(650*resolution/10),int64(880*resolution/10)))
			}
			p[Zufallszahl(0,26)].Draw(resolution,x,y)                                  	//Hierfindet die zufällige Auswahl der Plättchen und ihre Darstellung statt
		i = (i+1)%18																	//Dies sorgt dafür, dass jedes gedachte Kästchen zwei Plättchen erhällt.	
	}
}
