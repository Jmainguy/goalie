# goalie
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/goalie)](https://goreportcard.com/badge/github.com/Jmainguy/goalie)
[![Release](https://img.shields.io/github/release/Jmainguy/goalie.svg?style=flat-square)](https://github.com/Jmainguy/goalie/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/goalie/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/goalie?branch=main)

A golang CLI to display various stats about Hockey teams and their players

## Usage
```/bin/bash
[jmainguy@jmainguy goalie]$ goalie -h
Usage of ./goalie:
  -captains
    	To print the captains of the team or not
  -roster
    	To print the team roster or not (default true)
  -team string
    	Team to lookup (default "Carolina Hurricanes")


[jmainguy@jmainguy goalie]$ goalie 
================ Forwards ================
Name, Number, Position, Age, Height, Weight, Goals, Assists, +-, Shots, Time On Ice
Jordan Staal, #11, Center, 33, 6' 4", 220, 1, 3, 2, 9, 69:02
Derek Stepan, #18, Center, 31, 5' 11", 196, 0, 0, 1, 8, 48:39
Nino Niederreiter, #21, Right Wing, 29, 6' 2", 218, 1, 0, 2, 11, 57:51
Jesper Fast, #71, Right Wing, 29, 6' 1", 191, 3, 0, 3, 5, 54:50
Vincent Trocheck, #16, Center, 28, 5' 10", 183, 2, 3, 3, 9, 70:13
Teuvo Teravainen, #86, Left Wing, 27, 5' 11", 191, 2, 4, 3, 14, 72:37
Jordan Martinook, #48, Left Wing, 29, 6' 0", 196, 1, 0, 1, 5, 41:29
Sebastian Aho, #20, Center, 24, 6' 0", 176, 3, 3, 2, 12, 75:41
Steven Lorentz, #78, Center, 25, 6' 4", 206, 0, 0, 1, 1, 37:55
Martin Necas, #88, Center, 22, 6' 2", 189, 0, 2, 0, 8, 64:15
Jesperi Kotkaniemi, #82, Center, 21, 6' 2", 201, 1, 1, 0, 8, 50:26
Andrei Svechnikov, #37, Right Wing, 21, 6' 2", 195, 4, 3, 3, 21, 65:39
Seth Jarvis, #24, Center, 19, 5' 10", 175, 0, 0, 0, 0, 
================ Defensemen ================
Name, Number, Position, Age, Height, Weight, Goals, Assists, +-, Shots, Time On Ice
Eric Gelinas, #23, Defenseman, 30, 6' 4", 215, 0, 0, 0, 0, 
Ian Cole, #28, Defenseman, 32, 6' 1", 225, 0, 0, 4, 3, 56:12
Brendan Smith, #7, Defenseman, 32, 6' 2", 200, 0, 0, 0, 0, 
Brady Skjei, #76, Defenseman, 27, 6' 3", 210, 0, 1, 1, 5, 81:41
Jaccob Slavin, #74, Defenseman, 27, 6' 3", 207, 0, 3, 3, 4, 90:28
Brett Pesce, #22, Defenseman, 26, 6' 3", 206, 0, 3, 0, 4, 89:52
Tony DeAngelo, #77, Defenseman, 26, 5' 11", 180, 0, 4, 4, 10, 68:06
Ethan Bear, #25, Defenseman, 24, 5' 11", 197, 0, 1, 2, 2, 73:30
================ Goalies ================
Name, Number, Position, Age, Height, Weight, Wins, Losses, SavePercentage, EvenStrengthSavePercentage, PowerPlaySavePercentage, ShortHandedSavePercentage, Time On Ice
Frederik Andersen, #31, Goalie, 32, 6' 4", 238, 4, 0, 0.944000, 94.845361, 88.888889, 100.000000, 239:59
Antti Raanta, #32, Goalie, 32, 6' 0", 195, 0, 0, 0.000000, 0.000000, 0.000000, 0.000000, 


[jmainguy@jmainguy goalie]$ goalie -captains
Jordan Staal : Captain, #11, Center, Age: 33, Height 6' 4", Weight 220
Jordan Martinook : Alternate Captain, #48, Left Wing, Age: 29, Height 6' 0", Weight 196
Jaccob Slavin : Alternate Captain, #74, Defenseman, Age: 27, Height 6' 3", Weight 207

```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/goalie/releases)

## Install

### Homebrew

```/bin/bash
brew install Jmainguy/tap/goalie
```

## Build
```/bin/bash
export GO111MODULE=on
go build
```
