# goping
# Go based network monitor with output file logging when the RTT threshold is higher that initially defined.
Use as this:


go run goping.go google.com 443 -1 2s 15

-1 - Count forever
2s - delay between checks, you can define values in following format 2s, 5s, 1m etc.
15 - threshold of millisecond. If the value of tested RTT is higher that this  threshold value, will output an alert on stdout and also will log in file output.txt in the following format:

1 Oct 22 04:47PM        36      google.com      443
