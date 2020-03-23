package ponteiro

import "fmt"

type Carro struct {
	cor    string
	modelo string
	valor  int
}

func Main() {
	Fiat := Carro{
		cor:    "Preto",
		modelo: "Palio",
		valor:  120000,
	}

	fmt.Printf("Fiat Palio: %+v\r\n", Fiat)

	mudaCarro(&Fiat)

	fmt.Printf("Fusca: %+v\r\n", Fiat)

}

func mudaCarro(carro *Carro) {
	carro.cor = "Azul"
	carro.modelo = "Fusca"
}
