import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.RequestHandler;
import com.amazonaws.services.lambda.runtime.events.S3Event;
import com.amazonaws.services.sqs.AmazonSQS;
import com.amazonaws.services.sqs.AmazonSQSClientBuilder;
import com.amazonaws.services.sqs.model.SendMessageRequest;

public class FileUploadLambda implements RequestHandler<S3Event, String> {

    private static final String SQS_QUEUE_URL = "your-sqs-queue-url";

    @Override
    public String handleRequest(S3Event s3Event, Context context) {
        for (S3Event.S3EventRecord record : s3Event.getRecords()) {
            String bucketName = record.getS3().getBucket().getName();
            String objectKey = record.getS3().getObject().getKey();
            String message = "File uploaded: " + bucketName + "/" + objectKey;

            sendSQSMessage(message);
        }

        return "File upload events sent to SQS";
    }

    private void sendSQSMessage(String message) {
        AmazonSQS sqsClient = AmazonSQSClientBuilder.defaultClient();
        SendMessageRequest request = new SendMessageRequest()
                .withQueueUrl(SQS_QUEUE_URL)
                .withMessageBody(message);

        sqsClient.sendMessage(request);
    }
}
