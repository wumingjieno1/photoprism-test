package photoprism

import (
	"image"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/entity"
	"github.com/photoprism/photoprism/internal/thumb"
	"github.com/photoprism/photoprism/pkg/fs"

	"github.com/stretchr/testify/assert"
)

func TestMediaFile_DateCreated(t *testing.T) {
	conf := config.TestConfig()

	t.Run("telegram_2020-01-30_09-57-18.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/telegram_2020-01-30_09-57-18.jpg")
		if err != nil {
			t.Fatal(err)
		}
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2020-01-30 09:57:18 +0000 UTC", date.String())
	})
	t.Run("Screenshot 2019-05-21 at 10.45.52.png", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Screenshot 2019-05-21 at 10.45.52.png")
		if err != nil {
			t.Fatal(err)
		}
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2019-05-21 10:45:52 +0000 UTC", date.String())
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2018-09-10 03:16:13 +0000 UTC", date.String())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2019-06-06 07:29:51 +0000 UTC", date.String())
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "2013-11-26 13:53:55 +0000 UTC", date.String())
	})
	t.Run("dog_created_1919.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/dog_created_1919.jpg")
		if err != nil {
			t.Fatal(err)
		}
		date := mediaFile.DateCreated().UTC()
		assert.Equal(t, "1919-05-04 05:59:26 +0000 UTC", date.String())
	})
}

func TestMediaFile_TakenAt(t *testing.T) {
	conf := config.TestConfig()
	t.Run("testdata/2018-04-12 19_24_49.gif", func(t *testing.T) {
		mediaFile, err := NewMediaFile("testdata/2018-04-12 19_24_49.gif")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "2018-04-12 19:24:49 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcName, src)
	})
	t.Run("testdata/2018-04-12 19_24_49.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile("testdata/2018-04-12 19_24_49.jpg")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "2018-04-12 19:24:49 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcName, src)
	})
	t.Run("telegram_2020-01-30_09-57-18.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/telegram_2020-01-30_09-57-18.jpg")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "2020-01-30 09:57:18 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcName, src)
	})
	t.Run("Screenshot 2019-05-21 at 10.45.52.png", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Screenshot 2019-05-21 at 10.45.52.png")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "2019-05-21 10:45:52 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcName, src)
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "2018-09-10 03:16:13 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcMeta, src)
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "2019-06-06 07:29:51 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcMeta, src)
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "2013-11-26 13:53:55 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcMeta, src)
	})
	t.Run("dog_created_1919.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/dog_created_1919.jpg")
		if err != nil {
			t.Fatal(err)
		}

		date, src := mediaFile.TakenAt()
		assert.Equal(t, "1919-05-04 05:59:26 +0000 UTC", date.String())
		assert.Equal(t, entity.SrcMeta, src)
	})
}

func TestMediaFile_HasTimeAndPlace(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.HasTimeAndPlace())
	})
	t.Run("/peacock_blue.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/peacock_blue.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.HasTimeAndPlace())
	})
}
func TestMediaFile_CameraModel(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "iPhone SE", mediaFile.CameraModel())
	})
	t.Run("/iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "iPhone 7", mediaFile.CameraModel())
	})
}

func TestMediaFile_CameraMake(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Apple", mediaFile.CameraMake())
	})
	t.Run("/peacock_blue.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/peacock_blue.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.CameraMake())
	})
}

func TestMediaFile_LensModel(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "iPhone SE back camera 4.15mm f/2.2", mediaFile.LensModel())
	})
	t.Run("/canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "EF24-105mm f/4L IS USM", mediaFile.LensModel())
	})
}

func TestMediaFile_LensMake(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Apple", mediaFile.LensMake())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.LensMake())
	})
}

func TestMediaFile_FocalLength(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 29, mediaFile.FocalLength())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 111, mediaFile.FocalLength())
	})
}

func TestMediaFile_FNumber(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, float32(2.2), mediaFile.FNumber())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, float32(10.0), mediaFile.FNumber())
	})
}

func TestMediaFile_Iso(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 32, mediaFile.Iso())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 200, mediaFile.Iso())
	})
}

func TestMediaFile_Exposure(t *testing.T) {
	t.Run("/cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/cat_brown.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "1/50", mediaFile.Exposure())
	})
	t.Run("/elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "1/640", mediaFile.Exposure())
	})
}

func TestMediaFileCanonicalName(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "20180111_110938_7D8F8A23", mediaFile.CanonicalName())
}

func TestMediaFileCanonicalNameFromFile(t *testing.T) {
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "beach_wood", mediaFile.CanonicalNameFromFile())
	})
	t.Run("/airport_grey", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/airport_grey")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "airport_grey", mediaFile.CanonicalNameFromFile())
	})
}

func TestMediaFile_CanonicalNameFromFileWithDirectory(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_wood.jpg")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, conf.ExamplesPath()+"/beach_wood", mediaFile.CanonicalNameFromFileWithDirectory())
}

func TestMediaFile_EditedFilename(t *testing.T) {
	conf := config.TestConfig()

	t.Run("IMG_4120.JPG", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.JPG")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, conf.ExamplesPath()+"/IMG_E4120.JPG", mediaFile.EditedName())
	})

	t.Run("fern_green.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/fern_green.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.EditedName())
	})
}

