import express from "express";

import errorRoutes from "./ErrorStore.js"

const route = express.Router()

route.use("/errors", errorRoutes)

export default route