# Golang

[TOC]

1. Hello & packages
2. Func, loop, slice maps
3. Structs, Interface & structs embedding
4. Errors
5. Goroutines

====

## 1. Hello & packages

```go
package main

import (
	"fmt"
    "math"
	"math/rand"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Hello, 世界")
    
    fmt.Println(math.pi)
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    
    //swap
    a, b := swap("hello", "world")
	fmt.Println(a, b)
}

```

## 2. Func, loop, slice maps

```go
func LoopTest(){
    // loop over an array/a slice
    for i, e := range a {
     // i is the index, e the element
    }
    // if you only need e:
    for _, e := range a {
     // e is the element
    }
    // ...and if you only need the index
    for i := range a {
    }
    // In Go pre-1.4, it is a compiler error to not use i and e.
    // Go 1.4 introduced a variable-free form:
    for range time.Tick(time.Second) {
     // do it once a sec
    }
}

func Test(){
    var m map[string]int
    m = make(map[string]int)
    m["key"] = 42
    fmt.Println(m["key"])
    delete(m, "key")
    elem, ok := m["key"] // test if key "key" is present, retrieve if so
    // map literal
    var m = map[string]Vertex{
     "Bell Labs": {40.68433, -74.39967},
     "Google": {37.42202, -122.08408},
    }
}

func main() {
	Test()
}
```



## 3. Structs, Interface & structs embedding

```go
There are no classes, only structs. Structs can have methods.
// A struct is a type. It's also a collection of fields
// Declaration
type Vertex struct {
 X, Y int
}
// Creating
var v = Vertex{1, 2}
var v = Vertex{X: 1, Y: 2} // Creates a struct by defining values
 // with keys
// Accessing members
v.X = 4

// interface declaration
type Awesomizer interface {
 Awesomize() string
}
// types do *not* declare to implement interfaces
type Foo struct {}
// instead, types implicitly satisfy an interface if they implement
all required methods
func (foo Foo) Awesomize() string {
 return "Awesome!"
}
```

## 4. Errors

```go
type error interface {
 Error() string
}

func doStuff() (int, error) {
}
func main() {
 result, error := doStuff()
 if (error != nil) {
 // handle error
 } else {
 // all is good, use result
 }
}
```



## 5. Goroutines

Goroutines are lightweight threads (managed by Go, not OS threads). go f(a, b) starts a new goroutine which runs f (given f is a function).

```go
// just a function (which can be later started as a goroutine)
func doStuff(s string) {
}
func main() {
    // using a named function in a goroutine
    go doStuff("foobar")
    // using an anonymous inner function in a goroutine
    go func (x int) {
        // function body goes here
    }(42)
}
```



## Next

Next list / cheats

1. Goroutine cases
2. var types
3. Compiling
4. Channel
5. Boilerplate
6. Http
7. File read & write
8. Defer, panic, recover

## refs

- http://www.cheat-sheets.org/saved-copy/go-lang-cheat-sheet-master.20181212/golang_refcard.pdf
- https://tour.golang.org/
- https://golangdocs.com/panic-defer-recover-in-golang
- https://medium.com/benchkram/less-verbose-error-handling-in-golang-10e2679707a2
