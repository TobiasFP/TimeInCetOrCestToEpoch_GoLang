package CetCest

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTimeConv(t *testing.T) {
	// Time conversion
	layout := "2006-01-02T15:04:00 MST"
	wintertime := "2018-03-01T07:45:00"
	wintertime = GetCETOrCEST(wintertime)
	summertime := "2018-04-01T07:45:00"
	summertime = GetCETOrCEST(summertime)

	// timeToConvert := "2018-03-31T07:45:00"
	// Convert timestamps
	// loc, _ := time.LoadLocation("Europe/Copenhagen")
	ti, err := time.Parse(layout, wintertime)
	if err != nil {
		fmt.Println(err)
		t.Error("Didnt Expect to receive: Expected: a time'")

	}

	if 1519886700 != ti.Unix() {
		fmt.Println(strconv.Itoa(int(ti.Unix()) - 1519886700))
		t.Error("Expected '1519886700' but got ", ti.Unix())
	}

	ti, err = time.Parse(layout, summertime)
	if err != nil {
		fmt.Println(err)
		t.Error("Didnt Expect to receive: Expected: a time'")
	}

	if 1522561500 != ti.Unix() {
		fmt.Println(strconv.Itoa(int(ti.Unix()) - 1522561500))
		t.Error("Expected '1522561500' but got ", ti.Unix())
	}
}

func TestGetCETOrCEST(t *testing.T) {

	wintertime := "2018-03-01T07:45:00"
	if GetCETOrCEST(wintertime) != "2018-03-01T07:45:00 CET" {
		t.Error("Expected '2018-03-01T07:45:00' to be CET, but got", GetCETOrCEST(wintertime))
	}

	summertime := "2018-03-28T07:45:00"
	if GetCETOrCEST(summertime) != "2018-03-28T07:45:00 CEST" {
		t.Error("Expected '2018-03-28T07:45:00' to be CEST, but got", GetCETOrCEST(summertime))
	}
}

func TestIsCETOrCEST(t *testing.T) {
	// Time conversion
	layout := "2006-01-02T15:04:00"
	wintertime := "2018-03-01T07:45:00"
	ti, _ := time.Parse(layout, wintertime)

	CET := IsCET(ti)
	if CET != true {
		t.Error("Expected 'true' but got ", CET)
	}

	summertime := "2018-03-28T07:45:00"

	ti, _ = time.Parse(layout, summertime)

	CEST := IsCET(ti)
	if CEST != false {
		t.Error("Expected 'true' but got ", CEST)
	}

	summertime = "2018-04-01T07:45:00"

	ti, _ = time.Parse(layout, summertime)

	CEST = IsCET(ti)
	if CEST != false {
		t.Error("Expected 'true' but got ", CEST)
	}

	wintertime = "2018-10-31T07:45:00"

	ti, _ = time.Parse(layout, wintertime)

	CET = IsCET(ti)
	if CET != true {
		t.Error("Expected 'true' but got ", CET)
	}

	wintertime = "2018-11-31T07:45:00"

	ti, _ = time.Parse(layout, wintertime)

	CET = IsCET(ti)
	if CET != true {
		t.Error("Expected 'true' but got ", CET)
	}
}

func TestGetLastSundayOfMonth(t *testing.T) {
	layout := "2006-01-02T15:04:00"
	wintertime := "2018-03-01T07:45:00"
	// summertime := "2018-04-01T07:45:00"
	ti, _ := time.Parse(layout, wintertime)
	lastSunday := getLastSundayOfMonth(ti)
	fmt.Println(getLastSundayOfMonth(ti))
	if lastSunday != 25 {
		t.Error("Expected '25' but got ", lastSunday)
	}

	wintertime = "2020-03-01T07:45:00"
	ti, _ = time.Parse(layout, wintertime)
	lastSunday = getLastSundayOfMonth(ti)
	fmt.Println(getLastSundayOfMonth(ti))
	if lastSunday != 29 {
		t.Error("Expected '29' but got ", lastSunday)
	}
}
