package main

import (
	"fmt"
	"os"
	"osplaza32/Prueba_Cloner/Utils"
)

func main() {
	var str = `C:\Users\Administrador\go\src\osplaza32\Prueba_Cloner\Evidencia debug\`
	str = Utils.TrimSuffix(str,string(os.PathSeparator))
	fmt.Println(str)
}