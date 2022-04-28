import pool from "../database/index.js";

class ErrorController {

  static async errorStore(req, res) {

    const { err, category } = req.body;

    console.log(err, category);

    const query = `INSERT INTO ${category} ( err ) VALUES ( $1 )`;

    try {
      await pool.query(query, [err]);
    } catch (e) {
        console.log(e);
        return res.sendStatus(400)
    }
    return res.sendStatus(200)
  }

}

export default ErrorController
