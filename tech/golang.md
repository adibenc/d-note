# Golang


[TOC]

## Research
## Basic
### Fundamental

[wiki] Go was designed at [Google](https://en.wikipedia.org/wiki/Google) in 2007 to improve [programming productivity](https://en.wikipedia.org/wiki/Programming_productivity) in an era of [multicore](https://en.wikipedia.org/wiki/Multi-core_processor), [networked](https://en.wikipedia.org/wiki/Computer_network) [machines](https://en.wikipedia.org/wiki/Computer) and large [codebases](https://en.wikipedia.org/wiki/Codebase).[[22\]](https://en.wikipedia.org/wiki/Go_(programming_language)#cite_note-22) The designers wanted to address criticism of other languages in use at [Google](https://en.wikipedia.org/wiki/Google), but keep their useful characteristics:[[23\]](https://en.wikipedia.org/wiki/Go_(programming_language)#cite_note-23)

- [static typing](https://en.wikipedia.org/wiki/Static_typing) and [run-time](https://en.wikipedia.org/wiki/Run_time_(program_lifecycle_phase)) efficiency (like [C](https://en.wikipedia.org/wiki/C_(programming_language))),
- [readability](https://en.wikipedia.org/wiki/Readability) and [usability](https://en.wikipedia.org/wiki/Usability) (like [Python](https://en.wikipedia.org/wiki/Python_(programming_language)) or [JavaScript](https://en.wikipedia.org/wiki/JavaScript)),[[24\]](https://en.wikipedia.org/wiki/Go_(programming_language)#cite_note-24)
- high-performance [networking](https://en.wikipedia.org/wiki/Computer_network) and [multiprocessing](https://en.wikipedia.org/wiki/Multiprocessing).

[doc] https://golang.org/doc/

### Pattern : Concurrency

Goroutines [https://www.youtube.com/watch?v=f6kdp27TYZs]

[https://github.com/adityamenon/Google-IO_2012_Go-Concurrency-Patterns]

1. Independent execution
2. Concurrency goal is not paralellism, but good structure
3. not equal to parallelism, but concurrency enables parallelism
4. it is not a thread, instead goroutines are multiplexed dynamicaly onto threads as needed to keep all the goroutines running

Channels : a channel in go provides a connection between two goroutines, allowing them to communicate.

Generator : function that returns a channel. a "Pattern".

Multiplexing : **fanIn func**. makes channels communicating or run independently

select : switch-like to control channel communication. except it blocks until a communication can proceed. Just like await in javascript but with condition of channel must be proccess-able (the **one channel** that is ready will be executed first). this gives you way to non-blocking communication. **example** : fanIn func in select statement.

No locks, no condition variables, no callbacks.

### Why

1. System software

2. Concurrency

   1. not equal to parallelism, but concurrency enables parallelism
   2. it is not a thread,

3. In just few simple transformations, go have concurrency primitives that able to convert a

   - slow, sequential, failure sensitive

     to program that is

   - fast, concurrent, replicated, robust

4. aa

### Why not

1. Government project ?
2. I just know PHP to make website ?

### Libraries

1. basic utilities
   1. cobra : go cli tools
2. web
   1. viper
   2. routing : chi
   3. jwt-go
   4. ozzo-validation
3. db : 
   1. go-pg
   2. go-mysql
4. tracker
   1. Logrus
5. mail
   1. gomail
   2. 

### 10 common cases

1. a

### 30 Common operations

1. IO
   1. cli stdin / stdout
   2. http request
2. channeling
3. goroutine

### 10 products implementation

1. https://en.wikipedia.org/wiki/Go_(programming_language)#Applications

### Official Doc & 3 refs
## Production
1. Basic api project skeleton & boilerplate

2. Basic api & frontend pluggable project skeleton & boilerplate

### Deploying
## FE
### Web HTML
### Desktop
### CLI
## Common cases
1. object to JSON, JSON to object

   ```go
   
   ```

2. file system operations
   ```go
   
   ```

3. db operations
   ```go
   
   ```
   



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
