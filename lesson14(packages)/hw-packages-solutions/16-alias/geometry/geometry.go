package geometry

// Pi - своя константа "числа Пи", хуже стандартного math.Pi по точности,
// но специально названа так же, чтобы показать конфликт имён пакетов.
const Pi = 3.14

// CircleArea считает площадь круга через собственный Pi.
func CircleArea(radius float64) float64 {
	return Pi * radius * radius
}
