import pool from "../database/index.js";

class ErrorController {
  static async errorStore(req, res) {
    const { err, category } = req.body;

    const query = `INSERT INTO ${category} ( err ) VALUES ( $1 )`;

    try {
      await pool.query(query, [err]);
      return res.sendStatus(200);
    } catch (e) {
      console.log("catch");
      return res.sendStatus(400);
    }
  }
}

export default ErrorController;
