/**
 * NOTE: This class is auto generated by the akka-scala (beta) swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen
 * For any issue or feedback, please open a ticket via https://github.com/swagger-api/swagger-codegen/issues/new
 */
package io.swagger.client.api

import io.swagger.client.model.Event_type_definition_create
import io.swagger.client.core._
import io.swagger.client.core.CollectionFormats._
import io.swagger.client.core.ApiKeyLocations._

object EventtypesApi {

  /**
   * 
   * Expected answers:
   *   code 201 :  (Event type Added)
   *   code 400 :  (Invalid Request - unable to interpret request)
   *   code 401 :  (Authentication required to add a event type)
   *   code 403 :  (Access denied - not authorized to add a event type)
   *   code 409 :  (Body of the request contains conflicting, malformed, or invalid data)
   * 
   * @param body Request Body
   */
  def createEventType(body: Event_type_definition_create): ApiRequest[Unit] =
    ApiRequest[Unit](ApiMethods.POST, "https://localhost", "/event-types", "application/json")
      .withBody(body)
      .withSuccessResponse[Unit](201)
      .withErrorResponse[Unit](400)
      .withErrorResponse[Unit](401)
      .withErrorResponse[Unit](403)
      .withErrorResponse[Unit](409)
        /**
   * 
   * Expected answers:
   *   code 204 :  (Event type successfully removed)
   *   code 400 :  (Invalid Request - unable to interpret request)
   *   code 401 :  (Authentication required to delete an event type)
   *   code 403 :  (Access denied - not authorized to delete an event type)
   *   code 404 :  (Event type not found)
   * 
   * @param domain domain
   * @param entity entity
   * @param eventType event_type
   */
  def deleteEventType(domain: String, entity: String, eventType: String): ApiRequest[Unit] =
    ApiRequest[Unit](ApiMethods.DELETE, "https://localhost", "/event-types/{domain}/{entity}/{event_type}", "application/json")
      .withPathParam("domain", domain)
      .withPathParam("entity", entity)
      .withPathParam("event_type", eventType)
      .withSuccessResponse[Unit](204)
      .withErrorResponse[Unit](400)
      .withErrorResponse[Unit](401)
      .withErrorResponse[Unit](403)
      .withErrorResponse[Unit](404)
        /**
   * 
   * Expected answers:
   *   code 200 :  (Successful)
   *   code 400 :  (Invalid Request - unable to interpret request)
   * 
   * @param domain domain
   * @param entity entity
   * @param eventType event_type
   */
  def eventTypeOptions(domain: String, entity: String, eventType: String): ApiRequest[Unit] =
    ApiRequest[Unit](ApiMethods.OPTIONS, "https://localhost", "/event-types/{domain}/{entity}/{event_type}", "application/json")
      .withPathParam("domain", domain)
      .withPathParam("entity", entity)
      .withPathParam("event_type", eventType)
      .withSuccessResponse[Unit](200)
      .withErrorResponse[Unit](400)
        /**
   * 
   * Expected answers:
   *   code 200 :  (Successful)
   *   code 400 :  (Invalid Request - unable to interpret request)
   */
  def eventTypesOptions(): ApiRequest[Unit] =
    ApiRequest[Unit](ApiMethods.OPTIONS, "https://localhost", "/event-types", "application/json")
      .withSuccessResponse[Unit](200)
      .withErrorResponse[Unit](400)
        /**
   * 
   * Expected answers:
   *   code 200 :  (event-types Retrieved)
   *   code 204 :  (No results found)
   *   code 400 :  (Unknown parameter in search criteria)
   * 
   * @param domain 
   * @param entity 
   * @param &#x60;event-type&#x60; 
   */
  def getEventTypes(domain: Option[String] = None, entity: Option[String] = None, &#x60;event-type&#x60;: Option[String] = None): ApiRequest[Unit] =
    ApiRequest[Unit](ApiMethods.GET, "https://localhost", "/event-types", "application/json")
      .withQueryParam("domain", domain)
      .withQueryParam("entity", entity)
      .withQueryParam("event-type", &#x60;event-type&#x60;)
      .withSuccessResponse[Unit](200)
      .withErrorResponse[Unit](204)
      .withErrorResponse[Unit](400)
        /**
   * 
   * Expected answers:
   *   code 200 :  (Event type successfully modified)
   *   code 400 :  (Invalid Request - unable to interpret request)
   *   code 401 :  (Authentication required to modify an event type)
   *   code 403 :  (Access denied - not authorized to modify an event type)
   *   code 409 :  (Body of the request contains conflicting, malformed, or invalid data)
   * 
   * @param domain domain
   * @param entity entity
   * @param eventType event_type
   * @param body Request Body
   */
  def modifyEventType(domain: String, entity: String, eventType: String, body: Event_type_definition_create): ApiRequest[Unit] =
    ApiRequest[Unit](ApiMethods.PUT, "https://localhost", "/event-types/{domain}/{entity}/{event_type}", "application/json")
      .withBody(body)
      .withPathParam("domain", domain)
      .withPathParam("entity", entity)
      .withPathParam("event_type", eventType)
      .withSuccessResponse[Unit](200)
      .withErrorResponse[Unit](400)
      .withErrorResponse[Unit](401)
      .withErrorResponse[Unit](403)
      .withErrorResponse[Unit](409)
      

}

