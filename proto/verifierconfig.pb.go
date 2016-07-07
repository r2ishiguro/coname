// Code generated by protoc-gen-gogo.
// source: verifierconfig.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/maditya/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/maditya/protobuf/gogoproto"

import bytes "bytes"

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

type VerifierConfig struct {
	ID                   uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SigningKeyID         string              `protobuf:"bytes,2,opt,name=signing_key_id,json=signingKeyId,proto3" json:"signing_key_id,omitempty"`
	Realm                string              `protobuf:"bytes,3,opt,name=realm,proto3" json:"realm,omitempty"`
	TLS                  *TLSConfig          `protobuf:"bytes,4,opt,name=tls" json:"tls,omitempty"`
	KeyserverAddr        string              `protobuf:"bytes,5,opt,name=keyserver_addr,json=keyserverAddr,proto3" json:"keyserver_addr,omitempty"`
	InitialKeyserverAuth AuthorizationPolicy `protobuf:"bytes,6,opt,name=initial_keyserver_auth,json=initialKeyserverAuth" json:"initial_keyserver_auth"`
	TreeNonce            []byte              `protobuf:"bytes,7,opt,name=tree_nonce,json=treeNonce,proto3" json:"tree_nonce,omitempty"`
}

func (m *VerifierConfig) Reset()                    { *m = VerifierConfig{} }
func (*VerifierConfig) ProtoMessage()               {}
func (*VerifierConfig) Descriptor() ([]byte, []int) { return fileDescriptorVerifierconfig, []int{0} }

func (m *VerifierConfig) GetTLS() *TLSConfig {
	if m != nil {
		return m.TLS
	}
	return nil
}

func (m *VerifierConfig) GetInitialKeyserverAuth() AuthorizationPolicy {
	if m != nil {
		return m.InitialKeyserverAuth
	}
	return AuthorizationPolicy{}
}

func init() {
	proto1.RegisterType((*VerifierConfig)(nil), "proto.VerifierConfig")
}
func (this *VerifierConfig) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*VerifierConfig)
	if !ok {
		that2, ok := that.(VerifierConfig)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *VerifierConfig")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *VerifierConfig but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *VerifierConfig but is not nil && this == nil")
	}
	if this.ID != that1.ID {
		return fmt.Errorf("ID this(%v) Not Equal that(%v)", this.ID, that1.ID)
	}
	if this.SigningKeyID != that1.SigningKeyID {
		return fmt.Errorf("SigningKeyID this(%v) Not Equal that(%v)", this.SigningKeyID, that1.SigningKeyID)
	}
	if this.Realm != that1.Realm {
		return fmt.Errorf("Realm this(%v) Not Equal that(%v)", this.Realm, that1.Realm)
	}
	if !this.TLS.Equal(that1.TLS) {
		return fmt.Errorf("TLS this(%v) Not Equal that(%v)", this.TLS, that1.TLS)
	}
	if this.KeyserverAddr != that1.KeyserverAddr {
		return fmt.Errorf("KeyserverAddr this(%v) Not Equal that(%v)", this.KeyserverAddr, that1.KeyserverAddr)
	}
	if !this.InitialKeyserverAuth.Equal(&that1.InitialKeyserverAuth) {
		return fmt.Errorf("InitialKeyserverAuth this(%v) Not Equal that(%v)", this.InitialKeyserverAuth, that1.InitialKeyserverAuth)
	}
	if !bytes.Equal(this.TreeNonce, that1.TreeNonce) {
		return fmt.Errorf("TreeNonce this(%v) Not Equal that(%v)", this.TreeNonce, that1.TreeNonce)
	}
	return nil
}
func (this *VerifierConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*VerifierConfig)
	if !ok {
		that2, ok := that.(VerifierConfig)
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
	if this.ID != that1.ID {
		return false
	}
	if this.SigningKeyID != that1.SigningKeyID {
		return false
	}
	if this.Realm != that1.Realm {
		return false
	}
	if !this.TLS.Equal(that1.TLS) {
		return false
	}
	if this.KeyserverAddr != that1.KeyserverAddr {
		return false
	}
	if !this.InitialKeyserverAuth.Equal(&that1.InitialKeyserverAuth) {
		return false
	}
	if !bytes.Equal(this.TreeNonce, that1.TreeNonce) {
		return false
	}
	return true
}
func (this *VerifierConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&proto.VerifierConfig{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "SigningKeyID: "+fmt.Sprintf("%#v", this.SigningKeyID)+",\n")
	s = append(s, "Realm: "+fmt.Sprintf("%#v", this.Realm)+",\n")
	if this.TLS != nil {
		s = append(s, "TLS: "+fmt.Sprintf("%#v", this.TLS)+",\n")
	}
	s = append(s, "KeyserverAddr: "+fmt.Sprintf("%#v", this.KeyserverAddr)+",\n")
	s = append(s, "InitialKeyserverAuth: "+strings.Replace(this.InitialKeyserverAuth.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "TreeNonce: "+fmt.Sprintf("%#v", this.TreeNonce)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringVerifierconfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringVerifierconfig(e map[int32]github_com_maditya_protobuf_proto.Extension) string {
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
func (m *VerifierConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *VerifierConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(m.ID))
	}
	if len(m.SigningKeyID) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(len(m.SigningKeyID)))
		i += copy(data[i:], m.SigningKeyID)
	}
	if len(m.Realm) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(len(m.Realm)))
		i += copy(data[i:], m.Realm)
	}
	if m.TLS != nil {
		data[i] = 0x22
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(m.TLS.Size()))
		n1, err := m.TLS.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.KeyserverAddr) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(len(m.KeyserverAddr)))
		i += copy(data[i:], m.KeyserverAddr)
	}
	data[i] = 0x32
	i++
	i = encodeVarintVerifierconfig(data, i, uint64(m.InitialKeyserverAuth.Size()))
	n2, err := m.InitialKeyserverAuth.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.TreeNonce) > 0 {
		data[i] = 0x3a
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(len(m.TreeNonce)))
		i += copy(data[i:], m.TreeNonce)
	}
	return i, nil
}

