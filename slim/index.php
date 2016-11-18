<?php
/**
 * Other Document
 * @version base
 */

require_once __DIR__ . '/vendor/autoload.php';

$app = new Slim\App();


/**
 * POST addPet
 * Summary: Add a new pet to the store
 * Notes: 

 */
$app->POST('/dogs', function($request, $response, $args) {
            
            
            
            
            $response->write('How about implementing addPet as a POST method ?');
            return $response;
            });



$app->run();
