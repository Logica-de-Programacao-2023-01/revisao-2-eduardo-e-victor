package main

import (
	"errors"
)

type Employee struct {
	ID           int
	Name         string
	Position     string
	BaseSalary   float64
	MonthlyBonus []float64
}

// Função para calcular o salário total do funcionário
func calculateTotalSalary(employee *Employee) (float64, error) {
	if employee == nil {
		return 0.0, errors.New("employee object is nil")
	}

	// Calcula a soma dos bônus mensais
	totalBonus := 0.0
	for _, bonus := range employee.MonthlyBonus {
		totalBonus += bonus
	}

	// Atualiza a titulação do funcionário se a soma dos bônus for maior que 1500.0
	if totalBonus > 1500.0 {
		employee.Position = "Senior " + employee.Position
	}

	// Calcula o salário total
	totalSalary := employee.BaseSalary + totalBonus

	return totalSalary, nil
}
