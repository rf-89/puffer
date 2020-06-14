package main

import (
	"flag"
	"log"
	"path/filepath"
	"sort"

	"github.com/teapod89/puffer/evaluate"
	"github.com/teapod89/puffer/fileinfo"
	"github.com/teapod89/puffer/hash"
	"github.com/teapod89/puffer/report"
	"github.com/teapod89/puffer/util/file"
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

	files, err := file.Glob(*in)

	if err != nil {
		log.Println("failed to get file error.")
	}

	fileInfos := hash.Calculate(files, *out, *num)

	dirfiles := fileinfo.GetDirFiles(fileInfos)

	sort.Slice(dirfiles, func(i, j int) bool {
		return fileInfos[i]["directory"] < dirfiles[j]["directory"]
	})

	dMaps, fnCount, dFnCount := evaluate.Duplicates(dirfiles)

	report.Out(*out, dMaps, fnCount, dFnCount)
}
