package vendor

import (
	"strconv"

	"DigitalAssistance/Enum"
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func CurrentVender(sender string) {
	msg := &waProto.Message{
		ListMessage: &waProto.ListMessage{
			Description: proto.String("أهلا و سهلا بشريك النجاح"),
			ButtonText:  proto.String("يرجى تحديد الهدف من التواصل"),
			ListType:    waProto.ListMessage_SINGLE_SELECT.Enum(),
			Sections: []*waProto.Section{
				{
					Rows: []*waProto.Row{
						{
							RowId: proto.String(strconv.Itoa(Enum.Discount)),
							Title: proto.String("تسجيل حملة ترويجية و خصومات"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.AdDocumentation)),
							Title: proto.String("توثيق حملة دعائية و تسويقية"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.VendorComplain)),
							Title: proto.String("رفع ملاحظة أو شكوى"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.ItemRegistration)),
							Title: proto.String("تسجيل صنف لدى قطوف و حلا"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.PriceList)),
							Title: proto.String("تحميل و إرسال عرض أسعار"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.PriceChange)),
							Title: proto.String("طلب تعديل أسعار"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.Location)),
							Title: proto.String("طلب الموقع لإرسال عينات"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.VAT)),
							Title: proto.String("طلب الرقم الضريبي و السجل التجاري"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.CallArrangement)),
							Title: proto.String("ترتيب إتصال مرئي أو هاتفي"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.VisitArrangement)),
							Title: proto.String("جدولة زيارة و لقاء شخصي"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.Email)),
							Title: proto.String("البريد الإلكتروني لمدير القسم"),
						},
					},
				},
			},
		},
	}

	jid, ok := global.ParseJID(sender)
	if !ok {
		return
	}
	send, err := global.Cli.SendMessage(jid, "", msg) // jid = recipient

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}

}
