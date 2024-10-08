CREATE TABLE IF NOT EXISTS salary (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    currency VARCHAR(10) NOT NULL,
    wage DECIMAL(15, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS employee (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    Name VARCHAR(64) NOT NULL,
    salary_id BIGINT UNSIGNED,
    country VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (salary_id) REFERENCES salary(id) ON DELETE SET NULL
);