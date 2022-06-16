package entity

import (
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	t.Run("UTC", func(t *testing.T) {
		if TimeStamp().Location() != time.UTC {
			t.Fatal("timestamp zone must be utc")
		}
	})
	t.Run("Past", func(t *testing.T) {
		if TimeStamp().After(time.Now().Add(time.Second)) {
			t.Fatal("timestamp should be in the past from now")
		}
	})
	t.Run("JSON", func(t *testing.T) {
		t1 := TimeStamp().Add(time.Nanosecond * 123456)

		if b, err := t1.MarshalJSON(); err != nil {
			t.Fatal(err)
		} else {
			t.Logf("JSON: %s", b)
		}
	})
	t.Run("UnixMicro", func(t *testing.T) {
		t1 := time.Date(-3000, 1, 1, 1, 1, 1, 0, time.UTC)
		t2 := TimeStamp().Add(time.Nanosecond * 123456)
		t3 := time.Date(3000, 1, 1, 1, 1, 1, 0, time.UTC)

		ms1 := t1.UnixMilli()
		ms2 := t2.UnixMilli()
		ms3 := t3.UnixMilli()

		m1 := t1.UnixMicro()
		m2 := t2.UnixMicro()
		m3 := t3.UnixMicro()

		t.Logf("MS1: %20d", ms1)
		t.Logf("MS2: %20d", ms2)
		t.Logf("MS3: %20d", ms3)

		t.Logf("U1: %20d", m1)
		t.Logf("U2: %20d", m2)
		t.Logf("U3: %20d", m3)

		i1, i2, i3 := 1e18-m1, 1e18-m2, 1e18-m3

		t.Logf("ZZ: %20d", 9223372036854775807)
		t.Logf("I1: %20d", i1)
		t.Logf("I2: %20d", i2)
		t.Logf("I3: %20d", i3)

		t.Logf("T1: %20d", 1e18-i1)
		t.Logf("T2: %20d", 1e18-i2)
		t.Logf("T3: %20d", 1e18-i3)

		t.Logf("D1: %s", time.UnixMicro(1e18-i1).String())
		t.Logf("D2: %s", time.UnixMicro(1e18-i2).String())
		t.Logf("D3: %s", time.UnixMicro(1e18-i3).String())
	})
}

func TestTimePointer(t *testing.T) {
	result := TimePointer()

	if result == nil {
		t.Fatal("result must not be nil")
	}

	if result.Location() != time.UTC {
		t.Fatal("timestamp zone must be utc")
	}

	if result.After(time.Now().Add(time.Second)) {
		t.Fatal("timestamp should be in the past from now")
	}
}

func TestSeconds(t *testing.T) {
	result := Seconds(23)

	if result != 23*time.Second {
		t.Error("must be 23 seconds")
	}
}

func TestYesterday(t *testing.T) {
	now := time.Now()
	result := Yesterday()

	t.Logf("yesterday: %s", result)

	if result.After(now) {
		t.Error("yesterday is not before now")
	}

	if !result.Before(now) {
		t.Error("yesterday is before now")
	}
}
