package main

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// 自定义 required 验证函数
func customRequired(fl validator.FieldLevel) bool {
	field := fl.Field()

	// 处理 nil 值
	if field.IsNil() {
		return false
	}

	// 处理不同类型的字段
	switch field.Kind() {
	case reflect.String:
		return field.Len() > 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() != 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint() != 0

	case reflect.Float32, reflect.Float64:
		return field.Float() != 0

	case reflect.Slice, reflect.Array, reflect.Map:
		return field.Len() > 0

	case reflect.Struct:
		// 结构体需要特殊处理，可以根据需求自定义逻辑
		return !field.IsZero()

	default:
		return true
	}
}

func main() {
	validate := validator.New()

	// 注册自定义 required 验证器
	validate.RegisterValidation("required", customRequired)

	// 定义测试结构体
	type User struct {
		Name string `json:"name" validate:"custom_required"`
	}

	// 测试数据
	user := User{
		Name: "", // 会失败
	}

	// 执行验证
	if err := validate.Struct(user); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("字段 %s 验证失败，标签: %s\n", err.Field(), err.Tag())
		}
	} else {
		fmt.Println("验证通过")
	}
}
