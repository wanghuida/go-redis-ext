package xhash

import (
	"bytes"
	"reflect"
	"strings"
	"unicode"
)

const (
	// XHashTag tag 的名称
	XHashTag    = "redis"

	// XHashTagSep tag 的分隔符
	XHashTagSep = ";"
)

type FieldTag struct {
	Name     string // 字段存储的名称
	IsIgnore bool   // 是否忽略该字段，不存储
}

// ParseTag 分析字段的 tag
func ParseTag(field reflect.StructField) *FieldTag {
	fieldTag := &FieldTag{
		Name:     Hump2underline(field.Name),
		IsIgnore: false,
	}

	tagStr := field.Tag.Get(XHashTag)
	tagGroup := strings.Split(tagStr, XHashTagSep)

	// split 默认会有一个，如果为空，直接按默认返回
	if len(tagGroup) == 1 && tagGroup[0] == "" {
		return fieldTag
	}

	// 忽略的情况
	if tagGroup[0] == "-" {
		fieldTag.IsIgnore = true
		return fieldTag
	}

	// 自定义命名
	fieldTag.Name = tagGroup[0]
	return fieldTag
}

// Hump2underline 将驼峰转为下划线
func Hump2underline(name string) string {
	buffer := bytes.NewBufferString("")
	for i, r := range name {
		if i != 0 && unicode.IsUpper(r) {
			buffer.WriteString("_")
		}
		buffer.WriteRune(unicode.ToLower(r))
	}

	return buffer.String()
}
