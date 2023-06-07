exports.handler = async (event) => {
    let policy = {
      principalId: "user",
      policyDocument: {
        Version: "2012-10-17",
        Statement: [],
      },
    };
  
    if (event.authorizationToken === event.authorizationToken) {
      policy.policyDocument.Statement.push({
        Action: "execute-api:Invoke",
        Effect: "Allow",
        Resource: event.methodArn,
      });
    }
    return policy;
  };