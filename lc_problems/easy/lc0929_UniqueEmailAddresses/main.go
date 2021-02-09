import "strings"

func numUniqueEmails(emails []string) int {
	parse := func(s string) string {
		fields := strings.Split(s, "@")
		locals := strings.Split(fields[0], "+")
		withoutDots := strings.Split(locals[0], ".")
		return strings.Join(withoutDots, "") + "@" + fields[1]
	}

	m := make(map[string]struct{})
	for i, n := 0, len(emails); i < n; i++ {
		m[parse(emails[i])] = struct{}{}
	}

	return len(m)
}