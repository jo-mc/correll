package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	file1, err := os.Open("../CigarRegSPlit/srt_regpol.txt")  //regpol.txt")
	//"pol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()
	file2, err := os.Open("../CigarRegSPlit/regnopol.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	var s1 []string
	var s2 []string
	var nomatch bool
	var nmCount uint32
	var nmReads uint32
	var lineNum uint32
	nmCount = 0
	nmReads = 0
	lineNum = 0

	fmt.Println("chrPos readLen from% to% Clip line")

	for scanner1.Scan() { // loop until EOF
		lineNum++ // to ref read ID if need to look up something ie obscure data point.
		s1 = strings.Split(scanner1.Text(), ",")
		if len(s1) < 8 {
			//fmt.Println("len s1:", len(s1))
			continue
		}
		// fmt.Print(s1[8], s1[9], s1[11])
		_, err = file2.Seek(0, io.SeekStart)
		scanner2 = bufio.NewScanner(file2)
		nomatch = true
		nmReads++
//		 		if nmReads > 90 {
//			break
//		} 
		
		for scanner2.Scan() {
			s2 = strings.Split(scanner2.Text(), ",")
			if len(s2) < 8 {
				//fmt.Println("len s2:", len(s2))
				continue
			}
			if (s2[8] + s2[9]) == (s1[8] + s1[9]) {
				//fmt.Print(" MATCH. ")
				//fmt.Println("M from: ", s2[1], " to ", s1[1], ", M count:", s1[2], " I:D", s2[3], "-", s1[3], " S:H (clip) ", s2[4], "-", s1[4])
				// for R
				if s1[9] == " chr6 " {
					cMi, err := strconv.Atoi(s1[2])
					cIDi, err := strconv.Atoi(s1[3])
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(s1[11], (cMi + cIDi), s2[1], s1[1], s1[4], lineNum) // id: s1[8])  will need later, convert to hex of source file line #
				}
				nomatch = false
				break
			}
		}
		if nomatch {
			// fmt.Println(" - NO MATCH - ")
			nmCount++
		}
	}

	fmt.Println("Polished reads not found: ", nmCount, " of ", nmReads, " reads.")

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
