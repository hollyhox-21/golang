#!/bin/sh

echo ">> Waiting for $MYSQL_SERVER to start"
while ! `nc -z $MYSQL_SERVER 3306`; do sleep 3; done
echo ">> $MYSQL_SERVER has started"

mysql --host=$MYSQL_SERVER --user=$MYSQL_USER --password=$MYSQL_PASSWORD --database=$MYSQL_DATABASE<<EOFMYSQL
drop table if exists snippets ;
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);
CREATE INDEX idx_snippets_created ON snippets(created);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Не имей сто рублей',
    'Не имей сто рублей,\nа имей сто друзей.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
 
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Лучше один раз увидеть',
    'Лучше один раз увидеть,\nчем сто раз услышать.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
 
INSERT INTO snippets (title, content, created, expires) VALUES (
    'Не откладывай на завтра',
    'Не откладывай на завтра,\nчто можешь сделать сегодня.',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);

EOFMYSQL

# sh