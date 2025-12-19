package main

func main() {
	//Write a reverse of a number
	number := 18484
	for number != 0 {
		digit := number % 10
		number /= 10
		print(digit)
	}

}
