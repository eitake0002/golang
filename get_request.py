import threading
import time
import datetime
import urllib2
import sys

class TestThread(threading.Thread):

    """docstring for TestThread"""

    def __init__(self, url_list):
        super(TestThread, self).__init__()
        self.url_list = url_list

    def run(self):
        start = time.time()
        print " === start sub thread (sub class) === "
        for i in self.url_list:
            try:
                response = urllib2.urlopen(i)
                html = response.read()
                print i
                print html[0:20]
            except:
                print "Error : " + i 
        print " === end sub thread (sub class) === "
        elapsed_time = time.time() - start
        print ("elapsed_time:{0}".format(elapsed_time)) + "[sec]"

#def hoge(n, t):
#    print " === start sub thread (method) === "
#    for i in range(n):
#        time.sleep(t)
#        print "[%s] sub thread (method) : " % threading.currentThread().getName() + str(datetime.datetime.today())
#    print " === end sub thread (method) === "

if __name__ == '__main__':
    start = time.time()
    url_list = []
    for line in open("url_list.csv"):
      url_list.append(line)

    th_cl = TestThread(url_list)
    th_cl.start()

    print " === start main thread (main) === "
    print " == end main thread === "
