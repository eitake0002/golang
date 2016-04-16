#!/usr/bin/env python
# -*- encoding:utf-8 -*-

import boto.sqs
from boto.sqs.message import Message

region = "ap-northeast-1"
conn = boto.sqs.connect_to_region(region)

# Create queue
# q = conn.create_queue("my_q")

q = conn.get_queue("test_q")
for i in range(1, 10):
  m = Message()
  m.set_body("my_message_" + str(i))
  q.write(m)

for i in range(1, 10):
  print i
  rs = q.get_messages()
  m = rs[0]
  print m.get_body()
  q.delete_message(m)
