// Code generated by capnpc-go. DO NOT EDIT.

package chf

import (
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Custom struct{ capnp.Struct }
type Custom_value Custom
type Custom_value_Which uint16

const (
	Custom_value_Which_uint32Val  Custom_value_Which = 0
	Custom_value_Which_float32Val Custom_value_Which = 1
	Custom_value_Which_strVal     Custom_value_Which = 2
	Custom_value_Which_uint64Val  Custom_value_Which = 3
	Custom_value_Which_addrVal    Custom_value_Which = 4
	Custom_value_Which_uint16Val  Custom_value_Which = 5
	Custom_value_Which_uint8Val   Custom_value_Which = 6
)

func (w Custom_value_Which) String() string {
	const s = "uint32Valfloat32ValstrValuint64ValaddrValuint16Valuint8Val"
	switch w {
	case Custom_value_Which_uint32Val:
		return s[0:9]
	case Custom_value_Which_float32Val:
		return s[9:19]
	case Custom_value_Which_strVal:
		return s[19:25]
	case Custom_value_Which_uint64Val:
		return s[25:34]
	case Custom_value_Which_addrVal:
		return s[34:41]
	case Custom_value_Which_uint16Val:
		return s[41:50]
	case Custom_value_Which_uint8Val:
		return s[50:58]

	}
	return "Custom_value_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Custom_TypeID is the unique identifier for the type Custom.
const Custom_TypeID = 0xed5d37861203d027

func NewCustom(s *capnp.Segment) (Custom, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Custom{st}, err
}

func NewRootCustom(s *capnp.Segment) (Custom, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Custom{st}, err
}

func ReadRootCustom(msg *capnp.Message) (Custom, error) {
	root, err := msg.RootPtr()
	return Custom{root.Struct()}, err
}

func (s Custom) String() string {
	str, _ := text.Marshal(0xed5d37861203d027, s.Struct)
	return str
}

func (s Custom) Id() uint32 {
	return s.Struct.Uint32(0)
}

func (s Custom) SetId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Custom) Value() Custom_value { return Custom_value(s) }

func (s Custom_value) Which() Custom_value_Which {
	return Custom_value_Which(s.Struct.Uint16(8))
}
func (s Custom_value) Uint32Val() uint32 {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != uint32Val")
	}
	return s.Struct.Uint32(4)
}

func (s Custom_value) SetUint32Val(v uint32) {
	s.Struct.SetUint16(8, 0)
	s.Struct.SetUint32(4, v)
}

func (s Custom_value) Float32Val() float32 {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != float32Val")
	}
	return math.Float32frombits(s.Struct.Uint32(4))
}

func (s Custom_value) SetFloat32Val(v float32) {
	s.Struct.SetUint16(8, 1)
	s.Struct.SetUint32(4, math.Float32bits(v))
}

func (s Custom_value) StrVal() (string, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != strVal")
	}
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Custom_value) HasStrVal() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Custom_value) StrValBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Custom_value) SetStrVal(v string) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetText(0, v)
}

func (s Custom_value) Uint64Val() uint64 {
	if s.Struct.Uint16(8) != 3 {
		panic("Which() != uint64Val")
	}
	return s.Struct.Uint64(16)
}

func (s Custom_value) SetUint64Val(v uint64) {
	s.Struct.SetUint16(8, 3)
	s.Struct.SetUint64(16, v)
}

func (s Custom_value) AddrVal() ([]byte, error) {
	if s.Struct.Uint16(8) != 4 {
		panic("Which() != addrVal")
	}
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Custom_value) HasAddrVal() bool {
	if s.Struct.Uint16(8) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Custom_value) SetAddrVal(v []byte) error {
	s.Struct.SetUint16(8, 4)
	return s.Struct.SetData(0, v)
}

func (s Custom_value) Uint16Val() uint16 {
	if s.Struct.Uint16(8) != 5 {
		panic("Which() != uint16Val")
	}
	return s.Struct.Uint16(4)
}

func (s Custom_value) SetUint16Val(v uint16) {
	s.Struct.SetUint16(8, 5)
	s.Struct.SetUint16(4, v)
}

func (s Custom_value) Uint8Val() uint8 {
	if s.Struct.Uint16(8) != 6 {
		panic("Which() != uint8Val")
	}
	return s.Struct.Uint8(4)
}

func (s Custom_value) SetUint8Val(v uint8) {
	s.Struct.SetUint16(8, 6)
	s.Struct.SetUint8(4, v)
}

func (s Custom) IsDimension() bool {
	return s.Struct.Bit(80)
}

func (s Custom) SetIsDimension(v bool) {
	s.Struct.SetBit(80, v)
}

// Custom_List is a list of Custom.
type Custom_List struct{ capnp.List }

// NewCustom creates a new list of Custom.
func NewCustom_List(s *capnp.Segment, sz int32) (Custom_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1}, sz)
	return Custom_List{l}, err
}

func (s Custom_List) At(i int) Custom { return Custom{s.List.Struct(i)} }

func (s Custom_List) Set(i int, v Custom) error { return s.List.SetStruct(i, v.Struct) }

