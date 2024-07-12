package kry

import (
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"github.com/deta/deta-go/service/drive"
)

type DetaStruct struct {
	deta_ref    *deta.Deta
	deta_loaded bool
}

func (ref *DetaStruct) getDetaKey() string {
	deta_key := os.Getenv("DETA_PROJECT_KEY")
	return deta_key
}

func (ref *DetaStruct) init() {
	ref.deta_ref, _ = deta.New(deta.WithProjectKey(ref.getDetaKey()))
	ref.deta_loaded = true
}
func (ref *DetaStruct) Base(base_name string) *base.Base {
	if !ref.deta_loaded {
		ref.init()
	}
	base_ref, _ := base.New(ref.deta_ref, base_name)
	return base_ref
}
func (ref *DetaStruct) Drive(drive_name string) *drive.Drive {
	if !ref.deta_loaded {
		ref.init()
	}
	drive_ref, _ := drive.New(ref.deta_ref, drive_name)
	return drive_ref
}

var Deta DetaStruct

type KeyValuePairInt struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

var PROD_ENV = os.Getenv("APP_ENV") == "production"
var DEV_ENV = !PROD_ENV

type ErrorInfoType struct {
	Error string `json:"error"`
}
