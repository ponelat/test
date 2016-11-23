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
# NOTE: This class is auto generated by the swagger code generator program. 
# Do not edit the class manually.
# Ref: https://github.com/swagger-api/swagger-codegen
#
package WWW::SwaggerClient::StoreApi;

require 5.6.0;
use strict;
use warnings;
use utf8; 
use Exporter;
use Carp qw( croak );
use Log::Any qw($log);

use WWW::SwaggerClient::ApiClient;
use WWW::SwaggerClient::Configuration;

use base "Class::Data::Inheritable";

__PACKAGE__->mk_classdata('method_documentation' => {});

sub new {
    my $class   = shift;
    my (%self) = (
        'api_client' => WWW::SwaggerClient::ApiClient->instance,
        @_
    );

    #my $self = {
    #    #api_client => $options->{api_client}
    #    api_client => $default_api_client
    #}; 

    bless \%self, $class;

}


#
# delete_order
#
# Delete purchase order by ID
# 
# @param int $order_id ID of the order that needs to be deleted (required)
{
    my $params = {
    'order_id' => {
        data_type => 'int',
        description => 'ID of the order that needs to be deleted',
        required => '1',
    },
    };
    __PACKAGE__->method_documentation->{ 'delete_order' } = { 
    	summary => 'Delete purchase order by ID',
        params => $params,
        returns => undef,
        };
}
# @return void
#
sub delete_order {
    my ($self, %args) = @_;

    # verify the required parameter 'order_id' is set
    unless (exists $args{'order_id'}) {
      croak("Missing the required parameter 'order_id' when calling delete_order");
    }

    # parse inputs
    my $_resource_path = '/store/order/{orderId}';
    $_resource_path =~ s/{format}/json/; # default format to json

    my $_method = 'DELETE';
    my $query_params = {};
    my $header_params = {};
    my $form_params = {};

    # 'Accept' and 'Content-Type' header
    my $_header_accept = $self->{api_client}->select_header_accept('application/xml', 'application/json');
    if ($_header_accept) {
        $header_params->{'Accept'} = $_header_accept;
    }
    $header_params->{'Content-Type'} = $self->{api_client}->select_header_content_type();

    # path params
    if ( exists $args{'order_id'}) {
        my $_base_variable = "{" . "orderId" . "}";
        my $_base_value = $self->{api_client}->to_path_value($args{'order_id'});
        $_resource_path =~ s/$_base_variable/$_base_value/g;
    }

    my $_body_data;
    # authentication setting, if any
    my $auth_settings = [qw()];

    # make the API Call
    $self->{api_client}->call_api($_resource_path, $_method,
                                           $query_params, $form_params,
                                           $header_params, $_body_data, $auth_settings);
    return;
}

#
# get_inventory
#
# Returns pet inventories by status
# 
{
    my $params = {
    };
    __PACKAGE__->method_documentation->{ 'get_inventory' } = { 
    	summary => 'Returns pet inventories by status',
        params => $params,
        returns => 'HASH[string,int]',
        };
}
# @return HASH[string,int]
#
sub get_inventory {
    my ($self, %args) = @_;

    # parse inputs
    my $_resource_path = '/store/inventory';
    $_resource_path =~ s/{format}/json/; # default format to json

    my $_method = 'GET';
    my $query_params = {};
    my $header_params = {};
    my $form_params = {};

    # 'Accept' and 'Content-Type' header
    my $_header_accept = $self->{api_client}->select_header_accept('application/json');
    if ($_header_accept) {
        $header_params->{'Accept'} = $_header_accept;
    }
    $header_params->{'Content-Type'} = $self->{api_client}->select_header_content_type();

    my $_body_data;
    # authentication setting, if any
    my $auth_settings = [qw(api_key )];

    # make the API Call
    my $response = $self->{api_client}->call_api($_resource_path, $_method,
                                           $query_params, $form_params,
                                           $header_params, $_body_data, $auth_settings);
    if (!$response) {
        return;
    }
    my $_response_object = $self->{api_client}->deserialize('HASH[string,int]', $response);
    return $_response_object;
}

