import nodemailer from "nodemailer"

const { GMAIL, GMAIL_PASSWORD } = process.env;

const createTrans = () => {
  const transport = nodemailer.createTransport({
    host: "smtp.gmail.com",
    port: 465,
    secure: true,
    auth: {
      user: GMAIL,
      pass: GMAIL_PASSWORD,
    },
  });
  return transport;
};

const sendMail = (email, service) => {
  const transporter = createTrans();
  const info = transporter.sendMail({
    from: GMAIL,
    to: email,
    subject: "Critical Server Crash!",
    html: `The service: ${service}, is not responding!`,
  });
};

export default sendMail