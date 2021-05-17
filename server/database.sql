DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    role     VARCHAR(30) NOT NULL
);

INSERT INTO users(username, password,role)
    VALUES
        ('cfabrica46',  '01234',        'admin'),
        ('arturo',      '12345',        'admin'),
        ('carlos',      'abcd',         'member'),
        ('luis',        'lolsito123',   'member');
