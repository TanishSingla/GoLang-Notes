package main

import "fmt"

func main() {
	fmt.Println("Slices are like vectors");
	vector := []int{1,2,3,4,5};
	fmt.Println(vector)
	fmt.Printf("Type of vector : %T \n",vector);
	// vector = append(vector, 1);
	// vector = append(vector, 10);
	// vector = append(vector, 10,12,9,5);
	fmt.Println("Length->",len(vector));
	fmt.Println("Capacity->",cap(vector));
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized
	//Slice is a variable-length sequence that stores elements of a similar type, you 
	// are not allowed to store different type of elements in the same slice. It is just 
	// like an array having an index value and length, but the size of the slice is resized 
	// they are not in fixed-size just like an array. Internally, slice and an array are 
	// connected with each other, a slice is a reference to an underlying array. It is allowed 
	// to store duplicate elements in the slice. 


	//slice using make
	nums := make([]int,3,5);
	// make([]datatype,length,capacity);
	fmt.Println(nums);
	fmt.Println(len(nums));
	fmt.Println(cap(nums));

	// right now the nums array is = [0,0,0]
	//as length is 3
	//if we append any element in nums array
	//then lets say 4 , then nums will look like this [0,0,0,4] (len=4,cap=5)

	nums = append(nums, 4,5,6);
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	// capacity get double every time if exceed the current capacity...

	var names []string;
	fmt.Println(names);

	frnds := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
	};
	fmt.Println(frnds);

	var slice1 = frnds[1:3];
	var slice2 = frnds[3:5];
	fmt.Println(slice1)
	fmt.Println(slice2)

	s := []struct{
		name string
		age int
	}{{"Adam",21},{"Maxim",23}};

	fmt.Println(s);

	/*
		When slicing, you may omit the high or low bounds to use their defaults instead.
		The default is zero for the low bound and the length of the slice for the high bound.
		For the array
		var a [10]int
		these slice expressions are equivalent:
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/

	s2 := []int{1,2,3,4,5}
	fmt.Println(len(s2));
	fmt.Println(cap(s2));
	
	// s2 = s2[:2]
	// fmt.Println(s2);
	// fmt.Println(len(s2));
	// fmt.Println(cap(s2));


	s2 = s2[2:]
	fmt.Println(s2);
	fmt.Println(len(s2));
	fmt.Println(cap(s2));

}