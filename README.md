# Gopher Pacman

A Pacman clone written in Go using [Ebitengine](https://ebitengine.org/).

![Pacman](./assets/sample.png)

## Features

* Keyboard input
  * Arrow up or W: go up
  * Arrow down or S: go down
  * Arrow left or A: go left
  * Arrow right or D: go right
* Sound effect
  * Play a WAV file when game starts
  * Play two WAV files when Pacman eats a dot
  * Play a WAV file when Pacman eats a power pellet
* Ghosts
  * Create the four ghosts when the game starts: Blinky, Pinky, Inky and Clyde
  * The ghosts moves randomly: every time it reaches a tile, the next tile is recalculated randomly
* Events
  * The game implements an event notification algorithm to make communication between the actors. For example: when the user press a key a event is dispatched and fires a listener function that updates the Pacman

## Roadmap

* Implement a Graph Search Algorithm to make ghost smarters
* Allow Pacman eats a ghost when it was powered up by a power pellet
* Implement the collision to check if Pacman was touched by a ghost
* Implement Pacman death effect when it is touched by a ghost
* Scoreboard showing lives left, current score e current level