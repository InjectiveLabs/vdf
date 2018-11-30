from flask import Flask
from flask import request
import time
from multiprocessing import Pool
import go_wrapper
# predefine security parameter
p="256"
app = Flask(__name__)
_pool = None

# all int args
def vdf(p, input, time):
    # do your expensive time consuming process
    v = go_wrapper.go_wrapper()
    y = v.Sloth_eval(p,input,time)
    return y

def verify (p,x,t,y):
    v = go_wrapper.go_wrapper()

    return v.Sloth_verify(p,x,t,y)=="true"

@app.route('/vdf', methods=['POST'])
def compute():
    input = request.form.get('input')
    time = request.form.get('time')
    print(input, time)
    # perform validation, invoke executable, retreive output, return
    if validate(input, time):
         f = _pool.apply_async(vdf, (p,input,time))
         r = f.get(timeout=2) # timeout probably should be a function of time with a max cap
         r=hex(int(r))
         return r
    return "Fail message" # change to HTTP response

@app.route('/verify', methods=['POST'])
def confirm():
    input = request.form.get('input')
    time = request.form.get('time')
    output = request.form.get('output')
    print(input, output, time)
    # perform validation, invoke executable, retreive output, return
    if validate(input, time):
         f = _pool.apply_async(verify, (p,input,time,output))
         r = f.get(timeout=2) # timeout probably should be a function of time with a max cap
         # result is a bool
         if r:
             return "verification successful"
         else:
             return "verification failed"

    return "Fail message" # change to HTTP response


# Make sure that input and time have the correct form
def validate(input, time):
    if not (input==None and time==None):
        return True


if __name__ == '__main__':
    _pool = Pool(processes=8)
    try:
        app.run()
    except KeyboardInterrupt:
        _pool.close()
        _pool.join()
