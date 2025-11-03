CREATE TABLE "users"(
    "id" SERIAL NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    "hash_password" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(15) NOT NULL,
    "email" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
CREATE TABLE "merch"(
    "id" SERIAL NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "count" INTEGER NOT NULL,
    "description" VARCHAR(2000) NOT NULL,
    "type_id" INTEGER NOT NULL,
    "size_id" INTEGER NOT NULL,
    "color_id" INTEGER NOT NULL,
    "photo_id" INTEGER NOT NULL,
    "material_id" INTEGER NOT NULL,
    "gender" INTEGER NULL
);
ALTER TABLE
    "merch" ADD PRIMARY KEY("id");
CREATE TABLE "order"(
    "id" SERIAL NOT NULL,
    "users_id" INTEGER NOT NULL,
    "merch_id" INTEGER NOT NULL,
    "time_order" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "finish_order" TIMESTAMP(0) WITHOUT TIME ZONE NULL
);
ALTER TABLE
    "order" ADD PRIMARY KEY("id");
CREATE TABLE "color_merch"(
    "id" SERIAL NOT NULL,
    "color" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "color_merch" ADD PRIMARY KEY("id");
ALTER TABLE
    "color_merch" ADD CONSTRAINT "color_merch_color_unique" UNIQUE("color");
CREATE TABLE "size_merch"(
    "id" SERIAL NOT NULL,
    "size" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "size_merch" ADD PRIMARY KEY("id");
ALTER TABLE
    "size_merch" ADD CONSTRAINT "size_merch_size_unique" UNIQUE("size");
CREATE TABLE "type_merch"(
    "id" SERIAL NOT NULL,
    "type" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "type_merch" ADD PRIMARY KEY("id");
ALTER TABLE
    "type_merch" ADD CONSTRAINT "type_merch_type_unique" UNIQUE("type");
CREATE TABLE "photo_merch"(
    "id" SERIAL NOT NULL,
    "photo" VARCHAR(1000) NOT NULL
);
ALTER TABLE
    "photo_merch" ADD PRIMARY KEY("id");
ALTER TABLE
    "photo_merch" ADD CONSTRAINT "photo_merch_photo_unique" UNIQUE("photo");
CREATE TABLE "material_merch"(
    "id" SERIAL NOT NULL,
    "material" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "material_merch" ADD PRIMARY KEY("id");
ALTER TABLE
    "material_merch" ADD CONSTRAINT "material_merch_material_unique" UNIQUE("material");
CREATE TABLE "gender"(
    "id" SERIAL NOT NULL,
    "gender" VARCHAR(15) NOT NULL
);
ALTER TABLE
    "gender" ADD PRIMARY KEY("id");
ALTER TABLE
    "gender" ADD CONSTRAINT "gender_gender_unique" UNIQUE("gender");
ALTER TABLE
    "merch" ADD CONSTRAINT "merch_size_id_foreign" FOREIGN KEY("size_id") REFERENCES "size_merch"("id");
ALTER TABLE
    "merch" ADD CONSTRAINT "merch_type_id_foreign" FOREIGN KEY("type_id") REFERENCES "type_merch"("id");
ALTER TABLE
    "merch" ADD CONSTRAINT "merch_gender_foreign" FOREIGN KEY("gender") REFERENCES "gender"("id");
ALTER TABLE
    "merch" ADD CONSTRAINT "merch_color_id_foreign" FOREIGN KEY("color_id") REFERENCES "color_merch"("id");
ALTER TABLE
    "merch" ADD CONSTRAINT "merch_material_id_foreign" FOREIGN KEY("material_id") REFERENCES "material_merch"("id");
ALTER TABLE
    "order" ADD CONSTRAINT "order_merch_id_foreign" FOREIGN KEY("merch_id") REFERENCES "merch"("id");
ALTER TABLE
    "order" ADD CONSTRAINT "order_users_id_foreign" FOREIGN KEY("users_id") REFERENCES "users"("id");
ALTER TABLE
    "merch" ADD CONSTRAINT "merch_photo_id_foreign" FOREIGN KEY("photo_id") REFERENCES "photo_merch"("id");