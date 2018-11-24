# This is a wrapper over the compiled binary vdf_interface
# sourcecode can be found in go_src directory
# This only contains the VDF sloth candidate (modular square root)

from ctypes import *
import os
print( os.path.dirname(os.path.realpath(__file__)))
class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]

class go_wrapper:



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
        self.lib.Sloth_fixed_delay.argtypes = [c_char_p, c_char_p, c_char_p]
        self.lib.Sloth_fixed_delay.restype = c_char_p
        # p = GoString(c_char_p(p.encode('utf-8')), len(p))
        print(c_char_p(x.encode('utf-8')))
        # x = GoString(c_char_p(x.encode('utf-8')), len(x))
        # t = GoString(c_char_p(t.encode('utf-8')), len(t))
        i=1

        y = self.lib.Sloth_fixed_delay(c_char_p(p.encode('utf-8')), c_char_p(x.encode('utf-8')),
                                       c_char_p(t.encode('utf-8')))
        while True:
            x=y
            y=self.lib.Sloth_fixed_delay(c_char_p(p.encode('utf-8')), x,c_char_p(t.encode('utf-8')))
            print(y)
            i+=1
            print(x)
            print("Total Iterations: ",int(iteration)*i)
    # go file should be in the root directory of the lib... unless changed otherwise
    def __init__(self):

        self.lib = cdll.LoadLibrary("./vdf_interface.so")

