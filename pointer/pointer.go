package pointer

import "syscall/js"

func PInt(value int) *int {
	return &value
}

func PString(value string) *string {
	return &value
}

func PJSString(value string) *js.Value {
	aux := js.Global().Get("String").New(value)
	return &aux
}

func SetPJS(value *js.Value, val interface{}) *js.Value {

	*value = js.ValueOf(val)

	return value
}

func PBool(value bool) *bool {
	return &value
}

func PInterface(value interface{}) *interface{} {
	return &value
}

/*
func PLabelPosition(value JEasyUI.LabelPosition) *JEasyUI.LabelPosition {
	return &value
}

func PAlertIcon(value JEasyUI.AlertIcon) *JEasyUI.AlertIcon {
	return &value
}

func PMessagerShowType(value JEasyUI.MessagerShowType) *JEasyUI.MessagerShowType {
	return &value
}
*/
