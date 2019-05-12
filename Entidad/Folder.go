package Entidad
import (
	"github.com/google/uuid"
	"os"
	"osplaza32/Prueba_Cloner/Utils"
	"strings"
	"time"
)
type Folder struct {
	Name string
	Files []File
	Path string
	Id uuid.UUID
	Created time.Time
	Deleted time.Time
	Status bool
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
func (f *Folder) new(s string, now time.Time) Folder {
	return Folder{Created:now,Id:uuid.New(),Name:f.GetFolderName(s),Path:f.GetPath(s)}
}
