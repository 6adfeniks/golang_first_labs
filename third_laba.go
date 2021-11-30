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
	fmt.Printf("The differnces between to dates in dyas is %f\n",
		math.Abs(float64((t1.day+t1.month*30+t1.year*365)-(t2.day+t2.month*30+t2.year*365))))
}

func check_month(monthh int) int {
	var tmonth int
	k := 0
	for i := monthh; i < 1 || i > 12; {
		fmt.Println("Wrong input, try again!")
		fmt.Scanln(&tmonth)
		i = tmonth
		k++
	}
	if k > 0 {
		return tmonth
	}
	return monthh
}

func check_day(dayy int) int {
	var tday int
	k := 0
	for i := dayy; i < 1 || i > 30; {
		fmt.Println("Wrong input, try again!")
		fmt.Scanln(&tday)
		i = tday
		k++
	}
	if k > 0 {
		return tday
	}
	return dayy
}

func check_year(yearr int) int {
	var tyear int
	k := 0
	for i := yearr; i < 0; {
		fmt.Println("Wrong input, try again!")
		fmt.Scanln(&tyear)
		i = tyear
		k++
	}
	if k > 0 {
		return tyear
	}
	return yearr
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
			d = check_day(d)

			fmt.Print("Enter the month: ")
			fmt.Scanln(&m)
			m = check_month(m)

			fmt.Print("Enter the year: ")
			fmt.Scanln(&y)
			y = check_year(y)

			dates = append(
				dates, date{d, m, y},
			)
		case 2:
			for i, v := range dates {
				fmt.Printf("%d date - %d/%d/%d\n", i+1, v.day, v.month, v.year)
			}
		case 3:
			fmt.Println("Choose the date u want to change:")
			for i, v := range dates {
				fmt.Printf("%d date - %d/%d/%d\n", i+1, v.day, v.month, v.year)
			}
			var ch int
			fmt.Scanln(&ch)

			var dayA int
			var monthA int
			fmt.Println("Enter the count of days and months u want to add: ")
			fmt.Scanln(&dayA, &monthA)
			dates[ch-1].add_day(dayA)
			dates[ch-1].add_month(monthA)

		case 4:
			fmt.Println("Choose two dates(ex. 1 3) u want to find sub between them:")
			var x int
			var y int
			fmt.Scanln(&x, &y)
			dates[x-1].differences(&dates[y-1])
		case 5:
			return
		default:
			fmt.Println("Wrong input! Try again!")

		}
	}
}

func main() {
	dates := []date{}
	start(dates)
}
