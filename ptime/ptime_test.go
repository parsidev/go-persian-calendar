package ptime_test

import (
	"testing"
	"time"
	. "github.com/yaa110/go-persian-calendar/ptime"
	"fmt"
)

type pMonthName struct {
	month Month
	name string
}

type amPmName struct {
	ap AM_PM
	name string
}

type pdate struct {
	year int
	month Month
	day int
}

type gdate struct {
	year int
	month time.Month
	day int
}

type dateConversion struct {
	persian pdate
	gregorian gdate
}

var month_persian_names = []pMonthName {
	{Farvardin, "فروردین"},
	{Ordibehesht, "اردیبهشت"},
	{Khordad, "خرداد"},
	{Tir, "تیر"},
	{Mordad, "مرداد"},
	{Shahrivar, "شهریور"},
	{Mehr, "مهر"},
	{Aban, "آبان"},
	{Azar, "آذر"},
	{Dey, "دی"},
	{Bahman, "بهمن"},
	{Esfand, "اسفند"},
}

var month_dari_names = []pMonthName {
	{Hamal, "حمل"},
	{Sur, "ثور"},
	{Jauza, "جوزا"},
	{Saratan, "سرطان"},
	{Asad, "اسد"},
	{Sonboleh, "سنبله"},
	{Mizan, "میزان"},
	{Aqrab, "عقرب"},
	{Qos, "قوس"},
	{Jady, "جدی"},
	{Dolv, "دلو"},
	{Hut, "حوت"},
}

var am_pm_names = []amPmName {
	{AM, "قبل از ظهر"},
	{PM, "بعد از ظهر"},
}

var am_pm_snames = []amPmName {
	{AM, "ق.ظ"},
	{PM, "ب.ظ"},
}

var date_conversions = []dateConversion {
	{
		persian:    pdate{1383, Tir, 15},
		gregorian:  gdate{2004, time.July, 5},
	},
	{
		persian:    pdate{1394, Dey, 11},
		gregorian:  gdate{2016, time.January, 1},
	},
	{
		persian:    pdate{1394, Esfand, 9},
		gregorian:  gdate{2016, time.February, 28},
	},
	{
		persian:    pdate{1394, Esfand, 11},
		gregorian:  gdate{2016, time.March, 1},
	},
	{
		persian:    pdate{1394, Esfand, 29},
		gregorian:  gdate{2016, time.March, 19},
	},
	{
		persian:    pdate{1395, Farvardin, 1},
		gregorian:  gdate{2016, time.March, 20},
	},
	{
		persian:    pdate{1395, Farvardin, 2},
		gregorian:  gdate{2016, time.March, 21},
	},
	{
		persian:    pdate{1395, Farvardin, 3},
		gregorian:  gdate{2016, time.March, 22},
	},
	{
		persian:    pdate{1395, Dey, 11},
		gregorian:  gdate{2016, time.December, 31},
	},
}

func TestPersianMonthName(t *testing.T)  {
	for _, p := range month_persian_names {
		if p.month.String() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.month.String(),
			)
		}
	}
}

func TestDariMonthName(t *testing.T)  {
	for _, p := range month_dari_names {
		if p.month.Dari() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.month.String(),
			)
		}
	}
}

func TestAmPmName(t *testing.T)  {
	for _, p := range am_pm_names {
		if p.ap.String() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.ap.String(),
			)
		}
	}
}

func TestAmPmShortName(t *testing.T)  {
	for _, p := range am_pm_snames {
		if p.ap.Short() != p.name {
			t.Error(
				"Expected", p.name,
				"got", p.ap.Short(),
			)
		}
	}
}

func TestLocations(t *testing.T) {
	if (Iran.String() != "Asia/Tehran") {
		t.Error(
			"For", "Iran",
			"expected", "Asia/Tehran",
			"got", Iran.String(),
		)
	}

	if (Afghanistan.String() != "Asia/Kabul") {
		t.Error(
			"For", "Afghanistan",
			"expected", "Asia/Kabul",
			"got", Iran.String(),
		)
	}
}

func TestPersianToGregorian(t *testing.T) {
	for _, p := range date_conversions {
		gt := Date(p.persian.year, p.persian.month, p.persian.day, 11, 59, 59, 0, Iran).Time()

		if (gt.Year() != p.gregorian.year || gt.Month() != p.gregorian.month || gt.Day() != p.gregorian.day) {
			t.Error(
				"For", fmt.Sprintf("%d %s %d", p.persian.year, p.persian.month.String(), p.persian.day),
				"expected", fmt.Sprintf("%d %s %d", p.gregorian.year, p.gregorian.month.String(), p.gregorian.day),
				"got", fmt.Sprintf("%d %s %d", gt.Year(), gt.Month().String(), gt.Day()),
			)
		}
	}
}

func TestGregorianToPersian(t *testing.T) {
	for _, p := range date_conversions {
		pt := New(time.Date(p.gregorian.year, p.gregorian.month, p.gregorian.day, 11, 59, 59, 0, Iran))

		if (pt.Year() != p.persian.year || pt.Month() != p.persian.month || pt.Day() != p.persian.day) {
			t.Error(
				"For", fmt.Sprintf("%d %s %d", p.gregorian.year, p.gregorian.month.String(), p.gregorian.day),
				"expected", fmt.Sprintf("%d %s %d", p.persian.year, p.persian.month.String(), p.persian.day),
				"got", fmt.Sprintf("%d %s %d", pt.Year(), pt.Month().String(), pt.Day()),
			)
		}
	}
}

func TestToUnixTimeStamp(t *testing.T) {
	pu := Now(Iran).Unix()
	tu := time.Now().In(Iran).Unix()
	if (pu != tu) {
		t.Error(
			"Expected", tu,
			"got", pu,
		)
	}
}

func TestFromUnixTimeStamp(t *testing.T) {
	tu := time.Now().In(Iran).Unix()
	now := Now(Iran)

	fu := Unix(tu, int64(now.Nanosecond()), Iran)

	if (fu.String() != now.String()) {
		t.Error(
			"Expected", now.String(),
			"got", fu.String(),
		)
	}
}