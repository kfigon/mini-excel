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

// shunting yard algorithm
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
			
			lVal, err := e.extract(lhs)
			if err != nil {
				return 0, err
			}
			rVal, err := e.extract(rhs)
			if err != nil {
				return 0, err
			}
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
		return e.extract(v)
	}

	return -1, fmt.Errorf("unknown error")
}

func (e *evaluator) extract(t token) (int, error) {
	if t.tokType == number {
		v, _ := strconv.Atoi(t.val)
		return v, nil
	} else if t.tokType == coord {
		c, ok := e.r.read(t.val)
		if !ok{
			return 0, fmt.Errorf("can't extract value from coordinate: %v", t.val)
		}

		switch cell := c.(type) {
		case numberCell: return int(cell), nil
		default:
			return 0, fmt.Errorf("can't read value from %v - unknown cell type", t)
		}
	}

	return 0, fmt.Errorf("can't extract value from token %v", t)
}