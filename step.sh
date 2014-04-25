the_quote=$(curl -s http://www.iheartquotes.com/api/v1/random)
echo $the_quote
echo "RANDOM_JOKE=\"$the_quote\"" >> $HOME/.bash_profile