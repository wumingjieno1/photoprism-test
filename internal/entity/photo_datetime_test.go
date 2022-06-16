package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPhoto_TrustedTime(t *testing.T) {
	t.Run("MissingTakenAt", func(t *testing.T) {
		m := Photo{ID: 1, TakenAt: time.Time{}, TakenAtLocal: TimeStamp(), TakenSrc: SrcMeta, TimeZone: "Europe/Berlin"}
		assert.False(t, m.TrustedTime())
	})
	t.Run("MissingTakenAtLocal", func(t *testing.T) {
		m := Photo{ID: 1, TakenAt: TimeStamp(), TakenAtLocal: time.Time{}, TakenSrc: SrcMeta, TimeZone: "Europe/Berlin"}
		assert.False(t, m.TrustedTime())
	})
	t.Run("MissingTimeZone", func(t *testing.T) {
		n := TimeStamp()
		m := Photo{ID: 1, TakenAt: n, TakenAtLocal: n, TakenSrc: SrcMeta, TimeZone: ""}
		assert.False(t, m.TrustedTime())
	})
	t.Run("SrcAuto", func(t *testing.T) {
		n := TimeStamp()
		m := Photo{ID: 1, TakenAt: n, TakenAtLocal: n, TakenSrc: SrcAuto, TimeZone: "Europe/Berlin"}
		assert.False(t, m.TrustedTime())
	})
	t.Run("SrcEstimate", func(t *testing.T) {
		n := TimeStamp()
		m := Photo{ID: 1, TakenAt: n, TakenAtLocal: n, TakenSrc: SrcEstimate, TimeZone: "Europe/Berlin"}
		assert.False(t, m.TrustedTime())
	})
	t.Run("SrcMeta", func(t *testing.T) {
		n := TimeStamp()
		m := Photo{ID: 1, TakenAt: n, TakenAtLocal: n, TakenSrc: SrcMeta, TimeZone: "Europe/Berlin"}
		assert.True(t, m.TrustedTime())
	})
}

