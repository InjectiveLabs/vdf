import go_wrapper as sloth


# same parameters for both Sloth_elapsed_time and Sloth_fixed_delay:
#  ( p , x , t ) all String
# p: security parameter, how many bits prime. The actual prime number is stored in go_src/vdf_interface.go
# x: starting value, or input value
#
# for Sloth_fixed_delay:
# t: the number of modular square root iterations
#
# for Sloth_elapsed_time:
# t: the number of iterations every loop (within an infinite loop)

p=sloth.go_wrapper()

y=p.Sloth_eval("1024" , "8239479" , "9")
print("Ending Value: ",y)
result= p.Sloth_verify("1024","8239479","9",y)
# keep in mind that result is a string
print("verified: ",result)
# ZRX ORDER SCHEMA
# {
#     "id": "/orderSchema",
#     "properties": {
#         "makerAddress": { "$ref": "/addressSchema" },
#         "takerAddress": { "$ref": "/addressSchema" },
#         "makerFee": { "$ref": "/wholeNumberSchema" },
#         "takerFee": { "$ref": "/wholeNumberSchema" },
#         "senderAddress": { "$ref": "/addressSchema" },
#         "makerAssetAmount": { "$ref": "/wholeNumberSchema" },
#         "takerAssetAmount": { "$ref": "/wholeNumberSchema" },
#         "makerAssetData": { "$ref": "/hexSchema" },
#         "takerAssetData": { "$ref": "/hexSchema" },
#         "salt": { "$ref": "/wholeNumberSchema" },
#         "exchangeAddress": { "$ref": "/addressSchema" },
#         "feeRecipientAddress": { "$ref": "/addressSchema" },
#         "expirationTimeSeconds": { "$ref": "/wholeNumberSchema" }
#     },
#     "required": [
#         "makerAddress",
#         "takerAddress",
#         "makerFee",
#         "takerFee",
#         "senderAddress",
#         "makerAssetAmount",
#         "takerAssetAmount",
#         "makerAssetData",
#         "takerAssetData",
#         "salt",
#         "exchangeAddress",
#         "feeRecipientAddress",
#         "expirationTimeSeconds"
#     ],
#     "type": "object"
# }



# {
#     "id": "/signedOrderSchema",
#     "allOf": [
#         { "$ref": "/orderSchema" },
#         {
#             "properties": {
#                 "signature": { "$ref": "/hexSchema" }
#             },
#             "required": ["signature"]
#         }
#     ]
# }

