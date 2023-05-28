const jsonResponse = require('./json-response');
const buildForm = require('./build-form');
const aws = require('aws-sdk');
const s3 = new aws.S3();
const uploadLimitInMB = parseInt(process.env.UPLOAD_LIMIT_IN_MB);

exports.lambdaHandler = async (event, context) => {
    const key = context.awsRequestId + '.jpg',
    uploadParams = {
        Bucket: process.env.UPLOAD_S3_BUCKET,
        Expires: 600,
        Conditions: [
            ['content-length-range', 1, uploadLimitInMB * 1000000]
          ],
        Fields: {
            acl: 'private',
            key: key
        }
    },
    uploadForm = s3.createPresignedPost(uploadParams),
    downloadParams = {
        Bucket: process.env.THUMBNAILS_S3_BUCKET,
        Key: key,
        expires: 600
    },
    downloadUrl = s3.getSignedUrl('getObject', downloadParams);



  return jsonResponse({
    upload: uploadForm,
    download: downloadUrl
  });
};
