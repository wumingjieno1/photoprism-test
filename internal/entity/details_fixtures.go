package entity

type DetailsMap map[string]Details

func (m DetailsMap) Get(name string, photoId uint) Details {
	if result, ok := m[name]; ok {
		result.PhotoID = photoId
		return result
	}

	return Details{PhotoID: photoId}
}

func (m DetailsMap) Pointer(name string, photoId uint) *Details {
	if result, ok := m[name]; ok {
		result.PhotoID = photoId
		return &result
	}

	return &Details{PhotoID: photoId}
}

var DetailsFixtures = DetailsMap{
	"lake": {
		PhotoID:      1000000,
		Keywords:     "nature, frog",
		Notes:        "notes",
		Subject:      "Lake",
		Artist:       "Hans",
		Copyright:    "copy",
		License:      "MIT",
		CreatedAt:    TimeStamp(),
		UpdatedAt:    TimeStamp(),
		KeywordsSrc:  "meta",
		NotesSrc:     "manual",
		SubjectSrc:   "meta",
		ArtistSrc:    "meta",
		CopyrightSrc: "manual",
		LicenseSrc:   "manual",
	},
	"blacklist": {
		PhotoID:      1000001,
		Keywords:     "screenshot, info",
		Notes:        "notes",
		Subject:      "Blacklist",
		Artist:       "Hans",
		Copyright:    "copy",
		License:      "MIT",
		CreatedAt:    TimeStamp(),
		UpdatedAt:    TimeStamp(),
		KeywordsSrc:  "",
		NotesSrc:     "",
		SubjectSrc:   "meta",
		ArtistSrc:    "meta",
		CopyrightSrc: "manual",
		LicenseSrc:   "manual",
	},
	"bridge": {
		PhotoID:      1000003,
		Keywords:     "bridge, nature",
		Notes:        "Some Notes!@#$",
		Subject:      "Bridge",
		Artist:       "Jens Mander",
		Copyright:    "Copyright 2020",
		License:      "n/a",
		CreatedAt:    TimeStamp(),
		UpdatedAt:    TimeStamp(),
		KeywordsSrc:  "meta",
		NotesSrc:     "manual",
		SubjectSrc:   "meta",
		ArtistSrc:    "meta",
		CopyrightSrc: "manual",
		LicenseSrc:   "manual",
	},
}
