# Env var for access.
# example:
# export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
# export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
# export AWS_DEFAULT_REGION=us-west-2

provider "aws" {
    region = "us-west-2"
}

resource "aws_instance" "my_web" {
    ami = "ami-02f147dfb8be58a10"
    instance_type = "t3.micro"
    vpc_security_group_ids = [aws_security_group.my_web_sg.id]
    user_data = file("user_data.sh")
    key_name = "revolman"

    tags = {
        Name = "My Web Server"
        Owner = "Revolman"
    }
}

resource "aws_security_group" "my_web_sg" {
    name = "Web Security Group"
    description = "My First SG"

    ingress {
        from_port = 80
        to_port = 80
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    ingress {
        from_port = 443
        to_port = 443
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }
        
    ingress {
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks = ["0.0.0.0/0"]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
    tags = {
        Name = "Web server SG"
        Owner = "Revolman"
    }
}