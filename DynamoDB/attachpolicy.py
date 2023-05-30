import boto3

# Create an IAM client
iam_client = boto3.client('iam')

# Specify the policy ARN and group name
policy_arn = 'arn:aws:iam::123456789012:policy/MyPolicy'
group_name = 'MyUserGroup'

# Attach the policy to the group
response = iam_client.attach_group_policy(
    PolicyArn=policy_arn,
    GroupName=group_name
)

# Check if the policy attachment was successful
if response['ResponseMetadata']['HTTPStatusCode'] == 200:
    print(f"Policy '{policy_arn}' attached to group '{group_name}' successfully.")
else:
    print("Failed to attach policy to group.")
