package main

import "fmt"

// error is standard word
func divide(a,b float64) (float64,error) {
	if b==0{
		return 0,fmt.Errorf("denominator shouldn't be zero");
	}
	return a/b,nil;
}

func main() {
	fmt.Println("boom");
	ans,e := divide(10,0);
	fmt.Println(ans)
	fmt.Println(e)

	ans2,_ := divide(10,2);
	fmt.Println(ans2);

	// if we want to discard any output given by any function we use 
	// ,underscore
}