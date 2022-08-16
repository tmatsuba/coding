#!/bin/bash
dir=$(head -n 1 ./contest)
echo -e CONTEST: $dir
echo -n QUESTION?
read q
rm -rf test
echo http://$dir.contest.atcoder.jp/tasks/$dir"_"$q
oj download http://$dir.contest.atcoder.jp/tasks/$dir"_"$q
cp test/sample-1.in input
