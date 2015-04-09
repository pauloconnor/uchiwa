package uchiwa

// BuildTags builds a slice of every client tag
func BuildTags() {
	for _, c := range tmpResults.Clients {
		m := c.(map[string]interface{})
		switch x := m["tags"].(type) {
		case map[string]interface{}:
			for k, _ := range x {
				if !isStringInSlice(tmpResults.Tags, k) {
					tmpResults.Tags = append(tmpResults.Tags, k)
				}
			}
		default:
			// I don't know.
		}
	}
}
