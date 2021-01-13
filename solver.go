package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// Solver ....
type Solver struct{}

// NewSolver ...
func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) number1(number int) {
	fmt.Printf("\nSOAL 1\n")
	fmt.Printf("Default number : %d\n", number)
	var (
		tmpStringNumber = strconv.Itoa(number)
		makeZeroString  = func(n int) string {
			var tmpStr []string
			for i := 0; i < n-1; i++ {
				tmpStr = append(tmpStr, "0")
			}
			return strings.Join(tmpStr, "")
		}
	)

	fmt.Println(tmpStringNumber)
	for i, num := range tmpStringNumber {
		tmpNum := int(num - '0')
		fmt.Printf("%d%s\n", tmpNum, makeZeroString(len(tmpStringNumber)-i))
	}
}
func (s *Solver) number2(arr []int) {
	fmt.Printf("\nSOAL 2\n")
	fmt.Printf("List number : %v\n", arr)
	var (
		counter   int
		groupData [][]int
		tmpCount  = len(arr)
		limit     = int(math.Ceil(float64(tmpCount) / float64(3)))
		ascSort   = func(tmpArry []int) []int {
			sort.Slice(tmpArry, func(i, j int) bool {
				return tmpArry[i] < tmpArry[j]
			})
			return tmpArry
		}
		min = func(arr []int) int {
			tmpMax := arr[0]
			for _, val := range arr {
				if val < tmpMax {
					tmpMax = val
				}
			}
			return tmpMax
		}
		max = func(arr []int) int {
			tmpMin := arr[0]
			for _, val := range arr {
				if val > tmpMin {
					tmpMin = val
				}
			}
			return tmpMin
		}
		calculateArray = func(tmpArray []int, group int, tMin, tMax func(arr []int) int) {
			sum := 0
			for i := 0; i < len(tmpArray); i++ {
				sum += i
			}
			fmt.Printf("Sum array group %d : {%v}\n", group, sum)
			fmt.Printf("Mean array group %d : {%v}\n", group, float64(float64(sum)/float64(len(tmpArray))))
			fmt.Printf("Max array group %d : {%v}\n", group, tMax(tmpArray))
			fmt.Printf("Min array group %d : {%v}\n", group, tMin(tmpArray))
		}
	)
	for i := 0; i < tmpCount; i += limit {
		counter += limit
		if counter > tmpCount {
			counter = tmpCount
		}
		if len(arr[i:counter]) != 0 {
			groupData = append(groupData, arr[i:counter])
		}
	}
	for index, val := range groupData {
		groupCount := index + 1
		ascSort(val)
		fmt.Printf("\nSort asc array group %d : {%v}\n", groupCount, val)
		calculateArray(val, groupCount, min, max)

	}
}
func (s *Solver) number3(text string) {
	fmt.Printf("\nSOAL 3\n")
	fmt.Printf("Text : %s\n", text)
	dupeRemove := make(map[rune]int)
	tmpString := []string{}
	stringValue := func(i rune) string {
		text := ""
		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for _, v := range chars {
			if i == v {
				text = string(v + 5)
			}
		}
		return text
	}
	for _, charText := range strings.ToLower(text) {
		tmpChar := string(charText)
		if _, ok := dupeRemove[charText]; ok {
			dupeRemove[charText]++
			continue
		}
		if charText > 0 {
			tmpChar = stringValue(charText)
		}
		tmpString = append(tmpString, tmpChar)
		dupeRemove[charText] = 1
	}
	for key, val := range dupeRemove {
		fmt.Println(string(key), "=", val)
	}
	fmt.Println("After move character + 5", tmpString)

}
func (s *Solver) number4(num int) {
	fmt.Printf("\nSOAL 4\n")
	fmt.Printf("Secret number %d\n", num)
	var (
		min = 1
		max = 100
		counter = 0

	)
	for {
		counter++
		guess := rand.Intn(rand.Intn(max-min)+1)
		if guess > num {
			fmt.Printf("Guess number %d higher than secret number, [SECRET_NUMBER = %d]\n", guess, num)
			continue
		}
		if guess < num {
			fmt.Printf("Guess number %d lower than secret number, [SECRET_NUMBER = %d]\n", guess, num)
			continue
		}
		fmt.Printf("\nGuess number %d same as secret number, after %d spin\n", guess, counter)
		break
	}
}
