package main

import (
	"fmt"
	"strconv"
)

type reader interface {
	read(string) (cell, bool)
}

type evaluator struct {
	r reader
}

func newEvaluator(r reader) *evaluator {
	return &evaluator{r: r}
}

// todo: shunting yard algorithm or recursive descent parser?
func (e *evaluator) eval(currentCoord string, exp expressionCell) (int, error) {
	tokens := convertInfix(parseExpression(exp))
	if len(tokens) < 2 {
		return -1, fmt.Errorf("invalid expression provided")
	}

	iter := &tokenIterator{tokens: tokens}
	iter.next() // =

	operandStack := newStack[token]()

	for iter.hasNext() {
		cur, _ := iter.currentToken()
		
		if cur.tokType.isOperand() {
			operandStack.push(cur)
		} else {
			rhs, ok := operandStack.pop()
			lhs, ok := operandStack.pop()

			if !ok {
				return -1, fmt.Errorf("too few operands")
			}
			lVal, _ := strconv.Atoi(lhs.val)
			rVal, _ := strconv.Atoi(rhs.val)
			newToken := token{tokType: number}

			switch cur.tokType {
			case plus:
				newToken.val = strconv.Itoa(lVal+rVal)
			case minus:
				newToken.val = strconv.Itoa(lVal-rVal)
			case multiply:
				newToken.val = strconv.Itoa(lVal*rVal)
			default:
				return -1, fmt.Errorf("unknown operator: %v", cur)
			}
			operandStack.push(newToken)
		}

		iter.next()
	}
	
	if len(operandStack.tab) == 1 {
		v, _ := operandStack.pop()
		res, _ := strconv.Atoi(v.val)
		return res, nil
	}

	return -1, fmt.Errorf("unknown error")
}