/*
This is an simple example of golang testing; for some routines
in the date/datetime sequence printing command-line utility.
*/

package main

import (
	"testing"
)

func TestPraseDate(t *testing.T) {
	var dt string = "20010101"
	var dth string = "2022010101"
	tmd, hd := parseDateStr(dt)
	if hd != false {
		t.Errorf("parseDateStr returned h = %t, hourly format instead of daily", hd)
	}
	d := tmd.Day()
	if d != 1 {
		t.Errorf("parseDateStr returned a day of %d, which != 1", d)
	}
	tmh, hh := parseDateStr(dth)
	if hh != true {
		t.Errorf("parseDateStr returned h = %t, daily format instead of hourly", hh)
	}
	h := tmh.Hour()
	if h != 1 {
		t.Errorf("parseDateStr returned an hour of %d, which != 1", h)
	}

}
