import pg from "pg"


const { Pool } = pg
const { USER, POSTGRESQL_URI, POSTGRESQL_PASSWORD } = process.env

const credentials = {
  user: USER || "postgres",
  host: POSTGRESQL_URI || "localhost",
  database: "Guardian_Errors",
  password: POSTGRESQL_PASSWORD || "", // Your password
  port: 5432,
}

const pool = new Pool(credentials)

export default pool