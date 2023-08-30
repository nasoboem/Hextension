package spiel

import ("../seite"
		"../spielfeld"
		"../gfx"
		"../feld"
		. "../zufallszahlen"
		"fmt"
		"../felder"
		"strconv"
		"../plättchen"
		"../button"
		"../highscore"
		"../dateien"
		)


func Spielen(resolution uint16) {

//Start-Menü

	var seiten []seite.Seite																			//Erzeugt einen Slice von Seiten, der dann die Seiten vom Seiten Generator aufnimmt (Start-Menü und seine Seiten).
	seiten = seite.SeitenGenerator(resolution)															
	var aktSeite seite.Seite																			//Eine Variable wird erzeugt, die immer die anzuzeigende Seite enthält
	gfx.Fenster(160*resolution,100*resolution)															//gfx-Fenster geöffnet
Anf:aktSeite = seite.SeitenAuswahl("Willkommen",seiten)													//Die erste Seite wird übergeben																
	for aktSeite.GetName()!="GUI" && aktSeite.GetName()!="Beenden" && aktSeite.GetName()!="Konsole" {	//Schleife, die einen durch das Start-Menü führt - Aufruf der Seiten GUI, Konsole oder Beenden - beenden das Start-Menü 
		gfx.UpdateAus()																					//Schaltet das Update aus, für ein flackerfreie darstellung der Seiten
		aktSeite.Draw(resolution)																		//Darstellung der Seiten
		seitenInhalt(resolution,aktSeite)																//Darstellung des Seiteninhalts 
		gfx.UpdateAn()																					//Anschalten des Updates					
		aktSeite = seite.SeitenAuswahl(aktSeite.PressButton(resolution),seiten)							//PressButton - Kümmert sich um die Darstellung und das Buttonhandling; SeitenAuswahl - Wählt die nächste Seite aus je nach dem Welcher Button gedrückt wurde.
	}
	
//Start-Menü beendet, Spiel beginnt

//Festlegung des Darstellungsmodus
																										 
	if aktSeite.GetName() != "Beenden" {											//Wenn im Start-Menü nicht auf Beenden gedrückt wurde, dann wird das Spiel gestartet
		var inKonsole bool															//Darstelluns Variable, die entscheidet, wie das Spiel dargestellt wird. InKonsole == true - dann wird das Spiel in der Konsole dargestellt, wenn inKonsole == false, dann wird es in der GUI dargestellt
		if aktSeite.GetName() == "Konsole" {
			inKonsole = true														//Umstellen auf Konsolen-Darstellung
			gfx.FensterAus()														//Ausschalten des Fensters der GUI
		}
		
		
//Erstellung des Spiels und Vorinitialisierung der Spielfelder und Spielernamen

		var game [19]uint8															//Initialisierung der Variablen game(spiel) - Ein Spiel zufällig generiert und bestimmt in welcher Reihenfolge welches Plättchen gezogen wird 
		game = Spiel()																//Weist game durch die Funktion Spiel() ein Feld aus Zufallszahlen [19]uint8 aus 0-26 zu
		var sfs []spielfeld.Spielfeld												//Ein Slice von Spielfeldern wird generiert
		var spielernamen []string													//Der Slice für die Spielernamen wird initialisiert, durch die Eingabe beschränkt auf 6

//Abfrage der Spieleranzahl und Spielernamen
		
		if inKonsole {																				
			spielernamen = spielereingabeKonsole ()									//Eingabe der Spielernamen über die Konsole											
		} else {
			spielernamen = spielereingabe(resolution)								//Eingabe der Spielernamen über die GUI
		}
		
//Erstellen eines Spielfeldes pro Spieler 
			 
		for i:=0;i<len(spielernamen);i++ {											//Schleife generiert ein Spielfeld pro Spieler und versetzt es in den Anfangszustand (besetzen des aktuellen Plättchens)
			var aktSpielfeld spielfeld.Spielfeld									//Ein neues Spielfeld wird initialisiert
			aktSpielfeld = spielfeld.New(resolution,spielernamen[i])				//und mit konkreten Werten belegt							
			aktSpielfeld.PlättchenZiehen(game)										//Das Aktuelle Plättchen wird belegt
			sfs = append(sfs,aktSpielfeld)											//Die Spielfelder werden an den Slice aus Spielfeldern angehangen
		}
		
		
//Eigentliche Spielmechanik beginnt

		for i:=0;i<len(game);i++ {													//Schleife, die das Spiel in Zügen durchläuft (geht die 19 Züge des Spielfelds durch)
			for j:=0;j<len(sfs);j++{												//Schleife, die die Spielfelder durchläuft in jedem Zug
			
//Der Spieler wählt ein Feld, auf dass er das Plättchen legt
				
				var f feld.Feld														//Variable für die Feld-Wahl des Spielrs wird generiert
				f = feld.New()
				
//Feldeingabe an der Konsole
				
Kon:			if inKonsole {														//Spielmechanik für die Konsole und Einspringpunkt für den Wechsel der Darstellungsformen - Nur erlaubt vor der Auswahl des Feldes
					fmt.Println(sfs[j]) 											//Darstellen des Spielfeldes vor der Wahl des Feldes - Damit der Spieler sich überlegen kann welches Feld er wählt 
					var guiTrue, printSF, quit bool									//Steuervariablen - guiTrue für den wechsel in graphische Benutzeroberfläche, printSF um das Spielfeld erneut anzeigen zu lassen und quit um das Spiel zu beenden

A:					f, guiTrue,printSF, quit = feldWahl()							//Einspringpunkt, wenn das Spielfeld erneut angezeigt wurde, oder ein bereits besetztes Feld gewählt wurd 
					if quit  {														//Beenden des Spieles bei Eingabe q
						var abfrage string
						fmt.Println("")
						fmt.Print("Wollen Sie wirklich das Spiel beenden? (y/n): ")	//Bestätigung zur Beendigung des Spiels
						fmt.Scanln(&abfrage)
						if abfrage=="y" {											//Alle anderen Eingaben führen automatisch zur wieder Aufnahme des Spiels nicht nur n
							gfx.Fenster(160*resolution,100*resolution)				//Öffnen des gfx-Fensters
							goto Anf												//Sprung zum Startbildschirm
						}
					}
					if guiTrue {													//Wechsel zur gui bei Einagbe g
						inKonsole = false
						gfx.Fenster(160*resolution,100*resolution)
						goto Kon													//Übergang zur GUI, über die erneute If-Abfrage
					}
					if printSF {													//Erneutes darstellen des Spielfeldes - Wird benötigt, wenn durch ungünstige Eingaben das Spielfeld auf dem Bildschirm nicht mehr zu sehen ist
						fmt.Print(sfs[j])
						goto A
					}

					if sfs[j].FeldFrei(f) {											//Wenn das gewählte Feld frei ist, dann wird das Feld gesetzt
						sfs[j].FeldSetzen(f)
					} else {														//Ansonsten wird der Spieler auf seinen Fehler hingewiesen und der Feldauswahl prozess beginnt von vorn
						fmt.Println("")
						fmt.Println("Dieses Feld ist bereits belegt!!! Sie müssen ein anderes wählen.")
						goto A
					}
					
//Feld ist ausgewählt - jetzt muss der Spieler es noch bestätigen, dass auch das Plättchen an diese Position gelegt wird, andern falls kann der Spieler sich noch um entscheiden (Feldauswahl beginnt von vorn)					
B:					var abfrage string			
					fmt.Println("")
					fmt.Print("Wollen Sie die das Plättchen wirklich auf das Feld ",f," legen? (y/n):")
					fmt.Scanln(&abfrage)
					if abfrage=="n" {						//Bei n - nein: möchte der Spieler ein anderes Feld wählen --> zurück zur Feldauswahl
						goto A
					} else if abfrage=="y" {				//Bei y - ja: der Spieler will das Plättchen an diese Stelle legen, dann werden alle entsprechenden änderungen am Spielfeld vorgenommen
						aktSeite = sfs[j].GibSeite1()		//Damit man frei zwischen gui und konsole wechseln kann müssen die Buttons der Legefläche entfernt werden, auf die Plättchen gelegt werden, sonst werden die liegenden Plättchen von Buttons übermalt
						aktSeite.RemoveButton(fmt.Sprint(f))//fmt.Sprint(f) - Greift auf die Stringfunktion des Feldes zurück und gibt den String aus, was dem "Seitenname" des zu entfernenden Buttons entspricht
						sfs[j].SetzeSeite1(aktSeite)		//Gibt die veränderte Seite 1 ans Spielfeld zurück
						aktSeite = sfs[j].GibSeite2()		//Das gleiche für Seite 2 
						aktSeite.RemoveButton(fmt.Sprint(f))
						sfs[j].SetzeSeite2(aktSeite)
						sfs[j].PlättchenSetzen(game)
					} else if abfrage=="s"{					//Bei s wird das Spielfeld erneut ausgegeben --> Erneute Abfrage der Bestätigung
						fmt.Print(sfs[j])
						goto B	
					} else {								//Bei allen anderen Eingaben wird der Spieler auf die richtige Eingabe hingewiesen --> Erneute Abfrage der Bestätigung
						fmt.Println("")
						fmt.Println("Machen Sie eine korrekte Eingabe: y - Bestätigung des Feldes; n - Neue Feldwahl und s - Spielfeld ausgeben.")
						goto B
					}
					
//Feldeingabe in der GUI

				} else {
Gui:				var buttonReturn string					//Rückgabevariable für die Aufnahme der Button-Rückgabe

//Darstellen des Spielfeldes
					gfx.UpdateAus()
					sfs[j].Draw(resolution)					//Darstellung des Spielfeldes
					sfs[j].DrawLegefläche(resolution)		//Darstellung der bereits gelegten Plättchen
					gfx.UpdateAn()
//Darstellen und Aktivieren der Buttons
					aktSeite = sfs[j].GibSeite1()			//Übergabe der 1. Seite des Spielfeldes
					buttonReturn = aktSeite.PressButton(resolution) //Start des Buttonhandlings mit PressButton und Entgegennahme des Rückgabewerts des gedrückten Buttons
					
					
//Abfangen aller Buttons, die nicht die Legefläche ausmachen - Beenden, Konsole, Menü
					
//Beenden - Zeigt zu erst die "Wirklich-Beenden-Seite" und bei weiterem bestätigen wird das Spiel beendet --> Startseite, sonst wird geht man zum Anfang der Darstellung des Spielfeldes zurück
					if buttonReturn == "Beenden2" {
						aktSeite = seite.SeitenAuswahl(buttonReturn,seiten)
						gfx.UpdateAus()												//Schaltet das Update aus, für ein flackerfreie darstellung der Seiten
						aktSeite.Draw(resolution)									//Darstellung der Seiten
						seitenInhalt(resolution,aktSeite)							//Darstellung des Seiteninhalts
						gfx.UpdateAn()
						buttonReturn = aktSeite.PressButton(resolution)				//Darstellung und Aktivierung der Buttons
						if buttonReturn == "Willkommen" {							//Im Falle der Bestätigung
							goto Anf						//--> Startseite
						} else {													//Im Falle der Verneinung
							goto Gui						//--> Einspringpunkt der Rundendarstellung in der Gui - Spielfelddarstellung
						}
						
//Konsole - Wechseln der Darstellungsform des Spiels - inKonsole wird auf true gesetzt, das Fenster geschlossen und an den Anfang der Konsolendarstellung des Spielfeldes gesprungen						
					} else if buttonReturn == "Konsole"{
						inKonsole = true
						gfx.FensterAus()							//Fenster schließen
						goto Kon									//--> Einspringpunkt - if inKonsole - jetzt true
						
//Menü - Ein modifiziertes Startmenü wird aufgerufen, in dem kein neues Spiel gestartet werden kann, sondern über einen Zurück-Button wieder zum Spiel zurück gekehrt werden kann - Alle weiteren Funktionen beleiben erhalten						
					} else if buttonReturn == "Menü" {
						var seiten2 []seite.Seite
						seiten2 = seite.SeitenGenerator2(resolution)						//Übergabe des alternativen Startmenüs
Menue2:					aktSeite = seite.SeitenAuswahl("Willkommen",seiten2)				//Aufruf der Startseite, des alternativen Startmenüs
						for aktSeite.GetName()!="Zurück" && aktSeite.GetName()!="Beenden" {	//Schleife, die einen durch das alternative Start-Menü führt - Drücken der Buttons Zurück und Beenden führen aus der Schleife raus 
							gfx.UpdateAus()													//Schaltet das Update aus, für ein flackerfreie darstellung der Seiten
							aktSeite.Draw(resolution)										//Darstellung der Seiten
							seitenInhalt(resolution,aktSeite)								//Darstellung der Seiteninhalte
							gfx.UpdateAn()													//Anschalten des Updates							
							aktSeite = seite.SeitenAuswahl(aktSeite.PressButton(resolution),seiten2)	//Verwaltet die Buttons und das weiterleiten durch Klicken auf die nächste Seite
						}
						if aktSeite.GetName()=="Beenden" {
							aktSeite = seite.SeitenAuswahl("Beenden2",seiten)				//Ruft die Wirklich-Beenden-Seite auf, da das Spiel beendet werden würde
							gfx.UpdateAus()													//Ablauf wie beim vorherigen Beenden nur mit anderem Ausstiegspunkt									
							aktSeite.Draw(resolution)
							seitenInhalt(resolution,aktSeite)
							gfx.UpdateAn()
							buttonReturn = aktSeite.PressButton(resolution)
							if buttonReturn == "Willkommen" {
								goto Anf								//--> wirkliche Startseite des Spiels
							} else {
								goto Menue2								//--> Startseite des alternativen Startmenüs
							}
						}
						goto Gui										//--> Einspringpunkt der Rundendarstellung in der Gui - Spielfelddarstellung									

//Ende des Firlefanz, zurück zur eigentlichen Spielmechanik

//Feld der Legefläche wurde gedrückt - Übergang von Seite 1 des Spielfeldes zur Seite 2 des Spiefeldes - hier kann sich der Spieler beliebig lange umentscheiden, bis er auf "Weiter" drückt

					} else {
						var vorherigerButton button.Button				//Ein Nullerbutton wird generiert - dieser wird von der Funktion ReturnButton ignoriert, wodurch im ersten Durchgang nichts passiert
						vorherigerButton = button.New(resolution,"",0,0,"",0,0)
						var lastButtonReturn string											//Diese Variable wird benötigt, da der letzte Button der gedrückt wird immer der "Weiter"-Button ist, man aber noch den Return des letzten Buttons der Legefläche braucht		
						for buttonReturn!="Weiter" {										//Die Feldabfrage wird solange wiederholt, bis der Weiter-Button gedrückt wird
							sfs[j].GibSeite2().ReturnButton(resolution,vorherigerButton) 	//Beim umentscheiden muss der Button, der vorher entfernt wurde zurück getan werden (nur beim ersten mal nicht!!)
							f = stringToFeld(buttonReturn)									//Der String des buttonReturns wird in ein Feld umgewandelt 
							sfs[j].FeldSetzen(f)											//Das vom Spieler gewählte Feld wird zum aktuellen Feld auf dem Spielfeld gemacht
							gfx.UpdateAus()
							sfs[j].Draw(resolution)											//Das Spielfeld wird gezeichnet
							sfs[j].DrawLegefläche(resolution)								//Die bisher gelegten Plättchen werden gezeichnet
							sfs[j].PlättchenAnzeigen(resolution)							//Das aktuelle zuspielende Plättchen wird an die gewählte stelle gezeichnet
							gfx.UpdateAn()
							aktSeite = sfs[j].GibSeite2()									//Von der 2. Seite wird der Button der Legefläche entfernt, der mit dem vom Spieler gewählten Feld korespondiert
							vorherigerButton =aktSeite.RemoveButton(buttonReturn)			//und dieser Button wird in vorherigerButton abgelegt
							lastButtonReturn = buttonReturn									//Übergabe der Inhalte von Buttonreturn an lastButtonReturn (erfolgt nur, wenn der Button nicht "Weiter" war wie in der for-Schleife spezifiziert 
							buttonReturn = aktSeite.PressButton(resolution)					//Darstellung und Aktivierung der Buttons der 2.Seite des Spielfeldes (Buttons verändern sich je nach dem welches Feld der Spieler wählt 
						}
						sfs[j].SetzeSeite2(aktSeite)										//Die 2.Seite wird so wie sie aus der Feldwahl des Spielers herauskommt dem Spielfeld zurück gegeben
						aktSeite = sfs[j].GibSeite1()										//Die Änderungen der Buttons der Legefläche müssen jetzt auch noch an der 1.Seite des Spielfeldes vorgenommen werden - 1.Seite wird geholt
						aktSeite.RemoveButton(lastButtonReturn)								//der entsprechende Button wird entfernt (Hier wird die Variable lastButtonReturn benötigt!!)
						sfs[j].SetzeSeite1(aktSeite)										//Die veränderte Seite wird dem Spielfeld zurück gegeben
						sfs[j].PlättchenSetzen(game)										//Das Plättchen wird final in das Spielfeld eingetragen, mit allen Veränderungen, die mit dioesem Eintrag einhergehen (Zug++, belegteFelder,gelegtePlättchen, neues aktuellesPlättchen, neuer Punktestand, Plättchen in die Reiheneintragen) 
														//Spielfeld des nächsten Spielers wird aufgerufen j++
					}
														//Nächster Zug wird aufgerufen i++
				}										
			}
		}
		
//Alle Spieler haben das Spiel beendet und die Spielfelder enthalten den Endpunktestand
		
//Spielergebnis - Die Namen der Spieler und ihre Punkte werden in zwei Slices absteigendt sortiert
		var punkte []int
		_,spielernamen,punkte = reihenfolgePunkte(sfs)	//Sortierung - Mergesort übergibt das Spielergebnis (Namen & Punkte) als zwei getrennte Slices nach Punkten absteigend sortiert (Entspricht Endergebnis diesen Spiels!!)
		var newHighscore highscore.Highscore			//Ein neuer Highscore wird generiert
		newHighscore = highscore.New()
		var datei1,datei2 dateien.Datei					//Neue Highscore-Dateien werden angelegt sollten sie nicht existieren 
		datei1 = dateien.Oeffnen("highscore-namen.txt",'x')
		datei2 = dateien.Oeffnen("highscore-punkte.txt",'x')
		if datei1.Groesse()==0 || datei2.Groesse()==0 {
			datei1.Schliessen()
			datei2.Schliessen()
			newHighscore.Speichern()					//und der neue Highscore in ihnen gespeichert
		} else {
			datei1.Schliessen()							//wenn sie schon existierten passiert nichts
			datei2.Schliessen()
		}
		newHighscore.Lesen()							//Hier könnte es zu Fehlern kommen, wenn die Dateien verändert würden und kein richtiger Highscore entsteht
		var plaetze []int								//Initialisierung der Variable, die das Platzierungshighlighting im Highscore ermöglicht - Indexzahlen von Einfügen
		var platz int									//Variable 
		for i:=0;i<len(spielernamen);i++{				//geht die Spielernamen/Punkte durch und fügt sie in den Highscore ein 
			platz = newHighscore.Einfügen(spielernamen[i],punkte[i])
			plaetze = append(plaetze,platz)				//Generiert ein Slice in denen die Indexzahlen der neu in den Highscore eingefügten Einträge stehen (max. 6 Einträg mit Index 0-5; wenn der Highscore erreicht wurde und 7 wenn er nicht erreicht wurde) 
		}
		newHighscore.Speichern()						//Der neue Highscore wird in die Dateien geschrieben
		
//Darstellung der Spielergebnisse in der Konsole
		if inKonsole {
			var tabelle string
			fmt.Println("")
			fmt.Println("- - Endergebnis - -")
			fmt.Println("")
			tabelle = endTable(spielernamen,punkte)		//endTabel produziert ein String aus den Spielergebnissen, die in den zwei Slices spielernamen und unkte gespeichert ist
			fmt.Println(tabelle)
			fmt.Println("")
			fmt.Print("Bitte drücken Sie die Enter-Taste! ")
			var dump string
			fmt.Scanln(&dump)							//Eine Eingabe wird abgewartet, so dass der Spieler sich die Ergebnisse angucken kann
// Darstellung der Highscoreliste
			fmt.Println("")
			fmt.Println("- - Highscore - -")
			fmt.Println("")
			fmt.Println(newHighscore)					//Der neue Highscore wird ausgegeben - hier könnte man noch ein Highlighting der Plätze einfügen
			fmt.Println("")
			fmt.Print("Bitte drücken Sie die Enter-Taste! ")
			fmt.Scanln(&dump)							//ebenso wird hier eine Eingabe abgewartet
			gfx.Fenster(160*resolution,100*resolution)	//Öffnen des gfx-Fensters
			goto Anf									// --> Startbildschirm
			
			
//Darstellung der Spielergebnisse im GUI

		} else {
			var ergebnisSeite, highscoreSeite seite.Seite	//Die Highscore und Ergebnisseiten werden erstellt - sollte eigendlch ausgelagert werden
			ergebnisSeite = seite.New("Ergebnis")
			highscoreSeite = seite.New("Highscore")
			ergebnisSeite.AddButton(resolution,"Highscore",1300,790,"Weiter",-35,-10)	//Die Seiten bekommen ihre Buttons
			highscoreSeite.AddButton(resolution,"Willkommen",1300,790,"Weiter",-35,-10)
			gfx.UpdateAus()
			ergebnisSeite.Draw(resolution)												//Ergebnis-Seite wird gezeichnet
			gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(6*resolution))
			gfx.Stiftfarbe(0,0,0)
			gfx.SchreibeFont(60*resolution,18*resolution,"Endstand")					//Überschrift
			gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(4*resolution))
			gfx.SchreibeFont(40*resolution,30*resolution,"Platz")						//Tabellenüberschriften
			gfx.SchreibeFont(67*resolution,30*resolution,"Spieler")
			gfx.SchreibeFont(100*resolution,30*resolution,"Punkte")
			gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(3*resolution))
			for i:=uint16(0);i<uint16(len(spielernamen));i++ {							//Inhalt der Tabelle wird geschrieben
				gfx.SchreibeFont(43*resolution,39*resolution+(7*resolution*i),fmt.Sprint(i+1)+".")
				gfx.SchreibeFont(70*resolution,39*resolution+(7*resolution*i),spielernamen[i])
				gfx.SchreibeFont(103*resolution,39*resolution+(7*resolution*i),fmt.Sprint(punkte[i]))
			}
			gfx.UpdateAn()
			ergebnisSeite.PressButton(resolution)										//Buttons der Ergebnisseite werden dargestellt und aktiviert
			
