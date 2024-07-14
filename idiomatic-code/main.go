package main

import (
	"fmt"
	"net"
	"sync"
)

// 1 - constant declarations
const (
	Scalar  = 0.1
	version = 1
)

// 2 - variable grouping
func Foo() int {
	var (
		x   = 100
		y   = 2
		foo = "foo"
	)

	fmt.Println(foo)

	return x + y
}

// 3 - functions that panic
func MustParseIntFromString(s string) int { // need add Must in the name
	// logic
	panic("oops")

	return 10
}

// 4- struct initialization
type Vector struct {
	x int
	y int
}

func v() {
	// wrong way
	_ = Vector{20, 10}

	// right way
	_ = Vector{
		x: 20,
		y: 10,
	}
}

// 5 - mutex grouping
type Server struct {
	listenAddr string
	isRunning  bool

	otherMu sync.RWMutex
	other   map[string]net.Conn

	peersMu sync.RWMutex
	peers   map[string]net.Conn
}

// 6 - interface declaration/naming
// type Storer interface { // need to end with -er
// 	Get()
// 	Put()
// 	Delete()
// 	Patch()
// }

// compose interfaces
type Getter interface {
	Get()
}

type Setter interface {
	Set()
}

type Storer interface {
	Getter
	Setter
}

// 7 - function grouping
func VeryImportantFuncExported() {}

func veryImportantFunc() {}

func simpleUtil() {}

// 8 - http handler naming
func handleGetUserById() {}

func handleResizeImage() {}

// 9 - enums
type Suit byte

const (
	SuitSpades Suit = iota
	SuitHearts
	SuitDiamonds
	SuitClubs
)

// 10 - contructor
type Order struct {
	Size float64
}

func NewOrder(size float64) *Order { // or only New if the package is order to be order.New() and not order.NewOrder()
	// Logic here
	return &Order{
		Size: size,
	}
}

func main() {}
