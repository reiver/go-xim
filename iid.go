package iid

// IID represents an interaction‐identifier — i.e., an interaction‐ID — i.e., an IID.
//
// It is a scheme used to generate quazi‐monotonically‐increasing‐unique.
type IID struct {
	value uint64
	loaded bool
}

func Nothing() IID {
	return IID{}
}

func something(value uint64) IID {
	return IID{
		value:value,
		loaded:true,
	}
}

// Generate creates a new interaction‐identifier — i.e., an interaction‐ID — i.e., an IID.
func Generate() IID {

	var value uint64 = generate()

	return something(value)
}

// Chaos returns the randomness that is embeddd in the interaction‐identifier — i.e., an interaction‐ID — i.e., an IID.
func (receiver IID) Chaos() (uint64, bool) {
	if Nothing() == receiver {
		return 0, false
	}

	_, value := decompile(receiver.value)

	return value, true
}

func (receiver IID) String() string {
	if Nothing() == receiver {
		return ""
	}

	var serialized string = serialize(receiver.value)

	return serialized
}

func (receiver IID) MarshalText() (text []byte, err error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	var serialized string = serialize(receiver.value)

	return []byte(serialized), nil
}

// UnixTime returns the unix‐time that is embeddd in the interaction‐identifier — i.e., an interaction‐ID — i.e., an IID.
func (receiver IID) UnixTime() (int64, bool) {
	if Nothing() == receiver {
		return 0, false
	}

	value, _ := decompile(receiver.value)

	return int64(value), true
}

func (receiver *IID) UnmarshalText(p []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var serialized string = string(p)

	value, successful := unserialize(serialized)
	if !successful {
		return  errBadRequest
	}

	*receiver = something(value)
	return nil
}
