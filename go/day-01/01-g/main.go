// golfed: 276 characters

package main;import(."bytes";."log";."os";."sort";."strconv");func main(){f,_:=
ReadFile(Args[2]);e:=[]int{0};for _,s:=range Split(f,[]byte("\n")){if len(s)==0{
e=append(e,0);continue};x,_:=Atoi(string(s));e[len(e)-1]+=x};Ints(e);l:=len(e)-1;
Printf("%d:%d",e[l],e[l]+e[l-1]+e[l-2])}

// un-golfed: 350

// package main

// import (
// 	. "bytes"
// 	. "log"
// 	. "os"
// 	. "sort"
// 	. "strconv"
// )

// func main() {
// 	f, _ := ReadFile(Args[2])
// 	e := []int{0}
// 	for _, s := range Split(f, []byte("\n")) {
// 		if len(s) == 0 {
// 			e = append(e, 0)
// 			continue
// 		}
// 		x, _ := Atoi(string(s))
// 		e[len(e)-1] += x
// 	}
// 	Ints(e)
// 	l := len(e) - 1
// 	Printf("%d:%d", e[l], e[l]+e[l-1]+e[l-2])
// }
