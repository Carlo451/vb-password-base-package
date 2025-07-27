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

func (o *Owner) GetName() string {
	return o.name
}

func (o *Owner) ReturnFileContents() string {
	return o.owner
}

func (o *Owner) UpdateFileContents(newOwner string) {
	o.owner = newOwner
}
func (o *Owner) GetDirectoryPath() string {
	return o.path
}

type Reader struct {
	path   string
	name   string "readers"
	reader []string
}

func NewReader(path string, readers []string) *Reader {
	return &Reader{path: path, name: "readers", reader: readers}
}

func (r *Reader) GetName() string {
	return r.name
}

func (r *Reader) ReturnFileContents() string {
	return strings.Join(r.reader, ",")
}

func (o *Reader) UpdateFileContents(newReaders string) {
	o.reader = strings.Split(newReaders, ",")
}
func (o *Reader) GetDirectoryPath() string {
	return o.path
}

type Writer struct {
	path   string
	name   string "writers"
	writer []string
}

func NewWriter(path string, writers []string) *Writer {
	return &Writer{path: path, name: "writers", writer: writers}
}

func (w *Writer) GetName() string {
	return w.name
}

func (w *Writer) ReturnFileContents() string {
	return strings.Join(w.writer, ",")
}

func (o *Writer) UpdateFileContents(newWriters string) {
	o.writer = strings.Split(newWriters, ",")
}

func (o *Writer) GetDirectoryPath() string {
	return o.path
}

type EncryptionId struct {
	path         string
	name         string
	encryptionId string
}

func NewEncryptionId(path, encryptionId string) *EncryptionId {
	return &EncryptionId{path: path, name: "encryptionId", encryptionId: encryptionId}
}

func (e *EncryptionId) GetName() string {
	return e.name
}

func (e *EncryptionId) ReturnFileContents() string {
	return e.encryptionId
}
func (o *EncryptionId) UpdateFileContents(newEncryptionId string) {
	o.encryptionId = newEncryptionId
}

func (o *EncryptionId) GetDirectoryPath() string {
	return o.path
}

type LastEdited struct {
	path       string
	name       string "lastEdited"
	lastEdited time.Time
}

func NewLastEdited(path string) *LastEdited {
	return &LastEdited{path: path, name: "lastEdited"}
}

func (l *LastEdited) GetName() string {
	return l.name
}

func (l *LastEdited) ReturnFileContents() string {
	return l.lastEdited.String()
}

func (o *LastEdited) UpdateFileContents(nil string) {
	o.lastEdited = time.Now()
}

func (o *LastEdited) GetDirectoryPath() string {
	return o.path
}
