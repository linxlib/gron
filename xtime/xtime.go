package xtime

import "time"

// extends time.Duration

const (
	//Second has 1 * 1e9 nanoseconds
	Second time.Duration = time.Second
	//Minute has 60 seconds
	Minute time.Duration = time.Minute
	//Hour has 60 minutes
	Hour time.Duration = time.Hour
	//Day has 24 hours
	Day time.Duration = time.Hour * 24
	//Week has 7 days
	Week time.Duration = Day * 7
)

var (
	BeginOfMonth     = time.Now().AddDate(0, 0, -time.Now().Day()+1)
	BeginOfMonthZero = time.Date(BeginOfMonth.Year(), BeginOfMonth.Month(), BeginOfMonth.Day(), 0, 0, 0, 0, BeginOfMonth.Location())
	EndOfMonth       = BeginOfMonth.AddDate(0, 1, -1)
	EndOfMonthZero   = time.Date(EndOfMonth.Year(), EndOfMonth.Month(), EndOfMonth.Day(), 0, 0, 0, 0, EndOfMonth.Location())
)
