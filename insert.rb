
require 'mysql'

start_time = Time.now

connection = Mysql::connect("localhost", "root", "", "dev_test")
connection.query("set character set utf8")
qry = "insert into test_table (name) values('ruby test')"

10000.times do |item|
  rs = connection.query(qry)
end
connection.close

p "ruby #{Time.now - start_time}s"