func (s Custom_List) String() string {
	str, _ := text.MarshalList(0xed5d37861203d027, s.List)
	return str
}

// Custom_Promise is a wrapper for a Custom promised by a client call.
type Custom_Promise struct{ *capnp.Pipeline }

func (p Custom_Promise) Struct() (Custom, error) {
	s, err := p.Pipeline.Struct()
	return Custom{s}, err
}

func (p Custom_Promise) Value() Custom_value_Promise { return Custom_value_Promise{p.Pipeline} }

// Custom_value_Promise is a wrapper for a Custom_value promised by a client call.
type Custom_value_Promise struct{ *capnp.Pipeline }

func (p Custom_value_Promise) Struct() (Custom_value, error) {
	s, err := p.Pipeline.Struct()
	return Custom_value{s}, err
}

type CHF struct{ capnp.Struct }

// CHF_TypeID is the unique identifier for the type CHF.
const CHF_TypeID = 0xa7ab5c68e4bc7b62

func NewCHF(s *capnp.Segment) (CHF, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 232, PointerCount: 14})
	return CHF{st}, err
}

func NewRootCHF(s *capnp.Segment) (CHF, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 232, PointerCount: 14})
	return CHF{st}, err
}

func ReadRootCHF(msg *capnp.Message) (CHF, error) {
	root, err := msg.RootPtr()
	return CHF{root.Struct()}, err
}

func (s CHF) String() string {
	str, _ := text.Marshal(0xa7ab5c68e4bc7b62, s.Struct)
	return str
}

func (s CHF) TimestampNano() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s CHF) SetTimestampNano(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s CHF) DstAs() uint32 {
	return s.Struct.Uint32(8)
}

func (s CHF) SetDstAs(v uint32) {
	s.Struct.SetUint32(8, v)
}

func (s CHF) DstGeo() uint32 {
	return s.Struct.Uint32(12)
}

func (s CHF) SetDstGeo(v uint32) {
	s.Struct.SetUint32(12, v)
}

func (s CHF) DstMac() uint32 {
	return s.Struct.Uint32(16)
}

func (s CHF) SetDstMac(v uint32) {
	s.Struct.SetUint32(16, v)
}

func (s CHF) HeaderLen() uint32 {
	return s.Struct.Uint32(20)
}

func (s CHF) SetHeaderLen(v uint32) {
	s.Struct.SetUint32(20, v)
}

func (s CHF) InBytes() uint64 {
	return s.Struct.Uint64(24)
}

func (s CHF) SetInBytes(v uint64) {
	s.Struct.SetUint64(24, v)
}

func (s CHF) InPkts() uint64 {
	return s.Struct.Uint64(32)
}

func (s CHF) SetInPkts(v uint64) {
	s.Struct.SetUint64(32, v)
}

func (s CHF) InputPort() uint32 {
	return s.Struct.Uint32(40)
}

func (s CHF) SetInputPort(v uint32) {
	s.Struct.SetUint32(40, v)
}

func (s CHF) IpSize() uint32 {
	return s.Struct.Uint32(44)
}

func (s CHF) SetIpSize(v uint32) {
	s.Struct.SetUint32(44, v)
}

func (s CHF) Ipv4DstAddr() uint32 {
	return s.Struct.Uint32(48)
}

func (s CHF) SetIpv4DstAddr(v uint32) {
	s.Struct.SetUint32(48, v)
}

func (s CHF) Ipv4SrcAddr() uint32 {
	return s.Struct.Uint32(52)
}

func (s CHF) SetIpv4SrcAddr(v uint32) {
	s.Struct.SetUint32(52, v)
}

func (s CHF) L4DstPort() uint32 {
	return s.Struct.Uint32(56)
}

func (s CHF) SetL4DstPort(v uint32) {
	s.Struct.SetUint32(56, v)
}

func (s CHF) L4SrcPort() uint32 {
	return s.Struct.Uint32(60)
}

func (s CHF) SetL4SrcPort(v uint32) {
	s.Struct.SetUint32(60, v)
}

func (s CHF) OutputPort() uint32 {
	return s.Struct.Uint32(64)
}

func (s CHF) SetOutputPort(v uint32) {
	s.Struct.SetUint32(64, v)
}

func (s CHF) Protocol() uint32 {
	return s.Struct.Uint32(68)
}

func (s CHF) SetProtocol(v uint32) {
	s.Struct.SetUint32(68, v)
}

func (s CHF) SampledPacketSize() uint32 {
	return s.Struct.Uint32(72)
}

func (s CHF) SetSampledPacketSize(v uint32) {
	s.Struct.SetUint32(72, v)
}

func (s CHF) SrcAs() uint32 {
	return s.Struct.Uint32(76)
}

func (s CHF) SetSrcAs(v uint32) {
	s.Struct.SetUint32(76, v)
}

func (s CHF) SrcGeo() uint32 {
	return s.Struct.Uint32(80)
}

func (s CHF) SetSrcGeo(v uint32) {
	s.Struct.SetUint32(80, v)
}

func (s CHF) SrcMac() uint32 {
	return s.Struct.Uint32(84)
}

func (s CHF) SetSrcMac(v uint32) {
	s.Struct.SetUint32(84, v)
}

