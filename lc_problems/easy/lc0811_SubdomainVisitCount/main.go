func subdomainVisits(cpdomains []string) []string {
	if cpdomains == nil || len(cpdomains) == 0 {
		return make([]string, 0)
	}

	hashMap := make(map[string]int)
	for _, domain := range cpdomains {
		addToMap(domain, hashMap)
	}

	res := make([]string, 0)
	for domain, count := range hashMap {
		res = append(res, strconv.Itoa(count)+" "+domain)
	}
	return res
}

func addToMap(domain string, hashMap map[string]int) {
	count, _ := strconv.Atoi(strings.Split(domain, " ")[0])
	subDomain := strings.Split(domain, " ")[1]
	for len(subDomain) > 0 {
		hashMap[subDomain] = hashMap[subDomain] + count // hashMap[subDomain] += count     also works
		if !strings.ContainsRune(subDomain, '.') {
			break
		}
		subDomain = subDomain[strings.IndexRune(subDomain, '.')+1:]
	}
}