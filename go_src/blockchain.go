package main

import "math/big"

type t_order struct {
	rate			uint32
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
type block struct {
	height				uint64
	m_order_bk	 		[]m_order//make order book
	t_order_bk	 		[]t_order//take order book
	e_t_order_bk 		[]e_t_order
	height_m_order_bk	uint64
	height_t_order_bk	uint64
	height_e_t_order_bk	uint64


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
