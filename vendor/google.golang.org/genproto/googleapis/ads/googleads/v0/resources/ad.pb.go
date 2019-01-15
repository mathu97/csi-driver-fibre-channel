// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import common "google.golang.org/genproto/googleapis/ads/googleads/v0/common"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An ad.
type Ad struct {
	// The ID of the ad.
	Id *wrappers.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The list of possible final URLs after all cross-domain redirects for the
	// ad.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,2,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// The list of possible final mobile URLs after all cross-domain redirects
	// for the ad.
	FinalMobileUrls []*wrappers.StringValue `protobuf:"bytes,16,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,12,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings that can be used to substitute custom parameter tags
	// in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,10,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The URL that appears in the ad description for some ad formats.
	DisplayUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=display_url,json=displayUrl,proto3" json:"display_url,omitempty"`
	// The type of ad.
	Type enums.AdTypeEnum_AdType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.AdTypeEnum_AdType" json:"type,omitempty"`
	// Indicates if this ad was automatically added by Google Ads and not by a
	// user. For example, this could happen when ads are automatically created as
	// suggestions for new ads based on knowledge of how existing ads are
	// performing.
	AddedByGoogleAds *wrappers.BoolValue `protobuf:"bytes,19,opt,name=added_by_google_ads,json=addedByGoogleAds,proto3" json:"added_by_google_ads,omitempty"`
	// The device preference for the ad. You can only specify a preference for
	// mobile devices. When this preference is set the ad will be preferred over
	// other ads when being displayed on a mobile device. The ad can still be
	// displayed on other device types, e.g. if no other ads are available.
	// If unspecified (no device preference), all devices are targeted.
	// This is only supported by some ad types.
	DevicePreference enums.DeviceEnum_Device `protobuf:"varint,20,opt,name=device_preference,json=devicePreference,proto3,enum=google.ads.googleads.v0.enums.DeviceEnum_Device" json:"device_preference,omitempty"`
	// The name of the ad. This is only used to be able to identify the ad. It
	// does not need to be unique and does not affect the served ad.
	Name *wrappers.StringValue `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// Details pertinent to the ad type. Exactly one value must be set.
	//
	// Types that are valid to be assigned to AdData:
	//	*Ad_TextAd
	//	*Ad_ExpandedTextAd
	//	*Ad_DynamicSearchAd
	//	*Ad_ResponsiveDisplayAd
	//	*Ad_CallOnlyAd
	//	*Ad_ExpandedDynamicSearchAd
	//	*Ad_HotelAd
	//	*Ad_ShoppingSmartAd
	//	*Ad_ShoppingProductAd
	//	*Ad_GmailAd
	//	*Ad_ImageAd
	AdData               isAd_AdData `protobuf_oneof:"ad_data"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad_81ea002d2bacbe35, []int{0}
}
func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (dst *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(dst, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ad) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *Ad) GetFinalMobileUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalMobileUrls
	}
	return nil
}

func (m *Ad) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Ad) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *Ad) GetDisplayUrl() *wrappers.StringValue {
	if m != nil {
		return m.DisplayUrl
	}
	return nil
}

func (m *Ad) GetType() enums.AdTypeEnum_AdType {
	if m != nil {
		return m.Type
	}
	return enums.AdTypeEnum_UNSPECIFIED
}

func (m *Ad) GetAddedByGoogleAds() *wrappers.BoolValue {
	if m != nil {
		return m.AddedByGoogleAds
	}
	return nil
}

