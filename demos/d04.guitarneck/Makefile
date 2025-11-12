include ../demos.mk

pdf: README.pdf
%.pdf: %.md
	@pandoc $< -o $@

clean::
	@rm -f README.pdf