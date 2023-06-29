package main

type Product struct {
	Code      int
	Name      string
	UnitPrice float64
}

type Sale struct {
	Code        int
	Date        string
	TotalAmount float64
	Products    []Product
}

// Função para calcular o valor total de todas as vendas
func calculateTotalSales(sales map[int]*Sale) float64 {
	total := 0.0

	for _, sale := range sales {
		total += sale.TotalAmount
	}

	return total
}
