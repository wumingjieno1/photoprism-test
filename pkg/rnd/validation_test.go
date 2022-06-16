package rnd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUUID(t *testing.T) {
	assert.True(t, ValidUUID("dafbfeb8-a129-4e7c-9cf0-e7996a701cdb"))
	assert.True(t, ValidUUID("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	assert.False(t, ValidUUID("55785BAC-9H4B-4747-B090-EE123FFEE437"))
	assert.True(t, ValidUUID("550e8400-e29b-11d4-a716-446655440000"))
	assert.False(t, ValidUUID("4B1FEF2D1CF4A5BE38B263E0637EDEAD"))
}

func TestSanitizeUUID(t *testing.T) {
	assert.Equal(t, "dafbfeb8-a129-4e7c-9cf0-e7996a701cdb", SanitizeUUID("  \"dafbfeb8-a129-4e7c-9cf0-e7996a701cdb\"  "))
	assert.Equal(t, "dafbfeb8-a129-4e7c-9cf0-e7996a701cdb", SanitizeUUID("  xmp:dafbfeb8-a129-4e7c-9cf0-e7996a701cdb  "))
	assert.Equal(t, "dafbfeb8-a129-4e7c-9cf0-e7996a701cdb", SanitizeUUID("dafbfeb8-a129-4e7c-9cf0-e7996a701cdb"))
	assert.Equal(t, "6ba7b810-9dad-11d1-80b4-00c04fd430c8", SanitizeUUID("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	assert.Equal(t, "", SanitizeUUID("55785BAC-9H4B-4747-B090-EE123FFEE437"))
	assert.Equal(t, "550e8400-e29b-11d4-a716-446655440000", SanitizeUUID("550e8400-e29b-11d4-a716-446655440000"))
	assert.Equal(t, "", SanitizeUUID("4B1FEF2D1CF4A5BE38B263E0637EDEAD"))
	assert.Equal(t, "", SanitizeUUID(""))
}
