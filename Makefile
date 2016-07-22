QUIET = @

help: FORCE
	@echo "Available targets:"
	@echo "  aligncheck  Run aligncheck across the source base."
	@echo "  build       Generate necessary source files and build. (default)"
	@echo "  install     Build and install the library."
	@echo "  lint        Run golint across the source base."

aligncheck: FORCE
	$(QUIET)find . -type d ! -ipath ./Demo/images ! -ipath "./.git*" -exec aligncheck \{\} \;

build: FORCE
	$(QUIET)go generate; go build -v

install: FORCE
	$(QUIET)go generate; go install -v

lint: FORCE
	$(QUIET)find . -type d ! -ipath ./Demo/images ! -ipath "./.git*" -exec golint \{\} \;

FORCE:
	