func TestMediaFile_RelatedFiles(t *testing.T) {
	conf := config.TestConfig()

	t.Run("example.tif", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/example.tif")

		if err != nil {
			t.Fatal(err)
		}

		related, err := mediaFile.RelatedFiles(true)

		if err != nil {
			t.Fatal(err)
		}

		assert.Len(t, related.Files, 7)
		assert.True(t, related.ContainsJpeg())

		for _, result := range related.Files {
			t.Logf("FileName: %s", result.FileName())

			filename := result.FileName()

			if len(filename) < 2 {
				t.Fatalf("filename not be longer: %s", filename)
			}

			extension := result.Extension()

			if len(extension) < 2 {
				t.Fatalf("extension should be longer: %s", extension)
			}

			relativePath := result.RelPath(conf.ExamplesPath())

			if len(relativePath) > 0 {
				t.Fatalf("relative path should be empty: %s", relativePath)
			}
		}
	})

	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")

		if err != nil {
			t.Fatal(err)
		}

		expectedBaseFilename := conf.ExamplesPath() + "/canon_eos_6d"

		related, err := mediaFile.RelatedFiles(true)

		if err != nil {
			t.Fatal(err)
		}

		assert.Len(t, related.Files, 3)
		assert.False(t, related.ContainsJpeg())

		for _, result := range related.Files {
			t.Logf("FileName: %s", result.FileName())

			filename := result.FileName()

			extension := result.Extension()

			baseFilename := filename[0 : len(filename)-len(extension)]

			assert.Equal(t, expectedBaseFilename, baseFilename)
		}
	})

	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

		if err != nil {
			t.Fatal(err)
		}

		expectedBaseFilename := conf.ExamplesPath() + "/iphone_7"

		related, err := mediaFile.RelatedFiles(true)

		if err != nil {
			t.Fatal(err)
		}

		assert.GreaterOrEqual(t, len(related.Files), 3)

		for _, result := range related.Files {
			t.Logf("FileName: %s", result.FileName())

			filename := result.FileName()
			extension := result.Extension()
			baseFilename := filename[0 : len(filename)-len(extension)]

			if result.IsJpeg() {
				assert.Contains(t, expectedBaseFilename, "examples/iphone_7")
			} else {
				assert.Equal(t, expectedBaseFilename, baseFilename)
			}
		}
	})

	t.Run("2015-02-04.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile("testdata/2015-02-04.jpg")

		if err != nil {
			t.Fatal(err)
		}

		related, err := mediaFile.RelatedFiles(true)

		if err != nil {
			t.Fatal(err)
		}

		if related.Main == nil {
			t.Fatal("main file must not be nil")
		}

		if len(related.Files) != 4 {
			t.Fatalf("length is %d, should be 4", len(related.Files))
		}

		t.Logf("FILE: %s, %s", related.Main.FileType(), related.Main.MimeType())

		assert.Equal(t, "2015-02-04.jpg", related.Main.BaseName())

		assert.Equal(t, "2015-02-04.jpg", related.Files[0].BaseName())
		assert.Equal(t, "2015-02-04(1).jpg", related.Files[1].BaseName())
		assert.Equal(t, "2015-02-04.jpg.json", related.Files[2].BaseName())
		assert.Equal(t, "2015-02-04.jpg(1).json", related.Files[3].BaseName())
	})

	t.Run("2015-02-04(1).jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile("testdata/2015-02-04(1).jpg")

		if err != nil {
			t.Fatal(err)
		}

		related, err := mediaFile.RelatedFiles(false)

		if err != nil {
			t.Fatal(err)
		}

		if related.Main == nil {
			t.Fatal("main file must not be nil")
		}

		if len(related.Files) != 1 {
			t.Fatalf("length is %d, should be 1", len(related.Files))
		}

		assert.Equal(t, "2015-02-04(1).jpg", related.Main.BaseName())

		assert.Equal(t, "2015-02-04(1).jpg", related.Files[0].BaseName())
	})

	t.Run("2015-02-04(1).jpg stacked", func(t *testing.T) {
		mediaFile, err := NewMediaFile("testdata/2015-02-04(1).jpg")

		if err != nil {
			t.Fatal(err)
		}

		related, err := mediaFile.RelatedFiles(true)

		if err != nil {
			t.Fatal(err)
		}

		if related.Main == nil {
			t.Fatal("main file must not be nil")
		}

		if len(related.Files) != 4 {
			t.Fatalf("length is %d, should be 4", len(related.Files))
		}

		assert.Equal(t, "2015-02-04.jpg", related.Main.BaseName())

		assert.Equal(t, "2015-02-04.jpg", related.Files[0].BaseName())
		assert.Equal(t, "2015-02-04(1).jpg", related.Files[1].BaseName())
		assert.Equal(t, "2015-02-04.jpg.json", related.Files[2].BaseName())
		assert.Equal(t, "2015-02-04.jpg(1).json", related.Files[3].BaseName())
	})
}

func TestMediaFile_RelatedFiles_Ordering(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.JPG")

	if err != nil {
		t.Fatal(err)
	}

	related, err := mediaFile.RelatedFiles(true)

	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, related.Files, 5)

	assert.Equal(t, conf.ExamplesPath()+"/IMG_4120.AAE", related.Files[0].FileName())
	assert.Equal(t, conf.ExamplesPath()+"/IMG_4120.JPG", related.Files[1].FileName())

	for _, result := range related.Files {
		filename := result.FileName()
		t.Logf("FileName: %s", filename)
	}
}

func TestMediaFile_SetFilename(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/turtle_brown_blue.jpg")
	if err != nil {
		t.Fatal(err)
	}
	mediaFile.SetFileName("newFilename")
	assert.Equal(t, "newFilename", mediaFile.fileName)
	mediaFile.SetFileName("turtle_brown_blue")
	assert.Equal(t, "turtle_brown_blue", mediaFile.fileName)
}

func TestMediaFile_RootRelName(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")

	if err != nil {
		t.Fatal(err)
	}

	t.Run("examples_path", func(t *testing.T) {
		filename := mediaFile.RootRelName()
		assert.Equal(t, "tree_white.jpg", filename)
	})
}

func TestMediaFile_RootRelPath(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")
	mediaFile.fileRoot = entity.RootImport
	if err != nil {
		t.Fatal(err)
	}

	t.Run("examples_path", func(t *testing.T) {
		path := mediaFile.RootRelPath()
		assert.Equal(t, conf.ExamplesPath(), path)
	})
}

func TestMediaFile_RootPath(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")

	if err != nil {
		t.Fatal(err)
	}

	mediaFile.fileRoot = entity.RootImport
	t.Run("examples_path", func(t *testing.T) {
		path := mediaFile.RootPath()
		assert.Contains(t, path, "import")
	})
}

