package config

const username = "root"
const password = "runner"
const host = "db"
const port = "3306"
const dbName = "run"
const charset = "utf8"

const ConnectionString = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=" + charset + "&parseTime=true"

// root:workout@tcp(db:3306)/workout
const DriverName = "mysql"
