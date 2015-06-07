package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jasonmoo/wc"
)

var (
	multibytes = flag.Bool("m", false, "count the multibyte runes")
	lines      = flag.Bool("l", false, "count the lines")
	words      = flag.Bool("w", false, "count the words")
	bytes      = flag.Bool("c", false, "count the bytes")
)

func main() {

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("wc [-l] [-m] [-w] [-b] file [...fileN]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// default to stdin
	if flag.NArg() == 0 {
		c := wc.NewCounter(os.Stdin)

		err := c.Count(*multibytes, *bytes, *lines, *words)
		if err != nil {
			log.Fatal(err)
		}

		if *lines {
			fmt.Printf("% 10d ", c.Lines)
		}
		if *words {
			fmt.Printf("% 10d ", c.Words)
		}
		if *multibytes {
			fmt.Printf("% 10d ", c.Multibytes)
		}
		if *bytes {
			fmt.Printf("% 10d ", c.Bytes)
		}
	} else {
		var multibytes_total, lines_total, words_total, bytes_total uint64

		for _, filepath := range flag.Args() {

			file, err := os.Open(filepath)
			if err != nil {
				log.Fatal(err)
			}

			c := wc.NewCounter(file)

			err = c.Count(*multibytes, *bytes, *lines, *words)
			if err != nil {
				log.Fatal(err)
			}

			file.Close()

			if *lines {
				lines_total += c.Lines
				fmt.Printf("% 10d ", c.Lines)
			}
			if *words {
				words_total += c.Words
				fmt.Printf("% 10d ", c.Words)
			}
			if *multibytes {
				multibytes_total += c.Multibytes
				fmt.Printf("% 10d ", c.Multibytes)
			}
			if *bytes {
				bytes_total += c.Bytes
				fmt.Printf("% 10d ", c.Bytes)
			}

			fmt.Printf("%s\n", filepath)
		}

		if flag.NArg() > 1 {

			if *lines {
				fmt.Printf("% 10d ", lines_total)
			}
			if *words {
				fmt.Printf("% 10d ", words_total)
			}
			if *multibytes {
				fmt.Printf("% 10d ", multibytes_total)
			}
			if *bytes {
				fmt.Printf("% 10d ", bytes_total)
			}

			fmt.Print("total\n")

		}
	}

}