func (m *Ad) GetDevicePreference() enums.DeviceEnum_Device {
	if m != nil {
		return m.DevicePreference
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Ad) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type isAd_AdData interface {
	isAd_AdData()
}

type Ad_TextAd struct {
	TextAd *common.TextAdInfo `protobuf:"bytes,6,opt,name=text_ad,json=textAd,proto3,oneof"`
}

type Ad_ExpandedTextAd struct {
	ExpandedTextAd *common.ExpandedTextAdInfo `protobuf:"bytes,7,opt,name=expanded_text_ad,json=expandedTextAd,proto3,oneof"`
}

type Ad_DynamicSearchAd struct {
	DynamicSearchAd *common.DynamicSearchAdInfo `protobuf:"bytes,8,opt,name=dynamic_search_ad,json=dynamicSearchAd,proto3,oneof"`
}

type Ad_ResponsiveDisplayAd struct {
	ResponsiveDisplayAd *common.ResponsiveDisplayAdInfo `protobuf:"bytes,9,opt,name=responsive_display_ad,json=responsiveDisplayAd,proto3,oneof"`
}

type Ad_CallOnlyAd struct {
	CallOnlyAd *common.CallOnlyAdInfo `protobuf:"bytes,13,opt,name=call_only_ad,json=callOnlyAd,proto3,oneof"`
}

type Ad_ExpandedDynamicSearchAd struct {
	ExpandedDynamicSearchAd *common.ExpandedDynamicSearchAdInfo `protobuf:"bytes,14,opt,name=expanded_dynamic_search_ad,json=expandedDynamicSearchAd,proto3,oneof"`
}

type Ad_HotelAd struct {
	HotelAd *common.HotelAdInfo `protobuf:"bytes,15,opt,name=hotel_ad,json=hotelAd,proto3,oneof"`
}

type Ad_ShoppingSmartAd struct {
	ShoppingSmartAd *common.ShoppingSmartAdInfo `protobuf:"bytes,17,opt,name=shopping_smart_ad,json=shoppingSmartAd,proto3,oneof"`
}

type Ad_ShoppingProductAd struct {
	ShoppingProductAd *common.ShoppingProductAdInfo `protobuf:"bytes,18,opt,name=shopping_product_ad,json=shoppingProductAd,proto3,oneof"`
}

type Ad_GmailAd struct {
	GmailAd *common.GmailAdInfo `protobuf:"bytes,21,opt,name=gmail_ad,json=gmailAd,proto3,oneof"`
}

type Ad_ImageAd struct {
	ImageAd *common.ImageAdInfo `protobuf:"bytes,22,opt,name=image_ad,json=imageAd,proto3,oneof"`
}

func (*Ad_TextAd) isAd_AdData() {}

func (*Ad_ExpandedTextAd) isAd_AdData() {}

func (*Ad_DynamicSearchAd) isAd_AdData() {}

func (*Ad_ResponsiveDisplayAd) isAd_AdData() {}

func (*Ad_CallOnlyAd) isAd_AdData() {}

func (*Ad_ExpandedDynamicSearchAd) isAd_AdData() {}

func (*Ad_HotelAd) isAd_AdData() {}

func (*Ad_ShoppingSmartAd) isAd_AdData() {}

func (*Ad_ShoppingProductAd) isAd_AdData() {}

func (*Ad_GmailAd) isAd_AdData() {}

func (*Ad_ImageAd) isAd_AdData() {}

func (m *Ad) GetAdData() isAd_AdData {
	if m != nil {
		return m.AdData
	}
	return nil
}

func (m *Ad) GetTextAd() *common.TextAdInfo {
	if x, ok := m.GetAdData().(*Ad_TextAd); ok {
		return x.TextAd
	}
	return nil
}

func (m *Ad) GetExpandedTextAd() *common.ExpandedTextAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedTextAd); ok {
		return x.ExpandedTextAd
	}
	return nil
}

func (m *Ad) GetDynamicSearchAd() *common.DynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_DynamicSearchAd); ok {
		return x.DynamicSearchAd
	}
	return nil
}

func (m *Ad) GetResponsiveDisplayAd() *common.ResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveDisplayAd); ok {
		return x.ResponsiveDisplayAd
	}
	return nil
}

func (m *Ad) GetCallOnlyAd() *common.CallOnlyAdInfo {
	if x, ok := m.GetAdData().(*Ad_CallOnlyAd); ok {
		return x.CallOnlyAd
	}
	return nil
}

func (m *Ad) GetExpandedDynamicSearchAd() *common.ExpandedDynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedDynamicSearchAd); ok {
		return x.ExpandedDynamicSearchAd
	}
	return nil
}

func (m *Ad) GetHotelAd() *common.HotelAdInfo {
	if x, ok := m.GetAdData().(*Ad_HotelAd); ok {
		return x.HotelAd
	}
	return nil
}

func (m *Ad) GetShoppingSmartAd() *common.ShoppingSmartAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingSmartAd); ok {
		return x.ShoppingSmartAd
	}
	return nil
}

