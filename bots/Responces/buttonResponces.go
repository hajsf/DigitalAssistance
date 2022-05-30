package responces

import (
	"DigitalAssistance/Enum"
	"DigitalAssistance/bots/branches"
	"DigitalAssistance/bots/franchisee"
	"DigitalAssistance/bots/vendor"
	"DigitalAssistance/global"
)

func ButtonResponses(id int, sender string) {
	switch id {

	case Enum.Vendor:
		go vendor.WhichVender(sender)
	case Enum.ExistingVendor:
		global.Users = append(global.Users, map[string]int{
			sender: Enum.ExistingVendor,
		})
		go vendor.CurrentVender(sender)
	//	franchisee.ContactSuperIntendent(sender)
	//	franchisee.SendContact(sender)
	//	Bye(sender)
	case Enum.NewVendor:
		global.Users = append(global.Users, map[string]int{
			sender: Enum.NewVendor,
		})
		go vendor.NewVender(sender)
	case Enum.Supervisor:
		global.Users = append(global.Users, map[string]int{
			sender: Enum.Supervisor,
		})
		branches.BranchIssue(sender)
		//	franchisee.ContactSuperIntendent(sender)
		//	franchisee.SendContact(sender)
		Bye(sender)
	case Enum.Franchisee:
		global.Users = append(global.Users, map[string]int{
			sender: Enum.Franchisee,
		})
		franchisee.ContactSuperIntendent(sender)
		franchisee.SendContact(sender)
		Bye(sender)
	case Enum.Yes: // yes another service is required
		var userType int
		for _, v := range global.Users {
			//fmt.Println(k, "is:", v[sender])
			for key, value := range v {
				if key == sender {
					userType = value
					break
				}
			}
		}

		switch userType {
		case Enum.ExistingVendor:
			vendor.CurrentVender(sender)
		case Enum.NewVendor:
			vendor.NewVender(sender)
		case Enum.Franchisee:
		case Enum.Supervisor:
		}
	case Enum.No:
		Bye(sender)

	}
}
