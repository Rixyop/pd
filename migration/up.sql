DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "passwords" CASCADE;
DROP TABLE IF EXISTS "battalions" CASCADE;
DROP TABLE IF EXISTS "companies" CASCADE;
DROP TABLE IF EXISTS "clusters" CASCADE;
DROP TABLE IF EXISTS "courses" CASCADE;
DROP TABLE IF EXISTS "course_coaches" CASCADE;
DROP TABLE IF EXISTS "coach_absenteeism" CASCADE;
DROP TABLE IF EXISTS "classes" CASCADE;
DROP TABLE IF EXISTS "garrisons" CASCADE;

CREATE TABLE "garrisons" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "location" TEXT,
    "creator" VARCHAR(40) NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);

CREATE TABLE "users" (
    "id" VARCHAR(40) PRIMARY KEY,
    "first_name" TEXT NOT NULL,
    "last_name" TEXT NOT NULL,
    "username" TEXT NOT NULL,
    "role" TEXT NOT NULL,
    "avatar" TEXT NOT NULL,
    "garrison_id" INTEGER REFERENCES "garrisons" (id),
    "created_at" TIMESTAMP NOT NULL
);

CREATE TABLE "passwords" (
    "user_id" VARCHAR(40) PRIMARY KEY REFERENCES "users" (id),
    "password" TEXT NOT NULL,
    "last_update_at" TIMESTAMP NOT NULL
);

CREATE TABLE "battalions" (
    "id" SERIAL PRIMARY KEY,
    "garrison_id" INTEGER NOT NULL REFERENCES "garrisons" (id),
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);

CREATE TABLE "companies" (
    "id" SERIAL PRIMARY KEY,
    "garrison_id" INTEGER NOT NULL REFERENCES "garrisons" (id),
    "battalion_id" INTEGER NOT NULL REFERENCES "battalions" (id),
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);

CREATE TABLE "clusters" (
    "id" SERIAL PRIMARY KEY,
    "garrison_id" INTEGER NOT NULL REFERENCES "garrisons" (id),
    "battalion_id" INTEGER REFERENCES "battalions" (id), 
    "company_id" INTEGER REFERENCES "companies" (id),
    "name" TEXT NOT NULL,
    "created_at" TIMESTAMP NOT NULL
);

CREATE TABLE "courses"(
    "id" VARCHAR(40) PRIMARY KEY,
    "name" TEXT NOT NULL,
    "count" INTEGER NOT NULL,
    "priority" INTEGER NOT NULL,
    "class_time" INTEGER NOT NULL
);

CREATE TABLE "course_coaches"(
    "course_id" VARCHAR(40) NOT NULL REFERENCES "courses" (id),
    "coach_id" VARCHAR(40) REFERENCES "users" (id),
    UNIQUE (course_id, coach_id)
);

CREATE TABLE "coach_absenteeism"(
    "coach_id" VARCHAR(40) NOT NULL REFERENCES "users" (id),
    "week_days" TEXT[] NOT NULL
);

CREATE TABLE "classes"(
    "id" VARCHAR(40) PRIMARY KEY,
    "course_id" VARCHAR(40) NOT NULL,
    "start_date" TIMESTAMP NOT NULL,
    "end_date" TIMESTAMP NOT NULL,
    FOREIGN KEY (course_id) REFERENCES "courses" (id)
);



-- Insert data into tables
INSERT INTO "garrisons" ("name", "location", "creator", "created_at") 
VALUES ('پادگان1', 'بابل', 'admin', NOW());

INSERT INTO "users" ("id", "first_name", "last_name", "username", "role", "avatar", "garrison_id", "created_at")
VALUES ('6dXTEwafl0fjH9eg$LYTdJDLjYzh7Mo1M1PdaAfg', 'امیر', 'رضایی', 'amir', 'class_affairs', 'person', 
    (SELECT "id" FROM "garrisons" WHERE "name" = 'پادگان1'), NOW());

INSERT INTO "passwords" ("user_id", "password", "last_update_at")
VALUES ('6dXTEwafl0fjH9eg$LYTdJDLjYzh7Mo1M1PdaAfg', '$argon2id$v=19$m=65536,t=3,p=4$PhSype6dXTEwafl0fjH9eg$LYTdJDLjYzh7Mo1M1P/oSrylWYAzrsLCMLeDH9GsikI', NOW());

INSERT INTO "battalions" ("garrison_id", "name", "created_at")
VALUES ((SELECT "id" FROM "garrisons" WHERE "name" = 'پادگان1'), 'انتقامی', NOW());

INSERT INTO "companies" ("garrison_id", "battalion_id", "name", "created_at")
VALUES ((SELECT "id" FROM "garrisons" WHERE "name" = 'پادگان1'), 
    (SELECT "id" FROM "battalions" WHERE "name" = 'انتقامی'), 'گروهان 2', NOW());

INSERT INTO "clusters" ("garrison_id", "battalion_id", "company_id", "name", "created_at")
VALUES ((SELECT "id" FROM "garrisons" WHERE "name" = 'پادگان1'), 
    (SELECT "id" FROM "battalions" WHERE "name" = 'انتقامی'), 
    (SELECT "id" FROM "companies" WHERE "name" = 'گروهان 2'), 'دسته 1', NOW());