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


REFRESHER NOTES
- Naming variable uppercase first means they will be visible outside the package and lowercase means they will not be visible outside the package.
- Func init is a constructor in golang

TEMPLATES
- We can pass data into a template using the following notation
```
<h1>{{.}}</h1>
```
- We can also use $ notation to access a specific value



PARSING COMPOSITE DATA STRUCTURES
- We can use range in the template to iterate over the data


- We can also use variable for utterable types



- We can also access individual elements in a struct using the . notation
- Please remember that if members of a struct are lower case they will not be exported




Initialising the code that will populate multiple templates will be as follows
```
func init() {

    tpl = template.Must(template.ParseFiles("tpl.gohtml", "itr.gohtml", "struct.gohtml”))        // Constructor function that can be used to stage all template files for rendering   

}


func main() {

    gd := person{"Godfrey", "Bafana"}

    err := tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", gd)              // Runs the template to a writer of some sort. This example uses standard out

    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
}

```


In go we can create anonymous types also


- The above defines a type and initializes it at the same time

PASSING FUNCTIONS INTO TEMPLATES
- You can pass in your own functions or Global functions that are predefined.
- We can initialize the templates using the below
```
func init() {

    tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))

}

```



- We can now use the function directly in the html template



To ensure that we can hold anything in a variable we use interface{ }
- Funcs need to be parsed before template is passed
- Parse files or parse glob has to come after funcs are attached
- Go also supports chaining 
- 

PIPELINES IN TEMPLATES
- Go also uses pipes to format any data passed to its templates
```
{{ . | sqrt | cube }}
```


NESTING TEMPLATES
- You can define as many templates in 1 file as you want

- Templates can be directly referenced from other templates


You can pass the current piece of data to the template by using the following notation
```
<div>
    {{ template "polarbear" .}}
</div>
```
    

TEMPLATE COMPOSITIONS



HTML TEMPLATE PACKAGE
- Is built on top of text template
- Can escape unsafe characters on the web
- Ensures that we avoid xss


SERVERS IN GOLANG
- HTTP sits on top of TCP 
- You can set a deadline to a connection
```
Err := conn.SetDeadLine(time.Now().Add(10 * time.Second))
```


HTTP PACKAGES
- We have 2 methods that can be used:
	- http.ListenAndServe(adds, handler)
	- Http.ListenAndServeTLS(adds, cert, key, handler)
- Handling form submission on httpServe





- The response writer gives access to header updates



UNDERSTANDING SERVEMUX - ROUTING - MULTI-PLEXING
- Without having elegant routing we would have to use the following.


Implementation with a server mix


- We can decide to use default handlers for the servemux
```
func main() {
    var hd hotdog
    var ca bun
    http.Handle("/", hd)
    http.Handle("/animals/", ca)
    http.ListenAndServe(":8080", nil)
}
```

- We can also use the handle func method so we don’t create additional resources.


WHERE B AND C FOLOW:
```
func b(w http.ResponseWriter, req *http.Request)
func c(w http.ResponseWriter, r *http.Request)
```


SERVING FILES IN GOLANG
- We use io.Copy
- We can also use httpFileServer

This serves everything in this folder.

USING STRIP-PREFIX
- This takes a path and makes everything that comes beyond the path a file. Traverses directories to get to files.




- For static file hosting you need to have an index.html file available at root.


LOG FATAL vs HTTP ERROR
- We can wrap http.ListenAndServe in log.Fatal so that the function can automatically catch errors and log them
```
Log.Fatal(http.ListenAndServe(“:8080”, nil))
```
- Http error is tailor made for http errors
```
http.Error(responseWriter, “Status not found”, 404)
```

HTTP NOT FOUND HANDLER
- This is a handler that can used to server 404

