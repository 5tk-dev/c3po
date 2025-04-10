package c3po

// use 'validate' in tags
func Validate(sch, data any) Schema {
	f := ParseSchemaWithTag("validate", sch)
	return f.Decode(data)
}
