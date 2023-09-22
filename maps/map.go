// Package maps implements a couple of helpful iteration operations on generic maps.
package maps

// ToSingleton chooses a random element of the map.
// If the map is empty it returns the zero value for the maps respective types.
func ToSingleton[TKey comparable, TElem any](m map[TKey]TElem) (TKey, TElem) {
	for key, elem := range m {
		return key, elem
	}

	return *new(TKey), *new(TElem)
}

// Filter reduces m to a new map of all elements for which predicate evaluates to true.
func Filter[TKey comparable, TElem any](
	m map[TKey]TElem,
	predicate func(TKey, TElem) bool,
) map[TKey]TElem {
	reducedMap := make(map[TKey]TElem)
	for key, elem := range m {
		if predicate(key, elem) {
			reducedMap[key] = elem
		}
	}

	return reducedMap
}

// Map projects m into a new map using fn as the projection function for every key and element.
func Map[TKey, TNewKey comparable, TElem, TNewElem any](
	m map[TKey]TElem,
	fn func(TKey, TElem) (TNewKey, TNewElem),
) map[TNewKey]TNewElem {
	reducedMap := make(map[TNewKey]TNewElem)
	for key, elem := range m {
		newKey, newElem := fn(key, elem)
		reducedMap[newKey] = newElem
	}

	return reducedMap
}

// Fold executes a fold operation on m seeding the result for the first iteration with
// initialResult.
func Fold[TKey comparable, TElem, TResult any](
	m map[TKey]TElem,
	initialResult TResult,
	fn func(TResult, TKey, TElem) (TResult, error),
) (TResult, error) {
	var (
		err    error
		result = initialResult
	)

	for key, elem := range m {
		result, err = fn(result, key, elem)
		if err != nil {
			break
		}
	}

	return result, err
}

// ToSlice returns a slice of all values in m in random order.
func ToSlice[TKey comparable, TElem any](m map[TKey]TElem) []TElem {
	sl := make([]TElem, 0, len(m))

	for _, elem := range m {
		sl = append(sl, elem)
	}

	return sl
}

// Values is an alias for ToSlice.
// It returns a slice of all values in m in random order.
func Values[TKey comparable, TElem any](m map[TKey]TElem) []TElem {
	return ToSlice(m)
}

// Keys returns a slice of all keys in m in random order.
func Keys[TKey comparable, TElem any](m map[TKey]TElem) []TKey {
	sl := make([]TKey, 0, len(m))

	for key := range m {
		sl = append(sl, key)
	}

	return sl
}

// Union returns m1 with all elements and keys of m2 added.
// This includes overwriting elements from m1 with m2 on key collisions.
func Union[TKey comparable, TElem any](m1 map[TKey]TElem, m2 map[TKey]TElem) map[TKey]TElem {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

// Intersect returns a new map with all elements from m2 where their key is also in m1.
func Intersect[TKey comparable, TElem any](m1 map[TKey]TElem, m2 map[TKey]TElem) map[TKey]TElem {
	intersection := make(map[TKey]TElem)
	for k, v := range m2 {
		if _, hasKey := m1[k]; hasKey {
			intersection[k] = v
		}
	}
	return intersection
}
