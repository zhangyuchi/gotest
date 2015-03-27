package main

import ( 
	"fmt"
	"io"
	"bytes"
	"os"
)

//S's object must be pointer
type S struct { i int }
func (p *S) Get() int { return p.i }
func (p *S) Put(v int) int { tmp:=p.i; p.i=v; return tmp }

//R's object must be value
type R struct { i int }
func (p R) Get() int { return p.i }
func (p R) Put(v int)  int { tmp:=p.i; p.i=v; return tmp }

type I interface {
	Get() int
	Put(int) int
}

func f(p I) {
	fmt.Println(p.Put(p.Get()+1))
}

/*func e(p *I) {
	fmt.Println(p.Put(p.Get()+1))
}*/

func sort(i []interface{}) {
	switch i.(type) {
	case string:
		// ...
	case int:
		// ...
	}
	return /* ... */
}


func main() {
	//Array1 := [4] int {1,2,3,4}
	//Slice1 := Array1[0:2]
	
	var s S; 
        var r R;
 	for i:=0;i<10;i++ {
		f(r)
		f(&s)
	}

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)
	//fmt.Println(buf)
}
