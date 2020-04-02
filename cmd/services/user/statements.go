package user

const INSERT_NEW_USER = `INSERT into trivia.user (name) VALUES($1) RETURNING id`

