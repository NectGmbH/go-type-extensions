// Package slices implements a couple of helpful array/slice/list operations on generic slices.
package slices

// ToSingleton returns the first element of sl.
// If the slice is empty, it returns the zero value for the slices base type.
func ToSingleton[TElem any](sl []TElem) TElem {
	for _, elem := range sl {
		return elem
	}

	return *new(TElem)
}

// First is an alias for ToSingleton,
// it returns the first element of sl.
func First[TElem any](sl []TElem) TElem {
	return ToSingleton[TElem](sl)
}

// Filter reduces sl to a new slice of all elements for which predicate evaluates to true.
func Filter[TElem any](
	sl []TElem,
	predicate func(TElem) bool,
) []TElem {
	reducedSlice := make([]TElem, 0)
	for _, elem := range sl {
		if predicate(elem) {
			reducedSlice = append(reducedSlice, elem)
		}
	}

	return reducedSlice
}

// Map projects sl into a new slice using fn as the projection function for every element.
func Map[TElem, TNewElem any](
	sl []TElem,
	fn func(TElem) TNewElem,
) []TNewElem {
	reducedSl := make([]TNewElem, 0, len(sl))
	for _, elem := range sl {
		newElem := fn(elem)
		reducedSl = append(reducedSl, newElem)
	}

	return reducedSl
}

// Fold executes a fold operation on sl seeding the result for the first iteration with
// initialResult.
func Fold[TElem, TResult any](
	sl []TElem,
	initialResult TResult,
	fn func(TResult, TElem) (TResult, error),
) (TResult, error) {
	var (
		err    error
		result = initialResult
	)

	for _, elem := range sl {
		result, err = fn(result, elem)
		if err != nil {
			break
		}
	}

	return result, err
}

// ToMap generates a map from sl using keyFn to generate the map key for every element in sl.
func ToMap[TKey comparable, TElem any](sl []TElem, keyFn func(TElem) TKey) map[TKey]TElem {
	m := make(map[TKey]TElem)

	for _, elem := range sl {
		m[keyFn(elem)] = elem
	}

	return m
}
