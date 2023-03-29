package main

import (
	"fmt"
)

func Permutations(L []int, r int) [][]int {
	if r == 1 {
		//Convert every item in L to List and
		//Append it to List of List
		temp := make([][]int, 0)
		for _, rr := range L {
			t := make([]int, 0)
			t = append(t, rr)
			temp = append(temp, [][]int{t}...)
		}
		return temp
	} else {
		res := make([][]int, 0)
		for i := 0; i < len(L); i++ {
			//Create List Without L[i] element
			perms := make([]int, 0)
			perms = append(perms, L[:i]...)
			perms = append(perms, L[i+1:]...)
			//Call recursively to Permutations
			for _, x := range Permutations(perms, r-1) {
				t := append(x, L[i])
				res = append(res, [][]int{t}...)
			}
		}
		return res
	}
}
func main() {
	res := Permutations([]int{1, 2, 3, 4}, 2)
	fmt.Println(res, len(res))

}
