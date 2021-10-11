#!/usr/bin/env bash

DB=db
TABLE=paths
touch $DB
if [[ -f $DB ]]
  then
    sqlite3 $DB "create table $TABLE(shorthand varchar(20), fullpath varchar(50), priority int, extra varchar(50));" ".exit"
    echo "create table $TABLE"
fi

if [[ $1 == "test" ]]
  then
    sqlite3 $DB "insert into $TABLE VALUES (\"shorthand1\", \"fullpath1\", 0, \"\");" ".exit"
    echo "added testline to $TABLE"
fi