#
# get_order_by_id
#
# Find purchase order by ID
# 
# @param int $order_id ID of pet that needs to be fetched (required)
{
    my $params = {
    'order_id' => {
        data_type => 'int',
        description => 'ID of pet that needs to be fetched',
        required => '1',
    },
    };
    __PACKAGE__->method_documentation->{ 'get_order_by_id' } = { 
    	summary => 'Find purchase order by ID',
        params => $params,
        returns => 'Order',
        };
}
# @return Order
#
sub get_order_by_id {
    my ($self, %args) = @_;

    # verify the required parameter 'order_id' is set
    unless (exists $args{'order_id'}) {
      croak("Missing the required parameter 'order_id' when calling get_order_by_id");
    }

    # parse inputs
    my $_resource_path = '/store/order/{orderId}';
    $_resource_path =~ s/{format}/json/; # default format to json

    my $_method = 'GET';
    my $query_params = {};
    my $header_params = {};
    my $form_params = {};

    # 'Accept' and 'Content-Type' header
    my $_header_accept = $self->{api_client}->select_header_accept('application/xml', 'application/json');
    if ($_header_accept) {
        $header_params->{'Accept'} = $_header_accept;
    }
    $header_params->{'Content-Type'} = $self->{api_client}->select_header_content_type();

    # path params
    if ( exists $args{'order_id'}) {
        my $_base_variable = "{" . "orderId" . "}";
        my $_base_value = $self->{api_client}->to_path_value($args{'order_id'});
        $_resource_path =~ s/$_base_variable/$_base_value/g;
    }

    my $_body_data;
    # authentication setting, if any
    my $auth_settings = [qw()];

    # make the API Call
    my $response = $self->{api_client}->call_api($_resource_path, $_method,
                                           $query_params, $form_params,
                                           $header_params, $_body_data, $auth_settings);
    if (!$response) {
        return;
    }
    my $_response_object = $self->{api_client}->deserialize('Order', $response);
    return $_response_object;
}

#
# place_order
#
# Place an order for a pet
# 
# @param Order $body order placed for purchasing the pet (required)
{
    my $params = {
    'body' => {
        data_type => 'Order',
        description => 'order placed for purchasing the pet',
        required => '1',
    },
    };
    __PACKAGE__->method_documentation->{ 'place_order' } = { 
    	summary => 'Place an order for a pet',
        params => $params,
        returns => 'Order',
        };
}
# @return Order
#
sub place_order {
    my ($self, %args) = @_;

    # verify the required parameter 'body' is set
    unless (exists $args{'body'}) {
      croak("Missing the required parameter 'body' when calling place_order");
    }

    # parse inputs
    my $_resource_path = '/store/order';
    $_resource_path =~ s/{format}/json/; # default format to json

    my $_method = 'POST';
    my $query_params = {};
    my $header_params = {};
    my $form_params = {};

    # 'Accept' and 'Content-Type' header
    my $_header_accept = $self->{api_client}->select_header_accept('application/xml', 'application/json');
    if ($_header_accept) {
        $header_params->{'Accept'} = $_header_accept;
    }
    $header_params->{'Content-Type'} = $self->{api_client}->select_header_content_type();

    my $_body_data;
    # body params
    if ( exists $args{'body'}) {
        $_body_data = $args{'body'};
    }

    # authentication setting, if any
    my $auth_settings = [qw()];

    # make the API Call
    my $response = $self->{api_client}->call_api($_resource_path, $_method,
                                           $query_params, $form_params,
                                           $header_params, $_body_data, $auth_settings);
    if (!$response) {
        return;
    }
    my $_response_object = $self->{api_client}->deserialize('Order', $response);
    return $_response_object;
}

1;
