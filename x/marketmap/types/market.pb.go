// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: connect/marketmap/v2/market.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/skip-mev/connect/v2/pkg/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Market encapsulates a Ticker and its provider-specific configuration.
type Market struct {
	// Ticker represents a price feed for a given asset pair i.e. BTC/USD. The
	// price feed is scaled to a number of decimal places and has a minimum number
	// of providers required to consider the ticker valid.
	Ticker Ticker `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker"`
	// ProviderConfigs is the list of provider-specific configs for this Market.
	ProviderConfigs []ProviderConfig `protobuf:"bytes,2,rep,name=provider_configs,json=providerConfigs,proto3" json:"provider_configs"`
}

func (m *Market) Reset()      { *m = Market{} }
func (*Market) ProtoMessage() {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_54627e801f077fe4, []int{0}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetTicker() Ticker {
	if m != nil {
		return m.Ticker
	}
	return Ticker{}
}

func (m *Market) GetProviderConfigs() []ProviderConfig {
	if m != nil {
		return m.ProviderConfigs
	}
	return nil
}

// Ticker represents a price feed for a given asset pair i.e. BTC/USD. The price
// feed is scaled to a number of decimal places and has a minimum number of
// providers required to consider the ticker valid.
type Ticker struct {
	// CurrencyPair is the currency pair for this ticker.
	CurrencyPair types.CurrencyPair `protobuf:"bytes,1,opt,name=currency_pair,json=currencyPair,proto3" json:"currency_pair"`
	// Decimals is the number of decimal places for the ticker. The number of
	// decimal places is used to convert the price to a human-readable format.
	Decimals uint64 `protobuf:"varint,2,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// MinProviderCount is the minimum number of providers required to consider
	// the ticker valid.
	MinProviderCount uint64 `protobuf:"varint,3,opt,name=min_provider_count,json=minProviderCount,proto3" json:"min_provider_count,omitempty"`
	// Enabled is the flag that denotes if the Ticker is enabled for price
	// fetching by an oracle.
	Enabled bool `protobuf:"varint,14,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// MetadataJSON is a string of JSON that encodes any extra configuration
	// for the given ticker.
	Metadata_JSON string `protobuf:"bytes,15,opt,name=metadata_JSON,json=metadataJSON,proto3" json:"metadata_JSON,omitempty"`
}

func (m *Ticker) Reset()      { *m = Ticker{} }
func (*Ticker) ProtoMessage() {}
func (*Ticker) Descriptor() ([]byte, []int) {
	return fileDescriptor_54627e801f077fe4, []int{1}
}
func (m *Ticker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ticker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ticker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ticker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticker.Merge(m, src)
}
func (m *Ticker) XXX_Size() int {
	return m.Size()
}
func (m *Ticker) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticker.DiscardUnknown(m)
}

var xxx_messageInfo_Ticker proto.InternalMessageInfo

func (m *Ticker) GetCurrencyPair() types.CurrencyPair {
	if m != nil {
		return m.CurrencyPair
	}
	return types.CurrencyPair{}
}

func (m *Ticker) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *Ticker) GetMinProviderCount() uint64 {
	if m != nil {
		return m.MinProviderCount
	}
	return 0
}

func (m *Ticker) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Ticker) GetMetadata_JSON() string {
	if m != nil {
		return m.Metadata_JSON
	}
	return ""
}

type ProviderConfig struct {
	// Name corresponds to the name of the provider for which the configuration is
	// being set.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// OffChainTicker is the off-chain representation of the ticker i.e. BTC/USD.
	// The off-chain ticker is unique to a given provider and is used to fetch the
	// price of the ticker from the provider.
	OffChainTicker string `protobuf:"bytes,2,opt,name=off_chain_ticker,json=offChainTicker,proto3" json:"off_chain_ticker,omitempty"`
	// NormalizeByPair is the currency pair for this ticker to be normalized by.
	// For example, if the desired Ticker is BTC/USD, this market could be reached
	// using: OffChainTicker = BTC/USDT NormalizeByPair = USDT/USD This field is
	// optional and nullable.
	NormalizeByPair *types.CurrencyPair `protobuf:"bytes,3,opt,name=normalize_by_pair,json=normalizeByPair,proto3" json:"normalize_by_pair,omitempty"`
	// Invert is a boolean indicating if the BASE and QUOTE of the market should
	// be inverted. i.e. BASE -> QUOTE, QUOTE -> BASE
	Invert bool `protobuf:"varint,4,opt,name=invert,proto3" json:"invert,omitempty"`
	// MetadataJSON is a string of JSON that encodes any extra configuration
	// for the given provider config.
	Metadata_JSON string `protobuf:"bytes,15,opt,name=metadata_JSON,json=metadataJSON,proto3" json:"metadata_JSON,omitempty"`
}