func TestMediaFile_RelName(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")

	if err != nil {
		t.Fatal(err)
	}

	t.Run("directory with end slash", func(t *testing.T) {
		filename := mediaFile.RelName(conf.AssetsPath())
		assert.Equal(t, "examples/tree_white.jpg", filename)
	})

	t.Run("directory without end slash", func(t *testing.T) {
		filename := mediaFile.RelName(conf.AssetsPath())
		assert.Equal(t, "examples/tree_white.jpg", filename)
	})
	t.Run("directory not part of filename", func(t *testing.T) {
		filename := mediaFile.RelName("xxx/")
		assert.Equal(t, conf.ExamplesPath()+"/tree_white.jpg", filename)
	})
	t.Run("directory equals example path", func(t *testing.T) {
		filename := mediaFile.RelName(conf.ExamplesPath())
		assert.Equal(t, "tree_white.jpg", filename)
	})
}

func TestMediaFile_RelativePath(t *testing.T) {

	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")

	if err != nil {
		t.Fatal(err)
	}

	t.Run("directory with end slash", func(t *testing.T) {
		path := mediaFile.RelPath(conf.AssetsPath() + "/")
		assert.Equal(t, "examples", path)
	})
	t.Run("directory without end slash", func(t *testing.T) {
		path := mediaFile.RelPath(conf.AssetsPath())
		assert.Equal(t, "examples", path)
	})
	t.Run("directory equals filepath", func(t *testing.T) {
		path := mediaFile.RelPath(conf.ExamplesPath())
		assert.Equal(t, "", path)
	})
	t.Run("directory does not match filepath", func(t *testing.T) {
		path := mediaFile.RelPath("xxx")
		assert.Equal(t, conf.ExamplesPath(), path)
	})

	mediaFile, err = NewMediaFile(conf.ExamplesPath() + "/.photoprism/example.jpg")

	if err != nil {
		t.Fatal(err)
	}

	t.Run("hidden", func(t *testing.T) {
		path := mediaFile.RelPath(conf.ExamplesPath())
		assert.Equal(t, "", path)
	})
	t.Run("hidden empty", func(t *testing.T) {
		path := mediaFile.RelPath("")
		assert.Equal(t, conf.ExamplesPath(), path)
	})
	t.Run("hidden root", func(t *testing.T) {
		path := mediaFile.RelPath(filepath.Join(conf.ExamplesPath(), fs.HiddenPath))
		assert.Equal(t, "", path)
	})
}

func TestMediaFile_RelativeBasename(t *testing.T) {
	conf := config.TestConfig()

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tree_white.jpg")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("directory with end slash", func(t *testing.T) {
		basename := mediaFile.RelPrefix(conf.AssetsPath()+"/", true)
		assert.Equal(t, "examples/tree_white", basename)
	})
	t.Run("directory without end slash", func(t *testing.T) {
		basename := mediaFile.RelPrefix(conf.AssetsPath(), true)
		assert.Equal(t, "examples/tree_white", basename)
	})
	t.Run("directory equals example path", func(t *testing.T) {
		basename := mediaFile.RelPrefix(conf.ExamplesPath(), true)
		assert.Equal(t, "tree_white", basename)
	})

}

func TestMediaFile_Directory(t *testing.T) {
	t.Run("/limes.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/limes.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, conf.ExamplesPath(), mediaFile.Dir())
	})
}

func TestMediaFile_Basename(t *testing.T) {
	t.Run("/limes.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/limes.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "limes", mediaFile.BasePrefix(true))
	})
	t.Run("/IMG_4120 copy.JPG", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120 copy.JPG")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "IMG_4120", mediaFile.BasePrefix(true))
	})
	t.Run("/IMG_4120 (1).JPG", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120 (1).JPG")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "IMG_4120", mediaFile.BasePrefix(true))
	})
}

func TestMediaFile_MimeType(t *testing.T) {
	conf := config.TestConfig()

	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "image/jpeg", mediaFile.MimeType())
	})

	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "image/tiff", mediaFile.MimeType())

	})

	t.Run("iphone_7.xmp", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.MimeType())
	})

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.MimeType())
	})

	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "image/heif", mediaFile.MimeType())
	})

	t.Run("IMG_4120.AAE", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.AAE")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.MimeType())
	})

	t.Run("earth.mov", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "earth.mov")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, "video/quicktime", f.MimeType())
		}
	})

	t.Run("blue-go-video.mp4", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "blue-go-video.mp4")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, "video/mp4", f.MimeType())
		}
	})

	t.Run("earth.avi", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "earth.avi")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, "video/x-msvideo", f.MimeType())
		}
	})
}

func TestMediaFile_Exists(t *testing.T) {
	conf := config.TestConfig()

	exists, err := NewMediaFile(conf.ExamplesPath() + "/cat_black.jpg")

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, exists)
	assert.True(t, exists.Exists())

	missing, err := NewMediaFile(conf.ExamplesPath() + "/xxz.jpg")

	assert.NotNil(t, exists)
	assert.Error(t, err)
	assert.Equal(t, int64(-1), missing.FileSize())
}

func TestMediaFile_Move(t *testing.T) {
	conf := config.TestConfig()

	tmpPath := conf.CachePath() + "/_tmp/TestMediaFile_Move"
	origName := tmpPath + "/original.jpg"
	destName := tmpPath + "/destination.jpg"

	if err := os.MkdirAll(tmpPath, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tmpPath)

	f, err := NewMediaFile(conf.ExamplesPath() + "/table_white.jpg")

	if err != nil {
		t.Fatal(err)
	}

	if err := f.Copy(origName); err != nil {
		t.Fatal(err)
	}

	assert.True(t, fs.FileExists(origName))

	m, err := NewMediaFile(origName)
	if err != nil {
		t.Fatal(err)
	}

	if err = m.Move(destName); err != nil {
		t.Errorf("failed to move: %s", err)
	}

	assert.True(t, fs.FileExists(destName))
	assert.Equal(t, destName, m.FileName())
}

