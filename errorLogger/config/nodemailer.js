import nodemailer from "nodemailer";

const { GMAIL, GMAIL_PW } = process.env;

const createTrans = () => {
  if (!GMAIL || !GMAIL_PW) return;

  const transport = nodemailer.createTransport({
    host: "smtp.gmail.com",
    port: 465,
    secure: true,
    auth: {
      user: GMAIL,
      pass: GMAIL_PW,
    },
  });
  return transport;
};

const sendMail = (email, service) => {
  if (!email) return;

  const transporter = createTrans();
  const info = transporter.sendMail({
    from: GMAIL,
    to: email,
    subject: "Critical Server Crash!",
    html: `The service: ${service}, is not responding!`,
  });
};

export default sendMail;
