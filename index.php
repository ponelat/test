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


/**
 * POST addPet_1
 * Summary: Add a new pet to the store
 * Notes: 

 */
$app->POST('/pet', function($request, $response, $args) {
            
            
            
            
            $response->write('How about implementing addPet_1 as a POST method ?');
            return $response;
            });


/**
 * DELETE deletePet
 * Summary: Deletes a pet
 * Notes: 
 * Output-Formats: [application/xml, application/json]
 */
$app->DELETE('/pet/{petId}', function($request, $response, $args) {
            $headers = $request->getHeaders();
            
            
            
            $response->write('How about implementing deletePet as a DELETE method ?');
            return $response;
            });


/**
 * GET getPetById
 * Summary: Find pet by ID
 * Notes: Returns a single pet
 * Output-Formats: [application/xml, application/json]
 */
$app->GET('/pet/{petId}', function($request, $response, $args) {
            
            
            
            
            $response->write('How about implementing getPetById as a GET method ?');
            return $response;
            });


/**
 * PUT updatePet
 * Summary: Update an existing pet
 * Notes: 
 * Output-Formats: [application/xml, application/json]
 */
$app->PUT('/pet', function($request, $response, $args) {
            
            
            
            $body = $request->getParsedBody();
            $response->write('How about implementing updatePet as a PUT method ?');
            return $response;
            });


/**
 * POST updatePetWithForm
 * Summary: Updates a pet in the store with form data
 * Notes: 
 * Output-Formats: [application/xml, application/json]
 */
$app->POST('/pet/{petId}', function($request, $response, $args) {
            
            
            $name = $args['name'];    $status = $args['status'];    
            
            $response->write('How about implementing updatePetWithForm as a POST method ?');
            return $response;
            });



$app->run();
