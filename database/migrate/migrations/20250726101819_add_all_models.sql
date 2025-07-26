-- Create "ai_interactions" table
CREATE TABLE "ai_interactions" ("id" uuid NOT NULL, "clinic_id" character varying NOT NULL, "patient_whatsapp" character varying NULL, "interaction_type" character varying NOT NULL DEFAULT 'CHAT', "request_payload" jsonb NOT NULL, "response_payload" jsonb NOT NULL, "ai_model" character varying NULL, "response_time_ms" bigint NULL, "status" character varying NOT NULL DEFAULT 'SUCCESS', "error_message" text NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "aiinteraction_clinic_id" to table: "ai_interactions"
CREATE INDEX "aiinteraction_clinic_id" ON "ai_interactions" ("clinic_id");
-- Create index "aiinteraction_created_at" to table: "ai_interactions"
CREATE INDEX "aiinteraction_created_at" ON "ai_interactions" ("created_at");
-- Create index "aiinteraction_interaction_type" to table: "ai_interactions"
CREATE INDEX "aiinteraction_interaction_type" ON "ai_interactions" ("interaction_type");
-- Create index "aiinteraction_status" to table: "ai_interactions"
CREATE INDEX "aiinteraction_status" ON "ai_interactions" ("status");
-- Create "admin_audit_logs" table
CREATE TABLE "admin_audit_logs" ("id" uuid NOT NULL, "user_email" character varying NOT NULL, "operation" character varying NOT NULL, "model" character varying NOT NULL, "args" jsonb NOT NULL, "result" jsonb NULL, "query" text NULL, "params" text NULL, "source" character varying NOT NULL DEFAULT 'admin-panel', "duration_ms" bigint NULL, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create "queue_entries" table
CREATE TABLE "queue_entries" ("id" uuid NOT NULL, "clinic_id" character varying NOT NULL, "patient_id" character varying NOT NULL, "doctor_id" character varying NULL, "service_id" character varying NULL, "queue_number" bigint NOT NULL, "status" character varying NOT NULL DEFAULT 'WAITING', "estimated_time" timestamptz NULL, "called_at" timestamptz NULL, "completed_at" timestamptz NULL, "notes" text NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "queueentry_clinic_id_queue_number" to table: "queue_entries"
CREATE INDEX "queueentry_clinic_id_queue_number" ON "queue_entries" ("clinic_id", "queue_number");
-- Create index "queueentry_created_at" to table: "queue_entries"
CREATE INDEX "queueentry_created_at" ON "queue_entries" ("created_at");
-- Create index "queueentry_status" to table: "queue_entries"
CREATE INDEX "queueentry_status" ON "queue_entries" ("status");
-- Create "features" table
CREATE TABLE "features" ("id" uuid NOT NULL, "name" character varying NOT NULL, "display_name" character varying NOT NULL, "description" text NULL, "category" character varying NOT NULL DEFAULT 'CORE', "monthly_price" double precision NOT NULL DEFAULT 0, "active" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "features_name_key" to table: "features"
CREATE UNIQUE INDEX "features_name_key" ON "features" ("name");
-- Create "clinics" table
CREATE TABLE "clinics" ("id" uuid NOT NULL, "name" character varying NOT NULL, "type" character varying NOT NULL, "phone" character varying NULL, "email" character varying NULL, "address" text NULL, "business_hours" jsonb NULL, "whatsapp_number" character varying NULL, "subscription_plan" character varying NOT NULL DEFAULT 'basic', "enabled_features" jsonb NOT NULL, "active" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- Create index "clinic_active" to table: "clinics"
CREATE INDEX "clinic_active" ON "clinics" ("active");
-- Create index "clinic_whatsapp_number" to table: "clinics"
CREATE INDEX "clinic_whatsapp_number" ON "clinics" ("whatsapp_number");
-- Create "doctors" table
CREATE TABLE "doctors" ("id" uuid NOT NULL, "name" character varying NOT NULL, "specialization" character varying NOT NULL, "license_number" character varying NULL, "email" character varying NULL, "phone" character varying NULL, "bio" text NULL, "qualifications" jsonb NULL, "availability" jsonb NOT NULL, "consultation_duration" bigint NOT NULL DEFAULT 30, "consultation_fee" double precision NOT NULL DEFAULT 0, "active" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_doctors" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "doctors_clinics_doctors" FOREIGN KEY ("clinic_doctors") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "patients" table
CREATE TABLE "patients" ("id" uuid NOT NULL, "whatsapp_number" character varying NOT NULL, "name" character varying NULL, "email" character varying NULL, "phone" character varying NULL, "birth_date" timestamptz NULL, "gender" character varying NULL, "address" text NULL, "medical_history" jsonb NULL, "allergies" jsonb NULL, "emergency_contact_name" character varying NULL, "emergency_contact_phone" character varying NULL, "active" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_patients" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "patients_clinics_patients" FOREIGN KEY ("clinic_patients") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "patient_email" to table: "patients"
CREATE INDEX "patient_email" ON "patients" ("email");
-- Create index "patient_whatsapp_number" to table: "patients"
CREATE INDEX "patient_whatsapp_number" ON "patients" ("whatsapp_number");
-- Create "services" table
CREATE TABLE "services" ("id" uuid NOT NULL, "name" character varying NOT NULL, "description" text NULL, "category" character varying NULL, "price" double precision NOT NULL DEFAULT 0, "duration" bigint NOT NULL DEFAULT 30, "requirements" jsonb NULL, "requires_appointment" boolean NOT NULL DEFAULT true, "active" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_services" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "services_clinics_services" FOREIGN KEY ("clinic_services") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "appointments" table
CREATE TABLE "appointments" ("id" uuid NOT NULL, "appointment_date" timestamptz NOT NULL, "start_time" timestamptz NOT NULL, "end_time" timestamptz NOT NULL, "status" character varying NOT NULL DEFAULT 'SCHEDULED', "notes" text NULL, "symptoms" text NULL, "diagnosis" text NULL, "treatment_plan" text NULL, "prescriptions" jsonb NULL, "total_cost" double precision NULL, "payment_status" character varying NOT NULL DEFAULT 'PENDING', "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_appointments" uuid NULL, "doctor_appointments" uuid NULL, "patient_appointments" uuid NULL, "service_appointments" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "appointments_clinics_appointments" FOREIGN KEY ("clinic_appointments") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "appointments_doctors_appointments" FOREIGN KEY ("doctor_appointments") REFERENCES "doctors" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "appointments_patients_appointments" FOREIGN KEY ("patient_appointments") REFERENCES "patients" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "appointments_services_appointments" FOREIGN KEY ("service_appointments") REFERENCES "services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "appointment_appointment_date_start_time" to table: "appointments"
CREATE INDEX "appointment_appointment_date_start_time" ON "appointments" ("appointment_date", "start_time");
-- Create index "appointment_status" to table: "appointments"
CREATE INDEX "appointment_status" ON "appointments" ("status");
-- Create "appointment_reminders" table
CREATE TABLE "appointment_reminders" ("id" uuid NOT NULL, "type" character varying NOT NULL, "scheduled_time" timestamptz NOT NULL, "message" character varying NOT NULL, "status" character varying NOT NULL DEFAULT 'PENDING', "sent_at" timestamptz NULL, "created_at" timestamptz NOT NULL, "appointment_reminders" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "appointment_reminders_appointments_reminders" FOREIGN KEY ("appointment_reminders") REFERENCES "appointments" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "billing_records" table
CREATE TABLE "billing_records" ("id" uuid NOT NULL, "invoice_number" character varying NOT NULL, "amount" double precision NOT NULL, "tax_amount" double precision NOT NULL DEFAULT 0, "discount_amount" double precision NOT NULL DEFAULT 0, "total_amount" double precision NOT NULL, "currency" character varying NOT NULL DEFAULT 'IDR', "payment_method" character varying NULL, "payment_status" character varying NOT NULL DEFAULT 'PENDING', "line_items" jsonb NULL, "due_date" timestamptz NULL, "paid_at" timestamptz NULL, "notes" text NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_billing_records" uuid NULL, "patient_billing_records" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "billing_records_clinics_billing_records" FOREIGN KEY ("clinic_billing_records") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "billing_records_patients_billing_records" FOREIGN KEY ("patient_billing_records") REFERENCES "patients" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "billing_records_invoice_number_key" to table: "billing_records"
CREATE UNIQUE INDEX "billing_records_invoice_number_key" ON "billing_records" ("invoice_number");
-- Create index "billingrecord_due_date" to table: "billing_records"
CREATE INDEX "billingrecord_due_date" ON "billing_records" ("due_date");
-- Create index "billingrecord_invoice_number" to table: "billing_records"
CREATE INDEX "billingrecord_invoice_number" ON "billing_records" ("invoice_number");
-- Create index "billingrecord_payment_status" to table: "billing_records"
CREATE INDEX "billingrecord_payment_status" ON "billing_records" ("payment_status");
-- Create "chat_threads" table
CREATE TABLE "chat_threads" ("id" uuid NOT NULL, "whatsapp_thread_id" character varying NULL, "status" character varying NOT NULL DEFAULT 'ACTIVE', "context" jsonb NULL, "last_message_at" timestamptz NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_chat_threads" uuid NULL, "patient_chat_threads" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "chat_threads_clinics_chat_threads" FOREIGN KEY ("clinic_chat_threads") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "chat_threads_patients_chat_threads" FOREIGN KEY ("patient_chat_threads") REFERENCES "patients" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "chatthread_last_message_at" to table: "chat_threads"
CREATE INDEX "chatthread_last_message_at" ON "chat_threads" ("last_message_at");
-- Create index "chatthread_status" to table: "chat_threads"
CREATE INDEX "chatthread_status" ON "chat_threads" ("status");
-- Create index "chatthread_whatsapp_thread_id" to table: "chat_threads"
CREATE INDEX "chatthread_whatsapp_thread_id" ON "chat_threads" ("whatsapp_thread_id");
-- Create "chat_messages" table
CREATE TABLE "chat_messages" ("id" uuid NOT NULL, "whatsapp_message_id" character varying NULL, "sender_type" character varying NOT NULL, "message_type" character varying NOT NULL DEFAULT 'TEXT', "content" text NOT NULL, "metadata" jsonb NULL, "ai_tool_call" character varying NULL, "ai_tool_result" jsonb NULL, "is_read" boolean NOT NULL DEFAULT false, "created_at" timestamptz NOT NULL, "chat_thread_messages" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "chat_messages_chat_threads_messages" FOREIGN KEY ("chat_thread_messages") REFERENCES "chat_threads" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "chatmessage_created_at" to table: "chat_messages"
CREATE INDEX "chatmessage_created_at" ON "chat_messages" ("created_at");
-- Create index "chatmessage_sender_type" to table: "chat_messages"
CREATE INDEX "chatmessage_sender_type" ON "chat_messages" ("sender_type");
-- Create index "chatmessage_whatsapp_message_id" to table: "chat_messages"
CREATE INDEX "chatmessage_whatsapp_message_id" ON "chat_messages" ("whatsapp_message_id");
-- Create "doctor_schedules" table
CREATE TABLE "doctor_schedules" ("id" uuid NOT NULL, "date" timestamptz NOT NULL, "start_time" timestamptz NOT NULL, "end_time" timestamptz NOT NULL, "available" boolean NOT NULL DEFAULT true, "notes" character varying NULL, "created_at" timestamptz NOT NULL, "doctor_schedules" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "doctor_schedules_doctors_schedules" FOREIGN KEY ("doctor_schedules") REFERENCES "doctors" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "doctorschedule_date_start_time" to table: "doctor_schedules"
CREATE INDEX "doctorschedule_date_start_time" ON "doctor_schedules" ("date", "start_time");
-- Create "documents" table
CREATE TABLE "documents" ("id" uuid NOT NULL, "name" character varying NOT NULL, "type" character varying NOT NULL DEFAULT 'MEDICAL_RECORD', "file_path" character varying NOT NULL, "file_type" character varying NOT NULL, "file_size" bigint NOT NULL, "metadata" jsonb NULL, "document_date" timestamptz NULL, "is_confidential" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, "clinic_documents" uuid NULL, "patient_documents" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "documents_clinics_documents" FOREIGN KEY ("clinic_documents") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "documents_patients_documents" FOREIGN KEY ("patient_documents") REFERENCES "patients" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "document_document_date" to table: "documents"
CREATE INDEX "document_document_date" ON "documents" ("document_date");
-- Create index "document_type" to table: "documents"
CREATE INDEX "document_type" ON "documents" ("type");
-- Create "product_categories" table
CREATE TABLE "product_categories" ("id" uuid NOT NULL, "name" character varying NOT NULL, "description" text NULL, "image_url" character varying NULL, "active" boolean NOT NULL DEFAULT true, "sort_order" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_product_categories" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "product_categories_clinics_product_categories" FOREIGN KEY ("clinic_product_categories") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "productcategory_active" to table: "product_categories"
CREATE INDEX "productcategory_active" ON "product_categories" ("active");
-- Create index "productcategory_sort_order" to table: "product_categories"
CREATE INDEX "productcategory_sort_order" ON "product_categories" ("sort_order");
-- Create "products" table
CREATE TABLE "products" ("id" uuid NOT NULL, "sku" character varying NOT NULL, "name" character varying NOT NULL, "description" text NULL, "short_description" text NULL, "brand" character varying NULL, "images" jsonb NULL, "purchase_price" double precision NOT NULL DEFAULT 0, "selling_price" double precision NOT NULL DEFAULT 0, "discount_price" double precision NULL, "unit" character varying NOT NULL DEFAULT 'pcs', "min_stock_level" bigint NOT NULL DEFAULT 0, "current_stock" bigint NOT NULL DEFAULT 0, "track_inventory" boolean NOT NULL DEFAULT true, "prescription_required" boolean NOT NULL DEFAULT false, "specifications" jsonb NULL, "usage_instructions" jsonb NULL, "warnings" jsonb NULL, "expiry_date" timestamptz NULL, "batch_number" character varying NULL, "status" character varying NOT NULL DEFAULT 'ACTIVE', "tags" jsonb NULL, "weight" double precision NULL, "dimensions" jsonb NULL, "featured" boolean NOT NULL DEFAULT false, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_products" uuid NULL, "product_category_products" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "products_clinics_products" FOREIGN KEY ("clinic_products") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "products_product_categories_products" FOREIGN KEY ("product_category_products") REFERENCES "product_categories" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "product_current_stock" to table: "products"
CREATE INDEX "product_current_stock" ON "products" ("current_stock");
-- Create index "product_featured" to table: "products"
CREATE INDEX "product_featured" ON "products" ("featured");
-- Create index "product_prescription_required" to table: "products"
CREATE INDEX "product_prescription_required" ON "products" ("prescription_required");
-- Create index "product_sku" to table: "products"
CREATE INDEX "product_sku" ON "products" ("sku");
-- Create index "product_status" to table: "products"
CREATE INDEX "product_status" ON "products" ("status");
-- Create "inventory_movements" table
CREATE TABLE "inventory_movements" ("id" uuid NOT NULL, "type" character varying NOT NULL DEFAULT 'SALE', "quantity" bigint NOT NULL, "reference_id" character varying NULL, "notes" text NULL, "performed_by" character varying NULL, "movement_date" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, "clinic_inventory_movements" uuid NULL, "product_inventory_movements" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "inventory_movements_clinics_inventory_movements" FOREIGN KEY ("clinic_inventory_movements") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "inventory_movements_products_inventory_movements" FOREIGN KEY ("product_inventory_movements") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "inventorymovement_movement_date" to table: "inventory_movements"
CREATE INDEX "inventorymovement_movement_date" ON "inventory_movements" ("movement_date");
-- Create index "inventorymovement_reference_id" to table: "inventory_movements"
CREATE INDEX "inventorymovement_reference_id" ON "inventory_movements" ("reference_id");
-- Create index "inventorymovement_type" to table: "inventory_movements"
CREATE INDEX "inventorymovement_type" ON "inventory_movements" ("type");
-- Create "knowledge_bases" table
CREATE TABLE "knowledge_bases" ("id" uuid NOT NULL, "title" character varying NOT NULL, "category" character varying NOT NULL DEFAULT 'FAQ', "content" text NOT NULL, "tags" jsonb NULL, "metadata" jsonb NULL, "active" boolean NOT NULL DEFAULT true, "priority" bigint NOT NULL DEFAULT 0, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_knowledge_base" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "knowledge_bases_clinics_knowledge_base" FOREIGN KEY ("clinic_knowledge_base") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "knowledgebase_active" to table: "knowledge_bases"
CREATE INDEX "knowledgebase_active" ON "knowledge_bases" ("active");
-- Create index "knowledgebase_category" to table: "knowledge_bases"
CREATE INDEX "knowledgebase_category" ON "knowledge_bases" ("category");
-- Create index "knowledgebase_priority" to table: "knowledge_bases"
CREATE INDEX "knowledgebase_priority" ON "knowledge_bases" ("priority");
-- Create "orders" table
CREATE TABLE "orders" ("id" uuid NOT NULL, "order_number" character varying NOT NULL, "order_type" character varying NOT NULL DEFAULT 'MIXED', "status" character varying NOT NULL DEFAULT 'PENDING', "subtotal" double precision NOT NULL DEFAULT 0, "tax_amount" double precision NOT NULL DEFAULT 0, "discount_amount" double precision NOT NULL DEFAULT 0, "shipping_cost" double precision NOT NULL DEFAULT 0, "total_amount" double precision NOT NULL DEFAULT 0, "currency" character varying NOT NULL DEFAULT 'IDR', "payment_status" character varying NOT NULL DEFAULT 'PENDING', "payment_method" character varying NULL, "shipping_address" jsonb NULL, "billing_address" jsonb NULL, "delivery_method" character varying NOT NULL DEFAULT 'PICKUP', "expected_delivery_date" timestamptz NULL, "delivered_at" timestamptz NULL, "notes" text NULL, "cancellation_reason" text NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_orders" uuid NULL, "patient_orders" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "orders_clinics_orders" FOREIGN KEY ("clinic_orders") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "orders_patients_orders" FOREIGN KEY ("patient_orders") REFERENCES "patients" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "order_created_at" to table: "orders"
CREATE INDEX "order_created_at" ON "orders" ("created_at");
-- Create index "order_order_number" to table: "orders"
CREATE INDEX "order_order_number" ON "orders" ("order_number");
-- Create index "order_order_type" to table: "orders"
CREATE INDEX "order_order_type" ON "orders" ("order_type");
-- Create index "order_payment_status" to table: "orders"
CREATE INDEX "order_payment_status" ON "orders" ("payment_status");
-- Create index "order_status" to table: "orders"
CREATE INDEX "order_status" ON "orders" ("status");
-- Create index "orders_order_number_key" to table: "orders"
CREATE UNIQUE INDEX "orders_order_number_key" ON "orders" ("order_number");
-- Create "order_items" table
CREATE TABLE "order_items" ("id" uuid NOT NULL, "item_type" character varying NOT NULL DEFAULT 'PRODUCT', "item_name" character varying NOT NULL, "item_description" text NULL, "quantity" bigint NOT NULL DEFAULT 1, "unit_price" double precision NOT NULL, "discount_amount" double precision NOT NULL DEFAULT 0, "total_price" double precision NOT NULL, "item_metadata" jsonb NULL, "created_at" timestamptz NOT NULL, "appointment_order_items" uuid NULL, "order_order_items" uuid NULL, "product_order_items" uuid NULL, "service_order_items" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "order_items_appointments_order_items" FOREIGN KEY ("appointment_order_items") REFERENCES "appointments" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "order_items_orders_order_items" FOREIGN KEY ("order_order_items") REFERENCES "orders" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "order_items_products_order_items" FOREIGN KEY ("product_order_items") REFERENCES "products" ("id") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "order_items_services_order_items" FOREIGN KEY ("service_order_items") REFERENCES "services" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "orderitem_item_type" to table: "order_items"
CREATE INDEX "orderitem_item_type" ON "order_items" ("item_type");
-- Create "order_status_histories" table
CREATE TABLE "order_status_histories" ("id" uuid NOT NULL, "status" character varying NOT NULL, "notes" text NULL, "changed_by" character varying NULL, "created_at" timestamptz NOT NULL, "order_order_status_history" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "order_status_histories_orders_order_status_history" FOREIGN KEY ("order_order_status_history") REFERENCES "orders" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "orderstatushistory_created_at" to table: "order_status_histories"
CREATE INDEX "orderstatushistory_created_at" ON "order_status_histories" ("created_at");
-- Create index "orderstatushistory_status" to table: "order_status_histories"
CREATE INDEX "orderstatushistory_status" ON "order_status_histories" ("status");
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "email" character varying NOT NULL, "password_hash" character varying NOT NULL, "name" character varying NOT NULL, "role" character varying NOT NULL, "phone" character varying NULL, "active" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "clinic_users" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "users_clinics_users" FOREIGN KEY ("clinic_users") REFERENCES "clinics" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "user_email" to table: "users"
CREATE INDEX "user_email" ON "users" ("email");
-- Create index "user_role" to table: "users"
CREATE INDEX "user_role" ON "users" ("role");
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "sessions" table
CREATE TABLE "sessions" ("id" uuid NOT NULL, "refresh_token" character varying NOT NULL, "ip_address" character varying NULL, "user_agent" character varying NULL, "device_id" character varying NULL, "number_of_uses" bigint NOT NULL DEFAULT 0, "expires_at" timestamptz NOT NULL, "last_used_at" timestamptz NOT NULL, "created_at" timestamptz NOT NULL, "revoked_at" timestamptz NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "sessions_users_sessions" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "sessions_refresh_token_key" to table: "sessions"
CREATE UNIQUE INDEX "sessions_refresh_token_key" ON "sessions" ("refresh_token");
