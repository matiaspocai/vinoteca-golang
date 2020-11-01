package main

import "fmt"

type vino struct {
	nombre string
	marca  string
	precio int
}

var lista []vino

func crearVino(nom, mar string, pre int) {
	v := vino{nom, mar, pre}
	lista = append(lista, v)
}

func imprimirLista() {
	fmt.Println("Imprimir lista de vinos: ")
	for i, v := range lista {
		fmt.Println(i, v)
	}
}

func actualizarVino(nom, mar string, pre, pos int) {
	v := vino{nom, mar, pre}
	lista[pos] = v
}

func eliminarVino(pos int) {
	p2 := pos + 1
	lista = append(lista[:pos], lista[p2:]...)

}

func main() {

	crearVino("Elite", "Super24", 545)
	crearVino("Sutil", "NewDeal", 245)
	actualizarVino("Magic", "Speed", 178, 1)
	crearVino("Sampe", "Full", 150)
	imprimirLista()
	eliminarVino(0)
	imprimirLista()

}
