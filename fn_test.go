package gopp

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Test_MapArrayInt_1(t *testing.T) {
	isEven := func(v int) bool { return v%2 == 0 }
	nums := []int{1, 2, 3, 4, 5, 6}
	mapped := MapArrayInt(isEven, nums)
	fmt.Println(mapped)

	// test 2
	isEven2 := Must(NewFunc(isEven))
	nums2 := []interface{}{}
	for n := range nums {
		nums2 = append(nums2, n)
	}

	mapped2 := MapArrayAny(isEven2, nums2)
	fmt.Println(mapped2)

	// test 3
	nums3 := make(map[interface{}]interface{}, 0)
	for n := range nums {
		nums3[n] = n
		nums3["abc"] = n // 这个也支持，果然interface{}强大
	}

	nums4 := ArrayToAnyArray(nums)
	fmt.Println(nums4)

	nums5 := ArrayToAnyMap(nums)
	mapped3 := MapMapAny(isEven2, nums5)
	fmt.Println(mapped3)

	nums6 := ToAny(nums)
	mapped6 := MapAny(isEven2, nums6)
	fmt.Println(mapped6)

	real_type := func(v interface{}) {
		tv := reflect.TypeOf(v)
		vv := reflect.ValueOf(v)
		fmt.Println(tv.Kind().String(), vv.Kind().String(), tv.Elem().Kind().String())
		//fmt.Println(vv.MapKeys())
		// fmt.Println(vv.Bytes())
	}

	// v1 := make(map[int]int, 0)
	v2 := make([]int, 0)
	v1any := ToAny(v2)
	real_type(v1any)

	shifta := func(v int) int {
		return v + 1
	}
	f_shifta := Must(NewFunc(shifta))
	s1 := "abcdefg"
	mapped7 := MapAny(f_shifta, s1)
	fmt.Println(mapped7)
}

func Test_Maybe_1(t *testing.T) {
	s := "hello,abc"
	ns := MaybeFrom(s).Map(Must(NewFunc(strings.ToUpper))).
		Map(Must(NewFunc(strings.ToLower))).
		Map(Must(NewFunc(strings.Title)))
	fmt.Println(ns.Value)
}

func Test_Maybe_2(t *testing.T) {
	s := "hello,abc"
	ns := MaybeFrom(s).Do(strings.ToUpper,
		strings.ToLower, strings.Title)

	fmt.Println(ns.Value)
}

func Test_Many_1(t *testing.T) {
	m := ManyFrom("hello world 123", "good byte 456", "foo bar 789")
	fmt.Printf("%#v, %v\n", m, m.Count())
	toUpper := Must(NewFunc(strings.ToUpper))
	fields := Must(NewFunc(strings.Fields))
	res := m.Map(toUpper).Map(fields)
	fmt.Println("res111", res, res.Count())

	fres := res.Flat()
	fmt.Println("fres222", fres, len(fres))
}

func Test_Compose_1(t *testing.T) {
	f1 := Must(Compose(strings.ToUpper, strings.ToLower))
	Assert(f1, "")
	toUpper := MustFunc(strings.ToUpper)
	fields := MustFunc(strings.Fields)
	f2 := Must(ComposeFunc(toUpper, fields))
	Assert(f2, "")
}