func (m *Ad) GetShoppingProductAd() *common.ShoppingProductAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingProductAd); ok {
		return x.ShoppingProductAd
	}
	return nil
}

func (m *Ad) GetGmailAd() *common.GmailAdInfo {
	if x, ok := m.GetAdData().(*Ad_GmailAd); ok {
		return x.GmailAd
	}
	return nil
}

func (m *Ad) GetImageAd() *common.ImageAdInfo {
	if x, ok := m.GetAdData().(*Ad_ImageAd); ok {
		return x.ImageAd
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Ad) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Ad_OneofMarshaler, _Ad_OneofUnmarshaler, _Ad_OneofSizer, []interface{}{
		(*Ad_TextAd)(nil),
		(*Ad_ExpandedTextAd)(nil),
		(*Ad_DynamicSearchAd)(nil),
		(*Ad_ResponsiveDisplayAd)(nil),
		(*Ad_CallOnlyAd)(nil),
		(*Ad_ExpandedDynamicSearchAd)(nil),
		(*Ad_HotelAd)(nil),
		(*Ad_ShoppingSmartAd)(nil),
		(*Ad_ShoppingProductAd)(nil),
		(*Ad_GmailAd)(nil),
		(*Ad_ImageAd)(nil),
	}
}

func _Ad_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Ad)
	// ad_data
	switch x := m.AdData.(type) {
	case *Ad_TextAd:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextAd); err != nil {
			return err
		}
	case *Ad_ExpandedTextAd:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExpandedTextAd); err != nil {
			return err
		}
	case *Ad_DynamicSearchAd:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DynamicSearchAd); err != nil {
			return err
		}
	case *Ad_ResponsiveDisplayAd:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResponsiveDisplayAd); err != nil {
			return err
		}
	case *Ad_CallOnlyAd:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CallOnlyAd); err != nil {
			return err
		}
	case *Ad_ExpandedDynamicSearchAd:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExpandedDynamicSearchAd); err != nil {
			return err
		}
	case *Ad_HotelAd:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HotelAd); err != nil {
			return err
		}
	case *Ad_ShoppingSmartAd:
		b.EncodeVarint(17<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ShoppingSmartAd); err != nil {
			return err
		}
	case *Ad_ShoppingProductAd:
		b.EncodeVarint(18<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ShoppingProductAd); err != nil {
			return err
		}
	case *Ad_GmailAd:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GmailAd); err != nil {
			return err
		}
	case *Ad_ImageAd:
		b.EncodeVarint(22<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageAd); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Ad.AdData has unexpected type %T", x)
	}
	return nil
}

