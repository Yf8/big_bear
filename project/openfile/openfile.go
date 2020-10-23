package main

import ("fmt"
 		"os"
 		"bufio")


func writeFile(filename string)  {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()


	for i := 0; i < 20; i++{
		fmt.Fprintln(writer, i)
	}
}

func main()  {
	writeFile("fib.txt")
}