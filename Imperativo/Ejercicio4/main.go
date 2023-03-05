package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type producto struct {
	nombre      string
	descripcion string
	montoVenta  int32
}

type Productos []producto

var factura Productos

type Persona struct {
	nombre string
	edad   int
}
type Personas []Persona

var listaPersonas Personas

const rangoPagoImpuestos = 20000
const porcentajeImpuesto = 0.13

func (p *Personas) agregarPersonas(nombre string, edad int) {
	*p = append(*p, Persona{nombre, edad})
}

func (f *Productos) agregarProducto(nom string, desc string, pre int32) {
	idx := slices.IndexFunc(*f, func(p producto) bool { return p.nombre == nom })
	if idx == -1 {
		*f = append(*f, producto{nom, desc, pre})
	} else {
		fmt.Println("cant add existing product to the list")
	}
}

func (f *Productos) calcularImpuestoFactura() int32 {
	lista := filter2(*f, func(p producto) bool {
		return p.montoVenta > rangoPagoImpuestos
	})
	lista2 := map2(lista, func(p producto) int32 {
		return int32(float64(p.montoVenta) * porcentajeImpuesto)
	})
	return reduce(lista2)
}

func (f *Productos) calcularMontoFactura() int32 {
	lista := map2(*f, func(p producto) int32 {
		return p.montoVenta
	})
	return reduce(lista)
}

// funciones map y filter para aplicaciones específicas
func map1(list Productos, f func(producto) int32) []int32 {
	mapped := make([]int32, len(list))

	for i, e := range list {
		mapped[i] = f(e)
	}
	return mapped
}

func filter1(list Productos, f func(producto) bool) Productos {
	filtered := make(Productos, 0)

	for _, element := range list {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func reduce(list []int32) int32 {
	var result int32 = 0
	for _, v := range list {
		result += v
	}
	return result
}

func (p *Personas) calcularEdadPersonas() []Persona {
	lista := filter2(*p, func(p Persona) bool {
		return p.edad >= 18
	})
	return lista
}

// filters- maps de tipo Any

func map2[P1, P2 any](lista []P1, f func(P1) P2) []P2 {
	mapped := make([]P2, len(lista))
	for a, b := range lista {
		mapped[a] = f(b)
	}
	return mapped
}

func filter2[P1 any](lista []P1, f func(P1) bool) []P1 {
	filtered := make([]P1, 0)
	for _, element := range lista {
		if f(element) {
			filtered = append(filtered, element)
		}
	}
	return filtered
}

func main() {
	factura.agregarProducto("tarjeta madre", "Asus", 54200)
	factura.agregarProducto("mouse", "alámbrico", 15000)
	factura.agregarProducto("teclado", "gamer con luces", 30000)
	factura.agregarProducto("memoria ssd", "2 gb", 65000)
	factura.agregarProducto("cable video", "display port 1m", 18000)
	fmt.Print("Costo de la Factura: ")
	println(factura.calcularMontoFactura())
	fmt.Print("Impuestos aplicados: ")
	println(factura.calcularImpuestoFactura())

	fmt.Println("--------------------------")
	listaPersonas.agregarPersonas("Diego", 20)
	listaPersonas.agregarPersonas("Juan", 25)
	listaPersonas.agregarPersonas("Ana", 26)
	listaPersonas.agregarPersonas("Manolo", 15)
	listaPersonas.agregarPersonas("Pedro", 20)
	listaPersonas.agregarPersonas("Gloria", 18)
	listaPersonas.agregarPersonas("Maria", 24)
	listaPersonas.agregarPersonas("Sara", 16)
	listaPersonas.agregarPersonas("Pablo", 28)
	listaPersonas.agregarPersonas("Pirlo", 12)

	fmt.Println("Lista Personas", listaPersonas)
	fmt.Println("")
	fmt.Println("Lista de personas mayores de edad")
	fmt.Println(listaPersonas.calcularEdadPersonas())

}

/* Código en Ejecución

Costo de la Factura: 182200
Impuestos aplicados: 19396
--------------------------
Lista Personas [{Diego 20} {Juan 25} {Ana 26} {Manolo 15} {Pedro 20} {Gloria 18} {Maria 24} {Sara 16} {Pablo 28} {Pirlo 12}]

Lista de personas mayores de edad
[{Diego 20} {Juan 25} {Ana 26} {Pedro 20} {Gloria 18} {Maria 24} {Pablo 28}]

Process finished with the exit code 0

*/
