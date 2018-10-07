package datatypes

// Common data types that are used on the clipboard.
const (
	PlainText = `text/plain`
	RTFText   = "text/rtf"
)

// Data holds the data for a clipboard.
type Data struct {
	MimeType string
	Bytes    []byte
}
