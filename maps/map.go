package maps

func ToSingleton[TKey comparable, TElem any](m map[TKey]TElem) (TKey, TElem) {
	for key, elem := range m {
		return key, elem
	}

	return *new(TKey), *new(TElem)
}

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

func Fold[TKey comparable, TElem, TResult any](
	m map[TKey]TElem,
	fn func(TResult, TKey, TElem) (TResult, error),
) (TResult, error) {
	var (
		err    error
		result TResult
	)

	for key, elem := range m {
		result, err = fn(result, key, elem)
		if err != nil {
			break
		}
	}

	return result, err
}

func ToSlice[TKey comparable, TElem any](m map[TKey]TElem) []TElem {
	sl := make([]TElem, 0, len(m))

	for _, elem := range m {
		sl = append(sl, elem)
	}

	return sl
}

func Keys[TKey comparable, TElem any](m map[TKey]TElem) []TKey {
	sl := make([]TKey, 0, len(m))

	for key := range m {
		sl = append(sl, key)
	}

	return sl
}
