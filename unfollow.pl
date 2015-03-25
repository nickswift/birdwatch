#!/usr/bin/env perl -w
use strict;
use Getopt::Std;
use Net::Twitter;

# debug mode
my $DEBUG = 0;

# rate limiting modes
my $RLM_FOLLOWERS    = 0;
my $RLM_LOOKUP_USERS = 1;

# Twitter API information
my %api_info = (consumer=>{key=>'',secret=>''},access=>{token=>'',secret => ''});

# Suggest correct useage when the user calls script incorrectly
sub usage {
	print("usage here\n");
	exit 0;
}
# Display version for those who ask
sub version {
	my $maj = 0;
	my $min = 1;
	print("Unfollow $maj.$min\n");
}
# Print snazzy title
sub title {
print("                    __      _ _                  \n");
print("                   / _|    | | |                 \n");
print("       _   _ _ __ | |_ ___ | | | _____      __   \n");
print("      | | | | '_ \\|  _/ _ \\| | |/ _ \\ \\ /\\ / /   \n");
print("      | |_| | | | | || (_) | | | (_) \\ V  V /    \n");
print("       \\__,_|_| |_|_| \\___/|_|_|\\___/ \\_/\\_/     \n");
print("                                                 \n");
print("A More Ethical Way to Auto-Unfollow People on Twitter\n\n");
}

# Some helpers
sub print_debug($) {
	my ($msg) = @_;
	print("$msg\n") if $DEBUG;
}
sub reset_time($) {
	my ($rt) = @_;
	return ($rt-time()) + 1
}

# Observe Twitter's rate-limiting standards
sub rate_limit ($$) {
	my ($mode, $nt) = @_;
	my $m        = $nt->rate_limit_status;
	my $res      = $m->{'resources'};
	my $status   = $res->{'application'}->{'/application/rate_limit_status'};
	my $st_reset = $status->{'reset'};

	if($status->{'remaining'} == 0) {
		print_debug("Reached API limit. Waiting for $st_reset seconds.\n");
		sleep(reset_time($st_reset));
	}
	
	# select return values
	my ($t, $remaining, $reset);
	$t = ($mode == $RLM_FOLLOWERS) ? 
		$res->{'followers'}->{'followers/ids'} : ($mode == $RLM_LOOKUP_USERS) ?
		$res->{'users'}->{'/users/lookup'}     : undef;
	
	# screen out invalid modes
	if(!defined($t)){
		print_debug("rate_limit mode unrecognized") ;
		exit 1;
	}
	return {
		remaining => $t->{'remaining'},
		reset     => $t->{'reset'},
	};
}

# Wait for rate-limiter
sub rl_wait($$) {
	my $limit = rate_limit(shift, shift);
	if($limit->{'remaining'} == 0) {
		my $rt = reset_time($limit->{'reset'});
		print_debug("Reached API limit. Waiting for $rt seconds.");
		sleep($rt);
	}
}

# retrieve a reference to a list of a user's followers
sub followers($$) {
	my ($uname, $nt) = @_;
	my @followers;
	my $args;
	
	# consume items from whatever driver this thing is using
	for(my $c = -1, my $r; $c; $c = $r->{'next_cursor'}) {
		rl_wait($RLM_FOLLOWERS, $nt);
		
		# construct followers_ids args and invoke that function 
		$args = $uname ? {screen_name=>$uname,cursor=>$c} : {cursor=>$c};
		$r    = $nt->followers_ids($args);
		
		# add IDs to list. TODO: prevent duplicate IDs here
		push(@followers, @{$r->{ids}})
	}
	return \@followers;
}

# retrieve usernames
sub unames_from_ids ($$) {
	my ($u, $nt) = @_;
	my @ids = keys(%$u);

	while ($#ids > 0) {
		rl_wait($RLM_LOOKUP_USERS, $nt);
		my @ssids = splice(@ids, 0, 100);
		my $users = $nt->lookup_users({user_id=>\@ssids});
		map { $u->{$_->{'id'}}->{'name'} = $u->{'screen_name'}; } @{$users};
	}
}

sub main {
	# invoke getopts
	my %opts;
	getopts('vd', \%opts) or usage();

	# act upon opt args
	version() && exit(0) if $opts{v};
	$DEBUG = 1           if $opts{d};

	# TODO: initialize API authentication information
	
	# Create Net::Twitter object
	my $nt = Net::Twitter->new(
		traits          => [qw/API::RESTv1_1/],
		ssl             => 1,
		consumer_key    => $api_info{consumer}->{'key'},
		consumer_secret => $api_info{consumer}->{'secret'},
		access_token    => $api_info{access}->{'token'},
		access_secret   => $api_info{access}->{'secret'}
	);

	# Prompt the user for commands
	title();
	for (my $done = 0; $done == 0;) {
		print("unfollow > ");

		# pull apart user commands
		my $cmd = <STDIN>;

		# unary commands
		print("leaving...\n") && last if($cmd =~ m/quit/);
		version() && next             if($cmd =~ m/version/);
		
		# binary commands
	}
}

# start script
main();