func TestPhoto_SetTakenAt(t *testing.T) {
	t.Run("empty taken", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		m.SetTakenAt(time.Time{}, time.Time{}, "", SrcManual)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
	})
	t.Run("taken not from the same source", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		m.SetTakenAt(time.Date(2019, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2019, 12, 11, 9, 7, 18, 0, time.UTC), "", SrcAuto)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
	})
	t.Run("from name", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")
		m.TimeZone = ""
		m.TakenSrc = SrcAuto

		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)
		assert.Equal(t, "", m.TimeZone)
		assert.Equal(t, SrcAuto, m.TakenSrc)

		m.SetTakenAt(time.Date(2011, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2019, 11, 11, 10, 7, 18, 0, time.UTC), "America/New_York", SrcName)

		assert.Equal(t, "", m.TimeZone)
		assert.Equal(t, SrcName, m.TakenSrc)

		assert.Equal(t, time.Date(2011, 12, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2019, 11, 11, 10, 7, 18, 0, time.UTC), m.TakenAtLocal)
	})
	t.Run("success", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)

		m.SetTakenAt(time.Date(2019, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2019, 12, 11, 10, 7, 18, 0, time.UTC), "", SrcMeta)

		assert.Equal(t, time.Date(2019, 12, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2019, 12, 11, 10, 7, 18, 0, time.UTC), m.TakenAtLocal)
	})
	t.Run("fallback", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")
		assert.Equal(t, time.Date(2013, time.November, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2013, time.November, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)

		t.Logf("SRC, ZONE, UTC, LOCAL: %s / %s / %s /%s", m.TakenSrc, m.TimeZone, m.TakenAt, m.TakenAtLocal)

		m.SetTakenAt(time.Date(2019, time.December, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2019, time.December, 11, 10, 7, 18, 0, time.UTC), "", SrcAuto)

		t.Logf("SRC, ZONE, UTC, LOCAL: %s / %s / %s /%s", m.TakenSrc, m.TimeZone, m.TakenAt, m.TakenAtLocal)

		assert.Equal(t, time.Date(2013, time.November, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2013, time.November, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)

		newTime := time.Date(2013, time.November, 11, 9, 7, 18, 0, time.UTC)

		expected := time.Date(2013, time.November, 11, 8, 7, 18, 0, time.UTC)

		m.TimeZone = "Europe/Berlin"

		m.SetTakenAt(newTime, newTime, "", SrcName)

		assert.Equal(t, expected, m.TakenAt)
		assert.Equal(t, m.GetTakenAtLocal(), m.TakenAtLocal)
	})
	t.Run("time zone", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")

		zone := "Europe/Berlin"

		loc, _ := time.LoadLocation(zone)

		newTime := time.Date(2013, 11, 11, 9, 7, 18, 0, loc)

		m.SetTakenAt(newTime, newTime, zone, SrcName)

		assert.Equal(t, newTime.UTC(), m.TakenAt)
		assert.Equal(t, newTime, m.TakenAtLocal)
	})
	t.Run("time > max year", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)
		m.SetTakenAt(time.Date(2123, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2123, 12, 11, 10, 7, 18, 0, time.UTC), "", SrcManual)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)
	})
	t.Run("success with empty takenAtLocal", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo15")
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)
		m.SetTakenAt(time.Date(2019, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Time{}, "test", SrcXmp)
		assert.Equal(t, time.Date(2019, 12, 11, 9, 7, 18, 0, time.UTC), m.TakenAt)
		assert.Equal(t, time.Date(2019, 12, 11, 9, 7, 18, 0, time.UTC), m.TakenAtLocal)
	})
	t.Run("don't update older date", func(t *testing.T) {
		photo := &Photo{TakenAt: time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC)}
		photo.SetTakenAt(time.Date(2014, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2014, 12, 11, 10, 7, 18, 0, time.UTC), "", SrcAuto)
		assert.Equal(t, time.Date(2013, 11, 11, 9, 7, 18, 0, time.UTC), photo.TakenAt)
	})
	t.Run("set local time from utc", func(t *testing.T) {
		photo := &Photo{TakenAt: time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC), TimeZone: "Europe/Berlin"}
		photo.SetTakenAt(time.Date(2014, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2014, 12, 11, 10, 7, 18, 0, time.UTC), time.UTC.String(), SrcManual)
		assert.Equal(t, time.Date(2014, 12, 11, 9, 7, 18, 0, time.UTC), photo.TakenAt)
		assert.Equal(t, time.Date(2014, 12, 11, 10, 07, 18, 0, time.UTC), photo.TakenAtLocal)
	})
	t.Run("local is UTC", func(t *testing.T) {
		photo := &Photo{TakenAt: time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC), TimeZone: time.UTC.String()}
		photo.SetTakenAt(time.Date(2014, 12, 11, 9, 7, 18, 0, time.UTC),
			time.Date(2014, 12, 11, 10, 7, 18, 0, time.UTC), "", SrcManual)
		assert.Equal(t, time.Date(2014, 12, 11, 9, 7, 18, 0, time.UTC), photo.TakenAt)
		assert.Equal(t, time.Date(2014, 12, 11, 9, 07, 18, 0, time.UTC), photo.TakenAtLocal)
	})
}

