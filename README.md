# Generate bindings

First clone the smart-contracts repository from `git@github.com:ThingsIXFoundation/smart-contracts-evm.git`
and compile the contracts. This generated an `artifacts` directory that includes
the data to generate the bindings.

The shell script must be passed the directory to the smart contracts project and
will grab the relevant data from the `artifacts` directory and generated the
bindings in `bindings.go`.

## Example
```
./generate-bindings.sh $HOME/dev/thingsix/smart-contracts-evm
```