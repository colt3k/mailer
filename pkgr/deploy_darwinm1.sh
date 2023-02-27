#!/bin/bash

cp $name.bash /opt/homebrew/etc/bash_completion.d/$name
cp $name /usr/local/bin/$name
chmod +x /usr/local/bin/$name

