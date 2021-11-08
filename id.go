package xim

// ID represents a xim-id.
//
// It is a scheme used to generate quazi‐monotonically‐increasing‐unique identifier.
//
// This can be used for anything that needs a unique-identifier. And its serialization should be safe to use as a file-name, or a directory-name.
//
// Example usage:
//
//	var id xim.ID = xim.Generate()
//	
//	fmt.Println("xim-id =", id)
type ID struct {
	value uint64
	loaded bool
}

// ID is an option-type. Nothing returns the nothing value for the option-type.
//
// An example usage might be:
//
//	if xim.Nothing() == id {
//		//@TODO
//	}
func Nothing() ID {
	return ID{}
}

// ID is an option-type. something returns a something value for the option-type, that contains the value of the input variable ‘value’ inside of it.
func something(value uint64) ID {
	return ID{
		value:value,
		loaded:true,
	}
}

// Generate returns a new xim-id.
//
// Example usage:
//
//	var id xim.ID = xim.Generate()
//	
//	fmt.Println("xim-id =", id)
func Generate() ID {

	var value uint64 = generate()

	return something(value)
}

// Chaos returns the randomness that is embeddd in the xim-id.
//
// Example usage:
//
//	var id xim.ID = xim.Generate()
//	
//	chaos, successful := id.Chaos()
//	if !successful {
//		return errors.New("xim-id was not initialized")
//	}
//	
//	fmt.Printf("chaos = %#020b \n", chaos)
func (receiver ID) Chaos() (uint64, bool) {
	if Nothing() == receiver {
		return 0, false
	}

	_, value := decompile(receiver.value)

	return value, true
}

// String makes xim.ID fit the fmt.Stringer interface.
//
// String also returns the serialized for of a xim-id, in xim-notation.
func (receiver ID) String() string {
	if Nothing() == receiver {
		return ""
	}

	var serialized string = serialize(receiver.value)

	return serialized
}

// UnmarshalText makes xim.ID fit the encoding.TextMarshaler interface.
func (receiver ID) MarshalText() (text []byte, err error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	var serialized string = serialize(receiver.value)

	return []byte(serialized), nil
}

// UnixTime returns the unix‐time that is embeddd in the xim-id.
//
// Example usage:
//
//	var id xim.ID = xim.Generate()
//	
//	unixtime, successful := id.UnixTime()
//	if !successful {
//		return errors.New("xim-id was not initialized")
//	}
//	
//	var t time.Time = time.Unix(unixtime, 0)
//	
//	fmt.Println("xim-id was created on:", t)
func (receiver ID) UnixTime() (int64, bool) {
	if Nothing() == receiver {
		return 0, false
	}

	value, _ := decompile(receiver.value)

	return int64(value), true
}

// UnmarshalText makes xim.ID fit the encoding.TextUnmarshaler interface.
func (receiver *ID) UnmarshalText(p []byte) error {
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