//Darstellung des Highscores
			gfx.UpdateAus()
			highscoreSeite.Draw(resolution)												//Highscore-Seite wird gezeichnet
			gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(6*resolution))
			gfx.Stiftfarbe(0,0,0)
			gfx.SchreibeFont(59*resolution,18*resolution,"Highscore")					//Überschrift
			newHighscore.Draw(resolution,plaetze)										//Highscore mit Neueeintragshighlighting wird dargestellt
			highscoreSeite.PressButton(resolution)										//Buttons der Highscore-Seite werden dargestellt und aktiviert - Hier würde jeder Button zum Startbildschirm führen, da das Ergebnis nicht netgegen genommen wird.
			goto Anf																	//Spiel ist beendet --> Startbildschirm					
		}
	}
}


func spielereingabeKonsole () (erg []string) {
			var spielerAnzahl int																	//Initialisierung der Variablen, die die Anzahl der Spieler aufnimmt
			var saString string																		//Initialisierung der Variablen, die den String der Eingabe Übernimmt
			var err error																			//Variable, um den Fehler von strconv.Atoi abzufangen
SekA:		fmt.Print("Geben Sie die Anzahl der Spieler ein (max. 6): ")							//Abfrage der Spieleranzahl; Einspringpunkt, wenn der Spieler falsche Eingaben tätigt
			fmt.Scanln(&saString)
			spielerAnzahl,err = strconv.Atoi(saString)
			if err!=nil {
				goto SekA																			//Wiederholung der Abfrage, wenn ein Fehler bei strconv.Atoi auftritt
			}
			if spielerAnzahl<1 || spielerAnzahl>6 {
				goto SekA																			//Wiederholung der Abfrage, wenn andere Werte eingegeben werden als 1-6
			}
			var spieler string																		//Initialisierung der Variablen Spieler, die die Spielernamen entgegen nimmt
			for i:=0;i<spielerAnzahl;i++{															//Schleife, die Spielernamen erfragt
				fmt.Println("")																		//Einfügen einer Leerzeile, für die bessere Lesbarkeit
				fmt.Print("Geben Sie den Namen für Spieler ",i+1," ein (keine Leerzeichen && max. 15 Zeichen!!!): ") //Beschränkung auf 15 Zeichen ist nur für die Gui Endergebnisausgabe und Highscore wichtig
SekB:			fmt.Scanln(&spieler)																//Sollte der Spieler sich nicht an die Längenangabe gehalten haben wird hier wieder Eingesprungen
				var zähler int
				for range (spieler) {																//Eingaben länge wird in der Forschleife gezählt, range wird benutzt, das es tatsächlich buchstaben sind
					zähler++
				}
				if zähler>15 {																		//Falls der Spieler sich nicht dran gehalten hat, dann gibt es einen Hinweis für den Spieler und eine erneutes Auslesen des Spielernamens
					fmt.Println("")
					fmt.Print("Bitte geben Sie einen kürzeren Namen für Spieler ",i+1," ein (max. 15 Zeichen): ")
					goto SekB																		//Wiederholung der Spielereingabe
				}
				erg = append(erg,spieler)															//Einfügen des Spielernamens in den Slice der Spielernamen
			}
			return																					//Slice der Spielernamen wird zurückgegeben
}

