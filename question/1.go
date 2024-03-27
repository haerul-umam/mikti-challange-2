package question

import (
	"challange-2/question/helper"
	"fmt"
)

/*
for storing team's data
since there is an average score calculation, we use float for scores
*/
type Team struct {
  Scores []float64
  Name   string
}

/*
because we need method for add scores, we defined method on pointers
used slices instead float for params func to allow adding more scores
*/
func (team *Team) AddScore(value []float64) []float64 {
  team.Scores = append(team.Scores, value...)
  return team.Scores
}

func (team *Team) ResetScores() {
  team.Scores = []float64{}
}

func (team *Team) ScoreAvarage() float64 {
  var temp float64 = 0
  for _, i := range team.Scores {
    temp += i
  }
  var result = temp / float64(len(team.Scores))
  return result
}

func FirstQuestion() {
  var (
    koalaTeam = Team{ Name: "Koala" }
    dolphinTeam = Team{ Name: "Lumba-lumba " }
    teams = [...]*Team{ &dolphinTeam, &koalaTeam } // array teams, used pointer for update their score
    totalMatch = 3 // for flexibility in the number of matches
  )

  // impure function for reseting scores after the match is finished
  resetScores := func()  {
    for _, team := range teams {
      team.ResetScores()
    }
  }

  getUserInput := func() {
    resetScores()
     // loop each teams for user input interaction
    for _, team := range teams {
      fmt.Println("Skor Tim", team.Name)
      for i := 0; i < totalMatch; i++ {
        score := helper.ReadUserNumberInput(fmt.Sprintf("\tSkor ke-%d: ", i+1))
        team.AddScore([]float64{score})
      }
    }
  }
  
	fmt.Println("Soal pertama (interactive mode)")
  fmt.Println("Masukan skor untuk masing-masing Tim: (Data 1)")
  
  // input (data1)
  getUserInput()

  avgKoalaScore := koalaTeam.ScoreAvarage()
  avgDolphonScore := dolphinTeam.ScoreAvarage()

  // check score data 1
  if avgDolphonScore == avgKoalaScore {
    fmt.Println("Hasil Seri")
  } else if (avgDolphonScore > avgKoalaScore) {
    fmt.Println("Hasil menang: ", dolphinTeam.Name)
  } else {
    fmt.Println("Hasil menang: ", koalaTeam.Name)
  }

  fmt.Println("\nMasukan skor untuk masing-masing Tim: (Bonus 1)")
  // input (bonus1)
  getUserInput()

  var minimumScore = 100. // use float since avarage type data is float
  
  fmt.Printf(`
Bonus 1, menang jika:
- Minimum skor %g
- Lebih tinggi dari tim lain
  `, minimumScore)

  avgKoalaScore = koalaTeam.ScoreAvarage()
  avgDolphonScore = dolphinTeam.ScoreAvarage()

  fmt.Println("Hasil bonus 1")
  if avgKoalaScore >= minimumScore && avgDolphonScore >= minimumScore {
    if avgDolphonScore == avgKoalaScore {
      fmt.Println("Hasil Seri")
    } else if (avgDolphonScore > avgKoalaScore) {
      fmt.Println("Hasil menang: ", dolphinTeam.Name)
    } else {
      fmt.Println("Hasil menang: ", koalaTeam.Name)
    }
  } else {
    fmt.Println("Hasil bonus 1 : Tidak ada pemenang")
  }

  fmt.Println("\nMasukan skor untuk masing-masing Tim: (Bonus 2)")
  // input (bonus2)
  getUserInput()

  fmt.Printf(`
Bonus 2, seri jika:
- Skor sama
- Minimum skor %g atau lebih besar
  `, minimumScore)

  avgKoalaScore = koalaTeam.ScoreAvarage()
  avgDolphonScore = dolphinTeam.ScoreAvarage()

  if 
    (avgKoalaScore == avgDolphonScore) && 
    (avgKoalaScore >= minimumScore) && 
    (avgDolphonScore >= minimumScore) {
      fmt.Println("Hasil bonus 2 : Seri")
  } else {
    fmt.Println("Hasil bonus 2 : Tidak ada pemenang")
  }
  fmt.Println("")
}