func TestMediaFile_Copy(t *testing.T) {
	conf := config.TestConfig()

	tmpPath := conf.CachePath() + "/_tmp/TestMediaFile_Copy"

	if err := os.MkdirAll(tmpPath, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(tmpPath)

	mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/table_white.jpg")

	if err != nil {
		t.Fatal(err)
	}

	if err := mediaFile.Copy(tmpPath + "table_whitecopy.jpg"); err != nil {
		t.Fatal(err)
	}

	assert.True(t, fs.FileExists(tmpPath+"table_whitecopy.jpg"))
}

func TestMediaFile_Extension(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, ".json", mediaFile.Extension())
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, ".heic", mediaFile.Extension())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, ".dng", mediaFile.Extension())
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fs.ExtJPEG, mediaFile.Extension())
	})
}

func TestMediaFile_IsJpeg(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsJpeg())
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsJpeg())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsJpeg())
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsJpeg())
	})
}

func TestMediaFile_HasType(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.HasFileType("jpg"))
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.HasFileType("heif"))
	})
	t.Run("iphone_7.xmp", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.HasFileType("xmp"))
	})
}

func TestMediaFile_IsHEIF(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsHEIF())
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsHEIF())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsHEIF())
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsHEIF())
	})
}

func TestMediaFile_IsRaw(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsRaw())
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsRaw())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, true, mediaFile.IsRaw())
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsRaw())
	})
}

func TestMediaFile_IsPng(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsPng())
	})
	t.Run("tweethog.png", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tweethog.png")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, fs.ImagePNG, mediaFile.FileType())
		assert.Equal(t, "image/png", mediaFile.MimeType())
		assert.Equal(t, true, mediaFile.IsPng())
	})
}

func TestMediaFile_IsTiff(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fs.JsonFile, mediaFile.FileType())
		assert.Equal(t, "", mediaFile.MimeType())
		assert.Equal(t, false, mediaFile.IsTiff())
	})
	t.Run("purple.tiff", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/purple.tiff")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fs.ImageTIFF, mediaFile.FileType())
		assert.Equal(t, "image/tiff", mediaFile.MimeType())
		assert.Equal(t, true, mediaFile.IsTiff())
	})
	t.Run("example.tiff", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/example.tif")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fs.ImageTIFF, mediaFile.FileType())
		assert.Equal(t, "image/tiff", mediaFile.MimeType())
		assert.Equal(t, true, mediaFile.IsTiff())
	})
}

func TestMediaFile_IsImageOther(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsImageOther())
	})
	t.Run("purple.tiff", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/purple.tiff")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsImageOther())
	})
	t.Run("tweethog.png", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/tweethog.png")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsJpeg())
		assert.Equal(t, false, mediaFile.IsGif())
		assert.Equal(t, true, mediaFile.IsPng())
		assert.Equal(t, false, mediaFile.IsBitmap())
		assert.Equal(t, false, mediaFile.IsWebP())
		assert.Equal(t, true, mediaFile.IsImage())
		assert.Equal(t, true, mediaFile.IsImageNative())
		assert.Equal(t, true, mediaFile.IsImageOther())
		assert.Equal(t, false, mediaFile.IsVideo())
		assert.Equal(t, false, mediaFile.IsPlayableVideo())
	})
	t.Run("yellow_rose-small.bmp", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/yellow_rose-small.bmp")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, fs.ImageBMP, mediaFile.FileType())
		assert.Equal(t, "image/bmp", mediaFile.MimeType())
		assert.Equal(t, false, mediaFile.IsJpeg())
		assert.Equal(t, false, mediaFile.IsGif())
		assert.Equal(t, true, mediaFile.IsBitmap())
		assert.Equal(t, false, mediaFile.IsWebP())
		assert.Equal(t, true, mediaFile.IsImage())
		assert.Equal(t, true, mediaFile.IsImageNative())
		assert.Equal(t, true, mediaFile.IsImageOther())
		assert.Equal(t, false, mediaFile.IsVideo())
		assert.Equal(t, false, mediaFile.IsPlayableVideo())
	})
	t.Run("preloader.gif", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/preloader.gif")
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, fs.ImageGIF, mediaFile.FileType())
		assert.Equal(t, "image/gif", mediaFile.MimeType())
		assert.Equal(t, false, mediaFile.IsJpeg())
		assert.Equal(t, true, mediaFile.IsGif())
		assert.Equal(t, false, mediaFile.IsBitmap())
		assert.Equal(t, false, mediaFile.IsWebP())
		assert.Equal(t, true, mediaFile.IsImage())
		assert.Equal(t, true, mediaFile.IsImageNative())
		assert.Equal(t, true, mediaFile.IsImageOther())
		assert.Equal(t, false, mediaFile.IsVideo())
		assert.Equal(t, false, mediaFile.IsPlayableVideo())
	})
	t.Run("norway-kjetil-moe.webp", func(t *testing.T) {
		mediaFile, err := NewMediaFile("testdata/norway-kjetil-moe.webp")

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, fs.ImageWebP, mediaFile.FileType())
		assert.Equal(t, fs.MimeTypeWebP, mediaFile.MimeType())
		assert.Equal(t, false, mediaFile.IsJpeg())
		assert.Equal(t, false, mediaFile.IsGif())
		assert.Equal(t, false, mediaFile.IsBitmap())
		assert.Equal(t, true, mediaFile.IsWebP())
		assert.Equal(t, true, mediaFile.IsImage())
		assert.Equal(t, true, mediaFile.IsImageNative())
		assert.Equal(t, true, mediaFile.IsImageOther())
		assert.Equal(t, false, mediaFile.IsVideo())
		assert.Equal(t, false, mediaFile.IsPlayableVideo())
	})
}

func TestMediaFile_IsSidecar(t *testing.T) {
	conf := config.TestConfig()

	t.Run("iphone_7.xmp", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("IMG_4120.AAE", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/IMG_4120.AAE")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("test.xml", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.xml")

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("test.txt", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.txt")

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("test.yml", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.yml")

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("test.md", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.md")

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsSidecar())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")

		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsSidecar())
	})
}

