# Go

- [types]()
- [array]()
- [fmt]()
- [strings]()

---
## types
In go we have sign and unsign int, float
```go
    
    var age1 int8
    var numb uint8

    var num2 int16
	var age2 uint16
	
    var numb3 int32
    var numb4 uint32

	var num float32
	var num float64

```

---
## array

List with fix length
```go
    var ages [3]uint8 =[3] uint8 {2,2,2}

	// length is fix
	names :=[4]string{"root","mario"}
	names[1]="foo"
	fmt.Println(ages,len(ages))
```

#### slice
slice can be change
```go
    // define list
	var score=[]int {100,12}
	// change value
	score[1]=32
	// add to list
	score=append(score,23)
```

example
```go
//slice range
	name:=[]int{1,2,3,4,5,6}
	rangeOne:=name[1:4]
	rangeTwo:=name[2:]
	rangeThre:=name[:2]
```
---
## fmt


```go
	age := 12

    numFloat:=222.22

	fmt.Print("Hello \n")

	fmt.Println("Hi",age,"end")

	fmt.Printf("age %d",age)

	fmt.Printf("age %f",numFloat)

```

---
## strings

```go
import (
	"fmt"
	"strings"
)

    greeting :="Hello world!"

	// return boolean
	contain:=strings.Contains(greeting,"Hello")
	fmt.Println(contain)

	replace:=strings.ReplaceAll(greeting,"Hello","hi")

	toUpper:=strings.ToUpper(greeting)

	// return index word in string
	index:=strings.Index(greeting,"w")

	split:=strings.Split(greeting," ")

	fmt.Println(split)
```