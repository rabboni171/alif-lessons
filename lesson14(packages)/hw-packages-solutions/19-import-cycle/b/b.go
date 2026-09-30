package b

import "importcycle/common"

func Hello() string {
	return "b: " + common.Greeting
}
