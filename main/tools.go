package main

import (
	"errors"
	"fmt"
	"reflect"
)

//功能：移除切片中的重复元素，剩下元素保持原有顺序
//限制：
//1、切片中元素的类型必须可哈希（可比较），因此不能为：
//字典、切片、函数，或是含有这三种类型的结构体或数组
//2、必须传入切片的地址，而不是切片本身
//不遵循限制将会导致去重无效且返回错误
func RemoveDuplicate(in interface{}) error {
	var pseudo struct{}
	set := make(map[interface{}]struct{})
	pIn := reflect.ValueOf(in)
	vIn := reflect.Indirect(pIn)
	if !vIn.CanSet() {
		s := fmt.Sprintf("value cannot be set. Please pass the address of the slice.")
		return errors.New(s)
	}
	if vIn.Kind() == reflect.Slice {
		if vIn.Len() == 0 {
			return nil
		}
		//fmt.Println(pIn.Type().Elem())
		newV := reflect.MakeSlice(pIn.Type().Elem(), 0, vIn.Len())
		tOfElem := pIn.Elem().Type().Elem()
		if !tOfElem.Comparable() {
			s := fmt.Sprintf("the element's type is %s, which is not comparable", tOfElem.String())
			return errors.New(s)
		}
		for i := 0; i < vIn.Len(); i++ {
			//if already met
			if _, ok := set[vIn.Index(i).Interface()]; ok {
				continue
			} else {
				set[vIn.Index(i).Interface()] = pseudo
				newV = reflect.Append(newV, vIn.Index(i))
			}
		}
		vIn.Set(newV)
	} else {
		return errors.New("not a slice")
	}
	return nil
}

//判断value是否在array/slice里
//只允许array或是slice，否则会报错
func InSlice(value interface{}, container interface{}) (bool, error) {
	s := reflect.ValueOf(container)
	switch reflect.TypeOf(container).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true, nil
			}
		}
	default:
		//	如果不是slice或者array类型
		err := errors.New("you can only pass slice, array")
		return false, err
	}
	return false, nil
}
