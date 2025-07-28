package passwordstoreConfig

import (
	"strings"
	"time"
)

type Owner struct {
	path  string
	name  string
	owner string
}

func NewOwner(path, owner string) *Owner {
	return &Owner{path: path, name: "owner", owner: owner}
}

func (o *Owner) GetFileName() string {
	return o.name
}

func (o *Owner) ReturnFileContents() string {
	return o.owner
}

func (o *Owner) UpdateFileContents(newOwner string) {
	o.owner = newOwner
}
func (o *Owner) GetUnderlyingDirectoryPath() string {
	return o.path
}

type Reader struct {
	path   string
	name   string
	reader []string
}

func NewReader(path string, readers []string) *Reader {
	return &Reader{path: path, name: "readers", reader: readers}
}

func (r *Reader) GetFileName() string {
	return r.name
}

func (r *Reader) ReturnFileContents() string {
	return strings.Join(r.reader, ",")
}

func (r *Reader) UpdateFileContents(newReaders string) {
	r.reader = strings.Split(newReaders, ",")
}
func (r *Reader) GetUnderlyingDirectoryPath() string {
	return r.path
}

type Writer struct {
	path   string
	name   string
	writer []string
}

func NewWriter(path string, writers []string) *Writer {
	return &Writer{path: path, name: "writers", writer: writers}
}

func (w *Writer) GetFileName() string {
	return w.name
}

func (w *Writer) ReturnFileContents() string {
	return strings.Join(w.writer, ",")
}

func (w *Writer) UpdateFileContents(newWriters string) {
	w.writer = strings.Split(newWriters, ",")
}

func (w *Writer) GetUnderlyingDirectoryPath() string {
	return w.path
}

type EncryptionId struct {
	path         string
	name         string
	encryptionId string
}

func NewEncryptionId(path, encryptionId string) *EncryptionId {
	return &EncryptionId{path: path, name: "encryptionId", encryptionId: encryptionId}
}

func (e *EncryptionId) GetFileName() string {
	return e.name
}

func (e *EncryptionId) ReturnFileContents() string {
	return e.encryptionId
}
func (e *EncryptionId) UpdateFileContents(newEncryptionId string) {
	e.encryptionId = newEncryptionId
}

func (e *EncryptionId) GetUnderlyingDirectoryPath() string {
	return e.path
}

type LastEdited struct {
	path       string
	name       string
	lastEdited time.Time
}

func NewLastEdited(path string) *LastEdited {
	return &LastEdited{path: path, name: "lastEdited", lastEdited: time.Now()}
}

func (l *LastEdited) GetFileName() string {
	return l.name
}

func (l *LastEdited) ReturnFileContents() string {
	return l.lastEdited.String()
}

func (l *LastEdited) UpdateFileContents(nil string) {
	l.lastEdited = time.Now()
}

func (l *LastEdited) GetUnderlyingDirectoryPath() string {
	return l.path
}
