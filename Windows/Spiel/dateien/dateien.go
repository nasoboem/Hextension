package dateien

//Modell:

/* Eine Datei kann als Strom bekannter Länge von Bytes betrachtet werden, 
 * wobei jedes Byte im Strom über seinen Index (erstes Byte: Index 0)
 * adressiert werden kann.
 * (--> Man spricht in solch einem Fall von direktem Zugriff.)
 * Zu jeder Datei gehört eine Zeichenkette, der sogenannte Dateiname, über 
 * den man den Strom (später) wieder finden kann.
 * Ein Strom kann durch Anhängen von Bytes an sein Ende oder durch 
 * Überschreiben von Bytes an der zuvor auf "aktuell" gesetzten Stelle
 * modifiziert werden - das nennt man Schreiben.  
 * Zu jedem Zeitpunkt ist nur ein einziges Byte (aktuelles Byte) sichtbar,
 * es sei denn, man ist am Ende der Datei.
 * Der Zugriff auf das aktuelle Byte heißt Lesen, dadurch wird das folgende 
 * Byte sichtbar. Jeder Strom besitzt einen Modus: Er ist entweder
 * zum Lesen geöffnet oder zum Schreiben geöffnet oder zum Lesen UND Schreiben. */
 
type Datei interface {
	
	// Vor.: dateiname ist ein gültiger Dateiname. modus ist 'l', 's', 
	//       'a' oder 'x'.
	// Erg.: Eine Instanz vom Typ Datei ist initialisiert und geliefert.
	//       War modus 'l', so ist die zugehörige Datei zum Lesen geöffnet.
	//       Das allererste Byte des Stroms ist aktuell. 
	//       War modus 's', so ist die zugehörige Datei leer und zum
	//       Schreiben geöffnet.
	//       War modus 'a', so ist die zugehörige Datei zum Schreiben geöffnet,
	//       jedoch ist der alte Inhalt erhalten geblieben und es ist 
	//       gibt kein aktuelles Byte.
	//       In den Fällen 'l' und 'a' muss die Datei mit dem angegebenen
	//       Dateinamen muss schon existieren. Konnte die Datei nicht 
	//       geöffnet werden, ist das Programm abgebrochen.
	//		 War modus 'x', so ist die zugehörige Datei zum Lesen und Schreiben
	//       geöffnet. Gab es die Datei schon, so ist der alte Inhalt 
	//       erhalten geblieben und es gibt kein aktuelles Byte.
	//       Andernfalls ist eine neue Datei geöffnet worden.
	//       
	// Oeffnen (dateiname string, modus byte) Datei 
	// ^^^^^^^
	// entspricht dem sonst zu verwendenden New
	
	//Vor.: Die Datei ist geöffnet.
	//Erg.: die aktuelle Länge des Bytestroms der Datei
	Groesse () uint64
	
	//Vor.: Die Datei ist geöffnet.
	//Eff.: Das Byte des Bytestroms der Datei mit dem Index index ist 
	//      aktuell, wenn es dieses Byte gibt, andernfalls ist kein
	//      Byte des Datenstroms aktuell.
	Positionieren (index uint64)
	
	//Vor.: Die Datei ist zum Schreiben oder zum Lesen und Schreiben geöffnet.
	//Erg.: Das Byte an der aktuellen Indexposition ist durch b ersetzt
	//      und die aktuelle Indexposition ist nun eins größer, d.h., 
	//      das folgende Byte im Strom ist nun aktuell. Gab es keine 
	//      aktuelle Indexposition, so ist b an das Ende des Datenstroms
	//      angefügt. Wenn das Schreiben nicht möglich war, ist das
	//      Programm abgebrochen.
	Schreiben (b byte)
	
	// Vor.: Die Datei ist zum Lesen oder zum Lesen und Schreiben geöffnet.
	// Erg.: True ist geliefert, gdw. es kein aktuelles Byte gibt, d.h. 
	//       das Ende des Stroms ist erreicht. Misslingt der Test, so ist das
	//       Programm abgebrochen.
	Ende () bool
	
	// Vor.: Die Datei ist zum Lesen oder zum Lesen und Schreiben geöffnet.
	//       Es gibt ein aktuelles Byte.
	// Erg.: Das aktuelle Byte ist geliefert.
	// Eff.: Das darauf folgende Byte ist nun aktuell, wenn es eins gibt.
	//       Andernfalls ist das Ende des Stroms erreicht. Bei einem Fehler
	//       ist das Programm abgebrochen.
	Lesen () byte
	
	// Vor.: Die Datei wurde noch nicht geschlossen.
	// Eff.: Die Datei wurde geschlossen und steht zum Lesen und Schreiben
	//       nicht mehr zur Verfügung.Gab es beim Schliessen der Datei ein
	//       Problem, so ist das Programm abgebrochen!
	Schliessen ()
}
