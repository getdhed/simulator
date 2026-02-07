package models

type DayResult struct {
	Day         int
	Workers     int
	TotalEarned float64
	PerUser     map[int]float64
}
