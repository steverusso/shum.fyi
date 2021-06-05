---
title: Carnegie Mellon Study on XMR vs ZEC Shows That Optional Privacy Is a Fundamental Flaw
date: 2020-05-18
tags: [ zcash ]
srcs:
 - [ '', 'eprint.iacr.org/2020/593.pdf' ]
---

A group of researchers from Carnegie Mellon University published a study that
analyzed the traceability of Monero vs Zcash. In sum: Monero is untraceable;
Zcash is at least as traceable as Bitcoin. The following is an excerpt from the
abstract:

> We run some traceability experiments based on previously published papers for
> each coin. Results show that, introducing strict security and anonymity
> requirements into the cryptocurrency ecosystem makes the coin effectively
> untraceable, as shown by Monero. On the other hand, Zcash still hesitates to
> introduce changes that alter user behavior. Despite its strong cryptographic
> features, transactions are overall more traceable.

For Monero, unsurprisingly, it found that nearly zero transactions over the
last two years of the blockchain being analyzed (the beginning up until April
15, 2020) are traceable:

> Though we found that some transactions as recent as April 9, 2020, are fully
> deducible, the percentage of partially or fully deducible transactions has
> been nearly zero for over two years, as seen in Figure 12.

For Zcash, the results showed that the vast majority of transactions are just
as traceable as Bitcoin because users don't take advantage of the anonymity
features:

> From these two figures we can see how in the Zcash ecosystem, the majority of
> participants are not taking advantage of the privacy benefits of the protocol
> that implement zero-knowledge proofs aimed to increase anonymity. The
> majority of participants in the system are using Zcash public t-to-t
> transaction, which mirrors the Bitcoin ecosystem and its anonymity issues.

And it's not just that no one uses the optional privacy features. As it turns
out, the opt-in "privacy" itself might be weakening any shred of anonymity
that's even there in the first place:

> Since Zcash has this additional layer of obfuscation that Bitcoin does not in
> the form of shielded addresses, the use patterns of individuals within the
> shielded pool can go lengths to decreasing the anonymity of Zcash.

The study is only 24 pages long and well worth a read. It was indirectly
supported less than a month later when [Chainalysis added "investigation and
compliance" support for Zcash](/e/chainalysis-support-dash-zcash/).