func (s CHF) TcpFlags() uint32 {
	return s.Struct.Uint32(88)
}

func (s CHF) SetTcpFlags(v uint32) {
	s.Struct.SetUint32(88, v)
}

func (s CHF) Tos() uint32 {
	return s.Struct.Uint32(92)
}

func (s CHF) SetTos(v uint32) {
	s.Struct.SetUint32(92, v)
}

func (s CHF) VlanIn() uint32 {
	return s.Struct.Uint32(96)
}

func (s CHF) SetVlanIn(v uint32) {
	s.Struct.SetUint32(96, v)
}

func (s CHF) VlanOut() uint32 {
	return s.Struct.Uint32(100)
}

func (s CHF) SetVlanOut(v uint32) {
	s.Struct.SetUint32(100, v)
}

func (s CHF) Ipv4NextHop() uint32 {
	return s.Struct.Uint32(104)
}

func (s CHF) SetIpv4NextHop(v uint32) {
	s.Struct.SetUint32(104, v)
}

func (s CHF) MplsType() uint32 {
	return s.Struct.Uint32(108)
}

func (s CHF) SetMplsType(v uint32) {
	s.Struct.SetUint32(108, v)
}

func (s CHF) OutBytes() uint64 {
	return s.Struct.Uint64(112)
}

func (s CHF) SetOutBytes(v uint64) {
	s.Struct.SetUint64(112, v)
}

func (s CHF) OutPkts() uint64 {
	return s.Struct.Uint64(120)
}

func (s CHF) SetOutPkts(v uint64) {
	s.Struct.SetUint64(120, v)
}

func (s CHF) TcpRetransmit() uint32 {
	return s.Struct.Uint32(128)
}

func (s CHF) SetTcpRetransmit(v uint32) {
	s.Struct.SetUint32(128, v)
}

func (s CHF) SrcFlowTags() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s CHF) HasSrcFlowTags() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CHF) SrcFlowTagsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s CHF) SetSrcFlowTags(v string) error {
	return s.Struct.SetText(0, v)
}

func (s CHF) DstFlowTags() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s CHF) HasDstFlowTags() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CHF) DstFlowTagsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s CHF) SetDstFlowTags(v string) error {
	return s.Struct.SetText(1, v)
}

func (s CHF) SampleRate() uint32 {
	return s.Struct.Uint32(132)
}

func (s CHF) SetSampleRate(v uint32) {
	s.Struct.SetUint32(132, v)
}

func (s CHF) DeviceId() uint32 {
	return s.Struct.Uint32(136)
}

func (s CHF) SetDeviceId(v uint32) {
	s.Struct.SetUint32(136, v)
}

func (s CHF) FlowTags() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s CHF) HasFlowTags() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s CHF) FlowTagsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s CHF) SetFlowTags(v string) error {
	return s.Struct.SetText(2, v)
}

func (s CHF) Timestamp() int64 {
	return int64(s.Struct.Uint64(144))
}

func (s CHF) SetTimestamp(v int64) {
	s.Struct.SetUint64(144, uint64(v))
}

func (s CHF) DstBgpAsPath() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s CHF) HasDstBgpAsPath() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s CHF) DstBgpAsPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s CHF) SetDstBgpAsPath(v string) error {
	return s.Struct.SetText(3, v)
}

func (s CHF) DstBgpCommunity() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s CHF) HasDstBgpCommunity() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s CHF) DstBgpCommunityBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s CHF) SetDstBgpCommunity(v string) error {
	return s.Struct.SetText(4, v)
}

func (s CHF) SrcBgpAsPath() (string, error) {
	p, err := s.Struct.Ptr(5)
	return p.Text(), err
}

func (s CHF) HasSrcBgpAsPath() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s CHF) SrcBgpAsPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return p.TextBytes(), err
}

func (s CHF) SetSrcBgpAsPath(v string) error {
	return s.Struct.SetText(5, v)
}

func (s CHF) SrcBgpCommunity() (string, error) {
	p, err := s.Struct.Ptr(6)
	return p.Text(), err
}

func (s CHF) HasSrcBgpCommunity() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s CHF) SrcBgpCommunityBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(6)
	return p.TextBytes(), err
}

func (s CHF) SetSrcBgpCommunity(v string) error {
	return s.Struct.SetText(6, v)
}

func (s CHF) SrcNextHopAs() uint32 {
	return s.Struct.Uint32(140)
}

func (s CHF) SetSrcNextHopAs(v uint32) {
	s.Struct.SetUint32(140, v)
}

func (s CHF) DstNextHopAs() uint32 {
	return s.Struct.Uint32(152)
}

func (s CHF) SetDstNextHopAs(v uint32) {
	s.Struct.SetUint32(152, v)
}

func (s CHF) SrcGeoRegion() uint32 {
	return s.Struct.Uint32(156)
}

func (s CHF) SetSrcGeoRegion(v uint32) {
	s.Struct.SetUint32(156, v)
}

func (s CHF) DstGeoRegion() uint32 {
	return s.Struct.Uint32(160)
}

func (s CHF) SetDstGeoRegion(v uint32) {
	s.Struct.SetUint32(160, v)
}

