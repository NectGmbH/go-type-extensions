package slices

func ToSingleton[TElem any](sl []TElem) TElem {
	for _, elem := range sl {
		return elem
	}

	return *new(TElem)
}

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

func Fold[TElem, TResult any](
	sl []TElem,
	fn func(TResult, TElem) (TResult, error),
) (TResult, error) {
	var (
		err    error
		result TResult
	)

	for _, elem := range sl {
		result, err = fn(result, elem)
		if err != nil {
			break
		}
	}

	return result, err
}

func ToMap[TKey comparable, TElem any](sl []TElem, keyFn func(TElem) TKey) map[TKey]TElem {
	m := make(map[TKey]TElem)

	for _, elem := range sl {
		m[keyFn(elem)] = elem
	}

	return m
}
