---
title: Chainalysis Able to Analyze Darknet Fentanyl Dealer's Bitcoin Transactions
date: 2019-10-01
tags: [ chainalysis, darknet ]
srcs:
 - [ 'blog.chainalysis.com/reports/fentanyl-cryptocurrency-transactions', 'archive.ph/OtJmG' ]
 - [ 'justice.gov/usao-mdfl/press-release/file/982971/download', 'web.archive.org/web/20210309024802/https://www.justice.gov/usao-mdfl/press-release/file/982971/download' ]
---

In February of 2017, a woman suffered a fatal overdose on fentanyl which was
purchased on the darknet by her fiancè (who also suffered an overdose but
survived). The U.S. Drug Enforcement Agency (DEA) pursued an investigation
which led to the arrest of the dealer "ETIKING" in late June of 2017. He was
[indicted](https://archive.ph/HUrpL) less than a month later,
[convicted](https://archive.ph/d99Q6) in January of 2018, and [sentenced to
life in prison](https://archive.ph/vNJHS) in April of 2018.

The agents were able to focus their investigation on ETIKING after a
confidential source gave them a Bitcoin address that they linked to a
Coinbase.com account, which obviously contained some of his personally
identifying information. Nearly 18 months later, surveillance corporation
Chainalysis analyzed this case to show how many leads could have been generated
early on using just that lone Bitcoin address.

> Our goal was to learn if the tool would be helpful for similar investigations
> in the future, and we weren’t disappointed. <mark>Reactor surfaced a wealth
> of information and potential leads</mark> law enforcement could have pursued
> to identify ETIKING, <mark>starting with nothing more than his Bitcoin
> address.</mark>

Unsurprisingly, according to the Chainalysis breakdown, ETIKING was primarily
[receiving](https://archive.ph/OtJmG#selection-369.14-369.193) coins from two
darknet markets (AlphaBay and Dream Market) and primarily
[sending](https://archive.ph/OtJmG#selection-365.309-365.527) his coins to four
exchanges. However, by "looking more closely at an unusual transaction in
ETIKING’s sending exposure," Chainalysis was also [able to identify
transactions](https://archive.ph/OtJmG#selection-401.0-415.153) to a drug
testing lab in Barcelona named Energy Control International.  Additionally,
they were even [able to identify a whole new cluster of
transactions](https://archive.ph/OtJmG#selection-419.9-419.331) that _likely_
belong to ETIKING, not only based on the **similar general transaction
pattern**, but also a connection to the original address:

> [...] by backtracking ETIKING’s deposits, we can identify another cluster of
> addresses making deposits to the same addresses at three of the exchanges
> ETIKING favors (the green arrows), and receiving funds from the same darknet
> markets (the blue arrows). This new cluster of addresses is also likely to be
> controlled by ETIKING.

Chainalysis [concludes with the
obvious](https://archive.ph/OtJmG#selection-455.259-455.487): their products
can generate numerous possible leads from just one Bitcoin address, and **the
(centralized) exchanges "are the real goldmine here."** However, if this much
data can be gathered on a fentanyl dealer who uses Bitcoin, it can be done for
a journalist or political dissident who does as well.
