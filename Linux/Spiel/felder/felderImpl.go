package felder

import "../gfx"

var (
  // Globale Fontvariablen
  FONTPATH_DESC = "../fonts/LiberationSans-Regular.ttf"
  FONTPATH_FELD = "../fonts/LiberationMono-Bold.ttf"
  FONTHEIGHT    =  uint16(40)
  FONTWIDTH     =  uint16(25)
  // Globale Hintergrund-Farbvariablen
  R = uint8(0xDC)
  G = uint8(0xDC)
  B = uint8(0xDC)
)

const (
  ENTER     =  13
  BACKSPACE =   8
  TAB       =   9
  ESCAPE    =  27
)



type feld struct {
  x,y,anz,pos uint16
  fontpfad    string
  fontht      uint16
  fontwd      uint16
  titel       string
  text        string
  valid       string
  align       byte
  r,g,b       uint8
}



func Voreinstellungen (r,g,b uint8, groesse uint16) {
  R,G,B = r,g,b
  FONTWIDTH,FONTHEIGHT = regularSize (groesse)
}


func New (x,y,anz int, ausrichtung byte, titel string) *feld {
  var f *feld = new (feld)
  (*f).x,(*f).y = uint16(x),uint16(y)
  (*f).anz      = uint16(anz)
  (*f).titel    = titel
  (*f).pos      = 0
  (*f).fontpfad = FONTPATH_FELD
  (*f).fontht   = FONTHEIGHT
  (*f).fontwd   = FONTWIDTH
  (*f).text     = ""
  (*f).valid    = Ascii
  (*f).align    = ausrichtung
  (*f).r,(*f).g,(*f).b = R,G,B
  //refresh (f)
  return f
}

func (f *feld) SetzeTitel (nt string) {
	(*f).titel = nt
	refresh (f)
}

/*
func (f *feld) SetzeZeichensatz (fontpfad string, groesse int) {
  (*f).fontpfad = fontpfad
  (*f).fontht   = uint16(groesse)
  (*f).fontwd   = uint16(groesse) / 2 // TODO
}
*/

func (f *feld) SetzeErlaubteZeichen (erlaubt string) {
  (*f).valid = erlaubt
}

func (f *feld) Schreibe (text string) {
  var utf8 []rune

  utf8 = []rune (text)
  utf8 = utf8 [0:min2 (len(utf8),int((*f).anz))]
  (*f).text = string (utf8)
  (*f).pos = uint16 (len (utf8))
  refresh (f)
}

func (f *feld) Edit () string {
  var (
    gedrueckt     uint8
    zeichen,tiefe uint16
    utf8          []rune
  )

  // Eingabefeld hat den Focus, daher rot hinterlegt
//  (*f).r,(*f).g,(*f).b = 0xFF,0xCF,0xCF
  (*f).r,(*f).g,(*f).b = 0x9F,0x9F,0x9F
  refresh (f)
  gfx.TastaturpufferAn ()

  Eingabe:
  for ;; {
    zeichen,gedrueckt,tiefe = gfx.TastaturpufferLesen1 ()
    if (gedrueckt == 1) && !istSondertaste (zeichen) && istEnthalten (zeichen, tiefe, (*f).valid) {
      switch zeichen {
        case ENTER:
          // Noch Loslassen der Taste entgegennehmen und dann tschüss
          gfx.TastaturpufferLesen1 ()
          break Eingabe
        case ESCAPE:
          // Noch Loslassen der Taste entgegennehmen und dann tschüss
          gfx.TastaturpufferLesen1 ()
          break Eingabe
        case TAB:
          // Noch Loslassen der Taste entgegennehmen und dann tschüss
          gfx.TastaturpufferLesen1 ()
          break Eingabe
        case BACKSPACE:
          // Letztes Zeichen löschen, wenn nicht bereits am Anfang
          // Sonderbehandlung für Nicht-1-Byte-Zeichen (äöüß)
          if (*f).pos > 0 {
            (*f).pos-=1
            utf8 = []rune((*f).text)
            (*f).text = string (utf8 [0:len(utf8)-1])
          }
        default:
          if (*f).pos < uint16((*f).anz) {
            (*f).text = (*f).text + taste2char (zeichen, tiefe)
            (*f).pos+=1
          }
      }
    refresh (f)
    }
  }

  // Eingabefeld wieder als inaktiv markieren (=grau hinterlegt)
  (*f).r,(*f).g,(*f).b = R,G,B
  refresh (f)
  gfx.TastaturpufferAus ()

  return (*f).text
}

func (f *feld) Leere () {
  (*f).text = ""
  refresh (f)
}

func (f *feld) IstLeer () bool {
  return (*f).text == ""
}

func (f *feld) SetzeZeichensatzgroesse (groesse uint16) {
  (*f).fontwd,(*f).fontht = regularSize (groesse)
  refresh (f)
}

