# VDF
VDF Implementations

## Getting started
This is a simple implementation of Verifiable Delay Function. We have a Golang implementation of 
two candidates including Piertzak VDF (Simple VDF) and Sloth (Modular Square Root). You can locate the candidates in
``
cd go_src/candidates
``.
### Prerequisites

```
Python 2.7.1
```
Note: `vdf_interface` is built for MacOS. Rebuild VDF interface in go_src if you are operating on another OS.
### Run

In ``main.py`` we run an example VDF operation over a 128bit prime field.
```
python main.py
```
## Intro to VDF
VDF is a tuple of three algorithms:

<a href="https://www.codecogs.com/eqnedit.php?latex=Setup(\lambda,&space;T)\rightarrow&space;pp" target="_blank"><img src="https://latex.codecogs.com/gif.latex?Setup(\lambda,&space;T)\rightarrow&space;pp" title="Setup(\lambda, T)\rightarrow pp" /></a>

<a href="https://www.codecogs.com/eqnedit.php?latex=Eval(pp,x)&space;\rightarrow&space;(y,\pi)" target="_blank"><img src="https://latex.codecogs.com/gif.latex?Eval(pp,x)&space;\rightarrow&space;(y,\pi)" title="Eval(pp,x) \rightarrow (y,\pi)" /></a>

<a href="https://www.codecogs.com/eqnedit.php?latex=Verify(pp,x,y,\pi)&space;\rightarrow&space;\{accept,reject\}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?Verify(pp,x,y,\pi)&space;\rightarrow&space;\{accept,reject\}" title="Verify(pp,x,y,\pi) \rightarrow \{accept,reject\}" /></a>

### Modular Square Root

#### Evaluation
Sloth, or modular sqaure root, is one of the simplest yet effective candidate for verifiable delay function.

Given a prime field ``p``, starting value ``x``, and delay parameter ``t``:

We assume that finding ``x^2 % p`` is much easier than finding ``sqrt(x) % p``.
This is because the fastest existing algorithm to find modular square root of ``x``
is ``x^((p+1)/4) % p``, which requires a lot more operations than simply squaring ``x``.

This also implies that as the prime field ``p`` gets larger, more operations are required.

To implement Sloth, we first select a prime field ``p`` where ``p%4=3``. When we evaluate starting value ``x``, 
we have to first check whether ``x`` is a quadratic residue by checking whether 
``a^((p-1)/2)%p=1``. If not, we simply change starting value ``x`` to ``(-x)%p``.

#### Chaining

In Sloth, the ``t`` parameter specifies the number of iterations the VDF evaluation has undergone.
Since ``sqrt(x)->y``, we can extend the delay of Sloth by performing the same evaluation over ``y`` with 
``sqrt(x)-> y , sqrt(y)-> y_1 , sqrt(y_1)->y_2 ...``. This way long as the prime field (>256 bits)is large enough, 
the length of delay provided from adjusting ``t`` should be more than enough.

 

#### Proof Generation
Sloth is fairly straightforward when it comes to verification. We do not need to submit a proof 
for a basic Sloth implementation. Verifiers can simply square ``y`` once or ``y_t`` t+1 times to 
acquire the starting value ``x``.

For a faster verification, we can use SNARK proofs to be submitted along with ``y``. 
However, if we want better a verification vs. evaluation gap, there are better candidates 
for it. 

### Time Proof
In our protocol,``t`` is the time parameter for proof of elapsed time. In the context of Sloth, ``t`` 
is the number of modular square root evaluation. ``t`` can be computed incrementally for it to become a time proof.