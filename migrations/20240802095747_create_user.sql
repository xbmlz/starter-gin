-- +goose Up
-- +goose StatementBegin

CREATE TABLE "sys_user" (
  "id" INTEGER NOT NULL,
  "username" TEXT NOT NULL,
  "nickname" TEXT NOT NULL,
  "gender" INTEGER NOT NULL,
  "phone" TEXT,
  "password" TEXT NOT NULL,
  "email" TEXT UNIQUE NOT NULL,
  "avatar_url" TEXT,
  "status" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);

INSERT INTO sys_user (username, nickname, gender, phone, password, email, avatar_url, status)
VALUES ('admin', 'admin', 1, '13811112222', '$2a$10$05SaFGfrDvckPCV54xTNWezvpzL1JzJn8uwYAoXvuQR.Fe.0rqogy', 'admin@localhost.com', 'https://unpkg.com/outeres@0.0.5/demo/000.jpg', 1);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sys_user;
-- +goose StatementEnd
