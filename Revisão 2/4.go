package main

type Student struct {
	ID      int
	Name    string
	Grades  map[string]float64
	Average float64
}

// Função para atualizar a média geral de cada estudante
func updateAverage(students map[int]*Student) {
	for _, student := range students {
		sum := 0.0
		count := 0

		for _, grade := range student.Grades {
			sum += grade
			count++
		}

		student.Average = sum / float64(count)
	}
}
