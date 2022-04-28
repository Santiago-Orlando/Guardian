import pg from "pg"


const { Pool } = pg
const { USER, HOST, DATABASE, PASSWORD } = process.env

const credentials = {
  user: USER || "postgres",
  host: HOST || "localhost",
  database: DATABASE || "Guardian_Errors",
  password: PASSWORD || "", // Your password
  port: 5432,
}

const pool = new Pool(credentials)

export default pool