package Entidad
import (
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"osplaza32/Prueba_Cloner/Utils"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)
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
func (f *WorkPlace) MakeFileAndVersion(s string, created time.Time){
	var Carpeta Folder
	Carpeta.Id = uuid.New()
	Carpeta.Name = Carpeta.GetFolderName(s)
	Carpeta.Path = Carpeta.GetPath(s)
	Carpeta.Created = created
	Carpeta.Status = true
	Readfile, err := os.Open(s)
	Utils.Check(err)
	Carpeta.Files = append(Carpeta.Files, MakeFIrstFileAndVersion(Readfile,s,created))
	f.ListnerFolder = append(f.ListnerFolder,Carpeta)
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
	Carpeta.Id = uuid.New()
	Carpeta.Name = Carpeta.GetFolderName(s)
	Carpeta.Path = s
	Carpeta.Created = time.Now()
	Carpeta.Status = true

	files, err := ioutil.ReadDir(s)
	Utils.Check(err)
	for _, file := range files {
		if strings.Contains(filepath.Ext(file.Name()), ".") {
			ThisF := MakeEntity(file,s+string(os.PathSeparator)+file.Name(),time.Now())
			Carpeta.Files = append(Carpeta.Files,ThisF)
		}else {
			f.ChargeFilesInSystem(s+string(os.PathSeparator)+file.Name())
		}
	}
	f.ListnerFolder = append(f.ListnerFolder, Carpeta)
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
func (f *WorkPlace) AddFolder(s string, now time.Time) {
	var Carpeta Folder
	f.ListnerFolder = append(f.ListnerFolder,Carpeta.new(s,now))
}
func (f *WorkPlace) DeleteFile(s string,delete time.Time) string{
	idcarpeta,idfile,estado:=f.FileExist(s)
	if estado{
		id1, err := strconv.Atoi(idcarpeta)
		Utils.Check(err)
		id2, err := strconv.Atoi(idfile)
		Utils.Check(err)
		f.ListnerFolder[id1].Files[id2].Status = false
		f.ListnerFolder[id1].Files[id2].Deleteddate = delete
		f.DeleteAllVerions(id1,id2)
		return "Archivo Destruido puede volver a la vida cuando uds estime"
	}
	return "no Eliminado\no encontrado"
}
func (f *WorkPlace) FileExist(s string) (idcarpeta string,idfile string, estado bool) {
	for idFolder,Folder:= range f.ListnerFolder{
		for idFile,File	:= range Folder.Files{
			if strings.Contains(s,Folder.Path+string(os.PathSeparator)+File.Name){
				return strconv.Itoa(idFolder),strconv.Itoa(idFile),true
			}
		}
	}
	return "nil","nil",false
}
func (f *WorkPlace) DeleteAllVerions(IdFolder int, IdFile int) {
	for idVersion,_ := range f.ListnerFolder[IdFolder].Files[IdFile].Versiones{
		f.ListnerFolder[IdFolder].Files[IdFile].Versiones[idVersion].Status = false
		f.ListnerFolder[IdFolder].Files[IdFile].Versiones[idVersion].Deleted = time.Now()
	}
}
func (f *WorkPlace) DeleteFolder(s string, now time.Time) string {
	Idfolder,Estado := f.ThisFolderExist(s)
	if Estado {
		id1, err := strconv.Atoi(Idfolder)
		Utils.Check(err)
		f.ListnerFolder[id1].Status = false
		f.ListnerFolder[id1].Deleted = now
		f.deleteAllFile(id1)
		return  "Carpeta eliminada sus archivos y versiones puede traerla en la vida cuando estime"
	}
	return "Carpeta no encontada"
}
func (f *WorkPlace) ThisFolderExist(s string)(IdFolder string,Estado bool){
	for idFolder,Folder:= range f.ListnerFolder{
		if strings.Contains(s,Folder.Path){
			return strconv.Itoa(idFolder),true
		}
	}
	return "nil",false
}
func (f *WorkPlace) deleteAllFile(i int) {
	for idFile,_:= range f.ListnerFolder[i].Files{
		f.ListnerFolder[i].Files[idFile].Status=false
		f.ListnerFolder[i].Files[idFile].Deleteddate=time.Now()
		f.DeleteAllVerions(i,idFile)
	}
	}