func TestMediaFile_IsImage(t *testing.T) {
	t.Run("iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, false, mediaFile.IsImage())
	})
	t.Run("iphone_7.xmp", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.xmp")
		assert.Nil(t, err)
		assert.Equal(t, false, mediaFile.IsImage())
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsImage())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, true, mediaFile.IsImage())
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")
		assert.Nil(t, err)
		assert.Equal(t, true, mediaFile.IsImage())
	})
}

func TestMediaFile_IsVideo(t *testing.T) {
	conf := config.TestConfig()

	t.Run("christmas.mp4", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "christmas.mp4")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, false, f.IsImage())
			assert.Equal(t, true, f.IsVideo())
			assert.Equal(t, false, f.IsJson())
			assert.Equal(t, false, f.IsSidecar())
		}
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "canon_eos_6d.dng")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, true, f.IsImage())
			assert.Equal(t, false, f.IsVideo())
			assert.Equal(t, false, f.IsJson())
			assert.Equal(t, false, f.IsSidecar())
		}
	})
	t.Run("iphone_7.json", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "iphone_7.json")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, false, f.IsImage())
			assert.Equal(t, false, f.IsVideo())
			assert.Equal(t, true, f.IsJson())
			assert.Equal(t, true, f.IsSidecar())
		}
	})
}

func TestMediaFile_IsAnimated(t *testing.T) {
	conf := config.TestConfig()

	t.Run("example.gif", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "example.gif")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, true, f.IsImage())
			assert.Equal(t, false, f.IsVideo())
			assert.Equal(t, false, f.IsAnimated())
			assert.Equal(t, true, f.IsGif())
			assert.Equal(t, false, f.IsAnimatedGif())
			assert.Equal(t, false, f.IsSidecar())
		}
	})
	t.Run("pythagoras.gif", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "pythagoras.gif")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, true, f.IsImage())
			assert.Equal(t, false, f.IsVideo())
			assert.Equal(t, true, f.IsAnimated())
			assert.Equal(t, true, f.IsGif())
			assert.Equal(t, true, f.IsAnimatedGif())
			assert.Equal(t, false, f.IsSidecar())
		}
	})
	t.Run("christmas.mp4", func(t *testing.T) {
		if f, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "christmas.mp4")); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, false, f.IsImage())
			assert.Equal(t, true, f.IsVideo())
			assert.Equal(t, true, f.IsAnimated())
			assert.Equal(t, false, f.IsGif())
			assert.Equal(t, false, f.IsAnimatedGif())
			assert.Equal(t, false, f.IsSidecar())
		}
	})
}

func TestMediaFile_HasJpeg(t *testing.T) {
	t.Run("Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		f, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")

		if err != nil {
			t.Fatal(err)
		}

		assert.False(t, f.HasJpeg())
	})
	t.Run("ferriswheel_colorful.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		f, err := NewMediaFile(conf.ExamplesPath() + "/ferriswheel_colorful.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, f.HasJpeg())
	})
	t.Run("Random.docx with jpg", func(t *testing.T) {
		conf := config.TestConfig()

		f, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")
		f.hasJpeg = true
		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, f.HasJpeg())
	})
}

func TestMediaFile_Jpeg(t *testing.T) {
	t.Run("Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")

		if err != nil {
			t.Fatal(err)
		}

		file, err := mediaFile.Jpeg()

		if file != nil {
			t.Fatal("file should be nil")
		}

		if err == nil {
			t.Fatal("err should NOT be nil")
		}

		assert.Equal(t, "no jpeg found for "+mediaFile.FileName(), err.Error())
	})
	t.Run("ferriswheel_colorful.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/ferriswheel_colorful.jpg")

		if err != nil {
			t.Fatal(err)
		}

		file, err := mediaFile.Jpeg()

		if err != nil {
			t.Fatal(err)
		}

		assert.FileExists(t, file.fileName)
	})
	t.Run("iphone_7.json", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/test.md")

		if err != nil {
			t.Fatal(err)
		}

		file, err := mediaFile.Jpeg()

		if file != nil {
			t.Fatal("file should be nil")
		}

		if err == nil {
			t.Fatal("err should NOT be nil")
		}

		assert.Equal(t, "no jpeg found for "+mediaFile.FileName(), err.Error())
	})
}

func TestMediaFile_decodeDimension(t *testing.T) {
	t.Run("Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")

		if err != nil {
			t.Fatal(err)
		}

		decodeErr := mediaFile.decodeDimensions()

		assert.EqualError(t, decodeErr, "failed decoding dimensions of .docx file")
	})

	t.Run("clock_purple.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/clock_purple.jpg")

		if err != nil {
			t.Fatal(err)
		}

		if err := mediaFile.decodeDimensions(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

		if err != nil {
			t.Fatal(err)
		}

		if err := mediaFile.decodeDimensions(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("example.png", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/example.png")

		if err != nil {
			t.Fatal(err)
		}

		if err := mediaFile.decodeDimensions(); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 100, mediaFile.Width())
		assert.Equal(t, 67, mediaFile.Height())
	})

	t.Run("example.gif", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/example.gif")

		if err != nil {
			t.Fatal(err)
		}

		if err := mediaFile.decodeDimensions(); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 100, mediaFile.Width())
		assert.Equal(t, 67, mediaFile.Height())
	})

	t.Run("blue-go-video.mp4", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")

		if err != nil {
			t.Fatal(err)
		}

		if err := mediaFile.decodeDimensions(); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 1920, mediaFile.Width())
		assert.Equal(t, 1080, mediaFile.Height())
	})
	t.Run("blue-go-video.mp4 with orientation >4 and <8", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")
		mediaFile.metaData.Orientation = 5
		if err != nil {
			t.Fatal(err)
		}

		if err := mediaFile.decodeDimensions(); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 1080, mediaFile.Width())
		assert.Equal(t, 1920, mediaFile.Height())
	})
}

func TestMediaFile_Width(t *testing.T) {
	t.Run("Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")
		if err != nil {
			t.Fatal(err)
		}
		width := mediaFile.Width()
		assert.Equal(t, 0, width)
	})
	t.Run("elephant_mono.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephant_mono.jpg")
		if err != nil {
			t.Fatal(err)
		}
		width := mediaFile.Width()
		assert.Equal(t, 416, width)
	})
}

