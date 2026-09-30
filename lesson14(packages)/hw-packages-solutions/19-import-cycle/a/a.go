package a

import "importcycle/common"

func Hello() string {
	return "a: " + common.Greeting
}
