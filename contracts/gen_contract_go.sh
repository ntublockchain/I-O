#!/bin/bash
solc-select use 0.8.18
# Check if the folder name is provided
if [ -z "$1" ]; then
  echo "Please provide a folder name."
  exit 1
fi

# Navigate to the given folder
cd "$1" || exit

# Find the Solidity file
SOL_FILE=$(ls *.sol)
if [ "$(echo "$SOL_FILE" | wc -l)" -ne 1 ]; then
  echo "There must be exactly one Solidity file in the folder."
  exit 1
fi

# Extract contract name from the Solidity file
# Adjusted to use grep -E for extended regex, which is more widely supported
CONTRACT_NAME=$(grep -Eo 'contract[[:space:]]+[a-zA-Z_][a-zA-Z0-9_]*' "$SOL_FILE" | awk '{print $2}')

if [ -z "$CONTRACT_NAME" ]; then
  echo "Could not extract contract name from the Solidity file."
  exit 1
fi


# Compile the contract
solc --bin --abi --optimize -o build "$SOL_FILE" --overwrite

# Generate Go file using abigen
BIN_FILE="./build/${CONTRACT_NAME}.bin"
ABI_FILE="./build/${CONTRACT_NAME}.abi"
PKG_NAME=$(basename "$PWD")
OUT_FILE="${CONTRACT_NAME}_gen.go"

# Use the path to abigen that works for your setup; it might require an absolute path
../abigen --bin="$BIN_FILE" --abi="$ABI_FILE" --pkg="$PKG_NAME" --out="$OUT_FILE"

echo "Generation completed: $OUT_FILE"
