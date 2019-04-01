package gengo

import (
	"fmt"
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

	test("2019-05-02 00:00:00 +0900", Reiwa)
	test("2019-05-01 00:00:00 +0900", Reiwa)
	test("2019-04-30 23:59:59 +0900", Heisei)
	test("1989-01-08 00:00:00 +0900", Heisei)
	test("1989-01-07 23:00:00 +0800", Heisei)
	test("1989-01-07 23:59:59 +0900", Showa)
	test("1926-12-25 00:00:00 +0900", Showa)
	test("1926-12-24 23:59:59 +0900", Taisho)
	test("1912-07-30 00:00:00 +0900", Taisho)
	test("1912-07-29 23:59:59 +0900", Meiji)
}

func NowTest(t *testing.T) {
	g := Now()

	if g != Heisei {
		t.Errorf("expected(%s) but (%s) now", Heisei, g)
	}
}

func ExampleNow() {
	g := Now()

	fmt.Println(g)
	// Output:
	// 平成
}

func ExampleAt() {
	t, _ := time.Parse(DateTimeFormat, "1912-07-29 23:59:59 +0900")
	g := At(t)

	fmt.Println(g)
	// Output:
	// 明治
}
