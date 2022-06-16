package meta

import (
	"testing"
	"time"

	"github.com/photoprism/photoprism/pkg/video"

	"github.com/photoprism/photoprism/pkg/projection"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	t.Run("iphone-mov.json", func(t *testing.T) {
		data, err := JSON("testdata/iphone-mov.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, "20170323-083538-Berlin-Zoologischer-Garten-2017-2u4.mov", data.FileName)
		assert.Equal(t, CodecAvc1, data.Codec)
		assert.Equal(t, "3s", data.Duration.String())
		assert.Equal(t, "2018-09-08 19:20:14 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2018-09-08 17:20:14 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, 0, data.TakenNs)
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1080, data.ActualWidth())
		assert.Equal(t, 1920, data.ActualHeight())
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, float32(52.4587), data.Lat)
		assert.Equal(t, float32(13.4593), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone SE", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("gopher-telegram.json", func(t *testing.T) {
		data, err := JSON("testdata/gopher-telegram.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecAvc1, data.Codec)
		assert.Equal(t, "2s", data.Duration.String())
		assert.Equal(t, "2020-05-11 14:18:35 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2020-05-11 14:18:35 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, 0, data.TakenNs)
		assert.Equal(t, time.UTC.String(), data.TimeZone)
		assert.Equal(t, 270, data.Width)
		assert.Equal(t, 480, data.Height)
		assert.Equal(t, 270, data.ActualWidth())
		assert.Equal(t, 480, data.ActualHeight())
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("gopher-original.json", func(t *testing.T) {
		data, err := JSON("testdata/gopher-original.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecAvc1, data.Codec)
		assert.Equal(t, "2s", data.Duration.String())
		assert.Equal(t, "2020-05-11 16:16:48 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2020-05-11 14:16:48 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1080, data.ActualWidth())
		assert.Equal(t, 1920, data.ActualHeight())
		assert.Equal(t, float32(0.56), data.AspectRatio())
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, float32(52.4596), data.Lat)
		assert.Equal(t, float32(13.3218), data.Lng)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("berlin-landscape.json", func(t *testing.T) {
		data, err := JSON("testdata/berlin-landscape.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecAvc1, data.Codec)
		assert.Equal(t, "4s", data.Duration.String())
		assert.Equal(t, "2020-05-14 13:34:41 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2020-05-14 11:34:41 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, 0, data.TakenNs)
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(52.4649), data.Lat)
		assert.Equal(t, float32(13.3148), data.Lng)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("mp4.json", func(t *testing.T) {
		data, err := JSON("testdata/mp4.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecAvc1, data.Codec)
		assert.Equal(t, "4m25s", data.Duration.String())
		assert.Equal(t, "2019-11-23 13:51:49 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, 848, data.Width)
		assert.Equal(t, 480, data.Height)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("photoshop.json", func(t *testing.T) {
		data, err := JSON("testdata/photoshop.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)
		assert.Equal(t, "photoshop.xmp", data.FileName)
		assert.Equal(t, CodecXMP, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, float32(52.45969), data.Lat)
		assert.Equal(t, float32(13.321831), data.Lng)
		assert.Equal(t, "2020-01-01 16:28:23 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "2020-01-01 17:28:23 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, 899614000, data.TakenNs)
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, "Night Shift / Berlin / 2020", data.Title)
		assert.Equal(t, "Michael Mayer", data.Artist)
		assert.Equal(t, "Example file for development", data.Description)
		assert.Equal(t, "This is an (edited) legal notice", data.Copyright)
		assert.Equal(t, "HUAWEI", data.CameraMake)
		assert.Equal(t, "ELE-L29", data.CameraModel)
		assert.Equal(t, "HUAWEI P30 Rear Main Camera", data.LensModel)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("canon_eos_6d.json", func(t *testing.T) {
		data, err := JSON("testdata/canon_eos_6d.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecXMP, data.Codec)
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "Canon EOS 6D", data.CameraModel)
		assert.Equal(t, "EF24-105mm f/4L IS USM", data.LensModel)
		assert.Equal(t, 0, data.TakenNs)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("gps-2000.json", func(t *testing.T) {
		data, err := JSON("testdata/gps-2000.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecUnknown, data.Codec)
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("ladybug.json", func(t *testing.T) {
		data, err := JSON("testdata/ladybug.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecUnknown, data.Codec)
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("iphone_7.json", func(t *testing.T) {
		data, err := JSON("testdata/iphone_7.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, CodecHeic, data.Codec)
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 7", data.CameraModel)
		assert.Equal(t, "Apple", data.LensMake)
		assert.Equal(t, "iPhone 7 back camera 3.99mm f/1.8", data.LensModel)
	})

	t.Run("uuid-original.json", func(t *testing.T) {
		data, err := JSON("testdata/uuid-original.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, "9bafc58c-6c82-4e66-a45f-c13f915f99c5", data.DocumentID)
		assert.Equal(t, "", data.InstanceID)
		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, "2018-12-06 12:32:26 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2018-12-06 11:32:26 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, 3024, data.Width)
		assert.Equal(t, 4032, data.Height)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(48.300003), data.Lat)
		assert.Equal(t, float32(8.929067), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone SE", data.CameraModel)
		assert.Equal(t, "iPhone SE back camera 4.15mm f/2.2", data.LensModel)
	})

	t.Run("uuid-copy.json", func(t *testing.T) {
		data, err := JSON("testdata/uuid-copy.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, "", data.DocumentID)
		assert.Equal(t, "dafbfeb8-a129-4e7c-9cf0-e7996a701cdb", data.InstanceID)
		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, "2018-12-06 12:32:26 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2018-12-06 11:32:26 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, 1024, data.Width)
		assert.Equal(t, 1365, data.Height)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(48.300003), data.Lat)
		assert.Equal(t, float32(8.929067), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone SE", data.CameraModel)
		assert.Equal(t, "iPhone SE back camera 4.15mm f/2.2", data.LensModel)
	})

	t.Run("uuid-imagemagick.json", func(t *testing.T) {
		data, err := JSON("testdata/uuid-imagemagick.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("DATA: %+v", data)

		assert.Equal(t, "9bafc58c-6c82-4e66-a45f-c13f915f99c5", data.DocumentID)
		assert.Equal(t, "", data.InstanceID)
		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, "2018-12-06 12:32:26 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2018-12-06 11:32:26 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, 1125, data.Width)
		assert.Equal(t, 1500, data.Height)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(48.300003), data.Lat)
		assert.Equal(t, float32(8.929067), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone SE", data.CameraModel)
		assert.Equal(t, "iPhone SE back camera 4.15mm f/2.2", data.LensModel)
	})

	t.Run("orientation.json", func(t *testing.T) {
		data, err := JSON("testdata/orientation.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 326, data.Width)
		assert.Equal(t, 184, data.Height)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("gphotos-1.json", func(t *testing.T) {
		data, err := JSON("testdata/gphotos-1.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "2015-12-06 16:18:30 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2015-12-06 15:18:30 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, float32(52.508522), data.Lat)
		assert.Equal(t, float32(13.443206), data.Lng)
		assert.Equal(t, 40, data.Altitude)
		assert.Equal(t, 0, data.Views)

		assert.Equal(t, "", data.DocumentID)
		assert.Equal(t, "", data.InstanceID)
		assert.Equal(t, CodecUnknown, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, 0, data.Width)
		assert.Equal(t, 0, data.Height)
		assert.Equal(t, 0, data.Orientation)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("gphotos-2.json", func(t *testing.T) {
		data, err := JSON("testdata/gphotos-2.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "", data.Title)
		assert.Equal(t, "A photo description", data.Description)
		assert.Equal(t, "2019-05-18 12:06:45 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2019-05-18 10:06:45 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, float32(52.510796), data.Lat)
		assert.Equal(t, float32(13.456387), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, 1118, data.Views)
	})

	t.Run("gphotos-3.json", func(t *testing.T) {
		data, err := JSON("testdata/gphotos-3.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Bei den Landungsbrücken", data.Title)
		assert.Equal(t, "In Hamburg", data.Description)
		assert.Equal(t, "2011-11-07 21:34:34 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2011-11-07 21:34:34 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "", data.TimeZone)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, 177, data.Views)
	})

	t.Run("gphotos-4.json", func(t *testing.T) {
		data, err := JSON("testdata/gphotos-4.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "2012-12-11 00:07:15 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2012-12-10 23:07:15 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, float32(52.49967), data.Lat)
		assert.Equal(t, float32(13.422334), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, 0, data.Views)
	})

	t.Run("gphotos-album.json", func(t *testing.T) {
		data, err := JSON("testdata/gphotos-album.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, data.TakenAtLocal.IsZero())
		assert.True(t, data.TakenAt.IsZero())
		assert.Equal(t, 0, data.Views)

		if len(data.Albums) == 1 {
			assert.Equal(t, "iPhone", data.Albums[0])
		} else {
			assert.Len(t, data.Albums, 1)
		}
	})

	t.Run("panorama360.json", func(t *testing.T) {
		data, err := JSON("testdata/panorama360.json", "panorama360.jpg")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2020-05-24T08:55:21Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2020-05-24T11:55:21Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "panorama", data.Keywords.String())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 3600, data.Height)
		assert.Equal(t, 7200, data.Width)
		assert.Equal(t, float32(59.84083), data.Lat)
		assert.Equal(t, float32(30.51), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "1/1250", data.Exposure)
		assert.Equal(t, "SAMSUNG", data.CameraMake)
		assert.Equal(t, "SM-C200", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, projection.Equirectangular.String(), data.Projection)
	})

	t.Run("P7250006.json", func(t *testing.T) {
		data, err := JSON("testdata/P7250006.json", "P7250006.MOV")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2018-07-25T11:18:42Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2018-07-25T11:18:42Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords.String())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "", data.Exposure)
		assert.Equal(t, "OLYMPUS DIGITAL CAMERA", data.CameraMake)
		assert.Equal(t, "E-PL7", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
	})

	t.Run("P9150300.json", func(t *testing.T) {
		data, err := JSON("testdata/P9150300.json", "P9150300.MOV")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "OLYMPUS DIGITAL CAMERA", data.CameraMake)
		assert.Equal(t, "E-M10MarkII", data.CameraModel)
	})

	t.Run("GOPR0533.json", func(t *testing.T) {
		data, err := JSON("testdata/GOPR0533.json", "GOPR0533.MP4")

		if err != nil {
			t.Fatal(err)
		}

		// No make or model in this file...
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
	})

	t.Run("digikam.json", func(t *testing.T) {
		data, err := JSON("testdata/digikam.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "2020-10-17T15:48:24Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2020-10-17T17:48:24Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "Europe/Berlin", data.TimeZone)
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "berlin, shop", data.Keywords.String())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 375, data.Height)
		assert.Equal(t, 500, data.Width)
		assert.Equal(t, float32(52.46052), data.Lat)
		assert.Equal(t, float32(13.331403), data.Lng)
		assert.Equal(t, 84, data.Altitude)
		assert.Equal(t, "1/50", data.Exposure)
		assert.Equal(t, "HUAWEI", data.CameraMake)
		assert.Equal(t, "ELE-L29", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "", data.LensModel)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, int(data.Orientation))
	})

	t.Run("date.mov.json", func(t *testing.T) {
		data, err := JSON("testdata/date.mov.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, CodecAvc1, data.Codec)
		assert.Equal(t, "6s", data.Duration.String())
		assert.Equal(t, "2015-06-10 14:06:09 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2015-06-10 11:06:09 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Moscow", data.TimeZone)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1080, data.ActualWidth())
		assert.Equal(t, 1920, data.ActualHeight())
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, float32(55.5636), data.Lat)
		assert.Equal(t, float32(37.9824), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 6 Plus", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("date-creation.mov.json", func(t *testing.T) {
		data, err := JSON("testdata/date-creation.mov.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, string(video.CodecAVC), data.Codec)
		assert.Equal(t, "10s", data.Duration.String())
		assert.Equal(t, "2015-12-06 18:22:29 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2015-12-06 15:22:29 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Moscow", data.TimeZone)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1920, data.ActualWidth())
		assert.Equal(t, 1080, data.ActualHeight())
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(55.7579), data.Lat)
		assert.Equal(t, float32(37.6197), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 6 Plus", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("date-iphone8.mov.json", func(t *testing.T) {
		data, err := JSON("testdata/date-iphone8.mov.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, string(video.CodecHEVC), data.Codec)
		assert.Equal(t, "6s", data.Duration.String())
		assert.Equal(t, "2020-12-22 02:45:43 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2020-12-22 01:45:43 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "", data.TimeZone)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1080, data.ActualWidth())
		assert.Equal(t, 1920, data.ActualHeight())
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 8", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("date-iphonex.mov.json", func(t *testing.T) {
		data, err := JSON("testdata/date-iphonex.mov.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, string(video.CodecHEVC), data.Codec)
		assert.Equal(t, "2s", data.Duration.String())
		assert.Equal(t, "2019-12-12 20:47:21 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2019-12-13 01:47:21 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "America/New_York", data.TimeZone)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(40.7696), data.Lat)
		assert.Equal(t, float32(-73.9964), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone X", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("snow.json", func(t *testing.T) {
		data, err := JSON("testdata/snow.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, "2015-03-20 12:07:53 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2015-03-20 12:07:53 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, 0, data.TakenNs)
		assert.Equal(t, "", data.TimeZone)
		assert.Equal(t, 4608, data.Width)
		assert.Equal(t, 3072, data.Height)
		assert.Equal(t, 4608, data.ActualWidth())
		assert.Equal(t, 3072, data.ActualHeight())
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, "OLYMPUS IMAGING CORP.", data.CameraMake)
		assert.Equal(t, "TG-830", data.CameraModel)
		assert.Equal(t, "", data.LensModel)
	})

	t.Run("subject-1.json", func(t *testing.T) {
		data, err := JSON("testdata/subject-1.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, "2016-09-07 12:49:23.373 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2016-09-07 12:49:23.373 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, 373000000, data.TakenNs)
		assert.Equal(t, "", data.TimeZone)
		assert.Equal(t, 4032, data.Width)
		assert.Equal(t, 3024, data.Height)
		assert.Equal(t, 4032, data.ActualWidth())
		assert.Equal(t, 3024, data.ActualHeight())
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 6s", data.CameraModel)
		assert.Equal(t, "iPhone 6s back camera 4.15mm f/2.2", data.LensModel)
		assert.Equal(t, "holiday", data.Subject)
		assert.Equal(t, "holiday", data.Keywords.String())
	})

	t.Run("subject-2.json", func(t *testing.T) {
		data, err := JSON("testdata/subject-2.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "0s", data.Duration.String())
		assert.Equal(t, "2016-09-07 12:49:23.373 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2016-09-07 12:49:23.373 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "", data.TimeZone)
		assert.Equal(t, 4032, data.Width)
		assert.Equal(t, 3024, data.Height)
		assert.Equal(t, 4032, data.ActualWidth())
		assert.Equal(t, 3024, data.ActualHeight())
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, "Apple", data.CameraMake)
		assert.Equal(t, "iPhone 6s", data.CameraModel)
		assert.Equal(t, "iPhone 6s back camera 4.15mm f/2.2", data.LensModel)
		assert.Equal(t, "holiday, greetings", data.Subject)
		assert.Equal(t, "greetings, holiday", data.Keywords.String())
	})

	t.Run("newline.json", func(t *testing.T) {
		data, err := JSON("testdata/newline.json", "newline.jpg")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "Jens\r\tMander", data.Artist)
		assert.Equal(t, "2004-09-23T10:57:57Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2004-09-23T10:57:57Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "This is the title", data.Title)
		assert.Equal(t, "", data.Keywords.String())
		assert.Equal(t, "This is a\n\ndescription!", data.Description)
		assert.Equal(t, "This is the world.", data.Subject)
		assert.Equal(t, "© 2011 PhotoPrism", data.Copyright)
		assert.Equal(t, 567, data.Height)
		assert.Equal(t, 850, data.Width)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 30, data.Altitude)
		assert.Equal(t, "1/6", data.Exposure)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "Canon EOS-1DS", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "123456", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
	})

	t.Run("keywords.json", func(t *testing.T) {
		data, err := JSON("testdata/keywords.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, CodecJpeg, data.Codec)
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Artist)
		assert.Equal(t, Keywords{"alo", "cactus", "ever", "lang", "sonne"}, data.Keywords)
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "Canon EOS 7D", data.CameraModel)
		assert.Equal(t, "", data.LensMake)
		assert.Equal(t, "EF70-200mm f/4L IS USM", data.LensModel)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("quicktimeutc_on.json", func(t *testing.T) {
		data, err := JSON("testdata/quicktimeutc_on.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, string(video.CodecAVC), data.Codec)
		assert.Equal(t, "1s", data.Duration.String())
		assert.Equal(t, "2012-07-11 07:16:01 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2012-07-11 05:16:01 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Paris", data.TimeZone)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(43.5683), data.Lat)
		assert.Equal(t, float32(4.5645), data.Lng)
	})

	t.Run("quicktimeutc_off.json", func(t *testing.T) {
		data, err := JSON("testdata/quicktimeutc_off.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, string(video.CodecAVC), data.Codec)
		assert.Equal(t, "1s", data.Duration.String())
		assert.Equal(t, "2012-07-11 07:16:01 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2012-07-11 05:16:01 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Paris", data.TimeZone)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, float32(43.5683), data.Lat)
		assert.Equal(t, float32(4.5645), data.Lng)
	})

	t.Run("video_num_on.json", func(t *testing.T) {
		data, err := JSON("testdata/video_num_on.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, string(video.CodecAVC), data.Codec)
		assert.Equal(t, "1s", data.Duration.String())
		assert.Equal(t, "2012-07-11 07:16:01 +0000 UTC", data.TakenAtLocal.String())
		assert.Equal(t, "2012-07-11 05:16:01 +0000 UTC", data.TakenAt.String())
		assert.Equal(t, "Europe/Paris", data.TimeZone)
		assert.Equal(t, 6, data.Orientation)
		assert.Equal(t, float32(43.5683), data.Lat)
		assert.Equal(t, float32(4.5645), data.Lng)
	})

	t.Run("cr2_num_off.json", func(t *testing.T) {
		data, err := JSON("testdata/cr2_num_off.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2015-02-14T02:14:40Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2015-02-13T18:14:40Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2015-02-13T16:14:11Z", data.TakenGps.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2015-02-13 16:14:11.91", data.TakenGps.Format("2006-01-02 15:04:05.999999999"))
		assert.Equal(t, 3648, data.Height)
		assert.Equal(t, 5472, data.Width)
		assert.Equal(t, float32(32.843544), data.Lat)
		assert.Equal(t, float32(-117.28025), data.Lng)
		assert.Equal(t, 18, data.Altitude)
		assert.Equal(t, "1/500", data.Exposure)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "Canon EOS 6D", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "012324001432", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
	})

	t.Run("cr2_num_on.json", func(t *testing.T) {
		data, err := JSON("testdata/cr2_num_on.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2015-02-14T02:14:40Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2015-02-13T18:14:40Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, 3648, data.Height)
		assert.Equal(t, 5472, data.Width)
		assert.Equal(t, float32(32.843544), data.Lat)
		assert.Equal(t, float32(-117.28025), data.Lng)
		assert.Equal(t, 18, data.Altitude)
		assert.Equal(t, "0.002", data.Exposure)
		assert.Equal(t, "Canon", data.CameraMake)
		assert.Equal(t, "Canon EOS 6D", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "012324001432", data.CameraSerial)
		assert.Equal(t, 35, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
	})

	t.Run("pxl-mp4.json", func(t *testing.T) {
		data, err := JSON("testdata/pxl-mp4.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2021-07-12T22:56:37Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2021-07-12T22:56:37Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, time.UTC.String(), data.TimeZone)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("sony_mp4_exiftool.json", func(t *testing.T) {
		data, err := JSON("testdata/sony_mp4_exiftool.json", "")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "2021-07-06T13:51:36Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2021-07-06T13:51:36Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, time.UTC.String(), data.TimeZone)
		assert.Equal(t, 1080, data.Height)
		assert.Equal(t, 1920, data.Width)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, 1, data.Orientation)
	})

	t.Run("Iceland-P3.jpg", func(t *testing.T) {
		data, err := JSON("testdata/Iceland-P3.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "Nicolas Cornet", data.Artist)
		assert.Equal(t, "2012-08-08T22:07:18Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2012-08-08T22:07:18Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords.String())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "Nicolas Cornet", data.Copyright)
		assert.Equal(t, 400, data.Height)
		assert.Equal(t, 600, data.Width)
		assert.Equal(t, float32(65.05558), data.Lat)
		assert.Equal(t, float32(-16.625702), data.Lng)
		assert.Equal(t, 30, data.Altitude)
		assert.Equal(t, "1/8", data.Exposure)
		assert.Equal(t, "NIKON CORPORATION", data.CameraMake)
		assert.Equal(t, "NIKON D800E", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "6001440", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
		assert.Equal(t, "Display P3", data.ColorProfile)
	})

	t.Run("Iceland-P3-n.jpg", func(t *testing.T) {
		data, err := JSON("testdata/Iceland-P3-n.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "Nicolas Cornet", data.Artist)
		assert.Equal(t, "2012-08-08T22:07:18Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2012-08-08T22:07:18Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords.String())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "Nicolas Cornet", data.Copyright)
		assert.Equal(t, 400, data.Height)
		assert.Equal(t, 600, data.Width)
		assert.Equal(t, float32(65.05558), data.Lat)
		assert.Equal(t, float32(-16.625702), data.Lng)
		assert.Equal(t, 30, data.Altitude)
		assert.Equal(t, "0.125", data.Exposure)
		assert.Equal(t, "NIKON CORPORATION", data.CameraMake)
		assert.Equal(t, "NIKON D800E", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "6001440", data.CameraSerial)
		assert.Equal(t, 16, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
		assert.Equal(t, "Display P3", data.ColorProfile)
	})

	t.Run("Iceland-sRGB.jpg", func(t *testing.T) {
		data, err := JSON("testdata/Iceland-sRGB.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "Nicolas Cornet", data.Artist)
		assert.Equal(t, "2012-08-08T22:07:18Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "2012-08-08T22:07:18Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords.String())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "Nicolas Cornet", data.Copyright)
		assert.Equal(t, 400, data.Height)
		assert.Equal(t, 600, data.Width)
		assert.Equal(t, float32(65.05558), data.Lat)
		assert.Equal(t, float32(-16.625702), data.Lng)
		assert.Equal(t, 30, data.Altitude)
		assert.Equal(t, "1/8", data.Exposure)
		assert.Equal(t, "NIKON CORPORATION", data.CameraMake)
		assert.Equal(t, "NIKON D800E", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "6001440", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
		assert.Equal(t, "Display P3", data.ColorProfile)
	})

	t.Run("gif.json", func(t *testing.T) {
		data, err := JSON("testdata/gif.json", "")

		if err != nil {
			t.Fatal(err)
		}

		// t.Logf("all: %+v", data.All)

		assert.Equal(t, "", data.Artist)
		assert.Equal(t, "0001-01-01T00:00:00Z", data.TakenAt.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "0001-01-01T00:00:00Z", data.TakenAtLocal.Format("2006-01-02T15:04:05Z"))
		assert.Equal(t, "", data.Title)
		assert.Equal(t, "", data.Keywords.String())
		assert.Equal(t, "", data.Description)
		assert.Equal(t, "", data.Copyright)
		assert.Equal(t, 1533, data.Height)
		assert.Equal(t, 1917, data.Width)
		assert.Equal(t, 34, data.Frames)
		assert.Equal(t, 49*time.Second, data.Duration)
		assert.Equal(t, float32(0), data.Lat)
		assert.Equal(t, float32(0), data.Lng)
		assert.Equal(t, 0, data.Altitude)
		assert.Equal(t, "", data.Exposure)
		assert.Equal(t, "", data.CameraMake)
		assert.Equal(t, "", data.CameraModel)
		assert.Equal(t, "", data.CameraOwner)
		assert.Equal(t, "", data.CameraSerial)
		assert.Equal(t, 0, data.FocalLength)
		assert.Equal(t, 1, data.Orientation)
		assert.Equal(t, "", data.Projection)
		assert.Equal(t, "", data.ColorProfile)
	})

	t.Run("iptc-fields-500", func(t *testing.T) {
		data, err := JSON("testdata/iptc-fields-500.json", "")

		if err != nil {
			t.Fatal(err)
		}
		//t.Logf("all: %+v", data.exif)

		assert.Equal(t, "creator A, creator B", data.Artist)
		assert.Equal(t, "my image headline", data.Title)
		assert.Equal(t, "my iptc description", data.Description)
		assert.Equal(t, "my iptc copyright", data.Copyright)
		//TODO
		//assert.Equal(t, "zqdtcxt1q9wrxnur", data.DocumentID)
	})
}
