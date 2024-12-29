# mylayer
**mylayer** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

installing ignite to your codebase with ubantu with a software known as wsl.
```
curl https://get.ignite.com/cli! | bash

```
creating our own file for changes and u=installing cosmos sdk.
```
ignite scaffold chain mylayer
cd mylayer

```
communicating with the blockchain .
```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

#### Commands for interacting with blockchain.

Use the following command to see the address tied to the mnemonic:
```
mylayerd keys list
```
To see how many tokens you have:
```
mylayerd query bank balances <your-address>
```
If you want to send tokens to another address:
```
mylayerd tx bank send <your-address> <to-address> 100token --chain-id mylayer --gas auto --fees 1token

```

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

#### online Use the Faucet

Open your browser and navigate to:
```
http://localhost:4500
```
to see your transaction:
```
mylayerd query txs --query "message.sender='<your-address>'"
```

### Optional Flags

Pagination: If there are many transactions, you can use the --page and --limit flags:
```
mylayerd query txs --query "message.sender='<your-address>'" --page 1 --limit 50

```
Output in JSON: If you want the output in JSON format for easier parsing:
```
mylayerd query txs --query "message.sender='<your-address>'" -o json

```

### Adding additional module
Run the following command to scaffold a module named gasprice:
```
ignite scaffold module <module-name>

```

### Add Governance Support

Enable governance proposals to update the gas price:
```
ignite scaffold proposal SetMinGasPrice minGasPrice:dec --module gasprice

```

### Test Your Chain

Run the Chain:
```
ignite chain serve
```
Submit a Proposal:
```
mylayerd tx gov submit-proposal set-min-gas-price --from validator --gas 200000 --chain-id mylayer
```
Verify Transactions: Ensure transactions below the minimum gas price are rejected.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/mylayer@latest! | sudo bash
```
`username/mylayer` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
