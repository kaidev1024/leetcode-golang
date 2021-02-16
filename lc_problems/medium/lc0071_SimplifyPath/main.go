func simplifyPath(path string) string {
	stack := make([]string, 0)
	for _, c := range strings.Split(path, "/") {
		if c == "" || c == "." {
			continue
		}

		if c == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		stack = append(stack, c)
	}

	if len(stack) == 0 {
		return "/"
	}

	builder := strings.Builder{}
	for _, v := range stack {
		builder.WriteString("/")
		builder.WriteString(v)
	}

	return builder.String()
}