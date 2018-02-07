// Code generated by protoc-gen-gogo.
// source: schema.proto
// DO NOT EDIT!

package tipb

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
)
import math "math"

// discarding unused import gogoproto "gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TableInfo struct {
	TableId          int64         `protobuf:"varint,1,opt,name=table_id" json:"table_id"`
	Columns          []*ColumnInfo `protobuf:"bytes,2,rep,name=columns" json:"columns,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *TableInfo) Reset()         { *m = TableInfo{} }
func (m *TableInfo) String() string { return proto.CompactTextString(m) }
func (*TableInfo) ProtoMessage()    {}

func (m *TableInfo) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *TableInfo) GetColumns() []*ColumnInfo {
	if m != nil {
		return m.Columns
	}
	return nil
}

type ColumnInfo struct {
	ColumnId         int64    `protobuf:"varint,1,opt,name=column_id" json:"column_id"`
	Tp               int32    `protobuf:"varint,2,opt,name=tp" json:"tp"`
	Collation        int32    `protobuf:"varint,3,opt,name=collation" json:"collation"`
	ColumnLen        int32    `protobuf:"varint,4,opt,name=columnLen" json:"columnLen"`
	Decimal          int32    `protobuf:"varint,5,opt,name=decimal" json:"decimal"`
	Flag             int32    `protobuf:"varint,6,opt,name=flag" json:"flag"`
	Elems            []string `protobuf:"bytes,7,rep,name=elems" json:"elems,omitempty"`
	DefaultVal       []byte   `protobuf:"bytes,8,opt,name=default_val" json:"default_val,omitempty"`
	PkHandle         bool     `protobuf:"varint,21,opt,name=pk_handle" json:"pk_handle"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ColumnInfo) Reset()         { *m = ColumnInfo{} }
func (m *ColumnInfo) String() string { return proto.CompactTextString(m) }
func (*ColumnInfo) ProtoMessage()    {}

func (m *ColumnInfo) GetColumnId() int64 {
	if m != nil {
		return m.ColumnId
	}
	return 0
}

func (m *ColumnInfo) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

func (m *ColumnInfo) GetCollation() int32 {
	if m != nil {
		return m.Collation
	}
	return 0
}

func (m *ColumnInfo) GetColumnLen() int32 {
	if m != nil {
		return m.ColumnLen
	}
	return 0
}

func (m *ColumnInfo) GetDecimal() int32 {
	if m != nil {
		return m.Decimal
	}
	return 0
}

func (m *ColumnInfo) GetFlag() int32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

func (m *ColumnInfo) GetElems() []string {
	if m != nil {
		return m.Elems
	}
	return nil
}

func (m *ColumnInfo) GetDefaultVal() []byte {
	if m != nil {
		return m.DefaultVal
	}
	return nil
}

func (m *ColumnInfo) GetPkHandle() bool {
	if m != nil {
		return m.PkHandle
	}
	return false
}