func (m *ProviderConfig) Reset()         { *m = ProviderConfig{} }
func (m *ProviderConfig) String() string { return proto.CompactTextString(m) }
func (*ProviderConfig) ProtoMessage()    {}
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_54627e801f077fe4, []int{2}
}
func (m *ProviderConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderConfig.Merge(m, src)
}
func (m *ProviderConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProviderConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderConfig proto.InternalMessageInfo

func (m *ProviderConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProviderConfig) GetOffChainTicker() string {
	if m != nil {
		return m.OffChainTicker
	}
	return ""
}

func (m *ProviderConfig) GetNormalizeByPair() *types.CurrencyPair {
	if m != nil {
		return m.NormalizeByPair
	}
	return nil
}

func (m *ProviderConfig) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *ProviderConfig) GetMetadata_JSON() string {
	if m != nil {
		return m.Metadata_JSON
	}
	return ""
}

// MarketMap maps ticker strings to their Markets.
type MarketMap struct {
	// Markets is the full list of tickers and their associated configurations
	// to be stored on-chain.
	Markets map[string]Market `protobuf:"bytes,1,rep,name=markets,proto3" json:"markets" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MarketMap) Reset()      { *m = MarketMap{} }
func (*MarketMap) ProtoMessage() {}
func (*MarketMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_54627e801f077fe4, []int{3}
}
func (m *MarketMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarketMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarketMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarketMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketMap.Merge(m, src)
}
func (m *MarketMap) XXX_Size() int {
	return m.Size()
}
func (m *MarketMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketMap.DiscardUnknown(m)
}

var xxx_messageInfo_MarketMap proto.InternalMessageInfo

func (m *MarketMap) GetMarkets() map[string]Market {
	if m != nil {
		return m.Markets
	}
	return nil
}

func init() {
	proto.RegisterType((*Market)(nil), "connect.marketmap.v2.Market")
	proto.RegisterType((*Ticker)(nil), "connect.marketmap.v2.Ticker")
	proto.RegisterType((*ProviderConfig)(nil), "connect.marketmap.v2.ProviderConfig")
	proto.RegisterType((*MarketMap)(nil), "connect.marketmap.v2.MarketMap")
	proto.RegisterMapType((map[string]Market)(nil), "connect.marketmap.v2.MarketMap.MarketsEntry")
}

func init() { proto.RegisterFile("connect/marketmap/v2/market.proto", fileDescriptor_54627e801f077fe4) }

var fileDescriptor_54627e801f077fe4 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x8e, 0xd7, 0xd2, 0xad, 0x5e, 0xd7, 0x16, 0x6b, 0x42, 0x51, 0x85, 0xb2, 0x52, 0x76, 0xc8,
	0x61, 0x24, 0x28, 0x5c, 0x50, 0x8f, 0xad, 0x38, 0x50, 0x31, 0x98, 0x02, 0x48, 0x88, 0x4b, 0xe4,
	0xa6, 0x6e, 0x67, 0xb5, 0xb1, 0x23, 0xc7, 0x8d, 0x28, 0x27, 0x7e, 0x02, 0x47, 0x8e, 0x5c, 0xf8,
	0x19, 0xdc, 0x77, 0xdc, 0x0d, 0x0e, 0x08, 0xa1, 0x56, 0xe2, 0x77, 0xa0, 0x38, 0x4e, 0xd7, 0x4a,
	0x45, 0xda, 0xed, 0xbd, 0xe7, 0xcf, 0xdf, 0xf3, 0xf7, 0xf9, 0x3d, 0xf8, 0x20, 0xe4, 0x8c, 0x91,
	0x50, 0xba, 0x11, 0x16, 0x53, 0x22, 0x23, 0x1c, 0xbb, 0xa9, 0xa7, 0x13, 0x27, 0x16, 0x5c, 0x72,
	0x74, 0xac, 0x21, 0xce, 0x1a, 0xe2, 0xa4, 0x5e, 0xeb, 0x78, 0xc2, 0x27, 0x5c, 0x01, 0xdc, 0x2c,
	0xca, 0xb1, 0xad, 0xd3, 0x82, 0x4e, 0x2e, 0x62, 0x92, 0x64, 0x54, 0xe1, 0x5c, 0x08, 0xc2, 0xc2,
	0x45, 0x10, 0x63, 0x2a, 0x72, 0x54, 0xe7, 0x1b, 0x80, 0x95, 0x73, 0x45, 0x86, 0xba, 0xb0, 0x22,
	0x69, 0x38, 0x25, 0xc2, 0x04, 0x6d, 0x60, 0x1f, 0x7a, 0xf7, 0x9d, 0x5d, 0xdd, 0x9c, 0x37, 0x0a,
	0xd3, 0x2b, 0x5f, 0xfd, 0x3e, 0x31, 0x7c, 0x7d, 0x03, 0xbd, 0x85, 0xcd, 0x58, 0xf0, 0x94, 0x8e,
	0x88, 0x08, 0x42, 0xce, 0xc6, 0x74, 0x92, 0x98, 0x7b, 0xed, 0x92, 0x7d, 0xe8, 0x9d, 0xee, 0x66,
	0xb9, 0xd0, 0xe8, 0xbe, 0x02, 0x6b, 0xb6, 0x46, 0xbc, 0x55, 0x4d, 0xba, 0x07, 0x5f, 0xbe, 0x9e,
	0x18, 0x9f, 0x7e, 0xb5, 0x8d, 0xce, 0x5f, 0x00, 0x2b, 0x79, 0x67, 0xf4, 0x1c, 0x1e, 0x6d, 0x29,
	0xd1, 0xcf, 0xb5, 0xd6, 0x8d, 0x94, 0xe0, 0xac, 0x49, 0x5f, 0xc3, 0x2e, 0x30, 0x2d, 0x1e, 0x5c,
	0x0b, 0x37, 0x6a, 0xa8, 0x05, 0x0f, 0x46, 0x24, 0xa4, 0x11, 0x9e, 0x65, 0xcf, 0x05, 0x76, 0xd9,
	0x5f, 0xe7, 0xe8, 0x0c, 0xa2, 0x88, 0xb2, 0x60, 0x43, 0xd6, 0x9c, 0x49, 0xb3, 0xa4, 0x50, 0xcd,
	0x88, 0xb2, 0x1b, 0x05, 0x73, 0x26, 0x91, 0x09, 0xf7, 0x09, 0xc3, 0xc3, 0x19, 0x19, 0x99, 0xf5,
	0x36, 0xb0, 0x0f, 0xfc, 0x22, 0x45, 0x0f, 0xe1, 0x51, 0x44, 0x24, 0x1e, 0x61, 0x89, 0x83, 0xc1,
	0xeb, 0x57, 0x2f, 0xcd, 0x46, 0x1b, 0xd8, 0x55, 0xbf, 0x56, 0x14, 0xb3, 0xda, 0x86, 0xd0, 0x1f,
	0x00, 0xd6, 0xb7, 0xcd, 0x41, 0x08, 0x96, 0x19, 0x8e, 0x88, 0xd2, 0x59, 0xf5, 0x55, 0x8c, 0x6c,
	0xd8, 0xe4, 0xe3, 0x71, 0x10, 0x5e, 0x62, 0xca, 0x02, 0xfd, 0x6d, 0x7b, 0xea, 0xbc, 0xce, 0xc7,
	0xe3, 0x7e, 0x56, 0xd6, 0x76, 0x0d, 0xe0, 0x5d, 0xc6, 0x45, 0x84, 0x67, 0xf4, 0x23, 0x09, 0x86,
	0xda, 0xb2, 0xd2, 0x6d, 0x2c, 0xf3, 0x1b, 0xeb, 0x8b, 0xbd, 0xdc, 0xaf, 0x7b, 0xb0, 0x42, 0x59,
	0x4a, 0x84, 0x34, 0xcb, 0x4a, 0xa4, 0xce, 0x6e, 0xa5, 0xb1, 0xf3, 0x1d, 0xc0, 0x6a, 0x3e, 0x6a,
	0xe7, 0x38, 0x46, 0x2f, 0xe0, 0x7e, 0x3e, 0x10, 0x89, 0x09, 0xd4, 0xa0, 0x9c, 0xed, 0x1e, 0x94,
	0xf5, 0x0d, 0x1d, 0x25, 0xcf, 0x98, 0x14, 0x0b, 0xfd, 0x9b, 0x05, 0x45, 0xeb, 0x1d, 0xac, 0x6d,
	0x1e, 0xa3, 0x26, 0x2c, 0x4d, 0xc9, 0x42, 0x3b, 0x96, 0x85, 0xc8, 0x83, 0x77, 0x52, 0x3c, 0x9b,
	0x13, 0xe5, 0xd2, 0x7f, 0x87, 0x3b, 0x27, 0xf1, 0x73, 0x68, 0x77, 0xef, 0x29, 0xb8, 0xf9, 0x99,
	0xde, 0xe0, 0x6a, 0x69, 0x81, 0xeb, 0xa5, 0x05, 0xfe, 0x2c, 0x2d, 0xf0, 0x79, 0x65, 0x19, 0xd7,
	0x2b, 0xcb, 0xf8, 0xb9, 0xb2, 0x8c, 0xf7, 0x8f, 0x27, 0x54, 0x5e, 0xce, 0x87, 0x4e, 0xc8, 0x23,
	0x37, 0x99, 0xd2, 0xf8, 0x51, 0x44, 0x52, 0xb7, 0x58, 0xbf, 0xd4, 0x73, 0x3f, 0x6c, 0xac, 0xb4,
	0x72, 0x7a, 0x58, 0x51, 0xdb, 0xf7, 0xe4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xa9, 0x16,
	0x4f, 0xf4, 0x03, 0x00, 0x00,
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProviderConfigs) > 0 {
		for iNdEx := len(m.ProviderConfigs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProviderConfigs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Ticker.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Ticker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ticker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ticker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata_JSON) > 0 {
		i -= len(m.Metadata_JSON)
		copy(dAtA[i:], m.Metadata_JSON)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Metadata_JSON)))
		i--
		dAtA[i] = 0x7a
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	if m.MinProviderCount != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.MinProviderCount))
		i--
		dAtA[i] = 0x18
	}
	if m.Decimals != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.CurrencyPair.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMarket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ProviderConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata_JSON) > 0 {
		i -= len(m.Metadata_JSON)
		copy(dAtA[i:], m.Metadata_JSON)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Metadata_JSON)))
		i--
		dAtA[i] = 0x7a
	}
	if m.Invert {
		i--
		if m.Invert {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.NormalizeByPair != nil {
		{
			size, err := m.NormalizeByPair.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMarket(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OffChainTicker) > 0 {
		i -= len(m.OffChainTicker)
		copy(dAtA[i:], m.OffChainTicker)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.OffChainTicker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMarket(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarketMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarketMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarketMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Markets) > 0 {
		for k := range m.Markets {
			v := m.Markets[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMarket(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintMarket(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintMarket(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarket(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Ticker.Size()
	n += 1 + l + sovMarket(uint64(l))
	if len(m.ProviderConfigs) > 0 {
		for _, e := range m.ProviderConfigs {
			l = e.Size()
			n += 1 + l + sovMarket(uint64(l))
		}
	}
	return n
}

func (m *Ticker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CurrencyPair.Size()
	n += 1 + l + sovMarket(uint64(l))
	if m.Decimals != 0 {
		n += 1 + sovMarket(uint64(m.Decimals))
	}
	if m.MinProviderCount != 0 {
		n += 1 + sovMarket(uint64(m.MinProviderCount))
	}
	if m.Enabled {
		n += 2
	}
	l = len(m.Metadata_JSON)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	return n
}

func (m *ProviderConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	l = len(m.OffChainTicker)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	if m.NormalizeByPair != nil {
		l = m.NormalizeByPair.Size()
		n += 1 + l + sovMarket(uint64(l))
	}
	if m.Invert {
		n += 2
	}
	l = len(m.Metadata_JSON)
	if l > 0 {
		n += 1 + l + sovMarket(uint64(l))
	}
	return n
}

func (m *MarketMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Markets) > 0 {
		for k, v := range m.Markets {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovMarket(uint64(len(k))) + 1 + l + sovMarket(uint64(l))
			n += mapEntrySize + 1 + sovMarket(uint64(mapEntrySize))
		}
	}
	return n
}

func sovMarket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarket(x uint64) (n int) {
	return sovMarket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Ticker.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderConfigs = append(m.ProviderConfigs, ProviderConfig{})
			if err := m.ProviderConfigs[len(m.ProviderConfigs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func (m *Ticker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Ticker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ticker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrencyPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrencyPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinProviderCount", wireType)
			}
			m.MinProviderCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinProviderCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata_JSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata_JSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func (m *ProviderConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProviderConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OffChainTicker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OffChainTicker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NormalizeByPair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NormalizeByPair == nil {
				m.NormalizeByPair = &types.CurrencyPair{}
			}
			if err := m.NormalizeByPair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Invert", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Invert = bool(v != 0)
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata_JSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata_JSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func (m *MarketMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MarketMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarketMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMarket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMarket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Markets == nil {
				m.Markets = make(map[string]Market)
			}
			var mapkey string
			mapvalue := &Market{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMarket
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMarket
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthMarket
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthMarket
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMarket
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthMarket
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthMarket
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Market{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipMarket(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthMarket
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Markets[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func skipMarket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarket
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMarket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarket = fmt.Errorf("proto: unexpected end of group")
)