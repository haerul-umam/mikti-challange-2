package question

import "fmt"

type Body struct {
	Height float32
	Weight float32
}

func calculateBMI(height, weight float32) float32 {
	return weight / (height * height)
}

func SecondQuestion()  {
	fmt.Println("Soal kedua (Hard coded)")

	//multidimensional array
	trainingData := [2][2][2]float32{
		{
			{1.69, 78}, //mark
			{1.95, 92}, //john
		},
		{
			{1.88, 95}, //mark
			{1.76, 85}, //john
		},
	}

	// loop training data to avoid repetition code if more training data
	for i := 0; i < len(trainingData); i++ {
		fmt.Println("Data", i+1)
		data := trainingData[i]
		var (
			mark = Body{ Height: data[0][0], Weight: data[0][1] }
			john = Body{ Height: data[1][0], Weight: data[1][1] }
			markHigherBMI bool
		)
	
		markBmi := calculateBMI(mark.Height, mark.Weight)
		johnBmi := calculateBMI(john.Height, john.Weight)
	
		markHigherBMI = markBmi > johnBmi
	
		fmt.Println("BMI Mark :", markBmi)
		fmt.Println("BMI John :", johnBmi)
		fmt.Println("BMI Mark lebih tinggi dari pada John : ", markHigherBMI)
	}
}