package xhash

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"time"
)

func Model2map(origin interface{}) (map[string]interface{}, error) {

	originValue := reflect.ValueOf(origin).Elem()

	// 循环处理每一个字段
	result := make(map[string]interface{})
	for i := 0; i < originValue.NumField(); i++ {
		// 获取 tag
		field := originValue.Type().Field(i)
		tag := ParseTag(field)
		// 忽略直接跳过
		if tag.IsIgnore {
			continue
		}

		value, err := getValue(originValue, field)
		if err != nil {
			return nil, err
		}

		result[tag.Name] = value
	}
	return result, nil
}

// ----------------------------------------
// 根据字段类型，转换成可用类型
// ----------------------------------------
func getValue(originValue reflect.Value, field reflect.StructField) (interface{}, error) {
	fieldValue := originValue.FieldByName(field.Name)

	// 指针有专门的处理
	if field.Type.Kind() == reflect.Ptr {
		return getPtrValue(originValue, field)
	}

	switch field.Type.Kind() {
	// 处理所有 Int 类型
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return fieldValue.Int(), nil
	// 处理所有 Uint 类型
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		return fieldValue.Uint(), nil
	// 处理字符串类型
	case reflect.String:
		return fieldValue.Interface(), nil
	// 处理布尔类型
	case reflect.Bool:
		return fieldValue.Bool(), nil
	// 处理浮点类型
	case reflect.Float64, reflect.Float32:
		return fieldValue.Float(), nil
	// 处理切片类型
	case reflect.Slice:
		return getJsonValue(fieldValue)
	// 处理结构体类型
	case reflect.Struct, reflect.Map:
		if field.Type.String() == "time.Time" {
			return getTimeValue(fieldValue)
		} else {
			return getJsonValue(fieldValue)
		}
	// 处理 interface 类型
	case reflect.Interface:
		return fieldValue.Interface(), nil
	default:
		errMsg := fmt.Sprintf("unsupported type name=%s type=%s", field.Name, field.Type)
		return nil, errors.New(errMsg)
	}
}

func getTimeValue(fieldValue reflect.Value) (interface{}, error) {
	t := fieldValue.Interface().(time.Time)
	value := t.Local().Format("2006-01-02 15:04:05")
	return value, nil
}

func getJsonValue(fieldValue reflect.Value) (interface{}, error) {
	return json.Marshal(fieldValue.Interface())
}

// ----------------------------------------
// 根据字段类型，转换成可用类型, 专供指针
// ----------------------------------------
func getPtrValue(originValue reflect.Value, field reflect.StructField) (interface{}, error) {

	fieldValue := originValue.FieldByName(field.Name)

	if fieldValue.IsNil() {
		return nil, nil
	}

	switch fieldValue.Type().Elem().Kind() {
	// 处理所有 Int 指针类型
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
		return fieldValue.Elem().Interface(), nil
	// 处理所有 Uint 指针类型
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint:
		return fieldValue.Elem().Interface(), nil
	// 处理字符串指针, 布尔指针类型
	case reflect.String, reflect.Bool:
		return fieldValue.Elem().Interface(), nil
	// 处理布尔指针类型
	case reflect.Float64, reflect.Float32:
		return fieldValue.Elem().Interface(), nil
	// 处理自定义结构体指针
	default:
		if fieldValue.Type().Elem().String() == "time.Time" {
			return getPtrTimeValue(fieldValue)
		} else {
			return getPtrJsonValue(fieldValue)
		}
	}
}

func getPtrTimeValue(fieldValue reflect.Value) (interface{}, error) {
	t := fieldValue.Elem().Interface().(time.Time)
	value := t.Local().Format("2006-01-02 15:04:05")
	return value, nil
}

func getPtrJsonValue(fieldValue reflect.Value) (interface{}, error) {
	return json.Marshal(fieldValue.Elem().Interface())
}
