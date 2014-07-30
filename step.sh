#!/bin/bash

formatted_output_file_path="$BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH"

ruby ./random_quote.rb >> "$formatted_output_file_path"

exit $?