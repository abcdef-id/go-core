package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// LIMITATION JIKA FUNCTION FORMAT TERDAPAT VARIABLE GOLANG
func replace(in string) (has24HourNoZero bool, out string) {
	out = in
	has24HourNoZero = false
	var old string
	for _, ph := range Placeholder {
		out = strings.Replace(out, ph.find, ph.subst, -1)
		if ph.find == "H" {
			if old != out {
				has24HourNoZero = true
			}
		}
		old = out
	}
	return
}

func replaceFormat(in string) (out string) {
	out = in
	for _, wk := range Workaround {
		out = strings.Replace(out, wk.find, wk.subst, -1)
	}
	return
}

func clean(in string) (out string) {
	out = in
	for _, wk := range Workaround {
		i := strings.Index(out, wk.subst)
		if i > -1 {
			out = strings.Replace(out, wk.subst, wk.find, -1)
		}
	}
	return
}

// DateTimeFormat formats a date based on Microsoft Excel (TM) conventions
func DateTimeFormat(format string, date time.Time) (out string) {
	if format == "" {
		return strconv.FormatInt(date.Unix(), 10)
	}
	format = replaceFormat(format)
	has24HourNoZero, _format := replace(format)
	if has24HourNoZero && date.Hour() <= 12 {
		_format = strings.Replace(_format, "15", "3", -1)
	}
	out = date.Format(_format)
	out = clean(out)
	return
}

// DateTimeParse parses a value to a date based on Microsoft Excel (TM) formats
func DateTimeParse(format string, value string) (t time.Time, e error) {
	if format == "" {
		return t, fmt.Errorf("Empty DateTime Format")
	}
	_, _format := replace(format)
	return time.Parse(_format, value)
}

type p struct{ find, subst string }
type w struct{ find, subst string }

//Placeholder converter
var Placeholder = []p{
	{"HH", "15"},
	{"H", "15"},
	{"hh", "03"},
	{"h", "3"},
	{"mm", "04"},
	{"m", "4"},
	{"ss", "05"},
	{"s", "5"},
	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},
	{"M", "1"},
	{"A", "PM"},
	{"a", "pm"},
	{"ZZZZ", "-0700"},
	{"ZZZ", "MST"},
	{"ZZ", "Z07:00"},
	{"YYYY", "2006"},
	{"YY", "06"},
	{"DDDD", "Monday"},
	{"DDD", "Mon"},
	{"DD", "02"},
	{"D", "2"},
}

//Workaround Workaround
var Workaround = []w{
	{"January", "BBBBBB"},
	{"Jan", "BBBBB"},
	{"15", "BBBB"},
	{"01", "BBB"},
	{"1", "BB"},
	{"Mon", "B"},
	{"Monday", "CCCCCC"},
	{"__2", "CCCCC"},
	{"_2", "CCCC"},
	{"002", "CCC"},
	{"02", "CC"},
	{"2", "C"},
	{"03", "EEEEEE"},
	{"3", "EEEEE"},
	{"04", "EEEE"},
	{"4", "EEE"},
	{"05", "EE"},
	{"5", "E"},
	{"2006", "FFFFFF"},
	{"06", "FFFFF"},
	{"PM", "FFFF"},
	{"pm", "FFF"},
	{"MST", "FF"},
	{"Z070000", "F"},
	{"Z0700", "GGGGGG"},
	{"Z07", "GGGGG"},
	{"Z07,00,00", "GGGG"},
	{"Z07,00", "GGG"},
	{"-070000", "GG"},
	{"-0700", "G"},
	{"-07,00,00", "JJJJJJ"},
	{"-07,00", "JJJJJ"},
	{"-07", "JJJJ"},
	{".00", "JJJ"},
	{".0", "JJ"},
	{".99", "J"},
	{".9", "K"},
}
