-- +goose Up
-- +goose StatementBegin
CREATE TABLE "nit_user" (
  "id" SERIAL PRIMARY KEY,
  "fb_id" INT UNIQUE NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "nit_user";
-- +goose StatementEnd
