package benchmark

// CodePart represents basic bulding block of a benchmark
// swagger:model
type CodePart struct {
	Type     string            `json:"type" binding:"required"`
	Title    string            `json:"title" binding:"omitempty"`
	Language string            `json:"language" binding:"required"`
	Code     string            `json:"code" binding:"required"`
	Options  map[string]string `json:"options" binding:"omitempty"`
}

// CreateRequest represents request data for POST /benchmark
// swagger:model
type CreateRequest struct {
	// The benchmark title
	// example: My benchmark
	Title string `json:"title" binding:"required"`

	// The benchmark setup code
	// example: { "js": "const a = 'test';" }
	SetupCode map[string]string `json:"setup_code" binding:"required"`

	// A platform to run the benchmark
	// example: browser
	Platform string `json:"platform" binding:"required"`

	// Array of code parts
	CodeParts []CodePart `json:"code_parts" binding:"required"`
}