func Spiel () [19]uint8 {			//Ein Spiel besteht aus 19 zufällig gezogenen, sich nicht wiederholende Zahlen aus dem Zahlenraum von 0-26 - Sie dienen als Abfolge der gezogenen Plättchen (Abfolge, ist die Reihenfolge der Zahlen und die Plättchen werden identifiziert durch die Zahlen, die gespeichert sind)
	var erg [19]uint8				//Ergebnisvektor der gezogenen Plättchen IDs
	Randomisieren()					//Zufallszahlen werden durch die Systemzeit mit einem neuen Seed versehen
	for i:=0;i<len(erg);i++{		//Schleife, die durch die Positionen des Ergebnisvektors durchgeht um Zufällig die Plättchen-ID von 1-27 durchzugehen
A:		var zz int64				//Veriable, die die Zufallszahl (zz) speichert
		zz = Zufallszahl(0,26)		//Zuweiseung der zz
		for j:=0;j<i;j++ {			//überprüfung ob die zz bereits im Ergebnisvektor enthalten ist, in dem über eine Schleife die bisherigen Elemente in erg mit zz verglichen wird
			if erg[j]!=uint8(zz){	//Vergleich des aktuellen Elements in Erg mit zz (Typanpassung von zz auf uint8!)
				continue			//bei Ungleichheit wird die nächste Zahl in Erg mit zz verglichen
			} else {
				goto A				//bei Gleichheit muss eine neue zz gewählt werden, sprung zu A
			}
		}
		erg[i] = uint8(zz)			//ist zz noch nicht in erg enthalten, so wird zz an der Position i eingefügt
	}
	return erg						//Der Ergebnisvektor erg mit den zu Spielenden Plättchen-IDs wird zurückgegeben
}


