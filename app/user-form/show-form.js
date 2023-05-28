const jsonResponse = require('./json-response');
const errorResponse = require('./error-response');
const RequestProcessor = require('./request-processor');
const S3PolicySigner = require('./s3-policy-signer');
const aws = require('aws-sdk');
const s3 = new aws.S3();
const uploadLimitInMB = parseInt(process.env.UPLOAD_LIMIT_IN_MB);

exports.lambdaHandler = async (event, context) => {
  try {
    const uploadSigner = new S3PolicySigner(process.env.UPLOAD_S3_BUCKET);
    const downloadSigner = new S3PolicySigner(process.env.THUMBNAILS_S3_BUCKET);
    const requestProcessor = new RequestProcessor(uploadSigner, downloadSigner, uploadLimitInMB, process.env.ALLOWED_IMAGE_EXTENSIONS.split(','));
    const result = requestProcessor.processRequest(context.awsRequestId, event.pathParameters.extension,);
    return jsonResponse(result, process.env.CORS_ORIGIN);
  }
  catch (e) {
    return errorResponse(e, process.env.CORS_ORIGIN);
  }
};





