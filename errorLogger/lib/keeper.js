import sendMail from "../config/nodemailer.js"
import fetch from "node-fetch"

const { AUTH_URI, FILES_URI } = process.env
let { AUTH_STATUS, FILES_STATUS } = process.env
const { EMAIL } = process.env


// ------ Main Function ---- //


const mainKeeper = () => {
    setInterval(authKeeper, 60000)
    setInterval(filesKeeper, 60000)
}


// ----- Auxiliary functions ---- //

const authKeeper = async () => {

    if (AUTH_STATUS === "crashed") return

    const filesURI = FILES_URI || "http://localhost:3002/imAlive"

    try {
        await fetch(filesURI)
        
    } catch (e) {
        AUTH_STATUS = "crashed"
        sendMail(EMAIL, `FileStorage`)
    }

} 

const filesKeeper = async () => {

    if (FILES_STATUS === "crashed") return

    const authURI = AUTH_URI || "http://localhost:3001/imAlive"

    try {
        await fetch(authURI)
    } catch (e) {
        FILES_STATUS = "crashed"
        sendMail(EMAIL, `Authentication`)
    }
} 

export default mainKeeper