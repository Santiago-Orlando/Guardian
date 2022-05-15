import pg from "pg"


const { Pool } = pg
const { POSTGRESQL_URI } = process.env

const credentials = {
  host: POSTGRESQL_URI,
  database: "Guardian_Errors",
  port: 3006,
}

const pool = new Pool(credentials)

export default pool