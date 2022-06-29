package main

import (
	"context"
	"fmt"
)

type test01String string
type test02String string

func main() {
	// context实例
	ctx := context.Background()
	test01(ctx)
}

func test01(ctx context.Context) {
	// context实例
	ctx01 := context.WithValue(ctx, test01String("test01"), "hello1")
	test02(ctx01)
}

func test02(ctx context.Context) {
	// context实例
	ctx02 := context.WithValue(ctx, test02String("test02"), "hello2")
	test03(ctx02)
}

func test03(ctx context.Context) {
	fmt.Println(ctx.Value(test01String("test01")))
	fmt.Println(ctx.Value(test02String("test02")))
}
