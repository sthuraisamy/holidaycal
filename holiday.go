package holidycal

import "time"

type ObservedRule int
const (
	ObservedNearest ObservedRule = iota 			// nearest weekday (Friday or Monday)
	ObservedExact                       			// the exact day only
	ObservedMonday                      			// Monday always
)

/*
US holiday calculation.
 */
func calculateUSThanksgivingday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixedWeekdayHoliday(year, time.November, time.Thursday, 4, loc)
}

func calculateUSPresidentsday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixedWeekdayHoliday(year, time.February, time.Monday, 3, loc)
}

func calculateUSIndependenceday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixeddateHoliday(year, time.July, 4, ObservedNearest, loc)
}

func calculateUSColumbusday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixedWeekdayHoliday(year, time.October, time.Monday, 2, loc)
}

func calculateUSVeteransday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixeddateHoliday(year, time.November, 11, ObservedExact, loc)
}

func calculateUSMemorialday(year int, loc *time.Location) (time.Month, int) {
	hd := time.Date(year,time.May,31,0,0,0,1,loc)
	for hd.Weekday() != time.Monday {
		hd = hd.AddDate(0,0,-1)
	}

	return hd.Month(), hd.Day()
}

/*
Canadian holiday calculation.
 */
func calculateCAIndependence(year int, loc *time.Location) (time.Month, int) {
	return calculateFixeddateHoliday(year, time.July, 1, ObservedMonday, loc)
}

func calculateCAFamilyday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixedWeekdayHoliday(year, time.February, time.Monday, 3, loc)
}

func calculateCAVictoriaDay(year int, loc *time.Location)  (time.Month, int) {
	hd := time.Date(year,time.May,24,0,0,0,1,loc)
	for hd.Weekday() != time.Monday {
		hd = hd.AddDate(0,0,-1)
	}

	return hd.Month(), hd.Day()
}

func calculateCAThanksgivingDay(year int, loc *time.Location)  (time.Month, int) {
	hd := time.Date(year,time.October,11,0,0,0,1,loc)
	td := hd
	if td.Weekday() != time.Monday {
		daydiff  := 1
		for {
			td = hd.AddDate(0,0,(0-daydiff))
			if td.Weekday() == time.Monday {
				break
			}
			td = hd.AddDate(0,0,daydiff)
			if td.Weekday() == time.Monday {
				break
			}
			daydiff++
		}
	}
	return td.Month(), td.Day()
}

/*
Common holiday calculation.
 */
func calculateGoodFriday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	//Go the the day before yesterday
	gf := easter.AddDate(0, 0, -2)
	return gf.Month(), gf.Day()
}

func calculateEasterMonday(year int, loc *time.Location) (time.Month, int) {
	easter := calculateEaster(year, loc)
	//Go the the day after Easter
	em := easter.AddDate(0, 0, +1)
	return em.Month(), em.Day()
}

func calculateLabourday(year int, loc *time.Location) (time.Month, int){
	return calculateFixedWeekdayHoliday(year, time.September, time.Monday, 1, loc)
}

func calculateChristmasday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixeddateHoliday(year, time.December, 25, ObservedExact, loc)
}

func calculateBoxingday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixeddateHoliday(year, time.December, 26, ObservedExact, loc)
}

func calculateNewyearday(year int, loc *time.Location) (time.Month, int) {
	return calculateFixeddateHoliday(year, time.January, 1, ObservedMonday,  loc)
}

func calculateEaster(year int, loc *time.Location) time.Time {
	// Meeus/Jones/Butcher algorithm
	y := year
	a := y % 19
	b := y / 100
	c := y % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451

	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
}


/*
Holiday utilities.
 */
func calculateFixedWeekdayHoliday(year int, month time.Month, day time.Weekday, offset int, loc *time.Location) (time.Month, int) {
	hd := time.Date(year,month,1,0,0,0,1,loc)
	for hd.Weekday() != day {
		hd = hd.AddDate(0,0,1)
	}

	if offset <2 {
		return hd.Month(), hd.Day()
	}

	tdays := 7 * (offset-1)
	hd = hd.AddDate(0,0,tdays)
	return hd.Month(), hd.Day()
}

// If the holiday is fixed and based observe rule i.e holiday will be moved to Monday or Friday if the holiday is in weekend
func calculateFixeddateHoliday(year int, month time.Month, date int, observerule ObservedRule, loc *time.Location) (time.Month, int){
	hd := time.Date(year,month,date,0,0,0,1,loc)

	if observerule == ObservedExact || !IsWeekend(hd) {
		return hd.Month(), hd.Day()
	}

	offset := 1
	if hd.Weekday() == time.Saturday {
		if observerule == ObservedNearest {
			offset -= 2
		}
		if observerule == ObservedMonday {
			offset++
		}

	}

	hd = hd.AddDate(0,0,offset)
	return hd.Month(), hd.Day()
}

// IsWeekend reports whether the given date falls on a weekend.
func IsWeekend(date time.Time) bool {
	day := date.Weekday()
	return day == time.Saturday || day == time.Sunday
}
