# slice
go slice tools

# sliceutil

golang sliceutil, including methods like TakeWhile, Filter, ForEach and so on, which were inpired by js slice

# examples

```go
package main

import (
	"github.com/conbanwa/slice"
)

func main() {
	println(slice.SubString(`1234567890`, 2))       //34567890
	println(slice.SubString(`1234567890`, -2))      //90
	println(slice.SubString(`1234567890`, 0, -4))   //123456
	println(slice.SubString(`1234567890`, 2, -4))   //3456
	println(slice.SubString(`1234567890`, 6, 99))   //7890
	println(slice.SubString(`1234567890`, -5, 4))   //empty
	println(slice.SubString(`1234567890`, -5, 99))  //67890
	println(slice.SubString(`1234567890`, -99, 99)) //1234567890

	println(slice.Slice([]byte(`1234567890`), 2))       //34567890
	println(slice.Slice([]byte(`1234567890`), -2))      //90
	println(slice.Slice([]byte(`1234567890`), 0, -4))   //123456
	println(slice.Slice([]byte(`1234567890`), 2, -4))   //3456
	println(slice.Slice([]byte(`1234567890`), 6, 99))   //7890
	println(slice.Slice([]byte(`1234567890`), -5, 4))   //empty
	println(slice.Slice([]byte(`1234567890`), -5, 99))  //67890
	println(slice.Slice([]byte(`1234567890`), -99, 99)) //1234567890
}
```

# install

import github.com/conbanwa/slice
