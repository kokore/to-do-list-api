package response

const _internalCode = 100000

const (
	InternalError uint64 = _internalCode + iota
	InvalidRequestJSONString
	InvalidRequestRouteParams
	InvalidQueryParams
	UnableToFindToDoList
	UnableToCreateToDoList
	UnableToUpdateToDoList
)
