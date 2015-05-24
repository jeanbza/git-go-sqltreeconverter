# Example usage: php

Example usage using very light php to fetch data (use any language of your choice for this) and some javascript to render the tree.

## Running

(this guide assumes you're using a mac)

1. [Install xampp](https://www.apachefriends.org/index.html) - it should install at /Applications/XAMPP
1. Modify `/Applications/XAMPP/etc/php.ini` to include:
    - `mysql.default_socket=/tmp/mysql.sock`
    - `mysqli.default_socket=/tmp/mysql.sock`
1. `ln -s /Applications/XAMPP/htdocs/sqltree /path/to/git-go-sqltreeconverter`
1. Start xampp apache server (mac client is called manager-osx)
1. Load `members.sql` into a database called `tree_example` (feel free to change code to fit your db). Example:
    1. `mysql -uroot`
    1. `create database tree_example`
    1. `use tree_example`
    1. `source /path/to/this/repo/members.sql`
1. Open browser and navigate to `http://localhost/sqltree/example_usage/`