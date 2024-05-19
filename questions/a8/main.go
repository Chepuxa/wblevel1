package main

import "fmt"

/*
Вопрос: В чем разница make и new?

Ответ:
new - выделяет память и присваивает нулевое значение для типа, возвращая указатель.
Т.к. ссылочные типы (slice, map, chan) - указатели на лежащую в них структуру, при использовании new она не будет инициализирована и вернувшийся указатель не будет на что-то ссылаться
make - используется только для создания slice, chan и map. Выделяет память под переданный тип, а также инициализирует внутрилежащие структуры. Возвращает тот же тип, что и был передан, а не указатель

make больше похож на композитный литерал, когда new - на неинициализированную переменную
make помимо типа принимает численные значения, указывающие, в зависимости от типа, на len и cap создаваемой структуры
*/
func main() {
    fmt.Println("-- MAKE --")
    a := make([]int, 0)
    aPtr := &a
    fmt.Println("pointer == nil :", *aPtr == nil)
    fmt.Printf("pointer value: %p\n\n", *aPtr)

    fmt.Println("-- COMPOSITE LITERAL --")
    b := []int{}
    bPtr := &b
    fmt.Println("pointer == nil :", *bPtr == nil)
    fmt.Printf("pointer value: %p\n\n", *bPtr)

    fmt.Println("-- NEW REFERENCE TYPE --")
    cPtr := new([]int)
    fmt.Println("pointer == nil :", *cPtr == nil)
    fmt.Printf("pointer value: %p\n\n", *cPtr)

    fmt.Println("-- VAR REFERENCE TYPE (not initialized) --")
    var d []int
    dPtr := &d
    fmt.Println("pointer == nil :", *dPtr == nil)
    fmt.Printf("pointer value: %p\n\n", *dPtr)


    fmt.Println("-- NEW PRIMITIVE --")
    ePtr := new(int)
    fmt.Println("pointer == nil :", ePtr == nil)
    fmt.Printf("pointer value: %v\n", ePtr)
}