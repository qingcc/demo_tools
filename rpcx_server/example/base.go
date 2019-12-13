package example

import (
	"context"
	"fmt"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}
type Reply struct {
	C int `msg:"c"`
}
type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}

type Arith1 int

func (t *Arith1) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B * 100
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}
func (t *Arith) Error(ctx context.Context, args *Args, reply *Reply) error {
	panic("ERROR")
}
