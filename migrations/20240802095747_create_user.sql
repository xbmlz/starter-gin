-- +goose Up
-- +goose StatementBegin

CREATE TABLE "sys_user" (
  "id" INTEGER NOT NULL,
  "username" TEXT,
  "password" TEXT,
  "email" TEXT,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);

INSERT INTO sys_user (username, password, email) VALUES ('admin', '$2a$10$05SaFGfrDvckPCV54xTNWezvpzL1JzJn8uwYAoXvuQR.Fe.0rqogy', 'admin@localhost');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sys_user;
-- +goose StatementEnd
