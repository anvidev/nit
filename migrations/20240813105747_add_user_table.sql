-- +goose Up
-- +goose StatementBegin
CREATE TABLE nit_user (
  id SERIAL PRIMARY KEY,
  facebook_id BIGINT UNIQUE NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE nit_user;
-- +goose StatementEnd
