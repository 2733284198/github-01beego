Benchmarks
There are 3 benchmark types, trying to simulate real-life usage for small, medium and large JSON payloads. For each metric, the lower value is better. Time/op is in nanoseconds. Values better than standard encoding/json marked as bold text. Benchmarks run on standard Linode 1024 box.

Compared libraries:


https://golang.org/pkg/encoding/json
https://github.com/Jeffail/gabs
https://github.com/a8m/djson
https://github.com/bitly/go-simplejson
https://github.com/antonholmquist/jason
https://github.com/mreiferson/go-ujson
https://github.com/ugorji/go/codec
https://github.com/pquerna/ffjson
https://github.com/mailru/easyjson
https://github.com/buger/jsonparser

Small payload
Each test processes 190 bytes of http log as a JSON record. It should read multiple fields. https://github.com/buger/jsonparser/blob/master/benchmark/benchmark_small_payload_test.go

Library	time/op	bytes/op	allocs/op
encoding/json struct	7879	880	18
encoding/json interface{}	8946	1521	38
Jeffail/gabs	10053	1649	46
bitly/go-simplejson	10128	2241	36
antonholmquist/jason	27152	7237	101
github.com/ugorji/go/codec	8806	2176	31
mreiferson/go-ujson	7008	1409	37
a8m/djson	3862	1249	30
pquerna/ffjson	3769	624	15
mailru/easyjson	2002	192	9
buger/jsonparser	1367	0	0
buger/jsonparser (EachKey API)	809	0	0