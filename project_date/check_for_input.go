package main

import "fmt"

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