func _Ad_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Ad)
	switch tag {
	case 6: // ad_data.text_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.TextAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_TextAd{msg}
		return true, err
	case 7: // ad_data.expanded_text_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ExpandedTextAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ExpandedTextAd{msg}
		return true, err
	case 8: // ad_data.dynamic_search_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.DynamicSearchAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_DynamicSearchAd{msg}
		return true, err
	case 9: // ad_data.responsive_display_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ResponsiveDisplayAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ResponsiveDisplayAd{msg}
		return true, err
	case 13: // ad_data.call_only_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.CallOnlyAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_CallOnlyAd{msg}
		return true, err
	case 14: // ad_data.expanded_dynamic_search_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ExpandedDynamicSearchAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ExpandedDynamicSearchAd{msg}
		return true, err
	case 15: // ad_data.hotel_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.HotelAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_HotelAd{msg}
		return true, err
	case 17: // ad_data.shopping_smart_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ShoppingSmartAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ShoppingSmartAd{msg}
		return true, err
	case 18: // ad_data.shopping_product_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ShoppingProductAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ShoppingProductAd{msg}
		return true, err
	case 21: // ad_data.gmail_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.GmailAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_GmailAd{msg}
		return true, err
	case 22: // ad_data.image_ad
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.ImageAdInfo)
		err := b.DecodeMessage(msg)
		m.AdData = &Ad_ImageAd{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Ad_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Ad)
	// ad_data
	switch x := m.AdData.(type) {
	case *Ad_TextAd:
		s := proto.Size(x.TextAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ExpandedTextAd:
		s := proto.Size(x.ExpandedTextAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_DynamicSearchAd:
		s := proto.Size(x.DynamicSearchAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ResponsiveDisplayAd:
		s := proto.Size(x.ResponsiveDisplayAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_CallOnlyAd:
		s := proto.Size(x.CallOnlyAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ExpandedDynamicSearchAd:
		s := proto.Size(x.ExpandedDynamicSearchAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_HotelAd:
		s := proto.Size(x.HotelAd)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ShoppingSmartAd:
		s := proto.Size(x.ShoppingSmartAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ShoppingProductAd:
		s := proto.Size(x.ShoppingProductAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_GmailAd:
		s := proto.Size(x.GmailAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Ad_ImageAd:
		s := proto.Size(x.ImageAd)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Ad)(nil), "google.ads.googleads.v0.resources.Ad")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad.proto", fileDescriptor_ad_81ea002d2bacbe35)
}

var fileDescriptor_ad_81ea002d2bacbe35 = []byte{
	// 841 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdf, 0x6e, 0xe4, 0x34,
	0x14, 0xc6, 0x99, 0xd9, 0xd2, 0xd9, 0xba, 0xbb, 0xfd, 0xe3, 0xa1, 0x6c, 0x54, 0x10, 0xea, 0x22,
	0x21, 0x55, 0xad, 0xe4, 0x19, 0xcd, 0xb2, 0x70, 0x51, 0x71, 0x91, 0xa1, 0x55, 0x5b, 0x24, 0xc4,
	0x68, 0xfa, 0xe7, 0x02, 0x15, 0x22, 0x37, 0xf6, 0xa4, 0x11, 0x8e, 0x1d, 0xd9, 0x49, 0xe9, 0xf0,
	0x38, 0x5c, 0xf2, 0x28, 0x08, 0xf1, 0x4c, 0x2b, 0x1f, 0x27, 0xae, 0xba, 0xed, 0x34, 0xb9, 0xb3,
	0x9d, 0xef, 0xf7, 0x7d, 0x67, 0x8e, 0x1d, 0x67, 0xd0, 0x5e, 0xa2, 0x54, 0x22, 0xf8, 0x80, 0x32,
	0x33, 0x70, 0x43, 0x3b, 0xba, 0x1d, 0x0e, 0x34, 0x37, 0xaa, 0xd4, 0x31, 0x37, 0x03, 0xca, 0x48,
	0xae, 0x55, 0xa1, 0xf0, 0x5b, 0x27, 0x20, 0x94, 0x19, 0xe2, 0xb5, 0xe4, 0x76, 0x48, 0xbc, 0x76,
	0x7b, 0xb4, 0xc8, 0x2e, 0x56, 0x59, 0xa6, 0xe4, 0x80, 0xb2, 0xa8, 0x98, 0xe7, 0x3c, 0x4a, 0xe5,
	0x4c, 0x19, 0x67, 0xbb, 0xfd, 0xbe, 0x81, 0x89, 0x4b, 0x53, 0xa8, 0x2c, 0xca, 0xa9, 0xa6, 0x19,
	0x2f, 0xb8, 0xae, 0xb0, 0xfd, 0x45, 0x18, 0x97, 0x65, 0x66, 0xea, 0xa4, 0x4a, 0xbc, 0xf7, 0xbc,
	0x98, 0xf1, 0xdb, 0x34, 0xae, 0xb5, 0x5f, 0x55, 0x5a, 0x98, 0x5d, 0x97, 0xb3, 0xc1, 0x9f, 0x9a,
	0xe6, 0x39, 0xd7, 0x55, 0xbd, 0x5f, 0xff, 0xff, 0x1a, 0x75, 0x43, 0x86, 0xf7, 0x51, 0x37, 0x65,
	0x41, 0x67, 0xa7, 0xb3, 0xbb, 0x3a, 0xfa, 0xa2, 0xea, 0x07, 0xa9, 0x19, 0x72, 0x2a, 0x8b, 0xef,
	0xbe, 0xbd, 0xa4, 0xa2, 0xe4, 0xd3, 0x6e, 0xca, 0xf0, 0x01, 0x42, 0xb3, 0x54, 0x52, 0x11, 0x95,
	0x5a, 0x98, 0xa0, 0xbb, 0xf3, 0x62, 0x77, 0x75, 0xf4, 0xe5, 0x23, 0xe8, 0xac, 0xd0, 0xa9, 0x4c,
	0x1c, 0xb5, 0x02, 0xfa, 0x0b, 0x2d, 0x0c, 0x3e, 0x41, 0x9b, 0x0e, 0xce, 0xd4, 0x75, 0x2a, 0xb8,
	0xf3, 0xd8, 0x68, 0xe1, 0xb1, 0x0e, 0xd8, 0xcf, 0x40, 0x81, 0xd3, 0x04, 0x6d, 0x15, 0x9a, 0xc6,
	0x7f, 0xa4, 0x32, 0xb1, 0x2e, 0x51, 0xc1, 0xb3, 0x5c, 0xd0, 0x82, 0x07, 0xaf, 0xe0, 0x67, 0x3c,
	0xef, 0xd6, 0xaf, 0xd1, 0x0b, 0x2d, 0xce, 0x2b, 0x10, 0xc7, 0x68, 0xcb, 0x1a, 0x7d, 0xbc, 0x47,
	0x26, 0x40, 0x50, 0xdf, 0x80, 0x2c, 0x3a, 0x33, 0x6e, 0x73, 0xc9, 0x8f, 0x00, 0x4e, 0x6a, 0x6e,
	0xda, 0x2f, 0xb5, 0xf8, 0x68, 0xcd, 0xe0, 0x1f, 0xd0, 0x2a, 0x4b, 0x4d, 0x2e, 0xe8, 0xdc, 0x56,
	0x1d, 0x2c, 0xb5, 0x28, 0x16, 0x55, 0xc0, 0x85, 0x16, 0xf8, 0x10, 0x2d, 0xd9, 0xa3, 0x10, 0x7c,
	0xba, 0xd3, 0xd9, 0x5d, 0x1b, 0x0d, 0x17, 0x96, 0x04, 0x67, 0x81, 0x84, 0xec, 0x7c, 0x9e, 0xf3,
	0x23, 0x59, 0x66, 0xd5, 0x70, 0x0a, 0x34, 0x3e, 0x45, 0x7d, 0xca, 0x18, 0x67, 0xd1, 0xf5, 0x3c,
	0x72, 0x58, 0x44, 0x99, 0x09, 0xfa, 0x50, 0xcc, 0xf6, 0xa3, 0x62, 0xc6, 0x4a, 0x09, 0x57, 0xca,
	0x06, 0x60, 0xe3, 0xf9, 0x31, 0x28, 0x42, 0x66, 0xf0, 0x6f, 0x68, 0xd3, 0x9d, 0xb8, 0x28, 0xd7,
	0x7c, 0xc6, 0x35, 0x97, 0x31, 0x0f, 0x3e, 0x6b, 0x55, 0xdd, 0x21, 0x70, 0x50, 0x9d, 0x1b, 0x4e,
	0x37, 0x9c, 0xd5, 0xc4, 0x3b, 0xe1, 0x21, 0x5a, 0x92, 0x34, 0xe3, 0xc1, 0x9b, 0x16, 0x7d, 0x02,
	0x25, 0x3e, 0x42, 0xbd, 0x82, 0xdf, 0x15, 0x11, 0x65, 0xc1, 0x32, 0x40, 0x7b, 0x4d, 0xfb, 0x76,
	0xce, 0xef, 0x8a, 0x90, 0x9d, 0xca, 0x99, 0x3a, 0xf9, 0x64, 0xba, 0x5c, 0xc0, 0x0c, 0xff, 0x8e,
	0x36, 0xf8, 0x5d, 0x4e, 0xa5, 0xed, 0x52, 0xed, 0xd7, 0x03, 0xbf, 0x51, 0x93, 0xdf, 0x51, 0xc5,
	0x3d, 0xf0, 0x5d, 0xe3, 0x0f, 0x56, 0x31, 0x45, 0x9b, 0x6c, 0x2e, 0x69, 0x96, 0xc6, 0x91, 0xe1,
	0x54, 0xc7, 0x37, 0x36, 0xe0, 0x25, 0x04, 0xbc, 0x6b, 0x0a, 0x38, 0x74, 0xe0, 0x19, 0x70, 0x3e,
	0x61, 0x9d, 0x3d, 0x5c, 0xc6, 0x19, 0xda, 0xd2, 0xdc, 0xe4, 0x4a, 0x9a, 0xf4, 0x96, 0x47, 0xf5,
	0xa9, 0xa3, 0x2c, 0x58, 0x81, 0x98, 0xef, 0x9b, 0x62, 0xa6, 0x1e, 0x3e, 0x74, 0xac, 0x8f, 0xea,
	0xeb, 0xc7, 0x8f, 0xf0, 0x14, 0xbd, 0x8a, 0xa9, 0x10, 0x91, 0x92, 0x02, 0x52, 0x5e, 0x43, 0x0a,
	0x69, 0x7c, 0x6b, 0xa8, 0x10, 0xbf, 0x48, 0x71, 0x6f, 0x8e, 0x62, 0xbf, 0x82, 0xff, 0x42, 0xdb,
	0x7e, 0x17, 0x1e, 0xb7, 0x6b, 0x0d, 0x12, 0x0e, 0xda, 0xee, 0xc7, 0xd3, 0x6d, 0x7b, 0xc3, 0x9f,
	0x7e, 0x8c, 0x4f, 0xd0, 0xcb, 0x1b, 0x55, 0x70, 0x61, 0x93, 0xd6, 0x21, 0x69, 0xbf, 0x29, 0xe9,
	0xc4, 0xea, 0xbd, 0x73, 0xef, 0xc6, 0x4d, 0xed, 0x5e, 0x9b, 0x1b, 0x95, 0xe7, 0xf6, 0xaa, 0x32,
	0x19, 0xd5, 0x70, 0x98, 0x36, 0xdb, 0xed, 0xf5, 0x59, 0x05, 0x9e, 0x59, 0xee, 0x7e, 0xaf, 0xcd,
	0xc3, 0x65, 0x9c, 0xa0, 0xbe, 0x8f, 0xc8, 0xb5, 0x62, 0x65, 0x0c, 0x21, 0x18, 0x42, 0xde, 0xb7,
	0x0d, 0x99, 0x38, 0xd2, 0xc7, 0xf8, 0xb2, 0xfd, 0x03, 0xdb, 0x95, 0x24, 0xa3, 0x29, 0x74, 0x65,
	0xab, 0x5d, 0x57, 0x8e, 0xad, 0xfe, 0xbe, 0x2b, 0x89, 0x9b, 0x5a, 0xa7, 0x34, 0xa3, 0x89, 0xbd,
	0x7a, 0x82, 0xcf, 0xdb, 0x39, 0x9d, 0x5a, 0xfd, 0xbd, 0x53, 0xea, 0xa6, 0xe3, 0x15, 0xd4, 0xa3,
	0x2c, 0x62, 0xb4, 0xa0, 0xe3, 0xff, 0x3a, 0xe8, 0x9b, 0x58, 0x65, 0xa4, 0xf1, 0xf3, 0x3e, 0xee,
	0x85, 0x6c, 0x62, 0x6f, 0x91, 0x49, 0xe7, 0xd7, 0x9f, 0x2a, 0x75, 0xa2, 0x04, 0x95, 0x09, 0x51,
	0x3a, 0x19, 0x24, 0x5c, 0xc2, 0x1d, 0x53, 0x7f, 0x61, 0xf3, 0xd4, 0x3c, 0xf3, 0xbf, 0xe2, 0xc0,
	0x8f, 0xfe, 0xee, 0xbe, 0x38, 0x0e, 0xc3, 0x7f, 0xba, 0x6f, 0xdd, 0x0d, 0x49, 0x42, 0x66, 0x88,
	0xbf, 0x2c, 0xc9, 0xe5, 0xd0, 0xbe, 0x54, 0x4e, 0xf9, 0x6f, 0xad, 0xb9, 0x0a, 0x99, 0xb9, 0xf2,
	0x9a, 0xab, 0xcb, 0xe1, 0x95, 0xd7, 0x5c, 0x2f, 0x43, 0x11, 0xef, 0x3e, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0xb0, 0x28, 0xdd, 0xdb, 0x08, 0x00, 0x00,
}