type IndexInfo struct {
	TableId          int64         `protobuf:"varint,1,opt,name=table_id" json:"table_id"`
	IndexId          int64         `protobuf:"varint,2,opt,name=index_id" json:"index_id"`
	Columns          []*ColumnInfo `protobuf:"bytes,3,rep,name=columns" json:"columns,omitempty"`
	Unique           bool          `protobuf:"varint,4,opt,name=unique" json:"unique"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *IndexInfo) Reset()         { *m = IndexInfo{} }
func (m *IndexInfo) String() string { return proto.CompactTextString(m) }
func (*IndexInfo) ProtoMessage()    {}

func (m *IndexInfo) GetTableId() int64 {
	if m != nil {
		return m.TableId
	}
	return 0
}

func (m *IndexInfo) GetIndexId() int64 {
	if m != nil {
		return m.IndexId
	}
	return 0
}

func (m *IndexInfo) GetColumns() []*ColumnInfo {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *IndexInfo) GetUnique() bool {
	if m != nil {
		return m.Unique
	}
	return false
}

// KeyRange is the encoded index key range, low is closed, high is open. (low <= x < high)
type KeyRange struct {
	Low              []byte `protobuf:"bytes,1,opt,name=low" json:"low,omitempty"`
	High             []byte `protobuf:"bytes,2,opt,name=high" json:"high,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyRange) Reset()         { *m = KeyRange{} }
func (m *KeyRange) String() string { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()    {}

func (m *KeyRange) GetLow() []byte {
	if m != nil {
		return m.Low
	}
	return nil
}

func (m *KeyRange) GetHigh() []byte {
	if m != nil {
		return m.High
	}
	return nil
}

func (m *TableInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *TableInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintSchema(data, i, uint64(m.TableId))
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			data[i] = 0x12
			i++
			i = encodeVarintSchema(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ColumnInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ColumnInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintSchema(data, i, uint64(m.ColumnId))
	data[i] = 0x10
	i++
	i = encodeVarintSchema(data, i, uint64(m.Tp))
	data[i] = 0x18
	i++
	i = encodeVarintSchema(data, i, uint64(m.Collation))
	data[i] = 0x20
	i++
	i = encodeVarintSchema(data, i, uint64(m.ColumnLen))
	data[i] = 0x28
	i++
	i = encodeVarintSchema(data, i, uint64(m.Decimal))
	data[i] = 0x30
	i++
	i = encodeVarintSchema(data, i, uint64(m.Flag))
	if len(m.Elems) > 0 {
		for _, s := range m.Elems {
			data[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	if m.DefaultVal != nil {
		data[i] = 0x42
		i++
		i = encodeVarintSchema(data, i, uint64(len(m.DefaultVal)))
		i += copy(data[i:], m.DefaultVal)
	}
	data[i] = 0xa8
	i++
	data[i] = 0x1
	i++
	if m.PkHandle {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *IndexInfo) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *IndexInfo) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintSchema(data, i, uint64(m.TableId))
	data[i] = 0x10
	i++
	i = encodeVarintSchema(data, i, uint64(m.IndexId))
	if len(m.Columns) > 0 {
		for _, msg := range m.Columns {
			data[i] = 0x1a
			i++
			i = encodeVarintSchema(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x20
	i++
	if m.Unique {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *KeyRange) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *KeyRange) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Low != nil {
		data[i] = 0xa
		i++
		i = encodeVarintSchema(data, i, uint64(len(m.Low)))
		i += copy(data[i:], m.Low)
	}
	if m.High != nil {
		data[i] = 0x12
		i++
		i = encodeVarintSchema(data, i, uint64(len(m.High)))
		i += copy(data[i:], m.High)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Schema(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Schema(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSchema(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *TableInfo) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.TableId))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ColumnInfo) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.ColumnId))
	n += 1 + sovSchema(uint64(m.Tp))
	n += 1 + sovSchema(uint64(m.Collation))
	n += 1 + sovSchema(uint64(m.ColumnLen))
	n += 1 + sovSchema(uint64(m.Decimal))
	n += 1 + sovSchema(uint64(m.Flag))
	if len(m.Elems) > 0 {
		for _, s := range m.Elems {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	if m.DefaultVal != nil {
		l = len(m.DefaultVal)
		n += 1 + l + sovSchema(uint64(l))
	}
	n += 3
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IndexInfo) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSchema(uint64(m.TableId))
	n += 1 + sovSchema(uint64(m.IndexId))
	if len(m.Columns) > 0 {
		for _, e := range m.Columns {
			l = e.Size()
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	n += 2
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *KeyRange) Size() (n int) {
	var l int
	_ = l
	if m.Low != nil {
		l = len(m.Low)
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.High != nil {
		l = len(m.High)
		n += 1 + l + sovSchema(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSchema(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TableInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TableId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnInfo{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSchema(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ColumnInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnId", wireType)
			}
			m.ColumnId = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ColumnId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tp", wireType)
			}
			m.Tp = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Tp |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collation", wireType)
			}
			m.Collation = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Collation |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnLen", wireType)
			}
			m.ColumnLen = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ColumnLen |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimal", wireType)
			}
			m.Decimal = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Decimal |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag", wireType)
			}
			m.Flag = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Flag |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elems", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elems = append(m.Elems, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultVal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultVal = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PkHandle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.PkHandle = bool(v != 0)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSchema(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *IndexInfo) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableId", wireType)
			}
			m.TableId = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TableId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexId", wireType)
			}
			m.IndexId = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.IndexId |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Columns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Columns = append(m.Columns, &ColumnInfo{})
			if err := m.Columns[len(m.Columns)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unique", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Unique = bool(v != 0)
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSchema(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func (m *KeyRange) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Low", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Low = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field High", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.High = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipSchema(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSchema
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	return nil
}
func skipSchema(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSchema(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSchema = fmt.Errorf("proto: negative length found during unmarshaling")
)