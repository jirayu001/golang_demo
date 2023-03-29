package report

import (
	"fmt"
	"reflect"
	"strings"
)

//Text : generate report string
func Text(v interface{}) string {
	//Type
	//Value
	val := reflect.ValueOf(v)
	var sb strings.Builder
	text(&sb, &val)
	return sb.String()
}

func text(sb *strings.Builder, v *reflect.Value) {
	switch v.Kind() {
	case reflect.Struct:
		structName := v.Type().Name()
		if len(structName) == 0 {
			structName = "anonymous"

		}
		//sb.WriteString("i am struct")
		sb.WriteString(structName + "(")
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i)
			if i > 0 {
				sb.WriteString(", ")
			}
			//sb.WriteString(v.Type().Field(i).Name + " : ")
			displayName := v.Type().Field(i).Name
			tagName, ok := v.Type().Field(i).Tag.Lookup("report")
			if ok {
				tags := strings.Split(tagName, ",")
				if len(tags) == 2 {
					displayName = tags[0]
					if tags[1] == "uppercase" {
						tmp := v.Field(i).String()
						upper := strings.ToUpper(tmp)
						val = reflect.ValueOf(upper)
					}
				}
				displayName = tagName
			}
			sb.WriteString(displayName + " : ")

			text(sb, &val)
		}
		sb.WriteString(")")
	case reflect.String:
		sb.WriteString(v.String())

	case reflect.Int:
		sb.WriteString(fmt.Sprintf("%d", v.Int()))
	}
}
