-- Modify "users" table
ALTER TABLE "users" DROP COLUMN "clinic_users";
-- Create "clinic_users" table
CREATE TABLE "clinic_users" ("id" uuid NOT NULL, "clinic_id" uuid NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "clinic_users_clinics_clinic_users" FOREIGN KEY ("clinic_id") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "clinic_users_users_clinic_users" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
