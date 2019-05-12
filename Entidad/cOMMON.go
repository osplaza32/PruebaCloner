package Entidad

import (
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"osplaza32/Prueba_Cloner/Utils"
	"time"
)

func MakeEntity(info os.FileInfo,Path string,Momento time.Time)File{
	var file File
	file.Name = info.Name()
	file.Type = file.Whois()
	file.Id = uuid.New()
	file.Status= true
	file.LastUpdate = Momento
	dat, err := ioutil.ReadFile(Path)
	Utils.Check(err)
	file.Versiones = append(file.Versiones, MakeVersion(Momento,dat))
	return file
}
func MakeVersion(created time.Time, bytes []byte) Version {
	return Version{Content:bytes,Version:created,Name:uuid.New(),Status:true}
}

func MakeFIrstFileAndVersion(file *os.File, s string,created time.Time) File{
	data,err := file.Stat()
	Utils.Check(err)
	return MakeEntity(data,s,created)
}
