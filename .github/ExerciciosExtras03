// Crie uma struct chamada Aluno com os campos "nome" , "idade" e "notas".
// O campo "notas" deve ser um slice de float64. Escreva uma função que receba um Aluno como
// parâmetro e calcule a média das suas notas.
package main

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}
func main() {
	a := Aluno{
		Nome:  "Wellington",
		Idade: 18,
		Notas: []float64{75.54, 86.23, 92.65, 43.71, 78.59},
	}
}fmt.Println("")

func media(a Aluno) float64 {
	var sum float64
	for _, nota := range a.Notas {
		sum += nota
	}
	return sum / float64(len(a.Notas))
}