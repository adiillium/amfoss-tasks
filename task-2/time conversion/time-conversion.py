#!/bin/python3

import os
import sys
import time
#
# Complete the timeConversion function below.
#
def timeConversion(s):
    format_12 = time.strptime(s, "%I:%M:%S%p")
    return time.strftime("%H:%M:%S", format_12)
    

if __name__ == '__main__':
    f = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    f.write(result + '\n')

    f.close()
s