func TestMediaFile_Height(t *testing.T) {
	t.Run("Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx")

		if err != nil {
			t.Fatal(err)
		}

		height := mediaFile.Height()
		assert.Equal(t, 0, height)
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")

		if err != nil {
			t.Fatal(err)
		}

		height := mediaFile.Height()
		assert.Equal(t, 331, height)
	})
}

func TestMediaFile_Megapixels(t *testing.T) {
	conf := config.TestConfig()

	t.Run("Random.docx", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/Random.docx"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("elephant_mono.jpg", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/elephant_mono.jpg"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("telegram_2020-01-30_09-57-18.jpg", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/telegram_2020-01-30_09-57-18.jpg"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 1, f.Megapixels())
		}
	})
	t.Run("6720px_white.jpg", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/6720px_white.jpg"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 30, f.Megapixels())
		}
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("example.bmp", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/example.bmp"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("panorama360.jpg", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/panorama360.jpg"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("panorama360.json", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/panorama360.json"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("2018-04-12 19_24_49.gif", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/2018-04-12 19_24_49.gif"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("2018-04-12 19_24_49.mov", func(t *testing.T) {
		if _, err := NewMediaFile("testdata/2018-04-12 19_24_49.mov"); err != nil {
			assert.EqualError(t, err, "'testdata/2018-04-12 19_24_49.mov' is empty")
		} else {
			t.Errorf("error expected")
		}
	})
	t.Run("rotate/6.png", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/rotate/6.png"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 1, f.Megapixels())
		}
	})
	t.Run("rotate/6.tiff", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/rotate/6.tiff"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
	t.Run("norway-kjetil-moe.webp", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/norway-kjetil-moe.webp"); err != nil {
			t.Fatal(err)
		} else {
			assert.Equal(t, 0, f.Megapixels())
		}
	})
}

func TestMediaFile_ExceedsFileSize(t *testing.T) {
	t.Run("norway-kjetil-moe.webp", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/norway-kjetil-moe.webp"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsFileSize(3)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
	t.Run("telegram_2020-01-30_09-57-18.jpg", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/telegram_2020-01-30_09-57-18.jpg"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsFileSize(-1)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
	t.Run("6720px_white.jpg", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/6720px_white.jpg"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsFileSize(0)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsFileSize(10)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
	t.Run("example.bmp", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/example.bmp"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsFileSize(10)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
}

func TestMediaFile_DecodeConfig(t *testing.T) {
	t.Run("6720px_white.jpg", func(t *testing.T) {
		f, err := NewMediaFile(conf.ExamplesPath() + "/6720px_white.jpg")

		if err != nil {
			t.Fatal(err)
		}

		cfg1, err1 := f.DecodeConfig()

		assert.Nil(t, err1)
		assert.IsType(t, &image.Config{}, cfg1)
		assert.Equal(t, 6720, cfg1.Width)
		assert.Equal(t, 4480, cfg1.Height)

		cfg2, err2 := f.DecodeConfig()

		assert.Nil(t, err2)
		assert.IsType(t, &image.Config{}, cfg2)
		assert.Equal(t, 6720, cfg2.Width)
		assert.Equal(t, 4480, cfg2.Height)

		cfg3, err3 := f.DecodeConfig()

		assert.Nil(t, err3)
		assert.IsType(t, &image.Config{}, cfg3)
		assert.Equal(t, 6720, cfg3.Width)
		assert.Equal(t, 4480, cfg3.Height)
	})
}

func TestMediaFile_ExceedsResolution(t *testing.T) {
	t.Run("norway-kjetil-moe.webp", func(t *testing.T) {
		if f, err := NewMediaFile("testdata/norway-kjetil-moe.webp"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsResolution(3)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
	t.Run("telegram_2020-01-30_09-57-18.jpg", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/telegram_2020-01-30_09-57-18.jpg"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsResolution(3)
			assert.False(t, result)
			assert.Equal(t, 1, actual)
		}
	})
	t.Run("6720px_white.jpg", func(t *testing.T) {
		f, err := NewMediaFile(conf.ExamplesPath() + "/6720px_white.jpg")

		if err != nil {
			t.Fatal(err)
		}

		exceeds3, actual3 := f.ExceedsResolution(3)

		assert.True(t, exceeds3)
		assert.Equal(t, 30, actual3)

		exceeds30, actual30 := f.ExceedsResolution(30)

		assert.False(t, exceeds30)
		assert.Equal(t, 30, actual30)

		exceeds33, actual33 := f.ExceedsResolution(33)

		assert.False(t, exceeds33)
		assert.Equal(t, 30, actual33)
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.dng"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsResolution(3)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
	t.Run("example.bmp", func(t *testing.T) {
		if f, err := NewMediaFile(conf.ExamplesPath() + "/example.bmp"); err != nil {
			t.Fatal(err)
		} else {
			result, actual := f.ExceedsResolution(3)
			assert.False(t, result)
			assert.Equal(t, 0, actual)
		}
	})
}

func TestMediaFile_AspectRatio(t *testing.T) {
	t.Run("iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

		if err != nil {
			t.Fatal(err)
		}

		ratio := mediaFile.AspectRatio()
		assert.Equal(t, float32(0.75), ratio)
	})
	t.Run("fern_green.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/fern_green.jpg")

		if err != nil {
			t.Fatal(err)
		}

		ratio := mediaFile.AspectRatio()
		assert.Equal(t, float32(1), ratio)
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")

		if err != nil {
			t.Fatal(err)
		}

		ratio := mediaFile.AspectRatio()
		assert.Equal(t, float32(1.5), ratio)
	})
}

func TestMediaFile_Orientation(t *testing.T) {
	t.Run("iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

		if err != nil {
			t.Fatal(err)
		}

		orientation := mediaFile.Orientation()
		assert.Equal(t, 6, orientation)
	})
	t.Run("turtle_brown_blue.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/turtle_brown_blue.jpg")

		if err != nil {
			t.Fatal(err)
		}

		orientation := mediaFile.Orientation()
		assert.Equal(t, 1, orientation)
	})
}

func TestMediaFile_Thumbnail(t *testing.T) {
	conf := config.TestConfig()

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)

	t.Run("elephants.jpg", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")

		if err != nil {
			t.Fatal(err)
		}

		thumbnail, err := image.Thumbnail(thumbsPath, "tile_500")

		if err != nil {
			t.Fatal(err)
		}

		assert.FileExists(t, thumbnail)
	})
	t.Run("invalid image format", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/canon_eos_6d.xmp")

		if err != nil {
			t.Fatal(err)
		}

		thumbnail, err := image.Thumbnail(thumbsPath, "tile_500")

		assert.EqualError(t, err, "media: failed creating thumbnail for canon_eos_6d.xmp (image: unknown format)")

		t.Log(thumbnail)
	})
	t.Run("invalid thumbnail type", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")

		if err != nil {
			t.Fatal(err)
		}

		thumbnail, err := image.Thumbnail(thumbsPath, "invalid_500")

		assert.EqualError(t, err, "media: invalid type invalid_500")

		t.Log(thumbnail)
	})
}

