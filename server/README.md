main.py defines a multi-processor VDF microservice that receives as input a starting value (e.g. trade hash hex) and time value t (number of VDF iterations) and computes t iterations of the modular square root function, returning an output. 

Run the VDF server with `python server/main.py` which spawns the server on `localhost:5000` and test it using 

```
curl  localhost:5000/vdf -d "input=0x3338be694f50c5f338814986cdf0686453a888b84f424d792af4b9202398f392&time=20555"  -X POST -i
```

### vdf(p,x,t)

Input:

`p` (string): security parameter, we decided to use prime numbers stored in the go file for now.
therefore p is the number of bits you want for your prime number. You can choose from
"64","128","256","512", and "1024".

`x` (string): starting value or input, it accepts both hexadecimal as well as decimal numbers. If the input is hexadecimal,
please ensure that the input has prefix "0x".

`t` (string): number of iterations in a decimal format, keep in mind that this is in string.

Output:

`y` (string): ending value or output of VDF Eval, in hexadecimal string format.

We do not need proof for this candidate algorithm. But for other VDF candidate algorithms, output will also have:

**not used:** `proof` (string): a list of hexadecimal numbers separated/delimited by `,`.


### verify(p,x,t,y)


Input:

`p` (string): security parameter, we decided to use prime numbers stored in the go file for now.
therefore p is the number of bits you want for your prime number. You can choose from
"64","128","256","512", and "1024".

`x` (string): starting value or input, it accepts both hexadecimal as well as decimal numbers. If the input is hexadecimal,
please ensure that the input has prefix "0x".

`t` (string): number of iterations in a decimal format, keep in mind that this is in string.


`y` (string): ending value or output of VDF Eval, in hexadecimal string format.

We do not need proof for this candidate algorithm. But for other VDF candidate algorithms, output will also have:

**not used:** `proof` (string): a list of hexadecimal numbers separated/delimited by `,`.

Output:

`true/false` (bool): the result of verification, seeing whether the computation was done correctly with the given inputs.
Even if one of the inputs is incorrect/inaccurate, the output will return `false`.
