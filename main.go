package main

import (
	"errors"
	"fmt"
)

/* ------------------------------------------------------------------------*/
/* opCodes                                                                 */
/* ------------------------------------------------------------------------*/

const (
	PUSH = iota
	ADD
	PRINT
	HALT
	JMPLT
	SUB
)

type opCode struct {
	name  string
	nargs int
}

var opCodes = map[int]opCode{
	PUSH:  opCode{"push", 1},
	ADD:   opCode{"add", 0},
	PRINT: opCode{"print", 0},
	HALT:  opCode{"halt", 0},
	JMPLT: opCode{"jmplt", 2},
	SUB:   opCode{"sub", 0},
}

/* ------------------------------------------------------------------------*/
/* Stack Implementation                                                    */
/* ------------------------------------------------------------------------*/

type Stack struct {
	stack []int
}

func newStack() Stack {
	return Stack{
		stack: []int{},
	}
}

func (s Stack) getLength() int {
	return len(s.stack)
}

func (s *Stack) push(element int) {
	s.stack = append(s.stack, element)
}

func (s *Stack) pop() (element int, err error) {
	if (*s).getLength() > 0 {
		element = (*s).stack[s.getLength()-1]
		s.stack = s.stack[:s.getLength()-1]
		return element, nil
	} else {
		return -1, errors.New("Pop() on empty evaluationStack!")
	}
}

func (s *Stack) peek() (element int, err error) {
	if (*s).getLength() > 0 {
		element = (*s).stack[s.getLength()-1]
		return element, nil
	} else {
		return -1, errors.New("Peek() on empty evaluationStack!")
	}
}

/* ------------------------------------------------------------------------*/
/* VM                                                                      */
/* ------------------------------------------------------------------------*/

type VM struct {
	code            []int
	pc              int // Program counter
	evaluationStack Stack
}

func newVM() VM {
	return VM{
		code:            []int{},
		pc:              0,
		evaluationStack: newStack(),
	}
}

// Private function, that can be activated by Exec call, useful for debugging
func (vm *VM) trace() {
	addr := vm.pc
	opCode := opCodes[vm.code[vm.pc]]
	args := vm.code[vm.pc+1 : vm.pc+opCode.nargs+1]
	stack := vm.evaluationStack
	fmt.Printf("%04d: %s %vm \t%vm\n", addr, opCode.name, args, stack)
}

func (vm *VM) Exec(c []int, trace bool) {

	vm.code = c

	// Infinite Loop until break called
	for {
		if trace {
			vm.trace()
		}

		// Fetch
		opCode := vm.code[vm.pc]
		vm.pc++

		// Decode
		switch opCode {
		case PUSH:
			val := vm.code[vm.pc]
			vm.pc++
			vm.evaluationStack.push(val)

		case ADD:
			right, _ := vm.evaluationStack.pop()
			left, _ := vm.evaluationStack.pop()
			vm.evaluationStack.push(left + right)

		case PRINT:
			val, _ := vm.evaluationStack.peek()
			fmt.Println(val)

		case SUB:
			right, _ := vm.evaluationStack.pop()
			left, _ := vm.evaluationStack.pop()
			vm.evaluationStack.push(left - right)

		case HALT:
			return
		}
	}
}

/* ------------------------------------------------------------------------*/
/* Execution Engine/Main                                                   */
/* ------------------------------------------------------------------------*/

func main() {

	code := []int{
		PUSH, 2, //0, 2
		PUSH, 3, //0, 3
		ADD, //1
		PRINT,
		PUSH, 2,
		SUB,
		PRINT, //2
		HALT,  //3
	}

	vm := newVM()
	vm.Exec(code, true)
}
