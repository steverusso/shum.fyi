---
title: Elliptic Adds "Blockchain Monitoring" Support for Zcash & Horizen
date: 2020-06-30
tags: [ elliptic, zcash ]
srcs:
 - [ 'elliptic.co/blog/achieving-aml-and-sanctions-compliance-with-privacy-coins', 'archive.ph/iReJt' ]
 - [ 'elliptic.co/media-center/elliptics-adds-privacy-coins-to-blockchain-monitoring-platform', 'archive.ph/65exu' ]
---

Surveillance corporation Elliptic announced that it added "blockchain
monitoring" support for "privacy coins" Zcash and Horizen. Elliptic's blog post
on this matter [describes](https://archive.ph/iReJt#selection-891.0-899.266)
how there are two "classes" of privacy coins: privacy-by-default ones like
Monero (which are essentially impossible to develop monitoring tools for) and
opt-in privacy ones like Zcash:

> The first class of privacy coins are <mark>private by default. This includes
> the likes of monero, which seeks to conceal the details of all transactions.
> It is unlikely that there will ever be blockchain monitoring tools that allow
> compliance professionals to trace monero transactions</mark> - if such a
> capability existed it would defeat the point of such an asset.
>
> The second class of privacy coins are those that have opt-in privacy. This
> includes the likes of Zcash (ZEC) and Horizen (ZEN), which allow users to
> choose whether to make their transactions visible on the blockchain or not.
> The vast majority of transactions that take place in these currencies look
> the same as bitcoin transactions - they are fully visible on the blockchain,
> and blockchain monitoring tools such as Elliptic’s can be used to trace the
> source and destination of funds, and assess their risk.

Therefore, as the post [goes on to
explain](https://archive.ph/iReJt#selection-911.0-911.493), services that use
surveillance products can just treat Zcash shielded transactions the same way
they treat Bitcoin transactions originating from a mix or coinjoin (in reality,
this treatment very much resembles 'guilty until proven innocent'):

> If blockchain monitoring indicates that a Zcash deposit ultimately originated
> from a shielded wallet, then as with funds from a bitcoin mixer, the funds
> cannot be traced any further. But in both cases, an exchange can still use
> solutions like Elliptic’s to assign risk scores to the transactions that
> reflect its risk appetite. And in both cases, the compliance analyst at the
> exchange can then use the same processes to assess the risk of this
> customer’s transaction and determine next steps.
