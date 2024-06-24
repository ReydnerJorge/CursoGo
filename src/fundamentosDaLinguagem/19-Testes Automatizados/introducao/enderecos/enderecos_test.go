package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"", "Tipo Inválido"},
		{"Avenida ABC", "Avenida"},
		{"Avenida JK", "Avenida"},
		{"RUA Iub", "Rua"},
		{"Rua Alavez", "Rua"},
		{"Estrada L", "Estrada"},
		{"Estrada Pantaneira", "Estrada"},
		{"RODOVIA A", "Rodovia"},
		{"Rodovia B", "Rodovia"},
		{"Rua C", "Rua"},
		{"ESTRADA R", "Estrada"},
		{"Praça Ourinhos", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}

// Teste de Unidade
// enderecoParaTeste := "Rua ABC"
// tipoDeEnderecoEsperado := "Avenida"
// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 	t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
// 		tipoDeEnderecoEsperado,
// 		tipoDeEnderecoRecebido,
// 	)
// }