func spielereingabe (resolution uint16) []string { //Spielereingabe für die GUI - stammt noch aus Zeiten vor ADT-Seite, daher nicht mit Seite implementiert
	gfx.UpdateAus()
	gfx.Stiftfarbe(255,255,255)						//Hintergrund wird weiß gemacht
	gfx.Cls()
	drawHintergrund(resolution)						//Plättchenrand wird gezeichnet
	gfx.UpdateAn()
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(4*resolution)) //Texte mit Anweisungen an den Spieler werden dargestellt
	gfx.SchreibeFont(25*resolution,20*resolution,"Wählen Sie die Anzahl der Spieler aus (max. 6)")	
	gfx.SchreibeFont(35*resolution,275*resolution/10,"und geben Sie sie nacheinander ein.")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(4*resolution))
	gfx.SchreibeFont(598*resolution/10,45*resolution,"Anzahl der Spieler")
	var f felder.Feld														//Zugriff auf den ADT-Felder von Herr Schäfer
	f = felder.New(int(787*resolution/10),int(365*resolution/10),1,'z',"")	//Das Feld für die Eingabe der Anzahl der Spieler wird Positioniert und seine Eigenschaften, wie Eingabelänge und Orientirung festgelegt
	f.SetzeZeichensatzgroesse (groessenauswahl(resolution))					//Zum einstellen der Schriftgröße wurde eine eigene Funktion geschrieben, da Herr Schäfer seine eigenen Standardgrößen definiert und nur indirekt auf gfx zurückgreift
	f.SetzeErlaubteZeichen("123456")										//Setzen der erlaubten Zeichen
	var eingabeAnz string													//Initialisierung der Variable, die den eingegeben String entgegen nimmt
	eingabeAnz = ""															//Wird auf einen leeren String gesetzt
	for eingabeAnz == "" {													//Abfangen leerer eingaben - führte früher zum Absturz
		eingabeAnz = f.Edit()												//Übernimmt das ganze Eingabefeldhandling + Übergabe des Strings an die Variable 
	}
	var anzahlSpieler int
	anzahlSpieler,_ = strconv.Atoi(eingabeAnz)								//Umwandeln des Strings in eine Zahl
	var spieler []string													//Namens-Slice wird generiert
	gfx.Archivieren()														//Das gfx-Fenster wird gespeicher, damit es für jede Spielereingabe neu restauriert werden kann 
	var g felder.Feld														//Einagbefeld für die Spieler namen wird generiert
	for i:=0;i<anzahlSpieler;i++ {											//Eingabe wird so oft wiederholt, bis alle Spielernamen eingegeben sind
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(4*resolution))
		gfx.SchreibeFont(598*resolution/10,625*resolution/10,"Name von Spieler " + fmt.Sprint(i+1) ) //fmt.Sprint wandelt i+1 in ein String um
		g = felder.New(int(657*resolution/10),int(55*resolution),12,'z',"") //Das Feld für die Eingabe der Spielernamen wird Positioniert und seine Eigenschaften, wie Eingabelänge und Orientirung festgelegt
		g.SetzeZeichensatzgroesse (groessenauswahl(resolution))				//Zeichensatz an die Auflösung angepasst
		var a string
		a = ""
		for a == "" {														//Abfangen von leeren Eingaben
			a = g.Edit()													//Eingabefeldhandling  (s.o.)
		}
		spieler = append(spieler,a)											//Spielernamen weden an den spieler-Slice angehangen
		if !(i==(anzahlSpieler-1)) {										//Wenn es nicht der letzte Spieler ist, dann
			g.Leere()														//Wird das Eingabefeld geleer
			gfx.UpdateAus()
			gfx.Stiftfarbe(255,255,255)
			gfx.Cls()														//das ganze Fenster weiß gemacht
			gfx.Restaurieren(1,1,160*resolution,100*resolution)				//das ganze Fenster resauriert und es beginnt die nächste Spielernameneingabe
			gfx.UpdateAn()
		}
	}
	return spieler															//Spielernamen-Slice wird zurückgegeben
}

func groessenauswahl (resolution uint16) uint16 {	//Diese Funktion ist nur dazu dar, die von Herrn Schäfer verwendeten Schriftgrößenauf meine Auflösungen zu "mapen" - Ich habe den ADT-Felder dahingehend verändert, das er jetzt mehr Schriftgrößen besitzt und die Breite entsprechend so angepasst, das die Buchstaben nicht mehr aus dem Feld laufen.
	switch resolution {
		case 1:			//Meine Auflösung
			return 12	//Zu verwendende Schriftgröße
		case 2:
			return 12
		case 3:
			return 12
		case 4:
			return 16
		case 5:
			return 20
		case 6:
			return 24
		case 7:
			return 28
		case 8:
			return 32
		case 9:
			return 36
		case 10:
			return 40
		default:
			return 40
	}
}

