package dominate

// interface{} -> string
func Interface2String(inter interface{}) (string, bool) {
	if nil != inter {
		if value, success := inter.(string); success {
			return value, true
		}
	}
	return "", false
}

// interface{} -> int
func Interface2Number(inter interface{}) (int, bool) {
	if nil != inter {
		if value, success := inter.(int); success {
			return value, true
		}
	}
	return 0, false
}

// interface{} -> true
func Interface2Boolean(inter interface{}) (bool, bool) {
	if nil != inter {
		if value, success := inter.(bool); success {
			return value, true
		}
	}
	return false, false
}

// interface{} -> []interface{}
func Interface2Array(inter interface{}) ([]interface{}, bool) {
	if nil != inter {
		if value, success := inter.([]interface{}); success {
			return value, true
		}
	}
	return nil, false
}

// interface{} -> map[string]interface{}
func Interface2Map(inter interface{}) (map[string]interface{}, bool) {
	if nil != inter {
		if value, success := inter.(map[string]interface{}); success {
			return value, true
		}
	}
	return nil, false
}
