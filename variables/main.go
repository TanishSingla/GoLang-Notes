package main

func main() {

	//variables :-
	var var1 bool = true
	var1 = false
	println(var1)

	var var2 = "string "
	println(var2)

	// van assign and declare variables like this as well

	age := 21
	println(age)

	// Note :- if we start any variable name with small letter then that variable will be private
	// but if we are starting it with capital letter than that variable is public
	// and can be accessed anywhere , same is applicable to functions as well.
	
	PublicVariable := "Public Variable Can be accessed in others packages as well"
	privateVariable := "Private Variable"
	println(PublicVariable)
	println(privateVariable)
}