package main

func main2() {
	buffer := make([]string, 4, 5)
	buffer[0] = "a"
	buffer[1] = "b"
	buffer[2] = "c"
	println(len(buffer), cap(buffer))

	buffer = append(buffer, "d")
	println(len(buffer), cap(buffer))

	arr := make([]int, 5)
	arr[0] = 11
	arr[1] = 22
	arr[2] = 33
	arr[3] = 44
	arr[4] = 55
	for idx, val := range arr[0:0] {
		println("---", idx, val)
	}

}
