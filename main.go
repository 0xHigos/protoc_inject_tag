package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//需要生成的 .proto 文件目录

	var inputFiles, xxxTags string
	var removeTagComment bool
	var omitempty bool

	flag.StringVar(&inputFiles, "input", "", "pattern to match input file(s)")
	flag.StringVar(&xxxTags, "XXX_skip", "", "tags that should be skipped (applies 'tag:\"-\"') for unknown fields (deprecated since protoc-gen-go v1.4.0)")
	//是否去掉从proto生成到.pb.go的注释，默认去掉，保持美观
	flag.BoolVar(&removeTagComment, "remove_tag_comment", false, "removes tag comments from the generated file(s)")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")
	flag.BoolVar(&omitempty, "omitempty", false, "whether remove 'omitempty' in generated .pb json")

	flag.Parse()

	var xxxSkipSlice []string
	if len(xxxTags) > 0 {
		logf("warn: deprecated flag '-XXX_skip' used")
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}

	if inputFiles == "" {
		log.Fatal("input file is mandatory, see: -help")
	}

	// Note: glob doesn't handle ** (treats as just one *). This will return
	// files and folders, so we'll have to filter them out.
	globResults, err := filepath.Glob(inputFiles)
	if err != nil {
		log.Fatal(err)
	}

	var matched int
	for _, path := range globResults {
		finfo, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		if finfo.IsDir() {
			continue
		}

		// It should end with ".go" at a minimum.
		if !strings.HasSuffix(strings.ToLower(finfo.Name()), ".go") {
			continue
		}

		matched++

		areas, err := parseFile(path, nil, xxxSkipSlice)
		if err != nil {
			log.Fatal(err)
		}
		if err = writeFile(path, areas, removeTagComment, omitempty); err != nil {
			log.Fatal(err)
		}
	}

	if matched == 0 {
		log.Fatalf("input %q matched no files, see: -help", inputFiles)
	}
}
