package session

const INSERT_NEW_SESSION = `INSERT into trivia.session (code, name, description, start) VALUES($1,$2,$3,$4) RETURNING id`
const COUNT_OF_CODE = `SELECT COUNT(*) as count FROM trivia.session WHERE code=$1`
