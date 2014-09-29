require 'open-uri'
require 'json'
load 'bash-string-escaper/bash_string_escape.rb'

$formatted_output_file_path = ENV['BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH']

def puts_string_to_formatted_output(text)
  open($formatted_output_file_path, 'a') { |f|
    f.puts(text)
  }
end

def puts_section_to_formatted_output(section_text)
  open($formatted_output_file_path, 'a') { |f|
    f.puts
    f.puts(section_text)
    f.puts
  }
end

puts_section_to_formatted_output('# Random Quote')


uri_string = "http://www.iheartquotes.com/api/v1/random?format=json&source=humorix_misc"
uri = URI.parse(uri_string)
begin
  json_result = uri.read

  parsed_result = JSON.parse(json_result)
  raw_quote_test = parsed_result['quote']

  md_prepared_quote = ""
  raw_quote_test.each_line {|line| md_prepared_quote << "    #{line}"}
  puts_section_to_formatted_output(md_prepared_quote)

  quote_text = raw_quote_test

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
  File.open(File.join(ENV['HOME'], '.bash_profile'), 'a') { |f| f.write("export RANDOM_QUOTE=\"#{escaped_quote_text}\"\n") }
rescue => ex
  puts "Request error: #{ex}"
  err_msg = "Exception happened: #{ex}"
  puts err_msg
  puts " [i] uri_string: #{uri_string}"
  puts_section_to_formatted_output("## Failed")
  unless err_msg.nil?
    puts_section_to_formatted_output(err_msg)
  end
  puts_section_to_formatted_output("Check the Logs for more details.")
end
