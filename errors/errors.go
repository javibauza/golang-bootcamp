package errors

import "fmt"

type InternalError string
type ThirdPartyError string
type OtherError string

func (e InternalError) Error() string {
	return "InternalError"
}
func (e ThirdPartyError) Error() string {
	return "ThirdPartyError"
}
func (e OtherError) Error() string {
	return "OtherError"
}

func Set(eType string) error {
	var ie InternalError
	var tpE ThirdPartyError
	var oE OtherError

	switch eType {
	case "Internal":
		return ie
	case "ThirdParty":
		return tpE
	case "Other":
		return oE
	default:
		return nil
	}
}

func Get(e error, eT string) bool {
	switch eType := e.(type) {
	case InternalError:
		if eT == "Internal" {
			return true
		}
	case ThirdPartyError:
		if eT == "ThirdParty" {
			return true
		}
	case OtherError:
		if eT == "Other" {
			return true
		}
	default:
		fmt.Printf("eType: %v", eType)
	}

	return false
}
