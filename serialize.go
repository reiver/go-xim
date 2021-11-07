package iid

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"io"
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

	var storage bytes.Buffer
	{

		err := binary.Write(&storage, binary.BigEndian, value)
		if nil != err {
			return ""
		}
	}

	var encoded string
	{
		encoded = base64encoding.EncodeToString(storage.Bytes())
	}

	return encoded
}

func unserialize(value string) (uint64, bool) {

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