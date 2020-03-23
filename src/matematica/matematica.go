package matematica

func Multiplicadores(x int, y int) int {
	return x * y
}

func Soma(x int, y int) int {
	return x + y
}

func Divisor(paramA int, paramB int) (resultadoDivisao int) {
	resultadoDivisao = paramA / paramB
	return
}

func DivisorComResto(paramA int, paramB int) (resultado int, resto int) {
	resultado = paramA / paramB
	resto = paramA % paramB
	return
}

func Calculo(funcao func(int, int) int, paramA int, paramB int) int {
	return funcao(paramA, paramB)
}
