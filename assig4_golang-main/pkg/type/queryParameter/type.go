package queryParameter

import (
	"Aibek/pkg/type/pagination"
	"Aibek/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}