//Wandelt ein String in ein Feld um - da Buttons nur Strings ausgeben, müssen diese wieder in Felder überführt werden,
// wenn ein Feld auf dem Spielfeld angesprochen werden soll 
func stringToFeld (stf string) feld.Feld { 
	var f feld.Feld						//ein neues Feld wird initialisiert
		f = feld.New()
	if len(stf)<3 {						//Es wird geschaut, ob die Länge des Strings den mindest anforderungen entspricht (Länge min. 3)
		panic ("Der Übergebene String ist zu kurz!! Erwarte 3-2!!") //panic wenn dem nicht so ist
	} else {
		var spalte, zeile uint8				//Variablen für spalte und zeile werden initialisiert
		var sInt,zInt int					//Variablen für die spalte und zeile als int-Wert
		var errs,errz error					//Variablen für den Error der jeweiligen Umwandlungen
		sInt,errs = strconv.Atoi(stf[:1])	//Konvertierung des ersten bytes des Strings
		zInt,errz = strconv.Atoi(stf[2:3])	//Konvertierung des dritten bytes des Strings --> 3ä4 wäre kein korrekte Eingabe, da ä mehr als ein byte verbraucht, hier würde das zweite byte von ä konvertiert werden
		if errs==nil && errz==nil {			//wenn es bei der umwandlungen keine Fehler gab
			spalte = uint8(sInt)			//werden die int-Werte in uint8-werte überführt und in die entsprechenden Variablen für spalte und zeile eingetragen
			zeile = uint8(zInt)
		} else {
			panic ("Eingabe konnte nicht in Int-Werte überführt werden!!!")	//panic falls es Fehler gab
		}
		f.SetzeFeld(spalte,zeile)			//Feld wird entsprechen von Spalte und Zeile verändert
	}
	return f								//und zurück gegeben
}


//feldWahl (für das spielen in der Konsole) erlaubt es ein Feld auszuwählen, in dem der Nutzer eine Eingabe aus Spalte und Zeile macht;
//Weitere sinnvolle Eingaben sind q für Beenden, g für den wechsel in das GUI s um das Spielfeld erneut auszugeben.
func feldWahl () (f feld.Feld, guiTrue bool, printSF bool, quit bool) {	
C:	var eingabe string									//Funktionsstart --> EInspringpunkt, wenn die Eingabe wiederholt werden muss.
	f = feld.New()										//Ein neues Nullerfeld wird als Ausgabe definiert 
	fmt.Println("")										//Leerzeilen für die bessere leesbarkeit
	fmt.Println("")
	fmt.Print("Wählen Sie ein Feld - Bsp. 3-5:")
	fmt.Scanln(&eingabe)								//Abfrage der Spielereingabe
	if len(eingabe)==0{									//Hat die Eingabe die Länge 0 - aka Spieler hat keine Ahnung was ermachen soll
		fmt.Println("")
		fmt.Println("Machen Sie eine Eingabe: q - Spiel beenden; s - Spielfeld ausgeben; g - zur grafischen Benutzeroberfläche wechseln, 3-2 - wählt das entsprechende Feld.")
		goto C											//Sagen dem Spieler was er machen kann und wiederholen die Eingabeabfrage
	} else if len(eingabe)<3{							//Ist die Länge der Eingabe kleiner als 3
		var selektor string
		selektor = string(eingabe[0])					//Wenn der Anfang der Einagbe ein ... war
		if selektor=="q" {								//Bei der Eingabe von "q" wird die Variable quit auf true gesetzt --> Auswahl: Spiel beenden
			quit = true
			return
		} else if selektor=="s" {						//Bei der Eingabe von "s" wird die Variable printSF auf true gesetzt --> Auswahl: Spielfeld ausgeben
			printSF = true
			return
		} else if selektor=="g" {						//Bei der Eingabe von "g" wird die Variable guiTrue auf true gesetzt --> Auswahl: 
			guiTrue = true
			return
		} else {
			goto C										//trifft keines der Fälle zu, wennd  die Eingabe kleiner als 3 ist wird die Eingabeabfrage wiederholt 
		}
	}
 
	var spalte,zeile string																			//Variablen für Spalte und Zeile werden initialisiert
	spalte = string(eingabe[0])																		//Index 0
	zeile = string(eingabe[2])																		//und Index 2 des Eingabestrings werden ausgelesen
	var spalteZahl, zeileZahl int 																	//Vorinitialisierung der Variablen, die die Zahl für Spalte und Zeile aufnehmen sollen
	var errs,errz error																				//Fehlerkontrolle der Konvertierung, ein Fehler (!nil) lässt die Variable correct auf false (errs - Fehler für Spalte; errz - Fehler für Zeile)
	spalteZahl, errs = strconv.Atoi(spalte)															//Konvertierung von Spalte zu int
	zeileZahl, errz = strconv.Atoi(zeile)															//Konvertierung von Zeile zu int
	if errs == nil && errz == nil && feld.IstKorrekteEingabe(uint8(spalteZahl),uint8(zeileZahl)) {	//Überprüft auf Fehler (errs & errz) und auf das richtige Eingabeintervall
		f.SetzeFeld(uint8(spalteZahl),uint8(zeileZahl))												//Feld wird auf die eingegeben Werte gesetzt
		return																						//und zurückgegeben
	} else {																						//Tritt irgend ein Fehler auf, dann
		fmt.Println("")
		fmt.Println("Machen Sie eine korrekte Eingabe: q - Spiel beenden; s - Spielfeld ausgeben; g - zur grafischen Benutzeroberfläche wechseln, 3-2 - wählt das entsprechende Feld.")
		goto C																						//Sag dem Spieler was er machen kann und wiederholen die Eingabeabfrage
	}
}


//Funktion bekommt alle Spielfelder und gibt die Spielergebnisse sortiert nach erreichten Punkten absteigend in 3 Slices aus - Einem Index-Slice enthält die Spielernummer,
// Namens-Slice enthält die Spielernamen und dem Punkte-Slice enthält die erreichten Punkte 

func reihenfolgePunkte (sf []spielfeld.Spielfeld) (indexReihenfolge []int, namenReihenfolge []string, punkteReihenfolge []int) {
		var endPunkte []int
	var sPunkte int
	for i:=0;i<len(sf);i++ {							//Sammelt alle Ergebnisse (Punktestände) in der Reihenfolge der Spielfelder ein 
		sPunkte = sf[i].GibPunkte()
		endPunkte = append(endPunkte,sPunkte)			//und speichert sie in einem Slice endPunkte
	}
	indexReihenfolge = append(indexReihenfolge,0)				//Eintragen der ersten Werte in die Slices für die Endergebnisse
	punkteReihenfolge = append(punkteReihenfolge,endPunkte[0])
	for i:=1;i<len(endPunkte);i++{								//Geht die Endpunktestände durch und mittels Mergsort werden sie in die richtige Reihenfolge gebracht
		for j:=0;j<len(punkteReihenfolge);j++{					//Geht durch die bereitsbestehende punkteReihenfolge durch und schaut, wo der aktuelle Endpunktestand endPunkte[i] einsortiert werden soll
			if endPunkte[i] <= punkteReihenfolge[j] {			//Wenn der Wert in endPunkte kleiner/gleich der Punkte in der punkteReihenfolge ist
				if j == len(punkteReihenfolge)-1 {				//dann wird getested ob das der letzte Eintrag in der punkteReihenfolge ist
					punkteReihenfolge = append(punkteReihenfolge,endPunkte[i])	//wenn ja, dann werden die Werte hinten angehangen 
					indexReihenfolge = append(indexReihenfolge,i)
					break									//und die j-Schleife beendet
				} else {
					continue								//ist das Ende des Slices punkteReihenfolge nicht erreicht und endPunkte kleiner, gehe weiter
				}
			} else {										//Wenn der Wert in endPunkte größer ist ... 
					var zS1,zS2 []int
					for k:=0;k<j;k++ {						     //Dann nehme alle Werte an denen du bereits vorbeigekommen bist (0 bis j-1) und kopiere sie in ein neues Slice 
						zS1 = append(zS1,punkteReihenfolge[k])
						zS2 = append(zS2,indexReihenfolge[k])
					}
					zS1 = append(zS1,endPunkte[i])        		//Füge den einzufügenden Wert an der stelle j ein
					zS2 = append(zS2,i)
					for k:=j;k<len(punkteReihenfolge);k++{		//hänge alle anderen Werte die noch im ursprünglichen Slice enthalten sind hinten in den neuen Slice ein 
						zS1 = append(zS1,punkteReihenfolge[k])
						zS2 = append(zS2,indexReihenfolge[k])
					}
					punkteReihenfolge = zS1						//Übergebe die so verlängerten Slices an die ursprünglichen zurück 
					indexReihenfolge = zS2
					break										//Da der Wert eingefügt wurde wird die j-Schleife aufgelöst und die i-Schleife rückt eins weiter
			}
			
		}												 
	}
	for i:=0;i<len(indexReihenfolge);i++{						//Die indexReihenfolge wird in die namenReihenfolge überführt
		namenReihenfolge = append(namenReihenfolge,sf[indexReihenfolge[i]].GibSpielername())
	}
	return	//Die drei generierten Slices werden zurück gegeben.
}

