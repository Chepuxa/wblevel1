package _test

import (
	"bytes"
	"strings"
	"testing"
)

/*
Вопрос: Какой самый эффективный способ конкатенации строк?

Ответ: 
1. Самый очевидный способ - с использованием +. Он плох из-за того, что строки неизменяемы, и при конкатенации будет аллоцироваться память под новую строку
2. Использование buffer - структура, в которой заложен слайс байтов. Предоставляет метод WriteString, записывающий строку в этот слайс. 
   Если известна длина конечной строки, можно избежать лишних аллокаций памяти, заранее вызвав Grow()
3. Builder появился в Go 1.10, и использует идентичный с buffer подход. В структуре хранится слайс байтов, куда и записываются строки при конкатенации. Также предоставляется метод Grow()
   По сравнению с buffer-ом, Builder имеет:
   - Механизм copyCheck, предотвращающий случайное копирование билдера
   - Метод Reset() в билдере присваивает внутрилежащему слайсу nil, когда в buffer Reset() опустошает слайс, сохраняя выделенную под него память для дальнейшего переиспользования
   - Builder оптимизирован под работу со строками, buffer в свою очередь является более общим механизмом, нацеленным на работу с байтами
4. Copy позволяет быстро конкатенировать строки если заранее известен размер конечной строки, также избегая дополнительной аллокации памяти
5. Если нужно соединить набор строк, можно использовать strings.Join. Под капотом используется Builder, память аллоцируется заранее через Grow()

Используя приведенный ниже бенчмарк, наибольшую скорость в обоих случаях (размер строки известен, размер строки неизвестен) показал strings.Builder
*/

func BenchmarkConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func BenchmarkBufferAllocated(b *testing.B) {
	var buffer bytes.Buffer
	buffer.Grow(b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); buffer.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}


func BenchmarkStringBuilderAllocated(b *testing.B) {
	var strBuilder strings.Builder
	strBuilder.Grow(b.N)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strBuilder.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); strBuilder.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", strBuilder.String(), s)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	var strBuilder strings.Builder

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		strBuilder.WriteString("x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); strBuilder.String() != s {
		b.Errorf("unexpected result; got=%s, want=%s", strBuilder.String(), s)
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
	b.StopTimer()

	if s := strings.Repeat("x", b.N); string(bs) != s {
		b.Errorf("unexpected result; got=%s, want=%s", string(bs), s)
	}
}

func BenchmarkJoin(b *testing.B) {
	ss := []string{}
	for i := 0; i < b.N; i++ {
		ss = append(ss, "x")
	}
	b.ResetTimer()
	strings.Join(ss, "")
	b.StopTimer()
}
