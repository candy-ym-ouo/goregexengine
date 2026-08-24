package regex

var sharedGroups []Group

func groupBuffer() []Group          { return sharedGroups[:0] }
func saveGroupBuffer(value []Group) { sharedGroups = value }
