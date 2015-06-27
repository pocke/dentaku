Dentaku
===========

yacc learning.

Installation
---------------

```sh
go get -d github.com/pocke/dentaku
cd $GOPATH/src/github.com/pocke/dentaku/parser
make
cd ..
go install
```

Usage
------

```sh
echo '1+1' | dentaku  # => 2
```

Syntax
-------

### Addition

```ruby
1+3   # => 4
2+0   # => 2
```

### Subtraction

```ruby
10-3  # => 7
5-10  # => -5
```

### Multiplication

```ruby
10*5  # => 50
```

### Division

```ruby
10/2  # => 5
14/6  # => 7/3
```

### Parenthesis

```ruby
3*5+2*6     # => 27
3*(5+2)*6   # => 126
```
