from flask import Flask
from flask import request
import time
from multiprocessing import Pool

app = Flask(__name__)
_pool = None


def vdf(x, t):
    # do your expensive time consuming process
    print('yay vdf')
    return x


@app.route('/vdf', methods=['POST'])
def compute():
    input = request.form.get('input')
    time = request.form.get('time')
    print(input, time)
    # perform validation, invoke executable, retreive output, return
    if validate(input, time):
         f = _pool.apply_async(vdf, (input,time))
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
