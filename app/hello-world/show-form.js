const htmlResponse = require('./html-response');
const formHtml = `
  <html>
  <head>
    <meta charset="utf-8"/>
  </head>
  <body>
    <form method="POST">
      Please enter your name:
      <input type="text" name="name"/>
      <br/>
      <input type="submit" />
    </form>
  </body>
  </html>
`;
exports.lambdaHandler = async (event, context) => {
  return htmlResponse(formHtml);
};
