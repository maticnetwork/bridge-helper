## Micro Services

Name | Purpose
--- | ---
`state-id-manager` | Keeps track of what's latest value of `lastStateId`, which was synced into child chain, which is to be used for checking whether a deposit tx went through or not
`check-point-tracker` | Keeps track of what's latest checkpoint's block range, which is to be used for checking whether a burn tx has been pushed to root chain or not
`pos-exit-checker` | Given burn txHash, checks whether it has been exited on root chain using POS bridge or not
