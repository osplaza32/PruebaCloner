package main

import (
	"github.com/subosito/gotenv"
	"os"
	"osplaza32/Prueba_Cloner/Entidad"
	"time"
)

func main() {
		gotenv.Load(".env")
		pathOftest := os.Getenv("ENV_PATH")
		var Work Entidad.WorkPlace
	 	Work.ChargeFilesInSystem(pathOftest)
	 	//Archivo el mismo nombre de uno ya creado
		Work.AddFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\ddll.txt",time.Now())
	 	//Archivo nuevo con carpeta no registrada crea la carpeta
		Work.AddFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\1.- Debug con el ide y la comprobacion de cada carpeta tiene su archivo con la recusrividad.PNG",time.Now())
		// Crea Carpeta Vacia y nueva
		Work.AddFolder("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\CarpetaSinNada",time.Now())
	}
