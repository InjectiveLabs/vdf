# This is a wrapper over the compiled binary vdf_interface
# sourcecode can be found in go_src directory
# This only contains the VDF sloth candidate (modular square root)

from ctypes import *
import os
class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

class go_wrapper:



    # isolated eval for more flexibility
    # input:
    # string p: security parameter, we decided to use prime numbers stored in the go file for now
    # therefore p is the number of bits you want for your prime number, there are 64,128,256,512,1024 to choose
    # string x: starting value or input, it accepts both hexadecimal as well as arbitrary numbers
    # string t: number of iterations, keep in mind that this is in string
    # output:
    # string y: ending value of VDF Eval
    # We do not need proof for this candidate
    def Sloth_eval (self, p , x, t):
        self.lib.Sloth_eval.argtypes = [GoString, GoString, GoString]
        self.lib.Sloth_eval.restype = c_char_p
        p = GoString(c_char_p(p.encode('utf-8')), len(p))
        x = GoString(c_char_p(x.encode('utf-8')), len(x))
        t = GoString(c_char_p(t.encode('utf-8')), len(t))
        y=self.lib.Sloth_eval(p, x, t)
        return y.decode('utf-8')

    # input:
    # beyond variables mentioned in Sloth_eval, we also have:
    # string y: ending value of VDF Eval
    # output:
    # string result: boolean "true" "false" in string
    def Sloth_verify(self, p , x, t , y):

        self.lib.Sloth_verify.argtypes = [GoString, GoString, GoString, GoString]
        self.lib.Sloth_verify.restype = c_char_p
        p = GoString(c_char_p(p.encode('utf-8')), len(p))
        x = GoString(c_char_p(x.encode('utf-8')), len(x))
        t = GoString(c_char_p(t.encode('utf-8')), len(t))
        y = GoString(c_char_p(y.encode('utf-8')), len(y))
        result = self.lib.Sloth_verify(p, x, t , y)
        return result.decode('utf-8')




    # go file should be in the root directory of the lib... unless changed otherwise

    def __init__(self):
        self.lib = cdll.LoadLibrary("./vdf_interface.so")


    # from go: func Sloth_fixed_delay (p_parameter string, starting_value string, iteration string)  string
    def Sloth_fixed_delay(self, p , x , t):
        self.lib.Sloth_fixed_delay.argtypes = [GoString, GoString, GoString]
        self.lib.Sloth_fixed_delay.restype = c_char_p
        p = GoString(c_char_p(p.encode('utf-8')), len(p))
        x = GoString(c_char_p(x.encode('utf-8')), len(x))
        t = GoString(c_char_p(t.encode('utf-8')), len(t))

        return self.lib.Sloth_fixed_delay(p,x,t)

    # This is simply a loop over fixed delay function call in go.
    # t is the number of iteration of the VDF operation
    def Sloth_elapsed_time(self, p , x , t):
        iteration=t
        self.lib.Sloth_fixed_delay.argtypes = [GoString, GoString, GoString]
        self.lib.Sloth_fixed_delay.restype = c_char_p
        p = GoString(c_char_p(p.encode('utf-8')), len(p))
        x = GoString(c_char_p(x.encode('utf-8')), len(x))
        t = GoString(c_char_p(t.encode('utf-8')), len(t))
        i=0
        while True:
            y=self.lib.Sloth_fixed_delay(p,x,t)
            x=GoString(y,len(y))
            i+=1
            print("Total Iterations: ",int(iteration)*i)
