.DEFAULT:all

ethspider:
	cd cmd/ethspider; go build
	@echo "Done building."

all:ethspider
	@echo "Done All." 


.PHONY: clean 
clean:
	cd cmd/ethspider; go clean -cache
