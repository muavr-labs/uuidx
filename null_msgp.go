package uuidx

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *NullUUID) DecodeMsg(dc *msgp.Reader) (err error) {
	if dc.IsNil() {
		z.Valid = false
		z.UUID = Nil
		return dc.Skip()
	}

	z.Valid = true
	return z.UUID.DecodeMsg(dc)
}

// EncodeMsg implements msgp.Encodable
func (z *NullUUID) EncodeMsg(en *msgp.Writer) (err error) {
	if !z.Valid {
		err = en.WriteNil()
	} else {
		err = z.UUID.EncodeMsg(en)
	}

	return
}

// MarshalMsg implements msgp.Marshaler
func (z *NullUUID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	if !z.Valid {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.UUID.MarshalMsg(o)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *NullUUID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	switch msgp.NextType(bts) {
	case msgp.NilType:
		bts, err = msgp.ReadNilBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		z.Valid = false
		z.UUID = Nil
	case msgp.ArrayType:
		bts, err = z.UUID.UnmarshalMsg(bts)
		if err != nil {
			return
		}
		z.Valid = true
	default:
		bts, err = msgp.Skip(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
	}

	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *NullUUID) Msgsize() (s int) {
	if z.Valid {
		s += msgp.NilSize
	} else {
		s += z.UUID.Msgsize()
	}

	return
}
