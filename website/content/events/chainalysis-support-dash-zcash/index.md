---
title: Chainalysis Adds "Investigation and Compliance" Support for Dash & Zcash
date: 2020-06-08
tags: [ chainalysis, zcash, dash ]
srcs:
 - [ 'blog.chainalysis.com/reports/introducing-chainalysis-investigation-compliance-support-dash-zcash', 'archive.ph/jO7se' ]
---

Chainalysis, a blockchain surveillance corporation, announced in a blog post
that they now support Dash and Zcash in their "investigation and compliance"
product(s). How can blockchain surveillance products support "privacy" coins?
Further into the article, under a section titled "Real world use of Dash and
Zcash privacy features," Chainalysis goes into the "privacy" features each coin
offers and how they are actually being used "in the wild."

[For Dash](https://archive.ph/jO7se#selection-381.0-385.272), PrivateSend is
essentially a basic coinjoin implementation, and no one uses it:

> Mixing transactions related to PrivateSend make up roughly 9% of all Dash
> transactions. [...] The percentage of Dash transactions that constitute
> actual transfers of funds using PrivateSend is less than 0.7%.

[For Zcash](https://archive.ph/jO7se#selection-417.0-421.190), users have the
option to sheild certain types transactions, and no one does:

> Roughly 14% of Zcash transactions involve one of Zcash’s two shielded pools
> in some way. But of the transactions that interact with a shielded pool, only
> 6% are completely shielded, i.e. sender, receiver, and transaction amount are
> all encrypted. That’s only 0.9% of all Zcash transactions.
>
> So even though the obfuscation on Zcash is stronger due to the zk-SNARK
> encryption, Chainalysis can still provide the transaction value and at least
> one address for over 99% of ZEC activity.

Therefore, because these "privacy" features are optional, and especially
because almost no one uses them, these coins are about as easy to surveil as
Bitcoin and any other crypto that has an open, transparent blockchain. Hence
corporations like Chainalysis can easily expand their surveillance product
lines by adding support for them.
