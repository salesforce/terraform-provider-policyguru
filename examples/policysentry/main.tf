data "policysentry_document" "example" {}

# Returns policy sentry document in json
output "policysentry_document_json" {
  value = data.policysentry_document.example.json
}