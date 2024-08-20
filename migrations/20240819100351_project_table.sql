-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS nit_project (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description VARCHAR(255),
  designer VARCHAR(255) NOT NULL,
  yarn VARCHAR(255) NOT NULL,
  size VARCHAR(255) NOT NULL,
  needles VARCHAR(255) NOT NULL,
  started DATE NOT NULL,
  ended DATE NOT NULL,
  user_id INTEGER NOT NULL,
  inserted TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES nit_user(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE nit_project;
-- +goose StatementEnd
