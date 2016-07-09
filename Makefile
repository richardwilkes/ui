def:
	@echo "Available targets:"
	@echo "  aligncheck  Run aligncheck across the source base."
	@echo "  lint        Run golint across the source base."

aligncheck:
	@find . -type d ! -ipath ./Demo/images ! -ipath "./.git*" -exec aligncheck \{\} \;

lint:
	@find . -type d ! -ipath ./Demo/images ! -ipath "./.git*" -exec golint \{\} \;
