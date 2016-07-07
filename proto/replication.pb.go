// Code generated by protoc-gen-gogo.
// source: replication.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"

import strings "strings"
import github_com_maditya_protobuf_proto "github.com/maditya/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// KeyserverStep denotes the input to a single step of the keyserver state
// machine. Serializable high-availability replication is achieved by
// replicating an in-order log of all steps and having each replica reproduce
// the state from them.
type KeyserverStep struct {
	UID uint64 `protobuf:"fixed64,1,opt,name=UID,json=uID,proto3" json:"UID,omitempty"`
	// TODO: should all fields in a oneof have their own types for extensibility?
	//
	// Types that are valid to be assigned to Type:
	//	*KeyserverStep_Update
	//	*KeyserverStep_EpochDelimiter
	//	*KeyserverStep_ReplicaSigned
	//	*KeyserverStep_VerifierSigned
	Type isKeyserverStep_Type `protobuf_oneof:"type"`
}

func (m *KeyserverStep) Reset()                    { *m = KeyserverStep{} }
func (*KeyserverStep) ProtoMessage()               {}
func (*KeyserverStep) Descriptor() ([]byte, []int) { return fileDescriptorReplication, []int{0} }

type isKeyserverStep_Type interface {
	isKeyserverStep_Type()
	Equal(interface{}) bool
	VerboseEqual(interface{}) error
	MarshalTo([]byte) (int, error)
	Size() int
}

type KeyserverStep_Update struct {
	Update *UpdateRequest `protobuf:"bytes,2,opt,name=update,oneof"`
}
type KeyserverStep_EpochDelimiter struct {
	EpochDelimiter *EpochDelimiter `protobuf:"bytes,3,opt,name=epoch_delimiter,json=epochDelimiter,oneof"`
}
type KeyserverStep_ReplicaSigned struct {
	ReplicaSigned *SignedEpochHead `protobuf:"bytes,4,opt,name=replica_signed,json=replicaSigned,oneof"`
}
type KeyserverStep_VerifierSigned struct {
	VerifierSigned *SignedEpochHead `protobuf:"bytes,5,opt,name=verifier_signed,json=verifierSigned,oneof"`
}

func (*KeyserverStep_Update) isKeyserverStep_Type()         {}
func (*KeyserverStep_EpochDelimiter) isKeyserverStep_Type() {}
func (*KeyserverStep_ReplicaSigned) isKeyserverStep_Type()  {}
func (*KeyserverStep_VerifierSigned) isKeyserverStep_Type() {}

func (m *KeyserverStep) GetType() isKeyserverStep_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *KeyserverStep) GetUpdate() *UpdateRequest {
	if x, ok := m.GetType().(*KeyserverStep_Update); ok {
		return x.Update
	}
	return nil
}

func (m *KeyserverStep) GetEpochDelimiter() *EpochDelimiter {
	if x, ok := m.GetType().(*KeyserverStep_EpochDelimiter); ok {
		return x.EpochDelimiter
	}
	return nil
}

func (m *KeyserverStep) GetReplicaSigned() *SignedEpochHead {
	if x, ok := m.GetType().(*KeyserverStep_ReplicaSigned); ok {
		return x.ReplicaSigned
	}
	return nil
}

func (m *KeyserverStep) GetVerifierSigned() *SignedEpochHead {
	if x, ok := m.GetType().(*KeyserverStep_VerifierSigned); ok {
		return x.VerifierSigned
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*KeyserverStep) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _KeyserverStep_OneofMarshaler, _KeyserverStep_OneofUnmarshaler, _KeyserverStep_OneofSizer, []interface{}{
		(*KeyserverStep_Update)(nil),
		(*KeyserverStep_EpochDelimiter)(nil),
		(*KeyserverStep_ReplicaSigned)(nil),
		(*KeyserverStep_VerifierSigned)(nil),
	}
}

