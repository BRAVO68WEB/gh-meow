install:
	@echo "> Installing meow..."
	gh extension install .

remove:
	@echo "< Removing meow..."
	gh extension remove meow

reload: remove install