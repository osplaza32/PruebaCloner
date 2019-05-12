package Entidad

import (
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"osplaza32/Prueba_Cloner/Utils"
	"path/filepath"
	"strings"
	"time"
)
type File struct {
	Name string
	Versiones []Version
	Type string
	Id uuid.UUID
	LastUpdate time.Time
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
func (f *WorkPlace)AddFile(path string, created time.Time)  bool{
	  file,estado := f.SearchFile(path)
	if estado {
		f.MakeVersion(file,created,path)
	}else{
		f.MakeFileAndVersion(path,created)
		}
	  return true
}
func (place *WorkPlace) SearchFile(s string) (Archivo File,Estado bool) {
	for _, element := range place.ListnerFolder {
		for _,ele :=range element.Files{
			//fmt.Println(ele.Name)
				if strings.Contains(s, ele.Name) {
					return ele,true
				}
			}
		}
	return File{},false
}
func (f *WorkPlace) MakeVersion(file File, created time.Time, s string) {
	dat, err := ioutil.ReadFile(s)
	Utils.Check(err)
	NewVersion :=MakeVersion(created,dat)
	f.NewVersionInThisFile(file.Id,NewVersion)

	}
func MakeVersion(created time.Time, bytes []byte) Version {
	return Version{Content:bytes,Version:created,Name:uuid.New(),Status:true}
}
func (f *WorkPlace) MakeFileAndVersion(s string, created time.Time){
	var Carpeta Folder
	Carpeta.Name = Carpeta.GetFolderName(s)
	Carpeta.Path = Carpeta.GetPath(s)
	Readfile, err := os.Open(s)
	Utils.Check(err)
	Carpeta.Files = append(Carpeta.Files, MakeFIrstFileAndVersion(Readfile,s))
	f.ListnerFolder = append(f.ListnerFolder,Carpeta)

}

func MakeFIrstFileAndVersion(file *os.File, s string) File{
	data,_ := file.Stat()
	return MakeEntity(data,s)
}

func (f *WorkPlace) NewVersionInThisFile(uuids uuid.UUID, version Version) {
	for idfolder,folder :=  range f.ListnerFolder{
		for idfile,file :=  range folder.Files{
			if file.Id == uuids {
				f.ListnerFolder[idfolder].Files[idfile].Versiones=	append(f.ListnerFolder[idfolder].Files[idfile].Versiones,version)
				return
			}
		}
	}
	f.DisablesaAllOldVersion(uuids)
	}
func (f *WorkPlace) ChargeFilesInSystem(s string) {
	var Carpeta Folder
	Carpeta.Name = Carpeta.GetFolderName(s)
	Carpeta.Path = s
	files, err := ioutil.ReadDir(s)
	Utils.Check(err)
	for _, file := range files {
		if strings.Contains(filepath.Ext(file.Name()), ".") {
			ThisF := MakeEntity(file,s+string(os.PathSeparator)+file.Name())
			Carpeta.Files = append(Carpeta.Files,ThisF)
		}else {
			 f.ChargeFilesInSystem(s+string(os.PathSeparator)+file.Name())
			 }
	}
	f.ListnerFolder = append(f.ListnerFolder, Carpeta)
	}
func (f *Folder) GetFolderName(s string) string {
	P := strings.Split(s, string(os.PathSeparator))
	return P[len(P) - 1]
	}

func (f *Folder) GetPath(s string) string {
	var Out string
	P := strings.Split(s, string(os.PathSeparator))
	for _,path := range P {
		if !strings.Contains(path,f.Name) {
			Out += path+ string(os.PathSeparator)
		}
		}
	Out = Utils.TrimSuffix(Out,string(os.PathSeparator))

	return Out
	}


func (f *WorkPlace) DisablesaAllOldVersion(uuids uuid.UUID) {
	for idfolder,folder :=  range f.ListnerFolder{
		for idfile,file :=  range folder.Files{
			if file.Id == uuids {
				howManyVersion := len(file.Versiones) - 1
					for idversiones,_ := range file.Versiones{
						if idversiones < howManyVersion {
							f.ListnerFolder[idfolder].Files[idfile].Versiones[idversiones].Status = false
							return
							}
				}
			}
		}
	}
}
func MakeEntity(info os.FileInfo,Path string)File{
	var file File
	file.Name = info.Name()
	file.Type = file.Whois()
	file.Id = uuid.New()
	file.LastUpdate = time.Now()
	dat, err := ioutil.ReadFile(Path)
	Utils.Check(err)
	file.Versiones = append(file.Versiones, MakeVersion(time.Now(),dat))
	return file
}

