/*
47. Итерационные переменные и замыкания в выражениях for

Самая распространённая проблема в Go.
Итерационные переменные в выражении for снова используются для каждой итерации.
Это значит, что каждое замыкание (aka функциональный литерал),
созданное в вашем цикле for, будет ссылаться на ту же переменную
(и они получат значение переменной в тот момент,
когда начнётся исполнение их горутин).
*/

package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data := []field{{"one"}, {"two"}, {"three"}}

	/* Некорректный способ */
	for _, v := range data {
		go v.print()
	}

	/* Правильно */
	for _, v := range data {
		vcopy := v
		go func() {
			fmt.Println(vcopy)
		}()
	}

	/* Правильно */
	for _, v := range data {
		go func(in field) {
			fmt.Println(in)
		}(v)
	}

	/* Еще интересный вариант */
	dataPtr := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range dataPtr {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}
