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
 * @interface CreateUserInput
 */
export interface CreateUserInput {
    /**
     * 
     * @type {string}
     * @memberof CreateUserInput
     */
    emailAddress: string;
    /**
     * 
     * @type {string}
     * @memberof CreateUserInput
     */
    lastName: string;
    /**
     * 
     * @type {string}
     * @memberof CreateUserInput
     */
    firstName: string;
    /**
     * 
     * @type {Date}
     * @memberof CreateUserInput
     */
    birthday: Date;
    /**
     * 
     * @type {string}
     * @memberof CreateUserInput
     */
    address: string;
}

export function CreateUserInputFromJSON(json: any): CreateUserInput {
    return CreateUserInputFromJSONTyped(json, false);
}

export function CreateUserInputFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateUserInput {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'emailAddress': json['email_address'],
        'lastName': json['last_name'],
        'firstName': json['first_name'],
        'birthday': (new Date(json['birthday'])),
        'address': json['address'],
    };
}

export function CreateUserInputToJSON(value?: CreateUserInput | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'email_address': value.emailAddress,
        'last_name': value.lastName,
        'first_name': value.firstName,
        'birthday': (value.birthday.toISOString().substr(0,10)),
        'address': value.address,
    };
}


