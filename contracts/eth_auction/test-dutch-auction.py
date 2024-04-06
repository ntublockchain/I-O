from solcx import compile_standard, install_solc
import json
import os
from web3 import Web3
import time
# from eth_account.messages import encode_defunct
import numpy as np
# from numpy.linalg import eig

# mnemonic: top bus drama shoulder build master apart arrange notice fancy truth ice
# 16 accounts

private_keys = [
    "0d9fb7ddd8a2e3c19b9c57166f6dce352c5c021fe79e6dd9654b3adfbfbaad26",
    "74e853dc72b18efb576929ccc84903c0d5044ee2445661e884033bf69eec2638",
    "d14fa190b8cbd656fa4f8fe9acbe531cc19ed8b80a39f1e40205816885f80753",
    "05f6739e9a66796ddb57310f2142b0534d10937a22475f744994ffc76953ce70",
    "9620e18cde7eae53c810ec7a80fd49e5b418dba6870958ed4f2ed15f2479022f",
    "57bee8409f420f3e8e7f1993b64f27fd6bc03368db749956550932304502a884",
    "0b087981241722083ebf113ef6d062f0b8af82c593cb92e3b3612ee54a2cece9",
    "dd45d7f2ce48c6aa6a61989eeac6ee744f371054ae89a1dc4982fe12a0d222ab",
    "1ae07aee2dede3a67e0f2e193e49481819d8701d9830e19bb57812c831ba963f",
    "f148c3fbdc45fa2554a97e7498b6112ab04ceeef87debb61b6624ce7fdee2e05",
    "d1ee0eadf33442ed9c61b6da201975d473dc0e019c172024caf6945040f78792",
    "a0fc20c4cb3b197716e42e4062a8a8e01cc5dd6759fe728fc2c5880a3573e494",
    "e000975a7ba4b2c010d0e98850632b1510cd7858a2db0ef9019370cce7cff5ce",
    "3855a3046094b021b0710e1d9ae3987f82d91837ba326c698ac172d57a5c8e83",
    "7a169b606afaccdccc397e49a44d56a68a2cfbb7f8f544376658554409121728",
    "249705240d48de99e7272441057c93559fe0663538f7385a0581a1936310a939",
]

start_time = time.time()

__location__ = os.path.realpath(
    os.path.join(os.getcwd(), os.path.dirname(__file__)))

install_solc("0.8.18")

w3 = Web3(Web3.HTTPProvider("http://127.0.0.1:7545"))
chain_id = 1337
node_address = w3.eth.accounts[0]
node_private_key = private_keys[0]

# first the coin
with open(os.path.join(__location__, "Coin.sol"), "r") as file:
    coin_contract_file = file.read()

compiled_coin_sol = compile_standard(
    {
        "language": "Solidity",
        "sources": {"Coin.sol": {"content": coin_contract_file}},
        "settings": {
            "outputSelection": {
                "*": {
                    "*": [
                        "abi",
                        "metadata",
                        "evm.bytecode",
                        "evm.bytecode.sourceMap",
                    ]  # output needed to interact with and deploy contract
                }
            }
        },
    },
    solc_version="0.8.18",
)

with open(os.path.join(__location__, "compiler_output.json"), "w") as file:
    json.dump(compiled_coin_sol, file)

## get bytecode
coin_bytecode = compiled_coin_sol["contracts"]["Coin.sol"]["Coin"]["evm"][
    "bytecode"
]["object"]

# ## get abi
coin_abi = json.loads(
    compiled_coin_sol["contracts"]["Coin.sol"]["Coin"]["metadata"]
)["output"]["abi"]

# create the contract in Python
coin_contract = w3.eth.contract(abi=coin_abi, bytecode=coin_bytecode)
# get the latest transaction
coin_nonce = w3.eth.get_transaction_count(node_address)

# create a transaction that deploys the contract
deploy_coin_transaction = coin_contract.constructor(0).build_transaction(
    {"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": coin_nonce}
)

