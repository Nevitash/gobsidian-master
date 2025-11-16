{{range .Files}}
---
# {{.Path}}
---
{{.ReadContent}}  // <--- Changed from {{.Content}}
---
{{end}}