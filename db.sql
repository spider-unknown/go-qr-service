create table if not exists document_qrs (
    id uuid primary key,
    document_id varchar UNIQUE NOT NULL,
    qr text NOT NULL
)