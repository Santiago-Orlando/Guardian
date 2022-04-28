import express from "express"

import routes from "./routes/index.js"
import mainKeeper from "./lib/keeper.js"

//const { PORT } = process.env


const app = express()

app.use(express.json())

app.use("/", routes)

//mainKeeper()

const PORT = 3003

app.listen(PORT, () => {
    console.log("Server start on port: " + PORT)
})
