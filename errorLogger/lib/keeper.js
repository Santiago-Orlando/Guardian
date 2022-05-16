import sendMail from "../config/nodemailer.js";
import fetch from "node-fetch";

const {
  PORT_AUTHENTICATION_SERVICE,
  PORT_FILESTORAGE_SERVICE,
} = process.env;
let { AUTH_STATUS, FILES_STATUS } = process.env;
const { GMAIL, GMAIL_PW, EMAIL } = process.env;

// ------ Main Function ---- //

const mainKeeper = () => {
  if (!GMAIL || !GMAIL_PW || !EMAIL) return;

  setInterval(authKeeper, 60000);
  setInterval(filesKeeper, 60000);
};

// ----- Auxiliary functions ---- //

const authKeeper = async () => {
  if (AUTH_STATUS === "crashed") return;

  const authURI = "AUTHENTICATION_URL" + PORT_AUTHENTICATION_SERVICE;

  try {
    await fetch(authURI);
  } catch (e) {
    AUTH_STATUS = "crashed";
    sendMail(EMAIL, `FileStorage`);
  }
};

const filesKeeper = async () => {
  if (FILES_STATUS === "crashed") return;

  const filesURI = "FILESTORAGE_URL" + PORT_FILESTORAGE_SERVICE;

  try {
    await fetch(filesURI);
  } catch (e) {
    FILES_STATUS = "crashed";
    sendMail(EMAIL, `Authentication`);
  }
};

export default mainKeeper;
