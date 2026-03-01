package milvus

// FieldSchema represents a field definition in a Milvus collection.
type FieldSchema struct {
	Name     string
	Type     string // Field data type (int8, int16, int32, int64, float32, float64, etc.)
	DataType string
}

// CollectionInfo represents metadata about a Milvus collection.
type CollectionInfo struct {
	Name        string
	Description string
}

// CollectionSchema represents the schema of a Milvus collection.
type CollectionSchema struct {
	Name   string
	Fields []FieldSchema
}

// DatabaseInfo represents metadata about a Milvus database.
type DatabaseInfo struct {
	Name string
}
