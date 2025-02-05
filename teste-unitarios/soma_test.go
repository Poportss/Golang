package main

import "testing"

func TestSoma(t *testing.T) { //ShouldSumCorrect
	//arrange
	teste := soma(3, 2, 1)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestSoma2(t *testing.T) { //ShouldSumIncorrect
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestMultiplica(t *testing.T) { //ShouldMultCorrect

	teste := multiplica(10, 10)

	resultado := 100

	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestMultiplica2(t *testing.T) { //ShouldMultIncorrect

	teste := multiplica(10, 10)

	resultado := 10

	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}
