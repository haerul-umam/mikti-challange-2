package helper

import "fmt"

func ReadUserNumberInput(question string) float64 {
  var num float64
  // infinite loop until user input correct number
  for {
    // question title
    fmt.Print(question)
    _, err := fmt.Scan(&num)
    if err != nil {
      fmt.Println("Invalid input. Hanya menerima input berupa angka.")
    } else {
      break
    }
  }
  return num
}