#!/usr/bin/env bash

DB=db
touch $DB
sqlite3 $DB "create table paths(shorthand varchar(20), fullpath varchar(50), priority int);" ".exit"