STATE MANAGEMENT
- We can use query strings and target them in the code using request.FormValue(“param") 
- We can pass values from forms also

ENCTYPE
- This is encoding type
- There are 3 types which are
	- Application/x-www-form-urlencoded
	- Multipart/form-data
	- Text/plain

REDIRECTS
- Some of the codes are
	- 303 - Always changes to GET
	- 301 - Moved permanently. Browser remembers and avoids multiple round trips
	- 307 - Temporary redirect.
func foo(w http.ResponseWriter, req *http.Request) {
  dt := req.FormValue("q")
  io.WriteString(w, "Do my search: "+dt)
}


We can also use an inbuilt http function for redirect
func rdt(w http.ResponseWriter, req *http.Request) {
  fmt.Println("You current method at redirect is: ", req.Method)
  http.Redirect(w, req, "/files", http.StatusMovedPermanently)
}


COOKIES 
- A cookie is a little file that can hold information that a server can write to the clients machine if the client allows.
- Allows us to uniquely id the user.
- Cookies are domain based.
- We always use HTTPS


func set(w http.ResponseWriter, req *http.Request) {
 http.SetCookie(w, &http.Cookie{
  Name: "Godhiza",
  Value: "Mbofana",
 })
 io.WriteString(w, "The cookie was written")
}

func read(w http.ResponseWriter, req *http.Request) {
 c, err := req.Cookie("Godhiza")
 if err != nil {
  http.Error(w, err.Error(), http.StatusNotFound)
 }
 fmt.Fprintln(w, "The cookie was found: ", c)
}


- To delete a cookie we set max age.

SESSIONS
 UUID
- We use this to create a unique id for sessions
func foo(w http.ResponseWriter, req *http.Request) {
 c, err := req.Cookie("session”) 
 if err != nil {
  id := uuid.NewString()
  c = &http.Cookie{
   Name: "session",
   Value: id,
  }
 http.SetCookie(w, c)
 }

 fmt.Println(c)
}


- Maps give zero value when requested key is not available.


USING BCRYPT
- This is used to hash passwords.
pwd, _ := bcrypt.GenerateFromPassword([]byte("Nyandoro23"), bcrypt.MinCost)
fmt.Println(string(pwd))


WEB DEV TOOLKIT

HMAC
- Harsh based message authentication code
h := hmac.New(sha256.New, []byte("harsh_key"))
io.WriteString(h, s)
return fmt.Sprintf("%x", h.Sum(nil))

BASE64 ENCODING
- Ensures consistent character set for data for secure storage
chr := "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890+/"
s64 := base64.NewEncoding(chr).EncodeToString([]byte(s))
return s64

- We can use the standard encorder
func b64Encode(s string) string {
  s64 := base64.StdEncoding.EncodeToString([]byte(s))
  return s64
}

CONTEXT
- Keeps track of all processes that are spawned from a specific request
- You can initialize a context with a value.


TLS & HTTPS
http.HandleFunc("/", ctx)
http.Handle("/favicon.ico", http.NotFoundHandler())
http.ListenAndServeTLS(":8080", "example.crt", "example.key", nil)

JSON WITH GO
- We have Marshal and UnMarshal, Decode and encode
- Inline structs do not export outside the go functions.


UNMASHARL JSON TO GO
data := person{}
ss := string(jsn)
err = json.Unmarshal([]byte(ss), data)


UNMASHARLING WITH TAGS
- These are mappings that tell which json fields belong to which struct members

type cities struct {
  Name string `json name`
  Longitude float64 `json lng`
  Latitude float64 `json lat`
}


- Unmashaling an empty items gives empty values as default.
- 

POSTGRES
- CROSS JOIN
	- Every member of 1 set is matched to every member of the other set
- OUTER JOINERS
	- Gives the full list from left, right or both regardless of relation.
- You can grant specific privileges to a user on Postgres. Does not have to be restricted to all
	- SELECT
	- INSERT
	- UPDATE
	- DELETE
	- RULE
	- ALL

GO AND POSTGRES
- Defer db.Close( ) should not be in init function otherwise it terminates DB connection before transactions. 




```
mongod --dbpath /usr/local/var/mongodb --logpath /usr/local/var/log/mongodb/mongo.log --fork
```
MONGODB - NOSQL DATABASES
- This is a schema less database
- Scales well bu has eventual consistency
There are four types of nodal databases:
- KEY - VALUE - Key value store often running in memory
- DOCUMENT - Stores documents like json or json or xml and stuff.
- GRAPH - For connected data
- COLUMN - For huge data structures


DATABASE
- To create a database you use the werd use
```
Use temp
```
- To insert data into a db you use the following syntax
```
db.<collectionName>.insert({json-data})
db.users.insert({"name":"Godzab"})
```
- To show all records in a collection
```
db.users.find()
```


COLLECTIONS
- You can implicitly create a collection on a database
- This is creating a collection without having defined it.
```
db.surf.insert({"Surfer": true})
```
- You can drop a collection by the following
```
db.surf.drop()
```
- You can also create a collection declaratively
```
db.createCollection(“surf”)

- you can add additional options
1.) capped -> caps the size (Bool)
2.) size -> Size of cap in bytes
3.) max -> maximum number of Documents allowed

db.createCollection(“src”, {capped: true, size: 655512, max: 1000000})
```

DOCUMENT
- This is just a json object
- You can also insert multiple documents at the same time

FIND AND QUERY
```
db.<collection>.findOne()           // Brings back one
db.users.find({name: "Godzab”})     // Find specific item
db.users.find({$and: [{name: "Godzab”}, {age: 32}]})   // uses the and operator
db.users.find({$or: [{name: "Godzab”}, {age: 32}]})   // uses the or operator
————
db.users.find({$and: [{name: "Godzab”}, {age: {$gt : 20}}]})   // uses the greater than clause on age
db.users.find({$and: [{name: "Godzab”}, {age: {$lt : 20}}]})   // uses the less than clause on age
```

- We can also nest ands and ors
- We can also use regex
```
db.users.find({name: {regex:”^G”}})
```

UPDATE AND SAVE
```
db.users.update({"name":"Godzab"}, {$set: {role: "citizen", name:"Godfrey", surname:"Bafana"}})   //updates based on name
db.users.update({_id:objectId(“6072fa5de490e624979884b3")}, {$set: {role: "citizen", name:"Godfrey", surname:"Bafana"}})
```

DELETE
```
db.users.remove({"name":"Godzab”}, 1)     // Deletes 1 record
db.users.remove({"name":"Godzab”})        // deletes all related
db.users.remove({})                       // deletes all
```


PROJECTION
- Selecting some off the data from a query
```
db.<collection>.find(<selection criteria>, <list of fields with toggle 0 or 1>)
db.users.find({name: “Godfrey”}, {_id: 0, name:1, surname:0})
```


LIMIT
```
db.users.find({name: {regex:”^G”}}).limit(3)     // get 3 records
```

SORT
- We use 1 for asc and -1 for desc
```
db.users.find({name: {regex:”^G”}}).limit(3).sort({title: 1})
```

INDEX
- Allows quick search
```
db.users.createIndex({name: 1})
```

AGGREGATION
- We have the count function
- We have distinct function also
```
Db.users.distinct(“dept")
```




























