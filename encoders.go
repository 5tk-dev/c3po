package c3po

// use 'validate' in tags
func Validate(sch, data any) Schema {
	return ParseSchemaWithTag("validate", sch).Decode(data)
}
