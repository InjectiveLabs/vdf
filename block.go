package main
import (
	"math/big"
	"time"
	"fmt"
)
type block struct {
	height				uint64
	m_order_bk	 		[]m_order//make order book
	t_order_bk	 		[]t_order//take order book
	e_t_order_bk 		[]e_t_order
	height_m_order_bk	uint64
	height_t_order_bk	uint64
	height_e_t_order_bk	uint64


}
type t_order struct {
	rate			uint32
	// yolo float, will modify this for actual implementation
	vol				uint32
	origin_address	[]byte
	sig				[]byte
	dest_address	[]byte
}
type m_order struct {
	rate			uint32
	// yolo float, will modify this for actual implementation
	vol				uint32
	origin_address	[]byte
	sig				[]byte
}
//everything is encrypted, formally into bytes but for demo purposes we use uint.
type e_t_order struct{
	rate			uint32
	// yolo float, will modify this for actual implementation
	vol				uint32
	origin_address	[]byte
	sig				[]byte
	dest_address	[]byte
}
type cipher_pair struct{
	c 			*big.Int
	positive	bool
}
func update_block(prev_block block, new_m_order_bk []m_order,
	new_t_order_bk []t_order,
	new_e_t_order_bk []e_t_order) (new_block block){
	prev_block.height 				=prev_block.height+1
	prev_block.m_order_bk 			=new_m_order_bk
	prev_block.t_order_bk 			=new_t_order_bk
	prev_block.e_t_order_bk 		=new_e_t_order_bk
	prev_block.height_m_order_bk	=prev_block.height_m_order_bk+1
	prev_block.height_t_order_bk 	=prev_block.height_t_order_bk+1
	prev_block.height_e_t_order_bk	=prev_block.height_e_t_order_bk+1
	return prev_block
}
func mod_exp(base, exponent, modulus *big.Int) *big.Int {
	return new(big.Int).Exp(base, exponent, modulus)
}
//p:=2733211297152089461847926391575359211106319068554693017579269301794909490925781940161593535758510845944723920774129059209700248319180915042180597186066517
func quad_res(x,p *big.Int) bool{
	t:=new(big.Int).Exp(x,new(big.Int).Div(Sub(p,big.NewInt(1)),big.NewInt(2)),p)
	return t.Cmp(big.NewInt(1))==0
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
func mod_sqrt(x,p *big.Int) *big.Int{
	y:=big.NewInt(0)
	if quad_res(x,p){
		y=new(big.Int).Exp(x,Div(Add(p,big.NewInt(1)),big.NewInt(4)),p)
	} else{
		x:=big.NewInt(0).Mod(big.NewInt(0).Neg(x),p)
		y=new(big.Int).Exp(x,Div(Add(p,big.NewInt(1)),big.NewInt(4)),p)
	}
	return y
}
func square(y,p *big.Int) *big.Int{
	return big.NewInt(0).Exp(y,big.NewInt(2),p)
}
func verify(x,y,p *big.Int) bool{
	if quad_res(x,p){
		z:=square(y,p)
		return x.Cmp(z)==0
	} else{
		x:=big.NewInt(0).Mod(big.NewInt(0).Neg(x),p)
		z:=square(y,p)
		return x.Cmp(z)==0
	}
}
//var x, _ =new(big.Int).SetString("48579348758743879",0)
//pretty useless function at this stage, will remove later
func modsqrt_op(t int,x,p *big.Int) *big.Int{
	y:=x
	for i:=0;i<t;i++{
		y=mod_sqrt(y,p)
	}
	//	println(verify(x,y,p))
	return y
}
func encode_32(t int, m uint32,p *big.Int) cipher_pair{
	encrypted_m:=big.NewInt(int64(m))
	for x := 0; x < t; x++{
		encrypted_m=square(encrypted_m,p)
	}
	if quad_res(big.NewInt(int64(m)),p){
		return cipher_pair{encrypted_m,true}
	} else {
		return cipher_pair{encrypted_m,false}
	}

}
func encode_byte(t int, m []byte,p *big.Int) cipher_pair{
	encrypted_m:=new(big.Int).SetBytes(m)
	for x := 0; x < t; x++{
		encrypted_m=square(encrypted_m,p)
	}
	if quad_res(new(big.Int).SetBytes(m),p){
		return cipher_pair{encrypted_m,true}
	} else {
		return cipher_pair{encrypted_m,false}
	}
}
func decode(t int ,pair cipher_pair,p *big.Int) *big.Int{
	c:=pair.c
	z:=modsqrt_op(t,c,p)
	if pair.positive{
		return z
	} else{
		return big.NewInt(0).Mod(big.NewInt(0).Neg(z),p)
	}
}
func test(){
	// t as the length of the hash chain
	// t as the length of the hash chain
	t:=5000
	address:=[]byte("Lets_buy_this_shit_coin")
	var p, _ =new(big.Int).SetString("173397306592529065770607141968619470475653545575095293960792220585634982267927742165967848777476838142554533411692097947431611647746270773477484159413151694046613751536944251917727745942837116899911211300281569339109672955750917474682700268528337927874597846606602536703955940895955837514429833279735894760411",0)
	start := time.Now()

	cipher:=encode_byte(t,address,p)
	cur := time.Now()
	println("Encode Elapsed",fmt.Sprintf("%.2f",cur.Sub(start).Seconds()),"sec")
	println("Encoded it into:",cipher.c)
	start = time.Now()
	message:=decode(t,cipher,p)
	cur = time.Now()
	println("Decode Elapsed: ",fmt.Sprintf("%.2f",cur.Sub(start).Seconds()),"sec")
	println(string(message.Bytes()))
}
func main(){
	t:=1000
	address:=[]byte("Lets_buy_this_shit_coin_salt_me_daddy")
	var p, _ =new(big.Int).SetString("173397306592529065770607141968619470475653545575095293960792220585634982267927742165967848777476838142554533411692097947431611647746270773477484159413151694046613751536944251917727745942837116899911211300281569339109672955750917474682700268528337927874597846606602536703955940895955837514429833279735894760411",0)
	start := time.Now()
	starting_value:=encode_byte(t,address,p)
	cur := time.Now()
	println("Encode Elapsed",fmt.Sprintf("%.2f",cur.Sub(start).Seconds()),"sec")
	t_list:=[]int{}
	for true{
		start = time.Now()
		starting_value.c=modsqrt_op(t,starting_value.c,p)
		cur = time.Now()
		elapsed:=cur.Sub(start).Seconds()
		println("Delay Elapsed: ",fmt.Sprintf("%.2f",elapsed),"sec")
		t=int(float64(t)*(2/elapsed))
		t_list=append(t_list,t)
		var total int = 0
		for _, value:= range t_list {
			total += value
		}
		t=total/int(len(t_list))

		println("t modified to",t)
	}
}