func TestMediaFile_Resample(t *testing.T) {
	conf := config.TestConfig()

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)
	t.Run("elephants.jpg", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")

		if err != nil {
			t.Fatal(err)
		}

		thumbnail, err := image.Resample(thumbsPath, thumb.Tile500)

		if err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, thumbnail)

	})
	t.Run("invalid type", func(t *testing.T) {
		image, err := NewMediaFile(conf.ExamplesPath() + "/elephants.jpg")

		if err != nil {
			t.Fatal(err)
		}

		thumbnail, err := image.Resample(thumbsPath, "xxx_500")

		if err == nil {
			t.Fatal("err should not be nil")
		}

		assert.Equal(t, "media: invalid type xxx_500", err.Error())
		assert.Empty(t, thumbnail)
	})

}

func TestMediaFile_RenderDefaultThumbs(t *testing.T) {
	conf := config.TestConfig()

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	m, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "elephants.jpg"))

	if err != nil {
		t.Fatal(err)
	}

	err = m.CreateThumbnails(thumbsPath, true)

	if err != nil {
		t.Fatal(err)
	}

	thumbFilename, err := thumb.FileName(m.Hash(), thumbsPath, thumb.Sizes[thumb.Tile50].Width, thumb.Sizes[thumb.Tile50].Height, thumb.Sizes[thumb.Tile50].Options...)

	if err != nil {
		t.Fatal(err)
	}

	assert.FileExists(t, thumbFilename)

	err = m.CreateThumbnails(thumbsPath, false)

	assert.Empty(t, err)
}

func TestMediaFile_FileType(t *testing.T) {
	m, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "this-is-a-jpeg.png"))

	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, m.IsJpeg())
	assert.Equal(t, "jpg", string(m.FileType()))
	assert.Equal(t, fs.ImageJPEG, m.FileType())
	assert.Equal(t, ".png", m.Extension())
}

func TestMediaFile_Stat(t *testing.T) {
	t.Run("iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

		if err != nil {
			t.Fatal(err)
		}

		size, time, err := mediaFile.Stat()

		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, int64(785743), size)
		assert.IsType(t, time, time)
	})
}

func TestMediaFile_FileSize(t *testing.T) {
	t.Run("iphone_7.heic", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/iphone_7.heic")

		if err != nil {
			t.Fatal(err)
		}

		size := mediaFile.FileSize()
		assert.Equal(t, int64(785743), size)
	})
}

func TestMediaFile_JsonName(t *testing.T) {
	t.Run("blue-go-video.mp4", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")

		if err != nil {
			t.Fatal(err)
		}

		name := mediaFile.SidecarJsonName()
		assert.True(t, strings.HasSuffix(name, "/assets/examples/blue-go-video.mp4.json"))
	})
}

func TestMediaFile_PathNameInfo(t *testing.T) {
	t.Run("blue-go-video.mp4", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")

		if err != nil {
			t.Fatal(err)
		}

		root, base, path, name := mediaFile.PathNameInfo(true)
		assert.Equal(t, "examples", root)
		assert.Equal(t, "blue-go-video", base)
		assert.Equal(t, "", path)
		assert.Equal(t, "blue-go-video.mp4", name)

	})

	t.Run("beach_sand sidecar", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_sand.jpg")

		if err != nil {
			t.Fatal(err)
		}

		initialName := mediaFile.FileName()
		mediaFile.SetFileName(".photoprism/beach_sand.jpg")

		root, base, path, name := mediaFile.PathNameInfo(true)
		assert.Equal(t, "", root)
		assert.Equal(t, "beach_sand", base)
		assert.Equal(t, "", path)
		assert.Equal(t, ".photoprism/beach_sand.jpg", name)
		mediaFile.SetFileName(initialName)
	})

	t.Run("beach_sand import", func(t *testing.T) {
		conf := config.TestConfig()
		t.Log(Config().SidecarPath())
		t.Log(Config().ImportPath())

		mediaFile, err := NewMediaFile(filepath.Join(conf.ExamplesPath(), "beach_sand.jpg"))

		if err != nil {
			t.Fatal(err)
		}

		initialName := mediaFile.FileName()
		t.Log(initialName)
		mediaFile.SetFileName(filepath.Join(conf.ImportPath(), "beach_sand.jpg"))

		root, base, path, name := mediaFile.PathNameInfo(true)
		assert.Equal(t, "import", root)
		assert.Equal(t, "beach_sand", base)
		assert.Equal(t, "", path)
		assert.Equal(t, "beach_sand.jpg", name)
		mediaFile.SetFileName(initialName)
	})

	t.Run("beach_sand unknown root", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/beach_sand.jpg")

		if err != nil {
			t.Fatal(err)
		}

		initialName := mediaFile.FileName()
		mediaFile.SetFileName("/go/src/github.com/photoprism/notExisting/xxx/beach_sand.jpg")

		root, base, path, name := mediaFile.PathNameInfo(false)
		assert.Equal(t, "", root)
		assert.Equal(t, "beach_sand", base)
		assert.Equal(t, "/go/src/github.com/photoprism/notExisting/xxx", path)
		assert.Equal(t, "/go/src/github.com/photoprism/notExisting/xxx/beach_sand.jpg", name)
		mediaFile.SetFileName(initialName)
	})
}

