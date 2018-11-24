package main
import "C"
import (
	"fmt"
	sloth "./candidates/modular_sqrt"
	"strconv"
)

var p1024 string= "26665316952145251691159678627219217222885850903741016853585447718947343212288750750268012668712469908106258613976547496870890438504017231007766799519535785905104605162203896873810538315838185502276890025696087480171103337359532995917850779890238106057070346163136946293278160601772800244012833993583077700483"
var p512 string	= "1428747867218506432894623188342974573745986827958686951828141301796511703204477877094047850395093527438571991358833787830431256534283107665764428020239091"
var p256 string	= "60464814417085833675395020742168312237934553084050601624605007846337253615407"
var p128 string	= "271387921886905605025992265577018698667"
var p64	 string	= "9853393445385562019"
//CLI, Comment if not used
func main() {
	////command line interfacce arguments, note that the [0] in the list is the path to the program, might be useful later on
	////arguments [ security parameter , starting value x, iteration count t ]
	//cli_args := os.Args[1:]
	//
	//if len(cli_args)<2{
	//	fmt.Println("Invalid Arguments From CLI")
	//	return
	//}
	//
	//if cli_args[0] == "64" {
	//	cli_args[0]=p64
	//}else if cli_args[0]=="128"{
	//	cli_args[0]=p128
	//}else if cli_args[0]=="256"{
	//	cli_args[0]=p256
	//}else if cli_args[0]=="512"{
	//	cli_args[0]=p512
	//}else if cli_args[0]=="1024"{
	//	cli_args[0]=p1024
	//} else {
	//	fmt.Println("Invalid Security Parameter")
	//	return
	//}
	//
	//if len(cli_args)<3{
	//	sloth.Elapsed_proof(cli_args)
	//} else {
	//	sloth.Fixed_delay(cli_args)
	//}
}

//all string args
//export Sloth_fixed_delay
func Sloth_fixed_delay (p_parameter string, starting_value string, iteration string)  *C.char {
	if p_parameter == "64" {
		p_parameter=p64
	}else if p_parameter=="128"{
		p_parameter=p128
	}else if p_parameter=="256"{
		p_parameter=p256
	}else if p_parameter=="512"{
		p_parameter=p512
	}else if p_parameter=="1024"{
		p_parameter=p1024
	} else {
		fmt.Println("Invalid Security Parameter")
		return C.CString("")
	}
	return C.CString(sloth.Fixed_delay([3]string{p_parameter,starting_value,iteration}))
}
//export Sloth_eval
func Sloth_eval (p_parameter string, starting_value string, iteration string) *C.char{
	if p_parameter == "64" {
		p_parameter=p64
	}else if p_parameter=="128"{
		p_parameter=p128
	}else if p_parameter=="256"{
		p_parameter=p256
	}else if p_parameter=="512"{
		p_parameter=p512
	}else if p_parameter=="1024"{
		p_parameter=p1024
	} else {
		fmt.Println("Invalid Security Parameter")
		return C.CString("")
	}
	return C.CString(sloth.Eval([3]string{p_parameter,starting_value,iteration}))
}
//export Sloth_verify
func Sloth_verify (p_parameter string, starting_value string, iteration string, ending_value string) *C.char{
	if p_parameter == "64" {
		p_parameter=p64
	}else if p_parameter=="128"{
		p_parameter=p128
	}else if p_parameter=="256"{
		p_parameter=p256
	}else if p_parameter=="512"{
		p_parameter=p512
	}else if p_parameter=="1024"{
		p_parameter=p1024
	} else {
		fmt.Println("Invalid Security Parameter")
		return C.CString("")
	}
	return C.CString(strconv.FormatBool(sloth.Verify([4]string{p_parameter,starting_value,iteration,ending_value})))
}