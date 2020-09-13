package maptostruct

import (
	"reflect"
	"testing"
	"time"
)

type StructTest struct {
	Foo string `json:"foo"`
}

type MapToStructTest struct {
	String string `json:"string"`
	Int int `json:"int"`
	Float float64 `json:"float"`
	Array []int `json:"array"`
	Bool bool `json:"bool"`
	TimeT time.Time `json:"time_time"`
	TimeD time.Duration `json:"time_duration"`
	Struct StructTest `json:"struct"`
	Map map[string]interface{} `json:"map"`
}


func TestMapToStruct(t *testing.T) {

	data := make(map[string]interface{})
	mpTest := make(map[string]interface{})
	mpTest["foo"] = "bar"

	data["string"] = "foo"
	data["int"] = 1
	data["float"] = 0.300004
	data["array"] = []int{1, 2, 3}
	data["bool"] = true
	data["time_time"] = time.Now()
	data["time_duration"] = time.Since(time.Unix(1600000000, 0))
	data["struct"] = StructTest{ Foo: "Bar" }
	data["map"] = mpTest

	var out MapToStructTest

	if err := MapToStruct(data, &out); err != nil {
		t.Error(err)
		t.Fail()
	}

	if out.String != data["string"] {
		t.Error("fail on string")
		t.Fail()
	}

	if out.Int != data["int"] {
		t.Error("fail on int")
		t.Fail()
	}

	if out.Float != data["float"] {
		t.Error("fail on float")
		t.Fail()
	}

	if !reflect.DeepEqual(out.Array, data["array"]) {
		t.Error("fail on array")
		t.Fail()
	}

	if out.Bool != data["bool"] {
		t.Error("fail on bool")
		t.Fail()
	}

	if !out.TimeT.Equal(data["time_time"].(time.Time)) {
		t.Error("fail on time.Time")
		t.Fail()
	}

	if out.TimeD != data["time_duration"] {
		t.Error("fail on time.Duration")
		t.Fail()
	}

	if out.Struct != data["struct"] {
		t.Error("fail on struct")
		t.Fail()
	}

	if !reflect.DeepEqual(out.Map, data["map"]) {
		t.Error("fail on map")
		t.Fail()
	}
}
