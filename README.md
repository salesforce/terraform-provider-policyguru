# terraform-provider-policy-sentry

# Run locally

go mod init terraform-provider-policy-sentry
go mod vendor
make build

# Test provider
make install
cd examples
terraform init
terraform apply

