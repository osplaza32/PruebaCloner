package Entidad

import (
	"github.com/google/uuid"
	"strings"
	"time"
)
type File struct {
	Name        string
	Versiones   []Version
	Type        string
	Id          uuid.UUID
	LastUpdate  time.Time
	Status      bool
	Deleteddate time.Time
}
func (file File) Whois() string {
	a :=strings.SplitAfter(file.Name, ".")
	return a[len(a) - 1]
}



