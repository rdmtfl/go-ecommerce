// syntax = "proto3";

// package pb;

// option go_package = "./pb";

// message Order {
//     string order_id = 1;
//     string customer_id = 2;
// }

// message OrderRequest {
//     Order orderEntry = 1;
// }

// message OrderResponse {
//     string result = 1;
// }

// service OrderService {
//     rpc PostOrder(OrderRequest) returns (OrderResponse);
// }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: order.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	OrderId         string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                       // "order-abcde"
	CustomerId      string                 `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`              // "CUST-7890"
	OrderDate       string                 `protobuf:"bytes,3,opt,name=order_date,json=orderDate,proto3" json:"order_date,omitempty"`                 // ISO 8601 format
	Status          string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`                                        // e.g. "processing"
	Items           []*Order_Item          `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`                                          // list of items
	Subtotal        float64                `protobuf:"fixed64,6,opt,name=subtotal,proto3" json:"subtotal,omitempty"`                                  // 72.48
	Tax             float64                `protobuf:"fixed64,7,opt,name=tax,proto3" json:"tax,omitempty"`                                            // 5.80
	ShippingCost    float64                `protobuf:"fixed64,8,opt,name=shipping_cost,json=shippingCost,proto3" json:"shipping_cost,omitempty"`      // 5.99
	Total           float64                `protobuf:"fixed64,9,opt,name=total,proto3" json:"total,omitempty"`                                        // 84.27
	ShippingMethod  string                 `protobuf:"bytes,10,opt,name=shipping_method,json=shippingMethod,proto3" json:"shipping_method,omitempty"` // "standard"
	PaymentMethod   string                 `protobuf:"bytes,11,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`    // "credit_card"
	BillingAddress  *Order_Address         `protobuf:"bytes,12,opt,name=billing_address,json=billingAddress,proto3" json:"billing_address,omitempty"`
	ShippingAddress *Order_Address         `protobuf:"bytes,13,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	Notes           string                 `protobuf:"bytes,14,opt,name=notes,proto3" json:"notes,omitempty"`         // Delivery instructions
	Discounts       []*Order_Discount      `protobuf:"bytes,15,rep,name=discounts,proto3" json:"discounts,omitempty"` // list of discounts
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Order) GetOrderDate() string {
	if x != nil {
		return x.OrderDate
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetItems() []*Order_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetSubtotal() float64 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

func (x *Order) GetTax() float64 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Order) GetShippingCost() float64 {
	if x != nil {
		return x.ShippingCost
	}
	return 0
}

func (x *Order) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Order) GetShippingMethod() string {
	if x != nil {
		return x.ShippingMethod
	}
	return ""
}

func (x *Order) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *Order) GetBillingAddress() *Order_Address {
	if x != nil {
		return x.BillingAddress
	}
	return nil
}

func (x *Order) GetShippingAddress() *Order_Address {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

func (x *Order) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Order) GetDiscounts() []*Order_Discount {
	if x != nil {
		return x.Discounts
	}
	return nil
}

type OrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderEntry    *Order                 `protobuf:"bytes,1,opt,name=orderEntry,proto3" json:"orderEntry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	mi := &file_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderRequest) GetOrderEntry() *Order {
	if x != nil {
		return x.OrderEntry
	}
	return nil
}

type OrderResponse struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Error         bool                        `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`    // Always true for error responses
	Message       string                      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // General error message ("missing fields")
	Fields        []string                    `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`   // Detailed error messages
	Data          *OrderResponse_ResponseData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	mi := &file_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *OrderResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *OrderResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OrderResponse) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *OrderResponse) GetData() *OrderResponse_ResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

type Order_Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`   // "PROD-1001"
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`                     // 2
	UnitPrice     float64                `protobuf:"fixed64,3,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"` // 29.99
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                              // "Wireless Headphones"
	Description   string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`                // Product details
	Sku           string                 `protobuf:"bytes,6,opt,name=sku,proto3" json:"sku,omitempty"`                                // "WH-2023-BLK"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order_Item) Reset() {
	*x = Order_Item{}
	mi := &file_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order_Item) ProtoMessage() {}

func (x *Order_Item) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order_Item.ProtoReflect.Descriptor instead.
func (*Order_Item) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Order_Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Order_Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Order_Item) GetUnitPrice() float64 {
	if x != nil {
		return x.UnitPrice
	}
	return 0
}

func (x *Order_Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Order_Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order_Item) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

type Order_Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Street        string                 `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`                           // "123 Main St"
	City          string                 `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`                               // "Austin"
	State         string                 `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`                             // "TX"
	PostalCode    string                 `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"` // "78701"
	Country       string                 `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`                         // "USA"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order_Address) Reset() {
	*x = Order_Address{}
	mi := &file_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order_Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order_Address) ProtoMessage() {}

func (x *Order_Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order_Address.ProtoReflect.Descriptor instead.
func (*Order_Address) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Order_Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Order_Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Order_Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Order_Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Order_Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Order_Discount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`                             // "SUMMER10"
	Amount        float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`                       // 10.0
	IsPercent     bool                   `protobuf:"varint,3,opt,name=is_percent,json=isPercent,proto3" json:"is_percent,omitempty"` // true
	Description   string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`               // "Summer Sale 10% Off"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order_Discount) Reset() {
	*x = Order_Discount{}
	mi := &file_order_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order_Discount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order_Discount) ProtoMessage() {}

func (x *Order_Discount) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order_Discount.ProtoReflect.Descriptor instead.
func (*Order_Discount) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Order_Discount) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Order_Discount) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order_Discount) GetIsPercent() bool {
	if x != nil {
		return x.IsPercent
	}
	return false
}

func (x *Order_Discount) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type OrderResponse_ResponseData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`                // generated order ID
	CreationDate  string                 `protobuf:"bytes,2,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"` // ISO 8601 timestamp
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderResponse_ResponseData) Reset() {
	*x = OrderResponse_ResponseData{}
	mi := &file_order_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderResponse_ResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse_ResponseData) ProtoMessage() {}

