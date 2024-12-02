/**
 * Interface representing a default response object.
 * @interface ResponseObjectDefaultInterface
 * @property {string} status - The status of the response.
 * @property {number} httpStatus - The HTTP status code of the response.
 * @property {string} message - The message associated with the response.
 * @property {any[any]} payload - The payload of the response.
 */

interface ResponseObjectDefaultInterface {
  status: string;
  httpStatus: number;
  message: string;
  payload: any[any];
}

/**
 * Builder class for creating a default response object.
 * @class ResponseObjectDefaultBuilder
 * @implements {ResponseObjectDefaultInterface}
 */
export class ResponseObjectDefaultBuilder {
  private readonly responseObjectDefault: ResponseObjectDefaultInterface;

  /**
   * Creates an instance of ResponseObjectDefaultBuilder.
   * Initializes the default response object with default values.
   */
  constructor() {
    this.responseObjectDefault = {
      status: 'success',
      httpStatus: 200,
      message: 'success',
      payload: [],
    };
  }

  /**
   * Sets the status of the response object.
   * @param {string} status - The status to set.
   * @returns {ResponseObjectDefaultBuilder} The builder instance.
   */
  public setStatus(status: string): ResponseObjectDefaultBuilder {
    this.responseObjectDefault.status = status;
    return this;
  }

  /**
   * Sets the HTTP status code of the response object.
   * @param {number} httpStatus - The HTTP status code to set.
   * @returns {ResponseObjectDefaultBuilder} The builder instance.
   */
  public setHttpStatus(httpStatus: number): ResponseObjectDefaultBuilder {
    this.responseObjectDefault.httpStatus = httpStatus;
    return this;
  }

  /**
   * Sets the message of the response object.
   * @param {string} message - The message to set.
   * @returns {ResponseObjectDefaultBuilder} The builder instance.
   */
  public setMessage(message: string): ResponseObjectDefaultBuilder {
    this.responseObjectDefault.message = message;
    return this;
  }

  /**
   * Sets the payload of the response object.
   * @param {any[any]} payload - The payload to set.
   * @returns {ResponseObjectDefaultBuilder} The builder instance.
   */
  public setPayload(payload: any[any]): ResponseObjectDefaultBuilder {
    this.responseObjectDefault.payload = [...payload];
    return this;
  }

  /**
   * Builds and returns the response object.
   * @returns {ResponseObjectDefaultInterface} The built response object.
   */
  public build(): ResponseObjectDefaultInterface {
    return this.responseObjectDefault;
  }
}
