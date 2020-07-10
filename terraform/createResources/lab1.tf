# Env var for access.
# example:
# export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
# export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
# export AWS_DEFAULT_REGION=us-west-2

provider "aws" {
    region = "us-west-2"
}

resource "aws_instance" "my_Ubuntu" {
    count = 2
    ami = "ami-003634241a8fcdec0"
    instance_type = "t3.micro"
}

resource "aws_instance" "my_Amazon" {
    ami = "ami-0b1e2eeb33ce3d66f"
    instance_type = "t3.micro"
}