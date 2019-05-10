package Entidad

import (
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
)

type File struct {
	Name string
	Versiones []Version
	Type string

}

func (file File) Whois() string {
	a :=strings.SplitAfter(file.Name, ".")
	return a[len(a) - 1]

}
type Version struct {
	Content []byte
	Version time.Time
	Name uuid.UUID
	Status bool

}
type Folder struct {
	Name string
	Files []File
	Path string
}
type WorkPlace struct {
	ListnerFolder []Folder
}

// Agregar una versión de un archivo (que podría o no existir previamente)
// por que en folder?, por el path despues se le el archivo en esa ubicacion se guarda el path pero
// el nombre del  file el retutn bool sirve como control de errores! y el root es la carpeta padre
// desde ese punto nos podemos mover a otro repo y el contenido de ese arechivo que abrimos en el path (byte)
// como version
/*
func (f *Folder)AddFile(path string, created time.Time)(resul bool,err error){

}
func (f *Folder)DeleteFile(path string, created time.Time)(resul bool,err error){

}
func (f *Folder)Exists(path string, created time.Time) bool{

}
*/
func (f *Folder) GetFolderName(s string) string {
	P := strings.Split(s, string(os.PathSeparator))
	return P[len(P) - 1]

}


