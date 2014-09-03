# A sample Guardfile
# More info at https://github.com/guard/guard#readme

# Add files and commands to this file, like the example:
#   watch(%r{file/path}) { `command(s)` }
#
guard :shell do
  watch(/(.*).go/) {|m| `go test -v ./...` }
end

notification :tmux, {
  display_message: true,
  timeout: 3, # in seconds
  default_message_format: "%s >> %s",
  default: "default",
  success: "default",
  failed: "colour1",
  # the first %s will show the title, the second the message
  # Alternately you can also configure *success_message_format*,
  # *pending_message_format*, *failed_message_format*
  line_separator: " > " # since we are single line we need a separator
}
