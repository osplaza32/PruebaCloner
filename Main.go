package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/subosito/gotenv"
	"io/ioutil"
	"os"
	"osplaza32/Prueba_Cloner/Entidad"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	gotenv.Load(".env")
	pathOftest := os.Getenv("ENV_PATH")
	var Work Entidad.WorkPlace
	 materia := ChargeFilesInSystem(pathOftest,Work)
	 fmt.Println(materia)
	}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ChargeFilesInSystem(s string,W Entidad.WorkPlace) Entidad.WorkPlace {
	var Carpeta Entidad.Folder
	Carpeta.Name = Carpeta.GetFolderName(s)
	Carpeta.Path = s
	files, err := ioutil.ReadDir(s)
	check(err)
	for _, file := range files {
		//MakeEntity(file)
		if strings.Contains(filepath.Ext(file.Name()), ".") {
			ThisF := MakeEntity(file,s+string(os.PathSeparator)+file.Name())
			Carpeta.Files = append(Carpeta.Files,ThisF)

		}else {
			a := ChargeFilesInSystem(s+string(os.PathSeparator)+file.Name(),W)
			W=a

		}
	}
	W.ListnerFolder = append(W.ListnerFolder, Carpeta)
	return W
}

func MakeEntity(info os.FileInfo,Path string)Entidad.File{
	var file Entidad.File
	file.Name = info.Name()
	file.Type = file.Whois()
	dat, err := ioutil.ReadFile(Path)
	check(err)
	file.Versiones = append(file.Versiones, MakeInitVersion(dat, time.Now(),uuid.New(), true))

	return file
}

func MakeInitVersion(bytes []byte, now time.Time, s uuid.UUID, b bool) Entidad.Version {
	return Entidad.Version{Content:bytes,Version:now,Status:b,Name:s}
}
