# GJ

## Dependencies


*GJ* is a cli-utility that helps you jump faster between different directories. GJ can learn from either looking at your shells history file, by explicitly mapping shorthands to paths or by running it in discovery mode.

## Usage

'gj SH': Let's gj jump to the shorthand it has learned for SH

'gj SH PATH': Adds PATH to gjs db with the shorthand SH 

'gj --discover [PATH]': Starts a discovery run with [PATH] as the root, adding all directories it finds

'gj --list': Prints out all PATH = SH pairs

For gj to be able to run, it needs to create a db it saves its mappings to. Per default this file is located in /home/$USER/.config/gj.db. If the db is not at this location, its path has to be given by the parameter **--db**

## Development
**gj** is developed via an MVP and TDD methodology in mind. This means, that a minimal viable product will firstly defined by setting up testcases for the planned interfaces. Then those interfaces will be implemented and tested. Every iteration will aim at providing additional functionality until all of the above listed commands work as expected.

### Iteration 0: Base
* Provide CRUD-capabilities for a db-file
* Add SH PATH per CLI
* **gj** is able to jump given a SH
