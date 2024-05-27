package main

/*
Вопрос: Сколько существует способов задать переменную типа slice или map?

Ответ:
1. Используя make
2. Используя композитный литерал []type{}
3. Для slice - путем разрезания массива или другого slice
4. Используя new
5. Через определение неинициализированной переменной
*/
func main() {
	n := [2]int{1, 2}

	_ = make([]int, 0)
	_ = []int{}
	_ = n[:]
	_ = new([]int)
	var _ []int

	_ = make(map[string]int)
	_ = map[string]int{"a": 1, "b": 2, "c": 3}
	_ = new(map[string]int)
	var _ map[string]int
}