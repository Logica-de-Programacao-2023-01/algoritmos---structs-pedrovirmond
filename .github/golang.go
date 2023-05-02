// Crie uma struct chamada Aluno com os campos "nome" , "idade" e "notas".
// O campo "notas" deve ser um slice de float64. Escreva uma função que receba um Aluno como
// parâmetro e calcule a média das suas notas.
package main

type Aluno struct {
	nome  string
	idade int
	notas float64
}
type Notas struct {
	nota1 float64
	nota2 float64
	nota3 float64
}

func main() {
	a := Aluno{
		nome:  "Wellington",
		idade: 18,
		notas: Notas{
			nota1: 67.43,
			nota2: 54.94,
			nota3: 83.71,
		},
	}

}
func UmAluno(a Aluno) float64 {
	return (a.notas.nota1 + a.notas.nota2 + a.notas.nota3) / 3
}

}
