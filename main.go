package main

func main() {
	// PACKAGE
	// fmt.Println("helo, belajar golang")
	// sentence := TestAja()

	// fmt.Println(sentence)

	// result := calculation.Add(8, 9)
	// fmt.Println(result)

	// resultMultiply := calculation.Multiply(5, 5)
	// fmt.Println(resultMultiply)

	// VARIABLE
	// var name string
	// name = "Golang"
	// age := 20

	// fmt.Println(name, age)

	//IF ELSE
	// age := 10

	// if age > 10 {
	// 	fmt.Println("boleh bermain")
	// } else {
	// 	fmt.Println("tidak boleh bermain")
	// }

	// score := 65

	// var grade string

	// if score <= 50 {
	// 	grade = "E"
	// } else if score <= 60 {
	// 	grade = "D"
	// } else if score < 70 {
	// 	grade = "C"
	// } else {
	// 	grade = "NULL"
	// }
	// fmt.Println(grade)

	// SWITCH CASE

	// number := 2

	// switch number {
	// case 1:
	// 	fmt.Println("SATU")
	// case 2:
	// 	fmt.Println("DUA")
	// case 3:
	// 	fmt.Println("TIGA")
	// default:
	// 	fmt.Println("DEFAULT")
	// }

	// LOOPING
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("saya sedang belajar: ", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("saya sedang belajar: ", i)
	// 	i++
	// }

	// title := "golang the best language"

	// for _, letter := range title {
	// 	fmt.Println("letter :", string(letter))
	// }

	// title := "golang the best language"

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index:", index, "letter: ", string(letter))
	// 	}
	// }

	// for index, letter := range title {
	// 	letterString := string(letter)

	// 	if letterString == "a" || letterString == "i" || letterString == "e" || letterString == "o" {
	// 		fmt.Println("index:", index, "letter: ", string(letter))

	// 	}

	// 	switch letterString {
	// 	case "a", "i", "u", "e", "o":
	// 		fmt.Println("index:", index, "letter: ", string(letter))

	// 	}
	// }

	// ARRAY
	// var languages [5]string
	// languages[0] = "GO"
	// languages[1] = "Ruby"
	// languages[2] = "js"
	// languages[3] = "c#"
	// languages[4] = "python"

	// languages := [...]string{
	// 	"ruby",
	// 	"python",
	// 	"js",
	// 	"c#",
	// 	"go",
	// 	"nodejs",
	// }

	// for index, lang := range languages {
	// 	fmt.Println("index: ", index, "language:", lang)
	// }

	// SLICE
	// var gamingConsole []string
	// gamingConsole := []string{"ps", "switch", "pc"}

	// gamingConsole = append(gamingConsole, "ps")
	// gamingConsole = append(gamingConsole, "switch")
	// gamingConsole = append(gamingConsole, "pc")

	// for _, console := range gamingConsole {
	// 	fmt.Println("gaming console : ", console)

	// }

	// MAP

	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["GO"] = 9
	// myMap["JS"] = 8

	// myMap := map[string]string{
	// 	"ruby": "is beautiful",
	// 	"go":   "is super fast",
	// 	"js":   "hype",
	// }

	// for key, value := range myMap {
	// 	fmt.Println("key:", key, "value: ", value)
	// }

	// fmt.Println("=========")

	// delete(myMap, "ruby")

	// for key, value := range myMap {
	// 	fmt.Println("key:", key, "value: ", value)
	// }

	// value, isAvailabe := myMap["python"]
	// fmt.Println(isAvailabe)
	// fmt.Println(value)

	// SLICE MAP

	// students := []map[string]string{
	// 	{"name": "agung", "score": "A"},
	// 	{"name": "link", "score": "B"},
	// 	{"name": "mario", "score": "C"},
	// }

	// for _, student := range students {
	// 	fmt.Println("name: ", student["name"], "Raport: ", student["score"])
	// }

	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	// var total int

	// for _, score := range scores {
	// 	total = total + score
	// }

	// length := len(scores)
	// average := float64(total) / float64(length)
	// fmt.Println(average)

	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// var goodScore []int

	// for _, score := range scores {
	// 	if score >= 90 {
	// 		goodScore = append(goodScore, score)
	// 	}
	// }

	// fmt.Println(goodScore)

}
