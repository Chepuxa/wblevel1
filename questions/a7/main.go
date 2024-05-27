package main

import "fmt"

/*
Вопрос:
В какой последовательности будут выведены элементы map[int]int?

Пример:
m[0]=1
m[1]=124
m[2]=281

Ответ:
При итерации по map, значения возвращаются в случайном порядке
Это задуманное поведение, созданное для того, чтобы разработчики не полагались на итерацию по map, но даже без задуманной случайности карта является unordered
*/
func main() {
	m := make(map[int]int)
	i := 0
	j := 0
	for z := 0; z < 10; z++ {
		m[i] = j
		i++
		j++
	}

	for z := 0; z < 10; z++ {
		s := [][]int{}
		for k, v := range m {
			s = append(s, []int{k, v})
		}
		fmt.Println(s)
	}
}