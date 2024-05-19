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
	s := [1]interface{}{}
	n := [2]int{1, 2}

	s[0] = make([]int, 0)
	s[0] = []int{}
	s[0] = n[:]
	s[0] = new([]int)
	var se []int
	s[0] = se


	s[0] = make(map[string]int)
	s[0] = map[string]int{}
	s[0] = new(map[string]int)
	var md map[string]int
	s[0] = md
}