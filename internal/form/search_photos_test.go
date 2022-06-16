package form

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPhotoSearchForm(t *testing.T) {
	form := &SearchPhotos{}

	assert.IsType(t, new(SearchPhotos), form)
}

func TestParseQueryString(t *testing.T) {
	t.Run("subjects", func(t *testing.T) {
		form := &SearchPhotos{Query: "subjects:\"Jens & Mander\""}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Jens & Mander", form.Subjects)
	})
	t.Run("aliases", func(t *testing.T) {
		form := &SearchPhotos{Query: "people:\"Jens & Mander\" folder:Foo person:Bar"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", form.Folder)
		assert.Equal(t, "", form.Person)
		assert.Equal(t, "", form.People)
		assert.Equal(t, "Foo", form.Path)
		assert.Equal(t, "Bar", form.Subject)
		assert.Equal(t, "Jens & Mander", form.Subjects)
	})
	t.Run("keywords", func(t *testing.T) {
		form := &SearchPhotos{Query: "keywords:\"Foo Bar\""}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Foo Bar", form.Keywords)
	})
	t.Run("and query", func(t *testing.T) {
		form := &SearchPhotos{Query: "\"Jens & Mander\" title:\"Tübingen\""}

		err := form.ParseQueryString()

		// log.Debugf("%+v\n", form)

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "jens & mander", form.GetQuery())
		assert.Equal(t, "Tübingen", form.Title)
	})
	t.Run("path", func(t *testing.T) {
		form := &SearchPhotos{Query: "path:123abc/,EFG"}

		err := form.ParseQueryString()

		// log.Debugf("%+v\n", form)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "123abc/,EFG", form.Path)
	})

	t.Run("folder", func(t *testing.T) {
		form := &SearchPhotos{Query: "folder:123abc/,EFG"}

		err := form.ParseQueryString()

		// log.Debugf("%+v\n", form)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "123abc/,EFG", form.Path)
	})
	t.Run("valid query", func(t *testing.T) {
		form := &SearchPhotos{Query: "label:cat query:\"fooBar baz\" before:2019-01-15 camera:23 favorite:false dist:25000 lat:33.45343166666667"}

		err := form.ParseQueryString()

		// log.Debugf("%+v\n", form)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "cat", form.Label)
		assert.Equal(t, "fooBar baz", form.Query)
		assert.Equal(t, "23", form.Camera)
		assert.Equal(t, time.Date(2019, 01, 15, 0, 0, 0, 0, time.UTC), form.Before)
		assert.Equal(t, false, form.Favorite)
		assert.Equal(t, uint(0x61a8), form.Dist)
		assert.Equal(t, float32(33.45343), form.Lat)
	})
	t.Run("valid query 2", func(t *testing.T) {
		form := &SearchPhotos{Query: "chroma:200 title:\"te:st\" after:2018-01-15 favorite:true lng:33.45343166666667"}

		err := form.ParseQueryString()

		// log.Debugf("%+v\n", form)

		if err != nil {
			t.Fatal("err should be nil")
		}

		assert.Equal(t, uint8(200), form.Chroma)
		assert.Equal(t, "te:st", form.Title)
		assert.Equal(t, time.Date(2018, 01, 15, 0, 0, 0, 0, time.UTC), form.After)
		assert.Equal(t, float32(33.45343), form.Lng)
	})
	t.Run("valid query with filter", func(t *testing.T) {
		form := &SearchPhotos{Query: "label:cat title:\"fooBar baz\"", Filter: "label:dog"}

		err := form.ParseQueryString()

		// log.Debugf("%+v\n", form)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "dog", form.Label)
		assert.Equal(t, "label:dog", form.Filter)
		assert.Equal(t, "fooBar baz", form.Title)
	})
	t.Run("valid query with umlauts", func(t *testing.T) {
		form := &SearchPhotos{Query: "title:\"tübingen\""}

		err := form.ParseQueryString()

		// log.Debugf("%+v\n", form)

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "tübingen", form.Title)
	})
	t.Run("query for invalid filter", func(t *testing.T) {
		form := &SearchPhotos{Query: "xxx:false"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Equal(t, "unknown filter: Xxx", err.Error())
	})
	t.Run("query for favorites with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "favorite:cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Favorite)
	})
	t.Run("query for primary with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "primary:&cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Primary)
	})
	t.Run("query for stack with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "stack:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Stack)
	})
	t.Run("query for unstacked with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "unstacked:'cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Unstacked)
	})
	t.Run("query for stackable with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "stackable:mother's day"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Stackable)
	})
	t.Run("query for video with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "video:|cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Video)
	})
	t.Run("AnimatedYes", func(t *testing.T) {
		form := &SearchPhotos{Query: "animated:yes"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}
		assert.False(t, form.Vector)
		assert.True(t, form.Animated)
	})
	t.Run("VectorYes", func(t *testing.T) {
		form := &SearchPhotos{Query: "vector:yes"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}
		assert.False(t, form.Animated)
		assert.True(t, form.Vector)
	})
	t.Run("query for photo with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "photo:cat>"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Photo)
	})
	t.Run("query for raw with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "raw:ca+(t"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Raw)
	})
	t.Run("query for live with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "live:cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Live)
	})
	t.Run("query for scan with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "scan:;cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Scan)
	})
	t.Run("query for panorama with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "panorama:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Panorama)
	})
	t.Run("query for error with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "error:^cat$#"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Error)
	})
	t.Run("query for hidden with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "hidden:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Hidden)
	})
	t.Run("query for archived with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "archived:`cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Archived)
	})
	t.Run("query for public with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "public:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Public)
	})
	t.Run("query for private with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "private:*c@t"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Private)
	})
	t.Run("query for unsorted with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "unsorted:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Unsorted)
	})
	t.Run("query for mono with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "mono:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Mono)
	})
	t.Run("query for portrait with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "portrait:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Portrait)
	})
	t.Run("PortraitLandscapeSquare", func(t *testing.T) {
		form := &SearchPhotos{Query: "portrait:true landscape:yes square:jo"}

		assert.False(t, form.Portrait)
		assert.False(t, form.Landscape)
		assert.False(t, form.Square)
		assert.False(t, form.Panorama)

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Portrait)
		assert.True(t, form.Landscape)
		assert.True(t, form.Square)
		assert.False(t, form.Panorama)
	})
	t.Run("query for geo with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "geo:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Geo)
	})
	t.Run("query for review with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "review:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Review)
	})
	t.Run("query for merged with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "merged:*cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Merged)
	})
	t.Run("query for landscape with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "landscape:test$5123"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Landscape)
	})
	t.Run("query for square with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "square:%abc"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Square)
	})
	t.Run("query for animated with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "animated:%abc"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Animated)
	})
	t.Run("query for vector with uncommon bool value", func(t *testing.T) {
		form := &SearchPhotos{Query: "vector:%abc"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, form.Vector)
	})
	t.Run("query for lat with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "lat:&cat"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for lng with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "lng:^>cat"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for dist with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "dist:c@t"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for fmin with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "fmin:=}cat{"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for fmax with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "fmax:ca#$t"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for chroma with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "chroma:&|cat"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for diff with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "diff:&cat;%"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for quality with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "quality:`cat"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for count with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "count:ca(%t"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("query for offset with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "offset:&cat"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Contains(t, err.Error(), "invalid syntax")
	})
	t.Run("CameraString", func(t *testing.T) {
		form := &SearchPhotos{Query: "camera:cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "cat", form.Camera)
	})
	t.Run("LensString", func(t *testing.T) {
		form := &SearchPhotos{Query: "lens:cat"}

		err := form.ParseQueryString()

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "cat", form.Lens)
	})
	t.Run("query for before with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "before:cat"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Equal(t, "Could not find format for \"cat\"", err.Error())
	})
	t.Run("query for after with invalid type", func(t *testing.T) {
		form := &SearchPhotos{Query: "after:cat"}

		err := form.ParseQueryString()

		if err == nil {
			t.Fatal(err)
		}

		// log.Debugf("%+v\n", form)

		assert.Equal(t, "Could not find format for \"cat\"", err.Error())
	})
}

func TestNewPhotoSearch(t *testing.T) {
	r := NewPhotoSearch("cat")
	assert.IsType(t, SearchPhotos{}, r)
}

func TestPhotoSearch_Serialize(t *testing.T) {
	form := SearchPhotos{
		Query:   "foo BAR",
		Private: true,
		Photo:   false,
		Lat:     1.5,
		Lng:     -10.33333,
		Year:    "2002",
		Chroma:  1,
		Diff:    424242,
		Before:  time.Date(2019, 01, 15, 0, 0, 0, 0, time.UTC),
	}

	result := form.Serialize()

	t.Logf("SERIALIZED: %s", result)

	assert.IsType(t, "string", result)
}

func TestPhotoSearch_SerializeAll(t *testing.T) {
	form := SearchPhotos{
		Query:   "foo BAR",
		Private: true,
		Photo:   false,
		Lat:     1.5,
		Lng:     -10.33333,
		Year:    "2002|2003",
		Chroma:  1,
		Diff:    424242,
		Before:  time.Date(2019, 01, 15, 0, 0, 0, 0, time.UTC),
	}

	result := form.SerializeAll()

	t.Logf("SERIALIZED: %s", result)

	assert.IsType(t, "string", result)
}
