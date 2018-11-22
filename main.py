import go_wrapper as sloth


p=sloth.go_wrapper("")
print(p.Sloth_elapsed_time("64" , "83478237" , "1024"))




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

