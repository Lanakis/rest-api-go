CREATE TABLE IF NOT EXISTS users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL,
                       password VARCHAR(100) NOT NULL,
                       role VARCHAR(20) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profiles (
                                        id SERIAL PRIMARY KEY,
                                        first_name VARCHAR(20) NOT NULL,
                                        middle_name VARCHAR(20),
                                        last_name VARCHAR(20),
                                        age INT,
                                        head BOOLEAN,
                                        user_id INT REFERENCES users(id) ON DELETE CASCADE,
                                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DO $$
    BEGIN
        IF NOT EXISTS (
                SELECT constraint_name
                FROM information_schema.table_constraints
                WHERE table_name = 'profiles'
                  AND constraint_name = 'fk_user_profile'
                  AND constraint_type = 'FOREIGN KEY'
            ) THEN
            ALTER TABLE profiles
                ADD CONSTRAINT fk_user_profile
                    FOREIGN KEY (user_id) REFERENCES users(id)
                        ON DELETE CASCADE;
        END IF;
    END $$;


INSERT INTO users (username, password, role)
SELECT 'admin', '123456', 'admin'
WHERE NOT EXISTS (
        SELECT 1 FROM users WHERE username = 'admin'
    );

INSERT INTO profiles (first_name, middle_name, last_name, age, head, user_id)
SELECT 'John', NULL, 'Doe', 30, true, 1
WHERE NOT EXISTS (
        SELECT 1 FROM profiles WHERE user_id = 1
    );

INSERT INTO users (username, password, role)
SELECT 'user', '123456', 'user'
WHERE NOT EXISTS (
        SELECT 2 FROM users WHERE username = 'user'
    );

INSERT INTO profiles (first_name, middle_name, last_name, age, head, user_id)
SELECT 'Jack', NULL, 'Couch', 15, true, 2
WHERE NOT EXISTS (
        SELECT 2 FROM profiles WHERE user_id = 2
    );