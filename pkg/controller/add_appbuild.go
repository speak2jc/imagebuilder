package controller

import (
	"github.com/speak2jc/imagebuilder/pkg/controller/appbuild"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, appbuild.Add)
}
