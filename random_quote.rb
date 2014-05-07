require 'open-uri'
require 'json'
load 'bash-string-escaper/bash_string_escape.rb'

uri_string = "http://www.iheartquotes.com/api/v1/random?format=json&source=humorix_misc"
uri = URI.parse(uri_string)
begin
  json_result = uri.read

  parsed_result = JSON.parse(json_result)
  quote_text = parsed_result['quote']
  # puts quote_text
  # quote_text.tr!('\\n', '<br>')
  quote_text.gsub!(/\r\n|\r|\n/, "<br>")
  quote_text.gsub!(/\t/, "")
  # puts "==-===="

  puts "Original:"
  puts quote_text

  puts

  escaped_quote_text = bash_string_escape(quote_text, false)
  puts "Escaped:"
  puts escaped_quote_text
  File.open(File.join(ENV['HOME'], '.bash_profile'), 'a') { |f| f.write("export RANDOM_QUOTE=\"#{quote_text}\"\n") }
rescue => ex
  puts "Request error: #{ex}"
  puts " [i] uri_string: #{uri_string}"
  exit 1
end