/*Обьекты дат с методами добавления дней/месяцев/годов
и другими методами

*/

package main

import (
	"fmt"
	"math"
)

type date struct {
	day   int
	month int
	year  int
}

func (t *date) add_day(day_ch int) {
	t.day += day_ch
	if t.day > 30 {
		t.month += t.day / 30
		t.day = t.day % 30
		if t.month > 12 {
			t.year += t.month / 12
			t.month = t.month%12 + 1
		}
	}
}

func (t *date) add_month(month_ch int) {
	t.month += month_ch
	if t.month > 12 {
		t.year += t.month / 12
		t.month = t.month%12 + 1
	}
}
func (t1 *date) differences(t2 *date) {
	fmt.Printf("The differnces between to dates in dyas is %f",
		math.Abs(float64((t1.day+t1.month*30+t1.year*365)-(t2.day+t2.month*30+t2.year*365))))
}

func start(dates []date) {
	for {
		fmt.Print("Choose option: \n" +
			"1- Add dates\n" +
			"2- Show dates\n" +
			"3- Add days/month/years to date\n" +
			"4- Find differences in days between to dates\n" +
			"5- Exit\n" +
			"->..")
		var ch int
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			var d int
			var m int
			var y int
			fmt.Print("Enter the day: ")
			fmt.Scanln(&d)
			fmt.Print("Enter the month: ")
			fmt.Scanln(&m)
			fmt.Print("Enter the year: ")
			fmt.Scanln(&y)

			dates = append(
				dates, date{d, m, y},
			)
		case 2:
			for i, v := range dates {
				fmt.Printf("%d date - %d/%d/%d\n", i+1, v.day, v.month, v.year)
			}
		case 3:
			fmt.Println("later")
		case 4:
			fmt.Println("later")
		case 5:
			return
		default:
			fmt.Println("Wrong input! Try again!")

		}
	}
}

func main() {
	/*var day1 *date = &date{22,5,2003}
	fmt.Print("Enter the month: ")
	var x int
	fmt.Scanln(&x)
	day1.month = x
	fmt.Printf("%d/%d/%d - your date\n", day1.day, day1.month, day1.year)
	day1.add_day(365)
	fmt.Printf("%d/%d/%d - your date\n", day1.day, day1.month, day1.year)
	day1.add_month(5)
	fmt.Printf("%d/%d/%d - your date\n", day1.day, day1.month, day1.year)
	var day2 *date = &date{12, 3, 2015}
	fmt.Printf("%d/%d/%d - your 2nd date\n", day2.day, day2.month, day2.year)
	day1.differences(day2)*/

	dates := []date{}
	start(dates)

}
