package xhash

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Map2model map 转模型，取得 hash 数据时用
func Map2model(origin map[string]string, target interface{}) error {

	targetValue := reflect.ValueOf(target).Elem()

	// 循环处理每一个字段
	for i := 0; i < targetValue.NumField(); i++ {
		// 获取 tag
		field := targetValue.Type().Field(i)
		tag := ParseTag(field)
		// 忽略直接跳过
		if tag.IsIgnore {
			continue
		}
		// map 中不包含直接跳过
		originVal, has := origin[tag.Name]
		if !has {
			continue
		}
		err := setValue(targetValue, field, originVal)
		if err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------
// 根据字段类型，填充值
// ----------------------------------------
func setValue(targetValue reflect.Value, field reflect.StructField, originVal string) error {
	fieldValue := targetValue.FieldByName(field.Name)
	// 指针有专门的处理
	if field.Type.Kind() == reflect.Ptr {
		return setPtrValue(targetValue, field, originVal)
	}

	switch field.Type.Kind() {
	// 处理所有 Int 类型
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return setIntValue(fieldValue, originVal)
	// 处理所有 Uint 类型
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		return setUintValue(fieldValue, originVal)
	// 处理字符串类型
	case reflect.String:
		return setStrValue(fieldValue, originVal)
	// 处理布尔类型
	case reflect.Bool:
		return setBoolValue(fieldValue, originVal)
	// 处理浮点类型
	case reflect.Float64, reflect.Float32:
		return setFloatValue(fieldValue, originVal)
	// 处理切片类型
	case reflect.Slice:
		return setSliceValue(fieldValue, originVal)
	// 处理结构体类型
	case reflect.Struct, reflect.Map:
		if field.Type.String() == "time.Time" {
			return setTimeValue(fieldValue, originVal)
		}
		return setStructValue(fieldValue, originVal)
	// 处理 interface 类型
	case reflect.Interface:
		fieldValue.Set(reflect.ValueOf(originVal))
		return nil
	default:
		errMsg := fmt.Sprintf("unsupported type name=%s type=%s", field.Name, field.Type)
		return errors.New(errMsg)
	}
}

func setIntValue(fieldValue reflect.Value, originVal string) error {
	intVal, err := strconv.ParseInt(originVal, 10, 64)
	if err != nil {
		return err
	}
	fieldValue.SetInt(intVal)
	return nil
}

func setUintValue(fieldValue reflect.Value, originVal string) error {
	uintVal, err := strconv.ParseUint(originVal, 10, 64)
	if err != nil {
		return err
	}
	fieldValue.SetUint(uintVal)
	return nil
}

func setStrValue(fieldValue reflect.Value, originVal string) error {
	if fieldValue.Type().Kind() == reflect.Ptr {
		fieldValue.Set(reflect.ValueOf(&originVal))
	} else {
		fieldValue.SetString(originVal)
	}
	return nil
}

func setBoolValue(fieldValue reflect.Value, originVal string) error {
	boolVal, err := strconv.ParseBool(originVal)
	if err != nil {
		return err
	}
	if fieldValue.Type().Kind() == reflect.Ptr {
		fieldValue.Set(reflect.ValueOf(&boolVal))
	} else {
		fieldValue.SetBool(boolVal)
	}
	return nil
}

func setFloatValue(fieldValue reflect.Value, originVal string) error {
	floatVal, err := strconv.ParseFloat(originVal, 64)
	if err != nil {
		return err
	}
	fieldValue.SetFloat(floatVal)
	return nil
}

func setSliceValue(fieldValue reflect.Value, originVal string) error {
	sliceType := reflect.SliceOf(fieldValue.Type().Elem())
	slice := reflect.New(sliceType)
	bytesVal := bytes.NewBufferString(originVal).Bytes()
	err := json.Unmarshal(bytesVal, slice.Interface())
	if err != nil {
		return err
	}
	fieldValue.Set(reflect.ValueOf(slice.Interface()).Elem())
	return nil
}

func setTimeValue(fieldValue reflect.Value, originVal string) error {
	timeTime, err := time.ParseInLocation("2006-01-02 15:04:05", originVal, time.Local)
	if err != nil {
		return err
	}
	fieldValue.Set(reflect.ValueOf(timeTime))
	return nil
}

func setStructValue(fieldValue reflect.Value, originVal string) error {
	obj := reflect.New(fieldValue.Type())
	bytesVal := bytes.NewBufferString(originVal).Bytes()
	err := json.Unmarshal(bytesVal, obj.Interface())
	if err != nil {
		return err
	}
	fieldValue.Set(obj.Elem())
	return nil
}

// ----------------------------------------
// 根据字段类型，填充值，这个专门负责指针类型
// ----------------------------------------
func setPtrValue(targetValue reflect.Value, field reflect.StructField, originVal string) error {
	fieldValue := targetValue.FieldByName(field.Name)
	switch fieldValue.Type().Elem().Kind() {
	// 处理所有 Int 指针类型
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return setPtrIntValue(fieldValue, originVal)
	// 处理所有 Uint 指针类型
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		return setPtrUintValue(fieldValue, originVal)
	// 处理字符串指针类型
	case reflect.String:
		return setStrValue(fieldValue, originVal)
	// 处理布尔指针类型
	case reflect.Bool:
		return setBoolValue(fieldValue, originVal)
	// 处理布尔指针类型
	case reflect.Float64, reflect.Float32:
		return setPtrFloatValue(fieldValue, originVal)
	// 处理自定义结构体指针
	default:
		if fieldValue.Type().Elem().String() == "time.Time" {
			return setPtrTimeValue(fieldValue, originVal)
		}
		return setPtrStructValue(fieldValue, originVal)
	}
}

func setPtrIntValue(fieldValue reflect.Value, originVal string) error {
	intVal, err := strconv.ParseInt(originVal, 10, 64)
	if err != nil {
		return err
	}
	switch fieldValue.Type().Elem().Kind() {
	case reflect.Int64:
		fieldValue.Set(reflect.ValueOf(&intVal))
	case reflect.Int32:
		tmpVal := int32(intVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	case reflect.Int16:
		tmpVal := int16(intVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	case reflect.Int8:
		tmpVal := int8(intVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	case reflect.Int:
		tmpVal := int(intVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	}
	return nil
}

func setPtrUintValue(fieldValue reflect.Value, originVal string) error {
	uintVal, err := strconv.ParseUint(originVal, 10, 64)
	if err != nil {
		return err
	}
	switch fieldValue.Type().Elem().Kind() {
	case reflect.Uint64:
		fieldValue.Set(reflect.ValueOf(&uintVal))
	case reflect.Uint32:
		tmpVal := uint32(uintVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	case reflect.Uint16:
		tmpVal := uint16(uintVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	case reflect.Uint8:
		tmpVal := uint8(uintVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	case reflect.Uint:
		tmpVal := uint(uintVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	}
	return nil
}

func setPtrFloatValue(fieldValue reflect.Value, originVal string) error {
	floatVal, err := strconv.ParseFloat(originVal, 64)
	if err != nil {
		return err
	}
	switch fieldValue.Type().Elem().Kind() {
	case reflect.Float64:
		fieldValue.Set(reflect.ValueOf(&floatVal))
	case reflect.Float32:
		tmpVal := float32(floatVal)
		fieldValue.Set(reflect.ValueOf(&tmpVal))
	}
	return nil
}

func setPtrTimeValue(fieldValue reflect.Value, originVal string) error {
	timeTime, err := time.ParseInLocation("2006-01-02 15:04:05", originVal, time.Local)
	if err != nil {
		return err
	}
	fieldValue.Set(reflect.ValueOf(&timeTime))
	return nil
}

func setPtrStructValue(fieldValue reflect.Value, originVal string) error {
	obj := reflect.New(fieldValue.Type().Elem())
	bytesVal := bytes.NewBufferString(originVal).Bytes()
	err := json.Unmarshal(bytesVal, obj.Interface())
	if err != nil {
		return err
	}
	fieldValue.Set(obj)
	return nil
}
