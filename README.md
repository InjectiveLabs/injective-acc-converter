# Injective Address Converter

This code is a simple command-line tool that converts an Injective address to a subaccount ID. It is written in Go and uses the Cosmos SDK and Ethereum libraries.

## Installation

To use this code, you will need to have Go installed on your machine. You can then download and compile the code using the following commands:

```bash
git clone https://github.com/InjectiveLabs/injective-acc-converter.git
cd injective-acc-converter
go run main.go --address=<account_address>
```

## Usage

To convert an Injective address to a subaccount ID, run the following command:

```bash
./injective-acc-converter --address=<account_address>
```

Replace <account_address> with the Injective chain account address that you want to convert. For example:

```go
./injective-acc-converter --address=inj1hkhdaj2a2clmq5jq6mspsggqs32vynpk228q3r
```

This will output the hexadecimal representation of the first 20 bytes of the account address, padded with zeros to 24 digits.

You ca