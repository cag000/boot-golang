package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
	- Jelaskan secara detail tentang bahasa pemrograman yang sedang anda kerjakan
		Golang (atau biasa disebut dengan Go) adalah bahasa pemrograman baru yang dikembangkan di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson pada tahun 2007 dan mulai diperkenalkan ke publik tahun 2009.
		Penciptaan bahasa Go didasari bahasa C dan C++, oleh karena itu gaya sintaksnya mirip.
		Go memiliki kelebihan dibanding bahasa lainnya, beberapa di antaranya:
			Mendukung konkurensi di level bahasa dengan pengaplikasian cukup mudah
			Mendukung pemrosesan data dengan banyak prosesor dalam waktu yang bersamaan (pararel processing)
			Memiliki garbage collector
			Proses kompilasi sangat cepat
			Bukan bahasa pemrograman yang hirarkial, menjadikan developer tidak perlu ribet memikirkan segmen OOP-nya
			Package/modul yang disediakan terbilang lengkap. Karena bahasa ini open source, banyak sekali developer yang juga mengembangkan modul-modul lain yang bisa dimanfaatkan
	
	- Jelaskan aturan tentang bahasa pemrograman yang sedang anda kerjakan
		untuk aturan golang dapat dilihat lebih lengkap pada
		https://github.com/a8m/golang-cheat-sheet
		https://gobyexample.com/
		https://tour.golang.org/
		https://blog.golang.org/
	
*/

const (
	// Author ...
	Author = "SYAHRUL AL-RASYID"
	// Repo ...
	Repo = "github.com/cag000"
	// Email ...
	Email = "syahrull75.sar@gmail.com"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	solver := NewSolver()

	fmt.Printf("%s || %s || %s\n", Author, Repo, Email)
	
	// All parameter is defaul question in each case
	solver.number1(1234567)
	solver.number2([]int{1,2,33,44,55,33,44,22,44,33,2,77,66,1,2,3,4,5,6,7,89,3,3,8,9,75,4,3,2,2,1,3})
	solver.number3("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras interdum mi eu magna fermentum, vel luctus tellus semper. Nunc dignissim eleifend ipsum, nec viverra mauris pellentesque non. Fusce auctor ex id mauris egestas, quis luctus nunc pharetra. Sed in dignissim nisi. Aliquam sed tempor urna, nec aliquam mi. Aenean eu feugiat lacus, vel dictum eros. Nulla condimentum porttitor aliquet. Vestibulum vehicula elit non arcu auctor maximus. Quisque est eros, maximus nec diam faucibus, mollis odio.")
	solver.number4(rand.Intn(100-1)+1)
	
}