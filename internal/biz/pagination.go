package biz

type ListOption func(*ListOptions)

type ListOptions struct {
	Filters map[string]string
	Tag     string
	Offset  int64
	Limit   int64
}

func ListFilter(filter map[string]string) ListOption {
	return func(o *ListOptions) {
		o.Filters = filter
	}
}

func ListOffset(offset int64) ListOption {
	return func(o *ListOptions) {
		o.Offset = offset
	}
}

func ListLimit(limit int64) ListOption {
	return func(o *ListOptions) {
		o.Limit = limit
	}
}
