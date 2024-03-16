# coding =utf-8
"""
    @author: quochai.kstn
"""
import time

MAX_RETRIES = 3
SLEEPING_TIME = 1

class TimeoutException(Exception):
    pass

class ExceededRetry(Exception):
    pass

def query_data():
    def __query_data():
        raise TimeoutException

    num_call = 0
    while True:
        try:       
            return __query_data()
        except TimeoutException as e:
            num_call = num_call + 1
            if num_call < MAX_RETRIES:
                time.sleep(SLEEPING_TIME)
            else:
                raise ExceededRetry from e

if __name__ == '__main__':
    query_data()