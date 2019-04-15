workflow "Assign moreal to all issues" {
  resolves = ["Assign issue", "Label issue"]
  on = "issues"
}

action "Assign issue" {
  uses = "actions/github@v1.0.0"
  args = "assign moreal --action=opened"
  secrets = ["GITHUB_TOKEN"]
}

action "Label issue" {
  uses = "actions/github@v1.0.0"
  needs = ["Assign issue"]
  args = "label todo --action=opened"
  secrets = ["GITHUB_TOKEN"]
}