package enderecos_test

import (
	. "AutomatedTests/enderecos"
	"testing"
)

type testScenario struct {
	inputAddress   string
	expectedReturn string
}

func TestAddressType(t *testing.T) {
	t.Parallel()
	scenariosTest := []testScenario{
		{"Rua Dos Bobos", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada Velha", "Estrada"},
		{"RODOVIA DOS IMIGRANTES", "Rodovia"},
		{"RUA SEM NOME", "Rua"},
		{"AVENIDA CENTRAL", "Avenida"},
		{"estrada do sol", "Estrada"},
		{"rodovia abc", "Rodovia"},
		{"praca da sé", "Tipo inválido"},
		{" ", "Tipo inválido"},
		{"", "Tipo inválido"},
	}

	for _, scenario := range scenariosTest {
		if AddressType(scenario.inputAddress) != scenario.expectedReturn {
			t.Errorf("\nTipo recebido: '%s'. \nÉ diferente do tipo esperado '%s'",
				scenario.inputAddress, scenario.expectedReturn)
		}
	}
}

func TestTest(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste quebrou!!!")
	}
}

func TestTest2(t *testing.T) {
	t.Parallel()
	if 2 > 3 {
		t.Error("Teste quebrou!!!")
	}
}
