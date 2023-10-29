-- Add a unique constraint to the 'email' column in the 'users' table
ALTER TABLE "public"."users"
ADD CONSTRAINT unique_email UNIQUE (email);