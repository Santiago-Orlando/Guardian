import pg from "pg"


const { Pool } = pg
const { PASSWORD } = process.env

const credentials = {
  host: "ERROR_LOGGER_DB_URL",
  database: "Guardian_Errors",
  port: 5432,
  user: "postgres",
  password: PASSWORD
}

const pool = new Pool(credentials)

export default pool