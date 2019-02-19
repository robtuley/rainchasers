// Code generated by gopkg.in/actgardner/gogen-avro.v5. DO NOT EDIT.
/*
 * SOURCE:
 *     gauge.avsc
 */

package avro

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type ByteWriter interface {
	Grow(int)
	WriteByte(byte) error
}

type StringWriter interface {
	WriteString(string) (int, error)
}

func encodeFloat(w io.Writer, byteCount int, bits uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}
	for i := 0; i < byteCount; i++ {
		if bw != nil {
			err = bw.WriteByte(byte(bits & 255))
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(bits&255))
		}
		bits = bits >> 8
	}
	if bw == nil {
		_, err = w.Write(bb)
		return err
	}
	return nil
}

func encodeInt(w io.Writer, byteCount int, encoded uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	// To avoid reallocations, grow capacity to the largest possible size
	// for this integer
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}

	if encoded == 0 {
		if bw != nil {
			err = bw.WriteByte(0)
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(0))
		}
	} else {
		for encoded > 0 {
			b := byte(encoded & 127)
			encoded = encoded >> 7
			if !(encoded == 0) {
				b |= 128
			}
			if bw != nil {
				err = bw.WriteByte(b)
				if err != nil {
					return err
				}
			} else {
				bb = append(bb, b)
			}
		}
	}
	if bw == nil {
		_, err := w.Write(bb)
		return err
	}
	return nil

}

func readArrayMeasure(r io.Reader) ([]*Measure, error) {
	var err error
	var blkSize int64
	var arr = make([]*Measure, 0)
	for {
		blkSize, err = readLong(r)
		if err != nil {
			return nil, err
		}
		if blkSize == 0 {
			break
		}
		if blkSize < 0 {
			blkSize = -blkSize
			_, err = readLong(r)
			if err != nil {
				return nil, err
			}
		}
		for i := int64(0); i < blkSize; i++ {
			elem, err := readMeasure(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, elem)
		}
	}
	return arr, nil
}

func readFloat(r io.Reader) (float32, error) {
	buf := make([]byte, 4)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint32(buf)
	val := math.Float32frombits(bits)
	return val, nil

}

func readInt(r io.Reader) (int32, error) {
	var v int
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		v |= int(b&127) << shift
		if b&128 == 0 {
			break
		}
	}
	datum := (int32(v>>1) ^ -int32(v&1))
	return datum, nil
}

func readLong(r io.Reader) (int64, error) {
	var v uint64
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		v |= uint64(b&127) << shift
		if b&128 == 0 {
			break
		}
	}
	datum := (int64(v>>1) ^ -int64(v&1))
	return datum, nil
}

func readMeasure(r io.Reader) (*Measure, error) {
	var str = &Measure{}
	var err error
	str.Event_time, err = readLong(r)
	if err != nil {
		return nil, err
	}
	str.Value, err = readFloat(r)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func readSnapshot(r io.Reader) (*Snapshot, error) {
	var str = &Snapshot{}
	var err error
	str.Data_url, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Human_url, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Name, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.River_name, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Lat, err = readFloat(r)
	if err != nil {
		return nil, err
	}
	str.Lg, err = readFloat(r)
	if err != nil {
		return nil, err
	}
	str.Unit, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Type, err = readTypeValues(r)
	if err != nil {
		return nil, err
	}
	str.Readings, err = readArrayMeasure(r)
	if err != nil {
		return nil, err
	}
	str.Processing_time, err = readLong(r)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func readString(r io.Reader) (string, error) {
	len, err := readLong(r)
	if err != nil {
		return "", err
	}

	// makeslice can fail depending on available memory.
	// We arbitrarily limit string size to sane default (~2.2GB).
	if len < 0 || len > math.MaxInt32 {
		return "", fmt.Errorf("string length out of range: %d", len)
	}

	bb := make([]byte, len)
	_, err = io.ReadFull(r, bb)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}

func readTypeValues(r io.Reader) (TypeValues, error) {
	val, err := readInt(r)
	return TypeValues(val), err
}

func readUnionMeasureSnapshot(r io.Reader) (UnionMeasureSnapshot, error) {
	field, err := readLong(r)
	var unionStr UnionMeasureSnapshot
	if err != nil {
		return unionStr, err
	}
	unionStr.UnionType = UnionMeasureSnapshotTypeEnum(field)
	switch unionStr.UnionType {
	case UnionMeasureSnapshotTypeEnumMeasure:
		val, err := readMeasure(r)
		if err != nil {
			return unionStr, err
		}
		unionStr.Measure = val
	case UnionMeasureSnapshotTypeEnumSnapshot:
		val, err := readSnapshot(r)
		if err != nil {
			return unionStr, err
		}
		unionStr.Snapshot = val

	default:
		return unionStr, fmt.Errorf("Invalid value for UnionMeasureSnapshot")
	}
	return unionStr, nil
}

func writeArrayMeasure(r []*Measure, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeMeasure(e, w)
		if err != nil {
			return err
		}
	}
	return writeLong(0, w)
}

func writeFloat(r float32, w io.Writer) error {
	bits := uint64(math.Float32bits(r))
	const byteCount = 4
	return encodeFloat(w, byteCount, bits)
}

func writeInt(r int32, w io.Writer) error {
	downShift := uint32(31)
	encoded := uint64((uint32(r) << 1) ^ uint32(r>>downShift))
	const maxByteSize = 5
	return encodeInt(w, maxByteSize, encoded)
}

func writeLong(r int64, w io.Writer) error {
	downShift := uint64(63)
	encoded := uint64((r << 1) ^ (r >> downShift))
	const maxByteSize = 10
	return encodeInt(w, maxByteSize, encoded)
}

func writeMeasure(r *Measure, w io.Writer) error {
	var err error
	err = writeLong(r.Event_time, w)
	if err != nil {
		return err
	}
	err = writeFloat(r.Value, w)
	if err != nil {
		return err
	}

	return nil
}
func writeSnapshot(r *Snapshot, w io.Writer) error {
	var err error
	err = writeString(r.Data_url, w)
	if err != nil {
		return err
	}
	err = writeString(r.Human_url, w)
	if err != nil {
		return err
	}
	err = writeString(r.Name, w)
	if err != nil {
		return err
	}
	err = writeString(r.River_name, w)
	if err != nil {
		return err
	}
	err = writeFloat(r.Lat, w)
	if err != nil {
		return err
	}
	err = writeFloat(r.Lg, w)
	if err != nil {
		return err
	}
	err = writeString(r.Unit, w)
	if err != nil {
		return err
	}
	err = writeTypeValues(r.Type, w)
	if err != nil {
		return err
	}
	err = writeArrayMeasure(r.Readings, w)
	if err != nil {
		return err
	}
	err = writeLong(r.Processing_time, w)
	if err != nil {
		return err
	}

	return nil
}

func writeString(r string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil {
		return err
	}
	if sw, ok := w.(StringWriter); ok {
		_, err = sw.WriteString(r)
	} else {
		_, err = w.Write([]byte(r))
	}
	return err
}

func writeTypeValues(r TypeValues, w io.Writer) error {
	return writeInt(int32(r), w)
}

func writeUnionMeasureSnapshot(r UnionMeasureSnapshot, w io.Writer) error {
	err := writeLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionMeasureSnapshotTypeEnumMeasure:
		return writeMeasure(r.Measure, w)
	case UnionMeasureSnapshotTypeEnumSnapshot:
		return writeSnapshot(r.Snapshot, w)

	}
	return fmt.Errorf("Invalid value for UnionMeasureSnapshot")
}