# sign the transaction
signed_deploy_coin_transaction = w3.eth.account.sign_transaction(deploy_coin_transaction, private_key=node_private_key)
print(f"Start Deploying Coin Contract!")
# send the transaction
deploy_coin_transaction_hash = w3.eth.send_raw_transaction(signed_deploy_coin_transaction.rawTransaction)
# wait for the transaction to be mined, and get the transaction receipt
print("Waiting for transaction to finish...")
deploy_coin_transaction_receipt = w3.eth.wait_for_transaction_receipt(deploy_coin_transaction_hash)
coin_address = deploy_coin_transaction_receipt.contractAddress
print(f"Done! Contract Coin deployed to {coin_address}")

deployed_coin_contract = w3.eth.contract(coin_address, abi=coin_abi)

# see then the auction
with open(os.path.join(__location__, "DutchAuction.sol"), "r") as file:
    auction_contract_file = file.read()

compiled_auction_sol = compile_standard(
    {
        "language": "Solidity",
        "sources": {"DutchAuction.sol": {"content": auction_contract_file}},
        "settings": {
            "outputSelection": {
                "*": {
                    "*": [
                        "abi",
                        "metadata",
                        "evm.bytecode",
                        "evm.bytecode.sourceMap",
                    ]  # output needed to interact with and deploy contract
                }
            }
        },
    },
    solc_version="0.8.18",
)

with open(os.path.join(__location__, "compiler_output.json"), "w") as file:
    json.dump(compiled_auction_sol, file)

## get bytecode
auction_bytecode = compiled_auction_sol["contracts"]["DutchAuction.sol"]["DutchAuction"]["evm"][
    "bytecode"
]["object"]

# ## get abi
auction_abi = json.loads(
    compiled_auction_sol["contracts"]["DutchAuction.sol"]["DutchAuction"]["metadata"]
)["output"]["abi"]

# create the contract in Python
auction_contract = w3.eth.contract(abi=auction_abi, bytecode=auction_bytecode)
# get the latest transaction
auction_nonce = w3.eth.get_transaction_count(node_address)

# create a transaction that deploys the contract
deploy_auction_transaction = auction_contract.constructor(coin_address).build_transaction(
    {"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": auction_nonce}
)

# sign the transaction
signed_deploy_auction_transaction = w3.eth.account.sign_transaction(deploy_auction_transaction, private_key=node_private_key)
print(f"Start Deploying DutchAuction Contract!")
# send the transaction
deploy_auction_transaction_hash = w3.eth.send_raw_transaction(signed_deploy_auction_transaction.rawTransaction)
# wait for the transaction to be mined, and get the transaction receipt
print("Waiting for transaction to finish...")
deploy_auction_transaction_receipt = w3.eth.wait_for_transaction_receipt(deploy_auction_transaction_hash)
auction_address = deploy_auction_transaction_receipt.contractAddress
print(f"Done! Contract DutchAuction deployed to {auction_address}")

deployed_auction_contract = w3.eth.contract(auction_address, abi=auction_abi)

## scenario ##

# mint coins

print("Mint 1000 coins for account 0... ")
add_request_nonce = w3.eth.get_transaction_count(node_address)
tx_hash = deployed_coin_contract.functions.mint(w3.eth.accounts[0], 1000).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

print("Mint 1000 coins for account 1... ")
add_request_nonce = w3.eth.get_transaction_count(node_address)
tx_hash = deployed_coin_contract.functions.mint(w3.eth.accounts[1], 1000).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

print("Mint 1000 coins for account 2... ")
add_request_nonce = w3.eth.get_transaction_count(node_address)
tx_hash = deployed_coin_contract.functions.mint(w3.eth.accounts[2], 1000).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

# create auctions

print("Create auction... ")
add_request_nonce = w3.eth.get_transaction_count(node_address)
tx_hash = deployed_auction_contract.functions.create("asset 1", 120, 10).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

print("Create another auction... ")
add_request_nonce = w3.eth.get_transaction_count(node_address)
tx_hash = deployed_auction_contract.functions.create("asset 2", 200, 10).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