func (s CHF) SrcGeoCity() uint32 {
	return s.Struct.Uint32(164)
}

func (s CHF) SetSrcGeoCity(v uint32) {
	s.Struct.SetUint32(164, v)
}

func (s CHF) DstGeoCity() uint32 {
	return s.Struct.Uint32(168)
}

func (s CHF) SetDstGeoCity(v uint32) {
	s.Struct.SetUint32(168, v)
}

func (s CHF) Big() bool {
	return s.Struct.Bit(1376)
}

func (s CHF) SetBig(v bool) {
	s.Struct.SetBit(1376, v)
}

func (s CHF) SampleAdj() bool {
	return s.Struct.Bit(1377)
}

func (s CHF) SetSampleAdj(v bool) {
	s.Struct.SetBit(1377, v)
}

func (s CHF) Ipv4DstNextHop() uint32 {
	return s.Struct.Uint32(176)
}

func (s CHF) SetIpv4DstNextHop(v uint32) {
	s.Struct.SetUint32(176, v)
}

func (s CHF) Ipv4SrcNextHop() uint32 {
	return s.Struct.Uint32(180)
}

func (s CHF) SetIpv4SrcNextHop(v uint32) {
	s.Struct.SetUint32(180, v)
}

func (s CHF) SrcRoutePrefix() uint32 {
	return s.Struct.Uint32(184)
}

func (s CHF) SetSrcRoutePrefix(v uint32) {
	s.Struct.SetUint32(184, v)
}

func (s CHF) DstRoutePrefix() uint32 {
	return s.Struct.Uint32(188)
}

func (s CHF) SetDstRoutePrefix(v uint32) {
	s.Struct.SetUint32(188, v)
}

func (s CHF) SrcRouteLength() uint8 {
	return s.Struct.Uint8(173)
}

func (s CHF) SetSrcRouteLength(v uint8) {
	s.Struct.SetUint8(173, v)
}

func (s CHF) DstRouteLength() uint8 {
	return s.Struct.Uint8(174)
}

func (s CHF) SetDstRouteLength(v uint8) {
	s.Struct.SetUint8(174, v)
}

func (s CHF) SrcSecondAsn() uint32 {
	return s.Struct.Uint32(192)
}

func (s CHF) SetSrcSecondAsn(v uint32) {
	s.Struct.SetUint32(192, v)
}

func (s CHF) DstSecondAsn() uint32 {
	return s.Struct.Uint32(196)
}

func (s CHF) SetDstSecondAsn(v uint32) {
	s.Struct.SetUint32(196, v)
}

func (s CHF) SrcThirdAsn() uint32 {
	return s.Struct.Uint32(200)
}

func (s CHF) SetSrcThirdAsn(v uint32) {
	s.Struct.SetUint32(200, v)
}

func (s CHF) DstThirdAsn() uint32 {
	return s.Struct.Uint32(204)
}

func (s CHF) SetDstThirdAsn(v uint32) {
	s.Struct.SetUint32(204, v)
}

func (s CHF) Ipv6DstAddr() ([]byte, error) {
	p, err := s.Struct.Ptr(7)
	return []byte(p.Data()), err
}

func (s CHF) HasIpv6DstAddr() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s CHF) SetIpv6DstAddr(v []byte) error {
	return s.Struct.SetData(7, v)
}

func (s CHF) Ipv6SrcAddr() ([]byte, error) {
	p, err := s.Struct.Ptr(8)
	return []byte(p.Data()), err
}

func (s CHF) HasIpv6SrcAddr() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s CHF) SetIpv6SrcAddr(v []byte) error {
	return s.Struct.SetData(8, v)
}

func (s CHF) SrcEthMac() uint64 {
	return s.Struct.Uint64(208)
}

func (s CHF) SetSrcEthMac(v uint64) {
	s.Struct.SetUint64(208, v)
}

func (s CHF) DstEthMac() uint64 {
	return s.Struct.Uint64(216)
}

func (s CHF) SetDstEthMac(v uint64) {
	s.Struct.SetUint64(216, v)
}

func (s CHF) Custom() (Custom_List, error) {
	p, err := s.Struct.Ptr(9)
	return Custom_List{List: p.List()}, err
}

func (s CHF) HasCustom() bool {
	p, err := s.Struct.Ptr(9)
	return p.IsValid() || err != nil
}

func (s CHF) SetCustom(v Custom_List) error {
	return s.Struct.SetPtr(9, v.List.ToPtr())
}

// NewCustom sets the custom field to a newly
// allocated Custom_List, preferring placement in s's segment.
func (s CHF) NewCustom(n int32) (Custom_List, error) {
	l, err := NewCustom_List(s.Struct.Segment(), n)
	if err != nil {
		return Custom_List{}, err
	}
	err = s.Struct.SetPtr(9, l.List.ToPtr())
	return l, err
}

func (s CHF) Ipv6SrcNextHop() ([]byte, error) {
	p, err := s.Struct.Ptr(10)
	return []byte(p.Data()), err
}

func (s CHF) HasIpv6SrcNextHop() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s CHF) SetIpv6SrcNextHop(v []byte) error {
	return s.Struct.SetData(10, v)
}

