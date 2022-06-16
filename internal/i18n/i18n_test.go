package i18n

import (
	"os"
	"testing"

	"github.com/leonelquinteros/gotext"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gotext.Configure(localeDir, string(locale), "default")

	code := m.Run()

	os.Exit(code)
}

func TestMsg(t *testing.T) {
	t.Run("already exists", func(t *testing.T) {
		msg := Msg(ErrAlreadyExists, "A cat")
		assert.Equal(t, "A cat already exists", msg)
	})

	t.Run("unexpected error", func(t *testing.T) {
		msg := Msg(ErrUnexpected, "A cat")
		assert.Equal(t, "Unexpected error, please try again", msg)
	})

	t.Run("already exists german", func(t *testing.T) {
		SetLocale("de")
		msgTrans := Msg(ErrAlreadyExists, "Eine Katze")
		assert.Equal(t, "Eine Katze existiert bereits", msgTrans)
		SetLocale("")
		msgDefault := Msg(ErrAlreadyExists, "A cat")
		assert.Equal(t, "A cat already exists", msgDefault)
	})

	t.Run("already exists polish", func(t *testing.T) {
		SetLocale("pl")
		msgTrans := Msg(ErrAlreadyExists, "Kot")
		assert.Equal(t, "Kot już istnieje", msgTrans)
		SetLocale("")
		msgDefault := Msg(ErrAlreadyExists, "A cat")
		assert.Equal(t, "A cat already exists", msgDefault)
	})

	t.Run("Brazilian Portuguese", func(t *testing.T) {
		SetLocale("pt_BR")
		msgTrans := Msg(ErrAlreadyExists, "Gata")
		assert.Equal(t, "Gata já existe", msgTrans)
		SetLocale("")
		msgDefault := Msg(ErrAlreadyExists, "A cat")
		assert.Equal(t, "A cat already exists", msgDefault)
	})
}

func TestError(t *testing.T) {
	t.Run("already exists", func(t *testing.T) {
		err := Error(ErrAlreadyExists, "A cat")
		assert.EqualError(t, err, "a cat already exists")
	})

	t.Run("unexpected error", func(t *testing.T) {
		err := Error(ErrUnexpected, "A cat")
		assert.EqualError(t, err, "unexpected error, please try again")
	})

	t.Run("already exists german", func(t *testing.T) {
		SetLocale("de")
		errGerman := Error(ErrAlreadyExists, "Eine Katze")
		assert.EqualError(t, errGerman, "eine katze existiert bereits")
		SetLocale("")
		errDefault := Error(ErrAlreadyExists, "A cat")
		assert.EqualError(t, errDefault, "a cat already exists")
	})
}
