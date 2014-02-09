# verify_path [![Build Status](https://drone.io/github.com/crahles/verify_path/status.png)](https://drone.io/github.com/crahles/verify_path/latest)

verify_path is a simple util to check if there are paths in your input which dont exist anymore (old entries in crontab).

## Usage

From testfile
```code
cat test_input.txt | verify_path

Path /var/www/xxxmail/releases/20130114112251 does not exists!
Path /var/www/yyy_to_qv/releases/20131022103053 does not exists!
```

Or from your cat/crontab/else stdin via pipe
```code
cat file | verify_path
crontab -l -u xxx | verify_path
```

## TODOs

- remove fixed regexp and parse it from commandline via flags
- only read from stdin if no file given (add option for commandline)
- provide packages via bintray

## Contributing

Contributions are very welcome. Whether it's an issue or even a pull request. For pull requests you can use the following flow:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

I'd also really appreciate spec only pull requests or bug reports with a failing spec/minimal example as this makes fixing it a lot easier.

Thanks in advance for all contributions of any kind!

## License

Released under the MIT license

Copyright (c) 2014 Christoph Rahles (crahles)
