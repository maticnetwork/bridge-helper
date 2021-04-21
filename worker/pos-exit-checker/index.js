const { config } = require('dotenv')
const { MaticPOSClient } = require('@maticnetwork/maticjs')
const Matic = require('@maticnetwork/maticjs').default
const Web3 = require('web3')
const { createServer } = require('http')
const app = require('express')()
const { json } = require('express')

app.use(json())
app.use((req, _, next) => {
    console.log(`${new Date().toISOString()} | '/' | ${req.ip}`)
    next()
})

// Reading configuration params from .env file present in this directory
config({ path: '.env', silent: true })

// Matic plasma client instance, to be used for checking plasma withdraw exitability
const matic = new Matic({
    network: process.env.Network || 'testnet',
    version: process.env.Version || 'mumbai',
    parentProvider: process.env.RootRPC.startsWith('http') ? new Web3.providers.HttpProvider(process.env.RootRPC) : new Web3.providers.WebsocketProvider(process.env.RootRPC),
    maticProvider: process.env.ChildRPC.startsWith('http') ? new Web3.providers.HttpProvider(process.env.ChildRPC) : new Web3.providers.WebsocketProvider(process.env.ChildRPC),
})

// If failed to initialize, simply kills self
matic.initialize().then(_ => { }).catch(e => { console.log(e); process.exit(1); })

// Obtaining instance of matic pos client, to be used for checking whether
// given `burnTxHash` has been exited on root chain or not
const client = new MaticPOSClient(
    {
        network: process.env.Network || 'testnet',
        version: process.env.Version || 'mumbai',
        parentProvider: process.env.RootRPC.startsWith('http') ? new Web3.providers.HttpProvider(process.env.RootRPC) : new Web3.providers.WebsocketProvider(process.env.RootRPC),
        maticProvider: process.env.ChildRPC.startsWith('http') ? new Web3.providers.HttpProvider(process.env.ChildRPC) : new Web3.providers.WebsocketProvider(process.env.ChildRPC),
    }
)

// POST endpoint to be exposed, which accepts burnTxHash from child chain &
// checks whether exit has been processed on root chain or not
app.post('/', (req, res) => {
    if (req.body.txHash === undefined || req.body.txHash === null) {
        return res.status(400).json({ msg: 'Bad Payload' }).end()
    }

    if (req.body.txHash.length != 66) {
        return res.status(400).json({ msg: 'Bad Payload' }).end()
    }

    client.isERC20ExitProcessed(req.body.txHash)
        .then(v => v ? res.status(200).json({ code: 1, msg: 'Exited' }).end() : res.status(200).json({ code: 0, msg: 'Not Exited' }).end())
        .catch(_ => res.status(400).json({ msg: 'Bad Payload' }).end())
})

// Checks whether withdraw is exitable or not
// i.e. is it still in challenge period or not ?
app.post('/exit-time', (req, res) => {
    if (req.body.burnTxHash === undefined || req.body.burnTxHash === null || req.body.confirmTxHash === undefined || req.body.confirmTxHash === null) {
        return res.status(400).json({ msg: 'Bad Payload' }).end()
    }

    if (req.body.burnTxHash.length != 66 && req.body.confirmTxHash.length != 66) {
        return res.status(400).json({ msg: 'Bad Payload' }).end()
    }

    matic.withdrawManager.getExitTime(req.body.burnTxHash, req.body.confirmTxHash)
        .then(v => v.exitable ? res.status(200).json({ code: 1, msg: (Date.now() / 1000).toString() }).end() : res.status(200).json({ code: 0, msg: v.exitTime.toString() }).end())
        .catch(_ => res.status(400).json({ msg: 'Bad Payload' }).end())
})

createServer(app).listen(process.env.PORT || 7003, process.env.HOST || '127.0.0.1', _ => {
    console.log(`[+] Ready to accept requests on :${process.env.PORT}`)
})
