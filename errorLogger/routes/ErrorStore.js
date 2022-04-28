import express from "express";

import ErrorController from "../controllers/ErrorController.js";

const route = express.Router()


route.post("/", ErrorController.errorStore)


export default route