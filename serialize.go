package xim

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"io"
	"strings"
)

const (
	serializationprefix = "xi"
	serializationsuffix = "m"
)

var (
	base64encoding = base64.NewEncoding(
		"-"+
		"0123456789"+
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ"+
		"_"+
		"abcdefghijklmnopqrstuvwxyz",
	).WithPadding(base64.NoPadding)
)

func serialize(value uint64) string {

	var binstorage bytes.Buffer
	{
		err := binary.Write(&binstorage, binary.BigEndian, value)
		if nil != err {
			return ""
		}
	}

	var encoded strings.Builder
	{
		encoded.WriteString(serializationprefix)
	}
	{
		var wc io.WriteCloser = base64.NewEncoder(base64encoding, &encoded)
		if nil == wc {
			return ""
		}
		wc.Write(binstorage.Bytes())
		wc.Close()

	}
	{
		encoded.WriteString(serializationsuffix)
	}

	return encoded.String()
}

func unserialize(value string) (uint64, bool) {

	{
		if !strings.HasPrefix(value, serializationprefix) {
			return badvalue, false
		}

		value = value[len(serializationprefix):]
	}

	{
		if !strings.HasSuffix(value, serializationsuffix) {
			return badvalue, false
		}

		value = value[:len(value)-len(serializationsuffix)]
	}

	var p []byte
	{
		var err error

		p, err = base64encoding.DecodeString(value)
		if nil != err {
			return badvalue, false
		}
	}

	var result uint64
	{
		var reader *bytes.Reader = bytes.NewReader(p)

		err := binary.Read(reader, binary.BigEndian, &result)
		if nil != err {
			return badvalue, false
		}

		if expected, actual := 0, reader.Len(); expected != actual {
			return badvalue, false
		}
	}

	return result, true
}
