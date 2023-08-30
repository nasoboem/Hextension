package dateien

import "os"
import "io"
import "fmt"

type datei struct {
	f *os.File
}
 	
func Oeffnen (dateiname string, modus byte) *datei {
	var d = new(datei)
	var err error
	
	switch modus {
		case 'l':
		((*d).f), err= os.Open(dateiname) 
		case 's':
		((*d).f), err= os.Create(dateiname)
		case 'a':
		((*d).f), err= os.OpenFile(dateiname,os.O_WRONLY | os.O_APPEND,0666)
		case 'x':
		((*d).f), err= os.OpenFile(dateiname,os.O_RDWR | os.O_CREATE,0666)
		default:
		panic("Falscher Modus-Parameter! Programmabbruch!")
	} 
	if err != nil {
		panic ("Fehler beim Öffnen der Datei! Programmabbruch!")
	}
	return d
}

func (d *datei) Positionieren (index uint64) {
	fileinfo,err :=os.Stat((*d).f.Name())
	if err != nil {panic ("Fehler beim Positionieren!")}
	if index > uint64(fileinfo.Size()) {
		index = uint64(fileinfo.Size())
	}
	(*d).f.Seek (int64(index), 0)
}

func (d *datei) Groesse () uint64 {
	fileinfo,err :=os.Stat((*d).f.Name())
	if err != nil {panic ("Fehler beim Ermitteln der Groesse!")}
	return uint64(fileinfo.Size())
}
	

func (d *datei)	Schreiben (b byte) {
	var b1 = make([]byte,1)
	var n int
	b1[0] = b
	n, _ = (*d).f.Write (b1)
	if n != 1 {
		panic("Konnte nicht in die Datei schreiben! Programmabbruch!")
	}
}
	
func (d *datei)	Ende () bool {
	var b1 = make([]byte,1)
	var n int
	var err error
	n, err =(*d).f.Read (b1)
	if n == 1 {
		(*d).f.Seek(-1,1)
		return false
	}
	if n == 0 && err == io.EOF {
		return true
	}
	panic("Fehler beim Test auf Ende! Programmabbruch!")
}
	
func (d *datei) Lesen () byte {
	var b1 = make([]byte,1)
	var n int
	n, _ = (*d).f.Read (b1)
	if n == 1 {
		return b1[0]
	} else {
		panic("Fehler beim Lesen! Programmabbruch!")
	}
}
	
func (d *datei) Schliessen () {
	var err error
	err = (*d).f.Close ()
	if err != nil {
		fmt.Println (err)
		panic ("Fehler beim Schließen der Datei! Programmabbruch!")
	}
}

