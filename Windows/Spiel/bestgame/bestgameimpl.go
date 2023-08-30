package bestgame

//Es gibt 19! möglichkeiten, was etwa 1,2 * 10¹⁷ ist
//Brut-Force -Algorithmus funktioniert, wenn man mehr 142.694 Jahre zur Verfügung hat
//Es wäre ein Algorithmus zu finden, der ca. 10¹⁰x schneller wäre und auch dann bräuchte man noch 7,5 min für die Berechnung 

import ("../plättchen"
		"../feld"
		"../spielfeld"
		"fmt")

func OptGame (game [19]uint8) int { //Funktion, die nach dem 
	var max int
	var f []feld.Feld																				//Slice, dem die spielbaren Felderübergeben werden <-- FeldGenerator()
	var zähler uint64																				//Zählt die Anzahl der Variationen enes voll belegten Spielfeldes
	var zehner uint64 = 1																			//Spechert die 10er-Potenzen
	f = feld.FeldGenerator()																		//Übergabe aller Felder
	var pis [27]plättchen.Plättchen  																//Plättchen im Spiel <-- PlättchenGenerator()
	var np,p1,p2,p3,p4,p5,p6,p7,p8,p9,p10,p11,p12,p13,p14,p15,p16,p17,p18,p19 plättchen.Plättchen	//Initzialisierung der 20 Plättchenvariablen ein Nullerplättchen (np) zum Vergleich und 19 Nullerplättchen um das Plättchen zu sperren
	np = plättchen.New(0,0,0)																		//Generierung des Nullerplättchen
	pis = plättchen.PlättchenGenerator()															//Übergabe aller möglichen Plättchen
	var p [19]plättchen.Plättchen																	//Inizialisierung eines Feldes von 19 Plättchen 
	for i:=0;i<len(game);i++{																		//Auswahl der 19 im Spiel verwendeten Plättchen aus den 27 möglichen Plättchen
		p[i] = pis[game[i]]																			//Befüllung des Slices mit den Plättchen, die tasächlich im Spiel vorkommen
	}
	var sf spielfeld.Spielfeld																		//Neues Spielfeld wird generiert
	sf = spielfeld.New(10,"")																		//und initialisiert
	sf.PlättchenAusdemSpielnehmen(game)																//Anzeige der Plättchem im Spiel wird angepasst
	for a:=0;a<19;a++{																				//Für jedes Feld wird ein for-Loop erstellt
		sf.FreiesLegen(f[0],p[a])																	//wo ein Plättchen auf das Spielfeld gelegt wird
		p1 = plättchen.New(0,0,0)																	//Sperrnullerplättchen wird generiert
		p[a] = p1																					//und an dem Index eingesetzt im Slice 
		for b:=0;b<19;b++{
			if plättchen.IstGleich(p[b],np) {														//hier wird getestet, ob es sich um ien Nullerplättchen handelt, dann wird es nicht eingesetzt
				continue
			} else { 
				sf.FreiesLegen(f[1],p[b])															//Nächstes richtige Plättchen wird gesetzt
				p2 = plättchen.New(0,0,0)															//und durch ein Nullerplättchen ersetzt usw.
				p[b] = p2
			}
			for c:=0;c<19;c++{
				if plättchen.IstGleich(p[c],np) {
					continue
				} else { 
					sf.FreiesLegen(f[2],p[c])
					p3 = plättchen.New(0,0,0)
					p[c] = p3
				}
				for d:=0;d<19;d++{
					if plättchen.IstGleich(p[d],np) {
						continue
					} else { 
						sf.FreiesLegen(f[3],p[d])
						p4 = plättchen.New(0,0,0)
						p[d] = p4
					}
					for e:=0;e<19;e++{
						if plättchen.IstGleich(p[e],np) {
						continue
						} else { 
							sf.FreiesLegen(f[4],p[e])
							p5 = plättchen.New(0,0,0)
							p[e] = p5
						}
						for fs:=0;fs<19;fs++{
							if plättchen.IstGleich(p[fs],np) {
								continue
							} else { 
								sf.FreiesLegen(f[5],p[fs])
								p6 = plättchen.New(0,0,0)
								p[fs] = p6
							}
							for g:=0;g<19;g++{
								if plättchen.IstGleich(p[g],np) {
									continue
								} else { 
									sf.FreiesLegen(f[6],p[g])
									p7 = plättchen.New(0,0,0)
									p[g] = p7
								}
								for h:=0;h<19;h++{
									if plättchen.IstGleich(p[h],np) {
										continue
									} else { 
										sf.FreiesLegen(f[7],p[h])
										p8 = plättchen.New(0,0,0)
										p[h] = p8
									}
									for i:=0;i<19;i++{
										if plättchen.IstGleich(p[i],np) {
											continue
										} else { 
											sf.FreiesLegen(f[8],p[i])
											p9 = plättchen.New(0,0,0)
											p[i] = p9
										}
										for j:=0;j<19;j++{
											if plättchen.IstGleich(p[j],np) {
												continue
											} else {
												sf.FreiesLegen(f[9],p[j])
												p10 = plättchen.New(0,0,0)
												p[j] = p10
											}
											for k:=0;k<19;k++{
												if plättchen.IstGleich(p[k],np) {
													continue
												} else { 
													sf.FreiesLegen(f[10],p[k])
													p11 = plättchen.New(0,0,0)
													p[k] = p11
												}
												for l:=0;l<19;l++{
													if plättchen.IstGleich(p[l],np) {
														continue
													} else { 
														sf.FreiesLegen(f[11],p[l])
														p12 = plättchen.New(0,0,0)
														p[l] = p12
													}
													for m:=0;m<19;m++{
														if plättchen.IstGleich(p[m],np) {
															continue
														} else { 
															sf.FreiesLegen(f[12],p[m])
															p13 = plättchen.New(0,0,0)
															p[m] = p13
														}
														for n:=0;n<19;n++{
															if plättchen.IstGleich(p[n],np) {
																continue
															} else { 
																sf.FreiesLegen(f[13],p[n])
																p14 = plättchen.New(0,0,0)
																p[n] = p14
															}
															for o:=0;o<19;o++{
																if plättchen.IstGleich(p[o],np) {
																	continue
																} else { 
																	sf.FreiesLegen(f[14],p[o])
																	p15 = plättchen.New(0,0,0)
																	p[o] = p15
																}
																for pf:=0;pf<19;pf++{
																	if plättchen.IstGleich(p[pf],np) {
																		continue
																	} else { 
																		sf.FreiesLegen(f[15],p[pf])
																		p16 = plättchen.New(0,0,0)
																		p[pf] = p16
																	}
																	for q:=0;q<19;q++{
																		if plättchen.IstGleich(p[q],np) {
																			continue
																		} else {
																			sf.FreiesLegen(f[16],p[q])
																			p17 = plättchen.New(0,0,0)
																			p[q] = p17
																		}
																		for r:=0;r<19;r++{
																			if plättchen.IstGleich(p[r],np) {
																				continue
																			} else { 
																				sf.FreiesLegen(f[17],p[r])
																				p18 = plättchen.New(0,0,0)
																				p[r] = p18
																			}
																			for s:=0;s<19;s++{
																				if plättchen.IstGleich(p[s],np) {
																					continue
																				} else {
																					sf.FreiesLegen(f[18],p[s])    //Das ganze 19 mal, jetzt ist die Legefläche voll
																					p19 = plättchen.New(0,0,0)
																					p[s] = p19
																					zähler++					//Zähler zählt um eins hoch
																					if zähler==zehner {
																						fmt.Println("Der Zähler hat ",zehner," erreicht!")
																						zehner = zehner*10
																					}
																					//fmt.Println(zähler)
																				}
																				if max < sf.GibPunkte() {		//Gibt es eine neue höchst Punktzahl
																					max = sf.GibPunkte()		//So wir diese gespeichert
																					fmt.Println(sf)				//Das Spielfeld ausgegeben und 
																					fmt.Println("Dieses Variation wurde gefunden nach ",zähler, " Versuchen!") // dieser Satz angezeigt
																				}
																				p19 = sf.LetztesPlättchenEntfernen() //Jetzt werden die Plättchen wieder zurück gelegt und die for-Loops beginnen von forn
																				p[s] = p19															
																			}
																			p18 = sf.LetztesPlättchenEntfernen()
																			p[r] = p18
																		}
																		p17 = sf.LetztesPlättchenEntfernen()
																		p[q] = p17
																	}
																	p16 = sf.LetztesPlättchenEntfernen()
																	p[pf] = p16
																}
																p15 = sf.LetztesPlättchenEntfernen()
																p[o] = p15
															}
															p14 = sf.LetztesPlättchenEntfernen()
															p[n] = p14
														}
														p13 = sf.LetztesPlättchenEntfernen()
														p[m] = p13
													}
													p12 = sf.LetztesPlättchenEntfernen()
													p[l] = p12
												}
												p11 = sf.LetztesPlättchenEntfernen()
												p[k] = p11
											}
											p10 = sf.LetztesPlättchenEntfernen()
											p[j] = p10
										}
										p9 = sf.LetztesPlättchenEntfernen()
										p[i] = p9
									}
									p8 = sf.LetztesPlättchenEntfernen()
									p[h] = p8
								}
								p7 = sf.LetztesPlättchenEntfernen()
								p[g] = p7
							}
							p6 = sf.LetztesPlättchenEntfernen()
							p[fs] = p6
						}
						p5 = sf.LetztesPlättchenEntfernen()
						p[e] = p5
					}
					p4 = sf.LetztesPlättchenEntfernen()
					p[d] = p4
				}
				p3 = sf.LetztesPlättchenEntfernen()
				p[c] = p3
			}
			p2 = sf.LetztesPlättchenEntfernen()
			p[b] = p2
		}
		p1 = sf.LetztesPlättchenEntfernen()
		p[a] = p1
	}
	return max											//Wenn er je erreicht werden würde, (~142.694 Jahre später) würde der max Wert zurück gegeben
}

