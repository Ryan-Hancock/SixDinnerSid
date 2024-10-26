-- migrations/000001_init_schema.up.sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cats (
                      id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                      name VARCHAR(255) NOT NULL,
                      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE meals (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       cat_id UUID NOT NULL REFERENCES cats(id),
                       fed_by_id UUID NOT NULL REFERENCES users(id),
                       timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
                       food_types TEXT[] NOT NULL,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_meals_timestamp ON meals(timestamp);
CREATE INDEX idx_meals_cat_id ON meals(cat_id);
CREATE INDEX idx_meals_fed_by_id ON meals(fed_by_id);

-- migrations/000001_init_schema.down.sql
DROP TABLE IF EXISTS meals;
DROP TABLE IF EXISTS cats;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";