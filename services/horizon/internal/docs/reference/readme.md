---
title: Overview
---

Horizon is an API server for the Spn ecosystem.  It acts as the interface between [spn-core](https://github.com/spn/spn-core) and applications that want to access the Spn network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Spn ecosystem](https://www.spn.org/developers/guides/) for details of where Horizon fits in. You can also watch a [talk on Horizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Spn.org developer Scott Fleckenstein:

[![Horizon: API webserver for the Spn network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Horizon: API webserver for the Spn network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Horizon provides a RESTful API to allow client applications to interact with the Spn network. You can communicate with Horizon using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Spn SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.spn.org/developers/js-spn-sdk/learn/index.html) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net: [https://horizon-testnet.spn.org/](https://horizon-testnet.spn.org/) and one that is connected to the public Spn network:
[https://horizon.spn.org/](https://horizon.spn.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/spn/js-spn-sdk)
- [Java](https://github.com/spn/java-spn-sdk)
- [Go](https://github.com/spn/go)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/spn/ruby-spn-sdk)
- [Python](https://github.com/SpnCN/py-spn-base)
- [C#](https://github.com/QuantozTechnology/csharp-spn-base)
