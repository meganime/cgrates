#!/bin/bash
# set -e
go clean --cache
results=()
results+=($?)
if [ "$#" -ne 0 ]; then
# to run for a single agent add `-run=*mysql` as argument
# ./calls_test.sh -run=TestFreeswitchCalls
echo "go test github.com/cgrates/cgrates/general_tests -tags=call $@"
go test github.com/cgrates/cgrates/general_tests -tags=call $@ -v
results+=($?)
else
echo "go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestFreeswitchCalls"
go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestFreeswitchCalls -v
results+=($?)
echo "go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestKamailioCalls"
go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestKamailioCalls -v
results+=($?)
echo "go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestOpensipsCalls"
go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestOpensipsCalls -v
results+=($?)
echo "go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestAsteriskCalls"
go test github.com/cgrates/cgrates/general_tests -tags=call -run=TestAsteriskCalls -v
results+=($?)
fi

pass=1
for val in ${results[@]}; do
   (( pass=$pass||$val))
done
exit $pass