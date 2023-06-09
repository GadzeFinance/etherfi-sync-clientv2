# Go compiler
GO := go

# Output directory
OUTPUT_DIR := output

# Binary name
BINARY_NAME := etherfi-sync-clientv2

.PHONY: all clean

all: $(BINARY_NAME)

# Compile the Go code
$(BINARY_NAME):
	rm -f $(BINARY_NAME)
	rm -rf $(OUTPUT_DIR)/*
	rm -f data.db
	$(GO) build

# Clean up the binary and output directory
clean:
	rm -f $(BINARY_NAME)
	rm -rf $(OUTPUT_DIR)/*
	rm -f data.db

