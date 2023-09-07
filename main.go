package main

import (
	"bufio"
	"fmt"
	"os"

	iconv "github.com/djimenez/iconv-go"
)

const (
	TARGET = "CP949"
)

func main() {
	if 2 > len(os.Args) {
		fmt.Printf("Usage: $ %s [File] \n", os.Args[0])
		return
	}

	for _, from := range FROM {
		convert, err := iconv.NewConverter(from, TARGET)
		if err != nil {
			panic(err)
		}
		defer convert.Close()

		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			output, err := convert.ConvertString(scanner.Text())
			if err != nil {
				if err.Error() != "invalid or incomplete multibyte or wide character" {
					fmt.Printf("%s\t%s\n", err, from)
				}

			} else {
				fmt.Printf("[%s] %s\n", from, output)
			}
		}
	}
}
