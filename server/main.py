from flask import Flask
from flask import request
import time
from multiprocessing import Pool
import go_wrapper as vdf
# predefine security parameter
p="256"
app = Flask(__name__)
_pool = None

# all int args
def vdf(p, input, time):
    # do your expensive time consuming process
    v = vdf.go_wrapper()
    y = v.Sloth_eval(p,x,t)
    return y

def verify (p,x,t,y):
    v = vdf.go_wrapper()

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
         return r
    return "Fail message" # change to HTTP response

# Make sure that input and time have the correct form
def validate(input, time):
    return True


if __name__ == '__main__':
    _pool = Pool(processes=8)
    try:
        app.run()
    except KeyboardInterrupt:
        _pool.close()
        _pool.join()