func (s CHF) Ipv6DstNextHop() ([]byte, error) {
	p, err := s.Struct.Ptr(11)
	return []byte(p.Data()), err
}

func (s CHF) HasIpv6DstNextHop() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s CHF) SetIpv6DstNextHop(v []byte) error {
	return s.Struct.SetData(11, v)
}

func (s CHF) Ipv6SrcRoutePrefix() ([]byte, error) {
	p, err := s.Struct.Ptr(12)
	return []byte(p.Data()), err
}

func (s CHF) HasIpv6SrcRoutePrefix() bool {
	p, err := s.Struct.Ptr(12)
	return p.IsValid() || err != nil
}

func (s CHF) SetIpv6SrcRoutePrefix(v []byte) error {
	return s.Struct.SetData(12, v)
}

func (s CHF) Ipv6DstRoutePrefix() ([]byte, error) {
	p, err := s.Struct.Ptr(13)
	return []byte(p.Data()), err
}

func (s CHF) HasIpv6DstRoutePrefix() bool {
	p, err := s.Struct.Ptr(13)
	return p.IsValid() || err != nil
}

func (s CHF) SetIpv6DstRoutePrefix(v []byte) error {
	return s.Struct.SetData(13, v)
}

func (s CHF) IsMetric() bool {
	return s.Struct.Bit(1378)
}

func (s CHF) SetIsMetric(v bool) {
	s.Struct.SetBit(1378, v)
}

func (s CHF) AppProtocol() uint32 {
	return s.Struct.Uint32(224) ^ 4294967295
}

func (s CHF) SetAppProtocol(v uint32) {
	s.Struct.SetUint32(224, v^4294967295)
}

// CHF_List is a list of CHF.
type CHF_List struct{ capnp.List }

// NewCHF creates a new list of CHF.
func NewCHF_List(s *capnp.Segment, sz int32) (CHF_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 232, PointerCount: 14}, sz)
	return CHF_List{l}, err
}

func (s CHF_List) At(i int) CHF { return CHF{s.List.Struct(i)} }

func (s CHF_List) Set(i int, v CHF) error { return s.List.SetStruct(i, v.Struct) }

func (s CHF_List) String() string {
	str, _ := text.MarshalList(0xa7ab5c68e4bc7b62, s.List)
	return str
}

// CHF_Promise is a wrapper for a CHF promised by a client call.
type CHF_Promise struct{ *capnp.Pipeline }

func (p CHF_Promise) Struct() (CHF, error) {
	s, err := p.Pipeline.Struct()
	return CHF{s}, err
}

type PackedCHF struct{ capnp.Struct }

// PackedCHF_TypeID is the unique identifier for the type PackedCHF.
const PackedCHF_TypeID = 0xb158a6a28e2d29c2

func NewPackedCHF(s *capnp.Segment) (PackedCHF, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PackedCHF{st}, err
}

func NewRootPackedCHF(s *capnp.Segment) (PackedCHF, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PackedCHF{st}, err
}

func ReadRootPackedCHF(msg *capnp.Message) (PackedCHF, error) {
	root, err := msg.RootPtr()
	return PackedCHF{root.Struct()}, err
}

func (s PackedCHF) String() string {
	str, _ := text.Marshal(0xb158a6a28e2d29c2, s.Struct)
	return str
}

func (s PackedCHF) Msgs() (CHF_List, error) {
	p, err := s.Struct.Ptr(0)
	return CHF_List{List: p.List()}, err
}

