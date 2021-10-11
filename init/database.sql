DROP TABLE IF EXISTS "cucumbers";
CREATE TABLE "cucumbers"
(
    "id" serial,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "symbol" text,
    "title" text,
    PRIMARY KEY ("id")
);
