## Purpose of the library
This library solves the problem of converting data from maps to other types in efficient way (see benchmark section).
## How to use it
The main scenario of using this library is when you know exactly to which type you are going to decode your map.
This allows you to prepare you decoder during application initialization and use this efficient decoder during runtime
Library has simple interface to create decoder:
```go
type Person struct {
	Name string `json:"name"`
}
// ...
// prepare decoder for Person type
// NewDecoder is a generic function so decoder is to decode only for type Person
decoder, err := m2o.NewDecoder(Person{})
if err != nil {...}

sourceData := map[string]interface{}{"name": "John"}

var target Person
err = decoder.Decode(sourceData, &target) // Decode(any, *Person) error
```
Built decoder has also method ```decoder.DecodeToAny(any, any) error``` as it's convenient in some scenarios where you build your decoders
for type you don't know:
```go
err = m2o.DecodeToAny(sourceData, &target)
```
also you can use ```decoder.Produce(any) ([T], error)``` which creates object for you:
```go
person, err := decoder.Produce(sourceData)
```
### options for decoder constructor
1. Squash embedded fields:
```go
decoder, err := m2o.NewDecoder(Person{}, m2o.WithSquashAnonymous())
```
2. Make all fields required:
```go
decoder, err := m2o.NewDecoder(Person{}, m2o.WithAllFieldsRequired())
```
3. Set fields to zero value if no present in source data:
```go
decoder, err := m2o.NewDecoder(Person{}, m2o.WithZeroOnEmpty())
```
4. If you have a complex structure of objects linked with each other and want to get every decoded value initialized with all fields (including unexported) from billet:
```go
var billet := newPerson() // person billet has unexported initialized fields
decoder, err := m2o.NewDecoder(Person{}, m2o.WithInitializeFromBillet())
err = decoder.Decode(sourceData, &target) // at this moment target has its unexported fields initialized as in billet
```
5. Use custom decoder for your type:
```go
type Person struct {
    Name string `json:"name"`
    Birthdate time.Time `json:"birthdate"`
}
decoder, err := m2o.NewDecooder(Person{}, m2o.WithTimeDecoder(time.RFC3339))
var person Person
sourceData := map[string]interface{} {
	"name": "John",
	"birthdate": "1985-04-12T23:20:50Z", // will be decoded to person.Birthdate
}
decoder.Decode(sourceData, &person)
```
m2o.WithTimeDecoder from example is just an alias of m2o.WithCustomDecoder(m2o.TimeDecoder(time.RFC3339)), and you can define
a custom decoder for you type with implementing CustomDecoder interface (see example in [time decoder](decoder_time.go))
### 
## How it works
1. it builds execution plan during initialization, so it doesn't need to investigate target object's structure in runtime.
Execution plan is just a tree of wired functions which call each other in correct order, so no unnecessary code is executed in decode time.
2. during running execution plan it mostly relies on memory management without usage of reflect package for most often use cases
decoding source values directly to memory by addresses
## Known limitations:
1. Unlike mapstructure cannot be used for encoding from structure to maps
2. Cannot check if there was unused data in source. This feature is upcoming
## Benchmarks
According to benchmarks in test folder it is twice slower than hand made implementation for specific structure and 15-30 times faster than mapstructure 
