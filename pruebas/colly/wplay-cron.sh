#!/bin/sh

#------------------------------------------------------------------------------------------
# This is a shell file developed to be run evety 5 mins with crontab. This is part of the 
# Betting Arbitrage project developed by Carlos A. Gorricho using Golang (gorricho2009@gmail.com)
#
#------------------------------------------------------------------------------------------

# Reminf crontab of the path 

PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin

# Run the Golang executable for Wplay betting house. In the future the same executables for other betting houses go in this section

./wplay-foot-live

# Sync data file with Github, so that it's accesible from other computers

git add FOOT-live-data.csv
git commit -m "update FOOT data"
git push