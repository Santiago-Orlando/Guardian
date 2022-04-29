import express from "express"

import routes from "./routes/index.js"
import mainKeeper from "./lib/keeper.js"

const { PORT_ERROR_SERVICE } = process.env


const app = express()

app.use(express.json())

app.use("/", routes)

mainKeeper()


app.listen(PORT_ERROR_SERVICE, () => {
    console.log("Server start on port: " + PORT_ERROR_SERVICE)
})