//Generiert einen String, der die formatierte Ergebnistabelle des Spiels enthält

func endTable (namen []string, punkte []int) (erg string) {
	var nn []string								//Slice von Strings, der die Namen aufnimmt, die mit Leerzeichen aufgefüllt
	var lname uint								//Variable, die die Länge des längsten Namen aufnimmt
	lname,nn = nameSize(namen)					//nameSize liefert die Länge des längsten Namen und die Namensliste mit Leerstellen aufgefüllt
	erg = erg + fmt.Sprint(" Platz   Spieler")	//Beginn des Aufbaus des Strings - Kopfzeile
	for i:=0;i<(int(lname)-7);i++ {				//Angepasst an die länge der Namen
		erg = erg + fmt.Sprint(" ")
	}
	erg = erg + fmt.Sprintln("   Punkte")		//Ende der Kopfzeile
	länge:=len(erg)								//Übergabe der Gesamtlänge
	for i:=0;i<länge-1;i++{						//Einziehen der Zwischenzeile für die bessere Lesbarkeit
		erg = erg + "+"
	}
	erg = erg + fmt.Sprintln("")				//Ende der Zwischenzeile
	for i:=0;i<len(nn);i++{						//Einfügen der Tabelleninhalt für jeden Spieler
		erg = erg + fmt.Sprint("   ")			//Einrücken der Plazierungsnummer
		erg = erg + fmt.Sprint(i+1)				//Plazierungnummer
		erg = erg + fmt.Sprint(".  |  ")		//Einfügen von Punkt und Trennlinie
		erg = erg + fmt.Sprint(nn[i])			//Einfügen des Spielernamens
		if lname<7 {							//Ausgleich, wenn alle Namen in der Liste kürzer sind als 7 
			for i:=uint(0);i<7-lname;i++{
				erg = erg + fmt.Sprint(" ")
			}
		}
		erg = erg + fmt.Sprint(" |  ")			//Einfügen der zweiten Trennlinie
		erg = erg + fmt.Sprintln(punkte[i])		//Einfügen der Punkte
	}											//Ende der Tabelleninhalte
	for i:=0;i<länge-2;i++{						//Abschließende Zwischenzeile
		erg = erg + "+"
	}
	return										//Rückgabe des Tabellenstrings
} 
	
	
//Funktion bekommt ein Slice mit verschiedenlangen Namen und gibt die Länge des längsten Namen und den Namens-Slice mit Leerzeichen aufgefüllt zurück 	
func nameSize (namen []string) (uint, []string) {
	var max uint								//Nimmt die maximale Länge auf
	var count []uint							//Nimmt die differenz zwischen Max.Länge und der Länge des Strings an dem Index auf
	var erg []string							//Neuer Ergebnis-String in dem alle Strings gleich lang sind
	if namen==nil {								//Abfangen der leeren Liste
		return max, erg
	}
	for i:=0;i<len(namen);i++{					//Zählt die Buchstaben für jeden Spielernamen
		var zähler uint
		for range namen[i] {
			zähler++
		}
		count = append(count,zähler)			//und speichert die Anzahl in dem count-Slice
	}
	max = count[0]								//Max wird mit dem ersten Wert von count befüllt
	for i:=1;i<len(count);i++ {					//Geht count durch und sucht die größte Zahl
		if max>=count[i] {
			continue
		} else {
			max = count[i]
		}
	}											//Jetzt enthält max die Länge des längsten Namens
	for i:=0;i<len(namen);i++{					//For-Schleife generiert das Ergebnis-Slice
		var lname string			
		lname = namen[i]						//Einem String wird der Name übergeben
		for j:=uint(0);j<(max-count[i]);j++ {
			lname = lname + " "					//dann werden soviele Leerzeichen angefügt, wie die Differenz aus max und der Länge des jeweiligen Namens
		}
		erg = append(erg,lname)					//Die so auf die gleiche Länge gebrachte String wird an den Ergebnis-Slice angehangen
	}
	return max, erg								//Rückgabe der Werte
}

func drawHintergrund (resolution uint16) {  					//Kommentar s. ADT-Seite
	var p [27]plättchen.Plättchen
	p = plättchen.PlättchenGenerator()
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	var i int
	for j:=0;j<36;j++{
		Randomisieren()
		var x,y uint16
		switch i {
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
			p[Zufallszahl(0,26)].Draw(resolution,x,y)
		i = (i+1)%18
	}
}


// Seiteninhalte - Zuordnung -----------------------------------------------------------------------------------

func seitenInhalt (resolution uint16, s seite.Seite) { //Ordnet den Seiten nach ihrem Namen die richtigen Inhalte zu 
	var name string
	name = s.GetName()
	switch  name {
			case "Willkommen":
				drawBegrüßungsText(resolution) //Aufruf der Funktionen, die die Inhalte darstellen s.u.
			case "Regeln1":
				regeln1(resolution)
			case "Regeln2":
				regeln2(resolution)
			case "Regeln3":
				regeln3(resolution)
			case "Regeln4":
				regeln4(resolution)
			case "About":
				aboutText(resolution)
			case "Beenden":
			gfx.FensterAus()
			case "Beenden2":
				beenden2(resolution)
			case "Highscore":
				highscoreSeite(resolution)
	}
}

// Seiteninhalte - Inhalte --------------------------------------------------------------------------------------------------

func aboutText (resolution uint16) { //Liefert den den formatierten Text, der auf der About-Seite ausgegeben wird
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(6*resolution))
	gfx.Stiftfarbe(0,0,0)
	gfx.SchreibeFont(20*resolution,15*resolution,"Hextension")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(32*resolution/10))
	gfx.SchreibeFont(20*resolution,24*resolution,"Das Spiel")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(20*resolution,28*resolution,"Hextension ist ein Brettspiel für eine prinzipiell unbegrenzte Anzahl von Spie-")
	gfx.SchreibeFont(20*resolution,305*resolution/10,"lern. Es wurde 1983 unter dem Namen Hextension von Peter Burley entwickelt und von")
	gfx.SchreibeFont(20*resolution,33*resolution,"Spear-Spiele vertrieben. Die Grafik der Originalversion gestaltete Franz Vohwinkel.")
	gfx.SchreibeFont(20*resolution,355*resolution/10,"Unter dem Namen Take-it-easy wird das Spiel seit 1994 vom Verlag F.X. Schmid, jetzt")
	gfx.SchreibeFont(20*resolution,38*resolution,"Ravensburger, vertrieben.")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(32*resolution/10))
	gfx.SchreibeFont(20*resolution,44*resolution,"Das Programm")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(20*resolution,48*resolution,"Das vorliegende Programm wurde von Jakob Hättig im Rahmen der Lehrerweiterbildung")
	gfx.SchreibeFont(20*resolution,505*resolution/10,"entwickelt. Der unbedarfte Programmierer sei gewarnt. Man sagt, dass selbst der")
	gfx.SchreibeFont(20*resolution,53*resolution,"Blick in die Implementierung einens einfachen ADTs, einen den Verstand kosten kann.")
	gfx.SchreibeFont(20*resolution,555*resolution/10,"Wer es wagen sollte gar Hand an den Quellcode zu legen, der muss damit rechnen")
	gfx.SchreibeFont(20*resolution,58*resolution,"so zu enden wie der Autor.")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(32*resolution/10))
	gfx.SchreibeFont(20*resolution,63*resolution,"Der Programmierer")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(20*resolution,67*resolution,"Jakob Hättig, so sagt man, war ein einfacher Lehrer für Chemie und Biologie,")
	gfx.SchreibeFont(20*resolution,695*resolution/10,"bis er anfing in Go Hextension zu programmieren. Am Ende wiederholte er")
	gfx.SchreibeFont(20*resolution,72*resolution,"immer nur diesen einen Satz:")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(35*resolution/10))
	gfx.SchreibeFont(20*resolution,755*resolution/10,"\"Haben Sie es ausprobiert!\"")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(20*resolution,80*resolution,"Dann verschwand er eines Tages und ward nicht mehr gesehen.")
}

