package main

import ( 
	"fmt"
	"./_obj/file"
	"os"
	"strings"
)

func cat(f *file.File) {
        const NBUF = 512
        var buf [NBUF]byte
        for {
		switch nr, er := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.String(), er.String())
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", f.String(), ew.String())
			}
		}
        }
}

func parseline(buf []byte) []string {
	//fmt.Printf("read:\n %v\n", string(buf))
	return strings.Split(string(buf), "\n", -1)
}

func parse(f *file.File) []string{
        const NBUF = 512
        var buf [NBUF]byte
	var bigbuf []byte
	
        for {
		switch nr, er := f.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "error from %s: %s\n", f.String(), er.String())
			os.Exit(1)
		case nr == 0: // EOF		
			return parseline(bigbuf)
		case nr > 0:
			//fmt.Printf("read %v bytes\n", nr)
			bigbuf = append(bigbuf, buf[0:nr]...)
		}
        }

	return parseline(bigbuf)
}

func main() {
	Array1 := [4] int {1,2,3,4}
	Slice1 := Array1[0:2]
	
	fmt.Printf("hello, world %v, %v \n", Array1, Slice1)

	Array1[1] = 155

	fmt.Printf("hello, world %v, %v \n", Array1, Slice1)


	hello := []byte("hello, world file\n")
        file.Stdout.Write(hello)
        f, err := file.Open("/data/cdcdata/ip_area.db",  0,  0)
        if f == nil {
		fmt.Printf("can't open file; err=%s\n",  err.String())
		os.Exit(1)
        }	
	defer f.Close()

	Array2 := parse(f)
/*
	for i,rec := range Array2 {	// Loop over values received from 'src'.
		fmt.Printf("rec[%v] = %v \n", i, rec)		
	}
*/
	fmt.Printf("find %v recs\n", len(Array2))		
	
	var s string = string(hello)
	os.Stdout.WriteString(s)

	input := []byte("iloveyou\nbutilovefreedombetter\n")
	Array3 := parseline(input)
	fmt.Printf("array3 is %v \n", Array3)
}
