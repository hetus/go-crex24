package util

import "time"

const ISO8601 string = "2006-01-02T15:04:05-0700"

func Date() (t time.Time) {
	t = time.Now().UTC()
	return
}

func DateFromString(d string) (t time.Time, err error) {
	t, err = time.Parse(ISO8601, d)
	t = t.UTC()
	return
}

func DateToString(t time.Time) (d string) {
	d = t.Format(ISO8601)
	return
}
