package entity

import (
	"testing"

	"github.com/photoprism/photoprism/pkg/rnd"
	"github.com/stretchr/testify/assert"
)

func TestNewLink(t *testing.T) {
	link := NewLink("st9lxuqxpogaaba1", true, false)
	assert.Equal(t, "st9lxuqxpogaaba1", link.ShareUID)
	assert.Equal(t, false, link.CanEdit)
	assert.Equal(t, true, link.CanComment)
	assert.Equal(t, 10, len(link.LinkToken))
	assert.Equal(t, 16, len(link.LinkUID))
}

func TestLink_Expired(t *testing.T) {
	const oneDay = 60 * 60 * 24

	link := NewLink("st9lxuqxpogaaba1", true, false)

	link.ModifiedAt = TimeStamp().Add(-7 * Day)
	link.LinkExpires = 0

	assert.False(t, link.Expired())

	link.LinkExpires = oneDay

	assert.True(t, link.Expired())

	link.LinkExpires = oneDay * 8

	assert.False(t, link.Expired())

	link.LinkExpires = oneDay * 300
	link.LinkViews = 9
	link.MaxViews = 10

	assert.False(t, link.Expired())

	link.Redeem()

	assert.True(t, link.Expired())
}

func TestLink_Redeem(t *testing.T) {
	link := NewLink(rnd.GenerateUID('a'), false, false)

	assert.Equal(t, uint(0), link.LinkViews)

	link.Redeem()

	assert.Equal(t, uint(1), link.LinkViews)

	if err := link.Save(); err != nil {
		t.Fatal(err)
	}

	link.Redeem()

	assert.Equal(t, uint(2), link.LinkViews)
}

func TestLink_SetSlug(t *testing.T) {
	link := Link{}
	assert.Equal(t, "", link.ShareSlug)
	link.SetSlug("test Slug")
	assert.Equal(t, "test-slug", link.ShareSlug)
}

func TestLink_SetPassword(t *testing.T) {
	link := Link{LinkUID: "dftjdfkvh"}
	assert.Equal(t, false, link.HasPassword)
	err := link.SetPassword("123")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, link.HasPassword)
}

func TestLink_InvalidPassword(t *testing.T) {
	t.Run("no password", func(t *testing.T) {
		link := Link{LinkUID: "dftjdfkvhjh", HasPassword: false}
		assert.False(t, link.InvalidPassword("123"))
	})
	t.Run("invalid password", func(t *testing.T) {
		link := NewLink("dhfjf", false, false)

		err := link.SetPassword("123")
		if err != nil {
			t.Fatal(err)
		}
		assert.False(t, link.InvalidPassword("123"))
	})
	t.Run("valid password", func(t *testing.T) {
		link := NewLink("dhfjfk", false, false)

		err := link.SetPassword("123kkljgfuA")
		if err != nil {
			t.Fatal(err)
		}
		assert.True(t, link.InvalidPassword("123"))
	})
}

func TestLink_Save(t *testing.T) {
	t.Run("invalid share uid", func(t *testing.T) {
		link := NewLink("dhfjfjh", false, false)

		assert.Error(t, link.Save())
	})
	t.Run("empty token", func(t *testing.T) {
		link := Link{ShareUID: "lpfjfjhffgtredft", LinkToken: ""}

		assert.Error(t, link.Save())
	})
	t.Run("success", func(t *testing.T) {
		link := NewLink("lhfjfjhffgtredft", false, false)

		err := link.Save()

		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestLink_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		link := NewLink("lhfjfjhffgtreoft", false, false)

		err := link.Delete()

		if err != nil {
			t.Fatal(err)
		}

	})
	t.Run("empty token", func(t *testing.T) {
		link := Link{ShareUID: "lpfjpjhffgtredft", LinkToken: ""}
		assert.Error(t, link.Delete())
	})
}

func TestFindLink(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		m := NewLink("lhfjfjhffgtrjoft", false, false)

		link := &m

		if err := link.Save(); err != nil {
			t.Fatal(err)
		}
		uid := link.LinkUID
		t.Logf("%#v", link)
		r := FindLink(uid)
		t.Log(r)
		//TODO Why does it fail?
		//assert.Equal(t, "1jxf3jfn2k", r.LinkToken)
	})
	t.Run("nil", func(t *testing.T) {
		r := FindLink("XXX")
		assert.Nil(t, r)
	})
}

func TestFindLinks(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := FindLinks("1jxf3jfn2k", "")
		assert.Equal(t, "at9lxuqxpogaaba8", r[0].ShareUID)
	})
	t.Run("not found", func(t *testing.T) {
		r := FindLinks("", "")
		assert.Empty(t, r)
	})
	t.Run("not found", func(t *testing.T) {
		r := FindLinks("lkjh", "")
		assert.Empty(t, r)
	})
}

func TestFindValidLinksLinks(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := FindValidLinks("1jxf3jfn2k", "")
		assert.Equal(t, "at9lxuqxpogaaba8", r[0].ShareUID)
	})
}

func TestLink_String(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		link := NewLink("jhgko", false, false)
		uid := link.LinkUID
		assert.Equal(t, uid, link.String())
	})
}
