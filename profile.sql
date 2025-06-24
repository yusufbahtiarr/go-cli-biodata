CREATE TYPE gender_type AS ENUM ('pria', 'wanita');

CREATE TABLE profiles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL CHECK (name <> ''),
  age INT NOT NULL CHECK (age > 0),  
  gender gender_type NOT NULL,
  created_at TIMESTAMP(0) DEFAULT NOW()
);

INSERT INTO profiles (name, age, gender) VALUES
  ('Putri Patricia', 30, 'wanita'),
  ('Joko Susilo', 28, 'pria'),
  ('Dewi Lestari', 22, 'wanita'),
  ('Agus Prabowo', 35, 'pria');
