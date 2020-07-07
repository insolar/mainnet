# Airdrop

Insolar airdrop is a simple CLI tool for airdrop distribution at Insolar Platform.

The airdrop automates transfer from one member to several others.


## Usage

To use the airdrop, do the following:

1. Build it. In your `insolar/mainnet/` directory, run:

   ```console
   make airdrop
   ```
   
   This builds a `.bin/airdrop` binary in the `insolar/mainnet/` directory.

2. Get hex private key of member, from who airdrop will be done.
   
3. Create `members.json` file with members wallets reference and amount of XNS tokens to transfer for each.
   For example:

   ```json
   [
     {
        "wallet": "insolar:1AtrwSmt7cy7zgGFNVD4Q72IV2vFeLREOTfwhgyOg5cc",
        "tokens": 10
     },
     {
        "wallet": "insolar:1AtrwaFcqV4DU8P7g_im-ynyWjOx8T4XKU6b0v9Nt4uc",
        "tokens": 4
     },
     ...
   ]
   ```
   
4. To run the airdrop, specify:
 
   - Insolar Platform RPC endpoint as a URL parameter.
   - Private key in hex of member from who airdrop will be done `-p` option's value.
   - Path to `members.json` as the `-m` option's value.
   
   For example:

   ```console
   ./bin/airdrop https://<endpoint>/api/rpc -p 45eefdac06474a1812ca44d30b48a326f5041f5b095a9e9513b114591fda8ac3 -m members.json 
   ```

## Airdrop options

    Insolar airdrop is a simple CLI tool for airdrop distribution at Insolar Platform
    
    Usage:
      airdrop <insolar_endpoint> [flags]
    
    Examples:
    ./bin/airdrop http://localhost:19101/api/rpc -p <from_member_hex_private_key> -m members.json 
    
    Flags:
      -p, --hexPrivate          Airdrop from member private key in hex
      -m, --membersPath         Path to a file with members to airdrop
      -h, --help                Help for airdrop
      -v, --verbose             Print request information