func drawBegrüßungsText (resolution uint16) { //Liefert die Inhalte der Startseite
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(4*resolution))
	gfx.Stiftfarbe(0,0,0)
	gfx.SchreibeFont(638*resolution/10,42*resolution,"      by      ")
	gfx.SchreibeFont(638*resolution/10,485*resolution/10," Jakob Hättig ")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(10*resolution))
	gfx.Stiftfarbe(0,0,0)
	gfx.SchreibeFont(502*resolution/10,25*resolution,"Hextension")
	var inaktiveButton1, inaktiveButton2 button.Button
	inaktiveButton1 = button.New(resolution,"",624,650,"GUI",-15,-10)		//Die inaktiven Buttons verstecken sich im Startmenü unter den tatsächlichen Buttons
	inaktiveButton2 = button.New(resolution,"",712,700,"Konsole",-39,-10)	//sie sind nur im alternativen Startmenü zu sehen, das man aus dem Spielheraus erreicht
	inaktiveButton1.Draw(resolution)
	inaktiveButton2.Draw(resolution)
}

func highscoreSeite (resolution uint16) {		//Liefert den Inhalt der Highscore-Seite
	var newHighscore highscore.Highscore
	var plaetze []int
	newHighscore = highscore.New()
	var datei1,datei2 dateien.Datei				//Konstrukt fängt den Fall nicht existenter Dateien ab - s. Spielen() 
	datei1 = dateien.Oeffnen("highscore-namen.txt",'x')
	datei2 = dateien.Oeffnen("highscore-punkte.txt",'x')
	if datei1.Groesse()==0 || datei2.Groesse()==0 {
		datei1.Schliessen()
		datei2.Schliessen()
		newHighscore.Speichern()
	} else {
		datei1.Schliessen()
		datei2.Schliessen()
	}
	newHighscore.Lesen()						//Einlesen des Highscores
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(6*resolution))
	gfx.Stiftfarbe(0,0,0)
	gfx.SchreibeFont(59*resolution,18*resolution,"Highscore")
	newHighscore.Draw(resolution,plaetze)		//Darstellen des Highscores
}

func beenden2 (resolution uint16) { 	//Inhalt der "Wirklich Beenden"-Seite
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(4*resolution))
	gfx.SchreibeFont(305*resolution/10,450*resolution/10,"Wollen Sie wirklich dieses Spiel beenden?")
}
		

// Inhalt der Regelseiten-------------------------------------------------------------------------------------
//Die Regelseiten werden mit Spiefeldern belegt um bestimmte Zustände auf dem Spielfeld darzustellen

func regeln1 (resolution uint16) {		//Zeigt das Spielfeld vor dem ziehen des aktuellen Plättchens
	var sf spielfeld.Spielfeld
	sf = spielfeld.New(resolution,"Max Mustermann")	//Spielfeld wir generiert und inizialisiert
	sf.Draw(resolution)								//Spielfeld wird dargestellt
	sf.GibSeite1().ShowButton(resolution)			//Buttons der Seite 1 gezeichnet (Buttons sind dadurch inaktiv)
	textFeldDraw(resolution)						//Hintergrund des Textfeldes wird dargestellt
	regeln1Text(resolution)							//Inhalt des Textfeldes wird dargestellt
}

func regeln1Text (resolution uint16) { //Inhalt des Textfeldes zur Regel-Seite 1
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(32*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,157*resolution/10,"Das Spielfeld")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,199*resolution/10,"Es besteht aus einer sechs-")
	gfx.SchreibeFont(1147*resolution/10,224*resolution/10,"eckigen Legefläche, mit 19")
	gfx.SchreibeFont(1147*resolution/10,249*resolution/10,"Feldern, Anzeigen für Zug,")
	gfx.SchreibeFont(1147*resolution/10,274*resolution/10,"Punktestand und Name des")
	gfx.SchreibeFont(1147*resolution/10,299*resolution/10,"Spielers, dessen Spielfeld")
	gfx.SchreibeFont(1147*resolution/10,324*resolution/10,"gerade angezeigt wird.")
	gfx.SchreibeFont(1147*resolution/10,359*resolution/10,"Das aktuelle Plättchen,")
	gfx.SchreibeFont(1147*resolution/10,384*resolution/10,"zeigt das bei diesem Zug")
	gfx.SchreibeFont(1147*resolution/10,409*resolution/10,"zu spielende Plättchen an.")
	gfx.SchreibeFont(1147*resolution/10,434*resolution/10,"Zur Zeit ist es noch leer.")
	gfx.SchreibeFont(1147*resolution/10,469*resolution/10,"Ganz unten sind die 27 noch")
	gfx.SchreibeFont(1147*resolution/10,494*resolution/10,"im Spiel befindlichen ")
	gfx.SchreibeFont(1147*resolution/10,519*resolution/10,"Plättchen zu sehen. Zur Zeit")
	gfx.SchreibeFont(1147*resolution/10,544*resolution/10,"sind das noch alle Plättchen.")
	gfx.SchreibeFont(1530*resolution/10,768*resolution/10,"1/4")
}

var spiel = [19]uint8 {6,15,24,8,26,0,18,19,4,25,17,14,23,22,21,20,16,13,12} //Das auf den Regel-Seiten verwendete Spiel

func regeln2 (resolution uint16) { //Zeigt das Spielfeld nach der Auswahl des ersten aktuellen Plättchens
	var fs  []feld.Feld
	var sf spielfeld.Spielfeld
	sf = spielfeld.SpielfeldGenerator(resolution,fs,spiel) 	//Spielfeld generator bring das Spielfeld bereits in den gewünschten zustand, daher vorher nicht nutzbar
	sf.Draw(resolution)										//Darstellung des Spielfeldes wie bei Regel-Seite 1
	sf.GibSeite1().ShowButton(resolution)					
	textFeldDraw(resolution)
	regeln2Text(resolution)									//Inhalt des Textfeldes für die Regel-Seite 2
}

func regeln2Text (resolution uint16) {	//Inhalt des Textfeldes für die Regel-Seite 2
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(32*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,157*resolution/10,"Der Zug")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,199*resolution/10,"Von den Plättchen im Spiel")
	gfx.SchreibeFont(1147*resolution/10,224*resolution/10,"wird zufällig eines aus-")
	gfx.SchreibeFont(1147*resolution/10,249*resolution/10,"gewählt, welches dann zum ")
	gfx.SchreibeFont(1147*resolution/10,274*resolution/10,"aktuellen Plättchen wird.")
	gfx.SchreibeFont(1147*resolution/10,309*resolution/10,"Der Spieler muss, durch")
	gfx.SchreibeFont(1147*resolution/10,334*resolution/10,"anklicken eines Feldes")
	gfx.SchreibeFont(1147*resolution/10,359*resolution/10,"der Legefläche, auswählen")
	gfx.SchreibeFont(1147*resolution/10,384*resolution/10,"wo er das Plättchen hin-")
	gfx.SchreibeFont(1147*resolution/10,409*resolution/10,"legen möchte.")
	gfx.SchreibeFont(1530*resolution/10,768*resolution/10,"2/4")
}