func TestPhoto_UpdateTimeZone(t *testing.T) {
	t.Run("PhotoTimeZone", func(t *testing.T) {
		m := PhotoFixtures.Get("PhotoTimeZone")

		takenLocal := time.Date(2015, time.May, 17, 23, 2, 46, 0, time.UTC)
		takenJerusalemUtc := time.Date(2015, time.May, 17, 20, 2, 46, 0, time.UTC)
		takenShanghaiUtc := time.Date(2015, time.May, 17, 15, 2, 46, 0, time.UTC)

		assert.Equal(t, "", m.TimeZone)
		assert.Equal(t, takenLocal, m.TakenAt)
		assert.Equal(t, takenLocal, m.TakenAtLocal)

		zone1 := "Asia/Jerusalem"

		m.UpdateTimeZone(zone1)

		assert.Equal(t, zone1, m.TimeZone)
		assert.Equal(t, takenJerusalemUtc, m.TakenAt)
		assert.Equal(t, takenLocal, m.TakenAtLocal)

		m.UpdateTimeZone(zone1)

		assert.Equal(t, zone1, m.TimeZone)
		assert.Equal(t, takenJerusalemUtc, m.TakenAt)
		assert.Equal(t, takenLocal, m.TakenAtLocal)

		zone2 := "Asia/Shanghai"

		m.UpdateTimeZone(zone2)

		assert.Equal(t, zone2, m.TimeZone)
		assert.Equal(t, takenShanghaiUtc, m.TakenAt)
		assert.Equal(t, takenLocal, m.TakenAtLocal)

		zone3 := "UTC"

		m.UpdateTimeZone(zone3)

		assert.Equal(t, zone2, m.TimeZone)
		assert.Equal(t, takenShanghaiUtc, m.TakenAt)
		assert.Equal(t, takenLocal, m.TakenAtLocal)
	})
	t.Run("VideoTimeZone", func(t *testing.T) {
		m := PhotoFixtures.Get("VideoTimeZone")

		takenUtc := time.Date(2015, 5, 17, 17, 48, 46, 0, time.UTC)
		takenJerusalem := time.Date(2015, time.May, 17, 20, 48, 46, 0, time.UTC)
		takenShanghaiUtc := time.Date(2015, time.May, 17, 12, 48, 46, 0, time.UTC)

		assert.Equal(t, "UTC", m.TimeZone)
		assert.Equal(t, takenUtc, m.TakenAt)
		assert.Equal(t, takenUtc, m.TakenAtLocal)

		zone1 := "Asia/Jerusalem"

		m.UpdateTimeZone(zone1)

		assert.Equal(t, zone1, m.TimeZone)
		assert.Equal(t, takenUtc, m.TakenAt)
		assert.Equal(t, takenJerusalem, m.TakenAtLocal)

		m.UpdateTimeZone(zone1)

		assert.Equal(t, zone1, m.TimeZone)
		assert.Equal(t, takenUtc, m.TakenAt)
		assert.Equal(t, takenJerusalem, m.TakenAtLocal)

		zone2 := "Asia/Shanghai"

		m.UpdateTimeZone(zone2)

		assert.Equal(t, zone2, m.TimeZone)
		assert.Equal(t, takenShanghaiUtc, m.TakenAt)
		assert.Equal(t, takenJerusalem, m.TakenAtLocal)

		zone3 := "UTC"

		m.UpdateTimeZone(zone3)

		assert.Equal(t, zone2, m.TimeZone)
		assert.Equal(t, takenShanghaiUtc, m.TakenAt)
		assert.Equal(t, takenJerusalem, m.TakenAtLocal)
	})
	t.Run("UTC", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo12")
		m.TimeZone = "UTC"

		zone := "Europe/Berlin"

		takenAt := m.TakenAt
		takenAtLocal := m.TakenAtLocal

		assert.Equal(t, takenAt, m.TakenAt)
		assert.Equal(t, takenAtLocal, m.TakenAtLocal)

		m.UpdateTimeZone(zone)

		assert.Equal(t, takenAt, m.TakenAt)
		assert.Equal(t, m.GetTakenAtLocal(), m.TakenAtLocal)
	})

	t.Run("Europe/Berlin", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo12")

		zone := "Europe/Berlin"

		takenAt := m.TakenAt
		takenAtLocal := m.TakenAtLocal

		assert.Equal(t, takenAt, m.TakenAt)
		assert.Equal(t, takenAtLocal, m.TakenAtLocal)
		assert.Equal(t, "", m.TimeZone)

		m.UpdateTimeZone(zone)

		assert.Equal(t, m.GetTakenAt(), m.TakenAt)
		assert.Equal(t, takenAtLocal, m.TakenAtLocal)
	})

	t.Run("America/New_York", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo12")
		m.TimeZone = "Europe/Berlin"
		m.TakenAt = m.GetTakenAt()

		zone := "America/New_York"

		takenAt := m.TakenAt
		takenAtLocal := m.TakenAtLocal

		assert.Equal(t, takenAt, m.TakenAt)
		assert.Equal(t, takenAtLocal, m.TakenAtLocal)
		assert.Equal(t, "Europe/Berlin", m.TimeZone)

		m.UpdateTimeZone(zone)

		assert.Equal(t, m.GetTakenAt(), m.TakenAt)
		assert.Equal(t, takenAtLocal, m.TakenAtLocal)
	})

	t.Run("manual", func(t *testing.T) {
		m := PhotoFixtures.Get("Photo12")
		m.TimeZone = "Europe/Berlin"
		m.TakenAt = m.GetTakenAt()
		m.TakenSrc = SrcManual

		zone := "America/New_York"

		takenAt := m.TakenAt
		takenAtLocal := m.TakenAtLocal

		assert.Equal(t, takenAt, m.TakenAt)
		assert.Equal(t, takenAtLocal, m.TakenAtLocal)
		assert.Equal(t, "Europe/Berlin", m.TimeZone)

		m.UpdateTimeZone(zone)

		assert.Equal(t, takenAt, m.TakenAt)
		assert.Equal(t, takenAtLocal, m.TakenAtLocal)
		assert.Equal(t, "Europe/Berlin", m.TimeZone)
	})
	t.Run("zone = UTC", func(t *testing.T) {
		photo := &Photo{TakenAt: time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC), TimeZone: "Europe/Berlin"}
		photo.UpdateTimeZone("")
		assert.Equal(t, time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC), photo.TakenAt)
		assert.Equal(t, "Europe/Berlin", photo.TimeZone)
	})
}