func (f *feld) SetzeHintergrundfarbe (r,g,b uint8) {
  (*f).r,(*f).g,(*f).b = r,g,b
  refresh (f)
}

// ------------------------------- nicht exportierte Funktionen ----------------------------------

func regularSize (size uint16) (w,h uint16) {
  switch size {
    case 12: w,h =  7,12
    case 14: w,h =  9,14
    case 16: w,h = 10,16
    case 20: w,h = 12,20
    case 22: w,h = 14,22
    case 24: w,h = 15,24
    case 28: w,h = 17,28
    case 32: w,h = 20,32
    case 36: w,h = 22,36
    case 40: w,h = 25,40
    default: w,h = 10,16
  }
  return
}

func istSondertaste (c uint16) bool {
  // Tastencodes aller Sondertasten außer BACKSPACE, ENTER, POS1, ENDE, EINF, Pfeil rechts, Pfeil links
  var liste =[]uint16{275,276,280,281,282,283,284,285,286,287,288,289,290,291,292,292,301,302,303,304,306,305,311,313,319}
  for _,x := range (liste) {
    if x == c { return true }
  }
  return false
}

func istEnthalten (c,t uint16, liste string) bool {
  if (c == ENTER) || (c == BACKSPACE) || (c == ESCAPE) || (c == TAB) { return true }
  for _,x := range (liste) {
//    if (c == uint16(x)) { return true }
    if (taste2char (c,t) == string(x)) { return true }
  }
  return false
}

func refresh (f *feld) {
  var (
    utf8    []rune
    offset  uint16
  )

  // Eingabefeld in der aktuellen Farbe anzeigen
  gfx.Stiftfarbe ((*f).r,(*f).g,(*f).b)
  gfx.Vollrechteck (uint16((*f).x-2),uint16((*f).y-2),uint16((*f).anz)*(*f).fontwd+4,(*f).fontht+4)
  // Texte werden immer in schwarz dargestellt
  gfx.Stiftfarbe (0x0, 0x0, 0x0)
  // Titel in 12-Punkt setzen
  gfx.SetzeFont (FONTPATH_DESC, 12)
  // Den Titel aber nur ausgeben, wenn er nicht leer ist.
  if len((*f).titel) > 0 {
    gfx.SchreibeFont (uint16((*f).x), uint16((*f).y)+(*f).fontht+2, (*f).titel)
  }
  // Ggf. gespeicherten Text schreiben
  gfx.SetzeFont ((*f).fontpfad, int((*f).fontht))
  utf8 = []rune ((*f).text)
  switch (*f).align {
    case 'r':
      offset = ((*f).anz-uint16 (len (utf8)))*((*f).fontwd)
    case 'z':
      offset = (((*f).anz-uint16 (len (utf8)))*((*f).fontwd)) / 2
    // default-Fall umfasst die Standardausrichtung 'l', die auch verwendet
    // wird, wenn der Benutzer eine ungültige Ausrichtung angegeben hat.
    default:
      offset = 0
  }
  gfx.SchreibeFont (uint16((*f).x)+offset, uint16((*f).y), (*f).text)
}

func taste2char (taste,tiefe uint16) string {
  var character string
  // In erster Näherung gilt dies:
  character = string (taste)
  // Falls SHIFT oder AltGr gedrückt ist, Sonderbehandlung
  // SHIFT-Taste gedrückt?
  if (tiefe%2 == 1) || ((tiefe>>1)%2 == 1) {
    switch {
      case (taste >= 'a') && (taste <= 'z'):
        character = string (taste-32)
      case taste == '^': character = "°"
      case taste == '1': character = "!"
      case taste == '2': character = "\""
      case taste == '3': character = "§"
      case taste == '4': character = "$"
      case taste == '5': character = "%"
      case taste == '6': character = "&"
      case taste == '7': character = "/"
      case taste == '8': character = "("
      case taste == '9': character = ")"
      case taste == '0': character = "="
      case taste == 'ß': character = "?"
      case taste == '´': character = "`"
      case taste == '+': character = "*"
      case taste == '#': character = "'"
      case taste == '-': character = "_"
      case taste == '.': character = ":"
      case taste == ',': character = ";"
      case taste == 'ä': character = "Ä"
      case taste == 'ö': character = "Ö"
      case taste == 'ü': character = "Ü"
      case taste == '<': character = ">"
    }
  }
  // AltGr-taste gedrückt?
  if (tiefe>>14)%2 == 1 {
    switch taste {
      case '7': character = "{"
      case '8': character = "["
      case '9': character = "]"
      case '0': character = "}"
      case 'ß': character = "\\"
      case '+': character = "~"
      case 'q': character = "@"
      case '<': character = "|"
      case 'e': character = "€"
    }
  }
  return character
}

func min2 (a,b int) int {
  if a < b { return a } else { return b }
}
