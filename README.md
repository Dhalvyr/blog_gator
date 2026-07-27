Introduction.
Welcome to the blog aggregator (or "gator", for short).
This program let's you aggregate multiple blogs and their entries into a single feed.

Prerequisites.
Remember that for this program is necessary to have installed both Postgres and Go.
You can install Postgres with the command:
"brew install postgresql@15" if you're using macOS
or
"sudo apt update
sudo apt install postgresql postgresql-contrib" if you're using Linux (or a Linux shell).

Installation.
To install the program you can run the line:
go install github.com/dhalvyr/blog_gator@latest
The compiled module will be located either at $GOPATH/bin or $GOBIN if set.

Configuration.
The program uses a configuration file located at home directory called ".gatorconfig.json".
Such config file stores the Username and the Postgres connection string, be sure to 
manually create the file before running the gator something like:
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}

Commands.
Once the setup's ready you can run several commands with 'gator <command>'.
Some of those commands need additional arguments.
register <name>: Registers an username for future login.
login <name>: Logs you as the given user.
reset: Resets the data in the program.
users: Prints a list of registered users.
agg <time>: Runs the aggregator continuosly, fetching the data once per interval (e.g. 1m is once a minute, 30s is once each 30 seconds).
addfeed <url>: Adds a feed.
feeds: Prints a list of the available feeds.
follow <url>: Follows the given feed.
following: Prints a list of the feeds followed by current user.
unfollow <url>: Unfollows the given feed for the current user.
browse <amount>: Prints a list of the <amount> last feeds. If no amount is given the default is 2.
