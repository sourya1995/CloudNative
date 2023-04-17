
package com.google.training.appdev.services.gcp.pubsub;

import com.fasterxml.jackson.databind.ObjectMapper;

import com.google.cloud.ServiceOptions;
import com.google.api.core.ApiFuture;
import com.google.api.core.ApiFutures;
import com.google.cloud.pubsub.v1.Publisher;
import com.google.cloud.pubsub.v1.TopicAdminClient;
import com.google.protobuf.ByteString;
import com.google.pubsub.v1.PubsubMessage;
import com.google.pubsub.v1.TopicName;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import com.google.training.appdev.services.gcp.domain.Feedback;

import org.springframework.stereotype.Service;

import java.io.IOException;


@Service
public class PublishService {
    private static final String PROJECT_ID = ServiceOptions.getDefaultProjectId();
    private static final String TOPIC_NAME = "feedback";

    public void publishFeedback(Feedback feedback) throws Exception {
        
        ObjectMapper mapper = new ObjectMapper();
        String feedbackMessage = mapper.writeValueAsString(feedback);

        TopicName topicName = TopicName.create(PROJECT_ID, TOPIC_NAME);
        Publisher publisher = null;
        ApiFuture<String> messageIdFuture = null;
        try {

            publisher = Publisher.defaultBuilder(topicName).build();
            
            ByteString data = ByteString.copyFromUtf8(feedbackMessage); //convert to bytestring
            PubsubMessage pubsubMessage = PubsubMessage.newBuilder().setData(data).build();

            messageIdFuture = publisher.publish(pubsubMessage); //get a future
        
        } finally {

            String messageId = messageIdFuture.get(); //extract message

            System.out.println("published with message ID: " + messageId);

            if (publisher != null) {
                // When finished with the publisher, shutdown to free up resources.
                publisher.shutdown();
            }
        }

        
    }

}