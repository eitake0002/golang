
# Description:
#   Run go file.
# Usage:
#   go-run <go_file>
# Example:
#   go-run main.go
function go-run()
{
  go run $1
}

# Description:
#   Build go file.
# Usage:
#   go-build <go_file>
# Example:
#   go-build main.go
function go-build()
{
  go build $1
}
