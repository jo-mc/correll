package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var errStream = os.Stderr
var errFile string
var err error

func main() {

	errFile = "errdata.txt"

	fmt.Println("Welcome to the correll!")
	errStream, err = os.Create(errFile)
	defer errStream.Close()
	fmt.Println("The time is", time.Now())

	file1, err := os.Open("pol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	file2, err := os.Open("nopol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file1)

	var s1 []string
	var s2 []string
	var nomatch bool
	var nmCount uint32
	var nmReads uint32
	nmCount = 0
	nmReads = 0

	for scanner1.Scan() { // loop until EOF
		s1 = strings.Split(scanner1.Text(), ",")
		fmt.Println(s1[8])
		_, err = file2.Seek(0, io.SeekStart)
		scanner2 = bufio.NewScanner(file2)
		nomatch = true
		nmReads += 1
		for scanner2.Scan() {
			s2 = strings.Split(scanner2.Text(), ",")
			if s2[8] == s1[8] {
				fmt.Print(s2[8], " - MATCH. ")
				fmt.Println("M from: ", s2[1], " to ", s1[1], ", M count:", s1[2], " I:D", s2[3], "-", s1[3], " S:H (clip) ", s2[4], "-", s1[4])

				nomatch = false
				break
			}
		}
		if nomatch {
			fmt.Println(" - NO MATCH - ")
			nmCount += 1
		}
	}

	fmt.Println("Polished reads not found: ", nmCount, " of ", nmReads)

	/* 	fmt.Println("file2::::::::::::")

	   	scanner2.Scan()
	   	//for scanner2.Scan() { // loop until EOF
	   	fmt.Println(scanner2.Text())
	   	//}
	   	fmt.Println("file 2 line 1 again::::::::::::::::")
	   	_, err = file2.Seek(0, io.SeekStart)
	   	fmt.Println("err: ", err)
	   	scanner2 = bufio.NewScanner(file2)
	   	scanner2.Scan()
	   	fmt.Println(scanner2.Text())
	   	//	for j := 1; j <= 28; j++ {
	   	//		if scanner.Scan() { //  test for EOF
	   	//			f = scanner.Text() */

}
