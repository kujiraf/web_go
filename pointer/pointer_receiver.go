package main

import (
	"fmt"
)

type Object struct {
	X int
	Y int
}

func (o *Object) multipleWithPointerReceievr(n int) {
	o.X *= n
	o.Y *= n
	fmt.Printf("X=%d, Y=%d\n", o.X, o.Y)
}

func (o Object) multiple(n int) {
	o.X *= n
	o.Y *= n
	fmt.Printf("X=%d, Y=%d\n", o.X, o.Y)
}

func main() {
	// pointer receiverとは。pointer receiverとvalue receiverの違い
	fmt.Println("--- what is pointer receiver ---")
	whatIsPointerReceiver()

	// コンパイラがvalue型からpointer型に暗黙的に変換してくれる例
	fmt.Println("--- value to pointer ---")
	valueToPointer()

	// コンパイラがpointer型からvalue型に暗黙的に変換してくれる例
	fmt.Println("--- pointer to value---")
	pointerToValue()
}

func whatIsPointerReceiver() {
	obj := Object{1, 2}

	fmt.Printf("before multiple :%v\n", obj)
	// value receiverでは、ポインタの値は変更されない
	obj.multiple(2)
	fmt.Printf("after multiple :%v\n", obj)

	fmt.Printf("before multipleWithPointer :%v\n", obj)
	// pointer receiverでは、そのポインタの値が変更される
	obj.multipleWithPointerReceievr(2)
	fmt.Printf("after multipleWithPointer :%v\n", obj)
}

func valueToPointer() {
	obj := Object{3, 4}
	// objは値型
	fmt.Println(obj)

	fmt.Printf("before multipleWithPointer :%v\n", obj)
	// 正しくはこのように(&obj).と書くことになる。ただ、明示的にポインタで呼び出してあげなくても大丈夫。
	(&obj).multipleWithPointerReceievr(2)
	fmt.Printf("after multipleWithPointer :%v\n", obj)

	fmt.Printf("before multipleWithPointer :%v\n", obj)
	// 値型->ポインタ型の変換はコンパイラが暗黙的にやってくれる
	obj.multipleWithPointerReceievr(2)
	fmt.Printf("after multipleWithPointer :%v\n", obj)
}

func pointerToValue() {
	obj := &Object{5, 6}
	// objはポインタ型
	fmt.Println(obj)

	fmt.Printf("before multiple :%v\n", obj)
	// 正しくはこのように(*obj).と書くことになる。ただ、明示的に値型で呼び出してあげなくても大丈夫。
	(*obj).multiple(2)
	fmt.Printf("after multiple :%v\n", obj)

	fmt.Printf("before multiple :%v\n", obj)
	// ポインタ型->値型の変換はコンパイラが暗黙的にやってくれる
	obj.multiple(2)
	fmt.Printf("after multiple :%v\n", obj)
}