func TestMediaFile_SubDirectory(t *testing.T) {
	t.Run("blue-go-video.mp4", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")

		if err != nil {
			t.Fatal(err)
		}

		subdir := mediaFile.SubDir("xxx")
		assert.True(t, strings.HasSuffix(subdir, "/assets/examples/xxx"))
	})
}

func TestMediaFile_HasSameName(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")

		if err != nil {
			t.Fatal(err)
		}

		mediaFile2, err := NewMediaFile(conf.ExamplesPath() + "/beach_sand.jpg")

		if err != nil {
			t.Fatal(err)
		}

		assert.False(t, mediaFile.HasSameName(nil))
		assert.False(t, mediaFile.HasSameName(mediaFile2))

	})
}

func TestMediaFile_IsJson(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")

		if err != nil {
			t.Fatal(err)
		}

		assert.False(t, mediaFile.IsJson())
	})
	t.Run("true", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.json")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, mediaFile.IsJson())
	})
}

func TestMediaFile_IsPlayableVideo(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.json")

		if err != nil {
			t.Fatal(err)
		}

		assert.False(t, mediaFile.IsPlayableVideo())
	})
	t.Run("true", func(t *testing.T) {
		conf := config.TestConfig()

		mediaFile, err := NewMediaFile(conf.ExamplesPath() + "/blue-go-video.mp4")

		if err != nil {
			t.Fatal(err)
		}

		assert.True(t, mediaFile.IsPlayableVideo())
	})
}

func TestMediaFile_RenameSidecars(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		conf := config.TestConfig()

		jpegExample := filepath.Join(conf.ExamplesPath(), "/limes.jpg")
		jpegPath := filepath.Join(conf.OriginalsPath(), "2020", "12")
		jpegName := filepath.Join(jpegPath, "foobar.jpg")

		if err := fs.Copy(jpegExample, jpegName); err != nil {
			t.Fatal(err)
		}

		mf, err := NewMediaFile(jpegName)

		if err != nil {
			t.Fatal(err)
		}

		if err := os.MkdirAll(filepath.Join(conf.SidecarPath(), "foo"), os.ModePerm); err != nil {
			t.Fatal(err)
		}

		srcName := filepath.Join(conf.SidecarPath(), "foo/bar.jpg.json")
		dstName := filepath.Join(conf.SidecarPath(), "2020/12/foobar.jpg.json")

		if err := os.WriteFile(srcName, []byte("{}"), 0666); err != nil {
			t.Fatal(err)
		}

		if renamed, err := mf.RenameSidecars(filepath.Join(conf.OriginalsPath(), "foo/bar.jpg")); err != nil {
			t.Fatal(err)
		} else if len(renamed) != 1 {
			t.Errorf("len should be 2: %#v", renamed)
		} else {
			t.Logf("renamed: %#v", renamed)
		}

		if fs.FileExists(srcName) {
			t.Errorf("src file still exists: %s", srcName)
		}

		if !fs.FileExists(dstName) {
			t.Errorf("dst file not found: %s", srcName)
		}

		_ = os.Remove(srcName)
		_ = os.Remove(dstName)
	})
}

func TestMediaFile_RemoveSidecars(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		conf := config.TestConfig()

		jpegExample := filepath.Join(conf.ExamplesPath(), "/limes.jpg")
		jpegPath := filepath.Join(conf.OriginalsPath(), "2020", "12")
		jpegName := filepath.Join(jpegPath, "foobar.jpg")

		if err := fs.Copy(jpegExample, jpegName); err != nil {
			t.Fatal(err)
		}

		mf, err := NewMediaFile(jpegName)

		if err != nil {
			t.Fatal(err)
		}

		sidecarName := filepath.Join(conf.SidecarPath(), "2020/12/foobar.jpg.json")

		if err := os.WriteFile(sidecarName, []byte("{}"), 0666); err != nil {
			t.Fatal(err)
		}

		if err := mf.RemoveSidecars(); err != nil {
			t.Fatal(err)
		} else if fs.FileExists(sidecarName) {
			t.Errorf("src file still exists: %s", sidecarName)
		}

		_ = os.Remove(sidecarName)
	})
}

func TestMediaFile_ColorProfile(t *testing.T) {
	c := config.TestConfig()

	t.Run("iphone_7.json", func(t *testing.T) {
		mediaFile, err := NewMediaFile(c.ExamplesPath() + "/iphone_7.json")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.ColorProfile())
	})
	t.Run("iphone_7.xmp", func(t *testing.T) {
		mediaFile, err := NewMediaFile(c.ExamplesPath() + "/iphone_7.xmp")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.ColorProfile())
	})
	t.Run("iphone_7.heic", func(t *testing.T) {
		mediaFile, err := NewMediaFile(c.ExamplesPath() + "/iphone_7.heic")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.ColorProfile())
	})
	t.Run("canon_eos_6d.dng", func(t *testing.T) {
		mediaFile, err := NewMediaFile(c.ExamplesPath() + "/canon_eos_6d.dng")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.ColorProfile())
	})
	t.Run("elephants.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(c.ExamplesPath() + "/elephants.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Adobe RGB (1998)", mediaFile.ColorProfile())
	})
	t.Run("/beach_wood.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(c.ExamplesPath() + "/beach_wood.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "", mediaFile.ColorProfile())
	})
	t.Run("/peacock_blue.jpg", func(t *testing.T) {
		mediaFile, err := NewMediaFile(c.ExamplesPath() + "/peacock_blue.jpg")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "sRGB IEC61966-2.1", mediaFile.ColorProfile())
	})
}
