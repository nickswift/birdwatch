# birdwatch

Birdwatch is a command line interface for Twitter. Specifically, the design goal of this project is to provide an easy
way to curate one's Twitter feed, relationships, and history.

## Installation

Currently, installing birdwatch requires you to set up a Golang environment. In the future there will be some sort of 
binary distribution either hosted in version control, or somewhere on the internet. 

Birdwatch pulls twitter API information out of the user's environment variables. There will be support for this 
information through command line arguments in the future, but this method is the recommended one. Create all of the
relevant keys and secrets by following Twitter's instructions after creating a developer account here:
https://dev.twitter.com/

Then, add the following to your bash profile at $HOME/.bashrc ($HOME/.profile for OSX users):

	export BW_CONSUMER_KEY=<key here>
	export BW_CONSUMER_SECRET=<secret here>
	export BW_ACCESS_TOKEN=<token here>
	export BW_ACCESS_SECRET=<secret here>
