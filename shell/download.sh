#!/bin/bash
dir=$(head -n 1 ./contest)
echo -e CONTEST: $dir
echo -n QUESTION?
read q
rm -rf test

if [ ${dir} == "tessoku_book" ]; then
    echo https://atcoder.jp/contests/tessoku-book/tasks/tessoku_book_$q
    oj download https://atcoder.jp/contests/tessoku-book/tasks/tessoku_book_$q
else
    echo http://$dir.contest.atcoder.jp/tasks/$dir"_"$q
    oj download http://$dir.contest.atcoder.jp/tasks/$dir"_"$q
fi

cp test/sample-1.in input
