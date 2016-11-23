=begin comment

Swagger Petstore

This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.

OpenAPI spec version: 1.0.0
Contact: apiteam@swagger.io
Generated by: https://github.com/swagger-api/swagger-codegen.git

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

=end comment

=cut

#
# NOTE: This class is auto generated by Swagger Codegen
# Please update the test cases below to test the API endpoints.
# Ref: https://github.com/swagger-api/swagger-codegen
#
use Test::More tests => 1; #TODO update number of test cases
use Test::Exception;

use lib 'lib';
use strict;
use warnings;

use_ok('WWW::SwaggerClient::UserApi');

my $api = WWW::SwaggerClient::UserApi->new();
isa_ok($api, 'WWW::SwaggerClient::UserApi');

#
# create_user test
#
{
    my $body = undef; # replace NULL with a proper value
    my $result = $api->create_user(body => $body);
}

#
# create_users_with_array_input test
#
{
    my $body = undef; # replace NULL with a proper value
    my $result = $api->create_users_with_array_input(body => $body);
}

#
# create_users_with_list_input test
#
{
    my $body = undef; # replace NULL with a proper value
    my $result = $api->create_users_with_list_input(body => $body);
}

#
# delete_user test
#
{
    my $username = undef; # replace NULL with a proper value
    my $result = $api->delete_user(username => $username);
}

#
# get_user_by_name test
#
{
    my $username = undef; # replace NULL with a proper value
    my $result = $api->get_user_by_name(username => $username);
}

#
# login_user test
#
{
    my $username = undef; # replace NULL with a proper value
    my $password = undef; # replace NULL with a proper value
    my $result = $api->login_user(username => $username, password => $password);
}

#
# logout_user test
#
{
    my $result = $api->logout_user();
}

#
# update_user test
#
{
    my $username = undef; # replace NULL with a proper value
    my $body = undef; # replace NULL with a proper value
    my $result = $api->update_user(username => $username, body => $body);
}


1;
