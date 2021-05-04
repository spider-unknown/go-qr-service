package model

type Document struct {
	Id         string `db:"id"`
	DocumentId string `db:"document_id"`
	QR         string `db:"qr"`
}

type Documents struct {
	Docs []*Document
}
