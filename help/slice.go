package help

func In[T comparable](data T, slice []T) bool {
	for _, t := range slice {
		if data == t {
			return true
		}
	}
	return false
}
