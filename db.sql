create table if not exists document_qrs
(
    id          varchar primary key,
    document_id varchar UNIQUE,
    qr          text
)