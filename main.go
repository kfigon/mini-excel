package main

import (
	"fmt"
	"os"

	"github.com/sqs/goreturns/returns"
)

func main() {
	fmt.Println("asd")
	if len(os.Args) < 2 {
		fmt.Println("no file provided")
		return
	}
	path := os.Args[1]
	fmt.Println(path)

	d, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	ss := spreadsheetParser{"\n", ","}.parse(string(d))
	for header,rows := range ss {
		for k,v := range rows {
			exp, ok := v.(expressionCell)
			if !ok {
				continue
			}
			v, err := newEvaluator(ss).eval(fmt.Sprintf("%v%v", header,k), exp)
			if err != nil {
				fmt.Println(err)
				return
			}
			ss[header][k] = numberCell(v)
		}
	}

	fmt.Println(ss)
}