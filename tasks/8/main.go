package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
func main() {
	var num int64
	var i int64
	var b int64
	fmt.Println("Enter decimal number")
	fmt.Scanf("%d", &num)
	fmt.Printf("Binary representation: %b\n", uint64(num))
	fmt.Println("Enter position of byte to change")
	fmt.Scanf("%d", &i)
	fmt.Println("Enter new byte value (1 or 0)")
	fmt.Scanf("%d", &b)
	
	/*
	Пример: num = 10 (1010 в двоич.), i = 2
	Справа выполняется побитовый сдвиг (отсчет с 0)
	(1 << i) - сдвигаем 1 на i позиций влево = 100 (в двоич.)
	Далее выполняется исключающее ИЛИ. Возвращает 1, если только один из соответствующих разрядов обоих чисел равен 1
	1010 ^ 0100 = 1110
	*/

	num ^= (1 << i)

    fmt.Printf("New value: %d\n", num)
	fmt.Printf("Binary representation: %b\n", uint64(num))
}