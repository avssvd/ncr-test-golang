CREATE TABLE "controllers" (
                               "serial" int PRIMARY KEY,
                               "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "indications" (
                               "id" BIGSERIAL PRIMARY KEY,
                               "indication" numeric(3,1) NOT NULL,
                               "controller_serial" int NOT NULL,
                               "sent_at" timestamp NOT NULL,
                               "created_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "indications" ADD FOREIGN KEY ("controller_serial") REFERENCES "controllers" ("serial") ON DELETE CASCADE;