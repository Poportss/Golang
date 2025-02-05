package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	//arrange
	teste := soma(3, 2, 1)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	//arrange
	teste := soma(3, 2, 1)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldMultiCorrect(t *testing.T) {
	//arrange
	teste := multiplica(10, 10)

	//act
	resultado := 100

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldMultiIncorrect(t *testing.T) {
	//arrange
	teste := multiplica(10, 10)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	//arrange
	teste := divide(100, 2, 5)

	//act
	resultado := 10

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldDivideIncorrect(t *testing.T) {
	//arrange
	teste := divide(100, 2, 5)

	//act
	resultado := 6

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	//arrange
	teste := subtrai(50, 10, 5)

	//act
	resultado := 35

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldSubIncorrect(t *testing.T) {
	//arrange
	teste := subtrai(50, 10, 5)

	//act
	resultado := 30

	//assert
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}
