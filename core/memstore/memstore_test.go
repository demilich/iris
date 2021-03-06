package memstore

import (
	"testing"
)

type myTestObject struct {
	name string
}

func TestMuttable(t *testing.T) {
	var p Store

	// slice
	p.Set("slice", []myTestObject{{"value 1"}, {"value 2"}})
	v := p.Get("slice").([]myTestObject)
	v[0].name = "modified"

	vv := p.Get("slice").([]myTestObject)
	if vv[0].name != "modified" {
		t.Fatalf("expected slice to be muttable but caller was not able to change its value")
	}

	// map

	p.Set("map", map[string]myTestObject{"key 1": myTestObject{"value 1"}, "key 2": myTestObject{"value 2"}})
	vMap := p.Get("map").(map[string]myTestObject)
	vMap["key 1"] = myTestObject{"modified"}

	vvMap := p.Get("map").(map[string]myTestObject)
	if vvMap["key 1"].name != "modified" {
		t.Fatalf("expected map to be muttable but caller was not able to change its value")
	}

	// object pointer of a value, it can change like maps or slices and arrays.
	p.Set("objp", &myTestObject{"value"})
	// we expect pointer here, as we set it.
	vObjP := p.Get("objp").(*myTestObject)

	vObjP.name = "modified"

	vvObjP := p.Get("objp").(*myTestObject)
	if vvObjP.name != "modified" {
		t.Fatalf("expected objp to be muttable but caller was able to change its value")
	}
}

func TestImmutable(t *testing.T) {
	var p Store

	// slice
	p.SetImmutable("slice", []myTestObject{{"value 1"}, {"value 2"}})
	v := p.Get("slice").([]myTestObject)
	v[0].name = "modified"

	vv := p.Get("slice").([]myTestObject)
	if vv[0].name == "modified" {
		t.Fatalf("expected slice to be immutable but caller was able to change its value")
	}

	// map
	p.SetImmutable("map", map[string]myTestObject{"key 1": myTestObject{"value 1"}, "key 2": myTestObject{"value 2"}})
	vMap := p.Get("map").(map[string]myTestObject)
	vMap["key 1"] = myTestObject{"modified"}

	vvMap := p.Get("map").(map[string]myTestObject)
	if vvMap["key 1"].name == "modified" {
		t.Fatalf("expected map to be immutable but caller was able to change its value")
	}

	// object value, it's immutable at all cases.
	p.SetImmutable("obj", myTestObject{"value"})
	vObj := p.Get("obj").(myTestObject)
	vObj.name = "modified"

	vvObj := p.Get("obj").(myTestObject)
	if vvObj.name == "modified" {
		t.Fatalf("expected obj to be immutable but caller was able to change its value")
	}

	// object pointer of a value, it's immutable at all cases.
	p.SetImmutable("objp", &myTestObject{"value"})
	// we expect no pointer here if SetImmutable.
	// so it can't be changed by-design
	vObjP := p.Get("objp").(myTestObject)

	vObjP.name = "modified"

	vvObjP := p.Get("objp").(myTestObject)
	if vvObjP.name == "modified" {
		t.Fatalf("expected objp to be immutable but caller was able to change its value")
	}
}
