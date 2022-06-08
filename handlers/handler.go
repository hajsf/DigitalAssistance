package handlers

import (
	responces "DigitalAssistance/bots/Responces"
	"DigitalAssistance/bots/branches"
	"DigitalAssistance/global"
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/abadojack/whatlanggo"
	"go.mau.fi/whatsmeow/types/events"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func Handler(rawEvt interface{}) {
	/* Set language translation */

	// Create a new i18n bundle with default language.
	bundle := i18n.NewBundle(language.English)

	// Register a toml unmarshal function for i18n bundle.
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load translations from toml files for non-default languages.
	bundle.MustLoadMessageFile("./lang/active.ar.toml")
	bundle.MustLoadMessageFile("./lang/active.es.toml")

	var lang string

	switch evt := rawEvt.(type) {
	case *events.Message:
		sender := evt.Info.Chat.User
		pushName := evt.Info.PushName

		//	if sender == "966138117381" {
		// branches.SendCard(sender, "6287029390129")
		//	}
		re := regexp.MustCompile(`(?i)^(card|كرت|بطاقه|بطاقة)\s(?P<barcode>\w*)`)
		barcodeIndex := re.SubexpIndex("barcode")

		//	msgType := evt.Message.GetMessageContextInfo()
		//	fmt.Println(msgType)
		switch {
		case evt.Message.ExtendedTextMessage != nil:

		//	note3 := evt.Message.GetExtendedTextMessage()
		//	fmt.Println("hi 3", note3)
		case evt.Message.Conversation != nil:
			var msgReceived string
			received := evt.Message.GetConversation()
			// convert numbers in Arabic scrtip to numbers in latin script
			for _, e := range received {
				if e >= 48 && e <= 57 {
					//	fmt.Println("Number in english script number")
					msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
				} else if e >= 1632 && e <= 1641 {
					//	fmt.Println("It is Arabic script")
					msgReceived = fmt.Sprintf("%s%v", msgReceived, global.NormalizeNumber(e))
				} else {
					//	fmt.Println("Dose not looks to be a number")
					msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
				}
			}

			matches := re.FindStringSubmatch(msgReceived)

			//	if sender == "966138117381" {
			//		check.SendDisposable(sender)
			//	}
			if len(matches) > 0 {
				//	fmt.Println("searching for barcode:", matches[barcodeIndex])
				go branches.SendCard(sender, matches[barcodeIndex])
			} else if !evt.Info.IsGroup && !evt.Info.IsFromMe && (sender != "966556888145" && // && !evt.Info.IsFromMe
				sender != "966505148268" && sender != "966531041222" && sender != "966577942979" &&
				sender != "966506888972" && sender != "966557776097" && sender != "966505360700" && sender != "966555786616" &&
				sender != "966508884337" && sender != "966508899479" && sender != "966530052201" && sender != "966558936645" &&
				sender != "966502887935" && sender != "971563451686") {

				info := whatlanggo.Detect(evt.Message.GetConversation())
				fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script], " Confidence: ", info.Confidence)

				switch whatlanggo.Scripts[info.Script] {
				case "Arabic":
					go WelcomeMessage(sender, pushName)
					lang = "ar"
				case "Latin":
					go WelcomeMessageLatin(sender, pushName)
				}

				// Create a new localizer.
				localizer := i18n.NewLocalizer(bundle, lang)
				// Set title message.
				helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: &i18n.Message{
						ID:    "HelloPerson",     // set translation ID
						Other: "Hello {{.Name}}", // set default translation
					},
					TemplateData: map[string]string{
						"Name": pushName,
					},
					PluralCount: nil,
				})

				fmt.Println(helloPerson)
				//	go WelcomeMessage(sender, pushName)
			}

		case evt.Message.ImageMessage != nil,
			evt.Message.AudioMessage != nil,
			evt.Message.VideoMessage != nil,
			evt.Message.DocumentMessage != nil:
			if !evt.Info.IsGroup && !evt.Info.IsFromMe && (sender != "966556888145" && // !evt.Info.IsFromMe &&
				sender != "966505148268" && sender != "966531041222" && sender != "966502887935" && sender != "966551775959" &&
				sender != "966550669344") {
				WelcomeMessage(sender, pushName)
			}
		case evt.Message.ButtonsResponseMessage != nil:
			fmt.Println("Button responce pressed")
			ButtonResponse := evt.Message.GetButtonsResponseMessage()
			id, _ := strconv.Atoi(ButtonResponse.GetSelectedButtonId())
			responces.ButtonResponses(id, sender)

		case evt.Message.ListResponseMessage != nil:
			fmt.Println("List responce pressed")
			ListResponse := evt.Message.GetListResponseMessage()
			id, _ := strconv.Atoi(ListResponse.SingleSelectReply.GetSelectedRowId())
			fmt.Println(id, sender)
			responces.ListResponces(id, sender, pushName)
		case evt.Message.LocationMessage != nil:
			Location := evt.Message.GetLocationMessage()
			fmt.Println(Location.GetDegreesLatitude())
			fmt.Println(Location.GetDegreesLongitude())
			fmt.Println(Location.GetAddress())

			latitude := Location.GetDegreesLatitude()
			longitud := Location.GetDegreesLongitude()
			address := Location.GetAddress()

			global.Locations = append(global.Locations, global.Location{
				Sender:    sender,
				PushName:  pushName,
				Latitude:  Location.GetDegreesLatitude(),
				Longitude: Location.GetDegreesLongitude(),
				Address:   Location.GetAddress(),
			})

			tx, err := global.Db.Begin()
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := tx.Prepare("insert into locations (name, jid, longitude, latitude, address) values(?, ?, ?, ?, ?)")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(pushName, sender, longitud, latitude, address)
			if err != nil {
				log.Fatal(err)
			}

			tx.Commit()

			fmt.Println(global.Locations)

		case evt.Message.ContactMessage != nil:
			Contact := evt.Message.GetContactMessage()
			fmt.Println(Contact.GetDisplayName())
			fmt.Println(Contact.GetVcard())
		} // End of switch
	}
}
