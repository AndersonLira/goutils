package str

import (
	"fmt"
	"testing"
	"time"
	"github.com/andersonlira/goutils/str"
)

var now = time.Now()
var year,month,day = now.Date()
var dayString = fmt.Sprintf("%02d", day)
var monthString = fmt.Sprintf("%02d", int(month))
var monthName = month.String()[:3]
var yearString = fmt.Sprintf("%04d", year)
var yearShort =	yearString[2:len(yearString)]
var hour = now.Hour()
var hourString = fmt.Sprintf("%02d", hour)
var minute = now.Minute()
var minuteString = fmt.Sprintf("%02d", minute)
var second = now.Second()
var secondString = fmt.Sprintf("%02d", second)




func TestDay(t *testing.T) {
	f := str.FmtDate("dd",now)
	if  f != dayString {
		t.Errorf("Day was wrong. It should be %s but %s", dayString , f )
	}
}

func TestMonth(t *testing.T) {
	f := str.FmtDate("MM",now)
	if  f != monthString {
		t.Errorf("Month was wrong. It should be %s but %s", monthString , f )
	}
}

func TestYear(t *testing.T) {
	f := str.FmtDate("yyyy",now)
	if  f != yearString {
		t.Errorf("Year was wrong. It should be %s but %s", yearString , f )
	}

	f = str.FmtDate("yy",now)
	if  f != yearShort {
		t.Errorf("Year was wrong. It should be %s but %s", yearShort , f )
	}

}

func TestMinute(t *testing.T) {
	f := str.FmtDate("mm",now)
	if  f != minuteString {
		t.Errorf("minute was wrong. It should be %s but %s", minuteString , f )
	}
}

func TestSecond(t *testing.T) {
	f := str.FmtDate("ss",now)
	if  f != secondString {
		t.Errorf("second was wrong. It should be %s but %s", secondString , f )
	}
}

func TestCombinedFormats(t *testing.T) {
	fmtA := fmt.Sprintf("%s/%s/%s",dayString,monthString,yearString)
	f := str.FmtDate("dd/MM/yyyy",now)

	if f != fmtA {
		t.Errorf("Pattern was wrong. It should be %s but %s", fmtA , f )
	}

	fmtA = fmt.Sprintf("%s-%s-%s",dayString,monthString,yearShort)
	f = str.FmtDate("dd-MM-yy",now)

	if f != fmtA {
		t.Errorf("Pattern was wrong. It should be %s but %s", fmtA , f )
	}

	fmtA = fmt.Sprintf("%s:%s:%s",hourString,minuteString,secondString)
	f = str.FmtDate("HH:mm:ss",now)

	if f != fmtA {
		t.Errorf("Pattern was wrong. It should be %s but %s", fmtA , f )
	}

	fmtA = fmt.Sprintf("%s/%s/%s %s:%s:%s",dayString,monthName,yearString,hourString,minuteString,secondString)
	f = str.FmtDate("dd/MMM/yyyy HH:mm:ss",now)

	if f != fmtA {
		t.Errorf("Pattern was wrong. It should be %s but %s", fmtA , f )
	}

}