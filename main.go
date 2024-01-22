package main

import (
	"fmt"

	"github.com/godesde0/ejercicios"
)

/*
func main() {
	estado, texto := variables.ConviertoTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s", os)
	}
}
*/
/*
func main() {
	switch numero, palabra := ejercicios.ConviertoString("130"); numero < 100 {
	case true:
		fmt.Println("es menor que 100")
	default:
		fmt.Println(palabra + " es mayor o igual a 100")
	}
}
*/

func main() {
	numero := ejercicios.IngresoPorTeclado()
	for i := 0; i < 11; i++ {
		fmt.Println(numero * i)
	}
}
