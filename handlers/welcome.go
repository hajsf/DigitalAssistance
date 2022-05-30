package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"DigitalAssistance/Enum"
	"DigitalAssistance/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func WelcomeMessage(sender, name string) {
	var content strings.Builder
	content.WriteString(fmt.Sprintf("مرحبا *%v* \n", name))
	content.WriteString("معك المساعد الرقمي لمدير المشتريات في شركة قطوف و حلا")
	msg := &waProto.Message{ButtonsMessage: &waProto.ButtonsMessage{
		HeaderType:  waProto.ButtonsMessage_EMPTY.Enum(),
		ContentText: proto.String(content.String()),
		FooterText:  proto.String("من معي؟                          "),
		Buttons: []*waProto.Button{
			{
				ButtonId: proto.String(strconv.Itoa(Enum.Vendor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("مورد"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			},
			{
				ButtonId: proto.String(strconv.Itoa(Enum.Supervisor)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("مشرف فرع"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			},
			/*	{
				ButtonId: proto.String(strconv.Itoa(Enum.Franchisee)),
				ButtonText: &waProto.ButtonText{
					DisplayText: proto.String("حامل إمتياز"),
				},
				Type: waProto.Button_RESPONSE.Enum(),
			}, */
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
