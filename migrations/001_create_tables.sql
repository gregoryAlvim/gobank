-- Migration for natural_person table
CREATE TABLE natural_person (
    id SERIAL PRIMARY KEY,
    monthly_income DECIMAL,
    age INT,
    full_name VARCHAR(255),
    phone_number VARCHAR(20),
    email VARCHAR(255),
    category VARCHAR(50),
    balance DECIMAL
);

-- Migration for legal_person table
CREATE TABLE legal_person (
    id SERIAL PRIMARY KEY,
    annual_revenue DECIMAL,
    age INT,
    trade_name VARCHAR(255),
    phone_number VARCHAR(20),
    corporate_email VARCHAR(255),
    category VARCHAR(50),
    balance DECIMAL
);
