#!/usr/bin/perl

use strict;
use warnings;
use ZeroMQ qw/:all/;
use Data::Dumper;

my $ctxt = ZeroMQ::Context->new;
my $socket = $ctxt->socket(ZMQ_REQ);
$socket->connect( "tcp://127.0.0.1:5555" );

$socket->send($ARGV[0]);
my $msg = $socket->recv;
print $msg->data;

print "\n";
