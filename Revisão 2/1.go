package main

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

// Função para mesclar os dados dos alunos de dois mapas em um novo mapa
func mergeStudentData(studentData1, studentData2 map[string]Student) map[string]Student {
	mergedData := make(map[string]Student)

	// Copiar os dados do studentData1 para o mergedData
	for key, student := range studentData1 {
		mergedData[key] = student
	}

	// Atualizar ou adicionar os dados do studentData2 no mergedData
	for key, student := range studentData2 {
		if _, ok := mergedData[key]; ok {
			// O aluno já existe no mergedData, então atualiza as matérias e notas
			for subject, grade := range student.Subjects {
				mergedData[key].Subjects[subject] = grade
			}
		} else {
			// O aluno não existe no mergedData, então adiciona todos os dados
			mergedData[key] = student
		}
	}

	return mergedData
}
