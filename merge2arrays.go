// Eduard Mazo - merge 2 sorted arrays
//
package main

import "fmt"
func merge(input1,input2 [] int) [] int{
    i1 := 0
    i2 := 0
    var out[] int
    for (i1 < len(input1) && i2 < len(input2)) {
	if (input1[i1]<input2[i2]) {
		out = append(out,input1[i1])
		i1++
	}else {	
		out = append(out,input2[i2])
		i2++	
	}
    }
    //add unproceeded tail to out
    for ( i1 < len(input1)) {
	out = append(out,input1[i1])
	i1++
    }	
    for ( i2 < len(input2)) {
	out = append(out,input2[i2])
	i2++
    }
    return out
}
func main() {
    input1 := []int{1,4,5,9,15,16,19}
    input2 := []int{2,3,7,8,11,12,13,14,20}
    fmt.Println("out:",merge(input1,input2))

 } 
