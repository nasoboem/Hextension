package felder

/*************************************************************************
  Paket zur Ein- und Ausgabe von Text unter Verwendung von (Formular-)Fel-
  dern. Felder und ihr Inhalt können dabei  wiederverwendet werden. Es ist
  sowohl möglich, in Ihnen Eingaben zu tätigen, die Eingaben zu einem spä-
  teren Zeitpunkt zu ändern oder sie zur Ausgabe zu  verwenden. Jedes Feld
  kann den Inhalt rechts- oder  linksbündig oder zentriert ausgeben. Schon
  bei der Eingabe kann man festlegen, welche  Zeichen in  diesem Feld ver-
  wendet werden dürfen. Dadurch können Fehleingaben bereits  in einer sehr
  frühen Phase vermieden werden.
  Die Implementierung  benutzt das gfx-Paket von  Stefan Schmidt  in einer
  Version, die mindestens ein Datum vom 12.03.2017 trägt.

  (c) Oliver Schäfer           Versionsdatum: So 11. Feb 23:42:22 CET 2018
*************************************************************************/

//  Konstruktor des ADT Feld
/*
    New (x,y,anz, ausrichtung byte, titel string) Feld
*/

//  Erzeugt ein (rechteckiges) und  grau (r,g,b)=(0xDC,0xDC,0xDC) hinter-
//  legtes Ein-/Ausgabefeld und bringt es auf dem Bildschirm zur Anzeige.
//  Die linke, obere Ecke liegt  in (x|y), das  Feld  erlaubt maximal anz
//  Zeichen.
//  Der Text ist gemäß ausrichtung  linksbündig ('l'), rechtsbündig ('r')
//  oder zentriert ('z') ausgerichtet. Als Zeichensatz wird ein Terminus-
//  Monospace-Font (der Größe 16) verwendet. Unterhalb des Feldes ist ei-
//  ne 12-Punkt große Beschriftung aus  titel  gesetzt. Wird das Feld zur
//  Eingabe verwendet, sind neben den Steuerzeichen ENTER  und  BACKSPACE
//  alle über die Tastatur erreichbaren (und auch druckbaren) Zeichen er-
//  laubt (Ascii).

//  Allgemeine Einstellungen, gültig für alle Felder, können aber indivi-
//  duell geändert werden. Als Zeichengrößen stehen die Werte 12, 14, 16,
//  20, 22, 24, 28 und 32 zur Verfügung.
/*
    Voreinstellungen (r,g,b uint8, groesse uint16)
*/

//  Setzt - abweichend  von den  Standardeinstellungen - die Hintergrund-
//  farbe als (rot,grün,blau)-Tripel und  die Größe der im Feld verwende-
//  ten Zeichen.

const (
  Digits    = "0123456789"
  Capitals  = "ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖÜ"
  Small     = "abcdefghijklmnopqrstuvwxyzäöüß"
  Specials  = "\"\\!§$%&/()=?´`+*~#'-_.:,;@^°[]{}|<> "
  Letters   = Capitals + Small
  Ascii     = Capitals + Small + Digits + Specials
)

type Feld interface {

	
  //  Vor.: Keine
  //  Eff.: Das Eingabefeld akzeptiert neben ENTER und BACKSPACE  nur die
  //        in erlaubt enthaltenen Zeichen, andere Eingaben  werden igno-
  //        riert
  SetzeErlaubteZeichen (erlaubt string)

  //  Vor.: Keine.
  //  Eff.: Für  die Zeichen innerhalb des  Feldes wird der Standard-Zei-
  //        chensatz Terminus mit der in groesse angegebenen Größe sowohl
  //        für Ein- als auch für Ausgaben verwendet.
  SetzeZeichensatzgroesse (groesse uint16)

  //  Vor.: Keine.
  //  Eff.: Die Hintergrundfarbe ist auf den angegebenen Wert (r,g,b) ge-
  //        setzt.
  SetzeHintergrundfarbe (r,g,b uint8)

  //  Vor.: Ein gfx-Fenster ist geöffnet, kein  anderer Thread hat gerade
  //        den Tastaturfokus.
  //  Eff.: Das Eingabefeld ist auf dem  Bildschirm dargestellt, der Tas-
  //        taturfokus liegt  in diesem Eingabefeld. Mit  dem Beenden der
  //        Eingabe über ENTER oder ESCAPE wird die gerade im Eingabefeld
  //        dargestellte Zeichenkette geliefert.
  Edit () string

  //  Vor.: Ein gfx-Fenster  ist geöffnet.
  //  Eff.: Der Text ist gemäß der  eingestellten Ausrichtung  des Feldes
  //        im Feld ausgegeben. Ist  der Text länger als  das Feld breit,
  //        wird die Ausgabe-Zeichenkette abgeschnitten. Die Zeichenkette
  //        text ist (maximal soviele Zeichen, wie das Feld breit ist) im
  //        Feld gespeichert und kann z. B. beim Editieren  wieder verän-
  //        dert werden.
  Schreibe (text string)

  //  Vor.: Ein gfx-Fenster ist geöffnet.
  //  Eff.: Das Feld ist geleert, sonst ist nichts geändert.
  Leere ()

  //  Vor.: Keine
  //  Eff.: Genau dann, wenn der Inhalt des Feldes leer ist, ist true ge-
  //        liefert.
  IstLeer () bool
  
  SetzeTitel (nt string)
}
