* Structs are a data structure that we use to model our data in GO.
* Struct is like objects in JS or hashes in Ruby or dictionaries in Python.
* struct fields gets assigned what is known as zero values for each field datatype
```go

type person struct {
	firstName string
	age int // or float
  married bool

}
	var julie person

	fmt.Printf("%+v", julie)

  /** OUTPUT
  
  {firstName: "", age: 0, married: false}
  **/
```