func encodeFixed64Verifierconfig(data []byte, offset int, v uint64) int {
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
func encodeFixed32Verifierconfig(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintVerifierconfig(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedVerifierConfig(r randyVerifierconfig, easy bool) *VerifierConfig {
	this := &VerifierConfig{}
	this.ID = uint64(uint64(r.Uint32()))
	this.SigningKeyID = randStringVerifierconfig(r)
	this.Realm = randStringVerifierconfig(r)
	if r.Intn(10) != 0 {
		this.TLS = NewPopulatedTLSConfig(r, easy)
	}
	this.KeyserverAddr = randStringVerifierconfig(r)
	v1 := NewPopulatedAuthorizationPolicy(r, easy)
	this.InitialKeyserverAuth = *v1
	v2 := r.Intn(100)
	this.TreeNonce = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.TreeNonce[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyVerifierconfig interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneVerifierconfig(r randyVerifierconfig) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringVerifierconfig(r randyVerifierconfig) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneVerifierconfig(r)
	}
	return string(tmps)
}
func randUnrecognizedVerifierconfig(r randyVerifierconfig, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldVerifierconfig(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldVerifierconfig(data []byte, r randyVerifierconfig, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateVerifierconfig(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateVerifierconfig(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateVerifierconfig(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *VerifierConfig) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovVerifierconfig(uint64(m.ID))
	}
	l = len(m.SigningKeyID)
	if l > 0 {
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	l = len(m.Realm)
	if l > 0 {
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	if m.TLS != nil {
		l = m.TLS.Size()
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	l = len(m.KeyserverAddr)
	if l > 0 {
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	l = m.InitialKeyserverAuth.Size()
	n += 1 + l + sovVerifierconfig(uint64(l))
	l = len(m.TreeNonce)
	if l > 0 {
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	return n
}

func sovVerifierconfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVerifierconfig(x uint64) (n int) {
	return sovVerifierconfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *VerifierConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VerifierConfig{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`SigningKeyID:` + fmt.Sprintf("%v", this.SigningKeyID) + `,`,
		`Realm:` + fmt.Sprintf("%v", this.Realm) + `,`,
		`TLS:` + strings.Replace(fmt.Sprintf("%v", this.TLS), "TLSConfig", "TLSConfig", 1) + `,`,
		`KeyserverAddr:` + fmt.Sprintf("%v", this.KeyserverAddr) + `,`,
		`InitialKeyserverAuth:` + strings.Replace(strings.Replace(this.InitialKeyserverAuth.String(), "AuthorizationPolicy", "AuthorizationPolicy", 1), `&`, ``, 1) + `,`,
		`TreeNonce:` + fmt.Sprintf("%v", this.TreeNonce) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringVerifierconfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *VerifierConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifierconfig
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
			return fmt.Errorf("proto: VerifierConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifierConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
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
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigningKeyID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Realm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
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
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Realm = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TLS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
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
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TLS == nil {
				m.TLS = &TLSConfig{}
			}
			if err := m.TLS.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyserverAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
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
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyserverAddr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialKeyserverAuth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
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
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialKeyserverAuth.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreeNonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
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
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreeNonce = append(m.TreeNonce[:0], data[iNdEx:postIndex]...)
			if m.TreeNonce == nil {
				m.TreeNonce = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerifierconfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifierconfig
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
func skipVerifierconfig(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerifierconfig
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
					return 0, ErrIntOverflowVerifierconfig
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
					return 0, ErrIntOverflowVerifierconfig
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
				return 0, ErrInvalidLengthVerifierconfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVerifierconfig
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
				next, err := skipVerifierconfig(data[start:])
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
	ErrInvalidLengthVerifierconfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerifierconfig   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorVerifierconfig = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x6b, 0xe2, 0x40,
	0x18, 0xc6, 0x4d, 0xfc, 0x87, 0xb3, 0x59, 0x57, 0x06, 0x91, 0x10, 0xd8, 0xb8, 0x2c, 0x2c, 0x2c,
	0x2c, 0xe8, 0xd2, 0x42, 0xe9, 0xd5, 0xb4, 0x97, 0x62, 0x29, 0x25, 0x16, 0xaf, 0x21, 0x26, 0x63,
	0x1c, 0x1a, 0x67, 0xca, 0x64, 0x2c, 0xa4, 0xa7, 0x7e, 0x9c, 0x7e, 0x84, 0x1e, 0xdb, 0x9b, 0x47,
	0x8f, 0x3d, 0x89, 0x7a, 0xea, 0xd1, 0x63, 0x8f, 0x7d, 0x33, 0x11, 0xeb, 0xe1, 0x61, 0xf2, 0x3c,
	0xf3, 0x9b, 0xe7, 0x7d, 0x09, 0x6a, 0xde, 0x13, 0x41, 0xc7, 0x94, 0x88, 0x80, 0xb3, 0x31, 0x8d,
	0x3a, 0x77, 0x82, 0x4b, 0x8e, 0xcb, 0xea, 0xb0, 0xfe, 0x47, 0x54, 0x4e, 0x66, 0xa3, 0x4e, 0xc0,
	0xa7, 0xdd, 0xa9, 0x1f, 0x52, 0x99, 0xfa, 0x5d, 0x75, 0x33, 0x9a, 0x8d, 0xbb, 0x11, 0x8f, 0xb8,
	0x32, 0xea, 0x2b, 0x7f, 0x68, 0xfd, 0x90, 0x71, 0x72, 0xd8, 0x64, 0x19, 0x41, 0x4c, 0x09, 0x93,
	0xb9, 0xfb, 0xfd, 0xaa, 0xa3, 0xfa, 0x70, 0x37, 0xf0, 0x4c, 0x61, 0xb8, 0x85, 0x74, 0x1a, 0x9a,
	0xda, 0x2f, 0xed, 0x6f, 0xc9, 0xa9, 0x6c, 0x96, 0x6d, 0xfd, 0xe2, 0xdc, 0x85, 0x04, 0x9f, 0xa0,
	0x7a, 0x42, 0x23, 0x46, 0x59, 0xe4, 0xdd, 0x92, 0xd4, 0x03, 0x46, 0x07, 0xa6, 0xe6, 0x34, 0x80,
	0x31, 0x06, 0xf9, 0x4d, 0x9f, 0xa4, 0x40, 0x1b, 0xc9, 0x97, 0x0b, 0x71, 0x13, 0x95, 0x05, 0xf1,
	0xe3, 0xa9, 0x59, 0xcc, 0x70, 0x37, 0x37, 0xf8, 0x1f, 0x2a, 0xc2, 0x66, 0x66, 0x09, 0xb2, 0x6f,
	0x47, 0x8d, 0x7c, 0x9b, 0xce, 0xcd, 0xe5, 0x20, 0x5f, 0xc2, 0xa9, 0x42, 0x69, 0x11, 0xac, 0x9b,
	0x51, 0xf8, 0x0f, 0xaa, 0xc3, 0xc8, 0x84, 0x08, 0xf8, 0x37, 0x9e, 0x1f, 0x86, 0xc2, 0x2c, 0xab,
	0xae, 0xef, 0xfb, 0xb4, 0x07, 0x21, 0x1e, 0xa2, 0x16, 0x65, 0x54, 0x52, 0x3f, 0xf6, 0x0e, 0xf0,
	0x99, 0x9c, 0x98, 0x15, 0x35, 0xc6, 0xda, 0x8d, 0xe9, 0x41, 0xc4, 0x05, 0x7d, 0xf0, 0x25, 0xe5,
	0xec, 0x9a, 0xc7, 0x34, 0x48, 0x9d, 0xd2, 0x7c, 0xd9, 0x2e, 0xb8, 0xcd, 0xdd, 0xfb, 0xfe, 0xbe,
	0x17, 0x50, 0xfc, 0x13, 0x21, 0x29, 0x08, 0xf1, 0x18, 0x67, 0x01, 0x31, 0xab, 0xd0, 0x65, 0xb8,
	0xb5, 0x2c, 0xb9, 0xca, 0x02, 0xe7, 0x74, 0xb1, 0xb6, 0x0b, 0x6f, 0xa0, 0xd5, 0xda, 0xd6, 0xb6,
	0xa0, 0x0f, 0xd0, 0xe3, 0xc6, 0xd6, 0x9e, 0x40, 0xcf, 0xa0, 0x17, 0xd0, 0x1c, 0xb4, 0x00, 0xad,
	0x40, 0xef, 0x1b, 0xbb, 0xb0, 0x85, 0x73, 0x54, 0x51, 0xfb, 0x1c, 0x7f, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x11, 0x30, 0x86, 0xb2, 0xf4, 0x01, 0x00, 0x00,
}
