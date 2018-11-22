

package simple_vdf
import (
"math/big"
"time"
"fmt"
"crypto/sha256"
"runtime"
)
type pair struct{
	a 			*big.Int
	b 			*big.Int
}
type trip struct{
	g1 			*big.Int
	h1 			*big.Int
	t 			int
}
type cipher_pair struct{
	c 			*big.Int
	proof		[]trip
}
func square(base, modulus *big.Int) *big.Int {
	return new(big.Int).Mod(new(big.Int).Mul(base, base),modulus)
}
func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}
func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func Div(x, y *big.Int) *big.Int {
	return big.NewInt(0).Div(x, y)
}
func sha(input *big.Int) *big.Int{
	h := sha256.New()
	h.Write(input.Bytes())
	return new(big.Int).SetBytes(h.Sum(nil))
}

func encode_byte(t int, m ,N *big.Int) cipher_pair{
	x:=new(big.Int).Mod(m,N)
	proof:=[]*big.Int{}
	start:= time.Now()
	proof=append(proof, x)
	h:=x
	for x := 0; x < t; x++{
		h=square(h,N)
		proof=append(proof, h)
	}
	cur:= time.Now()
	println("Eval Elapsed",fmt.Sprintf("%.2f",cur.Sub(start).Seconds()),"sec")


	start = time.Now()
	prooflist:=proof_list_build(proof,h,N,t,[]trip{})
	cur = time.Now()
	println("Prooflist Elapsed",fmt.Sprintf("%.2f",cur.Sub(start).Seconds()),"sec")


	return cipher_pair{h,prooflist}
	}

//g: base
//r: random
//N: field
//v: result of T/2
//h: result of T

func proof_list_build(proof []*big.Int,h,N *big.Int, T int, prooflist[] trip)[]trip{

	for t:=T;t>=2;t=t/2 {
		h := new(big.Int)
		if (t)%2 != 0 {
			t++
			h = square(proof[t-1], N)
		} else {
			h = proof[t]

		}
		g := proof[0]
		r := new(big.Int).Mod(sha(proof[t/2]), N)
		v := proof[t/2]
		out := proof_build(g, v, h, r, N)
		prooflist = append(prooflist,trip{out.a,out.b,t/2})
	}
	return prooflist
}

func proof_build (g,v,h,r,N *big.Int) pair{
	g1 := make(chan *big.Int)
	h1 := make(chan *big.Int)

	go func (g,r,v,N*big.Int, c chan *big.Int){
		exp:=new(big.Int).Exp(g,r,N)
		mul:=new(big.Int).Mul(exp,v)
		mod:=new(big.Int).Mod(mul,N)
		c <- mod
	}(g,r,v,N,g1)

	go func (v,r,h,N*big.Int, c chan *big.Int){
		exp:=new(big.Int).Exp(v,r,N)
		mul:=new(big.Int).Mul(exp,h)
		mod:=new(big.Int).Mod(mul,N)
		c <- mod
	}(v,r,h,N,h1)
	outh:=<-h1
	outg:=<-g1
	return pair{outg,outh}
}






func verify(cp cipher_pair,N *big.Int) []bool{
	prooflist:=cp.proof
	boolist:=[]bool{}
	c:=make(chan bool)
	start := time.Now()
	for _,item := range prooflist{
		go func(item trip,N *big.Int,c chan bool){
			g:=item.g1
			for x:=0;x<item.t;x++{
				g=square(g,N)
			}
			c<- g.Cmp(item.h1)==0
		}(item,N,c)
	}
	for range prooflist{
		boolist = append(boolist, <-c)
	}
	cur := time.Now()
	println("Verify Elapsed",fmt.Sprintf("%.2f",cur.Sub(start).Seconds()),"sec")

	println(len(boolist))
	return boolist
}

func Test(){
	println("GOMAXPROCS is:", runtime.GOMAXPROCS(0))
	t:=100000
	input,_:=new(big.Int).SetString("349590234923847372",0)
	var p, _ =new(big.Int).SetString("126493185890016866190387990037436305339",0)
	var q, _ =new(big.Int).SetString("237515677732435432578220196406645605033",0)
	//p*q=N
	var N=Mul(p,q)

	start := time.Now()
	starting_value:=encode_byte(t,input,N)
	cur := time.Now()
	println("Encode Elapsed",fmt.Sprintf("%.2f",cur.Sub(start).Seconds()),"sec")
	println(starting_value.c.String())
	verif :=true
	for _,statement := range verify(starting_value,N){
		if statement!=true{
			verif=false
		}
	}
	println(verif)
}

