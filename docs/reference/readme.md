---
title: Overview
---

Equator is an API server for the Zion ecosystem.  It acts as the interface between [zion-core](https://github.com/zion2100/zion-core) and applications that want to access the Zion network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Zion ecosystem](http://zionc.info/developers/guides/) for details of where Equator fits in. You can also watch a [talk on Equator](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Zion.org developer Scott Fleckenstein:

[![Equator: API webserver for the Zion network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Equator: API webserver for the Zion network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Equator provides a RESTful API to allow client applications to interact with the Zion network. You can communicate with Equator using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Zion SDK in the language of your client.
SDF provides a [JavaScript SDK](http://zionc.info/developers/js-zion-sdk/learn/index.html) for clients to use to interact with Equator.

SDF runs a instance of Equator that is connected to the test net: [https://equator-testnet.zion.org/](https://equator-testnet.zion.org/) and one that is connected to the public Zion network:
[https://equator.zion.org/](https://equator.zion.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/zion2100/js-zion-sdk)
- [Java](https://github.com/zion2100/java-zion-sdk)
- [Go](https://github.com/zion2100/go)

Community maintained libraries (in various states of completeness) for interacting with Equator in other languages:<br>
- [Ruby](https://github.com/zion2100/ruby-zion-sdk)
- [Python](https://github.com/ZionCN/py-zion-base)
- [C#](https://github.com/QuantozTechnology/csharp-zion-base)
