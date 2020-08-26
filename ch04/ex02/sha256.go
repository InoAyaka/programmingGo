package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	opt := flag.Int("opt", 256, "support SHA256,SHA384,SHA512")

	flag.Parse()

	switch *opt {
	case 256:
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		fmt.Println(sha256.Sum256([]byte(input.Text())))
	case 384:
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		fmt.Println(sha512.Sum384([]byte(input.Text())))
	case 512:
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		fmt.Println(sha512.Sum512([]byte(input.Text())))
	default:
		fmt.Fprintf(os.Stderr, "Wrong flag specified.\n")
		os.Exit(1)
	}

}
