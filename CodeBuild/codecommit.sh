# creating a codecommit repository
aws codecommit create-repository \
    --repository-name name \
    --repository-description "Description of what the repo does" \
    --tags "key=value"

# Adding files to a repository/branch
aws codecommit put-file \
    --repository-name name \
    --branch-name name \
    --file-content file:///usercode/new.sh \
    --file-path /new.sh \
    --parent-commit-id "commit-id" \ # Obtain a commit ID from the output of get-branch command
    --name "name" \
    --email "email" \
    --commit-message "A descriptive commit message"

# Deleting a repository
aws codecommit delete-repository \
    --repository-name name

    