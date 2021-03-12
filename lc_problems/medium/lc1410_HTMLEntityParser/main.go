func entityParser(text string) string {
	r := strings.NewReplacer("&quot;", "\"", "&apos;", "'", "&amp;", "&", "&gt;", ">", "&lt;", "<", "&frasl;", "/")
	return r.Replace(text)
}