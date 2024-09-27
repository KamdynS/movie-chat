CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "username" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL
)