package main

import (
	"fmt"
	"testing"
)

func TestSanityCheck(t *testing.T) {
	got := sanityCheck()
	want := "its working"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := multiply(5, 7)
	want := 35

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDogYears(t *testing.T) {
	got := dogYears(5)
	want := 35

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// Table driven tests: https://github.com/golang/go/wiki/TableDrivenTests
func TestRpsGame(t *testing.T) {
	tests := []struct {
		user, computer string
		want           string
	}{
		{"rock", "scissorcs", "you win!"},
		{"rock", "paper", "you lose!"},
		{"rock", "rock", "its a tie!"},
		{"paper", "rock", "you win!"},
		{"paper", "scissors", "you lose!"},
		{"paper", "paper", "its a tie!"},
		{"scissors", "paper", "you win!"},
		{"scissors", "rock", "you lose!"},
		{"scissors", "scissors", "its a tie!"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("user: %s, computer: %s", tt.user, tt.computer)
		t.Run(testname, func(t *testing.T) {
			got := game(tt.user, tt.computer)
			want := tt.want

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}

func TestMiles(t *testing.T) {
	got := miles(10)
	want := 6.2371

	if got != want {
		t.Errorf("got %d got %d", got, want)
	}
}

func TestFeet(t *testing.T) {
	got := feet(160)
	want := 5.2493438320209975

	if got != want {
		t.Errorf("got %d got %d", got, want)
	}
}

func TestAnnoyingSong(t *testing.T) {
	got := annoyingSong(5)
	want := fmt.Sprintf(
		"%d bottles of soda on the wall, %d bottles of soda, take one down pass it around %d bottles of soda on the wall",
		5, 5, 4)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		grade int
		want  string
	}{
		{90, "you got an A"},
		{95, "you got an A"},
		{80, "you got an B"},
		{85, "you got an B"},
		{70, "you got an C"},
		{75, "you got an C"},
		{75, "you got an C"},
		{60, "you got an D"},
		{65, "you got an D"},
		{59, "you got an F"},
		{20, "you got an F"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("user: %s, computer: %s", tt.user, tt.computer)
		t.Run(testname, func(t *testing.T) {
			got := grade(tt.grade)
			want := tt.want

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}
