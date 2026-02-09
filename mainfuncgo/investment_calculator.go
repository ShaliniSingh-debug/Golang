package main
import 
("fmt"
"math"
)


// var investedAmount float64
// 	var anualReturn float64
// 	var years float64
// 	var futureValue float64
// 	var futureRealValue float64
const inflation = 2.5
func main() {
    investedAmount, anualReturn, years := enterValue()

    futureValue, futureRealValue := returnAmount(
        investedAmount,
        anualReturn,
        years,
        inflation,
    )

    fmt.Printf("Future Value: %.2f\n", futureValue)
    fmt.Printf("Future Value after Inflation:%.2f", futureRealValue)
}

func enterValue() (float64, float64, float64) {
    var investedAmount float64
    var anualReturn float64
    var years float64

    fmt.Print("Please enter the investing amount: ")
    fmt.Scan(&investedAmount)

    fmt.Print("Please enter the annual return: ")
    fmt.Scan(&anualReturn)

    fmt.Print("Please enter the years: ")
    fmt.Scan(&years)

    return investedAmount, anualReturn, years
}

func returnAmount(investedAmount, anualReturn, years, inflation float64) (float64, float64) {
    futureValue := investedAmount * math.Pow(1+anualReturn/100, years)
    futureRealValue := futureValue / math.Pow(1+inflation/100, years)

    return futureValue, futureRealValue
}
