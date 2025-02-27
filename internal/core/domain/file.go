package domain

type File struct {
	ContentType string
	Content     []byte
}

func NewFile(contentType string, content []byte) *File {
	return &File{
		ContentType: contentType,
		Content:     content,
	}
}