# Bid 1 

print("Allow 100 coins from account 2 to be sent... ")
add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
tx_hash = deployed_coin_contract.functions.approve(auction_address, 100).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

try:
    print("User 2 issues a bid for the first auction... ")
    add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
    tx_hash = deployed_auction_contract.functions.bid(0, 100).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    print(f"Done! Gas use: {receipt.gasUsed}")
except:
    print("bid failed (correct)")

price = deployed_auction_contract.functions.getPrice(0).call()  # read the coin total supply - call means we are reading from the blockchain
print(price) 

try:
    print("User 2 issues a bid for the first auction... ")
    add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
    tx_hash = deployed_auction_contract.functions.bid(0, 100).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
    receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
    print(f"Done! Gas use: {receipt.gasUsed}")
except:
    print("bid failed (correct)")

time.sleep(1)

price = deployed_auction_contract.functions.getPrice(0).call()  # read the coin total supply - call means we are reading from the blockchain
print(f"Current price: {price}")

print("Create yet another auction (to increase block height)... ")
add_request_nonce = w3.eth.get_transaction_count(node_address)
tx_hash = deployed_auction_contract.functions.create("asset 2", 200, 10).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")
price = deployed_auction_contract.functions.getPrice(0).call()  # read the coin total supply - call means we are reading from the blockchain
print(f"Current price: {price}")

time.sleep(1)

print("Create yet another auction (to increase block height)... ")
add_request_nonce = w3.eth.get_transaction_count(node_address)
tx_hash = deployed_auction_contract.functions.create("asset 2", 200, 10).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": node_address, "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")
price = deployed_auction_contract.functions.getPrice(0).call()  # read the coin total supply - call means we are reading from the blockchain
print(f"Current price: {price}")

time.sleep(1)

print("User 2 issues a second bid for the first auction... ")
add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
tx_hash = deployed_auction_contract.functions.bid(0, 100).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

print("Close auction 1... ")
add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[0])
tx_hash = deployed_auction_contract.functions.closeAuction(0, False).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[0], "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

print("Winner commits auction 1... ")
add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
tx_hash = deployed_auction_contract.functions.commit(0, "").transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

print("Auctioneer withdraws winning bid... ")
add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[0])
tx_hash = deployed_auction_contract.functions.withdraw(0).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[0], "nonce": add_request_nonce})
receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
print(f"Done! Gas use: {receipt.gasUsed}")

# # Bid 2

# print("Allow 120 coins from account 1 to be sent... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[1])
# tx_hash = deployed_coin_contract.functions.approve(auction_address, 120).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[1], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")

# print("User 1 issues a higher bid for the first auction... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[1])
# tx_hash = deployed_auction_contract.functions.bid(0, 120).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[1], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")

# # Bid 3

# print("Allow 80 coins from account 2 to be sent... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
# tx_hash = deployed_coin_contract.functions.approve(auction_address, 80).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")

# print("User 2 issues a bid for the first auction... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
# tx_hash = deployed_auction_contract.functions.bid(1, 80).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")

# # Close auction

# print("Close auction 1... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[0])
# tx_hash = deployed_auction_contract.functions.closeAuction(0, False).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[0], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")

# print("Winner commits auction 1... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[1])
# tx_hash = deployed_auction_contract.functions.commit(0, "").transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[1], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")

# print("First bidder gets refund... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[2])
# tx_hash = deployed_auction_contract.functions.withdraw(0).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[2], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")

# print("Auctioneer withdraws winning bid... ")
# add_request_nonce = w3.eth.get_transaction_count(w3.eth.accounts[0])
# tx_hash = deployed_auction_contract.functions.withdraw(0).transact({"chainId": chain_id, "gasPrice": w3.eth.gas_price, "from": w3.eth.accounts[0], "nonce": add_request_nonce})
# receipt = w3.eth.wait_for_transaction_receipt(tx_hash)
# print(f"Done! Gas use: {receipt.gasUsed}")