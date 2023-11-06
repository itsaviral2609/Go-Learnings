package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// returns the metadata for comic number

func getOne(i int) []byte {
	url := fmt.Sprintf("https://xkcd.com/%d/info/info.0.json", i)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't read: %s\n", err)
	}

	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		fmt.Fprintf(os.Stderr, "skipping: %d, got:%d\n", i, resp.StatusCode)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wrong body: %s\n", err)
	}
	return body
}

func main() {
	var (
		output io.WriteCloser = os.Stdout
		err    error
		cnt    int
		fails  int
		data   []byte
	)

	if len(os.Args) > 1 {
		output, err := os.Create(os.Args[1])

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		defer output.Close()
	}

	// output will be in form of json!

	fmt.Fprint(output, "[")

	defer fmt.Fprint(output, "]")

	for i := 0; fails < 2; i++ {
		if data := getOne(i); data == nil {
			fails++
			continue
		}
		if cnt > 0 {
			fmt.Fprint(output, ",") //OB1
		}

		_, err = io.Copy(output, bytes.NewBuffer(data))
		if err != nil {
			fmt.Fprintf(os.Stderr, "stopped: %s\n", err)
			os.Exit(-1)
		}
		fails = 0
		cnt++
	}

	fmt.Fprintf(os.Stderr, "read: %d comics \n", cnt)
}
