
<img src="https://img.shields.io/github/last-commit/nahuelmol/godns"/>
<img src="https://img.shields.io/github/languages/code-size/nahuelmol/godns"/>
<img src="https://img.shields.io/github/languages/top/nahuelmol/godns"/>

## about this project
Trying to implement a basic DNS recursive resolver

The idea is to make this DNS server to ask for a domain to the root server, using the same request tha the client sends. This looks into its database and response with a list of TLD servers that control ".com".

Later, it takes the list and select an specific server. The next step is to ask to the selected server for "google.com" keeping the request used for the client, which will answer with a list of X.google.com.

Following, the same logic, our takes a server and ask to it for "www.google.com" which will response with the IP.

Note that, the request is the same, using "www.google.com". The server, depending on if it is authoritative for the zone.

## source

DNS query message

https://www.firewall.cx/networking/network-protocols/dns-protocol/protocols-dns-query.html

https://ns1.com/resources/dns-protocol

## cmd

```
got test -v ./test -run [function tester]
```



