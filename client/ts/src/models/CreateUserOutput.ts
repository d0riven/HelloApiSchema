/* tslint:disable */
/* eslint-disable */
/**
 * HelloApiSchema
 * Practice api schema
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: doriven@example.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface CreateUserOutput
 */
export interface CreateUserOutput {
    /**
     * 
     * @type {number}
     * @memberof CreateUserOutput
     */
    id: number;
    /**
     * 
     * @type {string}
     * @memberof CreateUserOutput
     */
    emailAddress: string;
    /**
     * 
     * @type {string}
     * @memberof CreateUserOutput
     */
    lastName: string;
    /**
     * 
     * @type {string}
     * @memberof CreateUserOutput
     */
    firstName: string;
    /**
     * 
     * @type {Date}
     * @memberof CreateUserOutput
     */
    birthday: Date;
    /**
     * 
     * @type {string}
     * @memberof CreateUserOutput
     */
    address: string;
}

export function CreateUserOutputFromJSON(json: any): CreateUserOutput {
    return CreateUserOutputFromJSONTyped(json, false);
}

export function CreateUserOutputFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateUserOutput {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'emailAddress': json['email_address'],
        'lastName': json['last_name'],
        'firstName': json['first_name'],
        'birthday': (new Date(json['birthday'])),
        'address': json['address'],
    };
}

export function CreateUserOutputToJSON(value?: CreateUserOutput | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'email_address': value.emailAddress,
        'last_name': value.lastName,
        'first_name': value.firstName,
        'birthday': (value.birthday.toISOString().substr(0,10)),
        'address': value.address,
    };
}


