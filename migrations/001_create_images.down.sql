DROP POLICY IF EXISTS "Enable read and write access for all users" ON "public"."images";

ALTER TABLE images DISABLE ROW LEVEL SECURITY;

DROP TABLE IF EXISTS "images";
