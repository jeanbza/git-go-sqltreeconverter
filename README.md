# git-go-sqltreeconverter
Convert adjacency model (id, parent_id) trees into nested set (id, left, right) trees.

[![Build Status](https://travis-ci.org/jadekler/git-go-sqltreeconverter.svg)](https://travis-ci.org/jadekler/git-go-sqltreeconverter)

--

Example running: 

1. `mysql -uroot -e 'select id, parent_id from some_db.some_table > boom.sql`
1. `go run main.go raw_node.go linked_node.go --input boom.sql --output bam.sql --regex '(\d+)\t(\d+)' --target 'some_db.some_table'` *
1. `mysql -uroot < bam.sql`

* Note: you will need to install go to run this program. See [installation documentation here](https://golang.org/doc/install)

### Flags

1. `input`: Specify the input file with `--input test_input.sql`
1. `output`: Specify the output file with `--output test_output.sql`
1. `target`: Specify the string that gets place in `alter table <target> values .. ` with `--target somedb.members`
1. `regex`: Specify the regex that parses the input file with `--regex (\w),(\w)`. Note that the first two matching subgroups must be the id and parent_id
1. `leftName`: Specify the name of the left column with `--leftName lft`
1. `rightName`: Specify the name of the right column with `--rightName rgt`

### Conditions

- A root node MUST have a parent_id of `null` or its own id

### Notes

- This app is case insensitive
- This app can handle multiple 'root' nodes. These are tree as a discrete set - roots might look like `id:1 lft:1 rgt:18`, `id:6 lft:19 rgt:20`, `id:7 lft:21 rgt:26` for instance