func _KeyserverStep_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*KeyserverStep)
	// type
	switch x := m.Type.(type) {
	case *KeyserverStep_Update:
		_ = b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *KeyserverStep_EpochDelimiter:
		_ = b.EncodeVarint(3<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.EpochDelimiter); err != nil {
			return err
		}
	case *KeyserverStep_ReplicaSigned:
		_ = b.EncodeVarint(4<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.ReplicaSigned); err != nil {
			return err
		}
	case *KeyserverStep_VerifierSigned:
		_ = b.EncodeVarint(5<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.VerifierSigned); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("KeyserverStep.Type has unexpected type %T", x)
	}
	return nil
}

func _KeyserverStep_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*KeyserverStep)
	switch tag {
	case 2: // type.update
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(UpdateRequest)
		err := b.DecodeMessage(msg)
		m.Type = &KeyserverStep_Update{msg}
		return true, err
	case 3: // type.epoch_delimiter
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(EpochDelimiter)
		err := b.DecodeMessage(msg)
		m.Type = &KeyserverStep_EpochDelimiter{msg}
		return true, err
	case 4: // type.replica_signed
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(SignedEpochHead)
		err := b.DecodeMessage(msg)
		m.Type = &KeyserverStep_ReplicaSigned{msg}
		return true, err
	case 5: // type.verifier_signed
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(SignedEpochHead)
		err := b.DecodeMessage(msg)
		m.Type = &KeyserverStep_VerifierSigned{msg}
		return true, err
	default:
		return false, nil
	}
}

