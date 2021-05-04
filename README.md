# Basic-Golang
These are the building blocks of writing a simple program in golang. This is the combination of multiple research and tutorials online to create a library of useful tools.
***

Simple helloworld program.

```
package main                       #

import "fmt"

func main(){

fmt.Println("Hi There!")
}
```

There are 5 basic questions:
-  How do we run the project - go run main.go

- Go run compiles and immediately executes but go build only compiles
- Go fmt formats all the code
- Go get is like composer install

GO PACKAGES
- This is a collection of files that work towards the same goal
- All the files that are part of a package must have package declaration at the top


TYPES OF PACKAGES
- EXECUTABLE PACKAGES
	- Generate a file that we can run
- REUSABLE PACKAGES
	- Code used as helpers. 
	- Good  place to put reusable logic
- The word main is used to create an executable package
- Any other package name creates a reusable package.

PACKAGES (GO)  -  NAMESPACE (PHP)


IMPORT STATEMENTS
- Give access too code which is written within another package
- We could import from other engineers which are 3rd party
- Fmt is a shortened form of format included in Golang


- We can find package documentation on standard libraries at golang.org/pkg.

FILE ORGANIZATION
- Func stands for function -

- We have package declaration
- We then have package imports 
- We the do the function declaration.

VARIABLE DECLARATIONS - TYPED LANGUAGE

Some of the basic data types in golang are
- Bool - 
- string
- int - 0, -10000, 99999 
- Float64 - 10.0000001, 0.00009, -100.003
```
card := "Ace of spades"  // trusts the go compiler to determine the type. This syntax is only used for initialization
```


FUNCTIONS AND RETURN TYPES
- The function should be explicit and define the return type.
```
func newCard() string {
    return "Five of diamonds"
}
```



SLICES AND FOR LOOPS
- ARRAY - This is a fixed length list of things
- SLICE - An array that can grow or shrink
	- Every record must be of the same type

```
func main() {
    card := []string{newCard(), "Ace of diamonds"}
    card = append(card, "Six of spades")
    for i, card := range card {
        fmt.Println(i, card)
    }
}
```




```
go run main.go deck.go   // running related files
```

RECEIVER FUNCTIONS
```
Func (d deck) print(){ … }     # This means any variable of type deck gets access to this function
```


    Whenever we have variables that we will not use we replace them with underscore



SLICE RANGE SYNTAX
- Slices are zero-indexed
- Individual items are accessed fruit[0]
- Slices are like python slices fruit[0:2] - starting and ending index


MULTIPLE RETURNS
- You can return multiple values from a function
```
func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}
```
- We can assign multiple values also

BYTE SLICES
- Every character corresponds to an ascii character code in the slice

TYPE CONVERSION
- This is changing from one type to the other. Type casting in GO 
```
[]byte(“hi there”)    => turns string into byte slice
```

JOINING A SLICE OF STRINGS
- 

SAVING DATA TO DISK
- ioutil package is used.
- Writefile , ReadFile

ERROR OBJECTS

- Nil is no value in go

RANDOM NUMBER GENERATION
```
source := rand.NewSource(time.Now().UnixNano())

r := rand.New(source)

for i := range d {

np := r.Intn(len(d) - 1)

d[i], d[np] = d[np], d[i]
```

TESTIN WITH GO
- To create a test we create a file with _test
- To run all tests we run go test

ORGANISING DATA
- The equivalent of structs or classes in go
- With structs, go assigns a zero value to variables that are not assigned on creation


- We can assign values to a struct



EMBEDDED STRUCTS


Initialising embedded structs.
- Every single line in an embedded struct must have a comma


Go is a pass by value language
- 

POINTERS IN GO - PASS BY REFERENCE

- The same style as C++


- & -> references to the address in memory where a variable is pointing
- * -> take the memory address and give me value that exists there.




- When dealing with a slice, values are passed by reference


- When we create a slice we have a data structure as well as the underlaying array

MAPS IN GO
- This is a collection of key value pairs
- Both keys and values are statically typed





MANIPULATING MAPS
- A way of declaring maps are as follows
```
var colors map[string]string
—-
colors := make(map[string]string)

```



- We can only use bracket syntax for assignment as opposed to . Format.
```
colors := make(map[string]string)

colors["white"] = "#fff"

delete(colors, "white”)                        # Delete a value from a map

fmt.Println(colors)
```

ITERATING OVER A MAP




The difference between Maps and structs


- We cannot iterate over a struct however we can do so on a map which is indexed.

INTERFACES
- We can remove receiver value if not used 
```
func (spanishBot) getGreeting() string {
    return "Hola amigo."
}
```






INTERFACE SIGNATURE





We can also define multiple functions to satisfy an interface


We have 2 kinds of types 
- Concrete types - types that you can create a value directly with
- Interface types - cannot create a value directly from these

Interfaces are implicit, you don’t manually match an interface to a concretion.
- Bit hard to keep track of
Interfaces can be combined to create other more complex interfaces



CONCURRENT PROGRAMS
- This is serial checking
- This results in blocking calls that halt progress until items are done in the sequence


We can use a parallel approach.
- We use go routines to do this


- When you run a program you have activated a go-routine
- We use the go keyword to create a routine.
```
go checkLink(link)         # on function call
```
- The go scheduler is used to assign new processes for concurrent tasks
- Go attempts to use only 1 cpu core
- CONCURRENCY IS NOT PARALLELISM




CHANNELS
- Makes sure that the main routine tracks progress for sub-routines
- These are types like any other variable
```
c := make(chan string)

func checkLink(link string, c chan string) { … }

```



BLOCKING CHANNELS
- Receiving messages from a channel is a blocking call



We can use c++ inspired for loops
```
for i := 0; i < len(links); i++ {
    fmt.Println(<-c)
}
```

REPEATING ROUTINES
- We can continually run routines (Status checking behavior)
```
for l := range c{
    go checkLink(l, c)         // Where c is a channel
}
```

SLEEPING A ROUTINE
- We use the sleep function to pause the go routine



FUNCTION LITERALS
- These are anonymous functions or from python these are lambda functons
- Same as javascript we can define and invoke functions















