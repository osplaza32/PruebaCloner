package main

import (
	"fmt"
	"github.com/subosito/gotenv"
	"os"
	"osplaza32/Prueba_Cloner/Entidad"
	"time"
)

func main() {
		gotenv.Load(".env")
		pathOftest := os.Getenv("ENV_PATH")
		var Work Entidad.WorkPlace
		//Cargar archivos de prueba
	 	Work.ChargeFilesInSystem(pathOftest)
	 	//Archivo el mismo nombre de uno ya creado se agrega una veriosn a ese archivo!
		Work.AddFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\ddll.txt",time.Now())
		Work.AddFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\ddll.txt",time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
		Work.AddFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\ddll.txt",time.Date(2010, time.November, 10, 23, 0, 0, 0, time.UTC))
		Work.AddFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\ddll.txt",time.Date(2011, time.November, 10, 23, 0, 0, 0, time.UTC))
		//Archivo nuevo con carpeta no registrada crea la carpeta
		Work.AddFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\1.- Debug con el ide y la comprobacion de cada carpeta tiene su archivo con la recusrividad.PNG",time.Now())
		// Crea Carpeta Vacia y nueva
		Work.AddFolder("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\EvidenciadebugYPruebasSobreAgregar\\CarpetaSinNada",time.Now())
		// Elimina archivo con mas de una Version
		fmt.Println(Work.DeleteFile("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\Test\\AlgoParaPensar\\ddll.txt",time.Now()))
		// Eliminar la carpeta mas poblada y cambia el estado de la carpeta , archivos y vesriones
	    fmt.Println(Work.DeleteFolder("C:\\Users\\Administrador\\go\\src\\osplaza32\\Prueba_Cloner\\Test\\AlgoParaPensar",time.Now()))

}
