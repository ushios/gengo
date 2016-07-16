package gengo

import (
	"testing"
	"time"
)

func AtTest(t *testing.T) {
	test := func(nowStr string, g Gengo) {
		now, err := time.Parse(DateTimeFormat, nowStr)
		if err != nil {
			t.Fatal(err)
		}

		gengo := At(now)
		if gengo != g {
			t.Errorf("gengo expected (%s) but (%s)", g, gengo)
		}
	}

	test("1989-01-08 00:00:00 +0900", Heisei)
	test("1989-01-07 23:00:00 +0800", Heisei)
	test("1989-01-07 23:59:59 +0900", Showa)
	test("1926-12-25 00:00:00 +0900", Showa)
	test("1926-12-24 23:59:59 +0900", Taisho)
	test("1912-07-30 00:00:00 +0900", Taisho)
	test("1912-07-29 23:59:59 +0900", Meiji)
}
