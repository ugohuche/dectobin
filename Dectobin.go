package main

import "fmt"

func main(){
  var Number int

  fmt.Println("Please type in a number !")
  fmt.Scan(&Number)
  fmt.Printf(" %d in binary format = %b\n", Number, Number)
  fmt.Println("Would you like to try another number ? yes/no")
  for {
    var answer string
    fmt.Scan(&answer)
    if answer == "yes"{
      fmt.Println("Please type in a number !")
      fmt.Scan(&Number)
      fmt.Printf(" %d in binary format = %b\n", Number, Number)
      fmt.Println("Would you like to try another number ? yes/no")
    } else if answer == "no"{
      fmt.Println("Goodbye")
      break
    } else if answer != "yes" || answer != "no" {
      fmt.Println("Please type in yes or no")
    }
          
  }

}
