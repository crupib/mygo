package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    "time"
//    "reflect"
)

func main() {
    // open file
    f, err := os.Open("problems.csv")
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    var answer int 
    var Number_of_wrong, Number_of_right, Number_of_questions int
    for {
        start := time.Now()
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }
       
        // do something with read line
        fmt.Printf("What is %+v\n", rec[0])
        fmt.Scanln(&answer)
        actual_answer, _ := strconv.Atoi(rec[1])
        if answer == actual_answer {
          fmt.Printf("Correct\n")
          Number_of_right++
        } else {
            fmt.Printf("Wrong\n")    
            Number_of_wrong++
        }
        Number_of_questions++;
        elapsed := time.Since(start)
        accumtime := elapsed.Seconds()
        if accumtime >= 30.0 {
           fmt.Println(accumtime)
           break
        }
    }
    tempvar := fmt.Sprintf("Number of questions %d, Number of right answers %d, and the Number of wrong answers %d\n",Number_of_questions, Number_of_right,Number_of_wrong)
    fmt.Printf(tempvar)
 }
