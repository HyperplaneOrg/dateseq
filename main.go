/*
A simple golang date/datetime sequence printing command-line utility. 
This is really a "hello world" toy go program/project but one might 
use it in some shell scipting. It can be useful for creating filename
patterns that contain dates or datetimes.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
   "regexp"
   "github.com/HyperplaneOrg/go-strftime"
)
   
const inpDateForm string = "20060102"
var inpDateFormRx *regexp.Regexp = regexp.MustCompile(`^[0-9]{8}$`)
const inpDateTimeForm string = "2006010215"
var inpDateTimeFormRx *regexp.Regexp = regexp.MustCompile(`^[0-9]{10}$`)

func printUsage() {
	var prgusage = `
  Usage: datesets [ OPTIONS ] date [ date_end ] 
	   Where date is of the format YYYYmmdd[HH], where YYYY=year, 
      mm=month, dd=day, HH=hour(24) e.g. 20060102 or 2006010215
	   -s <N> :: The date-time sequence stride, e.g. every N days if input format 
                is YYYYmmdd or every N hours if input format is YYYYmmddHH
	   -f <format-string> :: The traditional date-time/strftime print format string, e.g. %%Y%%m%%d`+"\n\n"
	fmt.Fprintf(os.Stderr, prgusage)
}

func parseDateStr(s string) (time.Time, bool) {
   var e error
   var d time.Time 
   var hrly bool = false 
   if inpDateFormRx.MatchString(s) {
	   d, e = time.Parse(inpDateForm, s)
   } else if inpDateTimeFormRx.MatchString(s) {
	   d, e = time.Parse(inpDateTimeForm, s)
      hrly = true
   } else {
		fmt.Fprintf(os.Stderr, "Bad date or datetime given, try -h for help.\n")
		os.Exit(1)
		return d, hrly 
   }
	if e != nil {
	   fmt.Fprintf(os.Stderr, "ERROR : "+e.Error()+"\n")
      os.Exit(1)
		return d, hrly 
	} else {
		return d, hrly 
	}
}

func printDateStr(dt time.Time, layout string) {
   fmt.Println(strftime.Format(layout, dt))
}

func printDateSeq(startd time.Time, endd time.Time, hourly bool, stride int, layout string) error {
   if hourly == false {
      for d := startd; d.Before(endd) || d.Equal(endd); d = d.AddDate(0, 0, stride) {
         printDateStr(d, layout)
      }
   } else {
      dur := time.Duration(3600000000000 * stride) 
      for d := startd; d.Before(endd) || d.Equal(endd); d = d.Add(dur) {
         printDateStr(d, layout)
      }
   }
   return nil
}

func main() {
	var ed, bd time.Time 
	flag.Usage = printUsage
   outfrmt := flag.String("f", "%Y%m%d", "The date-time print format.")
	dstride := flag.Int("s", 1, "The date-time sequence stride, e.g. every N days, or N hours")
	hrly := false
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "No date(s) given, try -h for help.\n")
		os.Exit(0)
	}
	bd, hrly = parseDateStr(flag.Arg(0))

	if flag.NArg() >= 2 {
	   ed, _ = parseDateStr(flag.Arg(1))
   } else {
	   ed = bd
   }
   if ed.Before(bd) {
      ed, bd = bd, ed
   }

   if *dstride < 0 {
      *dstride = -(*dstride)
   } else if *dstride == 0 {
      *dstride = 1
   }

   printDateSeq(bd, ed, hrly, *dstride, *outfrmt)
} // main 

