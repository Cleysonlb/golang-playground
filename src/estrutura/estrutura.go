package estrutura

import (
	"encoding/json"
	"fmt"
)

type Imovel struct {
	Nome   string `json:"nome"`
	Bairro string `json:"bairro"`
	Valor  int    `json:"valor"`
}

func Main() {
	casa := Imovel{
		Nome:   "Minha casa",
		Bairro: "Laranjeiras",
		Valor:  2500000,
	}

	fmt.Printf("Casa a venda: %+v\r\n", casa)
	casa.SetValor(999)
	fmt.Printf("GetValor da casa: %+v\r\n", casa.GetValor())
	objJson, _ := json.Marshal(casa)
	fmt.Printf(string(objJson))

}

func (i *Imovel) SetValor(valor int) {
	i.Valor = valor
}

func (i *Imovel) GetValor() int {
	return i.Valor
}