func _KeyserverStep_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*KeyserverStep)
	// type
	switch x := m.Type.(type) {
	case *KeyserverStep_Update:
		s := proto1.Size(x.Update)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *KeyserverStep_EpochDelimiter:
		s := proto1.Size(x.EpochDelimiter)
		n += proto1.SizeVarint(3<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *KeyserverStep_ReplicaSigned:
		s := proto1.Size(x.ReplicaSigned)
		n += proto1.SizeVarint(4<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *KeyserverStep_VerifierSigned:
		s := proto1.Size(x.VerifierSigned)
		n += proto1.SizeVarint(5<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type EpochDelimiter struct {
	EpochNumber uint64    `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	Timestamp   Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp"`
}

func (m *EpochDelimiter) Reset()                    { *m = EpochDelimiter{} }
func (*EpochDelimiter) ProtoMessage()               {}
func (*EpochDelimiter) Descriptor() ([]byte, []int) { return fileDescriptorReplication, []int{1} }

func (m *EpochDelimiter) GetTimestamp() Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return Timestamp{}
}

func init() {
	proto1.RegisterType((*KeyserverStep)(nil), "proto.KeyserverStep")
	proto1.RegisterType((*EpochDelimiter)(nil), "proto.EpochDelimiter")
}
func (this *KeyserverStep) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*KeyserverStep)
	if !ok {
		that2, ok := that.(KeyserverStep)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *KeyserverStep")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *KeyserverStep but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *KeyserverStep but is not nil && this == nil")
	}
	if this.UID != that1.UID {
		return fmt.Errorf("UID this(%v) Not Equal that(%v)", this.UID, that1.UID)
	}
	if that1.Type == nil {
		if this.Type != nil {
			return fmt.Errorf("this.Type != nil && that1.Type == nil")
		}
	} else if this.Type == nil {
		return fmt.Errorf("this.Type == nil && that1.Type != nil")
	} else if err := this.Type.VerboseEqual(that1.Type); err != nil {
		return err
	}
	return nil
}
func (this *KeyserverStep_Update) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*KeyserverStep_Update)
	if !ok {
		that2, ok := that.(KeyserverStep_Update)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *KeyserverStep_Update")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *KeyserverStep_Update but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *KeyserverStep_Update but is not nil && this == nil")
	}
	if !this.Update.Equal(that1.Update) {
		return fmt.Errorf("Update this(%v) Not Equal that(%v)", this.Update, that1.Update)
	}
	return nil
}
func (this *KeyserverStep_EpochDelimiter) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*KeyserverStep_EpochDelimiter)
	if !ok {
		that2, ok := that.(KeyserverStep_EpochDelimiter)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *KeyserverStep_EpochDelimiter")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *KeyserverStep_EpochDelimiter but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *KeyserverStep_EpochDelimiter but is not nil && this == nil")
	}
	if !this.EpochDelimiter.Equal(that1.EpochDelimiter) {
		return fmt.Errorf("EpochDelimiter this(%v) Not Equal that(%v)", this.EpochDelimiter, that1.EpochDelimiter)
	}
	return nil
}
func (this *KeyserverStep_ReplicaSigned) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*KeyserverStep_ReplicaSigned)
	if !ok {
		that2, ok := that.(KeyserverStep_ReplicaSigned)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *KeyserverStep_ReplicaSigned")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *KeyserverStep_ReplicaSigned but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *KeyserverStep_ReplicaSigned but is not nil && this == nil")
	}
	if !this.ReplicaSigned.Equal(that1.ReplicaSigned) {
		return fmt.Errorf("ReplicaSigned this(%v) Not Equal that(%v)", this.ReplicaSigned, that1.ReplicaSigned)
	}
	return nil
}
func (this *KeyserverStep_VerifierSigned) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*KeyserverStep_VerifierSigned)
	if !ok {
		that2, ok := that.(KeyserverStep_VerifierSigned)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *KeyserverStep_VerifierSigned")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *KeyserverStep_VerifierSigned but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *KeyserverStep_VerifierSigned but is not nil && this == nil")
	}
	if !this.VerifierSigned.Equal(that1.VerifierSigned) {
		return fmt.Errorf("VerifierSigned this(%v) Not Equal that(%v)", this.VerifierSigned, that1.VerifierSigned)
	}
	return nil
}
func (this *KeyserverStep) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KeyserverStep)
	if !ok {
		that2, ok := that.(KeyserverStep)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.UID != that1.UID {
		return false
	}
	if that1.Type == nil {
		if this.Type != nil {
			return false
		}
	} else if this.Type == nil {
		return false
	} else if !this.Type.Equal(that1.Type) {
		return false
	}
	return true
}
func (this *KeyserverStep_Update) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KeyserverStep_Update)
	if !ok {
		that2, ok := that.(KeyserverStep_Update)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Update.Equal(that1.Update) {
		return false
	}
	return true
}
func (this *KeyserverStep_EpochDelimiter) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KeyserverStep_EpochDelimiter)
	if !ok {
		that2, ok := that.(KeyserverStep_EpochDelimiter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.EpochDelimiter.Equal(that1.EpochDelimiter) {
		return false
	}
	return true
}
func (this *KeyserverStep_ReplicaSigned) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KeyserverStep_ReplicaSigned)
	if !ok {
		that2, ok := that.(KeyserverStep_ReplicaSigned)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.ReplicaSigned.Equal(that1.ReplicaSigned) {
		return false
	}
	return true
}
func (this *KeyserverStep_VerifierSigned) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KeyserverStep_VerifierSigned)
	if !ok {
		that2, ok := that.(KeyserverStep_VerifierSigned)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.VerifierSigned.Equal(that1.VerifierSigned) {
		return false
	}
	return true
}
func (this *EpochDelimiter) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*EpochDelimiter)
	if !ok {
		that2, ok := that.(EpochDelimiter)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *EpochDelimiter")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *EpochDelimiter but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *EpochDelimiter but is not nil && this == nil")
	}
	if this.EpochNumber != that1.EpochNumber {
		return fmt.Errorf("EpochNumber this(%v) Not Equal that(%v)", this.EpochNumber, that1.EpochNumber)
	}
	if !this.Timestamp.Equal(&that1.Timestamp) {
		return fmt.Errorf("Timestamp this(%v) Not Equal that(%v)", this.Timestamp, that1.Timestamp)
	}
	return nil
}
func (this *EpochDelimiter) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*EpochDelimiter)
	if !ok {
		that2, ok := that.(EpochDelimiter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.EpochNumber != that1.EpochNumber {
		return false
	}
	if !this.Timestamp.Equal(&that1.Timestamp) {
		return false
	}
	return true
}
func (this *KeyserverStep) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&proto.KeyserverStep{")
	s = append(s, "UID: "+fmt.Sprintf("%#v", this.UID)+",\n")
	if this.Type != nil {
		s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *KeyserverStep_Update) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&proto.KeyserverStep_Update{` +
		`Update:` + fmt.Sprintf("%#v", this.Update) + `}`}, ", ")
	return s
}
func (this *KeyserverStep_EpochDelimiter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&proto.KeyserverStep_EpochDelimiter{` +
		`EpochDelimiter:` + fmt.Sprintf("%#v", this.EpochDelimiter) + `}`}, ", ")
	return s
}
func (this *KeyserverStep_ReplicaSigned) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&proto.KeyserverStep_ReplicaSigned{` +
		`ReplicaSigned:` + fmt.Sprintf("%#v", this.ReplicaSigned) + `}`}, ", ")
	return s
}
func (this *KeyserverStep_VerifierSigned) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&proto.KeyserverStep_VerifierSigned{` +
		`VerifierSigned:` + fmt.Sprintf("%#v", this.VerifierSigned) + `}`}, ", ")
	return s
}
func (this *EpochDelimiter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&proto.EpochDelimiter{")
	s = append(s, "EpochNumber: "+fmt.Sprintf("%#v", this.EpochNumber)+",\n")
	s = append(s, "Timestamp: "+strings.Replace(this.Timestamp.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringReplication(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringReplication(e map[int32]github_com_maditya_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *KeyserverStep) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *KeyserverStep) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UID != 0 {
		data[i] = 0x9
		i++
		i = encodeFixed64Replication(data, i, uint64(m.UID))
	}
	if m.Type != nil {
		nn1, err := m.Type.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *KeyserverStep_Update) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.Update != nil {
		data[i] = 0x12
		i++
		i = encodeVarintReplication(data, i, uint64(m.Update.Size()))
		n2, err := m.Update.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *KeyserverStep_EpochDelimiter) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.EpochDelimiter != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintReplication(data, i, uint64(m.EpochDelimiter.Size()))
		n3, err := m.EpochDelimiter.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *KeyserverStep_ReplicaSigned) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.ReplicaSigned != nil {
		data[i] = 0x22
		i++
		i = encodeVarintReplication(data, i, uint64(m.ReplicaSigned.Size()))
		n4, err := m.ReplicaSigned.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *KeyserverStep_VerifierSigned) MarshalTo(data []byte) (int, error) {
	i := 0
	if m.VerifierSigned != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintReplication(data, i, uint64(m.VerifierSigned.Size()))
		n5, err := m.VerifierSigned.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *EpochDelimiter) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *EpochDelimiter) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EpochNumber != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintReplication(data, i, uint64(m.EpochNumber))
	}
	data[i] = 0x12
	i++
	i = encodeVarintReplication(data, i, uint64(m.Timestamp.Size()))
	n6, err := m.Timestamp.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func encodeFixed64Replication(data []byte, offset int, v uint64) int {
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
func encodeFixed32Replication(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintReplication(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedKeyserverStep(r randyReplication, easy bool) *KeyserverStep {
	this := &KeyserverStep{}
	this.UID = uint64(uint64(r.Uint32()))
	oneofNumber_Type := []int32{2, 3, 4, 5}[r.Intn(4)]
	switch oneofNumber_Type {
	case 2:
		this.Type = NewPopulatedKeyserverStep_Update(r, easy)
	case 3:
		this.Type = NewPopulatedKeyserverStep_EpochDelimiter(r, easy)
	case 4:
		this.Type = NewPopulatedKeyserverStep_ReplicaSigned(r, easy)
	case 5:
		this.Type = NewPopulatedKeyserverStep_VerifierSigned(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedKeyserverStep_Update(r randyReplication, easy bool) *KeyserverStep_Update {
	this := &KeyserverStep_Update{}
	this.Update = NewPopulatedUpdateRequest(r, easy)
	return this
}
func NewPopulatedKeyserverStep_EpochDelimiter(r randyReplication, easy bool) *KeyserverStep_EpochDelimiter {
	this := &KeyserverStep_EpochDelimiter{}
	this.EpochDelimiter = NewPopulatedEpochDelimiter(r, easy)
	return this
}
func NewPopulatedKeyserverStep_ReplicaSigned(r randyReplication, easy bool) *KeyserverStep_ReplicaSigned {
	this := &KeyserverStep_ReplicaSigned{}
	this.ReplicaSigned = NewPopulatedSignedEpochHead(r, easy)
	return this
}
func NewPopulatedKeyserverStep_VerifierSigned(r randyReplication, easy bool) *KeyserverStep_VerifierSigned {
	this := &KeyserverStep_VerifierSigned{}
	this.VerifierSigned = NewPopulatedSignedEpochHead(r, easy)
	return this
}
func NewPopulatedEpochDelimiter(r randyReplication, easy bool) *EpochDelimiter {
	this := &EpochDelimiter{}
	this.EpochNumber = uint64(uint64(r.Uint32()))
	v1 := NewPopulatedTimestamp(r, easy)
	this.Timestamp = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyReplication interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneReplication(r randyReplication) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringReplication(r randyReplication) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneReplication(r)
	}
	return string(tmps)
}
func randUnrecognizedReplication(r randyReplication, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldReplication(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldReplication(data []byte, r randyReplication, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateReplication(data, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		data = encodeVarintPopulateReplication(data, uint64(v3))
	case 1:
		data = encodeVarintPopulateReplication(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateReplication(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateReplication(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateReplication(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateReplication(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *KeyserverStep) Size() (n int) {
	var l int
	_ = l
	if m.UID != 0 {
		n += 9
	}
	if m.Type != nil {
		n += m.Type.Size()
	}
	return n
}

func (m *KeyserverStep_Update) Size() (n int) {
	var l int
	_ = l
	if m.Update != nil {
		l = m.Update.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	return n
}
func (m *KeyserverStep_EpochDelimiter) Size() (n int) {
	var l int
	_ = l
	if m.EpochDelimiter != nil {
		l = m.EpochDelimiter.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	return n
}
func (m *KeyserverStep_ReplicaSigned) Size() (n int) {
	var l int
	_ = l
	if m.ReplicaSigned != nil {
		l = m.ReplicaSigned.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	return n
}
func (m *KeyserverStep_VerifierSigned) Size() (n int) {
	var l int
	_ = l
	if m.VerifierSigned != nil {
		l = m.VerifierSigned.Size()
		n += 1 + l + sovReplication(uint64(l))
	}
	return n
}
func (m *EpochDelimiter) Size() (n int) {
	var l int
	_ = l
	if m.EpochNumber != 0 {
		n += 1 + sovReplication(uint64(m.EpochNumber))
	}
	l = m.Timestamp.Size()
	n += 1 + l + sovReplication(uint64(l))
	return n
}

func sovReplication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozReplication(x uint64) (n int) {
	return sovReplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *KeyserverStep) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyserverStep{`,
		`UID:` + fmt.Sprintf("%v", this.UID) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyserverStep_Update) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyserverStep_Update{`,
		`Update:` + strings.Replace(fmt.Sprintf("%v", this.Update), "UpdateRequest", "UpdateRequest", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyserverStep_EpochDelimiter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyserverStep_EpochDelimiter{`,
		`EpochDelimiter:` + strings.Replace(fmt.Sprintf("%v", this.EpochDelimiter), "EpochDelimiter", "EpochDelimiter", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyserverStep_ReplicaSigned) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyserverStep_ReplicaSigned{`,
		`ReplicaSigned:` + strings.Replace(fmt.Sprintf("%v", this.ReplicaSigned), "SignedEpochHead", "SignedEpochHead", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *KeyserverStep_VerifierSigned) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&KeyserverStep_VerifierSigned{`,
		`VerifierSigned:` + strings.Replace(fmt.Sprintf("%v", this.VerifierSigned), "SignedEpochHead", "SignedEpochHead", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *EpochDelimiter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&EpochDelimiter{`,
		`EpochNumber:` + fmt.Sprintf("%v", this.EpochNumber) + `,`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Timestamp", "Timestamp", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringReplication(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *KeyserverStep) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReplication
			}
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
		if wireType == 4 {
			return fmt.Errorf("proto: KeyserverStep: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyserverStep: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			m.UID = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			m.UID = uint64(data[iNdEx-8])
			m.UID |= uint64(data[iNdEx-7]) << 8
			m.UID |= uint64(data[iNdEx-6]) << 16
			m.UID |= uint64(data[iNdEx-5]) << 24
			m.UID |= uint64(data[iNdEx-4]) << 32
			m.UID |= uint64(data[iNdEx-3]) << 40
			m.UID |= uint64(data[iNdEx-2]) << 48
			m.UID |= uint64(data[iNdEx-1]) << 56
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Update", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReplication
				}
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
			if msglen < 0 {
				return ErrInvalidLengthReplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &UpdateRequest{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Type = &KeyserverStep_Update{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDelimiter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReplication
				}
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
			if msglen < 0 {
				return ErrInvalidLengthReplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EpochDelimiter{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Type = &KeyserverStep_EpochDelimiter{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaSigned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReplication
				}
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
			if msglen < 0 {
				return ErrInvalidLengthReplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignedEpochHead{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Type = &KeyserverStep_ReplicaSigned{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifierSigned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReplication
				}
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
			if msglen < 0 {
				return ErrInvalidLengthReplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SignedEpochHead{}
			if err := v.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Type = &KeyserverStep_VerifierSigned{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReplication(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EpochDelimiter) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReplication
			}
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
		if wireType == 4 {
			return fmt.Errorf("proto: EpochDelimiter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochDelimiter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
			}
			m.EpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReplication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.EpochNumber |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReplication
				}
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
			if msglen < 0 {
				return ErrInvalidLengthReplication
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timestamp.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReplication(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReplication(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReplication
			}
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
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReplication
				}
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
				if shift >= 64 {
					return 0, ErrIntOverflowReplication
				}
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
				return 0, ErrInvalidLengthReplication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowReplication
					}
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
				next, err := skipReplication(data[start:])
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
	ErrInvalidLengthReplication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReplication   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorReplication = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xbf, 0x6f, 0xda, 0x40,
	0x14, 0xc7, 0x6d, 0x30, 0x96, 0x7a, 0x80, 0xa1, 0xa7, 0xb6, 0xb2, 0x18, 0xdc, 0x96, 0xa9, 0x93,
	0xa9, 0xda, 0x0e, 0xdd, 0xda, 0x22, 0x2a, 0x81, 0x2a, 0x75, 0x30, 0x61, 0x46, 0xfe, 0xf1, 0x30,
	0x27, 0xe1, 0x1f, 0x31, 0xe7, 0x48, 0x6c, 0xf9, 0x4f, 0xb2, 0xe6, 0x4f, 0xc8, 0x98, 0x91, 0x91,
	0x31, 0x53, 0x04, 0x4c, 0x19, 0x19, 0x33, 0xe6, 0x71, 0x3e, 0x27, 0x62, 0xca, 0xf0, 0xd1, 0xf9,
	0xbe, 0xef, 0x7d, 0xbf, 0xbe, 0xf7, 0xc8, 0xdb, 0x0c, 0xd2, 0x05, 0xf3, 0x5d, 0xce, 0x92, 0xd8,
	0x4e, 0xb3, 0x84, 0x27, 0xb4, 0x26, 0x8e, 0xce, 0xd7, 0x90, 0xf1, 0x79, 0xee, 0xd9, 0x7e, 0x12,
	0xf5, 0x22, 0x37, 0x60, 0x7c, 0xe5, 0xf6, 0x44, 0xc5, 0xcb, 0x67, 0xbd, 0x30, 0x09, 0x13, 0x71,
	0x11, 0x5f, 0x85, 0xb1, 0xd3, 0xf0, 0x17, 0x0c, 0x62, 0x2e, 0x6f, 0x2d, 0xce, 0x22, 0x58, 0x72,
	0x37, 0x4a, 0x0b, 0xa1, 0x7b, 0x55, 0x21, 0xcd, 0x7f, 0xb0, 0x5a, 0x42, 0x76, 0x01, 0xd9, 0x98,
	0x43, 0x4a, 0xdb, 0xa4, 0x3a, 0x19, 0x0d, 0x4c, 0xf5, 0x93, 0xfa, 0x45, 0x77, 0xaa, 0xf9, 0x68,
	0x40, 0x6d, 0xa2, 0xe7, 0x69, 0xe0, 0x72, 0x30, 0x2b, 0x28, 0xd6, 0xbf, 0xbd, 0x2b, 0xbc, 0xf6,
	0x44, 0x88, 0x0e, 0x9c, 0xe7, 0x18, 0x39, 0x54, 0x1c, 0xd9, 0x45, 0x7f, 0x93, 0x16, 0xa4, 0x89,
	0x3f, 0x9f, 0x06, 0xb0, 0x60, 0x11, 0xe3, 0x90, 0x99, 0x55, 0x61, 0x7c, 0x2f, 0x8d, 0x7f, 0x8f,
	0xd5, 0x41, 0x59, 0x44, 0xa7, 0x01, 0x27, 0x0a, 0xfd, 0x45, 0x0c, 0xb9, 0x82, 0xe9, 0x92, 0x85,
	0x31, 0x04, 0xa6, 0x26, 0x02, 0x3e, 0xc8, 0x80, 0xb1, 0x10, 0x45, 0xcc, 0x10, 0xdc, 0x00, 0x13,
	0x9a, 0xb2, 0xbf, 0xa8, 0xd0, 0x3f, 0xa4, 0x85, 0xf3, 0xb0, 0x19, 0x83, 0xac, 0x4c, 0xa8, 0xbd,
	0x92, 0x60, 0x94, 0x86, 0xa2, 0xd4, 0xd7, 0x89, 0xc6, 0x57, 0x29, 0x74, 0x19, 0x31, 0x4e, 0xdf,
	0x4b, 0x3f, 0x93, 0x46, 0x31, 0x5f, 0x9c, 0x47, 0x1e, 0x0e, 0x77, 0x5c, 0x95, 0xe6, 0xd4, 0x85,
	0xf6, 0x5f, 0x48, 0xf4, 0x07, 0x79, 0xf3, 0xbc, 0x69, 0xb9, 0xb5, 0xb6, 0xfc, 0xf3, 0x59, 0xa9,
	0xf7, 0xb5, 0xf5, 0xfd, 0x47, 0xc5, 0x79, 0x69, 0xec, 0xff, 0xdc, 0xec, 0x2c, 0xe5, 0x0e, 0xd9,
	0xee, 0x2c, 0xf5, 0x80, 0x3c, 0x22, 0x97, 0x7b, 0x4b, 0xbd, 0x46, 0x6e, 0x90, 0x5b, 0x64, 0x8d,
	0x6c, 0x90, 0x2d, 0xf2, 0xb0, 0xb7, 0x94, 0x03, 0x9e, 0x9e, 0x2e, 0xb2, 0xbf, 0x3f, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x86, 0x63, 0x1f, 0x2e, 0x3a, 0x02, 0x00, 0x00,
}
