package main

import "fmt"

func dibujarFigura(cantidad int) {
	for inicio := 1; inicio <= cantidad; inicio++ {
		for segundo := cantidad - inicio; segundo > 0; segundo-- {
			fmt.Print(" ")
		}
		// empieza a dibujar la parte superior
		for tercero := 0; tercero < inicio; tercero++ {
			fmt.Print(" *")
		}
		inicio++ // se le suma uno màs para que al volver al ciclo comience con un valor impar.
		fmt.Println("")
	}
	for inicio := 1; inicio <= cantidad; inicio++ {
		for segundo := 0; segundo <= inicio; segundo++ {
			fmt.Print(" ")
		}
		// empieza a dibujar la parte inferior
		for tercero := cantidad - inicio - 1; tercero > 0; tercero-- {
			fmt.Print(" *")
		}
		inicio++ // se le suma uno màs para que al volver al ciclo comienza con un valor impar.
		fmt.Println("")
	}
}
func main() {
	fmt.Println("Se imprime una figura apartir de puntos impares. \n")
	var cantidadPuntos = 5
	dibujarFigura(cantidadPuntos)
}

/*    -- Código en Ejecución --

Se imprime una figura apartir de puntos impares.

     *
   * * *
 * * * * *
   * * *
     *


Process finished with the exit code 0



*/