func (x *OrderResponse_ResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse_ResponseData.ProtoReflect.Descriptor instead.
func (*OrderResponse_ResponseData) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2, 0}
}

func (x *OrderResponse_ResponseData) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderResponse_ResponseData) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

const file_order_proto_rawDesc = "" +
	"\n" +
	"\vorder.proto\x12\x02pb\"\xc8\a\n" +
	"\x05Order\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x1f\n" +
	"\vcustomer_id\x18\x02 \x01(\tR\n" +
	"customerId\x12\x1d\n" +
	"\n" +
	"order_date\x18\x03 \x01(\tR\torderDate\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06status\x12$\n" +
	"\x05items\x18\x05 \x03(\v2\x0e.pb.Order.ItemR\x05items\x12\x1a\n" +
	"\bsubtotal\x18\x06 \x01(\x01R\bsubtotal\x12\x10\n" +
	"\x03tax\x18\a \x01(\x01R\x03tax\x12#\n" +
	"\rshipping_cost\x18\b \x01(\x01R\fshippingCost\x12\x14\n" +
	"\x05total\x18\t \x01(\x01R\x05total\x12'\n" +
	"\x0fshipping_method\x18\n" +
	" \x01(\tR\x0eshippingMethod\x12%\n" +
	"\x0epayment_method\x18\v \x01(\tR\rpaymentMethod\x12:\n" +
	"\x0fbilling_address\x18\f \x01(\v2\x11.pb.Order.AddressR\x0ebillingAddress\x12<\n" +
	"\x10shipping_address\x18\r \x01(\v2\x11.pb.Order.AddressR\x0fshippingAddress\x12\x14\n" +
	"\x05notes\x18\x0e \x01(\tR\x05notes\x120\n" +
	"\tdiscounts\x18\x0f \x03(\v2\x12.pb.Order.DiscountR\tdiscounts\x1a\xa8\x01\n" +
	"\x04Item\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x02 \x01(\x05R\bquantity\x12\x1d\n" +
	"\n" +
	"unit_price\x18\x03 \x01(\x01R\tunitPrice\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12\x10\n" +
	"\x03sku\x18\x06 \x01(\tR\x03sku\x1a\x86\x01\n" +
	"\aAddress\x12\x16\n" +
	"\x06street\x18\x01 \x01(\tR\x06street\x12\x12\n" +
	"\x04city\x18\x02 \x01(\tR\x04city\x12\x14\n" +
	"\x05state\x18\x03 \x01(\tR\x05state\x12\x1f\n" +
	"\vpostal_code\x18\x04 \x01(\tR\n" +
	"postalCode\x12\x18\n" +
	"\acountry\x18\x05 \x01(\tR\acountry\x1aw\n" +
	"\bDiscount\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\x01R\x06amount\x12\x1d\n" +
	"\n" +
	"is_percent\x18\x03 \x01(\bR\tisPercent\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\"9\n" +
	"\fOrderRequest\x12)\n" +
	"\n" +
	"orderEntry\x18\x01 \x01(\v2\t.pb.OrderR\n" +
	"orderEntry\"\xdb\x01\n" +
	"\rOrderResponse\x12\x14\n" +
	"\x05error\x18\x01 \x01(\bR\x05error\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x16\n" +
	"\x06fields\x18\x03 \x03(\tR\x06fields\x122\n" +
	"\x04data\x18\x04 \x01(\v2\x1e.pb.OrderResponse.ResponseDataR\x04data\x1aN\n" +
	"\fResponseData\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12#\n" +
	"\rcreation_date\x18\x02 \x01(\tR\fcreationDate2@\n" +
	"\fOrderService\x120\n" +
	"\tPostOrder\x12\x10.pb.OrderRequest\x1a\x11.pb.OrderResponseB\x06Z\x04./pbb\x06proto3"

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData []byte
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)))
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_order_proto_goTypes = []any{
	(*Order)(nil),                      // 0: pb.Order
	(*OrderRequest)(nil),               // 1: pb.OrderRequest
	(*OrderResponse)(nil),              // 2: pb.OrderResponse
	(*Order_Item)(nil),                 // 3: pb.Order.Item
	(*Order_Address)(nil),              // 4: pb.Order.Address
	(*Order_Discount)(nil),             // 5: pb.Order.Discount
	(*OrderResponse_ResponseData)(nil), // 6: pb.OrderResponse.ResponseData
}
var file_order_proto_depIdxs = []int32{
	3, // 0: pb.Order.items:type_name -> pb.Order.Item
	4, // 1: pb.Order.billing_address:type_name -> pb.Order.Address
	4, // 2: pb.Order.shipping_address:type_name -> pb.Order.Address
	5, // 3: pb.Order.discounts:type_name -> pb.Order.Discount
	0, // 4: pb.OrderRequest.orderEntry:type_name -> pb.Order
	6, // 5: pb.OrderResponse.data:type_name -> pb.OrderResponse.ResponseData
	1, // 6: pb.OrderService.PostOrder:input_type -> pb.OrderRequest
	2, // 7: pb.OrderService.PostOrder:output_type -> pb.OrderResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_proto_rawDesc), len(file_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
