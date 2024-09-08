CREATE TYPE Roles_t as ENUM ('doctor','admin','user');
CREATE TYPE gender_t as ENUM ('male', 'female', 'other');

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    Created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE category(
    id BIGSERIAL PRIMARY KEY,
    category_name TEXT UNIQUE NOT NULL,
    Created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE doctors (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL REFERENCES category(id),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    phone TEXT NOT NULL,
    password TEXT NOT NULL,
    gender gender_t NOT NULL DEFAULT 'male',
    dob TEXT NOT NULL,
    img TEXT,
    shedule TEXT NOT NULL,
    degree TEXT NOT NULL,
    address TEXT NOT NULL,
    Roles Roles_t NOT NULL DEFAULT 'doctor',
    Created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE appointments (
    id BIGSERIAL PRIMARY KEY,
    users_id BIGINT NOT NULL REFERENCES users(id),
    doctor_id BIGINT NOT NULL REFERENCES doctors(id),
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    gender gender_t NOT NULL DEFAULT 'male',
    age TEXT NOT NULL,
    address TEXT NOT NULL,
    date TEXT NOT NULL,
    time TEXT  NOT NULL,
    Created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);    
