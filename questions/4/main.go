package main

/*
Вопрос: Чем отличаются буферизированные и не буферизированные каналы?

Ответ:
По умолчанию канал создается с буфером 0 - все, что пишется в канал, сразу доступно для чтения, и каждая операция отправки в канал блокирует горутину
Когда размер буфера больше 0, горутина не блокируется до тех пор, пока буфер не будет заполнен.
Операция чтения на буфферизированном канале не будет завершена до тех пор, пока буфер не будет опустошен.
Т.е. горутина будет считывать буфер канала без блокировки до тех пор, пока буфер не станет пустым.
Объявление буферизированного канала: c := make(chan Type, n)
Текущая горутина не будет заблокирована, пока в канал не будет передано n+1 данных.

При записи - горутина заблокируется, когда буфер будет заполнен (n + 1)
При чтении - горутина заблокируется, когда буфер опустеет
С буферизированным каналом возможно использовать цикл for range на закрытом канале, при этом будут считаны данные из буфера
*/

import (
	"fmt"
	"time"
)

func sender(c chan int) {
    c <- 1 // len 1, cap 3
    c <- 2 // len 2, cap 3
    c <- 3 // len 3, cap 3
    c <- 4 // <- goroutine blocks here
    close(c)
}

func reader(c chan int) {
	fmt.Println(<- c)
}

func main() {
    c := make(chan int, 3)

    go sender(c)

    fmt.Printf("Length of channel c is %v and capacity of channel c is %v\n", len(c), cap(c))

    go reader(c)

	time.Sleep(5 * time.Second)
}