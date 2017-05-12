package gopp

import (
	"reflect"
)

type Pair struct {
	Key interface{}
	Val interface{}
}

func Domap(ins interface{}, f func(interface{}) interface{}) (outs []interface{}) {
	outs = make([]interface{}, 0)

	tmpty := reflect.TypeOf(ins)
	// the same as DomapTypeField
	if tmpty.Kind() == reflect.Ptr && tmpty.String() == "*reflect.rtype" {
		insty := ins.(reflect.Type)

		for idx := 0; idx < insty.NumField(); idx++ {
			field := insty.Field(idx)
			out := f(field)
			outs = append(outs, out)
		}
	} else if tmpty.Kind() == reflect.Slice || tmpty.Kind() == reflect.Array {
		tmpv := reflect.ValueOf(ins)
		for i := 0; i < tmpv.Len(); i++ {
			out := f(tmpv.Index(i).Interface())
			outs = append(outs, out)
		}
	} else if tmpty.Kind() == reflect.Map {
		tmpv := reflect.ValueOf(ins)
		for _, vk := range tmpv.MapKeys() {
			out := f(tmpv.MapIndex(vk).Interface())
			outs = append(outs, &Pair{vk.Interface(), out})
		}
	} else {
		insRanger := ins.([]interface{})
		for _, in := range insRanger {
			out := f(in)
			outs = append(outs, out)
		}
	}

	return
}

func DomapTypeField(ty reflect.Type, f func(reflect.StructField) interface{}) (outs []interface{}) {
	outs = make([]interface{}, 0)

	for idx := 0; idx < ty.NumField(); idx++ {
		field := ty.Field(idx)
		out := f(field)
		outs = append(outs, out)
	}

	return
}

func Doreduce(ins interface{}, v interface{},
	f func(v, i interface{}) interface{}) interface{} {
	tmpty := reflect.TypeOf(ins)
	// the same as DomapTypeField
	if tmpty.Kind() == reflect.Ptr && tmpty.String() == "*reflect.rtype" {
		insty := ins.(reflect.Type)

		for idx := 0; idx < insty.NumField(); idx++ {
			field := insty.Field(idx)
			v = f(v, field)
		}
	} else if tmpty.Kind() == reflect.Slice || tmpty.Kind() == reflect.Array {
		tmpv := reflect.ValueOf(ins)
		for i := 0; i < tmpv.Len(); i++ {
			v = f(v, tmpv.Index(i).Interface())
		}
	} else if tmpty.Kind() == reflect.Map {
		tmpv := reflect.ValueOf(ins)
		for _, vk := range tmpv.MapKeys() {
			v = f(v, tmpv.MapIndex(vk).Interface())
		}
	} else {
		insRanger := ins.([]interface{})
		for _, in := range insRanger {
			v = f(v, in)
		}
	}

	return v
}

// interface vector to strings
func IV2Strings(items []interface{}) []string {
	if items == nil {
		return nil
	}

	rets := make([]string, 0)
	for idx := 0; idx < len(items); idx++ {
		rets = append(rets, items[idx].(string))
	}
	return rets
}
