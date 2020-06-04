package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/teapod89/puffer/evaluate"
	"github.com/teapod89/puffer/files"
	"github.com/teapod89/puffer/hash"
	"github.com/teapod89/puffer/report"
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

	if filepath.Ext(*out) != ".xlsx" {
		log.Fatalln("Please output the target file path.")
	}

	files, err := files.Glob(*in)

	if err != nil {
		log.Println("failed to get file error.")
	}
	var fileInfos = hash.Calculate(files, *out, *num)
	var duplicates = evaluate.Duplicates(fileInfos)
	report.Out(*out, duplicates)

}
