---
title: Horizon
---

Horizon is the server for the client facing API for the Spn ecosystem.  It acts as the interface between [spn-core](https://www.spn.org/developers/learn/spn-core) and applications that want to access the Spn network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Spn ecosystem](https://www.spn.org/developers/guides/) for more details.

You can interact directly with horizon via curl or a web browser but SDF provides a [JavaScript SDK](https://www.spn.org/developers/js-spn-sdk/learn/) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net [https://horizon-testnet.spn.org/](https://horizon-testnet.spn.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/spn/js-spn-sdk)
- [Java](https://github.com/spn/java-spn-sdk)
- [Go](https://github.com/spn/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/spn/ruby-spn-sdk)
- [Python](https://github.com/SpnCN/py-spn-base)
- [C# .NET 2.0](https://github.com/QuantozTechnology/csharp-spn-base)
- [C# .NET Core 2.x](https://github.com/elucidsoft/dotnetcore-spn-sdk)
- [C++](https://bitbucket.org/bnogal/spnqore/wiki/Home)
