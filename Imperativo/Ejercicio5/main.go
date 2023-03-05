package main

import (
	"fmt"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	validar := lProductos.buscarProducto(nombre)
	if validar == -1 {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	} else {
		lProductos.IncrementarProductos(nombre, cantidad, precio)

	}
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cantidadVender int) {
	var ListaAuxiliar listaProductos
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		nom := (*l)[prod].nombre
		fmt.Println("El producto a vender es: ", nombre, ",y la cantidad solicitada es: ", cantidadVender)
		if (*l)[prod].cantidad < cantidadVender {
			fmt.Println("No hay suficientes productos.. se le ha vendido una cantidad de:", ((*l)[prod].cantidad), "\n")
			var i int
			for i = 0; i < len(*l); i++ {
				if i == prod { // se salta la posicion del calzado a eliminar
					continue
				} else { // se guardan temporal en otra lista
					ListaAuxiliar = append(ListaAuxiliar, producto{nombre: (*l)[i].nombre, cantidad: (*l)[i].cantidad, precio: (*l)[i].precio})
				}
			}
			fmt.Println("El producto", nom, "ha sido eliminado de la lista..")
			lProductos = ListaAuxiliar

		} else { // se procede a vender y a eliminar del stock
			var i int
			for i = 0; i < len(*l); i++ {
				if i == prod { // se salta la posicion del calzado a eliminar
					continue
				} else { // se guardan temporal en otra lista
					ListaAuxiliar = append(ListaAuxiliar, producto{nombre: (*l)[i].nombre, cantidad: (*l)[i].cantidad, precio: (*l)[i].precio})
				}
			}
			fmt.Println("El producto ", nom, ", ha sido eliminado de la lista..")
			lProductos = ListaAuxiliar
		}

	} else {
		fmt.Println("El producto", nombre, "no se encuentra en la lista....")

	}

}
func (l *listaProductos) listarProductosMínimos() listaProductos { // debe retornar una nueva lista con productos con existencia mínima
	var listaAuxiliar listaProductos
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].cantidad < existenciaMinima {
			listaAuxiliar = append(listaAuxiliar, producto{nombre: (*l)[i].nombre, cantidad: (*l)[i].cantidad, precio: (*l)[i].precio})

		} else {
			continue
		}
	}
	return listaAuxiliar
}

func (l *listaProductos) AumentarInventarioDeminimos(listaMinimos listaProductos) { // se incrementa la cantidad de productos
	var i int
	producto := 0
	for i = 0; i < len(listaMinimos); i++ {
		producto = l.buscarProducto((listaMinimos)[i].nombre)
		(*l)[producto].cantidad = existenciaMinima
	}
}

func (l *listaProductos) IncrementarProductos(nombre string, cantidad int, precio int) {
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			(*l)[i].cantidad += cantidad
			if (*l)[i].precio != precio {
				(*l)[i].precio = precio
			}
		}
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 10, 1200)
	lProductos.agregarProducto("papas", 6, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("maiz", 3, 3500)
	lProductos.agregarProducto("pan", 7, 4500)
}
func main() {
	llenarDatos()
	fmt.Println()
	fmt.Println("Lista de productos:", lProductos)
	fmt.Println("................................")
	lProductos.venderProducto("leche", 12)
	fmt.Println(lProductos)
	fmt.Println("...................................")
	fmt.Println("Lista de productos con una cantidad menor a", existenciaMinima, ": ", lProductos.listarProductosMínimos())
	lista := lProductos.listarProductosMínimos()  // se busca la lista con productos minimos
	lProductos.AumentarInventarioDeminimos(lista) // Se envia por parametro
	fmt.Println("Se ha aumentado los productos con existencias minimas: ", lProductos)
}

/*
Lista de productos: [{arroz 15 2500} {frijoles 4 2000} {leche 10 1200} {papas 6 1200} {café 12 4500} {maiz 3 3500} {pan 7 4500}]
................................
El producto a vender es:  leche ,y la cantidad solicitada es:  12
No hay suficientes productos.. se le ha vendido una cantidad de: 10

El producto leche ha sido eliminado de la lista..
[{arroz 15 2500} {frijoles 4 2000} {papas 6 1200} {café 12 4500} {maiz 3 3500} {pan 7 4500}]
...................................
Lista de productos con una cantidad menor a 10 :  [{frijoles 4 2000} {papas 6 1200} {maiz 3 3500} {pan 7 4500}]
Se ha aumentado los productos con existencias minimas:  [{arroz 15 2500} {frijoles 10 2000} {papas 10 1200} {café 12 4500} {maiz 10 3500} {pan 10 4500}]

Process finished with the exit code 0

*/
