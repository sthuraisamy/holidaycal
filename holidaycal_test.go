package holidycal

import (
	"testing"
	"fmt"
	"time"
)

func TestGoodFridayDay(t *testing.T) {
	dt := time.Now()

	y := 2010
	for y < 2021 {
		m,i := calculateGoodFriday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestEasterMonday(t *testing.T) {
	dt := time.Now()

	y := 2010
	for y < 2021 {
		m,i := calculateEasterMonday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestLabourday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateLabourday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}
func TestChristmasday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateChristmasday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestBoxingday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateBoxingday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestNewyearday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateNewyearday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestVictoriaDay(t *testing.T) {
	dt := time.Now()

	y := 2010
	for y < 2021 {
		m,i := calculateCAVictoriaDay(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestThanksgivingDay(t *testing.T) {
	dt := time.Now()

	y := 2010
	for y < 2021 {
		m,i := calculateCAThanksgivingDay(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestFamilyDay(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateCAFamilyday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestCAIndependence(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateCAIndependence(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestUSThanksgivingDay(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateUSThanksgivingday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestMemorialDay(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateUSMemorialday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}


func TestUSPresidentsday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateUSPresidentsday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestUSIndependenceday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateUSIndependenceday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestUSColumbusday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateUSColumbusday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}

func TestUSVeteransday(t *testing.T) {
	dt := time.Now()
	y := 2010
	for y < 2021 {
		m,i := calculateUSVeteransday(y, dt.Location())
		fmt.Printf("For year %d/%d/%d\n", y,m, i)
		y++
	}
}