func (s PackedCHF) HasMsgs() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PackedCHF) SetMsgs(v CHF_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewMsgs sets the msgs field to a newly
// allocated CHF_List, preferring placement in s's segment.
func (s PackedCHF) NewMsgs(n int32) (CHF_List, error) {
	l, err := NewCHF_List(s.Struct.Segment(), n)
	if err != nil {
		return CHF_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// PackedCHF_List is a list of PackedCHF.
type PackedCHF_List struct{ capnp.List }

// NewPackedCHF creates a new list of PackedCHF.
func NewPackedCHF_List(s *capnp.Segment, sz int32) (PackedCHF_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return PackedCHF_List{l}, err
}

func (s PackedCHF_List) At(i int) PackedCHF { return PackedCHF{s.List.Struct(i)} }

func (s PackedCHF_List) Set(i int, v PackedCHF) error { return s.List.SetStruct(i, v.Struct) }

func (s PackedCHF_List) String() string {
	str, _ := text.MarshalList(0xb158a6a28e2d29c2, s.List)
	return str
}

// PackedCHF_Promise is a wrapper for a PackedCHF promised by a client call.
type PackedCHF_Promise struct{ *capnp.Pipeline }

func (p PackedCHF_Promise) Struct() (PackedCHF, error) {
	s, err := p.Pipeline.Struct()
	return PackedCHF{s}, err
}

const schema_c75f49ee0059f55d = "x\xdat\xd7kl\x1c\xd5\x15\x07\xf0\xff\x99\xd9\xf5\xec" +
	":~\x8d\xef\x0d\x98$\xc6Ip\xc8\x83$\xd8\x8e\x13" +
	"\x8c\x09\xf8\x91@\x13\x94\xd0\x1do\x88\x02\x0a*\x93\xdd" +
	"\xc5\x1eX\xefnw\xc6\xa9C\xa9\xa0\x15T\xb4\xa2-" +
	"\xad\xa8\x04\x15H\x14\xf5AUZ@\x0a\x12\xad\x82\x94" +
	"V\xa0\x06)\x95\x82\x94H\xa9\x94J\xa9\x94\x0f \x85" +
	"\xb6H\xa0> \xdd\xea\x7f\xed}\xd8\x98|\xca\xfe\xe6" +
	"\xdcs\xe7\x9e\xb9g<\xb7\xaf3>j\xf5\xc7'\x05" +
	"\xf0:\xe2M\x95#_?qi\xea\xf0\xaf\x7f\x09\xaf" +
	"Y\xba\xffu\xdf'\xf7\xfc}\xefW\xfe\x14os\x00" +
	"\xf7\xb9\xb2\xfb\x82\x03\xf4?\xf7j\x1b\xa4\xf2\xc7\x8d[" +
	"\xbe\xff\xd2/\x0e\xbd\x0e\xb7Yjq\xe2\x00\xaa\xdb}" +
	"I\xads\xf9\xbf5\xee\x08\xa4\xb2\xfe\x8c\xdd\xf9\xed\x9b" +
	"\xee\xfb\x90)\xed\x85\xa1\x9e\xfb]u\x8f\x09\xbd\xdb}" +
	"\x15\xf2\x8f+\x1f<\xf1\xc4\xc1\x17?\xf5\\\xb1\xeb\xc3" +
	"\xee\x16G\x1c\x89\xa9s\xee\xef!\xdb\xce\xb9?\x10H" +
	"%3\xf5\xc0\xd6\x8c_*\xa0\xa74\xbck\xcf\x1d)" +
	"\x91t\x8b\xd81 &\x80\xba\xdd)\x03\xe9\xdd\x8e-" +
	"\xe9\x94c\x89+\x96\x16\xfa~g\x00H\xef\xa1\x1f\xa0" +
	"[\xb6\x16\x8b\xf7\xe1\x0c\x03\xe9}\xf4Ct;\xa6\xc5" +
	"\xe6]\x19O\xd1\x0f\xd3cq-1@\xdd\xe3L\x00" +
	"\xe9C\xf4,=nk\x89\x03\xcaw\xc6\x81\xf4a\xfa" +
	"\x14\xbd)\xa6\xa5\x09P9\x93\xe7~z\x9e\xee4k" +
	"\xb3\xfe\xc0\xe4\x99\xa2G\xf4\xc42-\x09@}\xd5\xc4" +
	"\xe7\xe9\xb3\xf4d\x8b\x96$\xa0f\x9c#@:\xa2?" +
	"Fon\xd5\xd2\x0c\xa8o\x18\x7f\x84\xfe$}Y\x9b" +
	"\x96e\x80z\xc2\xe4\x7f\x9c\xfe4\xbd\xa5]K\x0b\xa0" +
	"\xbeg\xfc)\xfa\xb3\xf4\xd6\x0e-\xad\x80\xfa\xb1s/" +
	"\x90~\x86\xfe\"\xbd\xcd\xd5\xd2\x06\xa8\x17\x9c;\x81\xf4" +
	"\xf3\xf4\x97\xe9\xed\x9dZ\xda\x01\xf5s\xe7G@\xfae" +
	"\xfaqz\x87\xd2\xd2\x01\xa8\xd7M\x9d_\xa1\xbfIw" +
	"\xb5\x16\x17Po\x98u\xbdF?A\xef\\\xae\xa5\x13" +
	"P\xbf3~\x9c~\x92\xae\xae\xd2\xa2\x00\xf5\x96\x99\xf7" +
	"\x04\xfd\x14]_\xadE\x03\xea\x1dg-\x90>I?" +
	"M_\xde\xa5e9\xa0\xde5y\xde\xa6\x9f\xa1_u" +
	"\x8d\x96\xab\x00\xf5g\xf3\\N\xd1\xcf\xd2\xaf^\xa1\xe5" +
	"j@\xbdg\xeav\x86~\x81\xde\xb5RK\x17\xa0\xfe" +
	"b\xe6=O\xbfD\xbf\xa6M\xcb5\x80\xfa\x9b\xf1\x8b" +
	"\xf4\xcb\xf4\x15\xedZV\x00\xea\x03\x93\xff\x12\xfd\x9f\xf4" +
	"\x95\xab\xb5\xac\x04\xd4\x87f\x1f^\xa6\xff\xdb\xb1DV" +
	"iY\x05\xa8O\xcc\xb4\x1f;\xb6L$,q\xbbE" +
	"K7\xa0\xae\x18\xfe\x8c\xd1\x09\xfa\xb5k\xb4\\\x0b\xa8" +
	"x\x82O%\x96\xb0%\xddA\xefY\xab\xa5\x07P\xad" +
	"\x09\xdeM\x0b\xbd\x8b\xbe\xda\xd2\xb2\x1aP\xcb\x8dk\xfa" +
	"j\xfa\x9aN-k\xd8\x98\x09>\xf5U\xf4\x0d\xf4\xb5" +
	"\xb6\x96\xb5\x80Z\x97x\x10H\xf7\xd2\xfb\xe8\xd7\xc5\xb4" +
	"\\\x07\xa8-\x89o\x01\xe9\xcd\xf4!zo\\K/" +
	"\xa0\xb6\x9b\xf8A\xfa(}]\x93\x96u\x80\xba\xd5\xc4" +
	"\xef\xa4\xef\xa1_\x7f\x9d\x96\xeb\xd9\x8d&~7=E" +
	"_\x7f\xbd\x96\xf5\xecF\xe3\xfb\xe8\x87\xe8\x1b\xd6k\xd9" +
	"\xc0\xae3~\x80~?}\xe3\x06-\x1b\x01u\x9f\xf1" +
	"\xc3\xf4)\xfa\xa6\x8dZ6\xb1\xbbL}\xb2\xf4\x12\xfd" +
	"\x86MZn\x00\xd4\xb4\xf1<}6aI\xff\xe6\xfb" +
	"\xe3Z6\xb3\x8d\x12\xdc>%^x\x84\x17\xb6\xf8q" +
	"-[\x00u\xccTh\x96\x17\x1eg\xa6\xad\x9b\xb5l" +
	"\x05\xd47\x13\x0f\x03\xe9\xc7\xe8O\xd1o\xdc\xa2\xe5F" +
	"@}\xc7\xf8\x93\xf4g\xe8}[\xb5\xf4\x01\xea\x87\xc6" +
	"\x9f\xa6?O\xef\xbfQK?\xa0\x9e3\xfe,\xfdg" +
	"\xf4\x81\xdfh\x19\x00\xd4O\x8d\xbfH\x7f\x85\xbe\xed\xb7" +
	"Z\xb6\x01\xeaW\xc6_\xa6\x1f\xa7\x0f\xf6i\x19d\x7f" +
	"\x99J\xbcF?A\xdf\xde\xafe;\xfb\xc8\xf8\x9b\xf4" +
	"\xb7\xe9;\x06\xb4\xec\x00\xd4\x1f\x12\xdcY'\xe9\xa7\xe9" +
	"7m\xd3r\x13\xfb\xc5\xf8)\xfaY\xfa\x90\xa3e\x88" +
	"}a\xfc\x0c\xfd\x02\xfd\xe6\x84\x96\x9b\xd9\x17\xc6\xcf\xd3" +
	"/\xd1\x87Wh\x19f_\x98\xba]\xa4_\xa6\xdf\xb2" +
	"R\xcb-\xec\x0b\xe3\xef\xd3?\xa6\xefLj\xd9\x09\xa8" +
	"\x8f\x12\xec\xd3\xcb\xf4D\xd2\x12\xf7\xd6f-\xb7r\xa7" +
	"'\xb9\xdeX\x92;\x9d~\xdb2-\xb7q\xa7\x1bo" +
	"\xa1w\xd1GZ\xb4\x8cp\xa7'\x7f\x02\xa4\xbb\xe8\xbd" +
	"\xf4\xd1V-\xa3\xfc\xc3c\xbc\x97\xde\x97\xb4\xa4\x7f\xec" +
	"H\\\xcb\x18\xb7t\x92\xad\xb1\x99\x17\x868`|H" +
	"\xcb\xb8\x88\xda\x9e\xe4\xc2\x06\xe9\xa3IK*Q0\x9d" +
	"\x0b#\x7f\x1a=\xa5\xbb\xfcBQ\xe2\xb0$\x0e\xe9\xc9" +
	"\x86\xd1X(\x09X\x92\x80\x8cd\xc3\xe8K\xb9b\xe3" +
	"\xcf\xfd~\xa6\xfa\xb32\x95\xf3\xb3\xb9\xf2\xbe\x1c\xa4P" +
	"\xb5G\x83\xc2\xf8\xb1(\x17J\x12\x96$!#A!" +
	"\xf5PT\xfbY\x09\x0a\xa5\x99(U,C\xa2Z\xd6" +
	"\xa0\x94\x0e\x1e\xce\xd5\xb2\x06\xa5\xa3\x83\xbb\xc3h\x0cN" +
	"6[^\xa0\xe9rf\x91\xe6\x19\xb90]%\xcf\xb8" +
	"EV\x9c\x89\xcc\xbc\xb0\xcbu,\x95\x8bQ1S\xcc" +
	"\x03\xa8Y\xe8O\x97\xf2\xb9lJ\xfc\xccC\xb9(\x1d" +
	"<,\xb5\x1b\xeb\x09\xcb\x99\x86\xd2\x84\xe5Lci\xc2" +
	"r\xa6\xb14Q\xa6tG\xde\x9f\x0c\x1br;Q\xb1" +
	">\xfah\xde/\xec\xadW\x8d?\xbf<\x13-X\xed" +
	"]\xb9\xd9\x08\xce\x9eb\xa9\xa6\xd3\xa5|x\xe0X)" +
	"\xd7x\xc3\xc5\x99\xc8\x14\x9c6_\xe4G\x8b3\xd1\x82" +
	"\xa2G\x99\xd2D.*\xfb\xe8)\x84\xd3A}\x96\xb0" +
	"\x9c\xb9#_\xfc\xda\x018\xfed(-\xb0\xa4\x05R" +
	"\xc9\x86\xd1\x12:W\x98\x09\x1fvT\x7fT\xd9\xdc\xd1" +
	" \x93\xdb\x9bm\xbc\xa1\x078v~\xe5\xd5\xc1\xb5\xcd" +
	"&\xa5\xeaF\xe34\xe3\x93\xa5\xb1\x10\xed)?\x9aj" +
	"\x9c}|\xb2\xb4\xab(\xd3\xd33\x85 :VO\x12" +
	"\x963K\x0d\x98\xe3/\x18\xc0\x1a\xeeA{\xb1T\x7f" +
	"p\x9ca)\x9e{\x9e\x139\xb4O\x06\xc5Bc\xf4" +
	"R<\x17\xbd+\x80\x1d\x1d[\x14\xbb\x10\x9d#\xc1\xa4" +
	"\x08,\x91Z\x11\xc7\xb2\x90\x07k6\xbf\xdd\xef\xc2\x08" +
	"o\xaa\xe1i\xcf\xef\xf8\xcf_\x08\xcb\x99\x89\xe2L\x94" +
	"\xc3H\xaa\x9c{ \x98m\x9c\x7f\xe9\x0b\xf5\x11\xfbr" +
	"\x85\xc9hJ\x9a`I\xd3\x82\x11\x8b.\x84\xe5L:" +
	"\x97)\x16\xd0\x9e\x1d\x0b\x17Tc)\x0e\xcb\x99\x03S" +
	"A9\x0bgQ\xf0\x12\x1a\x94\x8e\xee\xa8\xb7w+," +
	"i\x9d\xd7z{W5,gn\x8f\xa6\xf6\xfb\x90L" +
	"m3g\xc3h\xb1\x8ddf\xc2\xa88-m\x90\x94" +
	"-\xd2Q\xff\xb2\x86\x10\xab\xd9\xeb\xa5l\x9cvA\xf1" +
	"\x17\xdd\xcf\x84\xb0:\xac\xa6\x1d\xcc.\x1e\xf5\x05\x17\xc3" +
	"\xfd\xb9\xa8\x1cd\xb8\xff\xab\xcf\xd8/\x95R|\xd9\xc0" +
	"\xc9\x14\xf3\xac\xc4G\x89J\xa5Ri\xf8\xca\x97\xd2p" +
	"\x8a\xaf\x9c\xac3\xf7\xad\xef\xc5\xaa\x9f\xfan\xeb&\xc0" +
	"K\xd8\xe2\xf5Z\xd2>\x1dN\x86\xf5u\xd6\xce2\xf3" +
	"\xebl\xc8\xb6\xcb\x94\x04`\xae\x96Z\xae\xdbW\x00\xde" +
	"\xa8-\xde>K\xf8\xafz\x1eq\xf7\x0e\xc0r\xad\x94" +
	"\xf9\xf8w\xfb\x8f\x00^\x9f-\xdeNK\xec [{" +
	"\x01\x1e\xf5\xf33\xb9J\x10\xee\x0e\xa6s\x85\x10\x0e\x1b" +
	"\xa2\xba\xc4\xcfM=\xb2\xd5\x84{\xab\xec\xd8\xaaJE" +
	"\xcc\xf1\xc4}c\x02\xf0\x8e\xdb\xe2\x9d\xb4\xa4[\xfeG" +
	"\xb6\x00\xf7\xad{\x01\xef\x84-\xde)KZ\xad+\x15" +
	"s6q\xdf\x19\x06\xbc\x93\xb6x\xa7-\xe9\xb6?\xab" +
	"X\xe6\x08\xe2\xbe\xcb\x1c\xa7l\xf1\xceZ\xd2\x1a\xfb\xb4" +
	"b\x0e \xee{\xe3\x80w\xda\x16\xef\xbc%\xdd\xf1\xff" +
	"2\x98G\xbas\x0c>k\x8bw\xd1\x92\xee\xa6\xffT" +
	"b\xe6\xf8\xe1\xfe\xf5N\xc0\xbb`\x8b\xf7\xbe%\x95\x99" +
	"\xa0\x10m\x1b8\xe8C\xf2\x8do1\x9f\x08\xdb\xcfK" +
	"3,i\xe6;>*\x1f\xf4\xf3\xb5\xf7\x0b\xc7\xed\x18" +
	"\x9c\x1bW}\xf5\xfa\xd9\xac\x89\xa9\xee\x06\xc6\xf4\xef\x98" +
	"\x8bq`\x893oC\x07}\xf3wg\xbe\xdf\xfe\x1f" +
	"\x00\x00\xff\xff\x1f\x00\xea "

func init() {
	schemas.Register(schema_c75f49ee0059f55d,
		0xa7ab5c68e4bc7b62,
		0xb158a6a28e2d29c2,
		0xed5d37861203d027,
		0xfba056008585e9fd)
}