func regeln3 (resolution uint16) { 	//Zeigt das Spielfeld nach der Wahl des Feldes durch den Spieler an
	var f feld.Feld					//Daher wird ein Feld generiert
	var fs  []feld.Feld				//und ein leerer Slice von Feldern, der die bereits gezogenen Züge enthält (hier noch leer)
	var sf spielfeld.Spielfeld		//Generieren eines Spielfeldes
	f = feld.New()					//Initialisierung des Feldes
	f.SetzeFeld(3,1)				//Belegen des Feldes mit den entsprechenden werten für Spalte und Zeile
	sf = spielfeld.SpielfeldGenerator(resolution,fs,spiel)	//Spielfeld inizialisieren
	sf.FeldSetzen(f)				//Feld als aktuelles Feldsetzen
	sf.Draw(resolution)						//Spielfeld wird dargestellt (nach Feldwahl des Spielers)
	sf.GibSeite2().ShowButton(resolution)	//Spielfeld Seite 2, Buttons nur darstellen (inaktiv)
	sf.PlättchenAnzeigen(resolution)		//aktuelles Plättchen an der Stelle anzeigen die durch das aktuelle Feld gewählt wurde
	textFeldDraw(resolution)				//Hintergrund des Texfeldes darstellen
	regeln3Text(resolution)					//Inhalt des Textfeldes für die Regel-Seite 3 darstellen
}

func regeln3Text (resolution uint16) {		//Inhalt des Textfeldes für die Regel-Seite 3
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(32*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,157*resolution/10,"Der Zug")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,199*resolution/10,"Der Spieler kann das")
	gfx.SchreibeFont(1147*resolution/10,224*resolution/10,"Plättchen solange an andere")
	gfx.SchreibeFont(1147*resolution/10,249*resolution/10,"Positionen auf der Lege-")
	gfx.SchreibeFont(1147*resolution/10,274*resolution/10,"fläche bewegen, bis er den")
	gfx.SchreibeFont(1147*resolution/10,299*resolution/10,"Weiter-Knopf anklickt.")
	gfx.SchreibeFont(1147*resolution/10,334*resolution/10,"Damit beendet der Spieler")
	gfx.SchreibeFont(1147*resolution/10,359*resolution/10,"seinen Zug.")
	gfx.SchreibeFont(1147*resolution/10,394*resolution/10,"Bei mehreren Spielern ist")
	gfx.SchreibeFont(1147*resolution/10,419*resolution/10,"jetzt entweder der nächste")
	gfx.SchreibeFont(1147*resolution/10,444*resolution/10,"Spieler dran, der auch das")
	gfx.SchreibeFont(1147*resolution/10,469*resolution/10,"Plättchen auf seiner Lege-")
	gfx.SchreibeFont(1147*resolution/10,494*resolution/10,"fläche plazieren muss, oder")
	gfx.SchreibeFont(1147*resolution/10,519*resolution/10,"es gibt eine neue Zugrunde")
	gfx.SchreibeFont(1147*resolution/10,544*resolution/10,"mit einem neuen aktuellen")
	gfx.SchreibeFont(1147*resolution/10,569*resolution/10,"Plättchen.")
	gfx.SchreibeFont(1147*resolution/10,604*resolution/10,"Dies wiederholt sich dann")
	gfx.SchreibeFont(1147*resolution/10,629*resolution/10,"so lange, bis die gesamte")
	gfx.SchreibeFont(1147*resolution/10,654*resolution/10,"Legefläche voll ist.")
	gfx.SchreibeFont(1147*resolution/10,689*resolution/10,"Es werden also nur 19 ")
	gfx.SchreibeFont(1147*resolution/10,714*resolution/10,"Plättchen von den ")
	gfx.SchreibeFont(1147*resolution/10,739*resolution/10,"27 möglichen gezogen.")
	gfx.SchreibeFont(1530*resolution/10,768*resolution/10,"3/4")
}

func regeln4 (resolution uint16) { 		//Ansicht des Spielfeldes nach 10 Zügen - also Darstellung des 11. Zugs des Spielers
	var f1,f2,f3,f4,f5,f6,f7,f8,f9,f10 feld.Feld 	//Es werden 10 Felder generiert
	var fs []feld.Feld								//und ein Slice von Feldern, die dann dem SpielfeldGenerator übergeben wird, der dann die 10 angegeben Züge macht
	var sf spielfeld.Spielfeld						//Generierung eines Spielfeldes
	f1 = feld.New()									//Die 10 Felder werden Initialisiert und mit Werten belegt und
	f1.SetzeFeld(3,1)
	f2 = feld.New()
	f2.SetzeFeld(3,2)
	f3 = feld.New()
	f3.SetzeFeld(3,3)
	f4 = feld.New()
	f4.SetzeFeld(3,4)
	f5 = feld.New()
	f5.SetzeFeld(3,5)
	f6 = feld.New()
	f6.SetzeFeld(5,1)
	f7 = feld.New()
	f7.SetzeFeld(5,2)
	f8 = feld.New()
	f8.SetzeFeld(5,3)
	f9 = feld.New()
	f9.SetzeFeld(2,4)
	f10 = feld.New()
	f10.SetzeFeld(4,3)
	fs = append(fs,f1,f2,f3,f4,f5,f6,f7,f8,f9,f10)	//An den Slice angehangen
	sf = spielfeld.SpielfeldGenerator(resolution,fs,spiel)	//Das Spielfeld wird in den Zustand gebracht nach 10 Zügen
	sf.Draw(resolution)										//Darstellen des Spielfeldes					
	sf.GibSeite1().ShowButton(resolution)					//Inaktive Buttons darstellen
	sf.DrawLegefläche(resolution)							//Bereits gelegten Plättchen anzeigen
	textFeldDraw(resolution)								//Textfeld Hintergrund
	regeln4Text(resolution)									// und Inhalt darstellen
}

func regeln4Text (resolution uint16) {	//Inhalt des Textfeldes für die Regel-Seite 4
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(32*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,157*resolution/10,"Ziel des Spiels")
	gfx.SetzeFont("../fonts/LiberationMono-Bold.ttf",int(25*resolution/10))
	gfx.SchreibeFont(1147*resolution/10,199*resolution/10,"Ziel des Spiels ist es")
	gfx.SchreibeFont(1147*resolution/10,224*resolution/10,"möglichst viele Punkte zu ")
	gfx.SchreibeFont(1147*resolution/10,249*resolution/10,"bekommen.")
	gfx.SchreibeFont(1147*resolution/10,284*resolution/10,"Punkte gibt es aber nur,")
	gfx.SchreibeFont(1147*resolution/10,309*resolution/10,"für vollständige, ununter-")
	gfx.SchreibeFont(1147*resolution/10,334*resolution/10,"brochene Reihen in der")
	gfx.SchreibeFont(1147*resolution/10,359*resolution/10,"selben Zahl und Farbe.")
	gfx.SchreibeFont(1147*resolution/10,394*resolution/10,"Dabei berechnen sich die")
	gfx.SchreibeFont(1147*resolution/10,419*resolution/10,"Punkte nach der Anzahl der")
	gfx.SchreibeFont(1147*resolution/10,444*resolution/10,"Plättchen mal deren Zahl.")
	gfx.SchreibeFont(1147*resolution/10,479*resolution/10,"Daher ergeben sich für die")
	gfx.SchreibeFont(1147*resolution/10,504*resolution/10,"1er-Reihe 3 Punkte und für")
	gfx.SchreibeFont(1147*resolution/10,529*resolution/10,"die 9er-Reihe 45 Punkte.")
	gfx.SchreibeFont(1147*resolution/10,554*resolution/10,"Die Diagonale bringt keine")
	gfx.SchreibeFont(1147*resolution/10,579*resolution/10,"Punkte, da hier verschiedene")
	gfx.SchreibeFont(1147*resolution/10,604*resolution/10,"Zahlen in einer Reihe liegen.")
	gfx.SchreibeFont(1147*resolution/10,624*resolution/10,"")
	gfx.SchreibeFont(1147*resolution/10,649*resolution/10,"")
	gfx.SchreibeFont(1147*resolution/10,674*resolution/10,"Viel Spaß beim spielen!")
	gfx.SchreibeFont(1530*resolution/10,768*resolution/10,"4/4")
}


func textFeldDraw (resolution uint16) {	//Textfeld Hintergrund für die Regel-Seiten
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(114*resolution,15*resolution,44*resolution,65*resolution)
	gfx.Stiftfarbe(255,255,255)
	gfx.Vollrechteck(1142*resolution/10,152*resolution/10,436*resolution/10,646*resolution/10)
}

// Inhalt der Regelseiten Ende -------------------------------------------------------------------------









