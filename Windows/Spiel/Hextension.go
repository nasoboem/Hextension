package main

import (
		"./spiel"
		)
		
func main() {
	spiel.Spielen(9) //Übergabe der Auflösung in Schritten: 
					//1 = 160x100 - What is this? A game for ants? 
					//2 = 320x200 - ...
					//3 = 480x300 - zu klein
					//4 = 640x400 - Schrift nicht vollständig lesbar
					//5 = 800x500 - ab hier spielbar
					//6 = 960x600
					//7 = 1120x700
					//8 = 1280x800
					//9 = 1440x900
					//10 = 1600x1000
					//11 = 1760x1100
					//12 = 1920x1200
					//ab 13 wird das gfx-Fenster nicht mehr größer. Daher werden Inhalte abgeschnitten.
}
