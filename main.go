package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var (
		in  = flag.String("in", "", "input directory path.")
		out = flag.String("out", "", "output file path.")
		num = flag.Int("num", 1, "maximum parallel number.")
	)

	flag.Parse()

	if *in == "" {
		log.Fatalln("Please input the target directory name.")
	}

	if *out == "" {
		log.Fatalln("Please output the target file path.")
	}

	files, err := getFiles(*in)
	if err != nil {
		log.Println("failed to get file error.")
	}
	fmt.Println(fileProc(files, *out, *num))
}
