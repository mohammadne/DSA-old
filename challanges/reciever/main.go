package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// go run main.go < in.txt > out.txt

func onLine() {
	var a, b int
	fmt.Scan(&a, &b)
}

func readFileWithReadString(fn string) (err error) {
	reader := bufio.NewReader(os.Stdin)
	var line string

	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		// Process the line here.
		fmt.Printf(" > Read %d characters\n", len(line))
		fmt.Printf(" > > %s\n", limitLength(line, 50))

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}

	return
}

func readFileWithScanner(fn string) (err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Process the line here.
		fmt.Printf(" > Read %d characters\n", len(line))
		fmt.Printf(" > > %s\n", limitLength(line, 50))
	}

	if scanner.Err() != nil {
		fmt.Printf(" > Failed with error %v\n", scanner.Err())
		return scanner.Err()
	}

	return
}

func readFileWithReadLine(fn string) (err error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		var buffer bytes.Buffer

		var l []byte
		var isPrefix bool

		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)
			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}
			// If we're at the EOF, break.
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
		}

		line := buffer.String()

		// Process the line here.
		fmt.Printf(" > Read %d characters\n", len(line))
		fmt.Printf(" > > %s\n", limitLength(line, 50))

		if err == io.EOF {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return err
	}

	return
}

func limitLength(s string, length int) string {
	if len(s) < length {
		return s
	}

	return s[:length]
}
