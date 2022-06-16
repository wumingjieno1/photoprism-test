package video

import "testing"

func TestCodecs(t *testing.T) {
	if val := Codecs[""]; val != UnknownCodec {
		t.Fatal("default codec should be UnknownCodec")
	}

	if val := Codecs["avc"]; val != CodecAVC {
		t.Fatal("codec should be CodecAVC")
	}

	if val := Codecs["av1"]; val != CodecAV1 {
		t.Fatal("codec should be CodecAV1")
	}
}
