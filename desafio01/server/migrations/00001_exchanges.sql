
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS exchanges (
  id string PRIMARY KEY,
  currency VARCHAR(3) NOT NULL,
  desired_currency VARCHAR(3) NOT NULL,
  bid VARCHAR(250) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd