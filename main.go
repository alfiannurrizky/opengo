package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)

    var grab int
    flag.IntVar(&grab, "t", 5, "Set the number of URLs to process")
    flag.Parse()

    if flag.NFlag() == 0 {
        grab = -1
    }

    for i := 0; (grab < 0 || i < grab) && sc.Scan(); i++ {
        urls := sc.Text()
        urlList := strings.Fields(urls)

        for _, url := range urlList {
            exec.Command("open", url).Run()

            fmt.Printf("Successfully opened %s\n", url)
        }
    }

    if err := sc.Err(); err != nil {
        fmt.Printf("Error reading input: %v\n", err)
    }
}
