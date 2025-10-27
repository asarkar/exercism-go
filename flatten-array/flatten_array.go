package flatten

// func Flatten(nested any) []any {
// 	result := make([]any, 0)
// 	// `reflect.ValueOf(x)` returns a `reflect.Value` that actually wraps
// 	// the runtime value of `x``.
// 	// With it, we can manipulate and inspect contents dynamically.
// 	// If we only used `reflect.TypeOf`, we'd know it's a slice, but we'd
// 	// have no way to extract each element — no `Len()` or `Index()` methods exist
// 	// on a `reflect.Type`.
// 	val := reflect.ValueOf(nested)

// 	switch val.Kind() { //nolint:exhaustive
// 	case reflect.Invalid:
// 		// discard
// 	case reflect.Slice, reflect.Array:
// 		for i := range val.Len() {
// 			// Get v's current value as an `interface{}`.
// 			item := val.Index(i).Interface()
// 			if item == nil {
// 				continue
// 			}
// 			result = append(result, Flatten(item)...)
// 		}
// 	default:
// 		result = append(result, nested)
// 	}

// 	return result
// }

func Flatten(nested any) []any {
	res := make([]any, 0)

	if nested == nil {
		return res
	}

	switch nested := nested.(type) {
	case []any:
		for _, e := range nested {
			if e == nil {
				continue
			}
			res = append(res, Flatten(e)...)
		}
	default:
		res = append(res, nested)
	}
	return res
}
