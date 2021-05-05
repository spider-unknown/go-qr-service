package repository

import (
	"context"
	"database/sql"
	"document-qr/internal/app/model"
	pb "document-qr/proto"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/unistack-org/micro/v3/codec"
)

type DocumentRepositorySQL struct {
	db    *sqlx.DB
	codec codec.Codec
}

func NewDocumentRepository(db *sqlx.DB, codec codec.Codec) DocumentRepository {
	return &DocumentRepositorySQL{db: db, codec: codec}
}
func (dr *DocumentRepositorySQL) FindById(ctx context.Context, request *pb.PostQRRequest) (*model.Document, error) {
	tx := dr.db.MustBeginTx(ctx, &sql.TxOptions{ReadOnly: true})

	document := &model.Document{}
	query := "select * from document_qrs where document_id = $1"
	sqlErr := tx.QueryRowContext(
		ctx,
		query,
		request.DocumentId,
	).Scan(&document.Id, &document.DocumentId, &document.QR)

	if sqlErr != nil {
		tx.Rollback()
		return nil, sqlErr
	}
	tx.Commit()
	return document, nil
}

func (dr *DocumentRepositorySQL) Create(ctx context.Context, request *pb.PostQRRequest, qr string) (*model.Document, error) {
	tx := dr.db.MustBeginTx(ctx, &sql.TxOptions{})
	document := &model.Document{}

	sqlErr := tx.QueryRowContext(
		ctx,
		"insert into document_qrs (id, document_id, qr) values ($1, $2, $3) returning id",
		uuid.New().String(),
		request.DocumentId,
		qr).Scan(&document.Id)
	if sqlErr != nil {
		tx.Rollback()
		return nil, &pb.Error{Msg: sqlErr.Error()}
	}
	tx.Commit()
	return document, nil
}

func (dr *DocumentRepositorySQL) CreateEmptyDocument(ctx context.Context) (*model.Document, error) {
	tx := dr.db.MustBeginTx(ctx, &sql.TxOptions{})
	document := &model.Document{}

	sqlErr := tx.QueryRowContext(
		ctx,
		"insert into document_qrs (id) values ($1) returning id",
		uuid.New().String()).Scan(&document.Id)
	if sqlErr != nil {
		tx.Rollback()
		return nil, &pb.Error{Msg: sqlErr.Error()}
	}
	tx.Commit()
	return document, nil
}
