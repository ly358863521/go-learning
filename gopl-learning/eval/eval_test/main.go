package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"eval"
)

func main() {

	env := eval.Env{
		"pi": math.Pi,
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		expr, err := eval.Parse(input.Text())
		if err != nil {
			fmt.Println(err) // parse error
			continue
		}
		fmt.Printf("%.6g\n", expr.Eval(env))
	}

}
