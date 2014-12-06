package main

import "fmt"
import "time"

func main() {

     i := 2
     fmt.Print("Write", i, " as ")
     switch i {
     case 1:
     	  fmt.Println("one")
     case 2:
     	  fmt.Println("two")
     case 3:
     	  fmt.Println("three")
     }

     switch time.Now().Weekday() {
     case time.Saturday, time.Sunday:
     	  fmt.Println("Weekend")
     default:
	  fmt.Println("Workday")
     }

     t := time.Now()
     switch {
     case t.Hour() < 12:
     	  fmt.Println(t, " - It's before noon")
     default:
	  fmt.Println(t, " - It's after noon")
     }
}