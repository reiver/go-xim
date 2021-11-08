package iid

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"io"
	"strings"
)

const (
	serializationprefix = "x"
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

	return encoded.String()
}

func unserialize(value string) (uint64, bool) {

	{
		if !strings.HasPrefix(value, serializationprefix) {
			return badvalue, false
		}

		value = value[len(serializationprefix):]
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
		var reader io.Reader = bytes.NewReader(p)

		err := binary.Read(reader, binary.BigEndian, &result)
		if nil != err {
			return badvalue, false
		}
	}

	return result, true
}
