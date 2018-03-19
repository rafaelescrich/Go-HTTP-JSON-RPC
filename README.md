# JSON-RPC

This is a work-in-progress implementation of HTTP JSON-RPC in Go. Any help is appreciated.

## Packages details

- httpjsonrpc - package for using HTTP JSON-RPC

  - httpjsonrpcClient.go - currently contains a simple method for calling JSON-RPC through a HTTP POST

  - httpjsonrpcServer.go - currently contains a simple method for receiving JSON-RPC through a HTTP POST

  - test.go - leftover code - sample use, sample call with predefined parameters

## Useful links

- Original code - <http://stackoverflow.com/questions/8918455/how-to-properly-call-json-rpc-in-go>
- Bitcoind JSON call list - <https://en.bitcoin.it/wiki/Original_Bitcoin_client/API_Calls_list>
- Example request - <https://en.bitcoin.it/wiki/Elis-API#Example_Request>

## Install

```bash
cd $GOPATH/src/github.com
mkdir rafaelescrich
git clone git@github.com:rafaelescrich/Go-HTTP-JSON-RPC.git
```